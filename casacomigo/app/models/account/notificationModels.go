package account

type (
	Notification struct {
		Apelido			string	`bson:"apelido"`
		Tipo			string	`bson:"tipo"`
		Descricao		string	`bson:"descricao"`
		Data		string	`bson:"data"`
	}

)
