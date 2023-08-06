package models

type Response struct {
	T_trend     float32 `json:"t_trend" db:"t_trend"`
	G_temp      float32 `json:"g_temp" db:"g_temp"`
	Response    string  `json:"response" db:"response"`
	Oopt        string  `json:"oopt" db:"oopt"`
	CoastDist   float32 `json:"coast_dist" db:"coast_dist"`
	CoastRate   float32 `json:"coast_rate" db:"coast_rate"`
	DepositName string  `json:"deposit_name" db:"name"`
	YedDist     float32 `json:"yed_dist" db:"yed_dist"`
	PermDistr   string  `json:"perm_distr" db:"perm_distr"`
}
