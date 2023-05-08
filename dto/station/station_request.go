package stationdto

type StationRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
	City string `json:"city" form:"city" validate:"required"`
}
