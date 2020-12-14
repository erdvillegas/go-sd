package caservice

import (
	"encoding/xml"
	"time"

	"github.com/hooklift/gowsdl/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type CreateObject struct {
	XMLName            xml.Name       `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createObject"`
	Sid                int32          `xml:"sid,omitempty"`
	ObjectType         string         `xml:"objectType,omitempty"`
	AttrVals           *ArrayOfString `xml:"attrVals,omitempty"`
	Attributes         *ArrayOfString `xml:"attributes,omitempty"`
	CreateObjectResult string         `xml:"createObjectResult,omitempty"`
	NewHandle          string         `xml:"newHandle,omitempty"`
}

type CreateObjectResponse struct {
	XMLName            xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createObjectResponse"`
	CreateObjectResult string   `xml:"createObjectResult,omitempty"`
	NewHandle          string   `xml:"newHandle,omitempty"`
}

type AddAssetLog struct {
	XMLName       xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addAssetLog"`
	Sid           int32    `xml:"sid,omitempty"`
	AssetHandle   string   `xml:"assetHandle,omitempty"`
	ContactHandle string   `xml:"contactHandle,omitempty"`
	LogText       string   `xml:"logText,omitempty"`
}

type AddAssetLogResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addAssetLogResponse"`
}

type CreateLrelRelationships struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createLrelRelationships"`

	Sid int32 `xml:"sid,omitempty"`

	ContextObject string `xml:"contextObject,omitempty"`

	LrelName string `xml:"lrelName,omitempty"`

	AddObjectHandles *ArrayOfString `xml:"addObjectHandles,omitempty"`
}

type CreateLrelRelationshipsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createLrelRelationshipsResponse"`
}

type AddMemberToGroup struct {
	XMLName       xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addMemberToGroup"`
	Sid           int32    `xml:"sid,omitempty"`
	ContactHandle string   `xml:"contactHandle,omitempty"`
	GroupHandle   string   `xml:"groupHandle,omitempty"`
}

type AddMemberToGroupResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addMemberToGroupResponse"`
}

type AttachChangeToRequest struct {
	XMLName        xml.Name       `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attachChangeToRequest"`
	Sid            int32          `xml:"sid,omitempty"`
	Creator        string         `xml:"creator,omitempty"`
	RequestHandle  string         `xml:"requestHandle,omitempty"`
	ChangeHandle   string         `xml:"changeHandle,omitempty"`
	ChangeAttrVals *ArrayOfString `xml:"changeAttrVals,omitempty"`
	Description    string         `xml:"description,omitempty"`
}

type AttachChangeToRequestResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attachChangeToRequestResponse"`

	AttachChangeToRequestReturn string `xml:"attachChangeToRequestReturn,omitempty"`
}

type CallServerMethod struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk callServerMethod"`

	Sid int32 `xml:"sid,omitempty"`

	MethodName string `xml:"methodName,omitempty"`

	FactoryName string `xml:"factoryName,omitempty"`

	FormatList string `xml:"formatList,omitempty"`

	Parameters *ArrayOfString `xml:"parameters,omitempty"`
}

type CallServerMethodResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk callServerMethodResponse"`

	CallServerMethodReturn string `xml:"callServerMethodReturn,omitempty"`
}

type ChangeStatus struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk changeStatus"`

	Sid int32 `xml:"sid,omitempty"`

	Creator string `xml:"creator,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	NewStatusHandle string `xml:"newStatusHandle,omitempty"`
}

type ChangeStatusResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk changeStatusResponse"`

	ChangeStatusReturn string `xml:"changeStatusReturn,omitempty"`
}

type ClearNotification struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk clearNotification"`

	Sid int32 `xml:"sid,omitempty"`

	LrObject string `xml:"lrObject,omitempty"`

	ClearBy string `xml:"clearBy,omitempty"`
}

type ClearNotificationResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk clearNotificationResponse"`

	ClearNotificationReturn int32 `xml:"clearNotificationReturn,omitempty"`
}

type CreateActivityLog struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createActivityLog"`

	Sid int32 `xml:"sid,omitempty"`

	Creator string `xml:"creator,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	LogType string `xml:"logType,omitempty"`

	TimeSpent int32 `xml:"timeSpent,omitempty"`

	Internal bool `xml:"internal,omitempty"`
}

type CreateActivityLogResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createActivityLogResponse"`

	CreateActivityLogReturn string `xml:"createActivityLogReturn,omitempty"`
}

type CreateAsset struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAsset"`

	Sid int32 `xml:"sid,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	CreateAssetResult string `xml:"createAssetResult,omitempty"`

	NewAssetHandle string `xml:"newAssetHandle,omitempty"`

	NewExtensionHandle string `xml:"newExtensionHandle,omitempty"`

	NewExtensionName string `xml:"newExtensionName,omitempty"`
}

type CreateAssetResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAssetResponse"`

	CreateAssetResult string `xml:"createAssetResult,omitempty"`

	NewAssetHandle string `xml:"newAssetHandle,omitempty"`

	NewExtensionHandle string `xml:"newExtensionHandle,omitempty"`

	NewExtensionName string `xml:"newExtensionName,omitempty"`
}

type CreateAssetParentChildRelationship struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAssetParentChildRelationship"`

	Sid int32 `xml:"sid,omitempty"`

	ParentHandle string `xml:"parentHandle,omitempty"`

	ChildHandle string `xml:"childHandle,omitempty"`
}

type CreateAssetParentChildRelationshipResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAssetParentChildRelationshipResponse"`

	CreateAssetParentChildRelationshipReturn string `xml:"createAssetParentChildRelationshipReturn,omitempty"`
}

type CreateAttachment struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAttachment"`

	Sid int32 `xml:"sid,omitempty"`

	RepositoryHandle string `xml:"repositoryHandle,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	FileName string `xml:"fileName,omitempty"`
}

type CreateAttachmentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAttachmentResponse"`

	CreateAttachmentReturn string `xml:"createAttachmentReturn,omitempty"`
}

type RemoveAttachment struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk removeAttachment"`

	Sid int32 `xml:"sid,omitempty"`

	AttHandle string `xml:"attHandle,omitempty"`
}

type RemoveAttachmentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk removeAttachmentResponse"`

	RemoveAttachmentReturn int32 `xml:"removeAttachmentReturn,omitempty"`
}

type CreateChangeOrder struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createChangeOrder"`

	Sid int32 `xml:"sid,omitempty"`

	CreatorHandle string `xml:"creatorHandle,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`

	PropertyValues *ArrayOfString `xml:"propertyValues,omitempty"`

	Template string `xml:"template,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	NewChangeHandle string `xml:"newChangeHandle,omitempty"`

	NewChangeNumber string `xml:"newChangeNumber,omitempty"`
}

type CreateChangeOrderResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createChangeOrderResponse"`

	CreateChangeOrderReturn string `xml:"createChangeOrderReturn,omitempty"`

	NewChangeHandle string `xml:"newChangeHandle,omitempty"`

	NewChangeNumber string `xml:"newChangeNumber,omitempty"`
}

type CreateIssue struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createIssue"`

	Sid int32 `xml:"sid,omitempty"`

	CreatorHandle string `xml:"creatorHandle,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`

	PropertyValues *ArrayOfString `xml:"propertyValues,omitempty"`

	Template string `xml:"template,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	NewIssueHandle string `xml:"newIssueHandle,omitempty"`

	NewIssueNumber string `xml:"newIssueNumber,omitempty"`
}

type CreateIssueResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createIssueResponse"`

	CreateIssueReturn string `xml:"createIssueReturn,omitempty"`

	NewIssueHandle string `xml:"newIssueHandle,omitempty"`

	NewIssueNumber string `xml:"newIssueNumber,omitempty"`
}

type CreateWorkFlowTask struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createWorkFlowTask"`

	Sid int32 `xml:"sid,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	CreatorHandle string `xml:"creatorHandle,omitempty"`

	SelectedWorkFlow string `xml:"selectedWorkFlow,omitempty"`

	TaskType string `xml:"taskType,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	CreateWorkFlowTaskResult string `xml:"createWorkFlowTaskResult,omitempty"`

	NewHandle string `xml:"newHandle,omitempty"`
}

type CreateWorkFlowTaskResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createWorkFlowTaskResponse"`

	CreateWorkFlowTaskResult string `xml:"createWorkFlowTaskResult,omitempty"`

	NewHandle string `xml:"newHandle,omitempty"`
}

