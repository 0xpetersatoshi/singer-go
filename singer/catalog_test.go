package singer

import (
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var stream *Entry = &Entry{
	Name:              "MyStream",
	ReplicationKey:    "LastUpdated",
	ReplicationMethod: "INCREMENTAL",
	KeyProperties:     []string{"a", "b", "c"},
	Metadata: []*Metadata{m1},
}

var stream2 *Entry = &Entry{
	Name:              "Stream2",
	ReplicationMethod: "FULL_TABLE",
	KeyProperties:     []string{"a", "b", "c"},
	Metadata: []*Metadata{m2},
}

var stream3 *Entry = &Entry{
	Name: "Stream3",
	ReplicationMethod: "INCREMENTAL",
	ReplicationKey: "modify_time",
	KeyProperties: []string{"id", "property_id", "attribute_id"},
	Metadata: []*Metadata{m1},
}

var m1 *Metadata = &Metadata{
	MetadataProps: mp1,
}

var m2 *Metadata = &Metadata{
	MetadataProps: mp2,
}

var mp1 = &MetadataProperties{
	Selected: true,
}

var mp2 = &MetadataProperties{
	Selected: false,
}

func TestGetSelectedStreamsCatalog(t *testing.T) {
	type test map[string]struct {
		catalog *Catalog
		want *Catalog
	}

	tests := test {
		"one selected stream": {
			catalog: &Catalog{Streams: []*Entry{stream, stream2}},
			want: &Catalog{Streams: []*Entry{stream}},
		},
		"two selected streams": {
			catalog: &Catalog{Streams: []*Entry{stream, stream2, stream3}},
			want: &Catalog{Streams: []*Entry{stream, stream3}},
		},
		"empty catalog error": {
			catalog: &Catalog{Streams: []*Entry{stream2}},
			want: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T){
			got, _ := GetSelectedStreamsCatalog(tc.catalog)
			diff := cmp.Diff(tc.want, got)
			if diff != "" {
				t.Fatal(diff)
			}
		})
	}
}

func TestGetKeyProperties(t *testing.T) {
	streamType := stream.GetKeyProperties()
	if reflect.DeepEqual(streamType, []string{"a", "b", "c"}) != true {
		t.Error("Expected 'a', 'b', 'c', got", streamType)
	}
}

func TestGetName(t *testing.T) {
	streamName := stream.GetName()
	if streamName != "MyStream" {
		t.Error("Expected MyStream, got", streamName)
	}
}

func TestGetReplicationMethod(t *testing.T) {
	replicationMethod := stream.GetReplicationMethod()
	if replicationMethod != "INCREMENTAL" {
		t.Error("Expected INCREMENTAL, got", replicationMethod)
	}
}

func TestGetReplicationKey(t *testing.T) {
	replicationKey := stream.GetReplicationKey()
	if replicationKey != "LastUpdated" {
		t.Error("Expected LastUpdated, got", replicationKey)
	}
}
