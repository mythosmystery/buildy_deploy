package main

import (
	"log"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type Webhook struct {
	Ref        string `json:"ref"`
	Before     string `json:"before"`
	After      string `json:"after"`
	CompareURL string `json:"compare_url"`
	Commits    []struct {
		ID        string `json:"id"`
		Message   string `json:"message"`
		Timestamp string `json:"timestamp"`
		URL       string `json:"url"`
		Author    struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		} `json:"author"`
	} `json:"commits"`
	Repository struct {
		Name        string `json:"name"`
		URL         string `json:"url"`
		Description string `json:"description"`
		Homepage    string `json:"homepage"`
		GitURL      string `json:"git_url"`
		SSHURL      string `json:"ssh_url"`
		CloneURL    string `json:"clone_url"`
	} `json:"repository"`
	Pusher struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"pusher"`
}

func main() {
	r := gin.Default()

	r.GET("/heartbeat", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/webhooks", func(c *gin.Context) {
		var jsonData Webhook

		if err := c.BindJSON(&jsonData); err != nil {
			c.JSON(400, gin.H{
				"status": "error",
				"error":  err,
			})
			return
		}

		if jsonData.Ref == "refs/heads/main" {
			log.Printf("Deploying...")
			data, err := exec.Command("sh", "-c", "cd /home/web/build_io && ./scripts/deploy.sh").Output()
			if err != nil {
				log.Printf("Error: %s", err)
			}
			log.Printf("Output: %s", data)
		}

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.Run("0.0.0.0:8001") // listen and serve on 0.0.0.0:8080
}
