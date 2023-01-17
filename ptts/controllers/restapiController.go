package controllers

import (
	"bufio"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

// "/restapi"
func RestApiGetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "restapi.html", gin.H{
			"title":   "RESTApi",
			"restapi": "active",
		})
	}
}

func RestApiMsgHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		SERVER_IP := os.Getenv("SERVER_IP")
		API_SERVER_PORT := os.Getenv("API_SERVER_PORT")

		realAddr := "http://" + SERVER_IP + ":" + API_SERVER_PORT

		// step 1: resolve proxy address, change scheme and host in requets
		req := c.Request
		proxy, err := url.Parse(realAddr)

		if err != nil {
			c.String(500, "error::remote url")
			return
		}

		req.URL.Scheme = proxy.Scheme
		req.URL.Host = proxy.Host

		// step 2: use http.Transport to do request to real server.
		transport := http.DefaultTransport
		resp, err := transport.RoundTrip(req)

		if err != nil {
			c.String(500, "error::roundtrip")
			return
		}

		// step 3: return real server response to upstream.
		for k, vv := range resp.Header {
			for _, v := range vv {
				c.Header(k, v)
			}
		}

		defer resp.Body.Close()

		bufio.NewReader(resp.Body).WriteTo(c.Writer)
	}
}
