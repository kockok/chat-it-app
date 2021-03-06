package test

import (
	"context"
	"log"
	"testing"

	"github.com/gabivlj/chat-it/internals/domain"
	"github.com/gabivlj/chat-it/internals/repository"
)

func BenchmarkFindByIDs(b *testing.B) {
	const benchmarkIterations = 1
	if b.N > benchmarkIterations {
		return
	}
	userRep, _, _, _ := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	users, err := userRep.FindByIDs(context.TODO(), []string{"5ecff485476916b6cc2d180f", "5ecff4b7fac656a1633f04f7"})
	if err != nil {
		b.Error(err)
		b.FailNow()
	}

	for _, user := range users {
		b.Log(*user)
	}
}

func BenchmarkRegister(b *testing.B) {
	const benchmarkIterations = 1
	if b.N > benchmarkIterations {
		return
	}
	userRep, _, _, _ := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	user, _, err := userRep.SaveUser(context.TODO(), &domain.User{Username: "gabivlj053", Password: "123456"})
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	b.Log(user)
}

func BenchmarkUserLog(b *testing.B) {
	userRep, _, _, _ := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	user, err := userRep.FindByID(context.TODO(), "5ecff485476916b6cc2d180f")
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	if user.ID != "5ecff485476916b6cc2d180f" {
		b.Errorf("ERROR, ids don't match")
		b.FailNow()
	}
	b.Log(user)
}

func BenchmarkUserLogIn(b *testing.B) {
	userRep, _, _, _ := repository.NewRepository()
	defer userRep.Disconnect(context.TODO())
	user, session, err := userRep.LogUser(context.TODO(), &domain.User{Password: "123456", Username: "gabivlj02"})
	if err != nil {
		b.Error(err)
		b.FailNow()
	}
	us, err := userRep.VerifySession(session)
	log.Println(user, session, us, err)
}
