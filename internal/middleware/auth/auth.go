package auth

/*
import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"time"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func SetUpAuth(r *gin.Engine) gin.HandlerFunc {
	log.Println("Setting up auth")
	authMiddleware, err := auth.GetAuthMiddleware()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authMiddlewareFunc := authMiddleware.MiddlewareFunc()

	r.NoRoute(authMiddlewareFunc, func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	r.Group("/auth").POST("/login", authMiddleware.LoginHandler).GET("/refresh_token", authMiddleware.RefreshHandler)

	return authMiddlewareFunc
}

func GetAuthMiddleware() (*jwt.GinJWTMiddleware, error) {

	//add os.Getenv() for some params
	return jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret"),
		Timeout:     15 * time.Minute,
		MaxRefresh:  20 * time.Minute,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			//function to authenticate (change to get users from database)
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			//add repository to validate user
			if (userID == "admin" && password == "admin") || (userID == "magneto" && password == "magneto") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			//function to authorize the client (could be by identifying the client)
			if v, ok := data.(*User); ok && (v.UserName == "admin" || v.UserName == "magneto"){
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		//TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenLookup: "header: Authorization",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
}
 */

