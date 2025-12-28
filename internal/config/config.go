package config

const (
	// 改为从环境变量读取
	// func init() {
	// 	secret := os.Getenv("JWT_SECRET_KEY")
	// 	if secret == "" {
	// 		panic("JWT_SECRET_KEY environment variable is required")
	// 	}
	// 	JWTSecretKey = secret
	// }
	// JWTSecretKey 用于 JWT 的签名密钥
	JWTSecretKey = "1AA070C9-AA2F-AE2F-E6D4-13FD1075B424@pepsi1145@!#ZMJ"
	// JWTIssuer 用于 JWT 的发行者字段
	JWTIssuer = "go-flash-sale"
)
