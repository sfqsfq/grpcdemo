package main

import (
	"context"
	pb "github.com/sfqsfq/ch5/deadline/order-service-gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io"
	"log"
)

const (
	address    = "localhost:50051"
	certFile   = "server.crt"
	serverName = "sfq.me"
)

func main() {
	creds, err := credentials.NewClientTLSFromFile(certFile, serverName)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
	}
	// Setting up a connection to the server.
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second * 5)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Add Order
	order1 := pb.Order{Id: "101", Items: []string{"iPhone XS", "Mac Book Pro"}, Destination: "San Jose, CA", Price: 2300.00}
	res, addErr := client.AddOrder(ctx, &order1)

	if addErr != nil {
		got := status.Code(addErr)
		log.Printf("Error Occured -> addOrder : , %v:", got)
	} else {
		log.Print("AddOrder Response -> ", res.Value)
	}

	//log.Println("Cancelling context... ")
	//cancel()

	// Following RPC calls should fail because the RPC context is already cancelled.

	// Get Order
	//retrievedOrder , getOrderErr := client.GetOrder(ctx, &wrapper.StringValue{Value: "106"})
	//
	//if getOrderErr != nil {
	//	log.Printf("Error Occured -> getOrder : , %v:", status.Code(getOrderErr))
	//} else {
	//	log.Print("GetOrder Response -> : ", retrievedOrder)
	//}

	// Search Order
	//searchStream, _ := client.SearchOrders(ctx, &wrapper.StringValue{Value: "Google"})
	//for {
	//	searchOrder, err := searchStream.Recv()
	//	if err == io.EOF {
	//		log.Print("EOF")
	//		break
	//	}
	//
	//	if err == nil {
	//		log.Print("Search Result : ", searchOrder)
	//	}
	//}

	// Update Orders

	//updOrder1 := pb.Order{Id: "102", Items:[]string{"Google Pixel 3A", "Google Pixel Book"}, Destination:"Mountain View, CA", Price:1100.00}
	//updOrder2 := pb.Order{Id: "103", Items:[]string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination:"San Jose, CA", Price:2800.00}
	//updOrder3 := pb.Order{Id: "104", Items:[]string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination:"Mountain View, CA", Price:2200.00}
	//
	//updateStream, _ := client.UpdateOrders(ctx)
	//_ = updateStream.Send(&updOrder1)
	//_ = updateStream.Send(&updOrder2)
	//_ = updateStream.Send(&updOrder3)
	//
	//
	//updateRes, _ := updateStream.CloseAndRecv()
	//log.Printf("Update Orders Res : ", updateRes)

	// Process Order
	//streamProcOrder, _ := client.ProcessOrders(ctx)
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value: "102"})
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value: "103"})
	//_ = streamProcOrder.Send(&wrapper.StringValue{Value: "104"})
	//
	//channel := make(chan bool, 1)
	//
	//go asncClientBidirectionalRPC(streamProcOrder, channel)
	//time.Sleep(time.Millisecond * 1000)
	//
	//// Cancelling the RPC
	//cancel()
	//log.Printf("RPC Status : %s", ctx.Err())
	//
	//err = streamProcOrder.Send(&wrapper.StringValue{Value: "101"})
	//x, ok := status.FromError(err)
	//
	//fmt.Println(x, ok)
	//_ = streamProcOrder.CloseSend()
	//<-channel
	//status.New(codes.OK, "").WithDetails()
}

func asncClientBidirectionalRPC(streamProcOrder pb.OrderManagement_ProcessOrdersClient, c chan bool) {
	for {
		combinedShipment, errProcOrder := streamProcOrder.Recv()
		if errProcOrder != nil {
			log.Printf("Error Receiving messages %v", errProcOrder)
			break
		} else {
			if errProcOrder == io.EOF {
				break
			}
			log.Printf("Combined shipment : %s", combinedShipment.OrdersList)
		}
	}
	c <- true
}
