package im

import "gorm.io/gorm"

type ChatroomType string

const P2P ChatroomType = "p2p"
const Group ChatroomType = "group"

type Chatroom struct {
	gorm.Model
	Name string       `json:"name"`
	Type ChatroomType `json:"type"`
}
