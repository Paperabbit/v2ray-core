package dns

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	router "v2ray.com/core/app/router"
	net "v2ray.com/core/common/net"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type DomainMatchingType int32

const (
	DomainMatchingType_Full      DomainMatchingType = 0
	DomainMatchingType_Subdomain DomainMatchingType = 1
	DomainMatchingType_Keyword   DomainMatchingType = 2
	DomainMatchingType_Regex     DomainMatchingType = 3
	DomainMatchingType_New       DomainMatchingType = 4
)

var DomainMatchingType_name = map[int32]string{
	0: "Full",
	1: "Subdomain",
	2: "Keyword",
	3: "Regex",
	4: "New",
}

var DomainMatchingType_value = map[string]int32{
	"Full":      0,
	"Subdomain": 1,
	"Keyword":   2,
	"Regex":     3,
	"New":       4,
}

func (x DomainMatchingType) String() string {
	return proto.EnumName(DomainMatchingType_name, int32(x))
}

func (DomainMatchingType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{0}
}

type Config_Fake_RegenerationType int32

const (
	// Discard the LRU (Least recently used) IP (Default)
	Config_Fake_LRU Config_Fake_RegenerationType = 0
	// Discard the oldest IP
	Config_Fake_Oldest Config_Fake_RegenerationType = 1
	// Don't regenerate IP
	Config_Fake_None Config_Fake_RegenerationType = 2
)

var Config_Fake_RegenerationType_name = map[int32]string{
	0: "LRU",
	1: "Oldest",
	2: "None",
}

var Config_Fake_RegenerationType_value = map[string]int32{
	"LRU":    0,
	"Oldest": 1,
	"None":   2,
}

func (x Config_Fake_RegenerationType) String() string {
	return proto.EnumName(Config_Fake_RegenerationType_name, int32(x))
}

func (Config_Fake_RegenerationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{1, 2, 0}
}

