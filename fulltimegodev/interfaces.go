package main

import "fmt"

// Interfaces

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type MongoDBNumberStore struct {
	// mongoDb connection
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(number int) error {
	fmt.Println("store the number into mongo database")
	return nil
}

// switching from mongoDB to Postgres
type PostgresNumberStore struct {
	// postgres connection
}

func (s PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3, 4, 5, 6}, nil
}

func (s PostgresNumberStore) Put(number int) error {
	fmt.Println("store the number into postgres database")
	return nil
}

type ApiServer struct {
	store NumberStorer
}

func mainInterface() {
	apiServer := ApiServer{
		store: PostgresNumberStore{},
	}

	// logic
	if err := apiServer.store.Put(1); err != nil {
		panic(err)
	}

	numbers, err := apiServer.store.GetAll()

	if err != nil {
		panic(err)
	}

	fmt.Println(numbers)
}
