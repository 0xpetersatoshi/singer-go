package singer

type Metadata struct {
	MetadataProps   *MetadataProperties `json:"metadata"`
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

func (m *Metadata) GetStandardMetadata(stream *Entry) {
	m.MetadataProps.StreamName = stream.GetName()
	m.MetadataProps.TableKeyProperties = stream.GetKeyProperties()
	m.MetadataProps.ForcedReplicationMethod = stream.GetReplicationMethod()
	m.MetadataProps.ValidReplicationKeys = []string{stream.GetReplicationKey()}
}

func (m *Metadata) toSlice() {}

func (m *Metadata) write() {}

func (m *Metadata) IsSelected() bool {
	return m.MetadataProps.Selected
}
