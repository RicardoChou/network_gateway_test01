package grpc_proxy_middleware

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhj/go_gateway/dao"
	"github.com/zhj/go_gateway/public"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"log"
	"strings"
)

// GrpcWhiteListMiddleware 服务黑名单列表
func GrpcWhiteListMiddleware(serviceDetail *dao.ServiceDetail) func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error{
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error{
		iplist := []string{}
		if serviceDetail.AccessControl.WhiteList!=""{
			iplist = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}

		peerCtx,ok:=peer.FromContext(ss.Context())
		if !ok{
			return errors.New("peer not found with context")
		}
		peerAddr:=peerCtx.Addr.String()
		addrPos:=strings.LastIndex(peerAddr,":")
		clientIP:=peerAddr[0:addrPos]
		if serviceDetail.AccessControl.OpenAuth == 1 && len(iplist) > 0 {
			if !public.InStringSlice(iplist, clientIP) {
				return errors.New(fmt.Sprintf("%s not in white ip list", clientIP))
			}
		}
		if err := handler(srv, ss);err != nil {
			log.Printf("RPC failed with error %v\n", err)
			return err
		}
		return nil
	}
}
