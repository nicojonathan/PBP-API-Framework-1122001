package controllers

import (
	m "belajar_gin/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	db := Connect()
	defer db.Close()

	err := c.Request.ParseForm()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form"})
        return
    }

    name := c.Request.Form.Get("name")
    age := c.Request.Form.Get("age")
    address := c.Request.Form.Get("address")
    email := c.Request.Form.Get("email")
    password := c.Request.Form.Get("password")

    query := "SELECT * FROM users WHERE 1"

    if name != "" {
        query += fmt.Sprintf(" AND name='%s'", name)
    }

    if age != "" {
        query += fmt.Sprintf(" AND age=%s", age)
    }

    if address != "" {
        query += fmt.Sprintf(" AND address='%s'", address)
    }

    if email != "" {
        query += fmt.Sprintf(" AND email='%s'", email)
    }

    if password != "" {
        query += fmt.Sprintf(" AND password='%s'", password)
    }

    fmt.Println(query)

    rows, err := db.Query(query)
    if err != nil {
        // c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error! Database Query Failed!"})
		sendErrorResponse(c, 500, "INternal Server Error! Database Query Failed!")
        return
    }
    defer rows.Close()

    var users []m.User // Change to slice of pointers to User

    for rows.Next() {
        var user m.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password); err != nil {
            sendErrorResponse(c, 500, "Internal Server Error! Fail to Scan Row")
            return
        } else {
            users = append(users, user) // Append pointer to User
        }
    }

    if len(users) == 0 {
        sendErrorResponse(c, 404, "Data not found!")
        return
    }

   sendUserSuccessResponse(c, "Get User Successful!", users)
}

func InsertUser(c *gin.Context) {
    db := Connect()
    defer db.Close()

    err := c.Request.ParseForm()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form"})
        return
    }

    query := "INSERT INTO users (`name`, `age`, `address`, `email`, `password`) VALUES (?,?,?,?,?)"

    name := c.Request.Form.Get("name")
    age := c.Request.Form.Get("age")
    address := c.Request.Form.Get("address")
    email := c.Request.Form.Get("email")
    password := c.Request.Form.Get("password")

	if (name == "" || age == "" || address == "" || email == "" || password == "") {
		sendErrorResponse(c, 400, "Bad request! You must include all parameters!")
		return
	}

	_, err = db.Exec(query, name, age, address, email, password)
	if err != nil {
        sendErrorResponse(c, 500, "Internal Server Error! Database query failed!")
		return
	}

    //fmt.Print(result)

	sendSuccessResponse(c, "Insert User Successful!")

}

func UpdateUser(c *gin.Context) {
    db := Connect()
	defer db.Close()

	queryGetUser := "SELECT * FROM users WHERE ID = ?"

	err := c.Request.ParseForm()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form"})
        return
    }

	id,_ := strconv.Atoi(c.Request.Form.Get("id"))

	result, errGetUser := db.Query(queryGetUser, id)
	if errGetUser != nil {
        fmt.Print(errGetUser)
		sendErrorResponse(c, 500, "Internal Server Error! Database query failed!")
		return
	}

	var user m.User
	var users []m.User
	
	if result.Next() {
		errScan := result.Scan(&user.ID, &user.Name, &user.Age, &user.Address, &user.Email, &user.Password)

		if errScan != nil {
			fmt.Println(errScan)
			sendErrorResponse(c, 500, "Internal Server Error! Fail to Scan!")
			return
		}
	}else{
		sendErrorResponse(c, 404, "account not found!")
		return
	}

	name := c.Request.Form.Get("name")
    age := c.Request.Form.Get("age")
    address := c.Request.Form.Get("address")
	email := c.Request.Form.Get("email")
	password := c.Request.Form.Get("password")

	if name != "" {
		user.Name = name
	}

    if age != "" {
		user.Age,_ = strconv.Atoi(age)
	}

    if address != "" {
		user.Address = address
	}

	if email != "" {
		user.Email = email
	}

	if password != "" {
		user.Password = password
	}

	users = append(users, user)

	queryUpdate := "UPDATE users SET Name = ?, Age = ?, Address = ?, Email = ?, Password = ? WHERE ID = ?"

	resultUpdate, errUpdate := db.Exec(queryUpdate, user.Name, user.Age, user.Address, user.Email, user.Password, id)

	if errUpdate != nil {
		fmt.Println(errUpdate)
		sendErrorResponse(c, 500, "Database Query Fail")
		return
	}

	rowsAffected,_:= resultUpdate.RowsAffected()
	if rowsAffected == 0 {
		sendErrorResponse(c, 404, "Data not found!")
		return
	}

	sendUserSuccessResponse(c, "Update Successful", users)
}

func DeleteUser(c *gin.Context) {
    db := Connect()
    defer db.Close()

    err := c.Request.ParseForm()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse form"})
        return
    }

    query := "DELETE FROM users WHERE id=?"

	// Read From Query Param
	id := c.Request.Form.Get("id")


	if (id == "") {
		sendErrorResponse(c, 400, "Bad Request! You need to insert id!")
		return
	}


	result, err := db.Exec(query, id)
	if (err != nil) {
		sendErrorResponse(c, 500, "Internal Server Error! Database Query Fail!")
		return
	}

	amountRowsAffected, _ := result.RowsAffected()

	if (amountRowsAffected == 0) {
		sendErrorResponse(c, 404, "Data not found!")
        return
	}else{
		sendSuccessResponse(c, "Delete User Successfull!")
	}
}
