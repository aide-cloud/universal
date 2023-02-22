package routes

import (
	"github.com/aide-cloud/universal/aerror"
	"github.com/aide-cloud/universal/gin/response"
	"github.com/gin-gonic/gin"
	"path"
)

func FileUpload(savePath string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取文件
		file, err := c.FormFile("file")
		if err != nil {
			//c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			response.JSON(c, nil, aerror.New(aerror.WithCode(aerror.ErrCodeInvalidParam), aerror.WithErr(err)))
			return
		}

		// 上传文件至指定目录
		err = c.SaveUploadedFile(file, path.Join(savePath, file.Filename))
		if err != nil {
			response.JSON(c, nil, aerror.New(aerror.WithCode(aerror.ErrCodeInvalidParam), aerror.WithErr(err)))
			return
		}

		response.JSON(c, gin.H{"filename": file.Filename}, nil)
	}
}
