package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator(t *testing.T) {
	t.Run("resolves portuguese when selected", func(t *testing.T) {
		t.Setenv("FOCUZEN_LANG", "pt-BR")
		tr := New()
		assert.Equal(t, "Sem áudio", tr.T("noAudio"))
	})

	t.Run("falls back to english", func(t *testing.T) {
		t.Setenv("FOCUZEN_LANG", "en")
		tr := New()
		assert.Equal(t, "No audio", tr.T("noAudio"))
	})

	t.Run("normalizes locale values like pt_BR.UTF-8", func(t *testing.T) {
		t.Setenv("FOCUZEN_LANG", "pt_BR.UTF-8")
		tr := New()
		assert.Equal(t, "Sem áudio", tr.T("noAudio"))
	})

	t.Run("returns the id for unknown messages", func(t *testing.T) {
		t.Setenv("FOCUZEN_LANG", "en")
		tr := New()
		assert.Equal(t, "missing.id", tr.T("missing.id"))
	})
}
