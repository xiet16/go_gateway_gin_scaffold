package dao

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/xiet16/go_gateway_gin_scaffold/dto"
	"github.com/xiet16/go_gateway_gin_scaffold/golang_common/lib"
	"github.com/xiet16/go_gateway_gin_scaffold/public"
)

type ServiceDetail struct {
	Info          *ServiceInfo   `json:"info" description:"info"`
	HttpRule      *HttpRule      `json:"http" description:"http"`
	TcpRule       *TcpRule       `json:"tcp" description:"tcp"`
	GrpcRule      *GrpcRule      `json:"grpc" description:"grpc"`
	LoadBalance   *LoadBalance   `json:"loadbalance" description:"loadbalance"`
	AccessControl *AccessControl `json:"accesscontrol" description:"accesscontrol"`
}

//通过handler 对外暴露
var ServiceManagerHandler *ServiceManager

func init() {
	ServiceManagerHandler = NewServiceManager()
}

type ServiceManager struct {
	ServiceMap   map[string]*ServiceDetail
	ServiceSlice []*ServiceDetail
	Locker       sync.RWMutex
	init         sync.Once
	err          error
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap:   map[string]*ServiceDetail{},
		ServiceSlice: []*ServiceDetail{},
		Locker:       sync.RWMutex{},
		init:         sync.Once{},
	}
}

func (s *ServiceManager) GetTcpServiceList() []*ServiceDetail {
	list := []*ServiceDetail{}
	for _, serviceItem := range s.ServiceSlice {
		tempItem := serviceItem
		if tempItem.Info.LoadType == public.LoadTypeTCP {
			list = append(list, tempItem)
		}
	}
	return list
}

func (s *ServiceManager) GetGrpcServiceList() []*ServiceDetail {
	list := []*ServiceDetail{}
	for _, serverItem := range s.ServiceSlice {
		tempItem := serverItem
		if tempItem.Info.LoadType == public.LoadTypeGRPC {
			list = append(list, tempItem)
		}
	}
	return list
}

func (s *ServiceManager) LoadOnce() error {
	s.init.Do(func() {
		//从db中取分页信息
		serviceInfo := &ServiceInfo{}
		c, _ := gin.CreateTestContext(httptest.NewRecorder())
		tx, err := lib.GetGormPool("default")
		if err != nil {
			s.err = err
			return
		}

		params := &dto.ServiceListInput{PageNo: 1, PageSize: 99999}
		list, _, err := serviceInfo.PageList(c, tx, params)
		if err != nil {
			s.err = err
			return
		}

		s.Locker.Lock()
		defer s.Locker.Unlock()
		for _, listItem := range list {
			tmpItem := listItem
			serviceDetail, err := tmpItem.ServiceDetail(c, tx, &listItem)
			if err != nil {
				s.err = err
				return
			}
			s.ServiceMap[tmpItem.ServiceName] = serviceDetail
			s.ServiceSlice = append(s.ServiceSlice, serviceDetail)
		}
	})
	return s.err
}

func (s *ServiceManager) HttpAccessMode(c *gin.Context) (*ServiceDetail, error) {
	//1. 前缀匹配 /abc ==>serviceSlice.rule
	//2. 域名匹配 www.test.com

	host := c.Request.Host //www.test.com:8080
	host = host[0:strings.Index(host, ":")]
	fmt.Println("host: ", host)
	path := c.Request.URL.Path

	for _, serviceItem := range s.ServiceSlice {
		//判断是不是http 服务
		if serviceItem.Info.LoadType != public.LoadTypeHTTP {
			continue
		}
		if serviceItem.HttpRule.RuleType == public.HTTPRuleTypeDomain {
			return serviceItem, nil
		}

		if serviceItem.HttpRule.RuleType == public.HTTPRuleTypePrefixURL {
			if strings.HasPrefix(path, serviceItem.HttpRule.Rule) {
				return serviceItem, nil
			}
		}
	}

	return nil, errors.New("not match any service")
}
