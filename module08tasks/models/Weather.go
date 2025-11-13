package models

import (
	"time"
)

type Weather struct {
	Date        time.Time
	Temperature int
	Description string
}
