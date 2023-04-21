package router

import (
	"ZeroPassBackend/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
)

type Member struct {
	Address    string
	Identity_1 []byte
	Identity_2 []byte
	Licence    []byte
	verify     bool
}

// @Router /upload [post]
func UploadHandler(c *gin.Context) {
	address := c.PostForm("address")

	identity1, err := c.FormFile("identity_1")
	if err != nil {
		c.JSON(400, gin.H{"error": "identity_1 file is required"})
		return
	}

	identity2, err := c.FormFile("identity_2")
	if err != nil {
		c.JSON(400, gin.H{"error": "identity_2 file is required"})
		return
	}

	licence, err := c.FormFile("licence")
	if err != nil {
		c.JSON(400, gin.H{"error": "licence file is required"})
		return
	}

	// Read the uploaded files' content
	identity1Content, err := identity1.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error reading identity_1 file"})
		return
	}
	defer identity1Content.Close()

	identity2Content, err := identity2.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error reading identity_2 file"})
		return
	}
	defer identity2Content.Close()

	licenceContent, err := licence.Open()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error reading licence file"})
		return
	}
	defer licenceContent.Close()

	// Insert the data into MongoDB
	member := Member{
		Address:    address,
		Identity_1: readAllBytes(identity1Content),
		Identity_2: readAllBytes(identity2Content),
		Licence:    readAllBytes(licenceContent),
		verify:     false,
	}
	client, err := utils.NewMongoClient(utils.GetEnv("DB_CONNECTION"))
	if err != nil {
		log.Printf(err.Error())
		c.JSON(500, gin.H{"error": "Error connect into MongoDB"})
		return
	}
	_, err = client.InsertDocument("zeroPass", "member", member)
	if err != nil {
		log.Printf(err.Error())
		c.JSON(500, gin.H{"error": "Error inserting data into MongoDB"})
		return
	}

	c.JSON(200, gin.H{"message": "Upload successful"})
}

func AllMemberHandler(c *gin.Context) []Member {
	var memberList []Member
	// 從 MongoDB 取得所有會員資料
	client, err := utils.NewMongoClient(utils.GetEnv("DB_CONNECTION"))
	cur, err := client.FindAllDocuments("zeroPass", "member", nil)
	if err != nil {
		fmt.Println(err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var member Member
		err := cur.Decode(&member)
		if err != nil {
			fmt.Println(err)
		}
		memberList = append(memberList, member)
	}

	if err := cur.Err(); err != nil {
		fmt.Println(err)
	}
	// 返回結果
	return memberList
}

// @Router /verify [post]
func VerifyHandler(c *gin.Context) {
	// 驗證 ZKP credential
	// 與智能合約交互
	// 返回結果
}

func readAllBytes(reader io.Reader) []byte {
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatalf("Error reading file content: %v", err)
	}
	return content
}