type DeleteWorkFlowTask struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteWorkFlowTask"`

	Sid int32 `xml:"sid,omitempty"`

	WorkFlowHandle string `xml:"workFlowHandle,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`
}

type DeleteWorkFlowTaskResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteWorkFlowTaskResponse"`
}

type CreateRequest struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createRequest"`

	Sid int32 `xml:"sid,omitempty"`

	CreatorHandle string `xml:"creatorHandle,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`

	PropertyValues *ArrayOfString `xml:"propertyValues,omitempty"`

	Template string `xml:"template,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	NewRequestHandle string `xml:"newRequestHandle,omitempty"`

	NewRequestNumber string `xml:"newRequestNumber,omitempty"`
}

type CreateRequestResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createRequestResponse"`

	CreateRequestReturn string `xml:"createRequestReturn,omitempty"`

	NewRequestHandle string `xml:"newRequestHandle,omitempty"`

	NewRequestNumber string `xml:"newRequestNumber,omitempty"`
}

type CreateTicket struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createTicket"`

	Sid int32 `xml:"sid,omitempty"`

	Description string `xml:"description,omitempty"`

	Problem_type string `xml:"problem_type,omitempty"`

	Userid string `xml:"userid,omitempty"`

	Asset string `xml:"asset,omitempty"`

	Duplication_id string `xml:"duplication_id,omitempty"`

	NewTicketHandle string `xml:"newTicketHandle,omitempty"`

	NewTicketNumber string `xml:"newTicketNumber,omitempty"`

	ReturnUserData string `xml:"returnUserData,omitempty"`

	ReturnApplicationData string `xml:"returnApplicationData,omitempty"`
}

type CreateTicketResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createTicketResponse"`

	CreateTicketReturn string `xml:"createTicketReturn,omitempty"`

	NewTicketHandle string `xml:"newTicketHandle,omitempty"`

	NewTicketNumber string `xml:"newTicketNumber,omitempty"`

	ReturnUserData string `xml:"returnUserData,omitempty"`

	ReturnApplicationData string `xml:"returnApplicationData,omitempty"`
}

type CreateQuickTicket struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createQuickTicket"`

	Sid int32 `xml:"sid,omitempty"`

	CustomerHandle string `xml:"customerHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	NewTicketHandle string `xml:"newTicketHandle,omitempty"`

	NewTicketNumber string `xml:"newTicketNumber,omitempty"`
}

type CreateQuickTicketResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createQuickTicketResponse"`

	CreateQuickTicketReturn string `xml:"createQuickTicketReturn,omitempty"`

	NewTicketHandle string `xml:"newTicketHandle,omitempty"`

	NewTicketNumber string `xml:"newTicketNumber,omitempty"`
}

type CloseTicket struct {
	XMLName      xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk closeTicket"`
	Sid          int32    `xml:"sid,omitempty"`
	Description  string   `xml:"description,omitempty"`
	TicketHandle string   `xml:"ticketHandle,omitempty"`
}

type CloseTicketResponse struct {
	XMLName           xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk closeTicketResponse"`
	CloseTicketReturn string   `xml:"closeTicketReturn,omitempty"`
}

type LogComment struct {
	XMLName      xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk logComment"`
	Sid          int32    `xml:"sid,omitempty"`
	TicketHandle string   `xml:"ticketHandle,omitempty"`
	Comment      string   `xml:"comment,omitempty"`
	InternalFlag int32    `xml:"internalFlag,omitempty"`
}

type LogCommentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk logCommentResponse"`
}

type GetPolicyInfo struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPolicyInfo"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetPolicyInfoResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPolicyInfoResponse"`

	GetPolicyInfoReturn string `xml:"getPolicyInfoReturn,omitempty"`
}

type DetachChangeFromRequest struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk detachChangeFromRequest"`

	Sid int32 `xml:"sid,omitempty"`

	Creator string `xml:"creator,omitempty"`

	RequestHandle string `xml:"requestHandle,omitempty"`

	Description string `xml:"description,omitempty"`
}

type DetachChangeFromRequestResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk detachChangeFromRequestResponse"`

	DetachChangeFromRequestReturn string `xml:"detachChangeFromRequestReturn,omitempty"`
}

type DoSelect struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk doSelect"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectType string `xml:"objectType,omitempty"`

	WhereClause string `xml:"whereClause,omitempty"`

	MaxRows int32 `xml:"maxRows,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type DoSelectResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk doSelectResponse"`

	DoSelectReturn string `xml:"doSelectReturn,omitempty"`
}

type DoQuery struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk doQuery"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectType string `xml:"objectType,omitempty"`

	WhereClause string `xml:"whereClause,omitempty"`
}

type DoQueryResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk doQueryResponse"`

	DoQueryReturn *ListResult `xml:"doQueryReturn,omitempty"`
}

type Escalate struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk escalate"`

	Sid int32 `xml:"sid,omitempty"`

	Creator string `xml:"creator,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	SetAssignee bool `xml:"setAssignee,omitempty"`

	NewAssigneeHandle string `xml:"newAssigneeHandle,omitempty"`

	SetGroup bool `xml:"setGroup,omitempty"`

	NewGroupHandle string `xml:"newGroupHandle,omitempty"`

	SetOrganization bool `xml:"setOrganization,omitempty"`

	NewOrganizationHandle string `xml:"newOrganizationHandle,omitempty"`

	SetPriority bool `xml:"setPriority,omitempty"`

	NewPriorityHandle string `xml:"newPriorityHandle,omitempty"`
}

type EscalateResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk escalateResponse"`

	EscalateReturn string `xml:"escalateReturn,omitempty"`
}

type FreeListHandles struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk freeListHandles"`

	Sid int32 `xml:"sid,omitempty"`

	Handles *ArrayOfInt `xml:"handles,omitempty"`
}

type FreeListHandlesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk freeListHandlesResponse"`
}

type GetAssetExtensionInformation struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAssetExtensionInformation"`

	Sid int32 `xml:"sid,omitempty"`

	AssetHandle string `xml:"assetHandle,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	GetAssetExtInfoResult string `xml:"getAssetExtInfoResult,omitempty"`

	ExtensionHandle string `xml:"extensionHandle,omitempty"`

	ExtensionName string `xml:"extensionName,omitempty"`
}

type GetAssetExtensionInformationResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAssetExtensionInformationResponse"`

	GetAssetExtInfoResult string `xml:"getAssetExtInfoResult,omitempty"`

	ExtensionHandle string `xml:"extensionHandle,omitempty"`

	ExtensionName string `xml:"extensionName,omitempty"`
}

type GetConfigurationMode struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getConfigurationMode"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetConfigurationModeResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getConfigurationModeResponse"`

	GetConfigurationModeReturn string `xml:"getConfigurationModeReturn,omitempty"`
}

type GetGroupMemberListValues struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getGroupMemberListValues"`

	Sid int32 `xml:"sid,omitempty"`

	WhereClause string `xml:"whereClause,omitempty"`

	NumToFetch int32 `xml:"numToFetch,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetGroupMemberListValuesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getGroupMemberListValuesResponse"`

	GetGroupMemberListValuesReturn string `xml:"getGroupMemberListValuesReturn,omitempty"`
}

type GetObjectTypeInformation struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getObjectTypeInformation"`

	Sid int32 `xml:"sid,omitempty"`

	Factory string `xml:"factory,omitempty"`
}

type GetObjectTypeInformationResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getObjectTypeInformationResponse"`

	GetObjectTypeInformationReturn string `xml:"getObjectTypeInformationReturn,omitempty"`
}

type GetHandleForUserid struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getHandleForUserid"`

	Sid int32 `xml:"sid,omitempty"`

	UserID string `xml:"userID,omitempty"`
}

type GetHandleForUseridResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getHandleForUseridResponse"`

	GetHandleForUseridReturn string `xml:"getHandleForUseridReturn,omitempty"`
}

type GetAccessTypeForContact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAccessTypeForContact"`

	Sid int32 `xml:"sid,omitempty"`

	ContactHandle string `xml:"contactHandle,omitempty"`
}

type GetAccessTypeForContactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAccessTypeForContactResponse"`

	GetAccessTypeForContactReturn string `xml:"getAccessTypeForContactReturn,omitempty"`
}

