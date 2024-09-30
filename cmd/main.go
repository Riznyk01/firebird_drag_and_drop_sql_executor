package main

import (
	"fmt"
	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"sql_executor/internal/config"
	"sql_executor/internal/repository"
	"time"
)

func main() {
	cfg := config.MustLoad()

	cyan := color.New(color.FgBlue).SprintFunc()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path as an argument.")
		<-time.After(cfg.InfoTimeout)
		os.Exit(1)
	}

	filePath := os.Args[1]
	fmt.Printf("The DB file is: %s\n", filePath)

	logFile, err := os.OpenFile("logfile.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error occurred while opening logfile:\n", err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	fmt.Printf("type the password to DB, please:\n")

	passByte, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		log.Printf("error reading password: %v\n", err)
		<-time.After(cfg.InfoTimeout)
	} else {
		fmt.Printf("password saved\n")
	}

	db, connectionString, err := repository.NewFirebirdDB(cfg, string(passByte), filePath)
	if err != nil {
		fmt.Println(err)
		<-time.After(cfg.InfoTimeout)
		os.Exit(1)
	} else {
		fmt.Printf("db connected, conn.str. is: %s\n", connectionString)
	}

	repo := repository.NewRepository(db)
	fmt.Printf("updating GBackDate...\n")

	err = repo.UpdateDBCorrectionDate(time.Now())
	if err != nil {
		log.Printf("error occurred while updating GBackDate: %v\n", err)
	} else {
		fmt.Printf("GBackDate has been updated\n")
	}

	fmt.Printf("reading the actual GBackDate...\n")

	t, err := repo.GetDBCorrectionDate()
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		<-time.After(cfg.InfoTimeout)
	} else {
		fmt.Printf("GBackDate is: %v\n", cyan(t))
		<-time.After(cfg.InfoTimeout)
	}
}
