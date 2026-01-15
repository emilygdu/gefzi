package routes

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Intervall festlegen (umrechnung der Uhrzeiten in Int weil Berechnung dann weniger Fehleranfällig ist)
type interval struct {
	Start int
	End   int
}

//Intervall Logik (sorieren der Zeitslots und zusammenziehen wenn sich diese Überlappen)

func mergeIntervals(in []interval) []interval {
	if len(in) == 0 {
		return nil
	}

	sort.Slice(in, func(i, j int) bool {
		return in[i].Start < in[j].Start
	})

	out := []interval{in[0]}
	for _, cur := range in[1:] {
		last := &out[len(out)-1]
		if cur.Start <= last.End { // overlap oder angrenzend
			if cur.End > last.End {
				last.End = cur.End
			}
		} else {
			out = append(out, cur)
		}
	}
	return out
}

// Arbeitszeit minus blockierte Zeitslots -> berechnung der freien Zeitslots
func subtractIntervals(work interval, blocked []interval) []interval {
	free := []interval{}
	cursor := work.Start

	for _, b := range blocked {
		if b.Start > cursor {
			free = append(free, interval{Start: cursor, End: b.Start})
		}
		if b.End > cursor {
			cursor = b.End
		}
	}

	if cursor < work.End {
		free = append(free, interval{Start: cursor, End: work.End})
	}

	return free
}

//Umrechnen der Zeiten von Int in String

func parseHHMM(s string) (int, error) {
	parts := strings.Split(s, ":")
	if len(parts) != 2 {
		return 0, strconv.ErrSyntax
	}

	h, err := strconv.Atoi(parts[0])
	if err != nil || h < 0 || h > 23 {
		return 0, strconv.ErrSyntax
	}

	m, err := strconv.Atoi(parts[1])
	if err != nil || m < 0 || m > 59 {
		return 0, strconv.ErrSyntax
	}

	return h*60 + m, nil
}

func formatHHMM(minutes int) string {
	h := minutes / 60
	m := minutes % 60
	return fmt.Sprintf("%02d:%02d", h, m)
}
