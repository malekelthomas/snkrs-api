package rest

import (
	"snkrs/pkg/create"
	"snkrs/pkg/get"
)

type Services struct {
	Get    get.Service
	Create create.Service
}
