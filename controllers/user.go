package controllers

import (
	"encoding/json"
	"fmt"
	"sharit-backend/models"
	"strconv"
	"time"
)

// UserController does everything related to steam login
type UserController struct {
	BaseController
}

// Login user
func (c *UserController) Login() {
	mail := c.GetString("mail")
	pass := c.GetString("pass")
	u, err := models.FindUserByMail(mail)
	if err == nil {
		if pass == u.Pass {
			var r reg
			r.Token = u.Token
			r.Iduser = u.IDuser
			c.Data["json"] = r
			c.ServeJSON()
		} else {
			c.Data["json"] = "wrong pass"
			c.ServeJSON()
		}

	} else {
		c.Data["json"] = "mail not registered"
		c.ServeJSON()
	}
}

type reg struct {
	Token  string `bson:"token,omitempty"`
	Iduser string `bson:"iduser,omitempty"`
}

// Register register
func (c *UserController) Register() {

	name := c.GetString("name")
	surname := c.GetString("surname")
	stars := "0"
	mail := c.GetString("mail")
	pass := c.GetString("pass")
	var u models.User
	u.IDuser = EncodeID64(mail, name, surname)
	u.Email = mail
	u.Pass = pass
	u.Name = name
	u.Stars = stars
	coordx, _ := c.GetInt("X")
	coordy, _ := c.GetInt("Y")
	u.X = coordx
	u.Y = coordy
	u.Token, _ = EncodeToken(u.IDuser, pass)
	u.Create()
	var r reg
	r.Token = u.Token
	r.Iduser = u.IDuser
	c.Data["json"] = r
	c.ServeJSON()
}

// RegisterDebug register
func (c *UserController) RegisterDebug() {

	name := c.GetString("name")
	surname := c.GetString("surname")
	stars := "0"
	mail := c.GetString("mail")
	pass := c.GetString("pass")
	var u models.User
	u.IDuser = EncodeID64(mail, name, surname)
	u.Email = mail
	u.Pass = pass
	u.Name = name
	u.Stars = stars
	u.Create()
	c.ServeJSON()
}

//EditProfile : only can update email and password
func (c *UserController) EditProfile() {

	mail := c.GetString("mail")
	myToken := c.GetString("token")
	id, err := DecodeToken(myToken)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = "error token id"
		c.ServeJSON()
	}
	coordx, _ := c.GetInt("X")
	coordy, _ := c.GetInt("Y")
	var u models.User
	u.IDuser = id
	u.Email = mail
	u.X = coordx
	u.Y = coordy
	err = u.UpdateUser()
	if err != nil {
		fmt.Println("error al fer update")
	} else {
		fmt.Println("update ok")
	}

	c.Data["json"] = "ok"
	c.ServeJSON()
}

// GetAll get all the users
func (c *UserController) GetAll() {
	users, _ := models.GetAllUsers()
	_, er := json.Marshal(users)
	if er != nil {
		//
		c.Data["json"] = "error no users"
	} else {
		c.Data["json"] = users
	}
	c.ServeJSON()
}

// Get get a user
func (c *UserController) Get() {

	id := c.GetString("id")

	u, err := models.FindUserByID(id)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// PutItem get a user
func (c *UserController) PutItem() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	stars := "0"
	image := c.GetString("image")
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	var i models.Item
	stt := token + name + time.Now().String()
	i.ID = hash(stt)
	i.ItemName = name
	i.Description = description
	i.Stars = stars
	i.Image = image
	i.LastSharit = time.Now()
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		fmt.Println("ok item")
		u.PutItemModel(i)
		c.Data["json"] = u
	}
	c.ServeJSON()

}

// GetItems return user items
func (c *UserController) GetItems() {
	token := c.GetString("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		c.Data["json"] = u.ItemsUser
	}
	c.ServeJSON()

}

// GetItemsRadi return user items
func (c *UserController) GetItemsRadi() {
	token := c.GetString("token")
	iduser, _ := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	fmt.Println(iduser)
	if err == nil {
		items, err := models.GetItemsRadi(u.X, u.Y)
		if err == nil {
			c.Data["json"] = items
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// GetItem return user items
func (c *UserController) GetItem() models.Item {
	token := c.GetString("token")
	iduser, _ := DecodeToken(token)
	idItem := c.GetString("idItem")
	u, err := models.FindUserByID(iduser)
	var item models.Item
	uintID, _ := strconv.ParseUint(idItem, 10, 32)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		items := u.ItemsUser
		for _, it := range items {
			if it.ID == uintID {
				item = it
			}
		}
	}
	return item
}

// PutItemDebug get a user
func (c *UserController) PutItemDebug() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	stars := "0"
	image := c.GetString("image")
	iduser := c.GetString("id")
	var i models.Item
	i.ItemName = name
	i.Description = description
	i.Stars = stars
	i.Image = image
	u, err := models.FindUserByID(iduser)
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		u.PutItemModel(i)
		c.Data["json"] = u
	}

	c.ServeJSON()

}

