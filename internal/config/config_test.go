package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// NOTE: the primeagen rights tests like these
func TestLoadConfigCreatesDefault(t *testing.T) {
	// Fake HOME so we don’t use the real system directory
	tmpHome := t.TempDir()
	t.Setenv("HOME", tmpHome)

	cfg, err := LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	// Expect empty default config
	assert.Equal(t, []string{}, cfg.CollectionPath)

	// Ensure file was actually created
	cfgFile, err := getConfigFile()
	require.NoError(t, err)
	_, err = os.Stat(cfgFile)
	require.NoError(t, err)
}

func TestSaveConfigAndReload(t *testing.T) {
	tmpHome := t.TempDir()
	t.Setenv("HOME", tmpHome)

	// Load once to create directory
	cfg, err := LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	original := &TuiConfig{
		CollectionPath: []string{"~/manga/blue archive", "~/manga/naruto", "soul eater"},
		MangaSeries: map[string][]string{
			"blue archive": {"ch1", "ch2", "ch3"},
			"naruto":       {"ch1", "ch2 stuff here"},
			"soul eater":   {"ch0", "very long chapter name"},
		},
	}
	require.NoError(t, SaveConfig(original))

	loaded, err := LoadConfig()
	require.NotNil(t, loaded)
	require.NoError(t, err)

	assert.Equal(t, original.CollectionPath, loaded.CollectionPath)
	assert.Equal(t, original.MangaSeries, loaded.MangaSeries)
}

func TestCorruptedConf(t *testing.T) {
	tmpHome := t.TempDir()
	t.Setenv("HOME", tmpHome)

	// this should create a directory w/ a default config.json
	cfg, err := LoadConfig()
	require.NoError(t, err)
	require.NotNil(t, cfg)

	// write the corrupted data to the config.json
	corruptData := []byte("{corrupted: json}")
	testPath, err := getConfigFile()
	require.NoError(t, err)
	err = os.WriteFile(testPath, corruptData, 0o644)

	corruptCfg, loadErr := LoadConfig()
	assert.Error(t, loadErr)
	assert.Nil(t, corruptCfg)
}
