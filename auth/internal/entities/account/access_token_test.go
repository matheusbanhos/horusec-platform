package account

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAccessToken(t *testing.T) {
	t.Run("should return no error when valid data", func(t *testing.T) {
		data := &AccessToken{
			AccessToken: "test",
		}

		assert.NoError(t, data.Validate())
	})
}

func TestToBytesAccessToken(t *testing.T) {
	t.Run("should success parse to bytes", func(t *testing.T) {
		data := &AccessToken{
			AccessToken: "test",
		}

		assert.NotEmpty(t, data.ToBytes())
	})
}