type GetListValues struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getListValues"`

	Sid int32 `xml:"sid,omitempty"`

	ListHandle int32 `xml:"listHandle,omitempty"`

	StartIndex int32 `xml:"startIndex,omitempty"`

	EndIndex int32 `xml:"endIndex,omitempty"`

	AttributeNames *ArrayOfString `xml:"attributeNames,omitempty"`
}

type GetListValuesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getListValuesResponse"`

	GetListValuesReturn string `xml:"getListValuesReturn,omitempty"`
}

type GetLrelLength struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getLrelLength"`

	Sid int32 `xml:"sid,omitempty"`

	ContextObject string `xml:"contextObject,omitempty"`

	LrelName string `xml:"lrelName,omitempty"`
}

type GetLrelLengthResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getLrelLengthResponse"`

	GetLrelLengthReturn int32 `xml:"getLrelLengthReturn,omitempty"`
}

type GetLrelValues struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getLrelValues"`

	Sid int32 `xml:"sid,omitempty"`

	ContextObject string `xml:"contextObject,omitempty"`

	LrelName string `xml:"lrelName,omitempty"`

	StartIndex int32 `xml:"startIndex,omitempty"`

	EndIndex int32 `xml:"endIndex,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetLrelValuesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getLrelValuesResponse"`

	GetLrelValuesReturn string `xml:"getLrelValuesReturn,omitempty"`
}

type GetNotificationsForContact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getNotificationsForContact"`

	Sid int32 `xml:"sid,omitempty"`

	ContactHandle string `xml:"contactHandle,omitempty"`

	QueryStatus int32 `xml:"queryStatus,omitempty"`
}

type GetNotificationsForContactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getNotificationsForContactResponse"`

	GetNotificationsForContactReturn *ListResult `xml:"getNotificationsForContactReturn,omitempty"`
}

type GetObjectValues struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getObjectValues"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetObjectValuesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getObjectValuesResponse"`

	GetObjectValuesReturn string `xml:"getObjectValuesReturn,omitempty"`
}

type GetPendingChangeTaskListForContact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPendingChangeTaskListForContact"`

	Sid int32 `xml:"sid,omitempty"`

	ContactHandle string `xml:"contactHandle,omitempty"`
}

type GetPendingChangeTaskListForContactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPendingChangeTaskListForContactResponse"`

	GetPendingChangeTaskListForContactReturn *ListResult `xml:"getPendingChangeTaskListForContactReturn,omitempty"`
}

type GetPendingIssueTaskListForContact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPendingIssueTaskListForContact"`

	Sid int32 `xml:"sid,omitempty"`

	ContactHandle string `xml:"contactHandle,omitempty"`
}

type GetPendingIssueTaskListForContactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPendingIssueTaskListForContactResponse"`

	GetPendingIssueTaskListForContactReturn *ListResult `xml:"getPendingIssueTaskListForContactReturn,omitempty"`
}

type GetPropertyInfoForCategory struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPropertyInfoForCategory"`

	Sid int32 `xml:"sid,omitempty"`

	CategoryHandle string `xml:"categoryHandle,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetPropertyInfoForCategoryResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPropertyInfoForCategoryResponse"`

	GetPropertyInfoForCategoryReturn string `xml:"getPropertyInfoForCategoryReturn,omitempty"`
}

type GetRelatedList struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getRelatedList"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	ListName string `xml:"listName,omitempty"`
}

type GetRelatedListResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getRelatedListResponse"`

	GetRelatedListReturn *ListResult `xml:"getRelatedListReturn,omitempty"`
}

type GetRelatedListValues struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getRelatedListValues"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	ListName string `xml:"listName,omitempty"`

	NumToFetch int32 `xml:"numToFetch,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	GetRelatedListValuesResult string `xml:"getRelatedListValuesResult,omitempty"`

	NumRowsFound int32 `xml:"numRowsFound,omitempty"`
}

type GetRelatedListValuesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getRelatedListValuesResponse"`

	GetRelatedListValuesResult string `xml:"getRelatedListValuesResult,omitempty"`

	NumRowsFound int32 `xml:"numRowsFound,omitempty"`
}

type GetWorkFlowTemplates struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getWorkFlowTemplates"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetWorkFlowTemplatesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getWorkFlowTemplatesResponse"`

	GetWorkFlowTemplatesReturn string `xml:"getWorkFlowTemplatesReturn,omitempty"`
}

type GetTaskListValues struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getTaskListValues"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetTaskListValuesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getTaskListValuesResponse"`

	GetTaskListValuesReturn string `xml:"getTaskListValuesReturn,omitempty"`
}

type GetValidTaskTransitions struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getValidTaskTransitions"`

	Sid int32 `xml:"sid,omitempty"`

	TaskHandle string `xml:"taskHandle,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type GetValidTaskTransitionsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getValidTaskTransitionsResponse"`

	GetValidTaskTransitionsReturn string `xml:"getValidTaskTransitionsReturn,omitempty"`
}

type GetValidTransitions struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getValidTransitions"`

	Sid int32 `xml:"sid,omitempty"`

	Handle string `xml:"handle,omitempty"`

	TicketFactory string `xml:"ticketFactory,omitempty"`
}

type GetValidTransitionsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getValidTransitionsResponse"`

	GetValidTransitionsReturn string `xml:"getValidTransitionsReturn,omitempty"`
}

type GetDependentAttrControls struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDependentAttrControls"`

	Sid int32 `xml:"sid,omitempty"`

	Handle string `xml:"handle,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`
}

type GetDependentAttrControlsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDependentAttrControlsResponse"`

	GetDependentAttrControlsReturn string `xml:"getDependentAttrControlsReturn,omitempty"`
}

type Login struct {
	XMLName  xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk login"`
	Username string   `xml:"username,omitempty"`
	Password string   `xml:"password,omitempty"`
}

type LoginResponse struct {
	XMLName     xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginResponse"`
	LoginReturn int32    `xml:"loginReturn,omitempty"`
}

type LoginService struct {
	XMLName  xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginService"`
	Username string   `xml:"username,omitempty"`
	Password string   `xml:"password,omitempty"`
	Policy   string   `xml:"policy,omitempty"`
}

type LoginServiceResponse struct {
	XMLName            xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginServiceResponse"`
	LoginServiceReturn int32    `xml:"loginServiceReturn,omitempty"`
}

type LoginServiceManaged struct {
	XMLName          xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginServiceManaged"`
	Policy           string   `xml:"policy,omitempty"`
	Encrypted_policy string   `xml:"encrypted_policy,omitempty"`
}

type LoginServiceManagedResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginServiceManagedResponse"`

	LoginServiceManagedReturn string `xml:"loginServiceManagedReturn,omitempty"`
}

type LoginWithArtifact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginWithArtifact"`

	Userid string `xml:"userid,omitempty"`

	Artifact string `xml:"artifact,omitempty"`
}

type LoginWithArtifactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk loginWithArtifactResponse"`

	LoginWithArtifactReturn int32 `xml:"loginWithArtifactReturn,omitempty"`
}

type Impersonate struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk impersonate"`

	Sid int32 `xml:"sid,omitempty"`

	Userid string `xml:"userid,omitempty"`
}

type ImpersonateResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk impersonateResponse"`

	ImpersonateReturn int32 `xml:"impersonateReturn,omitempty"`
}

type Logout struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk logout"`

	Sid int32 `xml:"sid,omitempty"`
}

type LogoutResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk logoutResponse"`
}

type NotifyContacts struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk notifyContacts"`

	Sid int32 `xml:"sid,omitempty"`

	Creator string `xml:"creator,omitempty"`

	ContextObject string `xml:"contextObject,omitempty"`

	MessageTitle string `xml:"messageTitle,omitempty"`

	MessageBody string `xml:"messageBody,omitempty"`

	NotifyLevel int32 `xml:"notifyLevel,omitempty"`

	Notifyees *ArrayOfString `xml:"notifyees,omitempty"`

	Internal bool `xml:"internal,omitempty"`
}

type NotifyContactsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk notifyContactsResponse"`

	NotifyContactsReturn string `xml:"notifyContactsReturn,omitempty"`
}

type RemoveLrelRelationships struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk removeLrelRelationships"`

	Sid int32 `xml:"sid,omitempty"`

	ContextObject string `xml:"contextObject,omitempty"`

	LrelName string `xml:"lrelName,omitempty"`

	RemoveObjectHandles *ArrayOfString `xml:"removeObjectHandles,omitempty"`
}

type RemoveLrelRelationshipsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk removeLrelRelationshipsResponse"`
}

