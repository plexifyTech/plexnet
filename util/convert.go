package util

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

// AsJson Takes any struct and serializes it into a byte array
func AsJson(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Err(err).Msg("error parsing data to json")
		return nil
	}
	return bytes
}

// FromJson Takes a json byte array and unmarshalls it into a given object. Pointer has to be passed
func FromJson(data []byte, target interface{}) {
	if err := json.Unmarshal(data, target); err != nil {
		log.Err(err).Msg("error parsing data to json")
	}
}
