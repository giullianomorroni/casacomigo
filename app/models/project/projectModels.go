package project

type (
	Project struct {
		Name			string	`bson:"name"`
		Description		string	`bson:"description"`
		Content			string	`bson:"content"`
		Source			string	`bson:"source"`
		Complexity		[]int	`bson:"complexity"`
		Url				string	`bson:"url"`
	}

)
