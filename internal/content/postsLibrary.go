package content

import (
	"time"

	"github.com/nickfoden/web-log/internal/models"
)

var PostsLibrary = map[string]models.Post{
	"reduce-the-cost-of-owning-software": {
		Title:          "Reduce the Cost of Owning Software",
		Content:        "",
		ContentPreview: `"We don't need to get better at authoring software, our challenge, the financial challenge to any business is not I need to write more code faster, It's I need to reduce the cost of owning my software"`,
		CreatedAt:      time.Date(2025, 7, 3, 0, 0, 0, 0, time.UTC),
		Slug:           "reduce-the-cost-of-owning-software",
	},
	"react-native-requires-current-ruby": {
		Title:          "Ruby for React Native",
		Content:        "",
		ContentPreview: "RN setup issues? Ensure Ruby is current",
		CreatedAt:      time.Date(2025, 1, 16, 0, 0, 0, 0, time.UTC),
		Slug:           "react-native-requires-current-ruby",
	},
	"no-bugs": {
		Title:          "No Bugs Here",
		Content:        "",
		ContentPreview: "C'mon just use a proper tsconfig",
		CreatedAt:      time.Date(2025, 1, 15, 0, 0, 0, 0, time.UTC),
		Slug:           "no-bugs",
	},
	"1": {
		Title:          "New Web Log Alert",
		Content:        "",
		ContentPreview: "A \"simple\" web log built with Go and html + htmx.",
		CreatedAt:      time.Date(2025, 1, 3, 0, 0, 0, 0, time.UTC),
		Slug:           "1",
	},
}
