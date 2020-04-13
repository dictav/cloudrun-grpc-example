package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"log"
	"os"

	"github.com/dictav/cloudrun-grpc-example/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var (
	addr  = flag.String("addr", "localhost:8080", "host")
	insecure = flag.Bool("insecure", false, "use insecure request")
	gType = flag.String("type", "Hello", "GreetType")
)

func main() {
	flag.Parse()

	name := flag.Arg(0)
	if name == "" {
		log.Fatal("name is required")
	}

	tp, ok := proto.GreetType_value[*gType]
	if !ok {
		log.Fatal("invalid type:", *gType)
	}

	opts := []grpc.DialOption{}

	if *insecure {
		opts = append(opts, grpc.WithInsecure())
	} else {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			log.Fatal("failed to load system root CA cert pool")
		}
		creds := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	log.Printf("Connecting to gRPC Service [%s]", *addr)

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	token := os.Getenv("TOKEN")
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer " + token)
	req := proto.GreetRequest{
		Type: proto.GreetType(tp),
		Name: name,
	}

	res, err := client.Greet(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("greeting:", res.GetGreeting())
}
