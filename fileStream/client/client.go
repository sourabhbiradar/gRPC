package main

import (
	pb "FileStream/protoc"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client pb.StreamUploadClient

func main() {
	conn, err := grpc.NewClient(":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	pb.NewStreamUploadClient(conn)

	mb := 1024 * 1024 * 2                 // 4mb chuncks will be sent at once
	uploadStreamFile("./1GBfile.bin", mb) // implemented below
}

func uploadStreamFile(path string, batchSize int) {
	t := time.Now() // to see how long it took to upload file
	file, err := os.Open(path)
	if err != nil {
		log.Println("error opening file")
	}

	buf := make([]byte, batchSize) // setting buffer size
	batchNumber := 1               // how batches it took to send file

	strm, err := client.Upload(context.TODO())
	if err != nil {
		panic(err)
	}

	for {
		num, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		chunk := buf[:num] // dividing file into chuncks

		if e := strm.Send(&pb.UploadReqt{FilePath: path, Chuncks: chunk}); e != nil {
			fmt.Println(e)
			return
		}
		log.Printf("sent - batch %v - size %v\n", batchNumber, len(chunk))
		batchNumber += 1
	}

	res, err := strm.CloseAndRecv()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(time.Since(t))
	log.Printf("sent - %v , bytes - %v\n", res.GetFileSize(), res.GetMessage())

}
