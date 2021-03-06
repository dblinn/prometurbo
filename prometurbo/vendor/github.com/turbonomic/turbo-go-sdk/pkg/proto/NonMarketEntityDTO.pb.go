// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NonMarketEntityDTO.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Enumerate non market entities
type NonMarketEntityDTO_NonMarketEntityType int32

const (
	NonMarketEntityDTO_CLOUD_SERVICE    NonMarketEntityDTO_NonMarketEntityType = 0
	NonMarketEntityDTO_ACCOUNT          NonMarketEntityDTO_NonMarketEntityType = 2
	NonMarketEntityDTO_PLAN_DESTINATION NonMarketEntityDTO_NonMarketEntityType = 3
)

var NonMarketEntityDTO_NonMarketEntityType_name = map[int32]string{
	0: "CLOUD_SERVICE",
	2: "ACCOUNT",
	3: "PLAN_DESTINATION",
}

var NonMarketEntityDTO_NonMarketEntityType_value = map[string]int32{
	"CLOUD_SERVICE":    0,
	"ACCOUNT":          2,
	"PLAN_DESTINATION": 3,
}

func (x NonMarketEntityDTO_NonMarketEntityType) Enum() *NonMarketEntityDTO_NonMarketEntityType {
	p := new(NonMarketEntityDTO_NonMarketEntityType)
	*p = x
	return p
}

func (x NonMarketEntityDTO_NonMarketEntityType) String() string {
	return proto.EnumName(NonMarketEntityDTO_NonMarketEntityType_name, int32(x))
}

func (x *NonMarketEntityDTO_NonMarketEntityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_NonMarketEntityType_value, data, "NonMarketEntityDTO_NonMarketEntityType")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_NonMarketEntityType(value)
	return nil
}

func (NonMarketEntityDTO_NonMarketEntityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0}
}

type NonMarketEntityDTO_CloudServiceData_PriceModel int32

const (
	// Pay for compute capacity by per hour or per second depending
	// on which instances are run.
	NonMarketEntityDTO_CloudServiceData_ON_DEMAND NonMarketEntityDTO_CloudServiceData_PriceModel = 0
	// Pay a discounted price for reserved instances with a significant
	// discount (up to 75%) compared to On-Demand instance pricing.
	NonMarketEntityDTO_CloudServiceData_RESERVED NonMarketEntityDTO_CloudServiceData_PriceModel = 1
	// Pay the Spot price that is in effect for the time period the
	// instances are running and prices are set by Amazon EC2 and
	// adjusted gradually based on long-term trends in supply and demand.
	NonMarketEntityDTO_CloudServiceData_SPOT NonMarketEntityDTO_CloudServiceData_PriceModel = 2
)

var NonMarketEntityDTO_CloudServiceData_PriceModel_name = map[int32]string{
	0: "ON_DEMAND",
	1: "RESERVED",
	2: "SPOT",
}

var NonMarketEntityDTO_CloudServiceData_PriceModel_value = map[string]int32{
	"ON_DEMAND": 0,
	"RESERVED":  1,
	"SPOT":      2,
}

func (x NonMarketEntityDTO_CloudServiceData_PriceModel) Enum() *NonMarketEntityDTO_CloudServiceData_PriceModel {
	p := new(NonMarketEntityDTO_CloudServiceData_PriceModel)
	*p = x
	return p
}

func (x NonMarketEntityDTO_CloudServiceData_PriceModel) String() string {
	return proto.EnumName(NonMarketEntityDTO_CloudServiceData_PriceModel_name, int32(x))
}

func (x *NonMarketEntityDTO_CloudServiceData_PriceModel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_CloudServiceData_PriceModel_value, data, "NonMarketEntityDTO_CloudServiceData_PriceModel")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_CloudServiceData_PriceModel(value)
	return nil
}

func (NonMarketEntityDTO_CloudServiceData_PriceModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0}
}

// Types of Billing Data, based on the whether the discovering probe is also
// creating the main Entities or only BillingData or both.
type NonMarketEntityDTO_CloudServiceData_BillingData_Type int32

