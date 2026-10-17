package main

type Stats struct {
	Count         int
	Sum, Min, Max int64
}

func Calc(nums []int64) Stats {
	var stats Stats

	for i := 1; i < len(nums); i++ {
		stats.addDelta(nums[i] - nums[i-1])
	}

	return stats
}

func (s *Stats) addDelta(delta int64) {
	if s.Count == 0 || delta < s.Min {
		s.Min = delta
	}
	if s.Count == 0 || delta > s.Max {
		s.Max = delta
	}

	s.Count++
	s.Sum += delta
}