type NameServer struct {
	Address              *net.Endpoint                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PrioritizedDomain    []*NameServer_PriorityDomain `protobuf:"bytes,2,rep,name=prioritized_domain,json=prioritizedDomain,proto3" json:"prioritized_domain,omitempty"`
	Geoip                []*router.GeoIP              `protobuf:"bytes,3,rep,name=geoip,proto3" json:"geoip,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *NameServer) Reset()         { *m = NameServer{} }
func (m *NameServer) String() string { return proto.CompactTextString(m) }
func (*NameServer) ProtoMessage()    {}
func (*NameServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{0}
}

func (m *NameServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServer.Unmarshal(m, b)
}
func (m *NameServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServer.Marshal(b, m, deterministic)
}
func (m *NameServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServer.Merge(m, src)
}
func (m *NameServer) XXX_Size() int {
	return xxx_messageInfo_NameServer.Size(m)
}
func (m *NameServer) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServer.DiscardUnknown(m)
}

var xxx_messageInfo_NameServer proto.InternalMessageInfo

func (m *NameServer) GetAddress() *net.Endpoint {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *NameServer) GetPrioritizedDomain() []*NameServer_PriorityDomain {
	if m != nil {
		return m.PrioritizedDomain
	}
	return nil
}

func (m *NameServer) GetGeoip() []*router.GeoIP {
	if m != nil {
		return m.Geoip
	}
	return nil
}

type NameServer_PriorityDomain struct {
	Type                 DomainMatchingType `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.dns.DomainMatchingType" json:"type,omitempty"` // Deprecated: Do not use.
	Domain               string             `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NameServer_PriorityDomain) Reset()         { *m = NameServer_PriorityDomain{} }
func (m *NameServer_PriorityDomain) String() string { return proto.CompactTextString(m) }
func (*NameServer_PriorityDomain) ProtoMessage()    {}
func (*NameServer_PriorityDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{0, 0}
}

func (m *NameServer_PriorityDomain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameServer_PriorityDomain.Unmarshal(m, b)
}
func (m *NameServer_PriorityDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameServer_PriorityDomain.Marshal(b, m, deterministic)
}
func (m *NameServer_PriorityDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameServer_PriorityDomain.Merge(m, src)
}
func (m *NameServer_PriorityDomain) XXX_Size() int {
	return xxx_messageInfo_NameServer_PriorityDomain.Size(m)
}
func (m *NameServer_PriorityDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_NameServer_PriorityDomain.DiscardUnknown(m)
}

var xxx_messageInfo_NameServer_PriorityDomain proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *NameServer_PriorityDomain) GetType() DomainMatchingType {
	if m != nil {
		return m.Type
	}
	return DomainMatchingType_Full
}

func (m *NameServer_PriorityDomain) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

type Config struct {
	// Nameservers used by this DNS. Only traditional UDP servers are support at the moment.
	// A special value 'localhost' as a domain address can be set to use DNS on local system.
	NameServers []*net.Endpoint `protobuf:"bytes,1,rep,name=NameServers,proto3" json:"NameServers,omitempty"` // Deprecated: Do not use.
	// NameServer list used by this DNS client.
	NameServer []*NameServer `protobuf:"bytes,5,rep,name=name_server,json=nameServer,proto3" json:"name_server,omitempty"`
	// Static hosts. Domain to IP.
	// Deprecated. Use static_hosts.
	Hosts map[string]*net.IPOrDomain `protobuf:"bytes,2,rep,name=Hosts,proto3" json:"Hosts,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` // Deprecated: Do not use.
	// Client IP for EDNS client subnet. Must be 4 bytes (IPv4) or 16 bytes (IPv6).
	ClientIp    []byte                `protobuf:"bytes,3,opt,name=client_ip,json=clientIp,proto3" json:"client_ip,omitempty"`
	StaticHosts []*Config_HostMapping `protobuf:"bytes,4,rep,name=static_hosts,json=staticHosts,proto3" json:"static_hosts,omitempty"`
	// Tag is the inbound tag of DNS client.
	Tag                  string                     `protobuf:"bytes,6,opt,name=tag,proto3" json:"tag,omitempty"`
	Fake                 *Config_Fake               `protobuf:"bytes,7,opt,name=fake,proto3" json:"fake,omitempty"`
	ExternalRules        map[string]*ConfigPatterns `protobuf:"bytes,8,rep,name=external_rules,json=externalRules,proto3" json:"external_rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{1}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Config) GetNameServers() []*net.Endpoint {
	if m != nil {
		return m.NameServers
	}
	return nil
}

func (m *Config) GetNameServer() []*NameServer {
	if m != nil {
		return m.NameServer
	}
	return nil
}

// Deprecated: Do not use.
func (m *Config) GetHosts() map[string]*net.IPOrDomain {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Config) GetClientIp() []byte {
	if m != nil {
		return m.ClientIp
	}
	return nil
}

func (m *Config) GetStaticHosts() []*Config_HostMapping {
	if m != nil {
		return m.StaticHosts
	}
	return nil
}

func (m *Config) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *Config) GetFake() *Config_Fake {
	if m != nil {
		return m.Fake
	}
	return nil
}

func (m *Config) GetExternalRules() map[string]*ConfigPatterns {
	if m != nil {
		return m.ExternalRules
	}
	return nil
}

type Config_HostMapping struct {
	Type   DomainMatchingType `protobuf:"varint,1,opt,name=type,proto3,enum=v2ray.core.app.dns.DomainMatchingType" json:"type,omitempty"` // Deprecated: Do not use.
	Domain string             `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Ip     [][]byte           `protobuf:"bytes,3,rep,name=ip,proto3" json:"ip,omitempty"`
	// ProxiedDomain indicates the mapped domain has the same IP address on this domain. V2Ray will use this domain for IP queries.
	// This field is only effective if ip is empty.
	ProxiedDomain        string   `protobuf:"bytes,4,opt,name=proxied_domain,json=proxiedDomain,proto3" json:"proxied_domain,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config_HostMapping) Reset()         { *m = Config_HostMapping{} }
func (m *Config_HostMapping) String() string { return proto.CompactTextString(m) }
func (*Config_HostMapping) ProtoMessage()    {}
func (*Config_HostMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{1, 1}
}

func (m *Config_HostMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_HostMapping.Unmarshal(m, b)
}
func (m *Config_HostMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_HostMapping.Marshal(b, m, deterministic)
}
func (m *Config_HostMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_HostMapping.Merge(m, src)
}
func (m *Config_HostMapping) XXX_Size() int {
	return xxx_messageInfo_Config_HostMapping.Size(m)
}
func (m *Config_HostMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_HostMapping.DiscardUnknown(m)
}

var xxx_messageInfo_Config_HostMapping proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *Config_HostMapping) GetType() DomainMatchingType {
	if m != nil {
		return m.Type
	}
	return DomainMatchingType_Full
}

func (m *Config_HostMapping) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Config_HostMapping) GetIp() [][]byte {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *Config_HostMapping) GetProxiedDomain() string {
	if m != nil {
		return m.ProxiedDomain
	}
	return ""
}

type Config_Fake struct {
	FakeRules            []string                     `protobuf:"bytes,1,rep,name=fake_rules,json=fakeRules,proto3" json:"fake_rules,omitempty"`
	FakeNet              string                       `protobuf:"bytes,2,opt,name=fake_net,json=fakeNet,proto3" json:"fake_net,omitempty"`
	Regeneration         Config_Fake_RegenerationType `protobuf:"varint,3,opt,name=regeneration,proto3,enum=v2ray.core.app.dns.Config_Fake_RegenerationType" json:"regeneration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *Config_Fake) Reset()         { *m = Config_Fake{} }
func (m *Config_Fake) String() string { return proto.CompactTextString(m) }
func (*Config_Fake) ProtoMessage()    {}
func (*Config_Fake) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{1, 2}
}

