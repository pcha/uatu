package server

import (
	"fmt"
	"net/http"

	"the-one/internal/app/server/queue"
	saver2 "the-one/internal/pkg/saver"

	"github.com/gin-gonic/gin"
)

type DefaultServer struct {
	Port uint16
	Q    *queue.Queue
}

func NewDefaultServer(saver saver2.Saver, port uint16) *DefaultServer {
	return &DefaultServer{
		Q:    queue.New(saver),
		Port: port,
	}
}

func (s *DefaultServer) Start() error {
	s.Q.StartListening()

	r := gin.Default()
	r.GET("/ping", s.ping())
	r.POST("/:bucket", s.write())
	if err := r.Run(fmt.Sprintf(":%v", s.Port)); err != nil {
		return err
	}
	return nil
}

func (s *DefaultServer) ping() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

func (s *DefaultServer) write() func(c *gin.Context) {
	return func(c *gin.Context) {
		b := c.Param("bucket")
		data := saver2.NewFact()
		err := c.ShouldBindJSON(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		s.Q.Receive(b, data)
		c.JSON(http.StatusAccepted, gin.H{
			"received": "ok",
		})
	}
}
