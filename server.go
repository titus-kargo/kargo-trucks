package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/titus-kargo/kargo-trucks/graph"
	"github.com/titus-kargo/kargo-trucks/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// in server.go
	// go to line 20

	resolver := &graph.Resolver{}
	resolver.Init()
	if len(resolver.Trucks) >= 10 {
		fmt.Println("Siap untuk mengirim email")
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	// // Sender data.
	// from := os.Getenv("username")
	// password := os.Getenv("password")

	// // Receiver email address.
	// to := []string{
	// 	os.Getenv("receiver1"),
	// }

	// // SMTP server configuration.
	// smtpHost := "smtp.gmail.com"
	// smtpPort := "587"

	// // Message.
	// message := []byte("this is a test email message.")

	// // Authentication
	// auth := smtp.PlainAuth("", from, password, smtpHost)

	// // Sending Email.
	// err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println("Email Sent Successfully!")

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initEnv() error {
	if err := godotenv.Load(); err != nil {
		return err
	}

	return nil
}

func records() {
	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "gardener"},
		{"Lucy", "Smith", "teacher"},
		{"Brian", "Bethamy", "programmer"},
	}
	f, err := os.Create("users.csv")
	defer f.Close()
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	w := csv.NewWriter(f)
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
