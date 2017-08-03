package client

type Params struct {
  Key string `url:"key,omitempty"`
}

type ErrorResponse struct {
  ErrorCode    int    `json:"error_code"`
  Error        string `json:"error"`
  ErrorMessage string `json:"error_message"`
}

type ListingResponse struct {
  Listing struct {
    Id             int64     `json:"id"`
    City           string    `json:"city"`
    UserId         int64     `json:"user_id"`
    Latitude       float64   `json:"lat"`
    Longitude      float64   `json:"lng"`
    Bathrooms      float64   `json:"bathrooms"`
    Bedrooms       float64   `json:"bedrooms"`
    Beds           float64   `json:"beds"`
    PersonCapacity int       `json:"person_capacity"`
    CountryCode    string    `json:"country_code"`
    Amenities      []string  `json:"amenities"`
  } `json:"listing"`
}


