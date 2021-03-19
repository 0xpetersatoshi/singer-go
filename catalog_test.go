package singer

import (
	"reflect"
	"testing"
)

var stream *Stream = &Stream{
	Name:              "MyStream",
	ReplicationKey:    "LastUpdated",
	ReplicationMethod: "INCREMENTAL",
	KeyProperties:     []string{"a", "b", "c"},
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
