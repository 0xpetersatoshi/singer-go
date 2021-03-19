package singer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Stream in a Catalog
type Stream struct {
	TapStreamID       string      `json:"tap_stream_id,omitempty"`
	Name              string      `json:"stream,omitempty"`
	KeyProperties     []string    `json:"key_properties,omitempty"`
	Schema            *Schema     `json:"schema,omitempty"`
	ReplicationKey    string      `json:"replication_key,omitempty"`
	ReplicationMethod string      `json:"replication_method,omitempty"`
	IsView            bool        `json:"is_view,omitempty"`
	Database          string      `json:"database,omitempty"`
	Table             string      `json:"table,omitempty"`
	RowCount          int         `json:"row_count,omitempty"`
	StreamAlias       string      `json:"stream_alias,omitempty"`
	Metadata          []*Metadata `json:"metadata"`
}

// Catalog contains streams
type Catalog struct {
	Streams []Stream `json:"streams,omitempty"`
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

func (s *Stream) GetKeyProperties() []string {
	return s.KeyProperties
}

func (s *Stream) GetReplicationMethod() string {
	return s.ReplicationMethod
}

func (s *Stream) GetReplicationKey() string {
	return s.ReplicationKey
}

func (s *Stream) GetName() string {
	return s.Name
}
