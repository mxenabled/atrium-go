package atrium

func AtriumClient(apiKey, clientId string) *APIClient {
	config := NewConfiguration()
	config.DefaultHeader["MX-Client-ID"] = clientId
	config.DefaultHeader["MX-API-Key"] = apiKey
	return newAPIClient(config)
}

func AtriumClientWithConfiguration(apiKey, clientId string, config *Configuration) *APIClient {
	config.DefaultHeader["MX-Client-ID"] = clientId
	config.DefaultHeader["MX-API-Key"] = apiKey
	return newAPIClient(config)
}
