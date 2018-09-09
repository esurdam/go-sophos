package sophos

// Definition represents a Swagger API endpoint
// You can use the Swagger API Documents to identify all the different RESTful API end
// points with descriptions for each object and node.
// GET /api/definitions
//
// This returns a list of possible Swagger API definitions and you can define the call so
// the results are specific to different objects or nodes:
// GET /api/definitions/network
type Definition struct {
	Description string
	Name        string
	Link        string
}

// GetSwag will use the Client to request its Swag
func (d *Definition) GetSwag(c *Client, options ...Option) (Swag, error) {
	var swag Swag
	r, err := c.Get(d.Link, options...)
	if err != nil {
		return swag, err
	}

	err = r.MarshalTo(&swag)
	return swag, err
}

// Swag represents a Swagger API document.
// The Swagger API document contains API endpoints along with parameters and object
// definitions for those endpoints. When objects have references to other objects the
// type is a regular string (REF_ string). Since not all references are allowed, each object
// has a description that states which subset of an object can be used as a reference. For
// example, the string REF(network/*) means that all network objects can be used as ref
// erences while REF(network/host) means that only network host objects can be used.
type Swag struct {
	Paths map[string]MethodMap
}

// MethodMap is a map of Methods -> MethodDescriptions
// e.g. get: MethodDescriptions{}
type MethodMap map[string]MethodDescriptions

// MethodDescriptions describes the endpoint
type MethodDescriptions struct {
	Description string
	Parameters  []Parameter
	Tags        []string
	Responses   map[int]struct{ Description string }
}

// A Parameter defines the arguments that can be passed to the endpoint
type Parameter struct {
	Name        string
	In          string
	Description string
	Type        string
	Required    bool
}
