package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Author struct {
	item.Item

	Name  string `json:"name"`
	Photo string `json:"photo"`
	Bio   string `json:"bio"`
}

// MarshalEditor writes a buffer of html to edit a Author within the CMS
// and implements editor.Editable
func (a *Author) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(a,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Author field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Name", a, map[string]string{
				"label":       "Name",
				"type":        "text",
				"placeholder": "Enter the Name here",
			}),
		},
		editor.Field{
			View: editor.File("Photo", a, map[string]string{
				"label":       "Photo",
				"placeholder": "Upload the Photo here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Bio", a, map[string]string{
				"label":       "Bio",
				"placeholder": "Enter the Bio here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Author editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Author"] = func() interface{} { return new(Author) }
}

// String defines how a Author is printed. Update it using more descriptive
// fields from the Author struct type
func (a *Author) String() string {
	return a.Name
}

