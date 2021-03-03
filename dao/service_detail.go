package dao

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"info"`
	HttpRule      *HttpRule      `json:"http" description:"http"`
	TcpRule       *TcpRule       `json:"tcp" description:"tcp"`
	GrpcRule      *GrpcRule      `json:"grpc" description:"grpc"`
	LoadBalance   *LoadBalance   `json:"loadbalance" description:"loadbalance"`
	AccessControl *AccessControl `json:"accesscontrol" description:"accesscontrol"`
}
