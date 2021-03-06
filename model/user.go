package model

import (
    "fmt"
    "crypto/md5"
    "strconv"
    "time"
    "golang123/config"
)

// User 用户
type User struct {
    ID             uint               `gorm:"primary_key" json:"id"`
    CreatedAt      time.Time          `json:"createdAt"`
    UpdatedAt      time.Time          `json:"updatedAt"`
    DeletedAt      *time.Time         `sql:"index" json:"deletedAt"`
    Name           string             `json:"name"`
    Pass           string             `json:"-"`
    Email          string             `json:"email"`
    Phone          string             `json:"phone"`
    Score          uint               `json:"score"`
    ArticleCount   uint               `json:"articleCount"`
    CommentCount   uint               `json:"commentCount"`
    Collects       []Collect          `json:"collects"`
    CollectCount   uint               `json:"collectCount"`
    Signature      string             `json:"signature"`
    Role           int                `json:"role"`
    AvatarURL      string             `json:"avatarURL"`
    Status         int                `json:"status"`    
}

// CheckPassword 验证密码是否正确
func (user User) CheckPassword(password string) bool {
    if password == "" || user.Pass == "" {
        return false
    }
    return user.EncryptPassword(password, user.Salt()) == user.Pass
}

// Salt 每个用户都有一个不同的盐
func (user User) Salt() string {
    var userSalt string
    if user.Pass == "" {
        userSalt = strconv.Itoa(int(time.Now().Unix()))
    } else {
        userSalt = user.Pass[0:10]   
    }
    return userSalt
}

// EncryptPassword 给密码加密
func (user User) EncryptPassword(password, salt string) (hash string) {
    password = fmt.Sprintf("%x", md5.Sum([]byte(password)))
    hash = salt + password + config.ServerConfig.PassSalt
    hash = salt + fmt.Sprintf("%x", md5.Sum([]byte(hash)))
    return
}

// PublicInfo 用户公开的信息
func (user User) PublicInfo() User {
    return User{
        ID             : user.ID,
        CreatedAt      : user.CreatedAt,
        Name           : user.Name,
        Score          : user.Score,
        ArticleCount   : user.ArticleCount,
        CommentCount   : user.CommentCount,
        CollectCount   : user.CollectCount,
        Signature      : user.Signature,
        AvatarURL      : user.AvatarURL,
    }
}

const (
    // UserRoleNormal 普通用户
    UserRoleNormal      = 1

    // UserRoleEditor 网站编辑
    UserRoleEditor      = 2

    // UserRoleAdmin 管理员
    UserRoleAdmin       = 3

    // UserRoleSuperAdmin 超级管理员
    UserRoleSuperAdmin  = 4
)

const (
    // UserStatusInActive 未激活
    UserStatusInActive  = 1

    // UserStatusActived 已激活
    UserStatusActived   = 2

    // UserStatusFrozen 已冻结
    UserStatusFrozen    = 3
)
