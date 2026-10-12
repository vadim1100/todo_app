package tasks_service

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateTitle(t *testing.T) {
	tests := []struct {
		name string
		title string
		needErr bool
	}{
		{"ok", "Сделать домашнее задание", false},
		{"empty", "", true},
		{"very long", strings.Repeat("a", 250), true},
		{"spaces", "    ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateTitle(tt.title)
			if tt.needErr{
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}