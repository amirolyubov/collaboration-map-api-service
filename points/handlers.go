package points

import (
	"cm-api/db"
	"cm-api/types"
	"context"
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetPoints(c *fiber.Ctx) error {
	chat := c.Query("chat")
	user := c.Query("user")

	if (len(chat) > 0 && len(user) > 0) {
		return c.SendString("you cannot use chat and user together");
	} else {
		if (len(chat) > 0) {
			return getPointsByChat(c)
		}
		if (len(user) > 0) {
			return getPointsByUser(c)
		}
		return c.SendString("poop")
	}
}

func getPointsByUser(c *fiber.Ctx) error {
	client := db.GetMongoClient()
	user := c.Query("user")

	pointsCollection := client.Database("points").Collection(user)

	filter := bson.D{}
	cursor, err := pointsCollection.Find(context.TODO(), filter)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("error in finding points")
	}

	var results []types.Point
	err = cursor.All(context.TODO(), &results)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("error in parsing points")
	}
	jsonBytes, _ := json.Marshal(results)

	c.Status(http.StatusOK)
	c.Write(jsonBytes)

	return nil
}
func getPointsByChat(c *fiber.Ctx) error {
	// client := db.GetMongoClient()
	
	chat := c.Query("chat")

	// points := client.Database()

	return c.SendString(chat)
}


func GetPoint(c *fiber.Ctx) error {
	return c.SendString("get point");
}

func EditPoint(c *fiber.Ctx) error {
	return c.SendString("edit point");
}

func DeletePoint(c *fiber.Ctx) error {
	return c.SendString("delete point");
}