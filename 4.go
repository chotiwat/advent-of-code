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
	time      time.Time
	eventType int
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

	eventsMap := map[int][]*event{}
	var id int
	for _, log := range logs {
		switch log.message[0:prefixLen] {
		case "Guard":
			fmt.Sscanf(log.message, "Guard #%d", &id)
		case "falls":
			eventsMap[id] = append(eventsMap[id], &event{
				time:      log.time,
				eventType: eventAsleep,
			})
		case "wakes":
			eventsMap[id] = append(eventsMap[id], &event{
				time:      log.time,
				eventType: eventAwake,
			})
		default:
			panic("unexpected event")
		}
	}
	maxAsleep := 0.0
	var maxID int
	for id, events := range eventsMap {
		asleep := 0.0
		for i := 0; i < len(events); i += 2 {
			asleep += (events[i+1].time.Sub(events[i].time)).Minutes()
			if asleep > maxAsleep {
				maxID, maxAsleep = id, asleep
			}
		}
	}

	fmt.Println(maxID, maxAsleep)
	maxEvents := eventsMap[maxID]
	sort.Slice(maxEvents, func(i, j int) bool {
		mi, mj := maxEvents[i].time.Minute(), maxEvents[j].time.Minute()
		if mi == mj {
			return maxEvents[i].eventType > maxEvents[j].eventType
		}
		return mi < mj
	})

	maxMinute, maxCount := -1, 0
	count := 0
	for _, e := range maxEvents {
		count += e.eventType
		if count > maxCount {
			maxMinute, maxCount = e.time.Minute(), count
		}
	}

	fmt.Println(maxMinute, maxCount)

	fmt.Println(maxID * maxMinute)
}
