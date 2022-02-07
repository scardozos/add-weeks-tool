package dbclient

import (
	"crypto/tls"
	"crypto/x509"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// CreateConn initiaes a new grpc ClientConnection taking in the hostname of the gRPC server
// and whether the connection should be encrypted (secure) or not
func (s *DbClientContext) CreateConn(grpcServerAddr string, secure bool) error {

	var opts []grpc.DialOption

	if grpcServerAddr != "" {
		opts = append(opts, grpc.WithAuthority(grpcServerAddr))
	}

	if secure {
		systemRoots, err := x509.SystemCertPool()
		if err != nil {
			return err
		}
		cred := credentials.NewTLS(&tls.Config{
			RootCAs: systemRoots,
		})
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	conn, err := grpc.Dial(grpcServerAddr, opts...)
	if err != nil {
		log.Printf("failed to connect: %v", err)
		return err
	}
	s.Conn = conn
	return nil
}
