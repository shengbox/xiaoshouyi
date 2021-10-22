package api

type BaseResp struct {
	Code      string        `json:"code"`
	Msg       string        `json:"msg"`
	ErrorInfo interface{}   `json:"errorInfo"`
	Ext       []interface{} `json:"ext"`
}

type ActivityResp struct {
	BaseResp `json:",inline"`
	Result   struct {
		TotalSize int           `json:"totalSize"`
		Count     int           `json:"count"`
		Records   []interface{} `json:"records"`
	} `json:"result"`
}

// ActivityRecord 活动记录
type ActivityRecord struct {
	Id                int64       `json:"id,omitempty"`
	ContactId         int64       `json:"contactId" bson:"contactId"`
	OwnerId           int64       `json:"ownerId,omitempty" bson:"ownerId"`
	ItemId            int64       `json:"itemId,omitempty"`
	StartTime         int64       `json:"startTime,omitempty" bson:"startTime"`
	IntentionalDegree interface{} `json:"intentionalDegree,omitempty"`
	NeedFollow        bool        `json:"needFollow,omitempty"`
	EntityType        int64       `json:"entityType,omitempty"`
	DimDepart         int64       `json:"dimDepart,omitempty"`
	BelongId          int         `json:"belongId,omitempty"`
	Content           string      `json:"content,omitempty"`
	ContactName       interface{} `json:"contactName,omitempty"`
	ContactPhone      interface{} `json:"contactPhone,omitempty"`
	DbcRelation26     int64       `json:"dbcRelation26,omitempty" bson:"dbcRelation26"`
	UpdatedAt         int64       `json:"updatedAt,omitempty" bson:"updatedAt"`
	CreatedAt         int64       `json:"createdAt,omitempty" bson:"createdAt"`
}

// Contact 联系人
type Contact struct {
	Id                       int64       `json:"id"`
	AccountId                int64       `json:"accountId"`
	Address                  interface{} `json:"address"`
	Comment                  interface{} `json:"comment"`
	ContactBirthday          interface{} `json:"contactBirthday"`
	ContactName              string      `json:"contactName"`
	ContactRole              *int        `json:"contactRole"`
	ContactScore             int         `json:"contactScore"`
	CreatedAt                int64       `json:"createdAt"`
	CreatedBy                int64       `json:"createdBy"`
	Depart                   interface{} `json:"depart"`
	DimDepart                int64       `json:"dimDepart"`
	DoNotDisturb             bool        `json:"doNotDisturb"`
	Email                    interface{} `json:"email"`
	EntityType               int64       `json:"entityType"`
	Gender                   *int        `json:"gender"`
	LockStatus               int         `json:"lockStatus"`
	Mobile                   *string     `json:"mobile"`
	OwnerId                  int64       `json:"ownerId"`
	Phone                    *string     `json:"phone"`
	Post                     interface{} `json:"post"`
	RecentActivityRecordTime interface{} `json:"recentActivityRecordTime"`
	State                    interface{} `json:"state"`
	UpdatedAt                int64       `json:"updatedAt"`
	UpdatedBy                int64       `json:"updatedBy"`
	ZipCode                  interface{} `json:"zipCode"`
}

// Document 文档
type Document struct {
	Id          int64  `json:"id"`
	FileName    string `json:"fileName"`
	Type        string `json:"type"`
	FileType    int    `json:"fileType"`
	BelongId    int    `json:"belongId"`
	ObjectId    int64  `json:"objectId"`
	FileLength  int    `json:"fileLength"`
	DirectoryId int    `json:"directoryId"`
	UpdatedAt   string `json:"updatedAt"`
}

// User 用户
type User struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Status int    `json:"status"`
}

