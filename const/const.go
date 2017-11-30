package _const

import "dotapp_server/models"


const (
	 CODE_SUCCESS ,CODE_FIAL= 1,0
	 JM = "llkj"
)
//AuthList 获取用户列表
var (
	AuthorList = make([]*models.User, 0)
)


func init() {
	list, err := new(models.User).List()
	if err != nil {
		panic(err)
	}
	for _, user := range list {
		AuthorList = append(AuthorList, user)
	}

}
func AddAuthor(args... *models.User)  {
	for _,user:=range args {
		AuthorList = append(AuthorList,user)
	}
}