const (
	// Default implies that the discovering probe is capable of discovering BillingData
	// as well as the Entities being billed. For example, Probe 'X' discovers EntityDTOs
	// *and* BillingData for those EntityDTOs marks the BillingData as DEFAULT.
	NonMarketEntityDTO_CloudServiceData_BillingData_DEFAULT NonMarketEntityDTO_CloudServiceData_BillingData_Type = 0
	// Proxy implies that the discovering probe fetches BillingData but is not capable
	// of discovering Entities for which the BillingData is being discovered. For this
	// type of BillingData, DTOs in virtual_machines need to be marked proxy.
	// Optionally, mainTargetId can be set to the target id for which this discovery
	// probe is the acting proxy.
	// For example, Probe 'X' discovers EntityDTOs, Probe 'Y' discovers BillingData,
	// Probe 'Y' marks BillingData as PROXY.
	NonMarketEntityDTO_CloudServiceData_BillingData_PROXY NonMarketEntityDTO_CloudServiceData_BillingData_Type = 1
	// Stitching implies that the discovering probe is capable of discovering Entities
	// but not the BillingData for the Entities. This type of BillingData would contain
	// stitching information in the virtual_machines list.
	// For example, Probe 'X' discovers EntityDTOs, Probe 'Y' discovers BillingData,
	// Probe 'X' marks BillingData as STITCHING. virtual_machines in this BillingData
	// contain VM Billing Id in the entityProperties list.
	NonMarketEntityDTO_CloudServiceData_BillingData_STITCHING NonMarketEntityDTO_CloudServiceData_BillingData_Type = 2
)

var NonMarketEntityDTO_CloudServiceData_BillingData_Type_name = map[int32]string{
	0: "DEFAULT",
	1: "PROXY",
	2: "STITCHING",
}

var NonMarketEntityDTO_CloudServiceData_BillingData_Type_value = map[string]int32{
	"DEFAULT":   0,
	"PROXY":     1,
	"STITCHING": 2,
}

func (x NonMarketEntityDTO_CloudServiceData_BillingData_Type) Enum() *NonMarketEntityDTO_CloudServiceData_BillingData_Type {
	p := new(NonMarketEntityDTO_CloudServiceData_BillingData_Type)
	*p = x
	return p
}

func (x NonMarketEntityDTO_CloudServiceData_BillingData_Type) String() string {
	return proto.EnumName(NonMarketEntityDTO_CloudServiceData_BillingData_Type_name, int32(x))
}

func (x *NonMarketEntityDTO_CloudServiceData_BillingData_Type) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_CloudServiceData_BillingData_Type_value, data, "NonMarketEntityDTO_CloudServiceData_BillingData_Type")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_CloudServiceData_BillingData_Type(value)
	return nil
}

func (NonMarketEntityDTO_CloudServiceData_BillingData_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0, 0}
}

// The NonMarketEntityDTO message represents an Entity discovered in the target that your probe is
// monitoring and this entity does not participate in the market.
//
// Each entity must have a unique ID to identify it in the Operations Manager repository.
// Many targets provide unique IDs for their entities, or you can generate your own.
// To guarantee that it's unique, you can give the ID a prefix that identifies your
// probe and the given target.
//
// Specify entity type by setting a 'NonMarketEntityType' value to the 'entityType' field.
//
// The 'displayName' value appears in the product GUI and in reports to identify the entity.
//
type NonMarketEntityDTO struct {
	EntityType  *NonMarketEntityDTO_NonMarketEntityType `protobuf:"varint,1,req,name=entityType,enum=common_dto.NonMarketEntityDTO_NonMarketEntityType" json:"entityType,omitempty"`
	Id          *string                                 `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	DisplayName *string                                 `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	Description *string                                 `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// Notifications associated with the entity
	Notification []*NotificationDTO `protobuf:"bytes,5,rep,name=notification" json:"notification,omitempty"`
	// list of <string, string, string> namespace, key, value triplets
	EntityProperties []*EntityDTO_EntityProperty `protobuf:"bytes,6,rep,name=entityProperties" json:"entityProperties,omitempty"`
	// Collection of entity type's specific data.
	//
	// Types that are valid to be assigned to EntityData:
	//	*NonMarketEntityDTO_CloudServiceData_
	//	*NonMarketEntityDTO_PlanDestinationData_
	EntityData           isNonMarketEntityDTO_EntityData `protobuf_oneof:"entity_data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *NonMarketEntityDTO) Reset()         { *m = NonMarketEntityDTO{} }
func (m *NonMarketEntityDTO) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO) ProtoMessage()    {}
func (*NonMarketEntityDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0}
}

