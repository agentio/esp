// Code generated by protoc-gen-go.
// source: google/api/service.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import google_protobuf5 "google/protobuf"
import google_protobuf4 "google/protobuf"
import google_protobuf6 "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// `Service` is the root object of the configuration schema. It
// describes basic information like the name of the service and the
// exposed API interfaces, and delegates other aspects to configuration
// sub-sections.
//
// Example:
//
//     type: google.api.Service
//     config_version: 1
//     name: calendar.googleapis.com
//     title: Google Calendar API
//     apis:
//     - name: google.calendar.Calendar
//     backend:
//       rules:
//       - selector: "*"
//         address: calendar.example.com
type Service struct {
	// The version of the service configuration. The config version may
	// influence interpretation of the configuration, for example, to
	// determine defaults. This is documented together with applicable
	// options. The current default for the config version itself is `3`.
	ConfigVersion *google_protobuf6.UInt32Value `protobuf:"bytes,20,opt,name=config_version,json=configVersion" json:"config_version,omitempty"`
	// The DNS address at which this service is available,
	// e.g. `calendar.googleapis.com`.
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// A unique ID for a specific instance of this message, typically assigned
	// by the client for tracking purpose. If empty, the server may choose to
	// generate one instead.
	Id string `protobuf:"bytes,33,opt,name=id" json:"id,omitempty"`
	// The product title associated with this service.
	Title string `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	// The id of the Google developer project that owns the service.
	// Members of this project can manage the service configuration,
	// manage consumption of the service, etc.
	ProducerProjectId string `protobuf:"bytes,22,opt,name=producer_project_id,json=producerProjectId" json:"producer_project_id,omitempty"`
	// A list of API interfaces exported by this service. Only the `name` field
	// of the [google.protobuf.Api][google.protobuf.Api] needs to be provided by the configuration
	// author, as the remaining fields will be derived from the IDL during the
	// normalization process. It is an error to specify an API interface here
	// which cannot be resolved against the associated IDL files.
	Apis []*google_protobuf5.Api `protobuf:"bytes,3,rep,name=apis" json:"apis,omitempty"`
	// A list of all proto message types included in this API service.
	// Types referenced directly or indirectly by the `apis` are
	// automatically included.  Messages which are not referenced but
	// shall be included, such as types used by the `google.protobuf.Any` type,
	// should be listed here by name. Example:
	//
	//     types:
	//     - name: google.protobuf.Int32
	Types []*google_protobuf4.Type `protobuf:"bytes,4,rep,name=types" json:"types,omitempty"`
	// A list of all enum types included in this API service.  Enums
	// referenced directly or indirectly by the `apis` are automatically
	// included.  Enums which are not referenced but shall be included
	// should be listed here by name. Example:
	//
	//     enums:
	//     - name: google.someapi.v1.SomeEnum
	Enums []*google_protobuf4.Enum `protobuf:"bytes,5,rep,name=enums" json:"enums,omitempty"`
	// Additional API documentation.
	Documentation *Documentation `protobuf:"bytes,6,opt,name=documentation" json:"documentation,omitempty"`
	// API backend configuration.
	Backend *Backend `protobuf:"bytes,8,opt,name=backend" json:"backend,omitempty"`
	// HTTP configuration.
	Http *Http `protobuf:"bytes,9,opt,name=http" json:"http,omitempty"`
	// Auth configuration.
	Authentication *Authentication `protobuf:"bytes,11,opt,name=authentication" json:"authentication,omitempty"`
	// Context configuration.
	Context *Context `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	// Configuration controlling usage of this service.
	Usage *Usage `protobuf:"bytes,15,opt,name=usage" json:"usage,omitempty"`
	// Configuration of per-consumer project properties.
	ProjectProperties *ProjectProperties `protobuf:"bytes,17,opt,name=project_properties,json=projectProperties" json:"project_properties,omitempty"`
	// Configuration for the service control plane.
	Control *Control `protobuf:"bytes,21,opt,name=control" json:"control,omitempty"`
	// Defines the logs used by this service.
	Logs []*LogDescriptor `protobuf:"bytes,23,rep,name=logs" json:"logs,omitempty"`
	// Defines the metrics used by this service.
	Metrics []*MetricDescriptor `protobuf:"bytes,24,rep,name=metrics" json:"metrics,omitempty"`
	// Defines the monitored resources used by this service. This is required
	// by the [Service.monitoring][google.api.Service.monitoring] and [Service.logging][google.api.Service.logging] configurations.
	//
	MonitoredResources []*MonitoredResourceDescriptor `protobuf:"bytes,25,rep,name=monitored_resources,json=monitoredResources" json:"monitored_resources,omitempty"`
	// Billing configuration of the service.
	Billing *Billing `protobuf:"bytes,26,opt,name=billing" json:"billing,omitempty"`
	// Logging configuration of the service.
	Logging *Logging `protobuf:"bytes,27,opt,name=logging" json:"logging,omitempty"`
	// Monitoring configuration of the service.
	Monitoring *Monitoring `protobuf:"bytes,28,opt,name=monitoring" json:"monitoring,omitempty"`
	// Configuration for system parameters.
	SystemParameters *SystemParameters `protobuf:"bytes,29,opt,name=system_parameters,json=systemParameters" json:"system_parameters,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor16, []int{0} }

func (m *Service) GetConfigVersion() *google_protobuf6.UInt32Value {
	if m != nil {
		return m.ConfigVersion
	}
	return nil
}

func (m *Service) GetApis() []*google_protobuf5.Api {
	if m != nil {
		return m.Apis
	}
	return nil
}

func (m *Service) GetTypes() []*google_protobuf4.Type {
	if m != nil {
		return m.Types
	}
	return nil
}

func (m *Service) GetEnums() []*google_protobuf4.Enum {
	if m != nil {
		return m.Enums
	}
	return nil
}

func (m *Service) GetDocumentation() *Documentation {
	if m != nil {
		return m.Documentation
	}
	return nil
}

func (m *Service) GetBackend() *Backend {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *Service) GetHttp() *Http {
	if m != nil {
		return m.Http
	}
	return nil
}

func (m *Service) GetAuthentication() *Authentication {
	if m != nil {
		return m.Authentication
	}
	return nil
}

func (m *Service) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *Service) GetUsage() *Usage {
	if m != nil {
		return m.Usage
	}
	return nil
}

func (m *Service) GetProjectProperties() *ProjectProperties {
	if m != nil {
		return m.ProjectProperties
	}
	return nil
}

func (m *Service) GetControl() *Control {
	if m != nil {
		return m.Control
	}
	return nil
}

func (m *Service) GetLogs() []*LogDescriptor {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *Service) GetMetrics() []*MetricDescriptor {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Service) GetMonitoredResources() []*MonitoredResourceDescriptor {
	if m != nil {
		return m.MonitoredResources
	}
	return nil
}

func (m *Service) GetBilling() *Billing {
	if m != nil {
		return m.Billing
	}
	return nil
}

func (m *Service) GetLogging() *Logging {
	if m != nil {
		return m.Logging
	}
	return nil
}

func (m *Service) GetMonitoring() *Monitoring {
	if m != nil {
		return m.Monitoring
	}
	return nil
}

func (m *Service) GetSystemParameters() *SystemParameters {
	if m != nil {
		return m.SystemParameters
	}
	return nil
}

func init() {
	proto.RegisterType((*Service)(nil), "google.api.Service")
}

func init() { proto.RegisterFile("google/api/service.proto", fileDescriptor16) }

var fileDescriptor16 = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x94, 0x6f, 0x4f, 0x1b, 0x39,
	0x10, 0xc6, 0x15, 0x48, 0xe0, 0x30, 0x7f, 0x8e, 0x98, 0x00, 0x26, 0x84, 0x53, 0xb8, 0x3b, 0x89,
	0x48, 0x15, 0x41, 0x02, 0x89, 0xb7, 0x15, 0x81, 0x4a, 0x45, 0xa2, 0x52, 0x64, 0x0a, 0xea, 0xbb,
	0xc8, 0xd9, 0x35, 0x8b, 0xdb, 0x5d, 0xdb, 0xb2, 0xbd, 0xb4, 0xf9, 0x10, 0xfd, 0xce, 0xd5, 0xda,
	0xde, 0xc4, 0x49, 0x96, 0x77, 0xd9, 0x79, 0x7e, 0xcf, 0x64, 0x6c, 0xcf, 0x0c, 0x40, 0x89, 0x10,
	0x49, 0x4a, 0x2f, 0x88, 0x64, 0x17, 0x9a, 0xaa, 0x37, 0x16, 0xd1, 0xbe, 0x54, 0xc2, 0x08, 0x08,
	0x9c, 0xd2, 0x27, 0x92, 0xb5, 0x3b, 0x01, 0x45, 0x38, 0x17, 0x86, 0x18, 0x26, 0xb8, 0x76, 0x64,
	0x7b, 0x3f, 0x54, 0x73, 0xf3, 0xea, 0xc3, 0x61, 0xea, 0x31, 0x89, 0x7e, 0x50, 0x1e, 0x57, 0x29,
	0x2c, 0x4d, 0x19, 0x4f, 0xbc, 0x72, 0x14, 0x28, 0x91, 0xe0, 0x3a, 0xcf, 0xa8, 0xaa, 0x30, 0x45,
	0x82, 0x1b, 0xfa, 0xcb, 0xbc, 0xa3, 0x28, 0x91, 0x7a, 0xe5, 0x9f, 0x40, 0x89, 0x45, 0x94, 0x67,
	0x94, 0xbb, 0xd2, 0x2b, 0x2a, 0x7f, 0x35, 0x46, 0xfa, 0xf0, 0x41, 0x10, 0x4e, 0xc9, 0x98, 0x96,
	0xe9, 0x5a, 0x61, 0x5c, 0x24, 0x15, 0x7f, 0x9f, 0x8a, 0x24, 0x99, 0x9d, 0xe6, 0x30, 0x50, 0x32,
	0x6a, 0x14, 0x8b, 0xbc, 0xf0, 0x5f, 0x28, 0x08, 0xce, 0x8c, 0x50, 0x34, 0x1e, 0x29, 0xaa, 0x45,
	0xae, 0xca, 0x07, 0x68, 0x1f, 0x2f, 0x43, 0xb3, 0xd4, 0xa7, 0xe1, 0xbb, 0x4d, 0xb4, 0xa1, 0xd9,
	0x48, 0x12, 0x45, 0x32, 0x6a, 0xa6, 0x17, 0x16, 0x9e, 0x22, 0xd7, 0x24, 0xa1, 0x0b, 0x77, 0x6c,
	0xbf, 0xc6, 0xf9, 0xcb, 0x05, 0xe1, 0x93, 0x77, 0x25, 0xc9, 0xbc, 0xd4, 0x5e, 0x94, 0xcc, 0x44,
	0xd2, 0x85, 0x6b, 0x9e, 0x6a, 0x3f, 0x15, 0x91, 0x92, 0x2a, 0xdf, 0x20, 0xff, 0xfe, 0xde, 0x00,
	0xeb, 0x8f, 0xae, 0xb9, 0xe0, 0x2d, 0xd8, 0x89, 0x04, 0x7f, 0x61, 0xc9, 0xe8, 0x8d, 0x2a, 0xcd,
	0x04, 0x47, 0xad, 0x6e, 0xad, 0xb7, 0x79, 0xd9, 0xe9, 0xfb, 0x7e, 0x2b, 0x93, 0xf4, 0x9f, 0xee,
	0xb9, 0xb9, 0xba, 0x7c, 0x26, 0x69, 0x4e, 0xf1, 0xb6, 0xf3, 0x3c, 0x3b, 0x0b, 0x84, 0xa0, 0xce,
	0x49, 0x46, 0x51, 0xad, 0x5b, 0xeb, 0x6d, 0x60, 0xfb, 0x1b, 0xee, 0x80, 0x15, 0x16, 0xa3, 0x53,
	0x1b, 0x59, 0x61, 0x31, 0x6c, 0x81, 0x86, 0x61, 0x26, 0xa5, 0x68, 0xc5, 0x86, 0xdc, 0x07, 0xec,
	0x83, 0x3d, 0xa9, 0x44, 0x9c, 0x47, 0x54, 0x8d, 0xa4, 0x12, 0xdf, 0x69, 0x64, 0x46, 0x2c, 0x46,
	0x07, 0x96, 0x69, 0x96, 0xd2, 0xd0, 0x29, 0xf7, 0x31, 0xec, 0x81, 0x3a, 0x91, 0x4c, 0xa3, 0xd5,
	0xee, 0x6a, 0x6f, 0xf3, 0xb2, 0xb5, 0x54, 0xe4, 0x8d, 0x64, 0xd8, 0x12, 0xf0, 0x03, 0x68, 0x14,
	0x57, 0xa2, 0x51, 0xdd, 0xa2, 0xfb, 0x4b, 0xe8, 0xd7, 0x89, 0xa4, 0xd8, 0x31, 0x05, 0x4c, 0x79,
	0x9e, 0x69, 0xd4, 0x78, 0x07, 0xfe, 0xc4, 0xf3, 0x0c, 0x3b, 0x06, 0x7e, 0x04, 0xdb, 0x73, 0xcd,
	0x8b, 0xd6, 0xec, 0x8d, 0x1d, 0xf5, 0x67, 0x13, 0xda, 0xbf, 0x0b, 0x01, 0x3c, 0xcf, 0xc3, 0x73,
	0xb0, 0xee, 0x07, 0x10, 0xfd, 0x65, 0xad, 0x7b, 0xa1, 0x75, 0xe0, 0x24, 0x5c, 0x32, 0xf0, 0x7f,
	0x50, 0x2f, 0x86, 0x01, 0x6d, 0x58, 0x76, 0x37, 0x64, 0x3f, 0x1b, 0x23, 0xb1, 0x55, 0xe1, 0x00,
	0xec, 0x14, 0xc3, 0x4e, 0xb9, 0x61, 0x91, 0x2b, 0x6b, 0xd3, 0xf2, 0xed, 0x90, 0xbf, 0x99, 0x23,
	0xf0, 0x82, 0xa3, 0x28, 0xcc, 0x8f, 0x32, 0xda, 0x5a, 0x2e, 0xec, 0xd6, 0x49, 0xb8, 0x64, 0xe0,
	0x19, 0x68, 0xd8, 0x46, 0x46, 0x7f, 0x5b, 0xb8, 0x19, 0xc2, 0x4f, 0x85, 0x80, 0x9d, 0x0e, 0x1f,
	0x00, 0x2c, 0x1f, 0x57, 0x2a, 0x21, 0xa9, 0x32, 0x8c, 0x6a, 0xd4, 0xb4, 0xae, 0x93, 0xd0, 0xe5,
	0x1f, 0x7a, 0x38, 0x85, 0x6c, 0x0f, 0xcc, 0x87, 0xca, 0x2a, 0x95, 0x48, 0xd1, 0x7e, 0x75, 0x95,
	0x4a, 0xa4, 0xb8, 0x64, 0xe0, 0x39, 0xa8, 0xa7, 0x22, 0xd1, 0xe8, 0xd0, 0x3e, 0xed, 0xdc, 0x2b,
	0x3d, 0x88, 0xe4, 0x8e, 0xea, 0x48, 0x31, 0x69, 0x84, 0xc2, 0x16, 0x83, 0xd7, 0x60, 0xdd, 0xed,
	0x06, 0x8d, 0x90, 0x75, 0x74, 0x42, 0xc7, 0x17, 0x2b, 0x05, 0xa6, 0x12, 0x86, 0xdf, 0xc0, 0xde,
	0xf2, 0xea, 0xd0, 0xe8, 0xc8, 0xe6, 0x38, 0x9b, 0xcb, 0x51, 0x62, 0xd8, 0x53, 0x41, 0x3a, 0x98,
	0x2d, 0x8a, 0xf6, 0xbc, 0x7e, 0x2b, 0xa3, 0x76, 0x45, 0xbb, 0x38, 0x09, 0x97, 0x4c, 0x81, 0xfb,
	0xb5, 0x87, 0x8e, 0x97, 0xf1, 0x07, 0x27, 0xe1, 0x92, 0x81, 0xd7, 0x00, 0xcc, 0xb6, 0x19, 0xea,
	0x58, 0xc7, 0x41, 0x45, 0xb9, 0x85, 0x29, 0x20, 0xe1, 0x3d, 0x68, 0x2e, 0x2e, 0x3a, 0x8d, 0x4e,
	0xe6, 0x77, 0x47, 0x61, 0x7f, 0xb4, 0xd0, 0x70, 0xca, 0xe0, 0x5d, 0xbd, 0x10, 0x19, 0x74, 0x8b,
	0x1d, 0x94, 0x05, 0xa6, 0xc1, 0x96, 0x5f, 0x4f, 0xc3, 0x62, 0xfe, 0x86, 0xb5, 0xf1, 0x9a, 0x1d,
	0xc4, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x57, 0x98, 0xf1, 0x1f, 0x07, 0x00, 0x00,
}
