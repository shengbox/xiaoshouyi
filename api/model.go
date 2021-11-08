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
	Id                      int64                    `json:"id,omitempty"`
	Content                 string                   `json:"content,omitempty" bson:"content"`                                 // 活动记录内容
	ContactId               int64                    `json:"contactId,omitempty" bson:"contactId"`                             // 联系人
	DimDepart               int64                    `json:"dimDepart,omitempty" bson:"dimDepart"`                             // 所属部门
	OwnerId                 int64                    `json:"ownerId,omitempty" bson:"ownerId"`                                 // 发布人
	CreatedAt               int64                    `json:"createdAt,omitempty" bson:"createdAt"`                             // 发布时间
	StartTime               int64                    `json:"startTime,omitempty" bson:"startTime"`                             // 活动时间
	ActivityRecordFrom      int64                    `json:"activityRecordFrom,omitempty" bson:"activityRecordFrom"`           // 来自
	EntityType              int64                    `json:"entityType,omitempty" bson:"entityType"`                           // 活动记录类型
	ActivityRecordFrom_data int64                    `json:"activityRecordFrom_data,omitempty" bson:"activityRecordFrom_data"` // 来自_data
	SaleStageId             int64                    `json:"saleStageId,omitempty" bson:"saleStageId"`                         // 商机阶段
	ContactName             string                   `json:"contactName,omitempty" bson:"contactName"`                         // 联系人姓名
	ContactPhone            string                   `json:"contactPhone,omitempty" bson:"contactPhone"`                       // 联系人电话
	IntentionalDegree       int                      `json:"intentionalDegree,omitempty" bson:"intentionalDegree"`             // 意向程度
	NeedFollow              bool                     `json:"needFollow,omitempty" bson:"needFollow"`                           // 是否再次跟进
	NextCallTime            int64                    `json:"nextCallTime,omitempty" bson:"nextCallTime"`                       // 下次通话时间
	Location                string                   `json:"location,omitempty" bson:"location"`                               // 位置
	LocationDetail          string                   `json:"locationDetail,omitempty" bson:"locationDetail"`                   // 位置详情
	Longitude               float64                  `json:"longitude,omitempty" bson:"longitude"`                             // 经度
	Latitude                float64                  `json:"latitude,omitempty" bson:"latitude"`                               // 纬度
	UpdatedAt               int64                    `json:"updatedAt,omitempty" bson:"updatedAt"`                             // 最新修改日
	DbcRelation26           int64                    `json:"dbcRelation26,omitempty" bson:"dbcRelation26"`                     // 客户
	ContentPic              []map[string]interface{} `json:"contentPic,omitempty" bson:"contentPic"`                           // 活动记录内容图片
	ContentFile             []map[string]interface{} `json:"contentFile,omitempty" bson:"contentFile"`                         // 活动记录内容附件
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
	Id                       int64   `json:"id"`
	OwnerId                  int64   `json:"ownerId,omitempty" bson:"ownerId"`                                   // 联系人所有人
	EntityType               int64   `json:"entityType,omitempty" bson:"entityType"`                             // 联系人类型
	ContactName              string  `json:"contactName,omitempty" bson:"contactName"`                           // 姓名
	AccountId                int64   `json:"accountId,omitempty" bson:"accountId"`                               // 公司名称
	Depart                   string  `json:"depart,omitempty" bson:"depart"`                                     // 联系人所在部门
	Post                     string  `json:"post,omitempty" bson:"post"`                                         // 职务
	Phone                    string  `json:"phone,omitempty" bson:"phone"`                                       // 电话
	Mobile                   string  `json:"mobile,omitempty" bson:"mobile"`                                     // 手机
	Email                    string  `json:"email,omitempty" bson:"email"`                                       // 电子邮件
	State                    string  `json:"state,omitempty" bson:"state"`                                       // 省份
	Address                  string  `json:"address,omitempty" bson:"address"`                                   // 地址
	ZipCode                  string  `json:"zipCode,omitempty" bson:"zipCode"`                                   // 邮政编码
	Gender                   int     `json:"gender,omitempty" bson:"gender"`                                     // 性别
	ContactBirthday          int64   `json:"contactBirthday,omitempty" bson:"contactBirthday"`                   // 出生日期
	RecentActivityRecordTime int64   `json:"recentActivityRecordTime,omitempty" bson:"recentActivityRecordTime"` // 最新活动记录时间
	CreatedAt                int64   `json:"createdAt,omitempty" bson:"createdAt"`                               // 创建日期
	CreatedBy                int64   `json:"createdBy,omitempty" bson:"createdBy"`                               // 创建人
	UpdatedAt                int64   `json:"updatedAt,omitempty" bson:"updatedAt"`                               // 最新修改日
	UpdatedBy                int64   `json:"updatedBy,omitempty" bson:"updatedBy"`                               // 最新修改人
	Comment                  string  `json:"comment,omitempty" bson:"comment"`                                   // 备注
	DimDepart                int64   `json:"dimDepart,omitempty" bson:"dimDepart"`                               // 所属部门
	ApplicantId              int64   `json:"applicantId,omitempty" bson:"applicantId"`                           // 审批提交人
	ApprovalStatus           int     `json:"approvalStatus,omitempty" bson:"approvalStatus"`                     // 审批状态
	LockStatus               int     `json:"lockStatus,omitempty" bson:"lockStatus"`                             // 锁定状态
	DoNotDisturb             bool    `json:"doNotDisturb,omitempty" bson:"doNotDisturb"`                         // 免打扰
	ContactRole              int     `json:"contactRole,omitempty" bson:"contactRole"`                           // 联系人角色
	ContactScore             float64 `json:"contactScore,omitempty" bson:"contactScore"`                         // 联系人得分
}

