package router

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markhuang1212/siribus/internal/businfo"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("EtaToDiamondHill", func(c *gin.Context) {
		ret, _ := businfo.GetTimeFromNow("286M", businfo.PoTaiStreetStation, "1")

		ans := ""

		if len(ret) >= 2 {
			ans = fmt.Sprintf(`The first bus to Diamond Hill will arrive in %.1f minutes, and the second bus will arrive in %.1f minutes.`,
				ret[0].Minutes(), ret[1].Minutes())
		}

		c.String(http.StatusOK, ans)
	})
	return r
}
