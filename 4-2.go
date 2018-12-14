package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type log struct {
	time    time.Time
	message string
}

const (
	eventAwake  = -1
	eventAsleep = 1
)

type event struct {
	time               time.Time
	guardID, eventType int
}

func (l *log) String() string {
	return fmt.Sprintf("%s: %s\n", l.time, l.message)
}

func main() {
	const timeFormat = "[2006-01-02 15:04]"
	timeLen := len(timeFormat)
	prefixLen := len("Guard") // == len("wakes") == len("falls")

	scanner := bufio.NewScanner(os.Stdin)
	var logs []*log
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		t, err := time.Parse(timeFormat, line[0:timeLen])
		if err != nil {
			panic(err)
		}
		logs = append(logs, &log{
			time:    t,
			message: strings.TrimSpace(line[timeLen:]),
		})
	}

	sort.Slice(logs, func(i, j int) bool {
		return logs[i].time.Before(logs[j].time)
	})

	var events []*event
	var id int
	for _, log := range logs {
		switch log.message[0:prefixLen] {
		case "Guard":
			fmt.Sscanf(log.message, "Guard #%d", &id)
		case "falls":
			events = append(events, &event{
				time:      log.time,
				guardID:   id,
				eventType: eventAsleep,
			})
		case "wakes":
			events = append(events, &event{
				time:      log.time,
				guardID:   id,
				eventType: eventAwake,
			})
		default:
			panic("unexpected event")
		}
	}

	sort.Slice(events, func(i, j int) bool {
		mi, mj := events[i].time.Minute(), events[j].time.Minute()
		if mi == mj {
			return events[i].eventType > events[j].eventType
		}
		return mi < mj
	})

	maxID, maxMinute, maxCount := -1, -1, 0
	counts := make(map[int]int)
	for _, e := range events {
		counts[e.guardID] += e.eventType
		if counts[e.guardID] > maxCount {
			maxID, maxMinute, maxCount = e.guardID, e.time.Minute(), counts[e.guardID]
		}
	}

	fmt.Println(maxID, maxMinute, maxCount)

	fmt.Println(maxID * maxMinute)
}
