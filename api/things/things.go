package things

type ThingsClient struct {
	lineAccessToken string
}

func NewThingsClient(token string) *ThingsClient {
	return &ThingsClient{token}
}

func (things *ThingsClient) GetLineAccessToken() string {
	return things.lineAccessToken
}

func (tc *ThingsClient) GetDevice(deviceID string) (*Device, error) {
	return nil, nil
}

func (tc *ThingsClient) GetUserLinkedDevice(deviceID, userID string) (*UserDevice, error) {
	return nil, nil
}

func (tc *ThingsClient) GetProductLinkedDevices(deviceID, userID string) (*UserDevices, error) {
	return nil, nil
}

func (tc *ThingsClient) CreateTrialProduct() (*TrialProduct, error) {
	return nil, nil
}

func (tc *ThingsClient) DeleteTrialProduct(productID string) error {
	return nil
}

func (tc *ThingsClient) GetTrialProducts() (*TrialProducts, error) {
	return nil, nil
}
