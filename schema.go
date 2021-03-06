package singer

// Schema defines the shape and properties of a stream
type Schema struct {
    Type []string
    AdditionalProperties bool
    Properties Property
    Selected bool
    Inclusion bool
    Description string
}

// Property defines the properties of a field
type Property struct {
    PropertyName string
    Type []string
    Format string
    AdditionalProperties bool
    MultipleOf float64
    Minimum float64
    Maximum float64
    Selected bool
    Inclusion bool
    Description string
}