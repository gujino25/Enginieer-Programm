package domain

import (
	"errors"
	"testing"
)

func TestNewSystem(t *testing.T) {
	tests := []struct {
		name    string
		medium  string
		wantErr error
	}{
		{
			name:    "Валидная среда - воздух",
			medium:  "air",
			wantErr: nil,
		},
		{
			name:    "Невалидная среда",
			medium:  "абобус",
			wantErr: ErrMediumInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewSystem("project-1", "П1", tt.medium, "офис")
			if tt.wantErr == nil {
				if err != nil {
					t.Errorf("неожиданная ошибка: %v", err)
				}
				return
			}
			if !errors.Is(err, tt.wantErr) {
				t.Errorf("err = %v, жидалось %v", err, tt.wantErr)
			}
		})
	}

}
