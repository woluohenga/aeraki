// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/metaprotocol/v1alpha1/metaprotocol_application_protocol.proto

// $schema: metaprotocol.aeraki.io.v1alpha1.ApplicationProtocol
// $title: Application Protocol
// $description: ApplicationProtocol defines an application protocol built on top of MetaProtocol.
//
// ApplicationProtocol defines an application protocol built on top of MetaProtocol.
//
// ```yaml
// apiVersion: metaprotocol.aeraki.io/v1alpha1
// kind: ApplicationProtocol
// metadata:
//   name: dubbo
//   namespace: istio-system
// spec:
//   protocol: dubbo
//   codec: aeraki.meta_protocol.codec.dubbo
// ```

package v1alpha1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ApplicationProtocol
func (this *ApplicationProtocol) MarshalJSON() ([]byte, error) {
	str, err := MetaprotocolApplicationProtocolMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ApplicationProtocol
func (this *ApplicationProtocol) UnmarshalJSON(b []byte) error {
	return MetaprotocolApplicationProtocolUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	MetaprotocolApplicationProtocolMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	MetaprotocolApplicationProtocolUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{AllowUnknownFields: true}
)
