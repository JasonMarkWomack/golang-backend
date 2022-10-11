package player

import (
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestGetPlayer(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetPlayer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetPlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSavePlayer(t *testing.T) {
	type args struct {
		c *fiber.Ctx
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SavePlayer(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("SavePlayer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
