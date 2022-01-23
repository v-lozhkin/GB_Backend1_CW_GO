package config

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildConfigFileNotFound(t *testing.T) {
	configFile := "../../config.yaml"
	c, err := BuildConfig(configFile)
	require.Error(t, err)
	assert.Nil(t, c)
}

func TestBuildConfigEmptyFile(t *testing.T) {
	configFile := "./testdata/empty.config.yml"
	c, err := BuildConfig(configFile)
	require.Error(t, err)
	assert.Nil(t, c)
}

func TestBuildConfigSuccessDecode(t *testing.T) {
	configFile := "../../configs/config.yml"
	c, err := BuildConfig(configFile)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, "debug", c.LogLevel)
	assert.Equal(t, "badabum", c.JWTSecret)
	assert.Equal(t, 8083, c.Port)
}

func TestBuildConfigSuccessEnvProcess(t *testing.T) {
	logLevel := "INFO"
	port := 8083
	jwtSecret := "secret"
	os.Clearenv()
	os.Setenv("APP_LOG_LEVEL", logLevel)
	os.Setenv("APP_PORT", strconv.Itoa(port))
	os.Setenv("APP_JWT_SECRET", jwtSecret)
	configFile := "../../configs/config.yml"
	c, err := BuildConfig(configFile)
	require.NoError(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, logLevel, c.LogLevel)
	assert.Equal(t, jwtSecret, c.JWTSecret)
	assert.Equal(t, port, c.Port)
}
