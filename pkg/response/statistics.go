package response

type StatisticsResponse struct {
	Totals Totals `json:"totals"`
}
type Totals struct {
	Downloads int64 `json:"downloads"`
	Packages  int   `json:"packages"`
	Versions  int   `json:"versions"`
}
