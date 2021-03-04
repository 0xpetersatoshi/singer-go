package catalog

import (
	"errors"
	"fmt"
)

// Stream in a Catalog
type Stream struct {
    TapStreamID string
    Stream string
    KeyProperties []string
    Schema string // TODO: make this a struct
    ReplicationKey string
    ReplicationMethod string
    IsView bool
    Database string
    Table string
    RowCount int
    StreamAlias string
    Metadata string // TODO: make this a struct
}

// Catalog contains streams
type Catalog struct {
    Streams []Stream
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