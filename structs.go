package mobile

// NetworkIdentifier 定义了网络识别号
type NetworkIdentifier struct {
	Code     string             `json:"code" xml:"code" bson:"code"`             // 网络识别码 如 130， 138， 183， 189等
	ISP      string             `json:"isp" xml:"isp" bson:"isp"`                // 服务供应商
	Type     string             `json:"type" xml:"type" bson:"type"`             // 卡类型
	Segments map[string]Segment `json:"segments" xml:"segments" bson:"segments"` // 号码段列表
}

// Segment 结构体定义了一个中国手机的标准数据类型
type Segment struct {
	Prefix   string `json:"prefix" xml:"prefix" bson:"prefix"`       // 前缀，比如 1390001, 1830002 等
	Province string `json:"province" xml:"province" bson:"province"` // 省份
	City     string `json:"city" xml:"city" bson:"city"`             // 城市
	ZipCode  string `json:"zipCode" xml:"zipCode" bson:"zipCode"`    // 邮政编码
	AreaCode string `json:"areaCode" xml:"areaCode" bson:"areaCode"` // 区号
}

type Mobile struct {
	Name     string `json:"name" xml:"name" bson:"name"`             // 持有者姓名
	Number   string `json:"number" xml:"number" bson:"number"`       // 手机号码
	Province string `json:"province" xml:"province" bson:"province"` // 省份
	City     string `json:"city" xml:"city" bson:"city"`             // 城市
	ZipCode  string `json:"zipCode" xml:"zipCode" bson:"zipCode"`    // 邮政编码
	AreaCode string `json:"areaCode" xml:"areaCode" bson:"areaCode"` // 区号
	ISP      string `json:"isp" xml:"isp" bson:"isp"`                // 服务供应商
	Type     string `json:"type" xml:"type" bson:"type"`             // 卡类型
}
