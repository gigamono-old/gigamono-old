package configs

import (
	"reflect"

	"github.com/gofrs/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

// Decodes unsupported types like UUID.
func getCustomDecoder() viper.DecoderConfigOption {
	return viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		func(_ reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			switch t {
			case reflect.TypeOf(UUID{}): // When UUID type is encountered.
				parsedID, err := uuid.FromString(data.(string)) // Get UUID from string
				if err != nil {
					return UUID{}, err
				}
				id := UUID(parsedID) // Wrap it in our custom UUID type.
				return id, nil
			}
			return data, nil
		},
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))
}
