package main

import "ride-sharing/shared/types"

type previewTripRequest struct {
	UserID      string           `jsonL"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
