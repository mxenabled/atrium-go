package atrium

func NewAPIClient(apiKey, clientId string) *APIClient {
  config := NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = clientId
  config.DefaultHeader["MX-API-Key"] = apiKey
  return newAPIClient(config)
}

func NewAPIClientWithConfiguration(apiKey, clientId string, config *Configuration) *APIClient {
  config.DefaultHeader["MX-Client-ID"] = clientId
  config.DefaultHeader["MX-API-Key"] = apiKey
  return newAPIClient(config)
}
