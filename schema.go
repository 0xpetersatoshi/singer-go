package singer

// Schema defines the shape and properties of a stream
type Schema struct {
    Type string `json:"type"`
    AdditionalProperties bool `json:"additionalProperties"`
    Properties map[string]Schema `json:"properties"`
    KeyProperties []string `json:"keyProperties"`
    Selected bool `json:"selected"`
    Inclusion string `json:"inclusion"`
    Description string `json:"description"`
    Minimum float32 `json:"minimum"`
    Maximum float32 `json:"maximum"`
    ExclusiveMinimum float32 `json:"exclusiveMinimum"`
    ExclusiveMaximum float32 `json:"exclusiveMaximum"`
    MultipleOf float32 `json:"multipleOf"`
    MaxLength int `json:"maxLength"`
    MinLength int `json:"minLength"`
    Format string `json:"format"`
    Anyof string `json:"anyOf"`
    PatternProperties string `json:"patternProperties"`
}
