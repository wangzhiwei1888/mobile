package mobile

// Mobile 结构体定义了一个中国手机的标准数据类型
//
type Mobile struct {
	Prefix   string `json:"prefix" xml:"prefix" bson:"prefix"`       // 前缀，比如 139, 183 等
	Segment  string `json:"segment" xml:"segment" bson:"segment"`    // 号码段
	ISP      string `json:"isp" xml:"isp" bson:"isp"`                // 服务供应商
	Province string `json:"province" xml:"province" bson:"province"` // 省份
	City     string `json:"city" xml:"city" bson:"city"`             // 城市
	ZipCode  string `json:"zipCode" xml:"zipCode" bson:"zipCode"`    // 邮政编码
	AreaCode string `json:"areaCode" xml:"areaCode" bson:"areaCode"` // 区号
	Type     string `json:"type" xml:"type" bson:"type"`             // 卡类型
}
