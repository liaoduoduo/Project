/**
  @Author : hanxiaodong
*/

package controller

import "Project/service"

type Application struct {
	Setup *service.ServiceSetup
}

type User struct {
	LoginName	string
	Password	string
	IsAdmin	string
}


var users []User

func init() {

	admin := User{LoginName:"liaoduoyue", Password:"123456", IsAdmin:"T"}
	ldy := User{LoginName:"xiaoming", Password:"123456", IsAdmin:"T"}
	bob := User{LoginName:"alice", Password:"123456", IsAdmin:"F"}
	jack := User{LoginName:"bob", Password:"123456", IsAdmin:"F"}

	users = append(users, admin)
	users = append(users, ldy)
	users = append(users, bob)
	users = append(users, jack)

}