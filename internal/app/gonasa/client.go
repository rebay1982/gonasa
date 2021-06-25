package gonasa

const marsAPIBaseURL = "https://api.nasa.gov/mars-photos/api/v1/rovers"

type MarsAPIClient struct {
	apiKey string
}

func NewMarsAPIClient(apiKey string) *MarsAPIClient {
	return &MarsAPIClient{apiKey: apiKey}
}
