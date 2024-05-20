// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

type (
	IMiniUser interface{}
)

var (
	localMiniUser IMiniUser
)

func MiniUser() IMiniUser {
	if localMiniUser == nil {
		panic("implement not found for interface IMiniUser, forgot register?")
	}
	return localMiniUser
}

func RegisterMiniUser(i IMiniUser) {
	localMiniUser = i
}
