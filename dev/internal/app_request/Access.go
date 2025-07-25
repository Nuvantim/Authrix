package request

type Permission struct {
	Name string "json:name"
}

type Role struct {
	Name         string  "json:name"
	PermissionID []int32 "json:permission_id"
}
