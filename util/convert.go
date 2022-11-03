package util

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
)

func AsJson(data interface{}) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		log.Err(err).Msg("error parsing data to json")
		return nil
	}
	return bytes
}
