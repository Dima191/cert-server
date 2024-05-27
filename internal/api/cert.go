package api

import (
	"context"
	pb "github.com/Dima191/cert-server/pkg/cert"
	"unsafe"
)

func (i *Implementation) Cert(ctx context.Context, req *pb.CertRequest) (*pb.CertResponse, error) {
	c, err := i.service.Cert(ctx, req.GetDomain())
	if err != nil {
		return nil, err
	}

	return &pb.CertResponse{
		PrivateKey:       unsafe.Slice(unsafe.StringData(c.PrivateKey), len(c.PrivateKey)),
		CertificateChain: unsafe.Slice(unsafe.StringData(c.CertificateChain), len(c.CertificateChain)),
	}, nil
}
