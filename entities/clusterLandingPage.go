package entities

type ClusterLandingPage struct {
	ID  string `gorm:"column:id;type:varchar(255)" json:"id"`
	Src string `gorm:"column:src;type:varchar(255)" json:"src"`
}

func (ClusterLandingPage) TableName() string {
	return "cluster_landing_page"
}
