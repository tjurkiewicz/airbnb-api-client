package client

import (
  "github.com/dghubble/sling"
)

const BASE_URL string = "https://www.airbnb.com/api/v1/listings/"

type AirBNB struct {
  ApiKey string
}

func (api* AirBNB) ReadListing(id string) (listingResponse ListingResponse, errorResponse ErrorResponse, err error) {
  params := &Params{Key: api.ApiKey}
  _, err = sling.New().Base(BASE_URL).Path(id).QueryStruct(params).Receive(&listingResponse, &errorResponse)
  return
}

