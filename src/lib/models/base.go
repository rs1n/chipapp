package models

import "github.com/sknv/chip/mng/odm/document"

// Base application model.
type Base struct {
	document.Timestampable `bson:",inline"`
}
