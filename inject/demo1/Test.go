package main

import "github.com/facebookgo/inject"

type DBEngine struct{}

func NewDBEngine() *DBEngine {
	return &DBEngine{}
}

type CacheEngine struct{}

func NewCacheEngine() *CacheEngine {
	return &CacheEngine{}
}

type UserDB struct {
	db    *DBEngine    `inject:""`
	cache *CacheEngine `inject:""`
}

type ItemDB struct {
	db    *DBEngine    `inject:""`
	cache *CacheEngine `inject:""`
}

type UserService struct {
	db *UserDB `inject:""`
}

type ItemService struct {
	db *ItemDB `inject:""`
}

type App struct {
	user *UserService `inject:""`
	item *ItemService `inject:""`
}

func main() {
	db := NewDBEngine()
	cache := NewCacheEngine()
	var g inject.Graph
	var app App
	g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: db},
		&inject.Object{Value: cache})
	g.Populate()
	// use app do something
}
