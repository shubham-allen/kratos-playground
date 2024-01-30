package user

import (
	"github.com/go-kratos/kratos/v2/errors"

	pb "github.com/Allen-Career-Institute/go-kratos-sample/api/user/v1"
)

type args struct {
	CreateUser *pb.CreateUserRequest
}

type testData struct {
	Name          string
	Args          args
	ExpectedError error
	Mock          func()
}

var request = map[string]args{
	"Create Request Not Present": {
		CreateUser: nil,
	},
	"Given Name empty case": {
		CreateUser: &pb.CreateUserRequest{
			GivenName:    "",
			FamilyName:   "Alex",
			MobileNumber: "8888833333",
		},
	},
	"Family Name empty case": {
		CreateUser: &pb.CreateUserRequest{
			GivenName:    "Root",
			FamilyName:   "",
			MobileNumber: "8888833333",
		},
	},
	"Mobile Number empty case": {
		CreateUser: &pb.CreateUserRequest{
			GivenName:    "Root",
			FamilyName:   "Alex",
			MobileNumber: "",
		},
	},
	"User success case": {
		CreateUser: &pb.CreateUserRequest{
			GivenName:    "Root",
			FamilyName:   "Alex",
			MobileNumber: "8888833333",
		},
	},
}

var errorResponse = map[string]error{
	"Create Request Not Present": errors.BadRequest("Create Request Not Present", "Create Request Validation Failed"),
	"Given Name empty case":      errors.BadRequest("Given Name is not present", "Create Request Validation Failed"),
	"Family Name empty case":     errors.BadRequest("Family Name is not present", "Create Request Validation Failed"),
	"Mobile Number empty case":   errors.BadRequest("Mobile No is not present", "Create Request Validation Failed"),
	"User success case":          nil,
}

var TestCases = []testData{
	{
		Name:          "Create Request Not Present",
		Args:          request["Create Request Not Present"],
		ExpectedError: errorResponse["Create Request Not Present"],
		Mock: func() {

		},
	}, {
		Name:          "Given Name empty case",
		Args:          request["Given Name empty case"],
		ExpectedError: errorResponse["Given Name empty case"],
		Mock: func() {

		},
	}, {
		Name:          "Family Name empty case",
		Args:          request["Family Name empty case"],
		ExpectedError: errorResponse["Family Name empty case"],
		Mock: func() {

		},
	}, {
		Name:          "Mobile Number empty case",
		Args:          request["Mobile Number empty case"],
		ExpectedError: errorResponse["Mobile Number empty case"],
		Mock: func() {

		},
	}, {
		Name:          "User success case",
		Args:          request["User success case"],
		ExpectedError: errorResponse["User success case"],
		Mock: func() {

		},
	},
}
