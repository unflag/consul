// Code generated by protoc-gen-go-binary. DO NOT EDIT.
// source: pbmulticluster/v2beta1/exported_services.proto

package multiclusterv2beta1

import (
	"google.golang.org/protobuf/proto"
)

// MarshalBinary implements encoding.BinaryMarshaler
func (msg *ExportedServices) MarshalBinary() ([]byte, error) {
	return proto.Marshal(msg)
}

// UnmarshalBinary implements encoding.BinaryUnmarshaler
func (msg *ExportedServices) UnmarshalBinary(b []byte) error {
	return proto.Unmarshal(b, msg)
}
