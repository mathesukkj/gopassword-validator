package tests

import (
	"reflect"
	"testing"

	"github.com/mathesukkj/gopassword-validator/services"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		want     []string
	}{
		{"Abcdefg1@", nil},
		{
			"abc",
			[]string{
				"the password contains less than 8 characters",
				"the password doesnt contain an uppercase character",
				"the password doesnt contain a number",
				"the password doesnt contain a special character",
			},
		},
		{"abcDEF@GH", []string{"the password doesnt contain a number"}},
		{"ABCDEFG1@", []string{"the password doesnt contain a lowercase character"}},
		{"abcdefg1@", []string{"the password doesnt contain an uppercase character"}},
		{"Abcdefghij1", []string{"the password doesnt contain a special character"}},
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			if got := services.ValidatePassword(tt.password); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidatePassword(%q) = %v, want %v", tt.password, got, tt.want)
			}
		})
	}
}
