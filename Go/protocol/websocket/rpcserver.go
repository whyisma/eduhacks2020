package websocket

import (
	"context"
	"eduhacks2020/Go/pkg/setting"
	"eduhacks2020/Go/protocol/pb"
	"eduhacks2020/Go/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

// CommonServiceServer 基础服务
type CommonServiceServer struct{}

// Send2Client 发送到客户端
func (s *CommonServiceServer) Send2Client(ctx context.Context, req *pb.Send2ClientReq) (*pb.Send2ClientReply, error) {
	log.WithFields(log.Fields{
		"host":     setting.GlobalSetting.LocalHost,
		"port":     setting.CommonSetting.Port,
		"clientId": req.ClientID,
	}).Info("received RPC specified client message")
	SendMessage2LocalClient(req.MessageID, req.ClientID, req.SendUserID, int(req.Code), req.Message, &req.Data)
	return &pb.Send2ClientReply{}, nil
}

// CloseClient 关闭客户端
func (s *CommonServiceServer) CloseClient(ctx context.Context, req *pb.CloseClientReq) (*pb.CloseClientReply, error) {
	log.WithFields(log.Fields{
		"host":     setting.GlobalSetting.LocalHost,
		"port":     setting.CommonSetting.Port,
		"clientId": req.ClientID,
	}).Info("received RPC to close the connection")
	CloseLocalClient(req.ClientID, req.SystemID)
	return &pb.CloseClientReply{}, nil
}

// BindGroup 添加分组到group
func (s *CommonServiceServer) BindGroup(ctx context.Context, req *pb.BindGroupReq) (*pb.BindGroupReply, error) {
	if client, err := Manager.GetByClientID(req.ClientID); err == nil {
		//添加到本地
		Manager.AddClient2LocalGroup(req.GroupName, client, req.UserID, req.Extend)
	} else {
		log.Error("failed to add group" + err.Error())
	}
	return &pb.BindGroupReply{}, nil
}

// Send2Group 发送到群组
func (s *CommonServiceServer) Send2Group(ctx context.Context, req *pb.Send2GroupReq) (*pb.Send2GroupReply, error) {
	log.WithFields(log.Fields{
		"host": setting.GlobalSetting.LocalHost,
		"port": setting.CommonSetting.Port,
	}).Info("received RPC sent packet message")
	Manager.SendMessage2LocalGroup(req.SystemID, req.MessageID, req.SendUserID, req.GroupName, int(req.Code), req.Message, &req.Data)
	return &pb.Send2GroupReply{}, nil
}

// Send2System 发送到系统
func (s *CommonServiceServer) Send2System(ctx context.Context, req *pb.Send2SystemReq) (*pb.Send2SystemReply, error) {
	log.WithFields(log.Fields{
		"host": setting.GlobalSetting.LocalHost,
		"port": setting.CommonSetting.Port,
	}).Info("received RPC sending system message")
	Manager.SendMessage2LocalSystem(req.SystemID, req.MessageID, req.SendUserID, int(req.Code), req.Message, &req.Data)
	return &pb.Send2SystemReply{}, nil
}

// GetGroupClients 获取分组在线用户列表
func (s *CommonServiceServer) GetGroupClients(ctx context.Context, req *pb.GetGroupClientsReq) (*pb.GetGroupClientsReply, error) {
	response := pb.GetGroupClientsReply{}
	response.List = Manager.GetGroupClientList(utils.GenGroupKey(req.SystemID, req.GroupName))
	return &response, nil
}

// InitGRPCServer 初始化 RPC
func InitGRPCServer() {
	go createGRPCServer(":" + setting.CommonSetting.RPCPort)
}

func createGRPCServer(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pb.RegisterCommonServiceServer(s, &CommonServiceServer{})

	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
}
