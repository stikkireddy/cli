package configure_test

import (
	"context"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/databricks/cli/cmd"
	"github.com/stretchr/testify/assert"
	"gopkg.in/ini.v1"
)

func assertKeyValueInSection(t *testing.T, section *ini.Section, keyName, expectedValue string) {
	key, err := section.GetKey(keyName)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedValue, key.Value())
	}
}

func setup(t *testing.T) string {
	tempHomeDir := t.TempDir()
	homeEnvVar := "HOME"
	if runtime.GOOS == "windows" {
		homeEnvVar = "USERPROFILE"
	}
	t.Setenv(homeEnvVar, tempHomeDir)
	t.Setenv("DATABRICKS_CONFIG_FILE", "")
	t.Setenv("DATABRICKS_TOKEN", "")
	return tempHomeDir
}

func getTempFileWithContent(t *testing.T, tempHomeDir string, content string) *os.File {
	inp, err := os.CreateTemp(tempHomeDir, "input")
	assert.NoError(t, err)
	_, err = inp.WriteString(content)
	assert.NoError(t, err)
	err = inp.Sync()
	assert.NoError(t, err)
	_, err = inp.Seek(0, 0)
	assert.NoError(t, err)
	return inp
}

func TestDefaultConfigureNoInteractive(t *testing.T) {
	ctx := context.Background()
	tempHomeDir := setup(t)
	inp := getTempFileWithContent(t, tempHomeDir, "token\n")
	oldStdin := os.Stdin
	defer inp.Close()
	t.Cleanup(func() {
		os.Stdin = oldStdin
	})
	os.Stdin = inp

	cmd := cmd.New()
	cmd.SetArgs([]string{"configure", "--token", "--host", "https://host"})

	err := cmd.ExecuteContext(ctx)
	assert.NoError(t, err)

	cfgPath := filepath.Join(tempHomeDir, ".databrickscfg")
	_, err = os.Stat(cfgPath)
	assert.NoError(t, err)

	cfg, err := ini.Load(cfgPath)
	assert.NoError(t, err)

	defaultSection, err := cfg.GetSection("DEFAULT")
	assert.NoError(t, err)

	assertKeyValueInSection(t, defaultSection, "host", "https://host")
	assertKeyValueInSection(t, defaultSection, "token", "token")
}

func TestConfigFileFromEnvNoInteractive(t *testing.T) {
	//TODO: Replace with similar test code from go SDK, once we start using it directly
	ctx := context.Background()
	tempHomeDir := setup(t)
	cfgPath := filepath.Join(tempHomeDir, ".databrickscfg")
	t.Setenv("DATABRICKS_CONFIG_FILE", cfgPath)

	inp := getTempFileWithContent(t, tempHomeDir, "token\n")
	defer inp.Close()
	oldStdin := os.Stdin
	t.Cleanup(func() { os.Stdin = oldStdin })
	os.Stdin = inp

	cmd := cmd.New()
	cmd.SetArgs([]string{"configure", "--token", "--host", "https://host"})

	err := cmd.ExecuteContext(ctx)
	assert.NoError(t, err)

	_, err = os.Stat(cfgPath)
	assert.NoError(t, err)

	cfg, err := ini.Load(cfgPath)
	assert.NoError(t, err)

	defaultSection, err := cfg.GetSection("DEFAULT")
	assert.NoError(t, err)

	assertKeyValueInSection(t, defaultSection, "host", "https://host")
	assertKeyValueInSection(t, defaultSection, "token", "token")
}

func TestCustomProfileConfigureNoInteractive(t *testing.T) {
	ctx := context.Background()
	tempHomeDir := setup(t)
	cfgPath := filepath.Join(tempHomeDir, ".databrickscfg")
	inp := getTempFileWithContent(t, tempHomeDir, "token\n")
	defer inp.Close()
	oldStdin := os.Stdin
	t.Cleanup(func() { os.Stdin = oldStdin })
	os.Stdin = inp

	cmd := cmd.New()
	cmd.SetArgs([]string{"configure", "--token", "--host", "https://host", "--profile", "CUSTOM"})

	err := cmd.ExecuteContext(ctx)
	assert.NoError(t, err)

	_, err = os.Stat(cfgPath)
	assert.NoError(t, err)

	cfg, err := ini.Load(cfgPath)
	assert.NoError(t, err)

	defaultSection, err := cfg.GetSection("CUSTOM")
	assert.NoError(t, err)

	assertKeyValueInSection(t, defaultSection, "host", "https://host")
	assertKeyValueInSection(t, defaultSection, "token", "token")
}
