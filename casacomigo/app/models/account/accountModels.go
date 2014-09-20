package account


type (
	Account struct {
		Noivo			string	`bson:"noivo"`
		Noiva			string	`bson:"noiva"`
		DataCasamento	string	`bson:"dataCasamento"`
		TelefoneNoivo	string	`bson:"telefoneNoivo"`
		TelefoneNoiva	string	`bson:"telefoneNoiva"`
		EmailNoivo		string	`bson:"emailNoivo"`
		EmailNoiva		string	`bson:"emailNoiva"`
		Apelido			string	`bson:"apelido"`
		Senha			string	`bson:"senha"`
		Status			string	`bson:"status"`
		Lucro			float64	`bson:"lucro"`
	}

)
