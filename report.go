package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Card struct {
	Agency           string   `json: "agency"`
	ProgramName      string   `json: "programName"`
	ProgramType      string   `json: "programType"`
	TargetPops       []string `json: "targetPops"`
	NumberBeds       int      `json: "numberBeds"`
	NumberUnits      int      `json: "numberUnits"`
	ServedClients    int      `json: "servedClients"`
	NewClients       int      `json: "newClients"`
	ExitedClients    int      `json: "exitedClients"`
	ServedHouseholds int      `json: "servedHouseholds"`

	ClientsPITOctober int `json:"clientsPITOctober"`
	ClientsPITJanuary int `json:"clientsPITJanuary"`
	ClientsPITApril   int `json:"clientsPITApril"`
	ClientsPITJuly    int `json:"clientsPITJuly"`

	HouseholdsPITOctober int `json:"householdsPITOctober"`
	HouseholdsPITJanuary int `json:"householdsPITJanuary"`
	HouseholdsPITApril   int `json:"householdsPITApril"`
	HouseholdsPITJuly    int `json:"householdsPITJuly"`

	FullName         int `json: "fullName"`
	SocialSecurity   int `json: "socialSecurity"`
	HeadHousehold    int `json: "headHousehold"`
	BirthDate        int `json: "birthDate"`
	Race             int `json: "race"`
	Ethnicity        int `json: "ethnicity"`
	Gender           int `json: "gender"`
	VeteranStatus    int `json: "veteranStatus"`
	DisabilityStatus int `json: "disabilityStatus"`
	SubstanceAbuse   int `json: "substanceAbuse"`
	PriorLiving      int `json: "priorLiving"`
	ClientZip        int `json: "clientZip"`
	ChronicityStatus int `json: "chronicityStatus"`
	QualityScore     float64
}

type PITs struct {
	October,
	January,
	April,
	July int
}

type Report struct {
	Title     string
	SubTitle  string
	Month     int
	Year      int
	Source    string
	Continuum Card
	Cards     []Card
	PITs
}

func NewReport(title, subtitle, source string, month, year int, pits PITs) Report {
	return Report{
		Title:    title,
		SubTitle: subtitle,
		Source:   source,
		Month:    month,
		Year:     year,
		PITs: PITs{
			October: pits.October,
			January: pits.January,
			April:   pits.April,
			July:    pits.July,
		},
	}
}

func (report *Report) ListenAndServe(port int) {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) { RootHandler(res, req, *report) })
	http.HandleFunc("/static/", StaticHandler)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		panic(err)
	}
}

func (report *Report) Run() Report {
	data, err := ioutil.ReadFile(report.Source)
	checkError(err)

	err = json.Unmarshal(data, &report.Cards)
	checkError(err)

	for i, card := range report.Cards {
		total := float64(card.FullName + card.SocialSecurity + card.HeadHousehold + card.BirthDate +
			card.Race + card.Ethnicity + card.Gender + card.VeteranStatus + card.DisabilityStatus +
			card.SubstanceAbuse + card.PriorLiving + card.ClientZip + card.ChronicityStatus)

		report.Cards[i].QualityScore = ((total / 13.0) / float64(card.ServedClients)) * 100

		report.Continuum.NumberBeds += card.NumberBeds
		report.Continuum.NumberUnits += card.NumberUnits
		report.Continuum.ServedClients += card.ServedClients
		report.Continuum.NewClients += card.NewClients
		report.Continuum.ExitedClients += card.ExitedClients
		report.Continuum.ServedHouseholds += card.ServedHouseholds
		report.Continuum.FullName += card.FullName
		report.Continuum.SocialSecurity += card.SocialSecurity
		report.Continuum.HeadHousehold += card.HeadHousehold
		report.Continuum.BirthDate += card.BirthDate
		report.Continuum.Race += card.Race
		report.Continuum.Ethnicity += card.Ethnicity
		report.Continuum.Gender += card.Gender
		report.Continuum.VeteranStatus += card.VeteranStatus
		report.Continuum.DisabilityStatus += card.DisabilityStatus
		report.Continuum.SubstanceAbuse += card.SubstanceAbuse
		report.Continuum.PriorLiving += card.PriorLiving
		report.Continuum.ClientZip += card.ClientZip
		report.Continuum.ChronicityStatus += card.ChronicityStatus

	}

	total := float64(report.Continuum.FullName + report.Continuum.SocialSecurity + report.Continuum.HeadHousehold + report.Continuum.BirthDate +
		report.Continuum.Race + report.Continuum.Ethnicity + report.Continuum.Gender + report.Continuum.VeteranStatus + report.Continuum.DisabilityStatus +
		report.Continuum.SubstanceAbuse + report.Continuum.PriorLiving + report.Continuum.ClientZip + report.Continuum.ChronicityStatus)

	report.Continuum.QualityScore = ((total / 13.0) / float64(report.Continuum.ServedClients)) * 100

	return *report
}
