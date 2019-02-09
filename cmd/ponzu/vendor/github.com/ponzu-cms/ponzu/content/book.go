package content

import (
	"fmt"
	"net/http"

	"github.com/bosssauce/reference"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Book struct {
	item.Item

	Title  string   `json:"title"`
	Author string   `json:"author"`
	Pages  int      `json:"pages"`
	Year   int      `json:"year"`
	Photos []string `json:"photos"`
}

// MarshalEditor writes a buffer of html to edit a Book within the CMS
// and implements editor.Editable
func (b *Book) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(b,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Book field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", b, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: reference.Select("Author", b, map[string]string{
				"label": "Author",
			},
				"Author",
				`{{ .name }} `,
			),
		},
		editor.Field{
			View: editor.Input("Pages", b, map[string]string{
				"label":       "Pages",
				"type":        "text",
				"placeholder": "Enter the Pages here",
			}),
		},
		editor.Field{
			View: editor.Input("Year", b, map[string]string{
				"label":       "Year",
				"type":        "text",
				"placeholder": "Enter the Year here",
			}),
		},
		editor.Field{
			View: editor.FileRepeater("Photos", b, map[string]string{
				"label":       "Photos",
				"placeholder": "Upload the Photos here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Book editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Book"] = func() interface{} { return new(Book) }
}

// String defines how a Book is printed. Update it using more descriptive
// fields from the Book struct type
func (b *Book) String() string {
	return fmt.Sprintf("%s (%d)", b.Title, b.Year)
}

func (b *Book) Push() []string {
	return []string{
		"author",
		"photos",
	}
}

func (b *Book) Create(res http.ResponseWriter, req *http.Request) error {
	return nil
}

func (b *Book) AutoApprove(res http.ResponseWriter, req *http.Request) error {
	return nil
}
