package shopper

type (
	Shopper struct {
		Nome		string	`bson:"nome"`
		Telefone	string	`bson:"telefone"`
		Email		string	`bson:"email"`
		Casal		string	`bson:"casal"`
		Total		float64	`bson:"total"`
	}

)
