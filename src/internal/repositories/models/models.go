package models

import "github.com/D-D-EINA-Calendar/CalendarServer/src/internal/core/domain"

type AuxAvaiableHours struct {
	Kind      int    `json:"tipo"`
	Subject   string `json:"name"`
	Remaining int    `json:"disponibles"`
	Max       int    `json:"totales"`
}

func AuxToReal(auxv AuxAvaiableHours) domain.AvailableHours {
	return domain.AvailableHours{Subject: domain.Subject{Kind: auxv.Kind, Name: auxv.Subject}, Remaining: auxv.Remaining, Max: auxv.Max}
}