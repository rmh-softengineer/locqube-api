package model

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Property struct {
	PropertyID     string   `json:"property_id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Price          int      `json:"price"`
	PropertyType   string   `json:"property_type"`
	Bedrooms       int      `json:"bedrooms"`
	Baths          int      `json:"baths"`
	HalfBaths      int      `json:"half_baths"`
	LivingAreaSqft int      `json:"living_area_sqft"`
	LotSizeSqft    int      `json:"lot_size_sqft"`
	Address        Address  `json:"address"`
	Images         []string `json:"images"`
	ListingURL     string   `json:"listing_url"`
	VideoURL       string   `json:"video_url"`
}
