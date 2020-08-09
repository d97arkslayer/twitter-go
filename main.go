package main


import (
	"github.com/d97arkslayer/twitter-go/Database"
	"github.com/d97arkslayer/twitter-go/Handlers"
	"log"
)

func main() {
	if Database.CheckConnection()==0{
		log.Fatal("No database connection")
		return
	}
	Handlers.Handlers()
}