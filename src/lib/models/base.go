package models

import "github.com/skkv/chip/mng/odm/document"

// Base application model.
type Base struct {
	document.Timestampable `bson:",inline"`
}