type RemoveMemberFromGroup struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk removeMemberFromGroup"`

	Sid int32 `xml:"sid,omitempty"`

	ContactHandle string `xml:"contactHandle,omitempty"`

	GroupHandle string `xml:"groupHandle,omitempty"`
}

type RemoveMemberFromGroupResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk removeMemberFromGroupResponse"`
}

type ServerStatus struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk serverStatus"`

	Sid int32 `xml:"sid,omitempty"`
}

type ServerStatusResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk serverStatusResponse"`

	ServerStatusReturn int32 `xml:"serverStatusReturn,omitempty"`
}

type UpdateObject struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk updateObject"`

	Sid int32 `xml:"sid,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	AttrVals *ArrayOfString `xml:"attrVals,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`
}

type UpdateObjectResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk updateObjectResponse"`

	UpdateObjectReturn string `xml:"updateObjectReturn,omitempty"`
}

type GetBopsid struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getBopsid"`

	Sid int32 `xml:"sid,omitempty"`

	Contact string `xml:"contact,omitempty"`
}

type GetBopsidResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getBopsidResponse"`

	GetBopsidReturn string `xml:"getBopsidReturn,omitempty"`
}

type GetArtifact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getArtifact"`

	Sid int32 `xml:"sid,omitempty"`

	Contact string `xml:"contact,omitempty"`

	Passwd string `xml:"passwd,omitempty"`
}

type GetArtifactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getArtifactResponse"`

	GetArtifactReturn string `xml:"getArtifactReturn,omitempty"`
}

type AttachURLLinkToTicket struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attachURLLinkToTicket"`

	Sid int32 `xml:"sid,omitempty"`

	TicketHandle string `xml:"ticketHandle,omitempty"`

	Url string `xml:"url,omitempty"`

	AttmntName string `xml:"attmntName,omitempty"`

	Description string `xml:"description,omitempty"`
}

type AttachURLLinkToTicketResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attachURLLinkToTicketResponse"`

	AttachURLLinkToTicketReturn string `xml:"attachURLLinkToTicketReturn,omitempty"`
}

type CreateAttmnt struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAttmnt"`

	Sid int32 `xml:"sid,omitempty"`

	RepositoryHandle string `xml:"repositoryHandle,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`

	ObjectHandle int32 `xml:"objectHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	FileName string `xml:"fileName,omitempty"`
}

type CreateAttmntResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createAttmntResponse"`

	CreateAttmntReturn string `xml:"createAttmntReturn,omitempty"`
}

type GetDocumentsByIDs struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDocumentsByIDs"`

	Sid int32 `xml:"sid,omitempty"`

	DocIds string `xml:"docIds,omitempty"`

	PropertyList string `xml:"propertyList,omitempty"`

	SortBy string `xml:"sortBy,omitempty"`

	Descending bool `xml:"descending,omitempty"`
}

type GetDocumentsByIDsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDocumentsByIDsResponse"`

	GetDocumentsByIDsReturn string `xml:"getDocumentsByIDsReturn,omitempty"`
}

type GetDecisionTrees struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDecisionTrees"`

	Sid int32 `xml:"sid,omitempty"`

	PropertyList string `xml:"propertyList,omitempty"`

	SortBy string `xml:"sortBy,omitempty"`

	Descending bool `xml:"descending,omitempty"`
}

type GetDecisionTreesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDecisionTreesResponse"`

	GetDecisionTreesReturn string `xml:"getDecisionTreesReturn,omitempty"`
}

type DeleteDocument struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteDocument"`

	Sid int32 `xml:"sid,omitempty"`

	DocId int32 `xml:"docId,omitempty"`
}

type DeleteDocumentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteDocumentResponse"`

	DeleteDocumentReturn int32 `xml:"deleteDocumentReturn,omitempty"`
}

type GetCategory struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getCategory"`

	Sid int32 `xml:"sid,omitempty"`

	CatId int32 `xml:"catId,omitempty"`

	GetCategoryPaths bool `xml:"getCategoryPaths,omitempty"`
}

type GetCategoryResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getCategoryResponse"`

	GetCategoryReturn string `xml:"getCategoryReturn,omitempty"`
}

type GetStatuses struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getStatuses"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetStatusesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getStatusesResponse"`

	GetStatusesReturn string `xml:"getStatusesReturn,omitempty"`
}

type GetBookmarks struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getBookmarks"`

	Sid int32 `xml:"sid,omitempty"`

	ContactId string `xml:"contactId,omitempty"`
}

type GetBookmarksResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getBookmarksResponse"`

	GetBookmarksReturn string `xml:"getBookmarksReturn,omitempty"`
}

type GetQuestionsAsked struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getQuestionsAsked"`

	Sid int32 `xml:"sid,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	Descending bool `xml:"descending,omitempty"`
}

type GetQuestionsAskedResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getQuestionsAskedResponse"`

	GetQuestionsAskedReturn string `xml:"getQuestionsAskedReturn,omitempty"`
}

type GetPriorities struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPriorities"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetPrioritiesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPrioritiesResponse"`

	GetPrioritiesReturn string `xml:"getPrioritiesReturn,omitempty"`
}

type GetDocumentTypes struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDocumentTypes"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetDocumentTypesResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDocumentTypesResponse"`

	GetDocumentTypesReturn string `xml:"getDocumentTypesReturn,omitempty"`
}

type GetTemplateList struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getTemplateList"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetTemplateListResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getTemplateListResponse"`

	GetTemplateListReturn string `xml:"getTemplateListReturn,omitempty"`
}

type GetWorkflowTemplateList struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getWorkflowTemplateList"`

	Sid int32 `xml:"sid,omitempty"`
}

type GetWorkflowTemplateListResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getWorkflowTemplateListResponse"`

	GetWorkflowTemplateListReturn string `xml:"getWorkflowTemplateListReturn,omitempty"`
}

type AddComment struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addComment"`

	Sid int32 `xml:"sid,omitempty"`

	Comment string `xml:"comment,omitempty"`

	DocId int32 `xml:"docId,omitempty"`

	Email string `xml:"email,omitempty"`

	Username string `xml:"username,omitempty"`

	ContactId string `xml:"contactId,omitempty"`
}

type AddCommentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addCommentResponse"`

	AddCommentReturn string `xml:"addCommentReturn,omitempty"`
}

type DeleteComment struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteComment"`

	Sid int32 `xml:"sid,omitempty"`

	CommentId int32 `xml:"commentId,omitempty"`
}

type DeleteCommentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteCommentResponse"`

	DeleteCommentReturn int32 `xml:"deleteCommentReturn,omitempty"`
}

type CreateDocument struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createDocument"`

	Sid int32 `xml:"sid,omitempty"`

	KdAttributes *ArrayOfString `xml:"kdAttributes,omitempty"`
}

type CreateDocumentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createDocumentResponse"`

	CreateDocumentReturn string `xml:"createDocumentReturn,omitempty"`
}

type ModifyDocument struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk modifyDocument"`

	Sid int32 `xml:"sid,omitempty"`

	DocId int32 `xml:"docId,omitempty"`

	KdAttributes *ArrayOfString `xml:"kdAttributes,omitempty"`
}

type ModifyDocumentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk modifyDocumentResponse"`

	ModifyDocumentReturn string `xml:"modifyDocumentReturn,omitempty"`
}

type RateDocument struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk rateDocument"`

	Sid int32 `xml:"sid,omitempty"`

	DocId int32 `xml:"docId,omitempty"`

	Rating int32 `xml:"rating,omitempty"`

	Multiplier int32 `xml:"multiplier,omitempty"`

	TicketPerId string `xml:"ticketPerId,omitempty"`

	OnTicketAccept bool `xml:"onTicketAccept,omitempty"`

	SolveUserProblem bool `xml:"solveUserProblem,omitempty"`

	IsDefault bool `xml:"isDefault,omitempty"`
}

type RateDocumentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk rateDocumentResponse"`

	RateDocumentReturn string `xml:"rateDocumentReturn,omitempty"`
}

type GetComments struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getComments"`

	Sid int32 `xml:"sid,omitempty"`

	DocIds string `xml:"docIds,omitempty"`
}

type GetCommentsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getCommentsResponse"`

	GetCommentsReturn string `xml:"getCommentsReturn,omitempty"`
}

