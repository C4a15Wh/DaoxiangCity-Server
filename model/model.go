package model

type User struct {
	ID                int    `gorm:"column:uid"`
	Mail              string `gorm:"column:email"`
	Nickname          string `gorm:"column:nickname"`
	Locale            string `gorm:"column:locale"`
	Score             string `gorm:"column:score"`
	Avatar            int    `gorm:"column:avatar"`
	Password          string `gorm:"column:password"`
	IP                string `gorm:"column:ip"`
	DarkMode          string `gorm:"column:is_dark_mode"`
	Permission        string `gorm:"column:permission"`
	RegisterAt        string `gorm:"column:register_at"`
	VerificationToken string `gorm:"column:verification_token"`
	RememberToken     string `gorm:"column:remember_token"`
	ForumUid          string `gorm:"column:forum_uid"`
}

func (User) TableName() string {
	return "DXC_users"
}

type TokenDecode struct {
	Aud string `json:"Aud"`
	Iss string `json:"Iss"`
}

type ClientConfig struct {
	GitAddress string `json:"GitAddress"`
}

type UpdateParams struct {
	Type      string `json:"Type"`    // Stable 稳定版  Test 测试版
	ClientVer string `json:"Version"` // exmpale: 1.0.1
	Token     string `json:"Token"`
}

type SkinServerResp struct {
	Token string `json:"token"` // 详见https://blessing.netlify.app/api/auth.html
}

type LoginParams struct {
	Email    string `json:"email"` // 详见https://blessing.netlify.app/api/auth.html
	Password string `json:"password"`
}

type ServerConf struct {
	SkinServer string    `yaml:"SkinServer"`
	Port       string    `yaml:"Port"`
	GitAddr    string    `yaml:"GitAddr"`
	SQL        SQLConfig `yaml:"SQLConfig"`
	JWT        JWT       `yaml:"JWT"`
}

type SQLConfig struct {
	Addr     string `yaml:"Address"`
	Port     string `yaml:"Port"`
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	DB       string `yaml:"DB"`
}

type JWT struct {
	Secret string `yaml:"Secret"`
}
