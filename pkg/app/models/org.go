package models

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

//go:generate goqueryset -in org.go

// gen:qs
type Org struct {
	gorm.Model

	Name        string
	DisplayName string

	Provider               string
	ProviderID             int
	ProviderPersonalUserID int

	Settings []byte
}

func (o *Org) IsFake() bool {
	return o.ProviderPersonalUserID != 0
}

func (o *Org) UnmarshalSettings(v interface{}) error {
	return errors.Wrapf(json.Unmarshal(o.Settings, v), "failed to unmarshal settings for org(%d)", o.ID)
}

func (o *Org) MarshalSettings(v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal %#v as settings", v)
	}
	o.Settings = data
	return nil
}