type FindContacts struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk findContacts"`

	Sid int32 `xml:"sid,omitempty"`

	UserName string `xml:"userName,omitempty"`

	LastName string `xml:"lastName,omitempty"`

	FirstName string `xml:"firstName,omitempty"`

	Email string `xml:"email,omitempty"`

	AccessType string `xml:"accessType,omitempty"`

	InactiveFlag int32 `xml:"inactiveFlag,omitempty"`
}

type FindContactsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk findContactsResponse"`

	FindContactsReturn string `xml:"findContactsReturn,omitempty"`
}

type GetPermissionGroups struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPermissionGroups"`

	Sid int32 `xml:"sid,omitempty"`

	GroupId int32 `xml:"groupId,omitempty"`
}

type GetPermissionGroupsResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getPermissionGroupsResponse"`

	GetPermissionGroupsReturn string `xml:"getPermissionGroupsReturn,omitempty"`
}

type GetContact struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getContact"`

	Sid int32 `xml:"sid,omitempty"`

	ContactId string `xml:"contactId,omitempty"`
}

type GetContactResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getContactResponse"`

	GetContactReturn string `xml:"getContactReturn,omitempty"`
}

type AddBookmark struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addBookmark"`

	Sid int32 `xml:"sid,omitempty"`

	ContactId string `xml:"contactId,omitempty"`

	DocId int32 `xml:"docId,omitempty"`
}

type AddBookmarkResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk addBookmarkResponse"`

	AddBookmarkReturn string `xml:"addBookmarkReturn,omitempty"`
}

type UpdateRating struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk updateRating"`

	Sid int32 `xml:"sid,omitempty"`

	BuId int32 `xml:"buId,omitempty"`

	Rate int32 `xml:"rate,omitempty"`
}

type UpdateRatingResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk updateRatingResponse"`

	UpdateRatingReturn string `xml:"updateRatingReturn,omitempty"`
}

type DoSelectKD struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk doSelectKD"`

	Sid int32 `xml:"sid,omitempty"`

	WhereClause string `xml:"whereClause,omitempty"`

	SortBy string `xml:"sortBy,omitempty"`

	Desc bool `xml:"desc,omitempty"`

	MaxRows int32 `xml:"maxRows,omitempty"`

	Attributes *ArrayOfString `xml:"attributes,omitempty"`

	Skip int32 `xml:"skip,omitempty"`
}

type DoSelectKDResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk doSelectKDResponse"`

	DoSelectKDReturn string `xml:"doSelectKDReturn,omitempty"`
}

type GetFolderList struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getFolderList"`

	Sid int32 `xml:"sid,omitempty"`

	ParentFolderId int32 `xml:"parentFolderId,omitempty"`

	RepId int32 `xml:"repId,omitempty"`
}

type GetFolderListResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getFolderListResponse"`

	GetFolderListReturn string `xml:"getFolderListReturn,omitempty"`
}

type GetFolderInfo struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getFolderInfo"`

	Sid int32 `xml:"sid,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`
}

type GetFolderInfoResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getFolderInfoResponse"`

	GetFolderInfoReturn string `xml:"getFolderInfoReturn,omitempty"`
}

type GetAttmntList struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAttmntList"`

	Sid int32 `xml:"sid,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`

	RepId int32 `xml:"repId,omitempty"`
}

type GetAttmntListResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAttmntListResponse"`

	GetAttmntListReturn string `xml:"getAttmntListReturn,omitempty"`
}

type GetAttmntInfo struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAttmntInfo"`

	Sid int32 `xml:"sid,omitempty"`

	AttmntId int32 `xml:"attmntId,omitempty"`
}

type GetAttmntInfoResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAttmntInfoResponse"`

	GetAttmntInfoReturn string `xml:"getAttmntInfoReturn,omitempty"`
}

type GetRepositoryInfo struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getRepositoryInfo"`

	Sid int32 `xml:"sid,omitempty"`

	RepositoryId int32 `xml:"repositoryId,omitempty"`
}

type GetRepositoryInfoResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getRepositoryInfoResponse"`

	GetRepositoryInfoReturn string `xml:"getRepositoryInfoReturn,omitempty"`
}

type CreateFolder struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createFolder"`

	Sid int32 `xml:"sid,omitempty"`

	ParentFolderId int32 `xml:"parentFolderId,omitempty"`

	RepId int32 `xml:"repId,omitempty"`

	FolderType int32 `xml:"folderType,omitempty"`

	Description string `xml:"description,omitempty"`

	FolderName string `xml:"folderName,omitempty"`
}

type CreateFolderResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk createFolderResponse"`

	CreateFolderReturn string `xml:"createFolderReturn,omitempty"`
}

type Faq struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk faq"`

	Sid int32 `xml:"sid,omitempty"`

	CategoryIds string `xml:"categoryIds,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	PropertyList string `xml:"propertyList,omitempty"`

	SortBy string `xml:"sortBy,omitempty"`

	Descending bool `xml:"descending,omitempty"`

	WhereClause string `xml:"whereClause,omitempty"`

	MaxDocIDs int32 `xml:"maxDocIDs,omitempty"`
}

type FaqResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk faqResponse"`

	FaqReturn string `xml:"faqReturn,omitempty"`
}

type AttmntFolderLinkCount struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attmntFolderLinkCount"`

	Sid int32 `xml:"sid,omitempty"`

	FolderId int32 `xml:"folderId,omitempty"`
}

type AttmntFolderLinkCountResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attmntFolderLinkCountResponse"`

	AttmntFolderLinkCountReturn int32 `xml:"attmntFolderLinkCountReturn,omitempty"`
}

type AttachURLLink struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attachURLLink"`

	Sid int32 `xml:"sid,omitempty"`

	DocId int32 `xml:"docId,omitempty"`

	Url string `xml:"url,omitempty"`

	AttmntName string `xml:"attmntName,omitempty"`

	Description string `xml:"description,omitempty"`
}

type AttachURLLinkResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk attachURLLinkResponse"`

	AttachURLLinkReturn int32 `xml:"attachURLLinkReturn,omitempty"`
}

type DeleteBookmark struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteBookmark"`

	Sid int32 `xml:"sid,omitempty"`

	ContactId string `xml:"contactId,omitempty"`

	DocId int32 `xml:"docId,omitempty"`
}

type DeleteBookmarkResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk deleteBookmarkResponse"`

	DeleteBookmarkReturn int32 `xml:"deleteBookmarkReturn,omitempty"`
}

type GetKDListPerAttmnt struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getKDListPerAttmnt"`

	Sid int32 `xml:"sid,omitempty"`

	AttmntId int32 `xml:"attmntId,omitempty"`
}

type GetKDListPerAttmntResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getKDListPerAttmntResponse"`

	GetKDListPerAttmntReturn string `xml:"getKDListPerAttmntReturn,omitempty"`
}

type GetAttmntListPerKD struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAttmntListPerKD"`

	Sid int32 `xml:"sid,omitempty"`

	DocId int32 `xml:"docId,omitempty"`
}

type GetAttmntListPerKDResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getAttmntListPerKDResponse"`

	GetAttmntListPerKDReturn string `xml:"getAttmntListPerKDReturn,omitempty"`
}

type IsAttmntLinkedKD struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk isAttmntLinkedKD"`

	Sid int32 `xml:"sid,omitempty"`

	AttmntId int32 `xml:"attmntId,omitempty"`
}

type IsAttmntLinkedKDResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk isAttmntLinkedKDResponse"`

	IsAttmntLinkedKDReturn int32 `xml:"isAttmntLinkedKDReturn,omitempty"`
}

type Transfer struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk transfer"`

	Sid int32 `xml:"sid,omitempty"`

	Creator string `xml:"creator,omitempty"`

	ObjectHandle string `xml:"objectHandle,omitempty"`

	Description string `xml:"description,omitempty"`

	SetAssignee bool `xml:"setAssignee,omitempty"`

	NewAssigneeHandle string `xml:"newAssigneeHandle,omitempty"`

	SetGroup bool `xml:"setGroup,omitempty"`

	NewGroupHandle string `xml:"newGroupHandle,omitempty"`

	SetOrganization bool `xml:"setOrganization,omitempty"`

	NewOrganizationHandle string `xml:"newOrganizationHandle,omitempty"`
}

type TransferResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk transferResponse"`

	TransferReturn string `xml:"transferReturn,omitempty"`
}

