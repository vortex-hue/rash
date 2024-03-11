package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/go-redis/redis/v8"
	"net/http"
	//"sync"
)

var (
	producer sarama.SyncProducer
	redisClient *redis.Client
)

func initKafka() {
	var err error
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	brokers := []string{"localhost:9092"}
	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic(err)
	}
}

func initRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func postDataToKafka(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Here, you would parse your request data (omitted for brevity)

	go func(data string) {
		// Producing to Kafka
		msg := &sarama.ProducerMessage{
			Topic: "yourTopic",
			Value: sarama.StringEncoder(data),
		}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			fmt.Printf("Failed to send to Kafka: %s\n", err)
		}
	}("your-data")

	fmt.Fprintln(w, "Data sent to Kafka")
}

func getDataFromRedis(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Retrieve data from Redis
	val, err := redisClient.Get(context.Background(), "yourKey").Result()
	if err != nil {
		fmt.Printf("Error getting data from Redis: %s\n", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Data from Redis: %s", val)
}

func main() {
	initKafka()
	initRedis()

	http.HandleFunc("/push", postDataToKafka)
	http.HandleFunc("/get", getDataFromRedis)

	fmt.Println("Server starting on port :8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Printf("Failed to start server: %s\n", err)
	}
}
