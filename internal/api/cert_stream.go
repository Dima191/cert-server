package api

import (
	"errors"
	"fmt"
	pb "github.com/Dima191/cert-server/pkg/cert"
	"io"
)

func (i *Implementation) CertStream(stream pb.Cert_CertStreamServer) error {
	for {
		req, err := stream.Recv()
		fmt.Println(req, err)
		if errors.Is(err, io.EOF) {
			return nil
		}
		if err != nil {
			return err
		}
		ctx := stream.Context()
		res, err := i.Cert(ctx, req)
		if err != nil {
			return err
		}

		if err = stream.Send(res); err != nil {
			return err
		}
	}
}
