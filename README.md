# Amoeba
 Amoeba provides a flexible and automatic transformation layer for complex data output,
 especial used in RESTful APIs.
 
# Install

```
go get github.com/IamBusy/amoeba
```

# Usage

Take user-role-permission model as an example.

First import the package, and define your custom-transformers.
Theses transformers should inherit the `go-transformer.Tran`

In the `init` func, register these transformers with their names. You can give any
string as the name.

```apple js
import github.com/IamBusy/amoeba"

type UserTransformer struct {
	amoeba.Tran
}

type RoleTransformer struct {
	amoeba.Tran
}

type PermissionTransformer struct {
	amoeba.Tran
}


func init() {
	go_transformer.RegisterTransformer("user",&UserTransformer{})
	go_transformer.RegisterTransformer("role",&RoleTransformer{})
	go_transformer.RegisterTransformer("permission",&PermissionTransformer{})
}

```

Then define the relation function. The first parameter of `Include` is the key word, 
which should be matched with `includeStr`.

You cann't get the result if you define the key word as "myrole" with `roles.permissions` as `includeStr`.
But `myrole.permissions` is ok.

You can retrieve any parameters from `args` to get the related objects.
And the args was passed when you call the entrance func `Collection` of `Item`


```

func (transformer *UserTransformer) RegisterIncluder() {
	transformer.Include("role", func(transformer go_transformer.Transformer, entity interface{}, includeStr string, args ...interface{}) interface{} {
		entity := entity.(User)
		db := args[0].(*gorm.DB)
		db.Model(&entity).Association("roles").Find(entity.Roles)
		return transformer.Collection(user.Roles,"role", args)
	})
}

func (transformer *RoleTransformer) RegisterIncluder() {
	transformer.Include("permissions", func(transformer go_transformer.Transformer, entity interface{}, includeStr string, args ...interface{}) interface{} {
		entity := entity.(Role)
		db := args[0].(*gorm.DB)
		db.Model(&entity).Association("permissions").Find(entity.Permissions)
		return transformer.Collection(user.Permissions,"permission", args)
	})
}

```

And that's all.


```
user := &User{1,"myname}
fmp.Println(amoeba.Item(user, "user", "roles.permissions"  [, args] ))
```

With the mock data, you can get the result which is a `map[string]inteface{}` struct on call above function.
such as:

```
map[ID:1 Name:william roles:[map[ID:1 Name:admin permissions:[map[ID:1 Name:update-user] map[ID:2 Name:delete-user]]] map[permissions:[map[ID:1 Name:update-user] map[ID:2 Name:delete-user]] ID:2 Name:editor]]]
```

json format:

```
{
	"ID":1,
	"Name":"william",
	"roles":[
		{
			"ID":1,
			"Name":"admin",
			"permissions":[
				{
					"ID":1,
					"Name":"update-user"
				},
				{
					"ID":2,
					"Name":"delete-user"
				}
			]
		},
		{
			"ID":1,
			"Name":"editor"
		}
	]
}

```