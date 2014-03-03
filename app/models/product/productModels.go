package product


type (
	Product struct {
		Titulo			string	`bson:"titulo"`
		Codigo			int		`bson:"codigo"`
		Preco			float64	`bson:"preco"`
		PrecoOriginal	float64	`bson:"preco_original"`
		Descricao		string	`bson:"descricao"`
		Informacoes		interface{} `bson:"informacoes"`
		Oferta			boolean `bson:"oferta"`
	}

)
