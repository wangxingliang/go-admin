package models

type SysContent struct {
	Model
	CateId  int    `json:"cateId" gorm:"type:int(11);comment:分类id"`
	Name    string `json:"name" gorm:"type:varchar(255);comment:名称"`
	Status  int    `json:"status" gorm:"type:int(1);comment:状态"`
	Img     string `json:"img" gorm:"type:varchar(255);comment:图片"`
	Content string `json:"content" gorm:"type:text;comment:内容"`
	Remark  string `json:"remark" gorm:"type:varchar(255);comment:备注"`
	Sort    int    `json:"sort" gorm:"type:int(4);comment:排序"`
	ControlBy
	ModelTime
}

func (SysContent) TableName() string {
	return "sys_content"
}
