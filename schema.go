package singer

// Schema defines the shape and properties of a stream
type Schema struct {
	Type                 string            `json:"type,omitempty"`
	AdditionalProperties bool              `json:"additionalProperties,omitempty"`
	Properties           map[string]Schema `json:"properties,omitempty"`
	KeyProperties        []string          `json:"keyProperties,omitempty"`
	Selected             bool              `json:"selected,omitempty"`
	Inclusion            string            `json:"inclusion,omitempty"`
	Description          string            `json:"description,omitempty"`
	Minimum              float32           `json:"minimum,omitempty"`
	Maximum              float32           `json:"maximum,omitempty"`
	ExclusiveMinimum     float32           `json:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum     float32           `json:"exclusiveMaximum,omitempty"`
	MultipleOf           float32           `json:"multipleOf,omitempty"`
	MaxLength            int               `json:"maxLength,omitempty"`
	MinLength            int               `json:"minLength,omitempty"`
	Format               string            `json:"format,omitempty"`
	Anyof                string            `json:"anyOf,omitempty"`
	PatternProperties    string            `json:"patternProperties,omitempty"`
}
