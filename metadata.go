package singer

type Metadata struct {
	Metadata   MetadataProperties `json:"metadata"`
	Breadcrumb []string           `json:"breadcrumb"`
}

type MetadataProperties struct {
	Inclusion               string   `json:"inclusion"`
	Selected                bool     `json:"selected"`
	TableKeyProperties      []string `json:"table-key-properties"`
	ForcedReplicationMethod string   `json:"forced-replication-method"`
	ValidReplicationKeys    []string `json:"valid-replication-keys"`
	StreamName              string   `json:"stream_name"`
}

func (m *Metadata) GetStandardMetadata(stream *Stream) {
	m.Metadata.StreamName = stream.GetName()
	m.Metadata.TableKeyProperties = stream.GetKeyProperties()
	m.Metadata.ForcedReplicationMethod = stream.GetReplicationMethod()
	m.Metadata.ValidReplicationKeys = []string{stream.GetReplicationKey()}
}

func (m *Metadata) toSlice() {}

func (m *Metadata) write() {}
