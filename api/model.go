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
