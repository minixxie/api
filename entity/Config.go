package entity

type Config struct {
	JwtSecret string             `json:"jwtSecret"`
	Dbs struct {
		Main string
		User string
	}
}
