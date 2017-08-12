package entity

type User struct {
    Id int64                     `json:"id"`
    Phone string                 `json:"phone"`
    Password string              `json:"password"`

    CreatedAtMicroseconds int64  `json:"createdAtMicroseconds"`
    CreatedAtISO string          `json:"createdAtISO"`
	UpdatedAtMicroseconds int64  `json:"updatedAtMicroseconds"`
    UpdatedAtISO string          `json:"updatedAtISO"`
}
