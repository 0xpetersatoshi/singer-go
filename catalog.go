package singer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Stream in a Catalog
type Stream struct {
	TapStreamID       string   `json:"tap_stream_id"`
	Stream            string   `json:"stream"`
	KeyProperties     []string `json:"key_properties"`
	Schema            *Schema  `json:"schema"`
	ReplicationKey    string   `json:"replication_key"`
	ReplicationMethod string   `json:"replication_method"`
	IsView            bool     `json:"is_view"`
	Database          string   `json:"database"`
	Table             string   `json:"table"`
	RowCount          int      `json:"row_count"`
	StreamAlias       string   `json:"stream_alias"`
	Metadata          string   `json:"metadata"` // TODO: make this a struct
}

// Catalog contains streams
type Catalog struct {
	Streams []Stream `json:"streams"`
}

// GetStream iterates through slice and returns Stream struct matching stream name
func (c Catalog) GetStream(streamID string) (Stream, error) {
	for _, s := range c.Streams {
		if s.TapStreamID == streamID {
			return s, nil
		}
	}

	errorMessage := fmt.Sprintf("Stream %s not found", streamID)
	return Stream{}, errors.New(errorMessage)
}

// Dump outputs all the streams in the Catalog to JSON
func (c Catalog) Dump() string {
	bs, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON from Catalog: %s", err)
	}

	return string(bs)
}