func (m *Config_Fake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config_Fake.Unmarshal(m, b)
}
func (m *Config_Fake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config_Fake.Marshal(b, m, deterministic)
}
func (m *Config_Fake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config_Fake.Merge(m, src)
}
func (m *Config_Fake) XXX_Size() int {
	return xxx_messageInfo_Config_Fake.Size(m)
}
func (m *Config_Fake) XXX_DiscardUnknown() {
	xxx_messageInfo_Config_Fake.DiscardUnknown(m)
}

var xxx_messageInfo_Config_Fake proto.InternalMessageInfo

func (m *Config_Fake) GetFakeRules() []string {
	if m != nil {
		return m.FakeRules
	}
	return nil
}

func (m *Config_Fake) GetFakeNet() string {
	if m != nil {
		return m.FakeNet
	}
	return ""
}

func (m *Config_Fake) GetRegeneration() Config_Fake_RegenerationType {
	if m != nil {
		return m.Regeneration
	}
	return Config_Fake_LRU
}

type ConfigPatterns struct {
	Patterns             []string `protobuf:"bytes,1,rep,name=patterns,proto3" json:"patterns,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigPatterns) Reset()         { *m = ConfigPatterns{} }
func (m *ConfigPatterns) String() string { return proto.CompactTextString(m) }
func (*ConfigPatterns) ProtoMessage()    {}
func (*ConfigPatterns) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed5695198e3def8f, []int{1, 3}
}

func (m *ConfigPatterns) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigPatterns.Unmarshal(m, b)
}
func (m *ConfigPatterns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigPatterns.Marshal(b, m, deterministic)
}
func (m *ConfigPatterns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigPatterns.Merge(m, src)
}
func (m *ConfigPatterns) XXX_Size() int {
	return xxx_messageInfo_ConfigPatterns.Size(m)
}
func (m *ConfigPatterns) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigPatterns.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigPatterns proto.InternalMessageInfo

func (m *ConfigPatterns) GetPatterns() []string {
	if m != nil {
		return m.Patterns
	}
	return nil
}

func init() {
	proto.RegisterEnum("v2ray.core.app.dns.DomainMatchingType", DomainMatchingType_name, DomainMatchingType_value)
	proto.RegisterEnum("v2ray.core.app.dns.Config_Fake_RegenerationType", Config_Fake_RegenerationType_name, Config_Fake_RegenerationType_value)
	proto.RegisterType((*NameServer)(nil), "v2ray.core.app.dns.NameServer")
	proto.RegisterType((*NameServer_PriorityDomain)(nil), "v2ray.core.app.dns.NameServer.PriorityDomain")
	proto.RegisterType((*Config)(nil), "v2ray.core.app.dns.Config")
	proto.RegisterMapType((map[string]*ConfigPatterns)(nil), "v2ray.core.app.dns.Config.ExternalRulesEntry")
	proto.RegisterMapType((map[string]*net.IPOrDomain)(nil), "v2ray.core.app.dns.Config.HostsEntry")
	proto.RegisterType((*Config_HostMapping)(nil), "v2ray.core.app.dns.Config.HostMapping")
	proto.RegisterType((*Config_Fake)(nil), "v2ray.core.app.dns.Config.Fake")
	proto.RegisterType((*ConfigPatterns)(nil), "v2ray.core.app.dns.Config.patterns")
}

func init() {
	proto.RegisterFile("v2ray.com/core/app/dns/config.proto", fileDescriptor_ed5695198e3def8f)
}

var fileDescriptor_ed5695198e3def8f = []byte{
	// 767 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0xae, 0x9d, 0xff, 0xe3, 0x6c, 0x64, 0xe6, 0xa2, 0x32, 0xe6, 0xaf, 0x6c, 0xd5, 0x10, 0x81,
	0x70, 0xc0, 0x45, 0x82, 0x72, 0x41, 0x45, 0xdb, 0x14, 0x22, 0x68, 0x1a, 0x4d, 0x17, 0x2e, 0x00,
	0x29, 0x9a, 0xda, 0xa7, 0x59, 0x6b, 0x93, 0x99, 0xd1, 0x78, 0xb2, 0x5d, 0xf3, 0x2c, 0x3c, 0x01,
	0x12, 0x4f, 0xc1, 0x15, 0x6f, 0x85, 0x66, 0xec, 0x6d, 0x7e, 0x77, 0x97, 0x1b, 0xee, 0x66, 0x8e,
	0xbf, 0xef, 0xfc, 0x7c, 0xdf, 0xb1, 0x0d, 0x77, 0xcf, 0x63, 0xc5, 0x8a, 0x28, 0x11, 0xcb, 0x61,
	0x22, 0x14, 0x0e, 0x99, 0x94, 0xc3, 0x94, 0xe7, 0xc3, 0x44, 0xf0, 0x57, 0xd9, 0x3c, 0x92, 0x4a,
	0x68, 0x41, 0xc8, 0x25, 0x48, 0x61, 0xc4, 0xa4, 0x8c, 0x52, 0x9e, 0x87, 0x1f, 0xed, 0x10, 0x13,
	0xb1, 0x5c, 0x0a, 0x3e, 0xe4, 0xa8, 0x87, 0x2c, 0x4d, 0x15, 0xe6, 0x79, 0x49, 0x0e, 0x3f, 0xb9,
	0x1a, 0x98, 0x62, 0xae, 0x33, 0xce, 0x74, 0x26, 0x78, 0x05, 0xee, 0x1f, 0x68, 0x47, 0x89, 0x95,
	0x46, 0xb5, 0xd5, 0xd1, 0xf1, 0xdf, 0x2e, 0xc0, 0x84, 0x2d, 0xf1, 0x05, 0xaa, 0x73, 0x54, 0xe4,
	0x01, 0xb4, 0xaa, 0xa2, 0x81, 0x73, 0xc7, 0x19, 0x78, 0xf1, 0x07, 0xd1, 0x46, 0xcb, 0x65, 0xc5,
	0x88, 0xa3, 0x8e, 0x46, 0x3c, 0x95, 0x22, 0xe3, 0x9a, 0x5e, 0xe2, 0xc9, 0x6f, 0x40, 0xa4, 0xca,
	0x84, 0xca, 0x74, 0xf6, 0x3b, 0xa6, 0xb3, 0x54, 0x2c, 0x59, 0xc6, 0x03, 0xf7, 0x4e, 0x6d, 0xe0,
	0xc5, 0x9f, 0x46, 0xfb, 0x83, 0x47, 0xeb, 0xb2, 0xd1, 0xb4, 0x24, 0x16, 0x4f, 0x2c, 0x89, 0xbe,
	0xb5, 0x91, 0xa8, 0x0c, 0x91, 0x18, 0x1a, 0x73, 0x14, 0x99, 0x0c, 0x6a, 0x36, 0xe1, 0xbb, 0xbb,
	0x09, 0xcb, 0xd9, 0xa2, 0xef, 0x50, 0x8c, 0xa7, 0xb4, 0x84, 0x86, 0xa7, 0xd0, 0xdb, 0x4e, 0x4c,
	0xbe, 0x81, 0xba, 0x2e, 0x24, 0xda, 0xd9, 0x7a, 0x71, 0xff, 0x50, 0x57, 0x25, 0xf2, 0x19, 0xd3,
	0xc9, 0x69, 0xc6, 0xe7, 0x27, 0x85, 0xc4, 0x47, 0x6e, 0xe0, 0x50, 0xcb, 0x23, 0xb7, 0xa1, 0xf9,
	0x66, 0x2e, 0x67, 0xd0, 0xa1, 0xd5, 0xed, 0xf8, 0xaf, 0x36, 0x34, 0x1f, 0x5b, 0x59, 0xc9, 0x08,
	0xbc, 0xf5, 0x60, 0x46, 0xc5, 0xda, 0x7f, 0x50, 0xd1, 0x96, 0xd8, 0xe4, 0x91, 0x87, 0xe0, 0x71,
	0xb6, 0xc4, 0x59, 0x6e, 0xef, 0x41, 0xc3, 0xa6, 0x79, 0xff, 0x7a, 0x19, 0x29, 0xf0, 0xb5, 0x93,
	0x0f, 0xa1, 0xf1, 0xbd, 0xc8, 0x75, 0x5e, 0x39, 0x70, 0xef, 0x10, 0xb5, 0x6c, 0x39, 0xb2, 0xb8,
	0x11, 0xd7, 0xaa, 0xb0, 0x7d, 0x94, 0x3c, 0xf2, 0x0e, 0x74, 0x92, 0x45, 0x86, 0x5c, 0xcf, 0xac,
	0xea, 0xce, 0xa0, 0x4b, 0xdb, 0x65, 0x60, 0x2c, 0xc9, 0x18, 0xba, 0xb9, 0x66, 0x3a, 0x4b, 0x66,
	0xa7, 0xb6, 0x48, 0xdd, 0x16, 0xe9, 0xdf, 0x50, 0xe4, 0x19, 0x93, 0x32, 0xe3, 0x73, 0xea, 0x95,
	0xdc, 0xb2, 0x8e, 0x0f, 0x35, 0xcd, 0xe6, 0x41, 0xd3, 0x0a, 0x6a, 0x8e, 0xe4, 0x3e, 0xd4, 0x5f,
	0xb1, 0x33, 0x0c, 0x5a, 0xfb, 0x1b, 0xb8, 0x93, 0xf4, 0x29, 0x3b, 0x43, 0x6a, 0xc1, 0xe4, 0x04,
	0x7a, 0x78, 0xa1, 0x51, 0x71, 0xb6, 0x98, 0xa9, 0xd5, 0x02, 0xf3, 0xa0, 0x7d, 0xf5, 0xea, 0x55,
	0xf4, 0x51, 0x45, 0xa0, 0x06, 0x6f, 0x05, 0xa0, 0x47, 0xb8, 0x19, 0x0b, 0x7f, 0x05, 0x58, 0xab,
	0x63, 0x5a, 0x3d, 0xc3, 0xc2, 0x6e, 0x4f, 0x87, 0x9a, 0x23, 0xf9, 0x12, 0x1a, 0xe7, 0x6c, 0xb1,
	0x42, 0xbb, 0x0f, 0x5e, 0xfc, 0xe1, 0x15, 0x3e, 0x8f, 0xa7, 0xcf, 0x55, 0xb5, 0xdb, 0x25, 0xfe,
	0x6b, 0xf7, 0x2b, 0x27, 0xfc, 0xc3, 0x01, 0x6f, 0x43, 0x96, 0xff, 0x6b, 0x3b, 0x49, 0x0f, 0xdc,
	0xea, 0xc5, 0xe9, 0x52, 0x37, 0x93, 0xe4, 0x1e, 0xf4, 0xa4, 0x12, 0x17, 0xd9, 0xfa, 0x2d, 0xad,
	0x5b, 0xfc, 0x51, 0x15, 0x2d, 0x8b, 0x84, 0xff, 0x38, 0x50, 0x37, 0x02, 0x93, 0xf7, 0x00, 0x8c,
	0xc4, 0x95, 0xac, 0x66, 0xa3, 0x3b, 0xb4, 0x63, 0x22, 0x56, 0x23, 0xf2, 0x36, 0xb4, 0xed, 0x63,
	0x8e, 0xba, 0x2a, 0xdc, 0x32, 0xf7, 0x09, 0x6a, 0x72, 0x02, 0x5d, 0x85, 0x73, 0xe4, 0xa8, 0xec,
	0xb7, 0xc9, 0xae, 0x51, 0x2f, 0xfe, 0xec, 0x06, 0x47, 0x23, 0xba, 0xc1, 0x31, 0x33, 0xd2, 0xad,
	0x2c, 0xc7, 0x9f, 0x83, 0xbf, 0x8b, 0x20, 0x2d, 0xa8, 0xfd, 0x48, 0x7f, 0xf2, 0x6f, 0x11, 0x80,
	0xe6, 0xf3, 0x85, 0xf9, 0x1e, 0xfa, 0x0e, 0x69, 0x43, 0x7d, 0x22, 0x38, 0xfa, 0x6e, 0xd8, 0x87,
	0xb6, 0x64, 0xda, 0x38, 0x9b, 0x93, 0x70, 0x7d, 0xae, 0x86, 0x79, 0x73, 0x0f, 0x11, 0xc8, 0xfe,
	0x52, 0x1c, 0xf0, 0xfd, 0xc1, 0xb6, 0xef, 0x77, 0xaf, 0x99, 0xe8, 0x32, 0xf7, 0x86, 0xf3, 0x1f,
	0x4f, 0x80, 0xec, 0x3b, 0x69, 0xda, 0x7d, 0xba, 0x5a, 0x2c, 0xfc, 0x5b, 0xe4, 0x08, 0x3a, 0x2f,
	0x56, 0x2f, 0x4b, 0x73, 0x7c, 0x87, 0x78, 0xd0, 0xfa, 0x01, 0x8b, 0xd7, 0x42, 0xa5, 0xbe, 0x4b,
	0x3a, 0xd0, 0x30, 0xd3, 0x5f, 0xf8, 0x35, 0x33, 0xf4, 0x04, 0x5f, 0xfb, 0xf5, 0x47, 0x5f, 0xc0,
	0xed, 0x44, 0x2c, 0x0f, 0x34, 0x31, 0x75, 0x7e, 0xa9, 0xa5, 0x3c, 0xff, 0xd3, 0x25, 0x3f, 0xc7,
	0x94, 0x15, 0xd1, 0x63, 0xf3, 0xec, 0x5b, 0x29, 0xa3, 0x27, 0x3c, 0x7f, 0xd9, 0xb4, 0xbf, 0x80,
	0xfb, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x2a, 0x6a, 0x1f, 0xbb, 0x06, 0x00, 0x00,
}
