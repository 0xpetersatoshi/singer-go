package singer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)


var mpOne = &MetadataProperties{
	Inclusion: "available",
	Selected: true,
	TableKeyProperties: []string{"id"},
	ForcedReplicationMethod: "FULL_TABLE",
	StreamName: "test-stream",
}
var mpTwo = &MetadataProperties{
	Inclusion: "available",
	Selected: false,
	TableKeyProperties: []string{"table_id"},
	ForcedReplicationMethod: "INCREMENTAL",
	ValidReplicationKeys: []string{"updated_at"},
	StreamName: "other-stream",
}

var mOne = &Metadata{
	Metadata: mpOne,
}
var mTwo = &Metadata{
	Metadata: mpTwo,
}

func TestIsSelected(t *testing.T) {
	type test map[string]struct {
		metadata *Metadata
		want bool
	}

	tests := test {
		"stream is selected": {mOne, true},
		"stream is not selected": {mTwo, false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T){
			got := tc.metadata.IsSelected()
			if got != tc.want {
				t.Errorf("test %s: want %#v, got %#v", name, tc.want, got)
			}
		})
	}
}

func TestGetStandardMetadata(t *testing.T) {
	type test map[string]struct {
		metadata *Metadata
		input *Entry
		want *MetadataProperties
	}

	entryOne := &Entry{
		Name: "some-stream",
		KeyProperties: []string{"id"},
		ReplicationMethod: "FULL_TABLE",
		ReplicationKey: "",
	}

	tests := test {
		"standard case": {
			metadata: &Metadata{
				Metadata: mpOne,
			},
			input: entryOne,
			want: &MetadataProperties{
				StreamName: "some-stream",
				TableKeyProperties: []string{"id"},
				ForcedReplicationMethod: "FULL_TABLE",
				ValidReplicationKeys: []string{""},
				Inclusion: "available",
				Selected: true,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T){
			tc.metadata.GetStandardMetadata(tc.input)
			diff := cmp.Diff(tc.want, tc.metadata.Metadata)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
