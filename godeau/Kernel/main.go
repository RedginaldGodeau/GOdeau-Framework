package Kernel

import (
	"GOdeau/config/route"
	"GOdeau/godeau/Gobase"
	"net/http"
)

type Kernel struct {
	DB Gobase.Database
}

func (k *Kernel) Entity(entity interface{}) error {
	err := k.DB.ORM.AutoMigrate(&entity)
	if err != nil {
		return err
	}

	return nil
}

func ListenAndServe() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func loader() Kernel {
	
	route.Init()
	return Kernel{}
}

func Run() {
	loader()
	ListenAndServe()
}
