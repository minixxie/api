package entity

type Order struct {
    Id int64               `json:"id"`
    Title string           `json:"title"`
    OrderCategoryId int64  `json:"orderCategoryId"`
    CreatedAtMicroseconds int64        `json:"createdAtMicroseconds"`
    CreatedAtISO string    `json:"createdAtISO"`
	UpdatedAtMicroseconds int64        `json:"updatedAtMicroseconds"`
    UpdatedAtISO string    `json:"updatedAtISO"`
}
