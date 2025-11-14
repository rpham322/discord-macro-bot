package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	nutritionixURL = "https://trackapi.nutritionix.com/v2/natural/nutrients"
	requestTimeout = 5 * time.Second
)

// Client handles Nutritionix API interactions
type Client struct {
	appID  string
	appKey string
	client *http.Client
}

// NewClient creates a new Nutritionix API client
func NewClient(appID, appKey string) *Client {
	return &Client{
		appID:  appID,
		appKey: appKey,
		client: &http.Client{Timeout: requestTimeout},
	}
}

// Food represents nutrition data for a food item
type Food struct {
	FoodName           string  `json:"food_name"`
	BrandName          *string `json:"brand_name,omitempty"`
	ServingQty         float64 `json:"serving_qty"`
	ServingUnit        string  `json:"serving_unit"`
	ServingWeightGrams float64 `json:"serving_weight_grams"`
	Calories           float64 `json:"nf_calories"`
	Protein            float64 `json:"nf_protein"`
	Carbs              float64 `json:"nf_total_carbohydrate"`
	Fat                float64 `json:"nf_total_fat"`
	Fiber              float64 `json:"nf_dietary_fiber"`
	Sugars             float64 `json:"nf_sugars"`
	Sodium             float64 `json:"nf_sodium"`
	Cholesterol        float64 `json:"nf_cholesterol"`
}

// NutritionResponse represents the API response
type NutritionResponse struct {
	Foods []Food `json:"foods"`
}

// GetNutrition queries the Nutritionix API for food nutrition data
func (c *Client) GetNutrition(query string) (*Food, error) {
	// Build request body
	body := map[string]string{"query": query}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", nutritionixURL, bytes.NewReader(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("Failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-app-id", c.appID)
	req.Header.Set("x-app-key", c.appKey)

	// Execute request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response: %w", err)
	}

	// Check HTTP status
	if resp.StatusCode/100 != 2 {
		snippet := string(raw)
		if len(snippet) > 300 {
			snippet = snippet[:300] + "..."
		}
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, snippet)
	}

	// Parse JSON response
	var data NutritionResponse
	if err := json.Unmarshal(raw, &data); err != nil {
		return nil, fmt.Errorf("Failed to parse response: %w", err)
	}

	if len(data.Foods) == 0 {
		return nil, fmt.Errorf("No results found for query: %s", query)
	}

	return &data.Foods[0], nil
}

