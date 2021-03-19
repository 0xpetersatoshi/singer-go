package singer

import (
	"reflect"
	"testing"
)

func TestGetKeyProperties(t *testing.T) {
    schema := &Schema{
        Type: "MySchema",
        AdditionalProperties: true,
        Description: "Description for my schema.",
        KeyProperties: []string{"a", "b", "c"},
    }

    schemaType := schema.GetKeyProperties()
    if reflect.DeepEqual(schemaType, []string{"a", "b", "c"}) != true {
        t.Error("Expected 'a', 'b', 'c', got", schemaType)
    }
}

func TestGetSelected(t *testing.T) {
    schema := &Schema{Selected: true}

    isSelected := schema.GetSelected()
    if isSelected != true {
        t.Error("Expected true, got", isSelected)
    }
}