package singer

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

// Entry in a Catalog
type Entry struct {
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
	Streams []*Entry `json:"streams,omitempty"`
}

// GetStream iterates through slice and returns Stream struct matching stream name
func (c *Catalog) GetStream(streamID string) (*Entry, error) {
	for _, s := range c.Streams {
		if s.TapStreamID == streamID {
			return s, nil
		}
	}

	errorMessage := fmt.Sprintf("Stream %s not found", streamID)
	return &Entry{}, errors.New(errorMessage)
}

// Dump outputs all the streams in the Catalog to JSON
func (c *Catalog) Dump() string {
	bs, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON from Catalog: %s", err)
	}

	return string(bs)
}

// GetSelectedStreamsCatalog returns a pointer to a new catalog containing only selected streams
func GetSelectedStreamsCatalog(catalog *Catalog) (*Catalog, error) {
	newCatalog := &Catalog{}
	for _, stream := range catalog.Streams {
		for _, metadata := range stream.Metadata {
			if metadata != nil {
				if metadata.IsSelected() {
					newCatalog.Streams = append(newCatalog.Streams, stream)
				}
			}
		}
	}

	if len(newCatalog.Streams) > 0 {
		return newCatalog, nil
	} else {
		return nil, errors.New("Catalog is empty, no streams selected")
	}
}

func (s *Entry) GetKeyProperties() []string {
	return s.KeyProperties
}

func (s *Entry) GetReplicationMethod() string {
	return s.ReplicationMethod
}

func (s *Entry) GetReplicationKey() string {
	return s.ReplicationKey
}

func (s *Entry) GetName() string {
	return s.Name
}
