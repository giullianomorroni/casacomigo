package account

type (
	Bank struct {
		Apelido			string	`bson:"apelido"`
		Data			string	`bson:"data"`
		Comprador		string	`bson:"comprador"`
		Valor			float64	`bson:"valor"`
	}

)
