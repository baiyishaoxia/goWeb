package grpc

type GrpcInfo struct {
}
type GrpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
