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
