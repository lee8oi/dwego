// accounts.go
package main

import (
//"fmt"
)

type Item struct {
	ID, Limit         uint32
	Name, Description string
}

type Space struct {
	Amount uint32
	Item   Item
}

type Storage struct {
	Size  int
	Slots []Space
}

type Character struct {
	ID          uint32
	Name, Title string
	Inv         Storage
}

type Account struct {
	ID                           uint32
	Email, Name, Title, Password string
	Chars                        []Character
}

func (a *Account) Save(path string) {

}

func (a *Account) Load(path string) {

}