func (m *NonMarketEntityDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO.Merge(m, src)
}
func (m *NonMarketEntityDTO) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO.Size(m)
}
func (m *NonMarketEntityDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO proto.InternalMessageInfo

func (m *NonMarketEntityDTO) GetEntityType() NonMarketEntityDTO_NonMarketEntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return NonMarketEntityDTO_CLOUD_SERVICE
}

func (m *NonMarketEntityDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *NonMarketEntityDTO) GetNotification() []*NotificationDTO {
	if m != nil {
		return m.Notification
	}
	return nil
}

func (m *NonMarketEntityDTO) GetEntityProperties() []*EntityDTO_EntityProperty {
	if m != nil {
		return m.EntityProperties
	}
	return nil
}

type isNonMarketEntityDTO_EntityData interface {
	isNonMarketEntityDTO_EntityData()
}

type NonMarketEntityDTO_CloudServiceData_ struct {
	CloudServiceData *NonMarketEntityDTO_CloudServiceData `protobuf:"bytes,500,opt,name=cloud_service_data,json=cloudServiceData,oneof"`
}

type NonMarketEntityDTO_PlanDestinationData_ struct {
	PlanDestinationData *NonMarketEntityDTO_PlanDestinationData `protobuf:"bytes,502,opt,name=plan_destination_data,json=planDestinationData,oneof"`
}

func (*NonMarketEntityDTO_CloudServiceData_) isNonMarketEntityDTO_EntityData() {}

func (*NonMarketEntityDTO_PlanDestinationData_) isNonMarketEntityDTO_EntityData() {}

func (m *NonMarketEntityDTO) GetEntityData() isNonMarketEntityDTO_EntityData {
	if m != nil {
		return m.EntityData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetCloudServiceData() *NonMarketEntityDTO_CloudServiceData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_CloudServiceData_); ok {
		return x.CloudServiceData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetPlanDestinationData() *NonMarketEntityDTO_PlanDestinationData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_PlanDestinationData_); ok {
		return x.PlanDestinationData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NonMarketEntityDTO) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NonMarketEntityDTO_CloudServiceData_)(nil),
		(*NonMarketEntityDTO_PlanDestinationData_)(nil),
	}
}

