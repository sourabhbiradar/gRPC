package main

import (
	pb "FileStream/protoc"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedStreamUploadServer
}

func (s *server) Upload(strm pb.StreamUpload_UploadServer) error {
	var fileBytes []byte
	var fileSize int64 = 0

	for {
		req, err := strm.Recv() // receiving req
		if err == io.EOF {
			break
		}
		chuncks := req.GetChuncks()
		fileBytes = append(fileBytes, chuncks...) // appending chuncks to fileBytes
		fileSize += int64(len(chuncks))           // incrementing fileSize as chuncks append
	}
	f, err := os.Create("./store.bin") // creating static file
	if err != nil {
		log.Println("error creating file", err)
	}
	defer f.Close()

	_, err2 := f.Write(fileBytes) // from fileBytes write to `f`
	if err2 != nil {
		log.Println("error writing to store.bin")
		return err2
	}

	// send fileSize & message to client
	return strm.SendAndClose(&pb.UploadResp{FileSize: fileSize, Message: "file written successfully"})
}

func main() {
	listnr, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	pb.RegisterStreamUploadServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listnr); e != nil {
		panic(e)
	}
}
