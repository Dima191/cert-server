package certserviceimpl

import (
	"context"
	"github.com/Dima191/cert-server/internal/models"
	certrepository "github.com/Dima191/cert-server/internal/repository"
	"go.uber.org/mock/gomock"
	"log/slog"
	"testing"
)

const (
	certificateChain = "-----BEGIN CERTIFICATE-----\nMIIDDzCCAfegAwIBAgIUHbWknIITlI5qsnqRdOBb8Xj0lf8wDQYJKoZIhvcNAQEL\nBQAwFDESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTI0MDEyODEwMTY0M1oXDTI0MDIy\nNzEwMTY0M1owFDESMBAGA1UEAwwJbG9jYWxob3N0MIIBIjANBgkqhkiG9w0BAQEF\nAAOCAQ8AMIIBCgKCAQEA4uxBEhaIGR+SKpiz2WW+ZGBop/WppdL4jpbCnz3Y5ODo\nkoPMvl3s+4k7CVNHeh0wFchjSFW0mUcOhqHc/ES53bvzskp/TTTidmuCtjiGQ1GA\n/9BKRV6lYzn0r6Em950Yc8G6o+RUqwrGRjlKBNtTg+NrkoY+PgjKX5SOzixwzkgb\nVSe2cESCLIvynsmGvHvJCgMoy/tAPUm/eFIhaxYUjj1rIV1tdCkP5OAmxAE8y/Co\nIzXYc2Afz2bW0qOdpaj6MIoKbhUCN+gY4oW132TyMC1NrlBjNYGdC4aLzBHxpu3R\noU1FNmKVs6hGT7Aq0UAWYIEQj6HDgdbm7lGKoFo/LQIDAQABo1kwVzAUBgNVHREE\nDTALgglsb2NhbGhvc3QwCwYDVR0PBAQDAgeAMBMGA1UdJQQMMAoGCCsGAQUFBwMB\nMB0GA1UdDgQWBBS/LeKyW67p+SWmdDdz9Xpc6A3jhjANBgkqhkiG9w0BAQsFAAOC\nAQEAv7QGIGSNuxXPqw8uXAM6RXCUErBNYc/i9A6qiY5kFAI6itPDonFJOlKr9+jA\nSbrRdkLg/MrVwNsaT2Vq397cdgkwpvLR8+RGxY0lqhSXBwnUK1sBhKbNk+4F96Gt\nlmaWulNDLrT6NP/nMtwnsbW/evy+Lx73hg0INHsrvHXYGT7CUdOZS2Dq0zw7ws4x\n34LvlU/f7JvHOFDS2YSNHwm8goQKigvsI0jSZg8MH+kYGQFkZAGnYyGM/jGp0dul\nj+lMqprDwh2WtcSGcLXFKLADw46U8aQcbnXH0A6Xr7XmxEc6MYZxanZsw6GlkeNr\nrokhPQq/VOIELIf0mRAlDfClRw==\n-----END CERTIFICATE-----\n"
	privateKey       = "-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDi7EESFogZH5Iq\nmLPZZb5kYGin9aml0viOlsKfPdjk4OiSg8y+Xez7iTsJU0d6HTAVyGNIVbSZRw6G\nodz8RLndu/OySn9NNOJ2a4K2OIZDUYD/0EpFXqVjOfSvoSb3nRhzwbqj5FSrCsZG\nOUoE21OD42uShj4+CMpflI7OLHDOSBtVJ7ZwRIIsi/KeyYa8e8kKAyjL+0A9Sb94\nUiFrFhSOPWshXW10KQ/k4CbEATzL8KgjNdhzYB/PZtbSo52lqPowigpuFQI36Bji\nhbXfZPIwLU2uUGM1gZ0LhovMEfGm7dGhTUU2YpWzqEZPsCrRQBZggRCPocOB1ubu\nUYqgWj8tAgMBAAECggEADw6HCMi28hfOroXgVRnzppxhBVf/EDIt7OQi+Mb1R8aG\nmAYnxS/dRBloced1NCyQnrljoQ6Sw6Lb22INFC4JOSW590gLJ1fNePcMPsQEYJ4d\nVZv/+ZdcmPdk/WxPdhmV4ERn+mzxk0HNQyaU+zqEUZl670d6BMq3ht0IFrULLRyq\nu2bs7SeMdOc9tTeE5A2V3TTpfKzSEPCKCpku6n3N7XXyjiBTEwFUoTaHWFXtLbcE\nLYG9W3bcgNkQh5ePBdMXcnyqN1USmBFom76srSP3ac6DP9Eb/u/mH1WhxPd14bKd\ns1Vjw1zBFAFv5uOozVFgYpoakkSH5vLGgtGoOcwKzwKBgQDxYKQ/zXzz4Emo6Rma\nJs5UAG114L2eY2Gfl5iz5O7wbFQtQ9RHE5NlHXKvAyIXDx8b+UjV/qKY5EoW0wSg\nw2bvYQDKGC4QLo/D8l1GrVyormPUFlnwxltAhreDM8qkv5IXySsuyPtivIffjBVW\nMvYXyZwnNJL7j2euh6EC3hen5wKBgQDwq3IMONfEqv/+JsPJZ7Noc4k8frk/Mp5y\n4K6BpchzvV3zxaQo6zOjn6aoTOjUsYiGRnupZt0GXI1voWwlZNOsOuFtjG3SrfKm\nwnUeSc7L61jK9pOVm1aw3fQlQlfLDgHWC+AWcSPxTxuye3R3A9SurGmCC3zcDXmk\nkiBaV7etywKBgQCQpTIoTdKoLmrVvsIp30Fbk1oE/qWCydlRkr3eZ71L9A7JhVEr\nOq7kNC5qdD00hkpFMDCWlF4Jsxw260Nlt1Ly9jVL5guMhOqAKLf+x5q0NrT5/l7t\na2B5nYFRLXMtIOPCPzoScjw64fGmY8LRgf55KMbbs6S0/S3Lp9kz57VTVwKBgAHK\nxQaiRbkJLO7PixWs2AEnGxaAOxDlZ5ijY6lDesKh3lk6V4aWecP2JF+Mcw9iYwnc\n7H7ObUbm2YbDRPLiVVEq/xK5wPeYo/3p5MVc91U9Y2PfginTBko63N6KJ+0zJYBa\nhkikfwgE9sfVc4CPXr4OOprlSwC4ePdovyTtEkkRAoGBAJeG8DgspkpFoGcD7wKr\nFL4NGRT6O/7uQ6KbFvowHTby8dWIU6bNLQWJDzUDZ0x7LUwmXTDsljz7hg046huF\nE/ex7Yy00/TU48qy7/JvgXDZIJHFXldnjs+pMmWc3M2eZiHD3dg7890qog/STB0x\nlXckBnuTisdpCR72NOzYW0Zy\n-----END PRIVATE KEY-----\n"
)

const domain = "docs.test.ru"

func BenchmarkServiceCert(b *testing.B) {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	mockRep := certrepository.NewMockRepository(ctrl)

	logger := slog.Default()

	s := New(mockRep, logger)

	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		b.StopTimer()

		mockRep.EXPECT().Cert(gomock.Any(), domain).Return(&models.Cert{
			PrivateKey:       privateKey,
			CertificateChain: certificateChain,
		}, nil)
		b.StartTimer()

		if _, err := s.Cert(context.Background(), domain); err != nil {
			b.Fatal(err)
		}
	}

}
