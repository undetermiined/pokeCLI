package pokeapi

func (c *Client) GetLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	return getAndCache[LocationAreasResponse](c, url)
}
