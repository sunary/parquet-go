// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inotify.proto

package hadoop_hdfs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EventType int32

const (
	EventType_EVENT_CREATE   EventType = 0
	EventType_EVENT_CLOSE    EventType = 1
	EventType_EVENT_APPEND   EventType = 2
	EventType_EVENT_RENAME   EventType = 3
	EventType_EVENT_METADATA EventType = 4
	EventType_EVENT_UNLINK   EventType = 5
	EventType_EVENT_TRUNCATE EventType = 6
)

var EventType_name = map[int32]string{
	0: "EVENT_CREATE",
	1: "EVENT_CLOSE",
	2: "EVENT_APPEND",
	3: "EVENT_RENAME",
	4: "EVENT_METADATA",
	5: "EVENT_UNLINK",
	6: "EVENT_TRUNCATE",
}
var EventType_value = map[string]int32{
	"EVENT_CREATE":   0,
	"EVENT_CLOSE":    1,
	"EVENT_APPEND":   2,
	"EVENT_RENAME":   3,
	"EVENT_METADATA": 4,
	"EVENT_UNLINK":   5,
	"EVENT_TRUNCATE": 6,
}

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}
func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}
func (x *EventType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EventType_value, data, "EventType")
	if err != nil {
		return err
	}
	*x = EventType(value)
	return nil
}
func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

type INodeType int32

const (
	INodeType_I_TYPE_FILE      INodeType = 0
	INodeType_I_TYPE_DIRECTORY INodeType = 1
	INodeType_I_TYPE_SYMLINK   INodeType = 2
)

var INodeType_name = map[int32]string{
	0: "I_TYPE_FILE",
	1: "I_TYPE_DIRECTORY",
	2: "I_TYPE_SYMLINK",
}
var INodeType_value = map[string]int32{
	"I_TYPE_FILE":      0,
	"I_TYPE_DIRECTORY": 1,
	"I_TYPE_SYMLINK":   2,
}

