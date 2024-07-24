package main

import (
	"context"
	"copilot/internal/router"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Opções de conexão
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectar ao MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Verificar a conexão
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conectado ao MongoDB!")

	// Lembre-se de desconectar quando terminar
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	r := router.New(client)
	r.Run() // por padrão, roda na porta 8080
}
