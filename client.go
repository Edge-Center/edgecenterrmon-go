package edgecenterrmon

import (
	"github.com/Edge-Center/edgecenteredgemon-go/channel"
	"github.com/Edge-Center/edgecenteredgemon-go/checkgroup"
	"github.com/Edge-Center/edgecenteredgemon-go/checks/checkdns"
	"github.com/Edge-Center/edgecenteredgemon-go/checks/checkhttp"
	"github.com/Edge-Center/edgecenteredgemon-go/checks/checkping"
	"github.com/Edge-Center/edgecenteredgemon-go/checks/checkrabbitmq"
	"github.com/Edge-Center/edgecenteredgemon-go/checks/checksmtp"
	"github.com/Edge-Center/edgecenteredgemon-go/checks/checktcp"
	"github.com/Edge-Center/edgecenteredgemon-go/edgecenter"
	"github.com/Edge-Center/edgecenteredgemon-go/statuspage"
)

type ClientService interface {
	Channel() channel.Service
	StatusPage() statuspage.Service
	CheckGroup() checkgroup.Service
	CheckDNS() checkdns.Service
	CheckHTTP() checkhttp.Service
	CheckPing() checkping.Service
	CheckRabbitMQ() checkrabbitmq.Service
	CheckSMTP() checksmtp.Service
	CheckTCP() checktcp.Service
}

var _ ClientService = (*Service)(nil)

type Service struct {
	r                    edgecenter.Requester
	channelService       channel.Service
	statusPageService    statuspage.Service
	checkGroupService    checkgroup.Service
	checkDNSService      checkdns.Service
	checkHTTPService     checkhttp.Service
	checkPingService     checkping.Service
	checkRabbitMQService checkrabbitmq.Service
	checkSMTPService     checksmtp.Service
	checkTCPService      checktcp.Service
}

func NewService(r edgecenter.Requester) *Service {
	return &Service{
		r:                    r,
		channelService:       channel.New(r),
		statusPageService:    statuspage.New(r),
		checkGroupService:    checkgroup.New(r),
		checkDNSService:      checkdns.New(r),
		checkHTTPService:     checkhttp.New(r),
		checkPingService:     checkping.New(r),
		checkRabbitMQService: checkrabbitmq.New(r),
		checkSMTPService:     checksmtp.New(r),
		checkTCPService:      checktcp.New(r),
	}
}

func (s *Service) Channel() channel.Service {
	return s.channelService
}

func (s *Service) StatusPage() statuspage.Service {
	return s.statusPageService
}

func (s *Service) CheckGroup() checkgroup.Service {
	return s.checkGroupService
}

func (s *Service) CheckDNS() checkdns.Service {
	return s.checkDNSService
}

func (s *Service) CheckHTTP() checkhttp.Service {
	return s.checkHTTPService
}

func (s *Service) CheckPing() checkping.Service {
	return s.checkPingService
}

func (s *Service) CheckRabbitMQ() checkrabbitmq.Service {
	return s.checkRabbitMQService
}

func (s *Service) CheckSMTP() checksmtp.Service {
	return s.checkSMTPService
}

func (s *Service) CheckTCP() checktcp.Service {
	return s.checkTCPService
}
