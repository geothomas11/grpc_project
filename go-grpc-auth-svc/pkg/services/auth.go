package services

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/geothomas11/go-grpc-auth-svc/pkg/db"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/models"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/pb"
	"github.com/geothomas11/go-grpc-auth-svc/pkg/utils"
	"gorm.io/gorm"
)

type Server struct {
	H   db.Handler
	JWT utils.JwtWrapper
	pb.UnimplementedAuthServiceServer
}

// func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
// 	var user models.User

// 	if result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user); result.Error == nil {
// 		return &pb.RegisterResponse{
// 			Status: http.StatusConflict,
// 			Error:  "E-mail already exists",
// 		}, nil
// 	}
// 	user.Email = req.Email
// 	user.Password = utils.HashPassword(req.Password)

// 	s.H.DB.Create(&user)
// 	//user return print

// 	var users []models.User
// 	s.H.DB.Find(&users)
// 	log.Println("Users:", users)

//		return &pb.RegisterResponse{
//			Status: http.StatusCreated,
//		}, nil
//	}
func (s *Server) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	// Check if the user already exists
	result := s.H.DB.Where(&models.User{Email: req.Email}).First(&user)
	if result.Error == nil {
		// User exists
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  "E-mail already exists",
		}, nil
	} else if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// Some other DB error
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  result.Error.Error(),
		}, nil
	}

	// User does not exist, create a new one
	user = models.User{
		Email:    req.Email,
		Password: utils.HashPassword(req.Password),
		Role:     "user", // Optional: assign default role
	}

	if err := s.H.DB.Create(&user).Error; err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Log all users (optional)
	var users []models.User
	s.H.DB.Find(&users)
	log.Println("Users:", users)

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if result := s.H.DB.Where("email = ?", req.Email).First(&user); result.Error != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid password",
		}, nil
	}

	token, err := s.JWT.GenerateToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to generate token",
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *Server) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	claims, err := s.JWT.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}
	var user models.User
	if result := s.H.DB.Where(&models.User{Email: claims.Email}).First(&user); result.Error != nil {
		return &pb.ValidateResponse{
			Status: http.StatusNotFound,
			Error:  "User not found",
		}, nil
	}
	return &pb.ValidateResponse{
		Status: http.StatusOK,
		UserId: user.Id,
	}, nil

}
