package service

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

// shouldIDeployToday returns the funny message that will be displayed in the page.
// This uses the `should I deploy today` API.
func ShouldIDeployToday() string {
	ctx := context.Background()
	baseURL := "https://shouldideploy.today/api?tz=America/Bogota"
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL, nil)
	if err != nil {
		return ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var data apiResponse
	if err := json.Unmarshal(b, &data); err != nil {
		return ""
	}

	return data.Message
}

type apiResponse struct {
	Timezone string    `json:"timezone"`
	Date     time.Time `json:"date"`
	Message  string    `json:"message"`
}
