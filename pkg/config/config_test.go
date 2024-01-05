package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		createTestConfigFile(".env", "APP_ENV=local\nAPP_TZ=Asia/Jakarta\nDATABASE_DRIVER=mysql\nDATABASE_HOST=127.0.0.1\n")

		defer cleanupTestEnvironment()

		config, err := NewConfig()
		assert.NoError(t, err, "Expected no error")
		assert.NotNil(t, config, "Expected non-nil config")

		assert.Equal(t, "local", config.AppEnv, "Unexpected value for Key1")
		assert.Equal(t, "Asia/Jakarta", config.AppTz, "Unexpected value for Key2")
	})

	t.Run("ErrorReadingConfig", func(t *testing.T) {
		createTestConfigFile(".env", "invalid_content")

		defer cleanupTestEnvironment()

		_, err := NewConfig()
		assert.Error(t, err, "Expected an error")
	})
}

func createTestConfigFile(fileName, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		panic(err)
	}
}

func cleanupTestEnvironment() {
	os.Remove(".env")
}
