package service

import (
	"context"
	"github.com/Allen-Career-Institute/go-kratos-sample/internal/biz"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	pb "github.com/Allen-Career-Institute/go-kratos-sample/api/user/v1"
)

// UserService Code to handle the transport layer logic and invoke the
// correct set of handlers in response post validation.
type UserService struct {
	pb.UnimplementedUserServer
	handler *biz.UserHandler
	log     *log.Helper
}

func NewUserService(handler *biz.UserHandler, logger log.Logger) *UserService {
	return &UserService{
		UnimplementedUserServer: pb.UnimplementedUserServer{},
		handler:                 handler,
		log:                     log.NewHelper(logger),
	}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	if err := validateUser(req); err != nil {
		return nil, err
	}
	userInfo := &pb.UserInfo{
		MobileNumber: req.MobileNumber,
		GivenName:    req.GivenName,
		FamilyName:   req.FamilyName,
	}

	result, err := s.handler.CreateUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserReply{UserInfo: result}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	userInfo := &pb.UserInfo{
		Name:       req.Name,
		GivenName:  req.GivenName,
		FamilyName: req.FamilyName,
	}

	result, err := s.handler.UpdateUser(ctx, userInfo)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateUserReply{UserInfo: result}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	err := s.handler.DeleteUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	u, err := s.handler.GetUser(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{UserInfo: u}, nil
}

// ListUser No Lint is just used for sample code. We should not replicate it
// nolint: revive
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	return &pb.ListUserReply{}, nil
}

func validateUser(req *pb.CreateUserRequest) error {
	if req == nil {
		return errors.BadRequest("Create Request Not Present", "Create Request Validation Failed")
	}
	if req.GivenName == "" {
		return errors.BadRequest("Given Name is not present", "Create Request Validation Failed")
	}
	if req.FamilyName == "" {
		return errors.BadRequest("Family Name is not present", "Create Request Validation Failed")
	}
	if req.MobileNumber == "" {
		return errors.BadRequest("Mobile No is not present", "Create Request Validation Failed")
	}
	return nil
}
