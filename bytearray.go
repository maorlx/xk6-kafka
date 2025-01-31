package kafka

import (
	"fmt"

	"github.com/riferrei/srclient"
)

const (
	ByteArray srclient.SchemaType = "BYTEARRAY"

	ByteArraySerializer   string = "org.apache.kafka.common.serialization.ByteArraySerializer"
	ByteArrayDeserializer string = "org.apache.kafka.common.serialization.ByteArrayDeserializer"
)

// SerializeByteArray serializes the given data into a byte array and returns it.
// If the data is not a byte array, an error is returned. The configuration, topic, element,
// schema and version are just used to conform with the interface.
func SerializeByteArray(configuration Configuration, topic string, data interface{}, element Element, schema string, version int) ([]byte, *Xk6KafkaError) {
	switch data.(type) {
	case []interface{}:
		bArray := data.([]interface{})
		arr := make([]byte, len(bArray))
		for i, u := range bArray {
			arr[i] = byte(u.(int64))
		}
		return arr, nil
	default:
		return nil, NewXk6KafkaError(
			invalidDataType,
			"Invalid data type provided for byte array serializer (requires []byte)",
			fmt.Errorf("Expected: []byte, got: %T", data))
	}
}

// DeserializeByteArray deserializes the given data from a byte array and returns it.
// It just returns the data as is. The configuration, topic, element, schema and version
// are just used to conform with the interface.
func DeserializeByteArray(configuration Configuration, topic string, data []byte, element Element, schema string, version int) (interface{}, *Xk6KafkaError) {
	return data, nil
}
