# Cert Server

Cert Server provides TLS certificates for secure data exchange between microservices.

### Functionality

- Manages generation and issuance of TLS certificates.

### Interaction with Other Microservices

- **xds**: Provides certificates for establishing secure gRPC connections.
- **route-server**: Can use domain and endpoint data for configuring certificate security.

### Technologies

- Interacts with a database for certificate storage.
