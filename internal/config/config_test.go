package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetConfigPath(t *testing.T) {
	// TODO: create a temp dir to mock the user's home directory.
	tests := []struct {
		name          string
		wantErr       bool
		wantFileExist bool
	}{}
}
