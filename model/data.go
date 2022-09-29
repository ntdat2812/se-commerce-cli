package model

type Response struct {
	Result []Data `json:"data"`
}

type Data struct {
	Name             string        `json:"name"`
	ShortDescription string        `json:"description"`
	ThumbnailUrl     string        `json:"thumbnail_url"`
	OptionColor      []OptionColor `json:"option_color"`
}

type OptionColor struct {
	Price        float64 `json:"price"`
	OriginPrice  float64 `json:"origin_price"`
	ThumbnailUrl string  `json:"thumbnail_url"`
	DisplayName  string  `json:"display_name"`
}