// PutPeticioRadi put peticio al radi
func (c *UserController) PutPeticioRadi() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	var p models.Peticio
	p.IDuser = iduser
	p.Name = name
	p.To = ""
	p.Descripcio = description
	p.X = u.X
	p.Y = u.Y
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		p.Create()
		c.Data["json"] = p
	}
	c.ServeJSON()
}

// PutPeticioRadiDebug get a user
func (c *UserController) PutPeticioRadiDebug() {
	//rebre el token i verificar si es coorrecte
	name := c.GetString("name")
	description := c.GetString("description")
	iduser := c.GetString("iduser")

	u, err := models.FindUserByID(iduser)
	var p models.Peticio
	p.IDuser = iduser
	p.Name = name
	p.To = ""
	p.Descripcio = description
	p.X = u.X
	p.Y = u.Y
	if err != nil {
		c.Data["json"] = "user not found"
	} else {
		p.Create()
		c.Data["json"] = p
	}
	c.ServeJSON()
}

// PutPeticioUsuari get a user
func (c *UserController) PutPeticioUsuari() {
	//fer una peticio especifica a un usuari
	token := c.GetString("token")
	iduser, _ := DecodeToken(token)
	userto := c.GetString("userTo")
	itemId := c.GetString("itemId")
	u, _ := models.FindUserByID(userto)
	uPet, _ := models.FindUserByID(iduser)
	var pet models.Peticio
	pet.Descripcio = c.GetString("description")
	pet.IDuser = iduser
	pet.Name = c.GetString("name")
	pet.To = userto
	pet.X = uPet.X
	pet.Y = uPet.Y
	pet.ItemID = itemId
	u.PutPeticio(pet)
	c.Data["json"] = "ok"
	c.ServeJSON()

}

// GetPeticionsRadiUser get a user
func (c *UserController) GetPeticionsRadiUser() {
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		peticions, err := models.GetPeticionsRadi(u.X, u.Y)
		if err == nil {
			c.Data["json"] = peticions
			c.ServeJSON()
		}
	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// GetPeticionsUsuari get a user
func (c *UserController) GetPeticionsUsuari() {
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		c.Data["json"] = u.PeticionsUser
		c.ServeJSON()

	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

// PutFavourite put a favourite to a user
func (c *UserController) PutFavourite() {
	iditem := c.GetString("idItem")
	idowner := c.GetString("idowner")
	token := c.GetString("token")
	idusuari, err := DecodeToken(token)
	//buscar owner
	o, err := models.FindUserByID(idowner)
	//buscar objecte dins owner
	item, err := o.FindFavouriteByID(iditem)
	//put objecte a usuari
	u, err := models.FindUserByID(idusuari)
	if err != nil {
		c.Data["json"] = "error user not found"
	} else {
		u.PutFavouriteModel(item, idowner)
		c.Data["json"] = u
	}
	c.ServeJSON()
}

//GetFavouritesUsuari get the user favourites
func (c *UserController) GetFavouritesUsuari() {
	token := c.GetString("token")
	iduser, err := DecodeToken(token)
	u, err := models.FindUserByID(iduser)
	if err == nil {
		c.Data["json"] = u.FavUser
		c.ServeJSON()

	} else {
		c.Data["json"] = "error a les petcions"
		c.ServeJSON()
	}
}

//PutCoordenades put cordenades for the user
func (c *UserController) PutCoordenades() {
	myToken := c.GetString("token")
	id, err := DecodeToken(myToken)
	if err != nil {
		c.Data["json"] = "error token id"
		c.ServeJSON()
	}
	u, err := models.FindUserByID(id)
	coordx, _ := c.GetInt("X")
	coordy, _ := c.GetInt("Y")
	u.X = coordx
	u.Y = coordy
	err = u.UpdateUserCoords()
	if err != nil {
		fmt.Println("error al fer update")
	} else {
		fmt.Println("update ok")
	}
}
