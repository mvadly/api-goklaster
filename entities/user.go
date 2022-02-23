package entities

type User struct {
	ID           int32  `gorm:"column:id;type:int(11) AUTO_INCREMENT NOT NULL;primaryKey" json:"id"`
	Username     string `gorm:"column:username;type:varchar(255) NOT NULL;unique" json:"username"`
	Password     string `gorm:"column:password;type:varchar(255) NOT NULL;" json:"password"`
	Uppwd        int    `gorm:"column:uppwd;type:tinyint(1) NOT NULL;" json:"uppwd"`
	Permission   int    `gorm:"column:permission;type:tinyint(1) NOT NULL;" json:"permission"`
	Notif        int    `gorm:"column:notif;type:tinyint(1) NOT NULL;" json:"notif"`
	ApproveLevel int    `gorm:"column:approve_level;type:tinyint(1) NOT NULL;" json:"approve_level"`
}

func (User) TableName() string {
	return "user"
}