type Search struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk search"`

	Sid int32 `xml:"sid,omitempty"`

	Problem string `xml:"problem,omitempty"`

	ResultSize int32 `xml:"resultSize,omitempty"`

	Properties string `xml:"properties,omitempty"`

	SortBy string `xml:"sortBy,omitempty"`

	Descending bool `xml:"descending,omitempty"`

	RelatedCategories bool `xml:"relatedCategories,omitempty"`

	SearchType int32 `xml:"searchType,omitempty"`

	MatchType int32 `xml:"matchType,omitempty"`

	SearchField int32 `xml:"searchField,omitempty"`

	CategoryPath string `xml:"categoryPath,omitempty"`

	WhereClause string `xml:"whereClause,omitempty"`

	MaxDocIDs int32 `xml:"maxDocIDs,omitempty"`
}

type SearchResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk searchResponse"`

	SearchReturn string `xml:"searchReturn,omitempty"`
}

type GetDocument struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDocument"`

	Sid int32 `xml:"sid,omitempty"`

	DocId int32 `xml:"docId,omitempty"`

	PropertyList string `xml:"propertyList,omitempty"`

	RelatedDoc bool `xml:"relatedDoc,omitempty"`

	GetAttmnt bool `xml:"getAttmnt,omitempty"`

	GetHistory bool `xml:"getHistory,omitempty"`

	GetComment bool `xml:"getComment,omitempty"`

	GetNotiList bool `xml:"getNotiList,omitempty"`
}

type GetDocumentResponse struct {
	XMLName xml.Name `xml:"http://www.ca.com/UnicenterServicePlus/ServiceDesk getDocumentResponse"`

	GetDocumentReturn string `xml:"getDocumentReturn,omitempty"`
}

type ArrayOfString struct {
	String []string `xml:"string,omitempty"`
}

type ListResult struct {
	ListHandle int32 `xml:"listHandle,omitempty"`

	ListLength int32 `xml:"listLength,omitempty"`
}

type ArrayOfInt struct {
	Integer []int32 `xml:"integer,omitempty"`
}

type USD_WebServiceSoap interface {
	CreateObject(request *CreateObject) (*CreateObjectResponse, error)

	AddAssetLog(request *AddAssetLog) (*AddAssetLogResponse, error)

	CreateLrelRelationships(request *CreateLrelRelationships) (*CreateLrelRelationshipsResponse, error)

	AddMemberToGroup(request *AddMemberToGroup) (*AddMemberToGroupResponse, error)

	AttachChangeToRequest(request *AttachChangeToRequest) (*AttachChangeToRequestResponse, error)

	CallServerMethod(request *CallServerMethod) (*CallServerMethodResponse, error)

	ChangeStatus(request *ChangeStatus) (*ChangeStatusResponse, error)

	ClearNotification(request *ClearNotification) (*ClearNotificationResponse, error)

	CreateActivityLog(request *CreateActivityLog) (*CreateActivityLogResponse, error)

	CreateAsset(request *CreateAsset) (*CreateAssetResponse, error)

	CreateAssetParentChildRelationship(request *CreateAssetParentChildRelationship) (*CreateAssetParentChildRelationshipResponse, error)

	CreateAttachment(request *CreateAttachment) (*CreateAttachmentResponse, error)

	RemoveAttachment(request *RemoveAttachment) (*RemoveAttachmentResponse, error)

	CreateChangeOrder(request *CreateChangeOrder) (*CreateChangeOrderResponse, error)

	CreateIssue(request *CreateIssue) (*CreateIssueResponse, error)

	CreateWorkFlowTask(request *CreateWorkFlowTask) (*CreateWorkFlowTaskResponse, error)

	DeleteWorkFlowTask(request *DeleteWorkFlowTask) (*DeleteWorkFlowTaskResponse, error)

	CreateRequest(request *CreateRequest) (*CreateRequestResponse, error)

	CreateTicket(request *CreateTicket) (*CreateTicketResponse, error)

	CreateQuickTicket(request *CreateQuickTicket) (*CreateQuickTicketResponse, error)

	CloseTicket(request *CloseTicket) (*CloseTicketResponse, error)

	LogComment(request *LogComment) (*LogCommentResponse, error)

	GetPolicyInfo(request *GetPolicyInfo) (*GetPolicyInfoResponse, error)

	DetachChangeFromRequest(request *DetachChangeFromRequest) (*DetachChangeFromRequestResponse, error)

	DoSelect(request *DoSelect) (*DoSelectResponse, error)

	DoQuery(request *DoQuery) (*DoQueryResponse, error)

	Escalate(request *Escalate) (*EscalateResponse, error)

	FreeListHandles(request *FreeListHandles) (*FreeListHandlesResponse, error)

	GetAssetExtensionInformation(request *GetAssetExtensionInformation) (*GetAssetExtensionInformationResponse, error)

	GetConfigurationMode(request *GetConfigurationMode) (*GetConfigurationModeResponse, error)

	GetGroupMemberListValues(request *GetGroupMemberListValues) (*GetGroupMemberListValuesResponse, error)

	GetObjectTypeInformation(request *GetObjectTypeInformation) (*GetObjectTypeInformationResponse, error)

	GetHandleForUserid(request *GetHandleForUserid) (*GetHandleForUseridResponse, error)

	GetAccessTypeForContact(request *GetAccessTypeForContact) (*GetAccessTypeForContactResponse, error)

	GetListValues(request *GetListValues) (*GetListValuesResponse, error)

	GetLrelLength(request *GetLrelLength) (*GetLrelLengthResponse, error)

	GetLrelValues(request *GetLrelValues) (*GetLrelValuesResponse, error)

	GetNotificationsForContact(request *GetNotificationsForContact) (*GetNotificationsForContactResponse, error)

	GetObjectValues(request *GetObjectValues) (*GetObjectValuesResponse, error)

	GetPendingChangeTaskListForContact(request *GetPendingChangeTaskListForContact) (*GetPendingChangeTaskListForContactResponse, error)

	GetPendingIssueTaskListForContact(request *GetPendingIssueTaskListForContact) (*GetPendingIssueTaskListForContactResponse, error)

	GetPropertyInfoForCategory(request *GetPropertyInfoForCategory) (*GetPropertyInfoForCategoryResponse, error)

	GetRelatedList(request *GetRelatedList) (*GetRelatedListResponse, error)

	GetRelatedListValues(request *GetRelatedListValues) (*GetRelatedListValuesResponse, error)

	GetWorkFlowTemplates(request *GetWorkFlowTemplates) (*GetWorkFlowTemplatesResponse, error)

	GetTaskListValues(request *GetTaskListValues) (*GetTaskListValuesResponse, error)

	GetValidTaskTransitions(request *GetValidTaskTransitions) (*GetValidTaskTransitionsResponse, error)

	GetValidTransitions(request *GetValidTransitions) (*GetValidTransitionsResponse, error)

	GetDependentAttrControls(request *GetDependentAttrControls) (*GetDependentAttrControlsResponse, error)

	Login(request *Login) (*LoginResponse, error)

	LoginService(request *LoginService) (*LoginServiceResponse, error)

	LoginServiceManaged(request *LoginServiceManaged) (*LoginServiceManagedResponse, error)

	LoginWithArtifact(request *LoginWithArtifact) (*LoginWithArtifactResponse, error)

	Impersonate(request *Impersonate) (*ImpersonateResponse, error)

	Logout(request *Logout) (*LogoutResponse, error)

	NotifyContacts(request *NotifyContacts) (*NotifyContactsResponse, error)

	RemoveLrelRelationships(request *RemoveLrelRelationships) (*RemoveLrelRelationshipsResponse, error)

	RemoveMemberFromGroup(request *RemoveMemberFromGroup) (*RemoveMemberFromGroupResponse, error)

	ServerStatus(request *ServerStatus) (*ServerStatusResponse, error)

	UpdateObject(request *UpdateObject) (*UpdateObjectResponse, error)

	GetBopsid(request *GetBopsid) (*GetBopsidResponse, error)

	GetArtifact(request *GetArtifact) (*GetArtifactResponse, error)

	AttachURLLinkToTicket(request *AttachURLLinkToTicket) (*AttachURLLinkToTicketResponse, error)

	CreateAttmnt(request *CreateAttmnt) (*CreateAttmntResponse, error)

	GetDocumentsByIDs(request *GetDocumentsByIDs) (*GetDocumentsByIDsResponse, error)

	GetDecisionTrees(request *GetDecisionTrees) (*GetDecisionTreesResponse, error)

	DeleteDocument(request *DeleteDocument) (*DeleteDocumentResponse, error)

	GetCategory(request *GetCategory) (*GetCategoryResponse, error)

	GetStatuses(request *GetStatuses) (*GetStatusesResponse, error)

	GetBookmarks(request *GetBookmarks) (*GetBookmarksResponse, error)

	GetQuestionsAsked(request *GetQuestionsAsked) (*GetQuestionsAskedResponse, error)

	GetPriorities(request *GetPriorities) (*GetPrioritiesResponse, error)

	GetDocumentTypes(request *GetDocumentTypes) (*GetDocumentTypesResponse, error)

	GetTemplateList(request *GetTemplateList) (*GetTemplateListResponse, error)

	GetWorkflowTemplateList(request *GetWorkflowTemplateList) (*GetWorkflowTemplateListResponse, error)

	AddComment(request *AddComment) (*AddCommentResponse, error)

	DeleteComment(request *DeleteComment) (*DeleteCommentResponse, error)

	CreateDocument(request *CreateDocument) (*CreateDocumentResponse, error)

	ModifyDocument(request *ModifyDocument) (*ModifyDocumentResponse, error)

	RateDocument(request *RateDocument) (*RateDocumentResponse, error)

	GetComments(request *GetComments) (*GetCommentsResponse, error)

	FindContacts(request *FindContacts) (*FindContactsResponse, error)

	GetPermissionGroups(request *GetPermissionGroups) (*GetPermissionGroupsResponse, error)

	GetContact(request *GetContact) (*GetContactResponse, error)

	AddBookmark(request *AddBookmark) (*AddBookmarkResponse, error)

	UpdateRating(request *UpdateRating) (*UpdateRatingResponse, error)

	DoSelectKD(request *DoSelectKD) (*DoSelectKDResponse, error)

	GetFolderList(request *GetFolderList) (*GetFolderListResponse, error)

	GetFolderInfo(request *GetFolderInfo) (*GetFolderInfoResponse, error)

	GetAttmntList(request *GetAttmntList) (*GetAttmntListResponse, error)

	GetAttmntInfo(request *GetAttmntInfo) (*GetAttmntInfoResponse, error)

	GetRepositoryInfo(request *GetRepositoryInfo) (*GetRepositoryInfoResponse, error)

	CreateFolder(request *CreateFolder) (*CreateFolderResponse, error)

	Faq(request *Faq) (*FaqResponse, error)

	AttmntFolderLinkCount(request *AttmntFolderLinkCount) (*AttmntFolderLinkCountResponse, error)

	AttachURLLink(request *AttachURLLink) (*AttachURLLinkResponse, error)

	DeleteBookmark(request *DeleteBookmark) (*DeleteBookmarkResponse, error)

	GetKDListPerAttmnt(request *GetKDListPerAttmnt) (*GetKDListPerAttmntResponse, error)

	GetAttmntListPerKD(request *GetAttmntListPerKD) (*GetAttmntListPerKDResponse, error)

	IsAttmntLinkedKD(request *IsAttmntLinkedKD) (*IsAttmntLinkedKDResponse, error)

	Transfer(request *Transfer) (*TransferResponse, error)

	Search(request *Search) (*SearchResponse, error)

	GetDocument(request *GetDocument) (*GetDocumentResponse, error)
}

