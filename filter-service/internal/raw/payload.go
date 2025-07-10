package raw

type RawRequest struct {
	Valid string `json:"valid" validate:"required"`
	Data  string `json:"data" validate:"required"`
}
