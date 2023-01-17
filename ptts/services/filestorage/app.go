package servicies

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ServiceUploadSingleFile(storage_path string) gin.HandlerFunc {
	return func(c *gin.Context) {

		name := c.PostForm("name")
		email := c.PostForm("email")

		fmt.Println("name:", name, ", email:", email)

		// Multipart form
		form, err := c.MultipartForm()

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		files := form.File["myfile"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			dest_file := storage_path + "/" + filename

			// Upload the file to specific dst.
			if err := c.SaveUploadedFile(file, dest_file); err != nil {
				c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
				return
			}
		}
		ServiceShowStorage(storage_path)(c)
	}
}

func ServiceUploadMultiFile(storage_path string) gin.HandlerFunc {
	return func(c *gin.Context) {

		// Multipart form
		form, err := c.MultipartForm()

		if err != nil {
			c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		files := form.File["myfiles"]

		for _, file := range files {

			filename := filepath.Base(file.Filename)
			dest_file := storage_path + "/" + filename

			// Upload the file to specific dst.
			if err := c.SaveUploadedFile(file, dest_file); err != nil {
				c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
				return
			}
		}
		ServiceShowStorage(storage_path)(c)
	}
}

func ServiceShowStorage(storage_path string) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("ServiceShowStorage:", storage_path)

		dir_info := [](map[string]string){}

		files, err := ioutil.ReadDir(storage_path)
		if err != nil {
			log.Fatal(err)
		}

		for _, file := range files {
			var one_info = map[string]string{}

			//abs_path := storage_path + "/" + file.Name()

			one_info["extension"] = filepath.Ext(file.Name())
			one_info["basename"] = filepath.Base(file.Name())
			one_info["filename"] = one_info["basename"][0 : len(one_info["basename"])-len(one_info["extension"])]
			if file.IsDir() {
				one_info["state"] = "dir"
			} else if (file.Mode() & os.ModeSymlink) != 0 {
				one_info["state"] = "link"
			} else {
				one_info["state"] = "file"
			}

			one_info["size"] = strconv.FormatInt(file.Size(), 10)
			one_info["mode"] = file.Mode().String()
			one_info["atime"] = ""
			one_info["mtime"] = ""
			one_info["ctime"] = ""
			one_info["birthtime"] = file.ModTime().String()

			dir_info = append(dir_info, one_info)
		}

		c.JSON(http.StatusOK, dir_info)
	}
}

func ServiceClearStorage(storage_path string) gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println("ServiceClearStorage:", storage_path)

		err := os.RemoveAll(storage_path)
		if err != nil {
			log.Fatal(err)
		}

		err = os.MkdirAll(storage_path, os.ModeDir)
		if err != nil {
			log.Fatal(err)
		}

		ServiceShowStorage(storage_path)(c)
	}
}
