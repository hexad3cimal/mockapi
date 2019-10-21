package exceptions

import "time"

type Exception struct {
	TypeOfException string
	Message         string
	Timestamp       time.Time

}

func  Error(exception *Exception) string { return exception.TypeOfException + " -------- "+ exception.Message + " ------- " + exception.Timestamp.String() }