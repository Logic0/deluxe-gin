package mysql

type Application struct{
	// 应用 ID
	Appid uint               `gorm:"PRIMARY_KEY;AUTO_INCREMENT:false"`

	Events []Event           `gorm:"many2many:application_event;association_autoupdate:false"`
}

