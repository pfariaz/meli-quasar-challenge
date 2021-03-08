package schemas

type SatelliteRequestSchema struct {
	Name     string   `json:"name" binding:"required"`
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type ProcessFullMessageRequestSchema struct {
	Sattelites []SatelliteRequestSchema `json:"satellites" binding:"required"`
}

type ProcessSplitMessageRequestSchema struct {
	Distance float64  `json:"distance" binding:"required"`
	Message  []string `json:"message" binding:"required"`
}

type PositionResponseSchema struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type ShipResponseSchema struct {
	Position PositionResponseSchema `json:"position"`
	Message  string                 `json:"message"`
}
