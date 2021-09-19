package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/slns/codebanck/domain"
	"github.com/slns/codebanck/infrastructure/repository"
	// "github.com/slns/codebanck/usecase"
)

func main() {
	db := setupDb()

	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "1234"
	cc.Name = "Sergio"
	cc.ExpirationYear = 2021
	cc.ExpirationMonth = 7
	cc.CVV = 123
	cc.Limit = 1000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDb(db)
	err :=repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}
}

// func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
// 	transactionRepository := repository.NewTransactionRepositoryDb(db)
// 	useCase := usecase.NewUseCaseTransaction(transactionRepository)
// //	useCase.KafkaProducer = producer
// 	return useCase
// }

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		// "172.23.0.1",
		"db",
		"5432",
		"postgres",
		"root",
		"codebank",
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}