func (x INodeType) Enum() *INodeType {
	p := new(INodeType)
	*p = x
	return p
}
func (x INodeType) String() string {
	return proto.EnumName(INodeType_name, int32(x))
}
func (x *INodeType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(INodeType_value, data, "INodeType")
	if err != nil {
		return err
	}
	*x = INodeType(value)
	return nil
}
func (INodeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

type MetadataUpdateType int32

const (
	MetadataUpdateType_META_TYPE_TIMES       MetadataUpdateType = 0
	MetadataUpdateType_META_TYPE_REPLICATION MetadataUpdateType = 1
	MetadataUpdateType_META_TYPE_OWNER       MetadataUpdateType = 2
	MetadataUpdateType_META_TYPE_PERMS       MetadataUpdateType = 3
	MetadataUpdateType_META_TYPE_ACLS        MetadataUpdateType = 4
	MetadataUpdateType_META_TYPE_XATTRS      MetadataUpdateType = 5
)

var MetadataUpdateType_name = map[int32]string{
	0: "META_TYPE_TIMES",
	1: "META_TYPE_REPLICATION",
	2: "META_TYPE_OWNER",
	3: "META_TYPE_PERMS",
	4: "META_TYPE_ACLS",
	5: "META_TYPE_XATTRS",
}
var MetadataUpdateType_value = map[string]int32{
	"META_TYPE_TIMES":       0,
	"META_TYPE_REPLICATION": 1,
	"META_TYPE_OWNER":       2,
	"META_TYPE_PERMS":       3,
	"META_TYPE_ACLS":        4,
	"META_TYPE_XATTRS":      5,
}

func (x MetadataUpdateType) Enum() *MetadataUpdateType {
	p := new(MetadataUpdateType)
	*p = x
	return p
}
func (x MetadataUpdateType) String() string {
	return proto.EnumName(MetadataUpdateType_name, int32(x))
}
func (x *MetadataUpdateType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MetadataUpdateType_value, data, "MetadataUpdateType")
	if err != nil {
		return err
	}
	*x = MetadataUpdateType(value)
	return nil
}
func (MetadataUpdateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

type EventProto struct {
	Type             *EventType `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.EventType" json:"type,omitempty"`
	Contents         []byte     `protobuf:"bytes,2,req,name=contents" json:"contents,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *EventProto) Reset()                    { *m = EventProto{} }
func (m *EventProto) String() string            { return proto.CompactTextString(m) }
func (*EventProto) ProtoMessage()               {}
func (*EventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *EventProto) GetType() EventType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return EventType_EVENT_CREATE
}

func (m *EventProto) GetContents() []byte {
	if m != nil {
		return m.Contents
	}
	return nil
}

type EventBatchProto struct {
	Txid             *int64        `protobuf:"varint,1,req,name=txid" json:"txid,omitempty"`
	Events           []*EventProto `protobuf:"bytes,2,rep,name=events" json:"events,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *EventBatchProto) Reset()                    { *m = EventBatchProto{} }
func (m *EventBatchProto) String() string            { return proto.CompactTextString(m) }
func (*EventBatchProto) ProtoMessage()               {}
func (*EventBatchProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *EventBatchProto) GetTxid() int64 {
	if m != nil && m.Txid != nil {
		return *m.Txid
	}
	return 0
}

func (m *EventBatchProto) GetEvents() []*EventProto {
	if m != nil {
		return m.Events
	}
	return nil
}

type CreateEventProto struct {
	Type             *INodeType         `protobuf:"varint,1,req,name=type,enum=hadoop.hdfs.INodeType" json:"type,omitempty"`
	Path             *string            `protobuf:"bytes,2,req,name=path" json:"path,omitempty"`
	Ctime            *int64             `protobuf:"varint,3,req,name=ctime" json:"ctime,omitempty"`
	OwnerName        *string            `protobuf:"bytes,4,req,name=ownerName" json:"ownerName,omitempty"`
	GroupName        *string            `protobuf:"bytes,5,req,name=groupName" json:"groupName,omitempty"`
	Perms            *FsPermissionProto `protobuf:"bytes,6,req,name=perms" json:"perms,omitempty"`
	Replication      *int32             `protobuf:"varint,7,opt,name=replication" json:"replication,omitempty"`
	SymlinkTarget    *string            `protobuf:"bytes,8,opt,name=symlinkTarget" json:"symlinkTarget,omitempty"`
	Overwrite        *bool              `protobuf:"varint,9,opt,name=overwrite" json:"overwrite,omitempty"`
	DefaultBlockSize *int64             `protobuf:"varint,10,opt,name=defaultBlockSize,def=0" json:"defaultBlockSize,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *CreateEventProto) Reset()                    { *m = CreateEventProto{} }
func (m *CreateEventProto) String() string            { return proto.CompactTextString(m) }
func (*CreateEventProto) ProtoMessage()               {}
func (*CreateEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

const Default_CreateEventProto_DefaultBlockSize int64 = 0

func (m *CreateEventProto) GetType() INodeType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return INodeType_I_TYPE_FILE
}

func (m *CreateEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *CreateEventProto) GetCtime() int64 {
	if m != nil && m.Ctime != nil {
		return *m.Ctime
	}
	return 0
}

func (m *CreateEventProto) GetOwnerName() string {
	if m != nil && m.OwnerName != nil {
		return *m.OwnerName
	}
	return ""
}

func (m *CreateEventProto) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *CreateEventProto) GetPerms() *FsPermissionProto {
	if m != nil {
		return m.Perms
	}
	return nil
}

func (m *CreateEventProto) GetReplication() int32 {
	if m != nil && m.Replication != nil {
		return *m.Replication
	}
	return 0
}

func (m *CreateEventProto) GetSymlinkTarget() string {
	if m != nil && m.SymlinkTarget != nil {
		return *m.SymlinkTarget
	}
	return ""
}

func (m *CreateEventProto) GetOverwrite() bool {
	if m != nil && m.Overwrite != nil {
		return *m.Overwrite
	}
	return false
}

func (m *CreateEventProto) GetDefaultBlockSize() int64 {
	if m != nil && m.DefaultBlockSize != nil {
		return *m.DefaultBlockSize
	}
	return Default_CreateEventProto_DefaultBlockSize
}

type CloseEventProto struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	FileSize         *int64  `protobuf:"varint,2,req,name=fileSize" json:"fileSize,omitempty"`
	Timestamp        *int64  `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CloseEventProto) Reset()                    { *m = CloseEventProto{} }
func (m *CloseEventProto) String() string            { return proto.CompactTextString(m) }
func (*CloseEventProto) ProtoMessage()               {}
func (*CloseEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *CloseEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *CloseEventProto) GetFileSize() int64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *CloseEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type TruncateEventProto struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	FileSize         *int64  `protobuf:"varint,2,req,name=fileSize" json:"fileSize,omitempty"`
	Timestamp        *int64  `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TruncateEventProto) Reset()                    { *m = TruncateEventProto{} }
func (m *TruncateEventProto) String() string            { return proto.CompactTextString(m) }
func (*TruncateEventProto) ProtoMessage()               {}
func (*TruncateEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{4} }

func (m *TruncateEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *TruncateEventProto) GetFileSize() int64 {
	if m != nil && m.FileSize != nil {
		return *m.FileSize
	}
	return 0
}

func (m *TruncateEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type AppendEventProto struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	NewBlock         *bool   `protobuf:"varint,2,opt,name=newBlock,def=0" json:"newBlock,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AppendEventProto) Reset()                    { *m = AppendEventProto{} }
func (m *AppendEventProto) String() string            { return proto.CompactTextString(m) }
func (*AppendEventProto) ProtoMessage()               {}
func (*AppendEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{5} }

const Default_AppendEventProto_NewBlock bool = false

func (m *AppendEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *AppendEventProto) GetNewBlock() bool {
	if m != nil && m.NewBlock != nil {
		return *m.NewBlock
	}
	return Default_AppendEventProto_NewBlock
}

type RenameEventProto struct {
	SrcPath          *string `protobuf:"bytes,1,req,name=srcPath" json:"srcPath,omitempty"`
	DestPath         *string `protobuf:"bytes,2,req,name=destPath" json:"destPath,omitempty"`
	Timestamp        *int64  `protobuf:"varint,3,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RenameEventProto) Reset()                    { *m = RenameEventProto{} }
func (m *RenameEventProto) String() string            { return proto.CompactTextString(m) }
func (*RenameEventProto) ProtoMessage()               {}
func (*RenameEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{6} }

func (m *RenameEventProto) GetSrcPath() string {
	if m != nil && m.SrcPath != nil {
		return *m.SrcPath
	}
	return ""
}

func (m *RenameEventProto) GetDestPath() string {
	if m != nil && m.DestPath != nil {
		return *m.DestPath
	}
	return ""
}

func (m *RenameEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type MetadataUpdateEventProto struct {
	Path             *string             `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Type             *MetadataUpdateType `protobuf:"varint,2,req,name=type,enum=hadoop.hdfs.MetadataUpdateType" json:"type,omitempty"`
	Mtime            *int64              `protobuf:"varint,3,opt,name=mtime" json:"mtime,omitempty"`
	Atime            *int64              `protobuf:"varint,4,opt,name=atime" json:"atime,omitempty"`
	Replication      *int32              `protobuf:"varint,5,opt,name=replication" json:"replication,omitempty"`
	OwnerName        *string             `protobuf:"bytes,6,opt,name=ownerName" json:"ownerName,omitempty"`
	GroupName        *string             `protobuf:"bytes,7,opt,name=groupName" json:"groupName,omitempty"`
	Perms            *FsPermissionProto  `protobuf:"bytes,8,opt,name=perms" json:"perms,omitempty"`
	Acls             []*AclEntryProto    `protobuf:"bytes,9,rep,name=acls" json:"acls,omitempty"`
	XAttrs           []*XAttrProto       `protobuf:"bytes,10,rep,name=xAttrs" json:"xAttrs,omitempty"`
	XAttrsRemoved    *bool               `protobuf:"varint,11,opt,name=xAttrsRemoved" json:"xAttrsRemoved,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *MetadataUpdateEventProto) Reset()                    { *m = MetadataUpdateEventProto{} }
func (m *MetadataUpdateEventProto) String() string            { return proto.CompactTextString(m) }
func (*MetadataUpdateEventProto) ProtoMessage()               {}
func (*MetadataUpdateEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{7} }

func (m *MetadataUpdateEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *MetadataUpdateEventProto) GetType() MetadataUpdateType {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return MetadataUpdateType_META_TYPE_TIMES
}

func (m *MetadataUpdateEventProto) GetMtime() int64 {
	if m != nil && m.Mtime != nil {
		return *m.Mtime
	}
	return 0
}

func (m *MetadataUpdateEventProto) GetAtime() int64 {
	if m != nil && m.Atime != nil {
		return *m.Atime
	}
	return 0
}

func (m *MetadataUpdateEventProto) GetReplication() int32 {
	if m != nil && m.Replication != nil {
		return *m.Replication
	}
	return 0
}

func (m *MetadataUpdateEventProto) GetOwnerName() string {
	if m != nil && m.OwnerName != nil {
		return *m.OwnerName
	}
	return ""
}

func (m *MetadataUpdateEventProto) GetGroupName() string {
	if m != nil && m.GroupName != nil {
		return *m.GroupName
	}
	return ""
}

func (m *MetadataUpdateEventProto) GetPerms() *FsPermissionProto {
	if m != nil {
		return m.Perms
	}
	return nil
}

func (m *MetadataUpdateEventProto) GetAcls() []*AclEntryProto {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *MetadataUpdateEventProto) GetXAttrs() []*XAttrProto {
	if m != nil {
		return m.XAttrs
	}
	return nil
}

func (m *MetadataUpdateEventProto) GetXAttrsRemoved() bool {
	if m != nil && m.XAttrsRemoved != nil {
		return *m.XAttrsRemoved
	}
	return false
}

type UnlinkEventProto struct {
	Path             *string `protobuf:"bytes,1,req,name=path" json:"path,omitempty"`
	Timestamp        *int64  `protobuf:"varint,2,req,name=timestamp" json:"timestamp,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UnlinkEventProto) Reset()                    { *m = UnlinkEventProto{} }
func (m *UnlinkEventProto) String() string            { return proto.CompactTextString(m) }
func (*UnlinkEventProto) ProtoMessage()               {}
func (*UnlinkEventProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{8} }

func (m *UnlinkEventProto) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

func (m *UnlinkEventProto) GetTimestamp() int64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type EventsListProto struct {
	Events           []*EventProto      `protobuf:"bytes,1,rep,name=events" json:"events,omitempty"`
	FirstTxid        *int64             `protobuf:"varint,2,req,name=firstTxid" json:"firstTxid,omitempty"`
	LastTxid         *int64             `protobuf:"varint,3,req,name=lastTxid" json:"lastTxid,omitempty"`
	SyncTxid         *int64             `protobuf:"varint,4,req,name=syncTxid" json:"syncTxid,omitempty"`
	Batch            []*EventBatchProto `protobuf:"bytes,5,rep,name=batch" json:"batch,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *EventsListProto) Reset()                    { *m = EventsListProto{} }
func (m *EventsListProto) String() string            { return proto.CompactTextString(m) }
func (*EventsListProto) ProtoMessage()               {}
func (*EventsListProto) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{9} }

func (m *EventsListProto) GetEvents() []*EventProto {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *EventsListProto) GetFirstTxid() int64 {
	if m != nil && m.FirstTxid != nil {
		return *m.FirstTxid
	}
	return 0
}

func (m *EventsListProto) GetLastTxid() int64 {
	if m != nil && m.LastTxid != nil {
		return *m.LastTxid
	}
	return 0
}

func (m *EventsListProto) GetSyncTxid() int64 {
	if m != nil && m.SyncTxid != nil {
		return *m.SyncTxid
	}
	return 0
}

func (m *EventsListProto) GetBatch() []*EventBatchProto {
	if m != nil {
		return m.Batch
	}
	return nil
}

func init() {
	proto.RegisterType((*EventProto)(nil), "hadoop.hdfs.EventProto")
	proto.RegisterType((*EventBatchProto)(nil), "hadoop.hdfs.EventBatchProto")
	proto.RegisterType((*CreateEventProto)(nil), "hadoop.hdfs.CreateEventProto")
	proto.RegisterType((*CloseEventProto)(nil), "hadoop.hdfs.CloseEventProto")
	proto.RegisterType((*TruncateEventProto)(nil), "hadoop.hdfs.TruncateEventProto")
	proto.RegisterType((*AppendEventProto)(nil), "hadoop.hdfs.AppendEventProto")
	proto.RegisterType((*RenameEventProto)(nil), "hadoop.hdfs.RenameEventProto")
	proto.RegisterType((*MetadataUpdateEventProto)(nil), "hadoop.hdfs.MetadataUpdateEventProto")
	proto.RegisterType((*UnlinkEventProto)(nil), "hadoop.hdfs.UnlinkEventProto")
	proto.RegisterType((*EventsListProto)(nil), "hadoop.hdfs.EventsListProto")
	proto.RegisterEnum("hadoop.hdfs.EventType", EventType_name, EventType_value)
	proto.RegisterEnum("hadoop.hdfs.INodeType", INodeType_name, INodeType_value)
	proto.RegisterEnum("hadoop.hdfs.MetadataUpdateType", MetadataUpdateType_name, MetadataUpdateType_value)
}

func init() { proto.RegisterFile("inotify.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 917 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x3f, 0x3b, 0x71, 0x2f, 0x99, 0x5c, 0xaf, 0xab, 0xa5, 0x80, 0x89, 0x4e, 0x10, 0x22, 0x90,
	0xa2, 0x4a, 0x04, 0x54, 0x78, 0xe1, 0xde, 0xdc, 0xd4, 0x95, 0x2c, 0x12, 0x37, 0xda, 0xb8, 0xc7,
	0xf5, 0xa9, 0xda, 0xb3, 0x37, 0x8d, 0x75, 0x8e, 0x6d, 0x79, 0xb7, 0x7f, 0xc2, 0x67, 0xe0, 0x03,
	0xf0, 0xc8, 0x33, 0x9f, 0x85, 0xef, 0xc2, 0x57, 0x40, 0xbb, 0xeb, 0xda, 0xf1, 0xe5, 0x50, 0xfb,
	0xc0, 0x9b, 0xe7, 0x37, 0xbf, 0x9d, 0xf9, 0x79, 0x76, 0x66, 0x6c, 0xd8, 0x8f, 0xd3, 0x4c, 0xc4,
	0xcb, 0xcd, 0x38, 0x2f, 0x32, 0x91, 0xe1, 0xde, 0x8a, 0x46, 0x59, 0x96, 0x8f, 0x57, 0xd1, 0x92,
	0xf7, 0xbb, 0x34, 0x4c, 0x34, 0xde, 0xef, 0xdd, 0x53, 0x21, 0x8a, 0xd2, 0x00, 0xe9, 0xd5, 0xcf,
	0xc3, 0x00, 0xc0, 0xbd, 0x65, 0xa9, 0x98, 0xab, 0xe3, 0x47, 0xd0, 0x16, 0x9b, 0x9c, 0xd9, 0xc6,
	0xc0, 0x1c, 0xbd, 0x3c, 0xfe, 0x6c, 0xbc, 0x15, 0x6d, 0xac, 0x68, 0xc1, 0x26, 0x67, 0x44, 0x71,
	0x70, 0x1f, 0x3a, 0x61, 0x96, 0x0a, 0x96, 0x0a, 0x6e, 0x9b, 0x03, 0x73, 0xf4, 0x82, 0x54, 0xf6,
	0xf0, 0x0d, 0x1c, 0x28, 0xfa, 0x09, 0x15, 0xe1, 0x4a, 0x87, 0xc6, 0xd0, 0x16, 0xf7, 0x71, 0xa4,
	0x42, 0xb7, 0x88, 0x7a, 0xc6, 0xdf, 0xc3, 0x1e, 0xbb, 0x2d, 0x03, 0xb4, 0x46, 0xbd, 0xe3, 0xcf,
	0x77, 0x13, 0xaa, 0xc3, 0xa4, 0xa4, 0x0d, 0xff, 0x31, 0x01, 0x4d, 0x0a, 0x46, 0x05, 0x7b, 0xa2,
	0x68, 0xcf, 0xcf, 0x22, 0xb6, 0x25, 0x1a, 0x43, 0x3b, 0xa7, 0x62, 0xa5, 0x04, 0x77, 0x89, 0x7a,
	0xc6, 0x87, 0x60, 0x85, 0x22, 0x5e, 0x33, 0xbb, 0xa5, 0xa4, 0x69, 0x03, 0xbf, 0x82, 0x6e, 0x76,
	0x97, 0xb2, 0xc2, 0xa7, 0x6b, 0x66, 0xb7, 0x15, 0xbd, 0x06, 0xa4, 0xf7, 0xba, 0xc8, 0x6e, 0x72,
	0xe5, 0xb5, 0xb4, 0xb7, 0x02, 0xf0, 0x4f, 0x60, 0xe5, 0xac, 0x58, 0x73, 0x7b, 0x6f, 0x60, 0x8e,
	0x7a, 0xc7, 0x5f, 0x36, 0x24, 0x9d, 0xf1, 0x39, 0x2b, 0xd6, 0x31, 0xe7, 0x71, 0x96, 0xea, 0xb7,
	0xd3, 0x64, 0x3c, 0x80, 0x5e, 0xc1, 0xf2, 0x24, 0x0e, 0xa9, 0x88, 0xb3, 0xd4, 0x7e, 0x3e, 0x30,
	0x46, 0x16, 0xd9, 0x86, 0xf0, 0x37, 0xb0, 0xcf, 0x37, 0xeb, 0x24, 0x4e, 0xdf, 0x07, 0xb4, 0xb8,
	0x66, 0xc2, 0xee, 0x0c, 0x8c, 0x51, 0x97, 0x34, 0x41, 0xa5, 0xfc, 0x96, 0x15, 0x77, 0x45, 0x2c,
	0x98, 0xdd, 0x1d, 0x18, 0xa3, 0x0e, 0xa9, 0x01, 0xfc, 0x1d, 0xa0, 0x88, 0x2d, 0xe9, 0x4d, 0x22,
	0x4e, 0x92, 0x2c, 0x7c, 0xbf, 0x88, 0x7f, 0x63, 0x36, 0x0c, 0x8c, 0x51, 0xeb, 0xb5, 0xf1, 0x03,
	0xd9, 0x71, 0x0d, 0xaf, 0xe0, 0x60, 0x92, 0x64, 0x7c, 0xbb, 0xde, 0x0f, 0x35, 0x34, 0xb6, 0x6a,
	0xd8, 0x87, 0xce, 0x32, 0x4e, 0x98, 0x8a, 0x66, 0xaa, 0x32, 0x56, 0xb6, 0xd4, 0x23, 0x2b, 0xca,
	0x05, 0x5d, 0xe7, 0x65, 0x8d, 0x6b, 0x60, 0xf8, 0x0e, 0x70, 0x50, 0xdc, 0xa4, 0x61, 0xf3, 0x4e,
	0xff, 0xdf, 0x1c, 0x1e, 0x20, 0x27, 0xcf, 0x59, 0x1a, 0x3d, 0x92, 0xe1, 0x6b, 0xe8, 0xa4, 0xec,
	0x4e, 0xbd, 0xbc, 0x6d, 0xca, 0xc2, 0xbd, 0xb6, 0x96, 0x34, 0xe1, 0x8c, 0x54, 0xf0, 0x70, 0x09,
	0x88, 0xb0, 0x94, 0xae, 0xb7, 0xc5, 0xda, 0xf0, 0x9c, 0x17, 0xe1, 0xbc, 0x8e, 0xf6, 0x60, 0x4a,
	0xc9, 0x11, 0xe3, 0x62, 0x5e, 0xb7, 0x5c, 0x65, 0x3f, 0x22, 0xf9, 0xaf, 0x16, 0xd8, 0x33, 0x26,
	0x68, 0x44, 0x05, 0xbd, 0xc8, 0xa3, 0xc7, 0xab, 0xf3, 0x63, 0x39, 0x05, 0xa6, 0x9a, 0x82, 0xaf,
	0x1a, 0x2d, 0xd7, 0x0c, 0xb4, 0x35, 0x0e, 0x87, 0x60, 0xad, 0xcb, 0xd6, 0x37, 0x64, 0xeb, 0x2b,
	0x43, 0xa2, 0x54, 0xa1, 0x6d, 0x8d, 0x2a, 0xe3, 0xc3, 0xf6, 0xb4, 0x76, 0xdb, 0xb3, 0x31, 0x32,
	0x7b, 0xaa, 0x35, 0xff, 0x6b, 0x64, 0x9e, 0x6b, 0xef, 0x47, 0x46, 0x46, 0xb6, 0xf4, 0x93, 0x47,
	0x66, 0x0c, 0x6d, 0x1a, 0x26, 0xdc, 0xee, 0xaa, 0xf5, 0xd1, 0x6f, 0x1c, 0x72, 0xc2, 0xc4, 0x4d,
	0x45, 0xb1, 0xd1, 0x07, 0x14, 0x4f, 0x2e, 0x9c, 0x7b, 0x47, 0x88, 0x82, 0xdb, 0xf0, 0x91, 0x85,
	0xf3, 0x56, 0xba, 0xca, 0x85, 0xa3, 0x69, 0x72, 0xe2, 0xf4, 0x13, 0x61, 0xeb, 0xec, 0x96, 0x45,
	0x76, 0x4f, 0xcd, 0x53, 0x13, 0x1c, 0x9e, 0x02, 0xba, 0x48, 0xe5, 0x04, 0x3e, 0x72, 0x47, 0x8d,
	0x2b, 0x37, 0x3f, 0xbc, 0xf2, 0xbf, 0x8d, 0x72, 0x6b, 0xf2, 0x69, 0xcc, 0xcb, 0x28, 0xf5, 0x86,
	0x34, 0x9e, 0xb4, 0x21, 0x65, 0x8a, 0x65, 0x5c, 0x70, 0x11, 0xc8, 0x5d, 0x5b, 0xa6, 0xa8, 0x00,
	0xd9, 0x8f, 0x09, 0x2d, 0x9d, 0xba, 0xe5, 0x2a, 0x5b, 0xfa, 0xf8, 0x26, 0x0d, 0x95, 0xaf, 0xad,
	0x7d, 0x0f, 0x36, 0x3e, 0x06, 0xeb, 0x9d, 0x5c, 0xe5, 0xb6, 0xa5, 0x54, 0xbc, 0xda, 0x55, 0x51,
	0x6f, 0x7a, 0xa2, 0xa9, 0x47, 0xbf, 0x1b, 0xd0, 0xad, 0xbe, 0x19, 0x18, 0xc1, 0x0b, 0xf7, 0x8d,
	0xeb, 0x07, 0x57, 0x13, 0xe2, 0x3a, 0x81, 0x8b, 0x9e, 0xe1, 0x03, 0xe8, 0x95, 0xc8, 0xf4, 0x7c,
	0xe1, 0x22, 0xa3, 0xa6, 0x38, 0xf3, 0xb9, 0xeb, 0x9f, 0x22, 0xb3, 0x46, 0x88, 0xeb, 0x3b, 0x33,
	0x17, 0xb5, 0x30, 0x86, 0x97, 0x1a, 0x99, 0xb9, 0x81, 0x73, 0xea, 0x04, 0x0e, 0x6a, 0xd7, 0xac,
	0x0b, 0x7f, 0xea, 0xf9, 0xbf, 0x20, 0xab, 0x66, 0x05, 0xe4, 0xc2, 0x9f, 0xc8, 0x74, 0x7b, 0x47,
	0x67, 0xd0, 0xad, 0x3e, 0x06, 0x32, 0xb7, 0x77, 0x15, 0x5c, 0xce, 0xdd, 0xab, 0x33, 0x6f, 0x2a,
	0xc5, 0x1c, 0x02, 0x2a, 0x81, 0x53, 0x8f, 0xb8, 0x93, 0xe0, 0x9c, 0x5c, 0x22, 0x43, 0xc6, 0x29,
	0xd1, 0xc5, 0xe5, 0x4c, 0xc5, 0x36, 0x8f, 0xfe, 0x30, 0x00, 0xef, 0xce, 0x13, 0xfe, 0x04, 0x0e,
	0xa4, 0x24, 0xcd, 0x0e, 0xbc, 0x99, 0xbb, 0x40, 0xcf, 0xf0, 0x17, 0xf0, 0x69, 0x0d, 0x12, 0x77,
	0x3e, 0xf5, 0x26, 0x4e, 0xe0, 0x9d, 0xfb, 0xc8, 0x68, 0xf2, 0xcf, 0x7f, 0xf5, 0x5d, 0x82, 0xcc,
	0x26, 0x38, 0x77, 0xc9, 0x6c, 0xa1, 0x5f, 0xb9, 0x06, 0x9d, 0xc9, 0x74, 0x81, 0xda, 0x52, 0x6e,
	0x8d, 0xbd, 0x75, 0x82, 0x80, 0x2c, 0x90, 0x75, 0xf2, 0x33, 0x7c, 0x9b, 0x15, 0xd7, 0x63, 0x9a,
	0xd3, 0x70, 0xc5, 0x1a, 0x57, 0xa4, 0xbe, 0xf5, 0x61, 0x56, 0xfe, 0x0d, 0x9c, 0xec, 0x7b, 0xfa,
	0xa7, 0x41, 0xdd, 0x17, 0xff, 0xd3, 0x30, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x80, 0xd0, 0x57,
	0x29, 0x47, 0x08, 0x00, 0x00,
}
