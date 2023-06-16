package web

import (
	"github.com/RoyJoel/LZStudyingProjectBackend/package/config"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/web/controller"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/web/interceptor"
	"github.com/gin-gonic/gin"
)

func RunHttp() {
	r := gin.Default()
	//增加拦截器
	r.Use(interceptor.HttpInterceptor())
	//解决跨域
	r.Use(config.CorsConfig())

	authInfo := r.Group("/")
	{
		authInfo.GET("auth", controller.NewLZControllerImpl().Auth)
	}

	//路由组
	userInfo := r.Group("/user")
	{
		userInfo.POST("/signIn", controller.NewLZControllerImpl().SignIn)
		userInfo.POST("/signUp", controller.NewLZControllerImpl().SignUp)
		userInfo.POST("/resetPassword", controller.NewLZControllerImpl().ResetPassword)
		userInfo.POST("/update", controller.NewLZControllerImpl().UpdateUser)
	}

	playerInfo := r.Group("/player")
	{
		playerInfo.POST("/update", controller.NewLZControllerImpl().UpdatePlayer)
		playerInfo.POST("/getInfo", controller.NewLZControllerImpl().GetPlayerInfo)
		playerInfo.POST("/add", controller.NewLZControllerImpl().AddPlayer)
		playerInfo.POST("/search", controller.NewLZControllerImpl().SearchPlayer)
	}

	friendInfo := r.Group("/friend")
	{
		friendInfo.POST("/add", controller.NewLZControllerImpl().AddFriend)
		friendInfo.POST("/search", controller.NewLZControllerImpl().SearchFriend)
		friendInfo.POST("/delete", controller.NewLZControllerImpl().DeleteFriend)
		friendInfo.POST("/getAll", controller.NewLZControllerImpl().GetAllFriends)
	}

	clubInfo := r.Group("/music")
	{
		clubInfo.POST("/getInfos", controller.NewLZControllerImpl().GetMusicInfos)
	}

	orderInfo := r.Group("/order")
	{
		orderInfo.POST("/add", controller.NewLZControllerImpl().AddOrder)
		orderInfo.POST("/delete", controller.NewLZControllerImpl().DeleteOrder)
		// orderInfo.POST("/search", controller.NewLZControllerImpl().GetOrderInfos)
		orderInfo.POST("/update", controller.NewLZControllerImpl().UpdateOrder)
		orderInfo.POST("/getInfos", controller.NewLZControllerImpl().GetOrderInfosByUserId)
		orderInfo.GET("/getAll", controller.NewLZControllerImpl().GetAllOrders)
	}

	addressInfo := r.Group("/address")
	{
		addressInfo.POST("/add", controller.NewLZControllerImpl().AddAddress)
		addressInfo.POST("/delete", controller.NewLZControllerImpl().DeleteAddress)
		addressInfo.POST("/getInfos", controller.NewLZControllerImpl().GetAddressInfos)
		addressInfo.POST("/update", controller.NewLZControllerImpl().UpdateAddress)
	}

	cartInfo := r.Group("/cart")
	{
		cartInfo.POST("/getInfo", controller.NewLZControllerImpl().GetCartInfo)
		cartInfo.POST("/addBill", controller.NewLZControllerImpl().AddBillToCart)
		cartInfo.POST("/deleteBill", controller.NewLZControllerImpl().DeleteBillInCart)
		cartInfo.POST("/assign", controller.NewLZControllerImpl().AssignCartForUser)
	}

	commodityInfo := r.Group("/commodity")
	{
		commodityInfo.POST("/add", controller.NewLZControllerImpl().AddCommodity)
		commodityInfo.POST("/delete", controller.NewLZControllerImpl().DeleteCommodity)
		// commodityInfo.POST("/search", controller.NewLZControllerImpl().getCommodityInfo)
		commodityInfo.POST("/update", controller.NewLZControllerImpl().UpdateCommodity)
		commodityInfo.GET("/getAll", controller.NewLZControllerImpl().GetAllCommodities)
	}

	optionInfo := r.Group("/option")
	{
		optionInfo.POST("/add", controller.NewLZControllerImpl().AddOption)
		optionInfo.POST("/update", controller.NewLZControllerImpl().UpdateOption)
		optionInfo.POST("/delete", controller.NewLZControllerImpl().DeleteOption)

	}

	// billInfo := r.Group("/bill")
	// {
	// billInfo.POST("/add", controller.NewLZControllerImpl().AddBill)
	// billInfo.POST("/delete", controller.NewLZControllerImpl().DeleteBill)
	// billInfo.POST("/search", controller.NewLZControllerImpl().getBillInfo)
	// billInfo.POST("/update", controller.NewLZControllerImpl().UpdateBill)
	// }

	r.Run(":8080")
}
