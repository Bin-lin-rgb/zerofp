package define

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

var (
	MysqlDSN   = ""
	EmqxKey    = "1f9c5b734fe27865"
	EmqxSecret = "lV9C2iefOp9Cr9BeiB5rr3N9CBolJjKk3HruhqEpHQxsuVD"
)

type UserClaim struct {
	Id       uint   `json:"id"`
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.RegisteredClaims
}

var (
	JwtKey = "iot-platform"
)

type M map[string]interface{}

func init() {
	cfg, err := ini.Load("F:\\GitRepository\\zerofp\\config\\config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}

	MysqlDSN = cfg.Section("mysql").Key("dsn").String()
}
