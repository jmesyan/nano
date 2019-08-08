package models

import (
	"hewolf/app/models/structure"
)

func AddUserOnline(data *structure.YlyOnline) {
	dbr.Insert(data)
}
