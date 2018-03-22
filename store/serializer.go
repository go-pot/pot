package store

import "encoding/json"

// SessionSerializer provides an interface hook for alternative serializers
type SessionSerializer interface {
	Deserialize(d []byte, ss interface{}) error
	Serialize(ss interface{}) ([]byte, error)
}

// JSONSerializer encode the session map to JSON.
type JSONSerializer struct{}

// Serialize to JSON. Will err if there are unmarshalable key values
func (s JSONSerializer) Serialize(ss interface{}) ([]byte, error) {
	return json.Marshal(ss)
}

// Deserialize back to map[string]interface{}
func (s JSONSerializer) Deserialize(d []byte, ss interface{}) error {
	return json.Unmarshal(d, &ss)
}
