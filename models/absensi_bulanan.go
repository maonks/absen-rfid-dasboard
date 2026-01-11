package models

type HariCell struct {
	Masuk  string
	Pulang string
	Status string // OK | LATE
}

type AbsensiRow struct {
	Nama string
	Hari map[int]*HariCell // key = tanggal (1..31)
}

type AbsensiBulananVM struct {
	Days []int
	Rows []AbsensiRow
}
