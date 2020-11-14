package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"strings"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 //8M
	router.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		v4 := uuid.NewV4().String()
		log.Println("---------", file.Filename, strings.Replace(v4, "-", "", -1))
		dest := fmt.Sprint("./" + file.Filename)
		_ = context.SaveUploadedFile(file, dest)
	})
	_ = router.Run(":18080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
