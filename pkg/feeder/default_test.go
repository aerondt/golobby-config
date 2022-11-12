package feeder_test

import (
	"testing"

	"github.com/golobby/config/v3/pkg/feeder"
	"github.com/stretchr/testify/assert"
)

func TestDefault_Feed(t *testing.T) {
	type config struct {
		App struct {
			Name string `env:"APP_NAME" default:"Shop"`
			Port int    `env:"APP_PORT" default:"8585"`
		}
		Debug bool `env:"DEBUG" default:"true"`
	}

	c := config{}
	f := feeder.Default{}

	err := f.Feed(&c)
	assert.NoError(t, err)

	assert.Equal(t, "Shop", c.App.Name)
	assert.Equal(t, 8585, c.App.Port)
	assert.Equal(t, true, c.Debug)
}

func TestDefault_Feed_With_Invalid_Struct_It_Should_Fail(t *testing.T) {
	c := struct {
		App struct {
			Name float64 `env:"APP_NAME" default:"string"`
		}
	}{}
	f := feeder.Default{}

	err := f.Feed(&c)
	assert.Error(t, err)
}

func TestDefault_Feed_With_Invalid_Inner_Struct_It_Should_Fail(t *testing.T) {
	type HttpConf struct {
		RequestTimeoutMs float64 `default:"string"`
	}
	type App struct {
		Name string `env:"APP_NAME" default:"string"`
		Http *HttpConf
	}
	c := App{
		Http: &HttpConf{},
	}

	f := feeder.Default{}
	err := f.Feed(&c)
	assert.Error(t, err)

}

func TestDefault_Feed_With_Non_Struct_Type_Should_Fail(t *testing.T) {
	c := 42
	f := feeder.Default{}

	err := f.Feed(c)
	assert.Error(t, err)
}
