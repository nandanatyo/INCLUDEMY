package midtrans

// import (
//     "os"

//     midtransgo "github.com/midtrans/midtrans-go"
// )

// type Interface interface {
//     CreatePaymentLink(orderID string, grossAmount int64) (string, error)
// }

// type midtransClient struct {
//     serverKey string
//     client    *midtransgo.Bank
// }

// func Init() Interface {
// 	client := &midtransgo.ConfigOptions{
// 		ServerKey: os.Getenv("MIDTRANS_SERVER_KEY"),
// 		ClientKey: os.Getenv("MIDTRANS_CLIENT_KEY"),
// 		Env:       midtransgo.Sandbox,
// 	}

// 	return client
// }
