package router

import (
	"ZeroPassBackend/models"
	"ZeroPassBackend/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// @Router /upload [post]
func UploadHandler(c *gin.Context) {
	address := c.Param("address")
	// print c value
	log.Println(c)

	identity1, err := c.FormFile("identity_1")
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(400, gin.H{"error": "identity_1 file is required"})
		return
	}

	identity2, err := c.FormFile("identity_2")
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(400, gin.H{"error": "identity_2 file is required"})
		return
	}

	licence, err := c.FormFile("licence")
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(400, gin.H{"error": "licence file is required"})
		return
	}

	// Read the uploaded files' content
	identity1Content, err := identity1.Open()
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(500, gin.H{"error": "Error reading identity_1 file"})
		return
	}
	defer identity1Content.Close()

	identity2Content, err := identity2.Open()
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(500, gin.H{"error": "Error reading identity_2 file"})
		return
	}
	defer identity2Content.Close()

	licenceContent, err := licence.Open()
	if err != nil {
		log.Fatalf(err.Error())
		c.JSON(500, gin.H{"error": "Error reading licence file"})
		return
	}
	defer licenceContent.Close()

	// Insert the data into MongoDB
	member := models.Member{
		Address:    address,
		Identity_1: readAllBytes(identity1Content),
		Identity_2: readAllBytes(identity2Content),
		Licence:    readAllBytes(licenceContent),
		Verify:     false,
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

func GetAllMembersHandler(c *gin.Context) {
	client, err := utils.NewMongoClient(utils.GetEnv("DB_CONNECTION"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot connect to database"})
		return
	}

	var users []models.Member
	err = client.FindAll("zeroPass", "member", bson.M{}, &users)
	if err != nil {
		log.Println("Failed to find all users:", err)
	}

	c.JSON(http.StatusOK, users)
}

func GetMember(c *gin.Context) {
	address := c.Param("address")
	log.Printf("address: %s", address)
	client, err := utils.NewMongoClient(utils.GetEnv("DB_CONNECTION"))
	if err != nil {
		log.Printf(err.Error())
		c.JSON(500, gin.H{"error": "Error connect into MongoDB"})
		return
	}
	var user models.Member
	err = client.FindOne("zeroPass", "member", bson.M{"address": address}, &user)
	if err != nil {
		log.Println("Failed to find all users:", err)
	}

	c.JSON(200, user)
}

func UpdateMember(c *gin.Context) {
	address := c.Param("address")

	client, err := utils.NewMongoClient(utils.GetEnv("DB_CONNECTION"))
	if err != nil {
		log.Printf(err.Error())
		c.JSON(500, gin.H{"error": "Error connect into MongoDB"})
		return
	}
	// Update one document
	filter := bson.M{"address": address}
	update := bson.M{"Verify": true}
	result, err := client.UpdateOne("zeroPass", "member", filter, update)
	if err != nil {
		log.Println("Failed to update user:", err)
	} else {
		log.Printf("Updated %v documents\n", result.ModifiedCount)
	}

	c.JSON(200, gin.H{"message": "Member updated successfully"})
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
