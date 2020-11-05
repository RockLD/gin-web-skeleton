package dao

type Roles struct {
	ID       int64  `gorm:"primary_key;column:id;type:bigint(20);not null"`
	RoleName string `gorm:"column:role_name;type:varchar(255);not null"`
}

func (role Roles) TableName() string {
	return "gws_roles"
}
