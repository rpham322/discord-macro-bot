package bot



import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"


	"github.com/bwmarrin/discordgo"

)


const nutritionURL = "https://trackapi.nutritionix.com/v2/natural/nutrients"




type NutritionixData struct {
	Foods []Food `json:"foods"`
}


type Food struct {
	// Basic
	FoodName   string  `json:"food_name"`
	BrandName  *string `json:"brand_name,omitempty"`

	// Serving
	ServingQty         float64 `json:"serving_qty"`
	ServingUnit        string  `json:"serving_unit"`
	ServingWeightGrams float64 `json:"serving_weight_grams"`

	// Macros
	Calories float64 `json:"nf_calories"`
	Protein  float64 `json:"nf_protein"`
	Carbs    float64 `json:"nf_total_carbohydrate"`
	Fat      float64 `json:"nf_total_fat"`
	Fiber    float64 `json:"nf_dietary_fiber"`
	Sugars   float64 `json:"nf_sugars"`
	Sodium      float64 `json:"nf_sodium"`
	Cholesterol float64 `json:"nf_cholesterol"`





}




func getNutrition(message string) *discordgo.MessageSend {
	// match food query using regex
	r := regexp.MustCompile(`(?i)^!macro\s+(.+)$`)
	m := r.FindStringSubmatch(strings.TrimSpace(message))

	if len(m) < 2 {
		return &discordgo.MessageSend{Content: "Please provide a food item after the command. Usage: `!macro <food>`"}
	}
	query := m[1]

	// build JSON body
	body := []byte(fmt.Sprintf(`{"query":"%s"}`, query))

	// create new HTTP request
	req, _ := http.NewRequest("POST", nutritionURL, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-app-id", os.Getenv("NUTRITIONIX_APP_ID"))
	req.Header.Set("x-app-key", os.Getenv("NUTRITIONIX_TOKEN"))

	// create new HTTP client & set timeout
	client := &http.Client{Timeout: 5 * time.Second}

	// query Nutritionix API
	resp, err := client.Do(req)
	if err != nil {
		return &discordgo.MessageSend{Content: "API error."}
	}
	defer resp.Body.Close()

	// open HTTP response body
	raw, _ := io.ReadAll(resp.Body)



	// http check status
	// if not 2xx, show API error payload to help debug
	if resp.StatusCode/100 != 2 {
		snippet := string(raw)
		if len(snippet) > 300 { snippet = snippet[:300] + "..." }
		return &discordgo.MessageSend{
			Content: fmt.Sprintf("Nutritionix HTTP %d\n```%s```", resp.StatusCode, snippet),
		}
	}



	// convert JSON
	var data struct{ Foods []Food `json:"foods"` }
	if err := json.Unmarshal(raw, &data); err != nil || len(data.Foods) == 0 {
		return &discordgo.MessageSend{Content: fmt.Sprintf("No results for **%s**", query)}
	}

	// pull out desired nutrition info
	f := data.Foods[0]
	name := strings.Title(f.FoodName)
	brand := ""
	if f.BrandName != nil && *f.BrandName != "" {
		brand = " (" + *f.BrandName + ")"
	}

	serving := fmt.Sprintf("%.0f %s (%.0fg)", f.ServingQty, f.ServingUnit, f.ServingWeightGrams)
	calories := fmt.Sprintf("%.0f kcal", f.Calories)
	protein := fmt.Sprintf("%.1f g", f.Protein)
	carbs := fmt.Sprintf("%.1f g", f.Carbs)
	fat := fmt.Sprintf("%.1f g", f.Fat)
	fiber := fmt.Sprintf("%.1f g", f.Fiber)
	sugars := fmt.Sprintf("%.1f g", f.Sugars)

	// build Discord embed response
	embed := &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Type:        discordgo.EmbedTypeRich,
				Title:       "Nutrition — " + name + brand,
				Description: "Serving: " + serving,
				Fields: []*discordgo.MessageEmbedField{
					{Name: "Calories", Value: calories, Inline: true},
					{Name: "Protein", Value: protein, Inline: true},
					{Name: "Carbs", Value: carbs, Inline: true},
					{Name: "Fat", Value: fat, Inline: true},
					{Name: "Fiber", Value: fiber, Inline: true},
					{Name: "Sugars", Value: sugars, Inline: true},
				},
			},
		},
	}

	return embed
}


