package ex_game

import (
	"errors"
	"fmt"
)

type User struct {
	id    string
	pass  string
	score int
}

func NewUsers(id string, pass string) *User { // 새로운 유저 생성
	user := User{id: id, pass: pass, score: 0}
	return &user
}

func (a *User) Scoreediter(newscore int) error { // 유저의 스코어 변경
	if newscore < 0 { // 스코어 변경값이 음수이면 에러 반환
		return errors.New("do not")
	}
	a.score = newscore
	return nil
}
func (a *User) Outputid() string { // User 인스턴스의 id값을 반환
	return a.id
}
func (a *User) Outputscore() int { // User 인스턴스의 score값을 반환
	return a.score
}
func (a User) String() string { // fmt 자동실행
	return fmt.Sprintln(a.Outputid(), "'s score is : ", a.Outputscore())
}