type uSD_WebServiceSoap struct {
	client *soap.Client
}

func NewUSD_WebServiceSoap(client *soap.Client) USD_WebServiceSoap {
	return &uSD_WebServiceSoap{
		client: client,
	}
}

func (service *uSD_WebServiceSoap) CreateObject(request *CreateObject) (*CreateObjectResponse, error) {
	response := new(CreateObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AddAssetLog(request *AddAssetLog) (*AddAssetLogResponse, error) {
	response := new(AddAssetLogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateLrelRelationships(request *CreateLrelRelationships) (*CreateLrelRelationshipsResponse, error) {
	response := new(CreateLrelRelationshipsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AddMemberToGroup(request *AddMemberToGroup) (*AddMemberToGroupResponse, error) {
	response := new(AddMemberToGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AttachChangeToRequest(request *AttachChangeToRequest) (*AttachChangeToRequestResponse, error) {
	response := new(AttachChangeToRequestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CallServerMethod(request *CallServerMethod) (*CallServerMethodResponse, error) {
	response := new(CallServerMethodResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) ChangeStatus(request *ChangeStatus) (*ChangeStatusResponse, error) {
	response := new(ChangeStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) ClearNotification(request *ClearNotification) (*ClearNotificationResponse, error) {
	response := new(ClearNotificationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateActivityLog(request *CreateActivityLog) (*CreateActivityLogResponse, error) {
	response := new(CreateActivityLogResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateAsset(request *CreateAsset) (*CreateAssetResponse, error) {
	response := new(CreateAssetResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateAssetParentChildRelationship(request *CreateAssetParentChildRelationship) (*CreateAssetParentChildRelationshipResponse, error) {
	response := new(CreateAssetParentChildRelationshipResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateAttachment(request *CreateAttachment) (*CreateAttachmentResponse, error) {
	response := new(CreateAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) RemoveAttachment(request *RemoveAttachment) (*RemoveAttachmentResponse, error) {
	response := new(RemoveAttachmentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateChangeOrder(request *CreateChangeOrder) (*CreateChangeOrderResponse, error) {
	response := new(CreateChangeOrderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateIssue(request *CreateIssue) (*CreateIssueResponse, error) {
	response := new(CreateIssueResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateWorkFlowTask(request *CreateWorkFlowTask) (*CreateWorkFlowTaskResponse, error) {
	response := new(CreateWorkFlowTaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DeleteWorkFlowTask(request *DeleteWorkFlowTask) (*DeleteWorkFlowTaskResponse, error) {
	response := new(DeleteWorkFlowTaskResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateRequest(request *CreateRequest) (*CreateRequestResponse, error) {
	response := new(CreateRequestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateTicket(request *CreateTicket) (*CreateTicketResponse, error) {
	response := new(CreateTicketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateQuickTicket(request *CreateQuickTicket) (*CreateQuickTicketResponse, error) {
	response := new(CreateQuickTicketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CloseTicket(request *CloseTicket) (*CloseTicketResponse, error) {
	response := new(CloseTicketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) LogComment(request *LogComment) (*LogCommentResponse, error) {
	response := new(LogCommentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetPolicyInfo(request *GetPolicyInfo) (*GetPolicyInfoResponse, error) {
	response := new(GetPolicyInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DetachChangeFromRequest(request *DetachChangeFromRequest) (*DetachChangeFromRequestResponse, error) {
	response := new(DetachChangeFromRequestResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DoSelect(request *DoSelect) (*DoSelectResponse, error) {
	response := new(DoSelectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DoQuery(request *DoQuery) (*DoQueryResponse, error) {
	response := new(DoQueryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Escalate(request *Escalate) (*EscalateResponse, error) {
	response := new(EscalateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) FreeListHandles(request *FreeListHandles) (*FreeListHandlesResponse, error) {
	response := new(FreeListHandlesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetAssetExtensionInformation(request *GetAssetExtensionInformation) (*GetAssetExtensionInformationResponse, error) {
	response := new(GetAssetExtensionInformationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetConfigurationMode(request *GetConfigurationMode) (*GetConfigurationModeResponse, error) {
	response := new(GetConfigurationModeResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetGroupMemberListValues(request *GetGroupMemberListValues) (*GetGroupMemberListValuesResponse, error) {
	response := new(GetGroupMemberListValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetObjectTypeInformation(request *GetObjectTypeInformation) (*GetObjectTypeInformationResponse, error) {
	response := new(GetObjectTypeInformationResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetHandleForUserid(request *GetHandleForUserid) (*GetHandleForUseridResponse, error) {
	response := new(GetHandleForUseridResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetAccessTypeForContact(request *GetAccessTypeForContact) (*GetAccessTypeForContactResponse, error) {
	response := new(GetAccessTypeForContactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetListValues(request *GetListValues) (*GetListValuesResponse, error) {
	response := new(GetListValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetLrelLength(request *GetLrelLength) (*GetLrelLengthResponse, error) {
	response := new(GetLrelLengthResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetLrelValues(request *GetLrelValues) (*GetLrelValuesResponse, error) {
	response := new(GetLrelValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetNotificationsForContact(request *GetNotificationsForContact) (*GetNotificationsForContactResponse, error) {
	response := new(GetNotificationsForContactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetObjectValues(request *GetObjectValues) (*GetObjectValuesResponse, error) {
	response := new(GetObjectValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetPendingChangeTaskListForContact(request *GetPendingChangeTaskListForContact) (*GetPendingChangeTaskListForContactResponse, error) {
	response := new(GetPendingChangeTaskListForContactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetPendingIssueTaskListForContact(request *GetPendingIssueTaskListForContact) (*GetPendingIssueTaskListForContactResponse, error) {
	response := new(GetPendingIssueTaskListForContactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetPropertyInfoForCategory(request *GetPropertyInfoForCategory) (*GetPropertyInfoForCategoryResponse, error) {
	response := new(GetPropertyInfoForCategoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetRelatedList(request *GetRelatedList) (*GetRelatedListResponse, error) {
	response := new(GetRelatedListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetRelatedListValues(request *GetRelatedListValues) (*GetRelatedListValuesResponse, error) {
	response := new(GetRelatedListValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetWorkFlowTemplates(request *GetWorkFlowTemplates) (*GetWorkFlowTemplatesResponse, error) {
	response := new(GetWorkFlowTemplatesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetTaskListValues(request *GetTaskListValues) (*GetTaskListValuesResponse, error) {
	response := new(GetTaskListValuesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetValidTaskTransitions(request *GetValidTaskTransitions) (*GetValidTaskTransitionsResponse, error) {
	response := new(GetValidTaskTransitionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetValidTransitions(request *GetValidTransitions) (*GetValidTransitionsResponse, error) {
	response := new(GetValidTransitionsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetDependentAttrControls(request *GetDependentAttrControls) (*GetDependentAttrControlsResponse, error) {
	response := new(GetDependentAttrControlsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Login(request *Login) (*LoginResponse, error) {
	response := new(LoginResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) LoginService(request *LoginService) (*LoginServiceResponse, error) {
	response := new(LoginServiceResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) LoginServiceManaged(request *LoginServiceManaged) (*LoginServiceManagedResponse, error) {
	response := new(LoginServiceManagedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) LoginWithArtifact(request *LoginWithArtifact) (*LoginWithArtifactResponse, error) {
	response := new(LoginWithArtifactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Impersonate(request *Impersonate) (*ImpersonateResponse, error) {
	response := new(ImpersonateResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Logout(request *Logout) (*LogoutResponse, error) {
	response := new(LogoutResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) NotifyContacts(request *NotifyContacts) (*NotifyContactsResponse, error) {
	response := new(NotifyContactsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) RemoveLrelRelationships(request *RemoveLrelRelationships) (*RemoveLrelRelationshipsResponse, error) {
	response := new(RemoveLrelRelationshipsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) RemoveMemberFromGroup(request *RemoveMemberFromGroup) (*RemoveMemberFromGroupResponse, error) {
	response := new(RemoveMemberFromGroupResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) ServerStatus(request *ServerStatus) (*ServerStatusResponse, error) {
	response := new(ServerStatusResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) UpdateObject(request *UpdateObject) (*UpdateObjectResponse, error) {
	response := new(UpdateObjectResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetBopsid(request *GetBopsid) (*GetBopsidResponse, error) {
	response := new(GetBopsidResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetArtifact(request *GetArtifact) (*GetArtifactResponse, error) {
	response := new(GetArtifactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AttachURLLinkToTicket(request *AttachURLLinkToTicket) (*AttachURLLinkToTicketResponse, error) {
	response := new(AttachURLLinkToTicketResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateAttmnt(request *CreateAttmnt) (*CreateAttmntResponse, error) {
	response := new(CreateAttmntResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetDocumentsByIDs(request *GetDocumentsByIDs) (*GetDocumentsByIDsResponse, error) {
	response := new(GetDocumentsByIDsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetDecisionTrees(request *GetDecisionTrees) (*GetDecisionTreesResponse, error) {
	response := new(GetDecisionTreesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DeleteDocument(request *DeleteDocument) (*DeleteDocumentResponse, error) {
	response := new(DeleteDocumentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetCategory(request *GetCategory) (*GetCategoryResponse, error) {
	response := new(GetCategoryResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetStatuses(request *GetStatuses) (*GetStatusesResponse, error) {
	response := new(GetStatusesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetBookmarks(request *GetBookmarks) (*GetBookmarksResponse, error) {
	response := new(GetBookmarksResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetQuestionsAsked(request *GetQuestionsAsked) (*GetQuestionsAskedResponse, error) {
	response := new(GetQuestionsAskedResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetPriorities(request *GetPriorities) (*GetPrioritiesResponse, error) {
	response := new(GetPrioritiesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetDocumentTypes(request *GetDocumentTypes) (*GetDocumentTypesResponse, error) {
	response := new(GetDocumentTypesResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetTemplateList(request *GetTemplateList) (*GetTemplateListResponse, error) {
	response := new(GetTemplateListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetWorkflowTemplateList(request *GetWorkflowTemplateList) (*GetWorkflowTemplateListResponse, error) {
	response := new(GetWorkflowTemplateListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AddComment(request *AddComment) (*AddCommentResponse, error) {
	response := new(AddCommentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DeleteComment(request *DeleteComment) (*DeleteCommentResponse, error) {
	response := new(DeleteCommentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateDocument(request *CreateDocument) (*CreateDocumentResponse, error) {
	response := new(CreateDocumentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) ModifyDocument(request *ModifyDocument) (*ModifyDocumentResponse, error) {
	response := new(ModifyDocumentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) RateDocument(request *RateDocument) (*RateDocumentResponse, error) {
	response := new(RateDocumentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetComments(request *GetComments) (*GetCommentsResponse, error) {
	response := new(GetCommentsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) FindContacts(request *FindContacts) (*FindContactsResponse, error) {
	response := new(FindContactsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetPermissionGroups(request *GetPermissionGroups) (*GetPermissionGroupsResponse, error) {
	response := new(GetPermissionGroupsResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetContact(request *GetContact) (*GetContactResponse, error) {
	response := new(GetContactResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AddBookmark(request *AddBookmark) (*AddBookmarkResponse, error) {
	response := new(AddBookmarkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) UpdateRating(request *UpdateRating) (*UpdateRatingResponse, error) {
	response := new(UpdateRatingResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DoSelectKD(request *DoSelectKD) (*DoSelectKDResponse, error) {
	response := new(DoSelectKDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetFolderList(request *GetFolderList) (*GetFolderListResponse, error) {
	response := new(GetFolderListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetFolderInfo(request *GetFolderInfo) (*GetFolderInfoResponse, error) {
	response := new(GetFolderInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetAttmntList(request *GetAttmntList) (*GetAttmntListResponse, error) {
	response := new(GetAttmntListResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetAttmntInfo(request *GetAttmntInfo) (*GetAttmntInfoResponse, error) {
	response := new(GetAttmntInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetRepositoryInfo(request *GetRepositoryInfo) (*GetRepositoryInfoResponse, error) {
	response := new(GetRepositoryInfoResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) CreateFolder(request *CreateFolder) (*CreateFolderResponse, error) {
	response := new(CreateFolderResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Faq(request *Faq) (*FaqResponse, error) {
	response := new(FaqResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AttmntFolderLinkCount(request *AttmntFolderLinkCount) (*AttmntFolderLinkCountResponse, error) {
	response := new(AttmntFolderLinkCountResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) AttachURLLink(request *AttachURLLink) (*AttachURLLinkResponse, error) {
	response := new(AttachURLLinkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) DeleteBookmark(request *DeleteBookmark) (*DeleteBookmarkResponse, error) {
	response := new(DeleteBookmarkResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetKDListPerAttmnt(request *GetKDListPerAttmnt) (*GetKDListPerAttmntResponse, error) {
	response := new(GetKDListPerAttmntResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetAttmntListPerKD(request *GetAttmntListPerKD) (*GetAttmntListPerKDResponse, error) {
	response := new(GetAttmntListPerKDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) IsAttmntLinkedKD(request *IsAttmntLinkedKD) (*IsAttmntLinkedKDResponse, error) {
	response := new(IsAttmntLinkedKDResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Transfer(request *Transfer) (*TransferResponse, error) {
	response := new(TransferResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) Search(request *Search) (*SearchResponse, error) {
	response := new(SearchResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *uSD_WebServiceSoap) GetDocument(request *GetDocument) (*GetDocumentResponse, error) {
	response := new(GetDocumentResponse)
	err := service.client.Call("", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
