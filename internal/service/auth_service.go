package service

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"go-ecommerce/internal/model"
	"go-ecommerce/internal/repository"
)

type AuthService interface {
	Register(email, password, fullName string) (*model.User, error)
	Login(email, password string) (string, error) // Trả về Token
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo: repo}
}

func (s *authService) Register(email, password, fullName string) (*model.User, error) {
	// 1. Kiểm tra email đã tồn tại chưa
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("email đã được sử dụng")
	}

	// 2. Băm mật khẩu (Hash password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 3. Tạo User mới (Mặc định role là customer)
	user := &model.User{
		Email:    email,
		Password: string(hashedPassword),
		FullName: fullName,
		Role:     "customer",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authService) Login(email, password string) (string, error) {
	// 1. Tìm user theo email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("email hoặc mật khẩu không đúng")
	}

	// 2. So sánh mật khẩu người dùng nhập với mật khẩu đã băm trong DB
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("email hoặc mật khẩu không đúng")
	}

	// 3. Tạo JWT Token
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Hết hạn sau 24h
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
