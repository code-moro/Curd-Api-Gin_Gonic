package Controllers

import (
	"Rest-api/Database"
	"Rest-api/Models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
    "strconv"
	"github.com/gin-gonic/gin"
)

var collection=Database.DatabaseConnect()

func GetBooks(c *gin.Context){

	var books []Models.Book
	cur, err := collection.Find(context.TODO(), bson.M{})
	
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"Data":err.Error()})
		return
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		var book Models.Book
		err := cur.Decode(&book) 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
            return
		}

		// add item our array
		books = append(books, book)
	}
	c.IndentedJSON(http.StatusOK,gin.H{"Data":books})
}

func GetBooksById(c *gin.Context){
	
	var book Models.Book
	i,_:=strconv.Atoi(c.Param("id"))

	filter := bson.M{"_id": i}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK,gin.H{"Data":book})
}

func  CreateBooks(c *gin.Context){

	var newbook Models.Book

	if err:=c.BindJSON(&newbook);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
            return
	}
	insertResult, err := collection.InsertOne(context.TODO(),newbook)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"Data":insertResult})
}

func UpdateBook(c *gin.Context){
	i,_:=strconv.Atoi(c.Param("id"))
	var newbook Models.Book
	if err := c.BindJSON(&newbook); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"data": err.Error()})
            return
        }
	update:= bson.M{"title": newbook.Title, "author": newbook.Author, "quantity": newbook.Quantity}
    result, err := collection.UpdateOne(context.TODO(), bson.M{"_id": i}, bson.M{"$set": update})
        if err != nil {
            c.JSON(http.StatusInternalServerError,gin.H{"data": err.Error()})
            return
        }
		var updatedBook Models.Book
        if result.MatchedCount == 1 {
            err := collection.FindOne(context.TODO(), bson.M{"_id": i}).Decode(&updatedBook)
            if err != nil {
                c.JSON(http.StatusInternalServerError,gin.H{"data": err.Error()})
                return
            }
        }
	c.IndentedJSON(http.StatusOK,updatedBook)
}

func DeleteBook(c* gin.Context){
	i,_:=strconv.Atoi(c.Param("id"))


	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": i})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}

	if result.DeletedCount < 1 {
		c.JSON(http.StatusNotFound,gin.H{"data": "User with specified ID not found!"})
		return
	}

	c.JSON(http.StatusOK,gin.H{"data": "User successfully deleted!"})
}