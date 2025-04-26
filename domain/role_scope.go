package domain

type RoleScope struct {
	RoleScopeId int64 `gorm:"primaryKey"`
	RoleId      int64
	ScopeId     int64
}

type RoleScopeRepository interface {
	Create(*RoleScope) (*RoleScope, error)
	FindByScopeId(int64) ([]*RoleScope, error)
}
