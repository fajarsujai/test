package main

import (
	"fmt"
	"log"
	"net/http"
	"populationdatago/auth"
	"populationdatago/handler"
	"populationdatago/helper"
	"populationdatago/kecamatan"
	"populationdatago/kelurahan"
	"populationdatago/kota"
	"populationdatago/memberdetail"
	"populationdatago/provinsi"
	"populationdatago/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "fajarsujai:Secret$123@tcp(127.0.0.1:3306)/test_abera?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	provinsiRepository := provinsi.NewRepository(db)
	kotaRepository := kota.NewRepository(db)
	kecamatanRepository := kecamatan.NewRepository(db)
	kelurahanRepository := kelurahan.NewRepository(db)
	memberdetailRepository := memberdetail.NewRepository(db)

	userService := user.NewService(userRepository)
	provinsiService := provinsi.NewService(provinsiRepository)
	kotaService := kota.NewService(kotaRepository)
	kecamatanService := kecamatan.NewService(kecamatanRepository)
	kelurahanService := kelurahan.NewService(kelurahanRepository)
	memberdetailService := memberdetail.NewService(memberdetailRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)
	provinsiHandler := handler.NewProvinsiHandler(provinsiService)
	kotaHandler := handler.NewKotaHandler(kotaService)
	kecamatanHandler := handler.NewKecamatanHandler(kecamatanService)
	kelurahanHandler := handler.NewKelurahanHandler(kelurahanService)
	memberdetailHandler := handler.NewMemberDetailHandler(memberdetailService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/profile_pict", authMiddleware(authService, userService), userHandler.UploadAvatar)
	api.GET("/users/fetch", authMiddleware(authService, userService), userHandler.FetchUser)

	api.GET("/provinsi/fetch", authMiddleware(authService, userService), provinsiHandler.GetProvinsis)
	api.GET("/provinsi/:id", authMiddleware(authService, userService), provinsiHandler.GetProvinsi)
	api.POST("/provinsi", authMiddleware(authService, userService), provinsiHandler.CreateProvinsi)
	api.PUT("/provinsi/:id", authMiddleware(authService, userService), provinsiHandler.UpdateProvinsi)
	api.DELETE("/provinsi/:id", authMiddleware(authService, userService), provinsiHandler.Deleted)

	api.GET("/kota/fetch", authMiddleware(authService, userService), kotaHandler.GetKotas)
	api.GET("/kota/:id", authMiddleware(authService, userService), kotaHandler.GetKota)
	api.POST("/kota", authMiddleware(authService, userService), kotaHandler.CreateKota)
	api.PUT("/kota/:id", authMiddleware(authService, userService), kotaHandler.UpdateKota)
	api.DELETE("/kota/:id", authMiddleware(authService, userService), kotaHandler.Deleted)

	api.GET("/kecamatan/fetch", authMiddleware(authService, userService), kecamatanHandler.GetKecamatans)
	api.GET("/kecamatan/:id", authMiddleware(authService, userService), kecamatanHandler.GetKecamatan)
	api.POST("/kecamatan", authMiddleware(authService, userService), kecamatanHandler.CreateKecamatan)
	api.PUT("/kecamatan/:id", authMiddleware(authService, userService), kecamatanHandler.UpdateKecamatan)
	api.DELETE("/kecamatan/:id", authMiddleware(authService, userService), kecamatanHandler.Deleted)

	api.GET("/kelurahan/fetch", authMiddleware(authService, userService), kelurahanHandler.GetKelurahans)
	api.GET("/kelurahan/:id", authMiddleware(authService, userService), kelurahanHandler.GetKelurahan)
	api.POST("/kelurahan", authMiddleware(authService, userService), kelurahanHandler.CreateKelurahan)
	api.PUT("/kelurahan/:id", authMiddleware(authService, userService), kelurahanHandler.UpdateKelurahan)
	api.DELETE("/kelurahan/:id", authMiddleware(authService, userService), kelurahanHandler.Deleted)

	api.GET("/memberdetail/fetch", authMiddleware(authService, userService), memberdetailHandler.GetMemberDetails)
	api.GET("/memberdetail/:id", authMiddleware(authService, userService), memberdetailHandler.GetMemberDetail)
	api.POST("/memberdetail", authMiddleware(authService, userService), memberdetailHandler.CreateMemberDetail)
	api.PUT("/memberdetail/:id", authMiddleware(authService, userService), memberdetailHandler.UpdateMemberDetail)
	api.DELETE("/memberdetail/:id", authMiddleware(authService, userService), memberdetailHandler.Deleted)

	router.Run()
}
func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		fmt.Println(user)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
