package api

type BaseResp struct {
	Code      string        `json:"code"`
	Msg       string        `json:"msg"`
	ErrorInfo interface{}   `json:"errorInfo"`
	Ext       []interface{} `json:"ext"`
}

type XOQLResp struct {
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
	Id                int64       `json:"id,omitempty"`
	ContactId         int64       `json:"contactId" bson:"contactId"`
	OwnerId           int64       `json:"ownerId,omitempty" bson:"ownerId"`
	StartTime         int64       `json:"startTime,omitempty" bson:"startTime"`
	IntentionalDegree interface{} `json:"intentionalDegree,omitempty"`
	NeedFollow        bool        `json:"needFollow,omitempty"`
	EntityType        int64       `json:"entityType,omitempty"`
	DimDepart         int64       `json:"dimDepart,omitempty"`
	Content           string      `json:"content,omitempty"`
	ContactName       interface{} `json:"contactName,omitempty"`
	ContactPhone      interface{} `json:"contactPhone,omitempty"`
	DbcRelation26     int64       `json:"dbcRelation26,omitempty" bson:"dbcRelation26"`
	UpdatedAt         int64       `json:"updatedAt,omitempty" bson:"updatedAt"`
	CreatedAt         int64       `json:"createdAt,omitempty" bson:"createdAt"`
}

// ActivityRecordReq 活动记录
type ActivityRecordReq struct {
	Id                int64       `json:"id,omitempty"`
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
	ContactRole              int         `json:"contactRole"`
	ContactScore             int         `json:"contactScore"`
	CreatedAt                int64       `json:"createdAt" bson:"createdAt"`
	CreatedBy                int64       `json:"createdBy"`
	Depart                   interface{} `json:"depart"`
	DimDepart                int64       `json:"dimDepart"`
	DoNotDisturb             bool        `json:"doNotDisturb"`
	Email                    interface{} `json:"email"`
	EntityType               int64       `json:"entityType"`
	Gender                   int         `json:"gender"`
	LockStatus               int         `json:"lockStatus"`
	Mobile                   string      `json:"mobile"`
	OwnerId                  int64       `json:"ownerId"`
	Phone                    string      `json:"phone"`
	Post                     interface{} `json:"post"`
	RecentActivityRecordTime interface{} `json:"recentActivityRecordTime"`
	State                    interface{} `json:"state"`
	UpdatedAt                int64       `json:"updatedAt"`
	UpdatedBy                int64       `json:"updatedBy"`
	ZipCode                  interface{} `json:"zipCode"`
}

// Lead 线索
type Lead struct {
	Id                       int64  `json:"id"`
	EntityType               int64  `json:"entityType"`               // 业务类型
	OwnerId                  int64  `json:"ownerId"`                  // 销售线索所有人
	Status                   int    `json:"status"`                   // 跟进状态
	Name                     string `json:"name"`                     // 姓名
	Gender                   int    `json:"gender"`                   // 性别
	CompanyName              string `json:"companyName"`              // 公司名称
	Depart                   string `json:"depart"`                   // 联系人所在部门
	Post                     string `json:"post"`                     // 联系人职务
	Phone                    string `json:"phone"`                    // 电话
	Mobile                   string `json:"mobile"`                   // 手机
	Email                    string `json:"email"`                    // 电子邮件
	Address                  string `json:"address"`                  // 地址
	LeadSourceId             int64  `json:"leadSourceId"`             // 线索来源
	RecentActivityRecordTime int64  `json:"recentActivityRecordTime"` // 最新活动记录时间
	CampaignId               int64  `json:"campaignId"`               //市场活动
	RecentActivityCreatedBy  int64  `json:"recentActivityCreatedBy"`  //最新跟进人
	CreatedAt                int64  `json:"createdAt"`                //创建日期
	CreatedBy                int64  `json:"createdBy"`                //创建人
	UpdatedAt                int64  `json:"updatedBt"`                //最新修改日
	UpdatedBy                int64  `json:"updatedBy"`                //最新修改人
	Comment                  string `json:"comment"`                  //备注
	LockStatus               int    `json:"lockStatus"`               //锁定状态
	HighSeaId                int64  `json:"highSeaId"`                //所属公海
	ClaimTime                int64  `json:"claimTime"`                //认领日期
	ExpireTime               int64  `json:"expireTime"`               //到期时间
	HighSeaStatus            int    //状态
	DimDepart                int64  //所属部门
	ApplicantId              int64  //审批提交人
	ApprovalStatus           int    //审批状态
	DoNotDisturb             bool   //免打扰
	OpportunityId            int64  //销售机会
	DuplicateFlg             bool   `json:"duplicateFlg"` //疑似查重
	ReleaseDefinition        string //退回原因说明
	ContactId                int64  `json:"contactId"`     //联系人
	AccountId                int64  `json:"accountId"`     //客户
	ReleaseReason            int64  `json:"releaseReason"` //退回原因
	ReleaseTime              int64  `json:"releaseTime"`   //退回时间
	LastOwnerId              int64  `json:"lastOwnerId"`   //最后所有人
	ReturnTimes              int64  `json:"returnTimes"`   //退回次数
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
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	Status         int    `json:"status"` // 状态
	LastestLoginAt int64  `json:"lastestLoginAt"`
	CreatedAt      int64  // 创建时间
	CreatedBy      int64  // 创建人
	UpdatedAt      int64  // 修改时间
	UpdatedBy      int64  // 修改人
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
	CreatedAt                 int64       `json:"createdAt,omitempty" bson:"createdAt"`
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
	UpdatedAt                 int64       `json:"updatedAt,omitempty" bson:"updatedAt"`
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
	CreatedAt        int64       `json:"createdAt" bson:"createdAt"`
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
	Id                int64   `json:"id"`
	ProductName       string  `json:"productName"`
	FscTrackedAsAsset int     // 可转资产
	ParentId          int64   // 产品目录
	OwnerId           int64   // 产品所有人
	PriceUnit         float64 // 标准价格
	Unit              string  // 销售单位
	Description       string  // 产品描述
	CreatedAt         int64   // 创建日期
	UpdatedAt         int64   // 最新修改日
	EnableStatus      int     // 启用状态
	CreatedBy         int64   // 创建人
	UpdatedBy         int64   // 最新修改人
	EntityType        int64   // 业务类型
	DimDepart         int64   // 所属部门
	ApprovalStatus    int     // 审批状态
	ApplicantId       int64   // 审批提交人
	LockStatus        int     // 锁定状态
	FscProductModel   string  // 产品型号
	FscProductSpec    string  // 产品规格
	FscProductType    int     // 产品类型
	FscNeedReturn     int     // 需要回收
	CustomItem139__c  string  // 标准报价
	CustomItem138__c  float64 // 收费比例
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
			MaxLength             int         `json:"maxLength"`
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
