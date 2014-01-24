package project

import (
	"fmt"
	"casacomigo/app/models/project"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//** PUBLIC FUNCTIONS
func FindProject(url_name string) (p *project.Project, err error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(err)
	}
	defer session.Close()

	c := session.DB("arduino_awkward").C("project")

	err = c.Find(bson.M{"url": url_name}).One(&p)
	if err != nil {
		fmt.Println(err)
	}

	content,err := project.ReadContent(p.Content);
	p.Content = content;

	return p, err
}
