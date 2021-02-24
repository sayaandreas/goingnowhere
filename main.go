package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"github.com/sayaandreas/goingnowhere/api"
	"github.com/sayaandreas/goingnowhere/storage"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	store := storage.NewStorageSession()
	// httpHandler := api.NewHandler(store)
	// log.Println("server listening : 3333")
	// http.ListenAndServe(":3333", httpHandler)

	addr := ":3333"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error occurred: %s", err.Error())
	}

	httpHandler := api.NewHandler(store)
	server := &http.Server{
		Handler: httpHandler,
	}
	go func() {
		server.Serve(listener)
	}()
	defer Stop(server)
	log.Printf("Started server on %s", addr)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	log.Println(fmt.Sprint(<-ch))
	log.Println("Stopping API server.")
	// store.GetBucketObjectList()

	// resp := store.GetBucketList()
	// for _, bucket := range resp.Buckets {
	// 	fmt.Println("Bucket name:", *bucket.Name)
	// }
}
func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}

// func multipleFiles() {
// 	fmt.Println("Download Started")
// 	urls := []string{"https://upload.wikimedia.org/wikipedia/commons/thumb/0/0a/Ambarawa_Bypass_Road_from_Eling_Bening%2C_2017-03-15.jpg/3240px-Ambarawa_Bypass_Road_from_Eling_Bening%2C_2017-03-15.jpg", "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"}
// 	bytes, err := download.DownloadMultipleFilesIntoBytes(urls)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for i := 0; i < len(urls); i++ {
// 		output, err := os.Create("files" + strconv.Itoa(i) + ".jpg")
// 		if err != nil {
// 			panic(err)
// 		}
// 		size, err := output.Write(bytes[i])
// 		fmt.Printf("wrote %d bytes\n", size)
// 	}
// }

// func withProgress() {
// 	fmt.Println("Download Started")

// 	fileURL := "https://upload.wikimedia.org/wikipedia/commons/d/d6/Wp-w4-big.jpg"
// 	err := download.DownloadFileWithProgress("avatar.jpg", fileURL)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Download Finished")
// }
