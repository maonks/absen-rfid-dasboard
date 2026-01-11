package utils

import "time"

func MonthList() []struct {
	Value int
	Label string
} {
	return []struct {
		Value int
		Label string
	}{
		{1, "Januari"},
		{2, "Februari"},
		{3, "Maret"},
		{4, "April"},
		{5, "Mei"},
		{6, "Juni"},
		{7, "Juli"},
		{8, "Agustus"},
		{9, "September"},
		{10, "Oktober"},
		{11, "November"},
		{12, "Desember"},
	}
}

func YearList() []int {
	now := time.Now().Year()
	var years []int
	for i := now - 4; i <= now+1; i++ {
		years = append(years, i)
	}
	return years
}