// Lead 线索
type Lead struct {
	Id                       int64   `json:"id"`
	EntityType               int64   `json:"entityType,omitempty" bson:"entityType"`                             // 业务类型
	OwnerId                  int64   `json:"ownerId,omitempty" bson:"ownerId"`                                   // 销售线索所有人
	Status                   int     `json:"status,omitempty" bson:"status"`                                     // 跟进状态
	Name                     string  `json:"name,omitempty" bson:"name"`                                         // 姓名
	Gender                   int     `json:"gender,omitempty" bson:"gender"`                                     // 性别
	CompanyName              string  `json:"companyName,omitempty" bson:"companyName"`                           // 公司名称
	Depart                   string  `json:"depart,omitempty" bson:"depart"`                                     // 联系人所在部门
	Post                     string  `json:"post,omitempty" bson:"post"`                                         // 联系人职务
	Phone                    string  `json:"phone,omitempty" bson:"phone"`                                       // 电话
	Mobile                   string  `json:"mobile,omitempty" bson:"mobile"`                                     // 手机
	Email                    string  `json:"email,omitempty" bson:"email"`                                       // 电子邮件
	FState                   int     `json:"fState,omitempty" bson:"fState"`                                     // 省
	FCity                    int     `json:"fCity,omitempty" bson:"fCity"`                                       // 市
	FDistrict                int     `json:"fDistrict,omitempty" bson:"fDistrict"`                               // 区
	Address                  string  `json:"address,omitempty" bson:"address"`                                   // 地址
	LeadSourceId             int64   `json:"leadSourceId,omitempty" bson:"leadSourceId"`                         // 线索来源
	RecentActivityRecordTime int64   `json:"recentActivityRecordTime,omitempty" bson:"recentActivityRecordTime"` // 最新活动记录时间
	CampaignId               int64   `json:"campaignId,omitempty" bson:"campaignId"`                             // 市场活动
	RecentActivityCreatedBy  int64   `json:"recentActivityCreatedBy,omitempty" bson:"recentActivityCreatedBy"`   // 最新跟进人
	CreatedAt                int64   `json:"createdAt,omitempty" bson:"createdAt"`                               // 创建日期
	CreatedBy                int64   `json:"createdBy,omitempty" bson:"createdBy"`                               // 创建人
	UpdatedAt                int64   `json:"updatedAt,omitempty" bson:"updatedAt"`                               // 最新修改日
	UpdatedBy                int64   `json:"updatedBy,omitempty" bson:"updatedBy"`                               // 最新修改人
	Comment                  string  `json:"comment,omitempty" bson:"comment"`                                   // 备注
	LockStatus               int     `json:"lockStatus,omitempty" bson:"lockStatus"`                             // 锁定状态
	HighSeaId                int64   `json:"highSeaId,omitempty" bson:"highSeaId"`                               // 所属公海
	ClaimTime                int64   `json:"claimTime,omitempty" bson:"claimTime"`                               // 认领日期
	ExpireTime               int64   `json:"expireTime,omitempty" bson:"expireTime"`                             // 到期时间
	HighSeaStatus            int     `json:"highSeaStatus,omitempty" bson:"highSeaStatus"`                       // 状态
	DimDepart                int64   `json:"dimDepart,omitempty" bson:"dimDepart"`                               // 所属部门
	BdType                   int     `json:"bdType,omitempty" bson:"bdType"`                                     // 大数据类型
	ApplicantId              int64   `json:"applicantId,omitempty" bson:"applicantId"`                           // 审批提交人
	ApprovalStatus           int     `json:"approvalStatus,omitempty" bson:"approvalStatus"`                     // 审批状态
	DoNotDisturb             bool    `json:"doNotDisturb,omitempty" bson:"doNotDisturb"`                         // 免打扰
	OpportunityId            int64   `json:"opportunityId,omitempty" bson:"opportunityId"`                       // 销售机会
	Longitude                float64 `json:"longitude,omitempty" bson:"longitude"`                               // 经度
	DuplicateFlg             bool    `json:"duplicateFlg,omitempty" bson:"duplicateFlg"`                         // 疑似查重
	Latitude                 float64 `json:"latitude,omitempty" bson:"latitude"`                                 // 纬度
	ReleaseDefinition        string  `json:"releaseDefinition,omitempty" bson:"releaseDefinition"`               // 退回原因说明
	ContactId                int64   `json:"contactId,omitempty" bson:"contactId"`                               // 联系人
	AccountId                int64   `json:"accountId,omitempty" bson:"accountId"`                               // 客户
	ReleaseReason            int64   `json:"releaseReason,omitempty" bson:"releaseReason"`                       // 退回原因
	ReleaseTime              int64   `json:"releaseTime,omitempty" bson:"releaseTime"`                           // 退回时间
	LastOwnerId              int64   `json:"lastOwnerId,omitempty" bson:"lastOwnerId"`                           // 最后所有人
	ReturnTimes              int64   `json:"returnTimes,omitempty" bson:"returnTimes"`                           // 退回次数
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
	CreatedAt      int64  `json:"createdAt" bson:"createdAt"` // 创建时间
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
	ExpireTime                interface{} `json:"expireTime,omitempty" bson:"expireTime"`
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
	EndDate          int64       `json:"endDate" bson:"endDate"`
	Id               int64       `json:"id"`
	OwnerId          int64       `json:"ownerId"`
	StartDate        int64       `json:"startDate" bson:"startDate"`
	Status           int         `json:"status"`
	Title            string      `json:"title"`
	UpdatedAt        int64       `json:"updatedAt" bson:"updatedAt"`
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
	Id                int64                    `json:"id"`
	FscTrackedAsAsset int                      `json:"fscTrackedAsAsset,omitempty" bson:"fscTrackedAsAsset"` // 可转资产
	ProductName       string                   `json:"productName,omitempty" bson:"productName"`             // 产品名称
	ParentId          int64                    `json:"parentId,omitempty" bson:"parentId"`                   // 产品目录
	OwnerId           int64                    `json:"ownerId,omitempty" bson:"ownerId"`                     // 产品所有人
	PriceUnit         float64                  `json:"priceUnit,omitempty" bson:"priceUnit"`                 // 标准价格
	Unit              string                   `json:"unit,omitempty" bson:"unit"`                           // 销售单位
	FileImage1        []map[string]interface{} `json:"fileImage1,omitempty" bson:"fileImage1"`               // 产品图片
	Description       string                   `json:"description,omitempty" bson:"description"`             // 产品描述
	CreatedAt         int64                    `json:"createdAt,omitempty" bson:"createdAt"`                 // 创建日期
	UpdatedAt         int64                    `json:"updatedAt,omitempty" bson:"updatedAt"`                 // 最新修改日
	EnableStatus      int                      `json:"enableStatus,omitempty" bson:"enableStatus"`           // 启用状态
	CreatedBy         int64                    `json:"createdBy,omitempty" bson:"createdBy"`                 // 创建人
	UpdatedBy         int64                    `json:"updatedBy,omitempty" bson:"updatedBy"`                 // 最新修改人
	EntityType        int64                    `json:"entityType,omitempty" bson:"entityType"`               // 业务类型
	DimDepart         int64                    `json:"dimDepart,omitempty" bson:"dimDepart"`                 // 所属部门
	ApprovalStatus    int                      `json:"approvalStatus,omitempty" bson:"approvalStatus"`       // 审批状态
	ApplicantId       int64                    `json:"applicantId,omitempty" bson:"applicantId"`             // 审批提交人
	LockStatus        int                      `json:"lockStatus,omitempty" bson:"lockStatus"`               // 锁定状态
	FscProductModel   string                   `json:"fscProductModel,omitempty" bson:"fscProductModel"`     // 产品型号
	FscProductSpec    string                   `json:"fscProductSpec,omitempty" bson:"fscProductSpec"`       // 产品规格
	FscProductType    int                      `json:"fscProductType,omitempty" bson:"fscProductType"`       // 产品类型
	FscNeedReturn     int                      `json:"fscNeedReturn,omitempty" bson:"fscNeedReturn"`         // 需要回收
	// CustomItem139__c  string                   `json:"customItem139__c,omitempty" bson:"customItem139__c"`   // 标准报价
	// CustomItem138__c  float64                  `json:"customItem138__c,omitempty" bson:"customItem138__c"`   // 收费比例

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

// Opportunity 销售机会
type Opportunity struct {
	Id                       int64   `json:"id,omitempty"`
	EntityType               int64   `json:"entityType,omitempty" bson:"entityType"`                             // 业务类型
	OwnerId                  int64   `json:"ownerId,omitempty" bson:"ownerId"`                                   // 销售机会所有人
	OpportunityName          string  `json:"opportunityName,omitempty" bson:"opportunityName"`                   // 机会名称
	AccountId                int64   `json:"accountId,omitempty" bson:"accountId"`                               // 客户名称
	PriceId                  int64   `json:"priceId,omitempty" bson:"priceId"`                                   // 价格表名称
	OpportunityType          int     `json:"opportunityType,omitempty" bson:"opportunityType"`                   // 机会类型
	Money                    float64 `json:"money,omitempty" bson:"money"`                                       // 销售金额
	SaleStageId              int64   `json:"saleStageId,omitempty" bson:"saleStageId"`                           // 销售阶段
	LostStageId              int64   `json:"lostStageId,omitempty" bson:"lostStageId"`                           // 输单阶段
	WinRate                  int64   `json:"winRate,omitempty" bson:"winRate"`                                   // 赢率
	Reason                   int     `json:"reason,omitempty" bson:"reason"`                                     // 输单原因
	ReasonDesc               string  `json:"reasonDesc,omitempty" bson:"reasonDesc"`                             // 输单描述
	CloseDate                int64   `json:"closeDate,omitempty" bson:"closeDate"`                               // 结单日期
	CommitmentFlg            int     `json:"commitmentFlg,omitempty" bson:"commitmentFlg"`                       // 承诺
	SourceId                 int64   `json:"sourceId,omitempty" bson:"sourceId"`                                 // 机会来源
	ProjectBudget            float64 `json:"projectBudget,omitempty" bson:"projectBudget"`                       // 项目预算
	ActualCost               float64 `json:"actualCost,omitempty" bson:"actualCost"`                             // 实际花费
	Product                  string  `json:"product,omitempty" bson:"product"`                                   // 产品
	StageUpdatedAt           int64   `json:"stageUpdatedAt,omitempty" bson:"stageUpdatedAt"`                     // 阶段更新时间
	RecentActivityRecordTime int64   `json:"recentActivityRecordTime,omitempty" bson:"recentActivityRecordTime"` // 最新活动记录时间
	CreatedAt                int64   `json:"createdAt,omitempty" bson:"createdAt"`                               // 创建日期
	CreatedBy                int64   `json:"createdBy,omitempty" bson:"createdBy"`                               // 创建人
	UpdatedAt                int64   `json:"updatedAt,omitempty" bson:"updatedAt"`                               // 最新修改日
	UpdatedBy                int64   `json:"updatedBy,omitempty" bson:"updatedBy"`                               // 最新修改人
	Comment                  string  `json:"comment,omitempty" bson:"comment"`                                   // 备注
	DimDepart                int64   `json:"dimDepart,omitempty" bson:"dimDepart"`                               // 所属部门
	ForecastCategory         int     `json:"forecastCategory,omitempty" bson:"forecastCategory"`                 // 阶段分类
	LockStatus               int     `json:"lockStatus,omitempty" bson:"lockStatus"`                             // 锁定状态
	CampaignId               int64   `json:"campaignId,omitempty" bson:"campaignId"`                             // 市场活动
	ApprovalStatus           int     `json:"approvalStatus,omitempty" bson:"approvalStatus"`                     // 审批状态
	Status                   int     `json:"status,omitempty" bson:"status"`                                     // 状态
	LeadId                   int64   `json:"leadId,omitempty" bson:"leadId"`                                     // 销售线索
	OpportunityScore         float64 `json:"opportunityScore,omitempty" bson:"opportunityScore"`                 // 商机得分
	DuplicateFlg             bool    `json:"duplicateFlg,omitempty" bson:"duplicateFlg"`                         // 疑似查重
	FcastMoney               float64 `json:"fcastMoney,omitempty" bson:"fcastMoney"`                             // 预测金额
	ApplicantId              int64   `json:"applicantId,omitempty" bson:"applicantId"`                           // 审批提交人
	// CustomItem158__c         float64 `json:"customItem158__c,omitempty" bson:"customItem158__c"`                 // 收费比例
	// CustomItem159__c         float64 `json:"customItem159__c,omitempty" bson:"customItem159__c"`                 // 预估营收费用
	// CustomItem160__c         float64 `json:"customItem160__c,omitempty" bson:"customItem160__c"`                 // 预估净收入
	// CustomItem161__c         int64   `json:"customItem161__c,omitempty" bson:"customItem161__c"`                 // 预估服务人数
	// CustomItem162__c         float64 `json:"customItem162__c,omitempty" bson:"customItem162__c"`                 // 服务费报价
	// CustomItem163__c         float64 `json:"customItem163__c,omitempty" bson:"customItem163__c"`                 // 预估服务费收入(月)
	// CustomItem164__c         int64   `json:"customItem164__c,omitempty" bson:"customItem164__c"`                 // 需求产品
	// CustomItem165__c         int64   `json:"customItem165__c,omitempty" bson:"customItem165__c"`                 // 预计启动时间
}

// PaymentApplicationPlan 收款计划
type PaymentApplicationPlan struct {
	Id                int64   `json:"id,omitempty"`
	Name              string  `json:"name,omitempty" bson:"name"`                           // 收款计划编号
	EntityType        int64   `json:"entityType,omitempty" bson:"entityType"`               // 业务类型
	TransactionDate   int64   `json:"transactionDate,omitempty" bson:"transactionDate"`     // 计划收款日期
	FirstPaymentDate  int64   `json:"firstPaymentDate,omitempty" bson:"firstPaymentDate"`   // 首次收款日期
	OwnerId           int64   `json:"ownerId,omitempty" bson:"ownerId"`                     // 所有人
	DimDepart         int64   `json:"dimDepart,omitempty" bson:"dimDepart"`                 // 所属部门
	LastedPaymentDate int64   `json:"lastedPaymentDate,omitempty" bson:"lastedPaymentDate"` // 最近收款日期
	AccountId         int64   `json:"accountId,omitempty" bson:"accountId"`                 // 客户
	WorkflowStageName string  `json:"workflowStageName,omitempty" bson:"workflowStageName"` // 工作流阶段名称
	ContractId        int64   `json:"contractId,omitempty" bson:"contractId"`               // 合同
	Stage             int     `json:"stage,omitempty" bson:"stage"`                         // 收款阶段
	Status            string  `json:"status,omitempty" bson:"status"`                       // 收款计划状态
	Period            int64   `json:"period,omitempty" bson:"period"`                       // 收款期次
	CreatedBy         int64   `json:"createdBy,omitempty" bson:"createdBy"`                 // 创建人
	CreatedAt         int64   `json:"createdAt,omitempty" bson:"createdAt"`                 // 创建日期
	Percentage        float64 `json:"percentage,omitempty" bson:"percentage"`               // 计划收款比例
	Amount            float64 `json:"amount,omitempty" bson:"amount"`                       // 计划收款金额
	UpdatedBy         int64   `json:"updatedBy,omitempty" bson:"updatedBy"`                 // 修改人
	UpdatedAt         int64   `json:"updatedAt,omitempty" bson:"updatedAt"`                 // 修改日期
	OverdueFlg        bool    `json:"overdueFlg,omitempty" bson:"overdueFlg"`               // 是否逾期
	OrderId           int64   `json:"orderId,omitempty" bson:"orderId"`                     // 订单
	InvoiceAmount     float64 `json:"invoiceAmount,omitempty" bson:"invoiceAmount"`         // 实际应收账金额
	PaymentAmount     float64 `json:"paymentAmount,omitempty" bson:"paymentAmount"`         // 已收款金额
	ApplicantId       int64   `json:"applicantId,omitempty" bson:"applicantId"`             // 审批提交人
	InvoicePercentage float64 `json:"invoicePercentage,omitempty" bson:"invoicePercentage"` // 实际应收账比例
	ApprovalStatus    int     `json:"approvalStatus,omitempty" bson:"approvalStatus"`       // 审批状态
	PaymentPercentage float64 `json:"paymentPercentage,omitempty" bson:"paymentPercentage"` // 实际收款比例
	InvoiceBalance    float64 `json:"invoiceBalance,omitempty" bson:"invoiceBalance"`       // 应收余额
	ReceiptAmount     float64 `json:"receiptAmount,omitempty" bson:"receiptAmount"`         // 已开票金额
	ReceiptBalance    float64 `json:"receiptBalance,omitempty" bson:"receiptBalance"`       // 未开票金额
	ActualPercentage  float64 `json:"actualPercentage,omitempty" bson:"actualPercentage"`   // 计划收款金额实际占比
}

// Receipt 发票
type Receipt struct {
	Id                       int64   `json:"id,omitempty"`
	ContactPhone             string  `json:"contactPhone,omitempty" bson:"contactPhone"`                         // 联系电话
	CompanyName              string  `json:"companyName,omitempty" bson:"companyName"`                           // 开票单位
	Number                   string  `json:"number,omitempty" bson:"number"`                                     // 票据号
	WorkflowStageName        string  `json:"workflowStageName,omitempty" bson:"workflowStageName"`               // 工作流阶段名称
	Address                  string  `json:"address,omitempty" bson:"address"`                                   // 邮寄地址
	ReceiptAmount            float64 `json:"receiptAmount,omitempty" bson:"receiptAmount"`                       // 开票金额
	Postcode                 string  `json:"postcode,omitempty" bson:"postcode"`                                 // 邮编
	CreatedBy                int64   `json:"createdBy,omitempty" bson:"createdBy"`                               // 创建人
	Status                   int     `json:"status,omitempty" bson:"status"`                                     // 发票状态
	Type                     int     `json:"type,omitempty" bson:"type"`                                         // 发票种类
	CreatedAt                int64   `json:"createdAt,omitempty" bson:"createdAt"`                               // 创建日期
	UpdatedBy                int64   `json:"updatedBy,omitempty" bson:"updatedBy"`                               // 修改人
	TransactionDate          int64   `json:"transactionDate,omitempty" bson:"transactionDate"`                   // 发票日期
	UpdatedAt                int64   `json:"updatedAt,omitempty" bson:"updatedAt"`                               // 修改日期
	EffectiveDate            int64   `json:"effectiveDate,omitempty" bson:"effectiveDate"`                       // 生效日期
	Description              string  `json:"description,omitempty" bson:"description"`                           // 发票描述
	ApplicantId              int64   `json:"applicantId,omitempty" bson:"applicantId"`                           // 审批提交人
	Amount                   float64 `json:"amount,omitempty" bson:"amount"`                                     // 发票金额
	CancellationDate         int64   `json:"cancellationDate,omitempty" bson:"cancellationDate"`                 // 作废日期
	ApprovalStatus           int     `json:"approvalStatus,omitempty" bson:"approvalStatus"`                     // 审批状态
	CancellationReason       int     `json:"cancellationReason,omitempty" bson:"cancellationReason"`             // 作废原因
	PaymentApplicationPlanId int64   `json:"paymentApplicationPlanId,omitempty" bson:"paymentApplicationPlanId"` // 收款计划
	ContactId                int64   `json:"contactId,omitempty" bson:"contactId"`                               // 联系人
	OrderId                  int64   `json:"orderId,omitempty" bson:"orderId"`                                   // 订单
	SourceType               int     `json:"sourceType,omitempty" bson:"sourceType"`                             // 来源
	InvoiceId                int64   `json:"invoiceId,omitempty" bson:"invoiceId"`                               // 应收单
}

// VisitRecord 拜访记录
type VisitRecord struct {
	Id                   int64                    `json:"id,omitempty"`
	Name                 string                   `json:"name,omitempty" bson:"name"`                                 // 拜访记录
	EntityType           int64                    `json:"entityType,omitempty" bson:"entityType"`                     // 业务类型
	OwnerId              int64                    `json:"ownerId,omitempty" bson:"ownerId"`                           // 所有人
	DimDepart            int64                    `json:"dimDepart,omitempty" bson:"dimDepart"`                       // 所属部门
	WorkflowStageName    string                   `json:"workflowStageName,omitempty" bson:"workflowStageName"`       // 工作流阶段名称
	ApplicantId          int64                    `json:"applicantId,omitempty" bson:"applicantId"`                   // 审批提交人
	CreatedBy            int64                    `json:"createdBy,omitempty" bson:"createdBy"`                       // 创建人
	CreatedAt            int64                    `json:"createdAt,omitempty" bson:"createdAt"`                       // 创建日期
	DoorPicture          []map[string]interface{} `json:"doorPicture,omitempty" bson:"doorPicture"`                   // 门头照
	UpdatedBy            int64                    `json:"updatedBy,omitempty" bson:"updatedBy"`                       // 修改人
	UpdatedAt            int64                    `json:"updatedAt,omitempty" bson:"updatedAt"`                       // 修改日期
	LockStatus           int                      `json:"lockStatus,omitempty" bson:"lockStatus"`                     // 锁定状态
	ApprovalStatus       int                      `json:"approvalStatus,omitempty" bson:"approvalStatus"`             // 审批状态
	SignInDate           int64                    `json:"signInDate,omitempty" bson:"signInDate"`                     // 签到时间
	SignInStatus         int                      `json:"signInStatus,omitempty" bson:"signInStatus"`                 // 签到定位状态
	SignOutDate          int64                    `json:"signOutDate,omitempty" bson:"signOutDate"`                   // 签退时间
	SignOutStatus        int                      `json:"signOutStatus,omitempty" bson:"signOutStatus"`               // 签退定位状态
	SignOutAddress       string                   `json:"signOutAddress,omitempty" bson:"signOutAddress"`             // 签退地址
	SignInAddress        string                   `json:"signInAddress,omitempty" bson:"signInAddress"`               // 签到地址
	Schedule             int64                    `json:"schedule,omitempty" bson:"schedule"`                         // 拜访日程
	VisitTime            string                   `json:"visitTime,omitempty" bson:"visitTime"`                       // 拜访时长
	Longitude            float64                  `json:"longitude,omitempty" bson:"longitude"`                       // 经度
	Latitude             float64                  `json:"latitude,omitempty" bson:"latitude"`                         // 纬度
	VisitDetailId        int64                    `json:"visitDetailId,omitempty" bson:"visitDetailId"`               // 关联拜访日程
	CompVisitObject      int64                    `json:"compVisitObject,omitempty" bson:"compVisitObject"`           // 拜访对象
	VisitModel           int                      `json:"visitModel,omitempty" bson:"visitModel"`                     // 拜访模板
	CompVisitObject_data int64                    `json:"compVisitObject_data,omitempty" bson:"compVisitObject_data"` // 拜访对象_data
	VisitType            int                      `json:"visitType,omitempty" bson:"visitType"`                       // 拜访类型
	VisitRecordStatus    int                      `json:"visitRecordStatus,omitempty" bson:"visitRecordStatus"`       // 拜访记录状态
	OriginalVisitRecord  int64                    `json:"originalVisitRecord,omitempty" bson:"originalVisitRecord"`   // 原拜访记录
	VisitLocationItemKey string                   `json:"visitLocationItemKey,omitempty" bson:"visitLocationItemKey"` // 拜访地址字段
	VisitObjectAddress   string                   `json:"visitObjectAddress,omitempty" bson:"visitObjectAddress"`     // 拜访对象地址
	SignOutLatitude      float64                  `json:"signOutLatitude,omitempty" bson:"signOutLatitude"`           // 签退纬度
	SignOutLongitude     float64                  `json:"signOutLongitude,omitempty" bson:"signOutLongitude"`         // 签退经度
	CustomItem1__c       int64                    `json:"customItem1__c,omitempty" bson:"customItem1__c"`             // 客户名称
}
