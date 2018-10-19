package json

import (
	"io"

	"github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func BorrowIterator(data []byte) *jsoniter.Iterator {
	return json.BorrowIterator(data)
}

func ReturnIterator(iter *jsoniter.Iterator) {
	json.ReturnIterator(iter)
}

func BorrowStream(writer io.Writer) *jsoniter.Stream {
	return json.BorrowStream(writer)
}

func ReturnStream(stream *jsoniter.Stream) {
	json.ReturnStream(stream)
}

// MarshalToString convenient method to write as string instead of []byte
func MarshalToString(v interface{}) (string, error) {
	return json.MarshalToString(v)
}

// Marshal adapts to json/encoding Marshal API
//
// Marshal returns the JSON encoding of v, adapts to json/encoding Marshal API
// Refer to https://godoc.org/encoding/json#Marshal for more information
func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// MarshalIndent same as json.MarshalIndent. Prefix is not supported.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

// UnmarshalFromString convenient method to read from string instead of []byte
func UnmarshalFromString(str string, v interface{}) error {
	return json.UnmarshalFromString(str, v)
}

// Unmarshal adapts to json/encoding Unmarshal API
//
// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v.
// Refer to https://godoc.org/encoding/json#Unmarshal for more information
func Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

// Get quick method to get value from deeply nested JSON structure
func Get(data []byte, path ...interface{}) jsoniter.Any {
	return json.Get(data, path...)
}

// NewEncoder same as json.NewEncoder
func NewEncoder(writer io.Writer) *jsoniter.Encoder {
	return json.NewEncoder(writer)
}

// NewDecoder adapts to json/stream NewDecoder API.
//
// NewDecoder returns a new decoder that reads from r.
//
// Instead of a json/encoding Decoder, an Decoder is returned
// Refer to https://godoc.org/encoding/json#NewDecoder for more information
func NewDecoder(reader io.Reader) *jsoniter.Decoder {
	return json.NewDecoder(reader)
}

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	return json.Valid(data)
}

func RegisterExtension(extension jsoniter.Extension) {
	json.RegisterExtension(extension)
}

func DecoderOf(typ reflect2.Type) jsoniter.ValDecoder {
	return json.DecoderOf(typ)
}

func EncoderOf(typ reflect2.Type) jsoniter.ValEncoder {
	return json.EncoderOf(typ)
}
