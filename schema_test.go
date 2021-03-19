package singer

import (
	"testing"
)

func TestGetSelected(t *testing.T) {
	schema := &Schema{Selected: true}

	isSelected := schema.GetSelected()
	if isSelected != true {
		t.Error("Expected true, got", isSelected)
	}
}
