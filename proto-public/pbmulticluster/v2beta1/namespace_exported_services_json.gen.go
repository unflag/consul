// Code generated by protoc-json-shim. DO NOT EDIT.
package multiclusterv2beta1

import (
	protojson "google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON is a custom marshaler for NamespaceExportedServices
func (this *NamespaceExportedServices) MarshalJSON() ([]byte, error) {
	str, err := NamespaceExportedServicesMarshaler.Marshal(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for NamespaceExportedServices
func (this *NamespaceExportedServices) UnmarshalJSON(b []byte) error {
	return NamespaceExportedServicesUnmarshaler.Unmarshal(b, this)
}

var (
	NamespaceExportedServicesMarshaler   = &protojson.MarshalOptions{}
	NamespaceExportedServicesUnmarshaler = &protojson.UnmarshalOptions{DiscardUnknown: false}
)
