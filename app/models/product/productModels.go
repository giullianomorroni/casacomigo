package product


type (
	Product struct {
		Titulo			string	`bson:"titulo"`
		Codigo			int	`bson:"codigo"`
		Preco			float64	`bson:"preco"`
		Descricao		string	`bson:"descricao"`
		Informacoes		interface{} `bson:"informacoes"`
	}

)
