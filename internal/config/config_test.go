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
	assert.Equal(t, []string{}, cfg.CollectionPaths)

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
		CollectionPaths: []string{"/home/manga", "/home/other-library"},
		MangaSeries: rootDir{
			"/home/manga": subDir{
				"blue-archive/hina": []string{
					"somefile.cbz",
					"somefile2.cbz",
					"somefile3.cbz",
				},
				"chainsaw-man": []string{
					"file1.cbz",
					"file2.cbz",
					"file3.cbz",
				},
			},
			"/home/other-library": subDir{
				"frieren": []string{
					"somefile.cbz",
					"somefile2.cbz",
					"somefile3.cbz",
				},
				"naruto": []string{
					"file1.cbz",
					"file2.cbz",
					"file3.cbz",
				},
			},
		},
	}
	require.NoError(t, SaveConfig(original))

	loaded, err := LoadConfig()
	require.NotNil(t, loaded)
	require.NoError(t, err)

	assert.Equal(t, original.CollectionPaths, loaded.CollectionPaths)
	assert.Equal(t, original.MangaSeries, loaded.MangaSeries)

	for _, p := range loaded.CollectionPaths {
		_, exists := loaded.MangaSeries[p]
		assert.True(t, exists)
	}
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