// Account 客户
type Account struct {
	AccountName               string      `json:"accountName,omitempty" bson:"accountName"`
	AccountScore              int         `json:"accountScore,omitempty"`
	ActualInvoicedAmount      int         `json:"actualInvoicedAmount,omitempty"`
	Address                   string      `json:"address,omitempty"`
	AmountUnbilled            int         `json:"amountUnbilled,omitempty"`
	AnnualRevenue             interface{} `json:"annualRevenue,omitempty"`
	ClaimTime                 interface{} `json:"claimTime,omitempty"`
	Comment                   interface{} `json:"comment,omitempty"`
	CreatedAt                 int64       `json:"createdAt,omitempty"`
	CreatedBy                 int64       `json:"createdBy,omitempty"`
	DimDepart                 int64       `json:"dimDepart,omitempty"`
	DoNotDisturb              bool        `json:"doNotDisturb,omitempty"`
	EcouponsAccountLabel      interface{} `json:"ecouponsAccountLabel,omitempty"`
	EmployeeNumber            interface{} `json:"employeeNumber,omitempty"`
	ExpireTime                interface{} `json:"expireTime,omitempty"`
	FCity                     int         `json:"fCity,omitempty"`
	FDistrict                 interface{} `json:"fDistrict,omitempty"`
	FState                    int         `json:"fState,omitempty"`
	Fax                       interface{} `json:"fax,omitempty"`
	HighSeaAccountSource      []int       `json:"highSeaAccountSource,omitempty"`
	HighSeaId                 int64       `json:"highSeaId,omitempty"`
	HighSeaStatus             int         `json:"highSeaStatus,omitempty" bson:"highSeaStatus"`
	Id                        int64       `json:"id,omitempty"`
	IndustryId                int         `json:"industryId,omitempty"`
	InvoiceBalance            int         `json:"invoiceBalance,omitempty"`
	IsCustomer                string      `json:"isCustomer,omitempty"`
	LeadId                    interface{} `json:"leadId,omitempty"`
	Level                     int         `json:"level,omitempty"`
	LockStatus                int         `json:"lockStatus,omitempty"`
	OutterDepartId            interface{} `json:"outterDepartId,omitempty"`
	OwnerId                   int64       `json:"ownerId,omitempty" bson:"ownerId"`
	PaidAmount                int         `json:"paidAmount,omitempty"`
	ParentAccountId           interface{} `json:"parentAccountId,omitempty"`
	Phone                     interface{} `json:"phone,omitempty"`
	RecentActivityCreatedBy   interface{} `json:"recentActivityCreatedBy,omitempty"`
	RecentActivityRecordTime  interface{} `json:"recentActivityRecordTime,omitempty"`
	SrcFlg                    int         `json:"srcFlg,omitempty"`
	TotalActiveOrders         int         `json:"totalActiveOrders,omitempty"`
	TotalOrderAmount          int         `json:"totalOrderAmount,omitempty"`
	TotalWonOpportunities     int         `json:"totalWonOpportunities,omitempty"`
	TotalWonOpportunityAmount int         `json:"totalWonOpportunityAmount,omitempty"`
	UnpaidAmount              int         `json:"unpaidAmount,omitempty"`
	UpdatedAt                 int64       `json:"updatedAt,omitempty"`
	UpdatedBy                 int64       `json:"updatedBy,omitempty"`
	Url                       interface{} `json:"url,omitempty"`
	VisitInplanCount          int         `json:"visitInplanCount,omitempty"`
	VisitLatestTime           interface{} `json:"visitLatestTime,omitempty"`
	VisitTotalCount           int         `json:"visitTotalCount,omitempty"`
	VisitUnvisitDay           interface{} `json:"visitUnvisitDay,omitempty"`
	ZipCode                   interface{} `json:"zipCode,omitempty"`
	EntityType                int64       `json:"entityType,omitempty"`
}

// Contract 合同
type Contract struct {
	AccountId        int64       `json:"accountId"`
	CreatedAt        int64       `json:"createdAt"`
	EndDate          int64       `json:"endDate"`
	Id               int64       `json:"id"`
	OwnerId          int64       `json:"ownerId"`
	StartDate        int64       `json:"startDate"`
	Status           int         `json:"status"`
	Title            string      `json:"title"`
	UpdatedAt        int64       `json:"updatedAt"`
	CustomItem158__C int64       `json:"customItem158__c"`
	CustomItem159__C string      `json:"customItem159__c"`
	CustomItem160__C int         `json:"customItem160__c"`
	CustomItem161__C interface{} `json:"customItem161__c"`
	CustomItem162__C interface{} `json:"customItem162__c"`
	CustomItem163__C interface{} `json:"customItem163__c"`
	CustomItem164__C interface{} `json:"customItem164__c"`
	CustomItem165__C interface{} `json:"customItem165__c"`
}

//Product 产品
type Product struct {
	Id          int64  `json:"id"`
	ProductName string `json:"productName"`
}

type DocumentResp struct {
	TotalSize int        `json:"totalSize"`
	Count     int        `json:"count"`
	Record    []Document `json:"record"`
}

type Description struct {
	BaseResp `json:",inline"`
	Data     struct {
		ApiKey      string `json:"apiKey"`
		Custom      bool   `json:"custom"`
		Label       string `json:"label"`
		Disabled    bool   `json:"disabled"`
		Createable  bool   `json:"createable"`
		Deletable   bool   `json:"deletable"`
		Updateable  bool   `json:"updateable"`
		Queryable   bool   `json:"queryable"`
		FeedEnabled bool   `json:"feedEnabled"`
		Fields      []struct {
			ApiKey                string      `json:"apiKey"`
			Label                 string      `json:"label"`
			Type                  string      `json:"type"`
			ItemType              string      `json:"itemType"`
			DefaultValue          *string     `json:"defaultValue"`
			Enabled               bool        `json:"enabled"`
			Required              bool        `json:"required"`
			Createable            bool        `json:"createable"`
			Updateable            bool        `json:"updateable"`
			Sortable              bool        `json:"sortable"`
			MinLength             interface{} `json:"minLength"`
			MaxLength             *int        `json:"maxLength"`
			DependentPropertyName interface{} `json:"dependentPropertyName"`
			Unique                bool        `json:"unique"`
			Encrypt               bool        `json:"encrypt"`
			ReferTo               struct {
				ApiKey string `json:"apiKey"`
			} `json:"referTo,omitempty"`
			Selectitem []struct {
				Label          string `json:"label"`
				Value          int    `json:"value"`
				ApiKey         string `json:"apiKey"`
				DependentValue []int  `json:"dependentValue"`
				IsActive       bool   `json:"isActive"`
			} `json:"selectitem"`
			Checkitem []struct {
				Label    string `json:"label"`
				Value    int    `json:"value"`
				ApiKey   string `json:"apiKey"`
				IsActive bool   `json:"isActive"`
			} `json:"checkitem"`
		} `json:"fields"`
	} `json:"data"`
}

type CommonResp struct {
	BaseResp `json:",inline"`
	Data     struct {
		Id           int64  `json:"id"`
		ObjectApiKey string `json:"objectApiKey"`
		BusiType     int64  `json:"busiType"`
		Name         string `json:"name"`
	} `json:"data"`
	ErrorInfo interface{} `json:"errorInfo"`
}
