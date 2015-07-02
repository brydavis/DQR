package main

// _ "github.com/mattn/go-sqlite3"

func main() {
	report := NewReport(
		"Data Quality Report",
		"Shelter and Housing Programs",
		"data/data.json",
		6,
		2015,
		struct {
			October,
			January,
			April,
			July int
		}{
			October: 29,
			January: 28,
			April:   29,
			July:    29,
		},
	)
	report.ListenAndServe(8080)
}
