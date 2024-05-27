package certserviceimpl

import (
	"context"
	"github.com/Dima191/cert-server/internal/models"
	certrepository "github.com/Dima191/cert-server/internal/repository"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"log/slog"
	"os"
	"testing"
)

func TestCert(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockedRep := certrepository.NewMockRepository(ctrl)

	mockedRep.EXPECT().Cert(gomock.Any(), domain).Return(&models.Cert{
		PrivateKey:       certificateChain,
		CertificateChain: privateKey,
	}, nil)

	// LOGGER
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}))

	s := New(mockedRep, logger)
	_, err := s.Cert(context.Background(), domain)
	assert.NoError(t, err)
}
