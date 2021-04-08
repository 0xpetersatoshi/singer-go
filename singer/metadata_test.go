package singer

import "testing"


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

var tests = []struct {
	name string
	metadata *Metadata
	expected bool
}{
	{"stream is selected", mOne, true},
	{"stream is not selected", mTwo, false},
}


func TestIsSelected(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T){
			result := tt.metadata.IsSelected()
			if result != tt.expected {
				t.Errorf("metadata.IsSelected() got %v, want %v", result, tt.expected)
			}
		})
	}
}
