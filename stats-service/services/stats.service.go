package services

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"stats-service/utils"
)

var STAT_KEYS = []string{"MemTotal", "MemFree", "MemAvailable", "SwapTotal", "SwapCached", "SwapFree"}

type StatsService interface {
	ReadMemory() (map[string]int, error)
}

type Stats struct{}

func (s *Stats) ReadMemory() (map[string]int, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	stats := map[string]int{}

	for scanner.Scan() {
		key, value := utils.ParseMemInfoLine(scanner.Text())
		if slices.Contains(STAT_KEYS, key) {
			stats[key] = value
		}
	}
	fmt.Println(stats)
	return stats, nil
}
