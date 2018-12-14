package atrium

func NewAtriumClient(apiKey, clientId string) *AtriumClientService {
  config := NewConfiguration()
  config.DefaultHeader["MX-Client-ID"] = clientId
  config.DefaultHeader["MX-API-Key"] = apiKey
  return newAPIClient(config).AtriumClient
}

func NewAtriumClientWithConfiguration(apiKey, clientId string, config *Configuration) *AtriumClientService {
  config.DefaultHeader["MX-Client-ID"] = clientId
  config.DefaultHeader["MX-API-Key"] = apiKey
  return newAPIClient(config).AtriumClient
}
