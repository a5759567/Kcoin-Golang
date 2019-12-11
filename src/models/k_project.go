package models

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"

	"github.com/astaxie/beego/orm"
)

//查询并以json形式返回所有的项目信息
func GetAllProjectsInfo() (string, error) {
	o := orm.NewOrm()
	o.Using("default")

	var projectsInfo ProjectInfo
	projectsInfo.ErrorCode = "default Error"

	/******************************************query all projects************************************************/
	queryProjectSql := `SELECT project_id , project_name , project_url , project_cover_url FROM "K_Project" `
	/***********************************************************************************************************/

	_, err := o.Raw(queryProjectSql).QueryRows(&projectsInfo.Data)

	/******************************************query menberList in one project**********************************/
	queryUsersInProjectSql := `select u.k_user_id,u.user_name,u.head_shot_url
								from "K_User" u left join "K_User_in_Project" up on u.k_user_id=up.user_id
       							where up.project_id=?`
	/**********************************************************************************************************/
	for _, v := range projectsInfo.Data {
		var memberList []*UserData
		_, err := o.Raw(queryUsersInProjectSql, v.ProjectId).QueryRows(&memberList)
		if err != nil {
			fmt.Println(err.Error())
			return fmt.Sprint(projectsInfo), err
		}
		v.MemberList = memberList
	}
	projectsInfo.ErrorCode = "0"
	res, err := json.Marshal(&projectsInfo)
	if err != nil {
		fmt.Println(err.Error())
		return fmt.Sprint(projectsInfo), err
	}
	return string(res), nil
}