type NonMarketEntityDTO_CloudServiceData struct {
	CloudProvider        *string                                          `protobuf:"bytes,1,req,name=cloud_provider,json=cloudProvider" json:"cloud_provider,omitempty"`
	BillingData          *NonMarketEntityDTO_CloudServiceData_BillingData `protobuf:"bytes,3,opt,name=billing_data,json=billingData" json:"billing_data,omitempty"`
	AccountId            *string                                          `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	PriceModels          []NonMarketEntityDTO_CloudServiceData_PriceModel `protobuf:"varint,5,rep,name=price_models,json=priceModels,enum=common_dto.NonMarketEntityDTO_CloudServiceData_PriceModel" json:"price_models,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData) Reset()         { *m = NonMarketEntityDTO_CloudServiceData{} }
func (m *NonMarketEntityDTO_CloudServiceData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_CloudServiceData) ProtoMessage()    {}
func (*NonMarketEntityDTO_CloudServiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Merge(m, src)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Size(m)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_CloudServiceData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_CloudServiceData) GetCloudProvider() string {
	if m != nil && m.CloudProvider != nil {
		return *m.CloudProvider
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetBillingData() *NonMarketEntityDTO_CloudServiceData_BillingData {
	if m != nil {
		return m.BillingData
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetPriceModels() []NonMarketEntityDTO_CloudServiceData_PriceModel {
	if m != nil {
		return m.PriceModels
	}
	return nil
}

type NonMarketEntityDTO_CloudServiceData_BillingData struct {
	VirtualMachines   []*EntityDTO                                          `protobuf:"bytes,1,rep,name=virtual_machines,json=virtualMachines" json:"virtual_machines,omitempty"`
	ReservedInstances []*EntityDTO                                          `protobuf:"bytes,2,rep,name=reserved_instances,json=reservedInstances" json:"reserved_instances,omitempty"`
	BillingDataUuid   *string                                               `protobuf:"bytes,3,opt,name=billingDataUuid" json:"billingDataUuid,omitempty"`
	Type              *NonMarketEntityDTO_CloudServiceData_BillingData_Type `protobuf:"varint,4,opt,name=type,enum=common_dto.NonMarketEntityDTO_CloudServiceData_BillingData_Type,def=0" json:"type,omitempty"`
	// Optionally set for Proxy Type BillingData. Represents the target id which is the
	// source of Entities whose BillingData is being fetched by the discovering probe.
	MainTargetId         *string  `protobuf:"bytes,5,opt,name=mainTargetId" json:"mainTargetId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) Reset() {
	*m = NonMarketEntityDTO_CloudServiceData_BillingData{}
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) String() string {
	return proto.CompactTextString(m)
}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) ProtoMessage() {}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Merge(m, src)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Size(m)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData proto.InternalMessageInfo

const Default_NonMarketEntityDTO_CloudServiceData_BillingData_Type NonMarketEntityDTO_CloudServiceData_BillingData_Type = NonMarketEntityDTO_CloudServiceData_BillingData_DEFAULT

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetVirtualMachines() []*EntityDTO {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetReservedInstances() []*EntityDTO {
	if m != nil {
		return m.ReservedInstances
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetBillingDataUuid() string {
	if m != nil && m.BillingDataUuid != nil {
		return *m.BillingDataUuid
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetType() NonMarketEntityDTO_CloudServiceData_BillingData_Type {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_NonMarketEntityDTO_CloudServiceData_BillingData_Type
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetMainTargetId() string {
	if m != nil && m.MainTargetId != nil {
		return *m.MainTargetId
	}
	return ""
}

type NonMarketEntityDTO_PlanDestinationData struct {
	// Set to true if Plan destination has exported data to Azure Migrate Project.
	HasExportedData      *bool    `protobuf:"varint,1,req,name=has_exported_data,json=hasExportedData" json:"has_exported_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonMarketEntityDTO_PlanDestinationData) Reset() {
	*m = NonMarketEntityDTO_PlanDestinationData{}
}
func (m *NonMarketEntityDTO_PlanDestinationData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_PlanDestinationData) ProtoMessage()    {}
func (*NonMarketEntityDTO_PlanDestinationData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 1}
}

func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Merge(m, src)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.Size(m)
}
func (m *NonMarketEntityDTO_PlanDestinationData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_PlanDestinationData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_PlanDestinationData) GetHasExportedData() bool {
	if m != nil && m.HasExportedData != nil {
		return *m.HasExportedData
	}
	return false
}

// Generic cost data struct used for sending cost/spend from probe to the platform.
type CostDataDTO struct {
	// UUID of the entity (Account/CloudService) or profile (VMTemplate/DBTemplate)
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// Whether the cost applies to an entity (default), or a profile.
	AppliesProfile *bool `protobuf:"varint,2,req,name=applies_profile,json=appliesProfile" json:"applies_profile,omitempty"`
	// Type of the cost data entity or the entity that the profile applies to.
	EntityType *EntityDTO_EntityType `protobuf:"varint,3,req,name=entity_type,json=entityType,enum=common_dto.EntityDTO_EntityType" json:"entity_type,omitempty"`
	// Business account id
	AccountId *string `protobuf:"bytes,4,req,name=account_id,json=accountId" json:"account_id,omitempty"`
	// Cost of the associated entity (e.g template/VM etc.)
	// If representing spend (top-down), then how much was spent for this account/service/template,
	// parsed from the bill.
	Cost                 *float32 `protobuf:"fixed32,5,req,name=cost" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostDataDTO) Reset()         { *m = CostDataDTO{} }
func (m *CostDataDTO) String() string { return proto.CompactTextString(m) }
func (*CostDataDTO) ProtoMessage()    {}
func (*CostDataDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{1}
}

func (m *CostDataDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostDataDTO.Unmarshal(m, b)
}
func (m *CostDataDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostDataDTO.Marshal(b, m, deterministic)
}
func (m *CostDataDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostDataDTO.Merge(m, src)
}
func (m *CostDataDTO) XXX_Size() int {
	return xxx_messageInfo_CostDataDTO.Size(m)
}
func (m *CostDataDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_CostDataDTO.DiscardUnknown(m)
}

var xxx_messageInfo_CostDataDTO proto.InternalMessageInfo

func (m *CostDataDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *CostDataDTO) GetAppliesProfile() bool {
	if m != nil && m.AppliesProfile != nil {
		return *m.AppliesProfile
	}
	return false
}

func (m *CostDataDTO) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *CostDataDTO) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *CostDataDTO) GetCost() float32 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return 0
}

func init() {
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_NonMarketEntityType", NonMarketEntityDTO_NonMarketEntityType_name, NonMarketEntityDTO_NonMarketEntityType_value)
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_CloudServiceData_PriceModel", NonMarketEntityDTO_CloudServiceData_PriceModel_name, NonMarketEntityDTO_CloudServiceData_PriceModel_value)
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_CloudServiceData_BillingData_Type", NonMarketEntityDTO_CloudServiceData_BillingData_Type_name, NonMarketEntityDTO_CloudServiceData_BillingData_Type_value)
	proto.RegisterType((*NonMarketEntityDTO)(nil), "common_dto.NonMarketEntityDTO")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData_BillingData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData.BillingData")
	proto.RegisterType((*NonMarketEntityDTO_PlanDestinationData)(nil), "common_dto.NonMarketEntityDTO.PlanDestinationData")
	proto.RegisterType((*CostDataDTO)(nil), "common_dto.CostDataDTO")
}

func init() { proto.RegisterFile("NonMarketEntityDTO.proto", fileDescriptor_a7ef4e88daee47ae) }

var fileDescriptor_a7ef4e88daee47ae = []byte{
	// 831 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x8e, 0xdb, 0x44,
	0x14, 0x5e, 0x3b, 0x49, 0x49, 0x8e, 0xf3, 0xe3, 0x9d, 0xa5, 0x92, 0xb5, 0x08, 0x11, 0x45, 0x20,
	0xa2, 0x4a, 0xa4, 0x52, 0xb8, 0x2b, 0x17, 0x34, 0x1b, 0xbb, 0x34, 0xb0, 0x6b, 0x5b, 0x13, 0x2f,
	0x05, 0xa1, 0x62, 0x66, 0xed, 0xd9, 0xee, 0xa8, 0xb6, 0xc7, 0xb2, 0x27, 0x5b, 0xf6, 0x81, 0x78,
	0x0b, 0x2e, 0x78, 0x25, 0x24, 0xe8, 0x35, 0x9a, 0xb1, 0x37, 0xc9, 0x66, 0xa3, 0x56, 0x15, 0x77,
	0x33, 0xdf, 0x9c, 0x9f, 0xef, 0x9c, 0x39, 0xe7, 0x03, 0xcb, 0xe5, 0xd9, 0x19, 0x29, 0x5e, 0x53,
	0xe1, 0x64, 0x82, 0x89, 0x1b, 0x3b, 0xf0, 0x26, 0x79, 0xc1, 0x05, 0x47, 0x10, 0xf1, 0x34, 0xe5,
	0x59, 0x18, 0x0b, 0x7e, 0x3c, 0x98, 0xab, 0xf3, 0xfa, 0x71, 0xf4, 0x87, 0x01, 0xe8, 0xbe, 0x27,
	0xc2, 0x00, 0x54, 0x5d, 0x82, 0x9b, 0x9c, 0x5a, 0xda, 0x50, 0x1f, 0xf7, 0xa7, 0xd3, 0xc9, 0x26,
	0xd0, 0x64, 0x4f, 0xb6, 0x1d, 0x48, 0x7a, 0xe2, 0xad, 0x28, 0xa8, 0x0f, 0x3a, 0x8b, 0x2d, 0x7d,
	0xa8, 0x8f, 0x3b, 0x58, 0x67, 0x31, 0x1a, 0x82, 0x11, 0xb3, 0x32, 0x4f, 0xc8, 0x8d, 0x4b, 0x52,
	0x6a, 0x35, 0x86, 0xda, 0xb8, 0x83, 0xb7, 0x21, 0x65, 0x41, 0xcb, 0xa8, 0x60, 0xb9, 0x60, 0x3c,
	0xb3, 0x9a, 0xb5, 0xc5, 0x06, 0x42, 0xdf, 0x42, 0x37, 0xe3, 0x82, 0x5d, 0xb2, 0x88, 0x28, 0x93,
	0xd6, 0xb0, 0x31, 0x36, 0xa6, 0x9f, 0xdc, 0x65, 0xba, 0x79, 0xb7, 0x03, 0x0f, 0xdf, 0x71, 0x40,
	0x3e, 0x98, 0x15, 0x45, 0xbf, 0xe0, 0x39, 0x2d, 0x04, 0xa3, 0xa5, 0xf5, 0x40, 0x05, 0xf9, 0x7c,
	0x3b, 0xc8, 0xa6, 0x4a, 0x67, 0xdb, 0xfa, 0x06, 0xdf, 0xf3, 0x46, 0xbf, 0x01, 0x8a, 0x12, 0xbe,
	0x8a, 0xc3, 0x92, 0x16, 0xd7, 0x2c, 0xa2, 0x61, 0x4c, 0x04, 0xb1, 0xfe, 0x91, 0xe5, 0x19, 0xd3,
	0xc7, 0xef, 0xe9, 0xe1, 0x5c, 0x7a, 0x2e, 0x2b, 0x47, 0x9b, 0x08, 0xf2, 0xfc, 0x00, 0x9b, 0xd1,
	0x0e, 0x86, 0x18, 0x3c, 0xcc, 0x13, 0x92, 0x85, 0x31, 0x2d, 0x05, 0xcb, 0x54, 0x1d, 0x55, 0x92,
	0xb7, 0x55, 0x92, 0xf7, 0x7d, 0x94, 0x9f, 0x90, 0xcc, 0xde, 0xf8, 0xd6, 0x79, 0x8e, 0xf2, 0xfb,
	0xf0, 0xf1, 0x5f, 0x2d, 0x30, 0x77, 0x39, 0xa1, 0x2f, 0xa0, 0x5f, 0x55, 0x98, 0x17, 0xfc, 0x9a,
	0xc5, 0xb4, 0x50, 0x03, 0xd2, 0xc1, 0x3d, 0x85, 0xfa, 0x35, 0x88, 0x7e, 0x85, 0xee, 0x05, 0x4b,
	0x12, 0x96, 0xbd, 0xaa, 0xd8, 0x55, 0xe4, 0xbe, 0xf9, 0xc0, 0x0e, 0x4c, 0x4e, 0xaa, 0x18, 0xf2,
	0x8c, 0x8d, 0x8b, 0xcd, 0x05, 0x7d, 0x0a, 0x40, 0xa2, 0x88, 0xaf, 0x32, 0x11, 0xb2, 0xb8, 0x1e,
	0x8e, 0x4e, 0x8d, 0x2c, 0x62, 0xf4, 0x12, 0xba, 0x79, 0x21, 0xfb, 0x9f, 0xf2, 0x98, 0x26, 0xa5,
	0x1a, 0x8d, 0xfe, 0xf4, 0xc9, 0x87, 0xa6, 0xf7, 0x65, 0x8c, 0x33, 0x19, 0x02, 0x1b, 0xf9, 0xfa,
	0x5c, 0x1e, 0xff, 0xad, 0x83, 0xb1, 0x45, 0x0d, 0x3d, 0x05, 0xf3, 0x9a, 0x15, 0x62, 0x45, 0x92,
	0x30, 0x25, 0xd1, 0x15, 0xcb, 0x68, 0x69, 0x69, 0x6a, 0x90, 0x1e, 0xee, 0x1d, 0x24, 0x3c, 0xa8,
	0xcd, 0xcf, 0x6a, 0x6b, 0x64, 0x03, 0x2a, 0xa8, 0x1c, 0x1a, 0x1a, 0x87, 0x2c, 0x2b, 0x05, 0xc9,
	0x22, 0x5a, 0x5a, 0xfa, 0xbb, 0x62, 0x1c, 0xde, 0x3a, 0x2c, 0x6e, 0xed, 0xd1, 0x18, 0x06, 0x5b,
	0x4d, 0x3a, 0x5f, 0xb1, 0xb8, 0xde, 0xac, 0x5d, 0x18, 0xbd, 0x84, 0xa6, 0x90, 0xdb, 0x2d, 0x3b,
	0xd7, 0x9f, 0x3e, 0xfd, 0x1f, 0xff, 0x32, 0x91, 0xfb, 0xfd, 0xe4, 0x23, 0xdb, 0x79, 0x36, 0x3b,
	0x3f, 0x0d, 0xb0, 0x0a, 0x8b, 0x46, 0xd0, 0x4d, 0x09, 0xcb, 0x02, 0x52, 0xbc, 0xa2, 0x62, 0x11,
	0x5b, 0x2d, 0xc5, 0xe2, 0x0e, 0x36, 0xfa, 0x0a, 0x9a, 0x4a, 0x1a, 0x0c, 0xb8, 0x75, 0x36, 0x0f,
	0x50, 0x07, 0x5a, 0x3e, 0xf6, 0x7e, 0xfa, 0xd9, 0xd4, 0x50, 0x0f, 0x3a, 0xcb, 0x60, 0x11, 0xcc,
	0x9f, 0x2f, 0xdc, 0xef, 0x4c, 0x7d, 0xf4, 0x35, 0xc0, 0xe6, 0x3b, 0xe4, 0xa3, 0xe7, 0x86, 0xb6,
	0x73, 0x36, 0x73, 0x6d, 0xf3, 0x00, 0x75, 0xa1, 0x8d, 0x9d, 0xa5, 0x83, 0x7f, 0x74, 0x6c, 0x53,
	0x43, 0x6d, 0x68, 0x2e, 0x7d, 0x2f, 0x30, 0xf5, 0xe3, 0x19, 0x1c, 0xed, 0x19, 0x78, 0xf4, 0x08,
	0x0e, 0xaf, 0x48, 0x19, 0xd2, 0xdf, 0x73, 0x5e, 0x08, 0x1a, 0x57, 0x23, 0x2a, 0xe7, 0xb8, 0x8d,
	0x07, 0x57, 0xa4, 0x74, 0x6a, 0x5c, 0xda, 0x8e, 0x7e, 0x81, 0xa3, 0x3d, 0xe2, 0x86, 0x0e, 0xa1,
	0x37, 0x3f, 0xf5, 0xce, 0xed, 0x50, 0x66, 0x5d, 0xcc, 0x1d, 0xf3, 0x40, 0x16, 0x32, 0x9b, 0xcf,
	0xbd, 0x73, 0x37, 0x30, 0x75, 0xf4, 0x31, 0x98, 0xfe, 0xe9, 0x4c, 0x52, 0x5c, 0x06, 0x0b, 0x77,
	0x16, 0x2c, 0x3c, 0xd7, 0x6c, 0x8c, 0x9a, 0x6d, 0xcd, 0xd4, 0x1e, 0xb5, 0x5f, 0x78, 0xf8, 0x87,
	0x67, 0xa7, 0xde, 0x8b, 0x93, 0x1e, 0x18, 0x95, 0x86, 0x28, 0x0a, 0xdf, 0x3f, 0x68, 0xff, 0xdb,
	0x30, 0xdf, 0x36, 0x70, 0xef, 0x0d, 0x2f, 0x5e, 0x5f, 0x26, 0xfc, 0x8d, 0x82, 0x47, 0x7f, 0x6a,
	0x60, 0xcc, 0x79, 0x29, 0x24, 0x1f, 0x29, 0xd0, 0x95, 0x98, 0x6a, 0x6b, 0x31, 0xfd, 0x12, 0x06,
	0x24, 0xcf, 0x13, 0x46, 0x4b, 0xb9, 0x95, 0x97, 0x2c, 0xa1, 0x4a, 0x69, 0xdb, 0xb8, 0x5f, 0xc3,
	0x7e, 0x85, 0xa2, 0xd9, 0x3a, 0x9d, 0xfa, 0xfc, 0x86, 0x92, 0xf6, 0xe1, 0xbb, 0xb4, 0xee, 0x9e,
	0x90, 0xef, 0x2e, 0x9e, 0x7e, 0x77, 0xf1, 0x10, 0x34, 0x23, 0x5e, 0x0a, 0xab, 0x35, 0xd4, 0xc7,
	0x3a, 0x56, 0xe7, 0x93, 0xc7, 0xf0, 0x59, 0xc4, 0xd3, 0xc9, 0x75, 0x2a, 0x56, 0xc5, 0x05, 0x9f,
	0xe4, 0x09, 0x11, 0x97, 0xbc, 0x48, 0xeb, 0xb4, 0x93, 0x58, 0xf0, 0x93, 0xee, 0xba, 0xc5, 0x76,
	0xe0, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0xce, 0xe5, 0xa3, 0xde, 0xcf, 0x06, 0x00, 0x00,
}
