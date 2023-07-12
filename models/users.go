package models

import (
	"context"
	"encoding/json"
	"fmt"

	"example.com/goRedis/db"
)

type User struct {
	Key      string `json:"key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Id       int    `json:"id"`
}

// 存储和读取redis时，不能直接存入结构体，必须进行序列号和反序列化

// 序列化
func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// 反序列化
func (u *User) UnmarshalBinary(data []byte) (err error) {
	return json.Unmarshal(data, u)
}

// 新增数据
func (u *User) Add(ctx *context.Context) (err error) {

	// HMSet 批量设置 map[string]interface{}{"name": "张", "password": "11111", "id": 123}

	fmt.Println("json.=", u)
	// 必须转为[]byte, redis不支持直接存结构体

	// 0 表示key永不过期
	_, err = db.Rdbs.Set(*ctx, u.Key, u, 0).Result()
	if err != nil {
		fmt.Println("db.Rdbs.Set,", err)
		return
	}
	return
}

// 根据key 获取到数据
func (u *User) GetUser(ctx *context.Context) (user1 User, err error) {

	err = db.Rdbs.Get(*ctx, u.Key).Scan(u)
	if err != nil {
		fmt.Println("GetUser error", err)
		return
	}
	user1 = *u
	return
}

// 查询所有
func (u *User) QueryAll(ctx *context.Context) (user1 []User, er error) {
	// keys, err := db.Rdbs.Keys(*ctx, "[1-9]*")).Result()
	keys, err := db.Rdbs.Keys(*ctx, "*").Result()
	fmt.Println("key=", keys)
	if err != nil {
		fmt.Println("db.Rdbs.Keys error:", err)
		return
	}
	for _, key := range keys {
		err = db.Rdbs.Get(*ctx, key).Scan(u)
		if err != nil {
			fmt.Print("Rdbs.Get error", err)
			fmt.Println("key=", key)
			return
		}
		fmt.Println("query=", *u)
		user1 = append(user1, *u)

	}
	return
}

// 删除
func (u *User) DeleteUser(ctx *context.Context) (err bool) {
	// 先判断key是否存在
	if db.Rdbs.Exists(*ctx, u.Key).Val() == 1 {
		_, er := db.Rdbs.Del(*ctx, u.Key).Result()
		if er != nil {
			fmt.Println("DeleteUser error: ", err)
			return false
		}
		return true
	}
	return false
}

// 修改
func (u *User) UpdateUser(ctx *context.Context) (err error) {
	// 先判断key是否存在
	if db.Rdbs.Exists(*ctx, u.Key).Val() == 1 {
		// 必须转为[]byte, redis不支持直接存结构体
		u1, _ := json.Marshal(u)
		_, err = db.Rdbs.Set(*ctx, u.Key, u1, 0).Result()
		if err != nil {
			fmt.Println("UpdateUser error:", err)
			return
		}
	}
	return
}
