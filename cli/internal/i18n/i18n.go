// Package i18n loads the CLI's translations from embedded TOML message files and
// resolves them for the language detected from the environment.
package i18n

import (
	"embed"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	goi18n "github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed locales/*.toml
var localesFS embed.FS

// Translator resolves message IDs to localized strings.
type Translator struct {
	localizer *goi18n.Localizer
}

// New builds a Translator for the language detected from the environment,
// falling back to English.
func New() *Translator {
	bundle := goi18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	for _, name := range []string{"locales/active.en.toml", "locales/active.pt-BR.toml"} {
		// Files are embedded and valid; a parse error would be a build-time bug.
		_, _ = bundle.LoadMessageFileFS(localesFS, name)
	}

	return &Translator{localizer: goi18n.NewLocalizer(bundle, detectLanguages()...)}
}

// T returns the localized string for id, falling back to the id itself.
func (t *Translator) T(id string) string {
	message, err := t.localizer.Localize(&goi18n.LocalizeConfig{MessageID: id})
	if err != nil {
		return id
	}
	return message
}

// detectLanguages reads language preferences from common environment variables,
// normalizing values like "pt_BR.UTF-8" to "pt-BR".
func detectLanguages() []string {
	for _, key := range []string{"FOCUZEN_LANG", "LC_ALL", "LC_MESSAGES", "LANG"} {
		if value := os.Getenv(key); value != "" {
			return []string{normalize(value)}
		}
	}
	return nil
}

func normalize(value string) string {
	if i := strings.IndexAny(value, ".@"); i >= 0 {
		value = value[:i]
	}
	return strings.ReplaceAll(value, "_", "-")
}
