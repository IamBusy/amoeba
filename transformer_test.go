package amoeba

import (
	"fmt"
	"testing"
)

type User struct {
	ID   uint64
	Name string
}

type Role struct {
	ID   uint64
	Name string
}

type Permission struct {
	ID   uint64
	Name string
}

type UserTransformer struct {
	Tran
}

type RoleTransformer struct {
	Tran
}

type PermissionTransformer struct {
	Tran
}

func (user *UserTransformer) RegisterIncluder() {
	user.Include("roles", func(transformer Transformer, entity interface{}, includeStr string, args ...interface{}) interface{} {
		roles := []Role{{1, "admin"}, {2, "editor"}}
		return user.Collection(roles, "role", includeStr, args)
	})
}

func (role *RoleTransformer) RegisterIncluder() {
	role.Include("permissions", func(transformer Transformer, entity interface{}, includeStr string, args ...interface{}) interface{} {
		permission := []Permission{{1, "update-user"}, {2, "delete-user"}}
		return role.Collection(permission, "permission", includeStr, args)
	})
}

func TestCollection(t *testing.T) {
	users := []User{{1, "william"}}
	RegisterTransformer("user", &UserTransformer{})
	RegisterTransformer("role", &RoleTransformer{})
	RegisterTransformer("permission", &PermissionTransformer{})
	res := Collection(users, "user", "roles.permissions")
	fmt.Println(res)
}
