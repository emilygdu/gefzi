package routes

import (
	"net/http"
	"strconv"
	"time"

	"gefzi/internal/database"
	"gefzi/internal/models"

	"github.com/gin-gonic/gin"
)

func RegisterAvailabilityRoutes(c *gin.Engine) {
	c.GET("/availability/month", GetAvailabilityMonth)
}

func GetAvailabilityMonth(c *gin.Context) {
	//Parameter
	gcID, _ := strconv.Atoi(c.Query("group_calendar_id"))
	year, _ := strconv.Atoi(c.Query("year"))
	monthInt, _ := strconv.Atoi(c.Query("month"))

	//Parameter Check, wenn einer fehlt -> 400
	if gcID == 0 || year == 0 || monthInt < 1 || monthInt > 12 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "use ?group_calendar_id=1&year=2026&month=11",
		})
		return
	}

	//Kalender laden
	var gc models.GroupCalendar
	if err := database.DB.First(&gc, gcID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "group calendar not found",
		})
		return
	}
	workStart, err1 := parseHHMM(gc.WorkStart)
	workEnd, err2 := parseHHMM(gc.WorkEnd)
	if err1 != nil || err2 != nil || workEnd <= workStart {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid work hours",
		})
		return
	}

	//Monatsspanne
	start := time.Date(year, time.Month(monthInt), 1, 0, 0, 0, 0, time.Local)
	end := start.AddDate(0, 1, -1)
	startStr := start.Format("2006-01-02")
	endStr := end.Format("2006-01-02")

	// Alle Events für den Monat holen
	var events []models.Event
	if err := database.DB.Where(
		"group_calendar_id = ? AND date >= ? AND date <= ?",
		gcID, startStr, endStr,
	).Find(&events).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	//Events nach Tagen gruppieren
	byDate := map[string][]models.Event{}
	for _, e := range events {
		byDate[e.Date] = append(byDate[e.Date], e)
	}

	//für jeden Tag freie Slots berechnen
	days := []gin.H{}

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		dateStr := d.Format("2006-01-02")

		// Wochenende blockiert?
		if gc.WeekendBlocked && (d.Weekday() == time.Saturday || d.Weekday() == time.Sunday) {
			days = append(days, gin.H{"date": dateStr, "free": []gin.H{}})
			continue
		}

		// Events -> blockierte Zeitslots (Intervalle)
		blocked := []interval{}
		for _, e := range byDate[dateStr] {
			s, errS := parseHHMM(e.StartTime)
			en, errE := parseHHMM(e.EndTime)
			if errS != nil || errE != nil || en <= s {
				continue
			}
			//Arbeitszeit
			if en <= workStart || s >= workEnd {
				continue
			}
			if s < workStart {
				s = workStart
			}
			if en > workEnd {
				en = workEnd
			}
			blocked = append(blocked, interval{Start: s, End: en})
		}

		// merge overlaps
		blocked = mergeIntervals(blocked)

		//Arbeitszeit minus blocked
		free := subtractIntervals(interval{Start: workStart, End: workEnd}, blocked)

		// in JSON umwandeln
		freeJSON := []gin.H{}
		for _, f := range free {
			freeJSON = append(freeJSON, gin.H{
				"start": formatHHMM(f.Start),
				"end":   formatHHMM(f.End),
			})
		}

		days = append(days, gin.H{"date": dateStr, "free": freeJSON})
	}

	c.JSON(http.StatusOK, gin.H{
		"group_calendar_id": gcID,
		"year":              year,
		"month":             monthInt,
		"days":              days,
	})
}
