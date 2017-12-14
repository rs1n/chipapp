package models

import "github.com/sknv/pgup/orm/record"

// Base application model.
type Base struct {
	record.Timestampable `db:",inline"`
}
