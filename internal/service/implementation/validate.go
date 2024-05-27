package certserviceimpl

import (
	"fmt"
	serviceerror "github.com/Dima191/cert-server/internal/service/error"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/status"
)

const (
	domainField = "Domain"
	hostField   = "Host"
	portField   = "Port"
)

func (s *service) validate(domain string) error {
	const hostnameConstraint = "hostname"
	if err := s.v.Var(domain, hostnameConstraint); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			switch err.StructField() {
			case domainField:
				fmt.Println(serviceerror.InvalidDomainError)
				return serviceerror.InvalidDomainError
			case hostField:
				fmt.Println(serviceerror.InvalidHostError)
				return serviceerror.InvalidHostError
			case portField:
				fmt.Println(serviceerror.InvalidPortError)
				return serviceerror.InvalidPortError
			default:
				return status.Error(3, fmt.Sprintf("invalid data: %s", err.Error()))
			}
		}
	}

	return nil
}
