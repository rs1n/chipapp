package models

import "github.com/sknv/mng/odm/document"

// Base model.
type Base struct {
	document.Timestampable `bson:",inline"`
}
