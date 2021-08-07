package api

type ActivityResp struct {
	Code   int           `json:"code"`
	Msg    string        `json:"msg"`
	Ext    []interface{} `json:"ext"`
	Result struct {
		TotalSize int           `json:"totalSize"`
		Count     int           `json:"count"`
		Records   []interface{} `json:"records"`
	} `json:"result"`
}

// ActivityRecord 活动记录
type ActivityRecord struct {
	Id                int64       `json:"id"`
	OwnerId           int64       `json:"ownerId"`
	StartTime         int64       `json:"startTime"`
	IntentionalDegree interface{} `json:"intentionalDegree"`
	NeedFollow        bool        `json:"needFollow"`
	EntityType        int64       `json:"entityType"`
	Content           string      `json:"content"`
	ContactName       interface{} `json:"contactName"`
	ContactPhone      interface{} `json:"contactPhone"`
	DbcRelation26     int64       `json:"dbcRelation26"`
	UpdatedAt         int         `json:"updatedAt"`
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
	AccountName               string      `json:"accountName"`
	AccountScore              int         `json:"accountScore"`
	ActualInvoicedAmount      int         `json:"actualInvoicedAmount"`
	Address                   string      `json:"address"`
	AmountUnbilled            int         `json:"amountUnbilled"`
	AnnualRevenue             interface{} `json:"annualRevenue"`
	ClaimTime                 interface{} `json:"claimTime"`
	Comment                   interface{} `json:"comment"`
	CreatedAt                 int64       `json:"createdAt"`
	CreatedBy                 int64       `json:"createdBy"`
	DimDepart                 int64       `json:"dimDepart"`
	DoNotDisturb              bool        `json:"doNotDisturb"`
	EcouponsAccountLabel      interface{} `json:"ecouponsAccountLabel"`
	EmployeeNumber            interface{} `json:"employeeNumber"`
	ExpireTime                interface{} `json:"expireTime"`
	FCity                     int         `json:"fCity"`
	FDistrict                 interface{} `json:"fDistrict"`
	FState                    int         `json:"fState"`
	Fax                       interface{} `json:"fax"`
	HighSeaAccountSource      []int       `json:"highSeaAccountSource"`
	HighSeaId                 int64       `json:"highSeaId"`
	HighSeaStatus             int         `json:"highSeaStatus"`
	Id                        int64       `json:"id"`
	IndustryId                int         `json:"industryId"`
	InvoiceBalance            int         `json:"invoiceBalance"`
	IsCustomer                string      `json:"isCustomer"`
	LeadId                    interface{} `json:"leadId"`
	Level                     int         `json:"level"`
	LockStatus                int         `json:"lockStatus"`
	OutterDepartId            interface{} `json:"outterDepartId"`
	OwnerId                   int64       `json:"ownerId"`
	PaidAmount                int         `json:"paidAmount"`
	ParentAccountId           interface{} `json:"parentAccountId"`
	Phone                     interface{} `json:"phone"`
	RecentActivityCreatedBy   interface{} `json:"recentActivityCreatedBy"`
	RecentActivityRecordTime  interface{} `json:"recentActivityRecordTime"`
	SrcFlg                    int         `json:"srcFlg"`
	TotalActiveOrders         int         `json:"totalActiveOrders"`
	TotalOrderAmount          int         `json:"totalOrderAmount"`
	TotalWonOpportunities     int         `json:"totalWonOpportunities"`
	TotalWonOpportunityAmount int         `json:"totalWonOpportunityAmount"`
	UnpaidAmount              int         `json:"unpaidAmount"`
	UpdatedAt                 int64       `json:"updatedAt"`
	UpdatedBy                 int64       `json:"updatedBy"`
	Url                       interface{} `json:"url"`
	VisitInplanCount          int         `json:"visitInplanCount"`
	VisitLatestTime           interface{} `json:"visitLatestTime"`
	VisitTotalCount           int         `json:"visitTotalCount"`
	VisitUnvisitDay           interface{} `json:"visitUnvisitDay"`
	ZipCode                   interface{} `json:"zipCode"`
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
