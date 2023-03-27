package theme

import (
	"fmt"
	"strings"

	"github.com/qor/qor"
	"github.com/qor/qor/resource"
)

type Themes map[string]*Theme

func (t Themes) FindManyHandler(results interface{}, context *qor.Context) error {
	switch result := results.(type) {
	// the framework calls this handler with an *int interface to get
	// the records count, for pagination purposes
	case *int:
		count := len(t)
		*result = count

	// used by the UI
	case *[]Config:
		keyword := context.Request.URL.Query().Get("keyword")
		for _, theme := range t {
			if keyword == "" || keyword != "" && strings.Contains(theme.Name, keyword) {
				*result = append(*result, Config{
					ID:      theme.Name,
					Name:    theme.Name,
					Author:  theme.Author,
					Version: theme.Version,
					Path:    theme.Path,
				})
			}
		}

	// used by the API
	case *[]*Config:
		for _, theme := range t {
			*result = append(*result, &Config{
				ID:      theme.Name,
				Name:    theme.Name,
				Author:  theme.Author,
				Version: theme.Version,
				Path:    theme.Path,
			})
		}

	default:
		return fmt.Errorf("unknown interface: %#v", results)
	}

	return nil
}

func (t Themes) FindOneHandler(i interface{}, mv *resource.MetaValues, ctx *qor.Context) error {
	switch result := i.(type) {
	case *Config:
		if ctx.ResourceID != "" {
			if theme, ok := t[ctx.ResourceID]; ok {
				*result = Config{
					ID:      theme.Name,
					Name:    theme.Name,
					Author:  theme.Author,
					Version: theme.Version,
					Path:    theme.Path,
				}
				return nil
			}

			return fmt.Errorf("themeFindOneHandler: not found: %s", ctx.ResourceID)
		}
	default:
		return fmt.Errorf("themeFindOneHandler: unknown interface: %#v", i)
	}
	return nil
}
