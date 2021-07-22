package data

import (
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	USER userType = iota
	ADMIN
)

type (
	User struct {
		ID       uint   `gorm:"primaryKey,autoIncrement" json:"id"`
		Name     string `gorm:"not null" json:"name" validate:"required"`
		Email    string `gorm:"not null;unique" json:"email" validate:"required,email"`
		Password string `gorm:"not null" json:"password" validate:"required"`
		UserType int32  `gorm:"default:0" json:"user_type" validate:"gte=0,lte=1"`
	}

	UserRepo struct {
		db *gorm.DB
	}
	userType int32
)

func NewUserRepo(connStr string) *UserRepo {

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connStr,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	return &UserRepo{db}

}

func (ur *UserRepo) CreateUser(u *User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return ur.db.Create(u).Error
}

func (ur *UserRepo) GetUser(email string) (User, error) {
	u := User{}
	err := ur.db.Where("email = ?", email).First(&u).Error
	return u, err
}

func (u *User) CheckPassword(provided string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(provided))
}

func (u *User) Validate() error {
	validator := validator.New()
	return validator.Struct(u)
}
