package account

import (
	"fmt"
	"casacomigo/app/models/account"
	"casacomigo/app/models/site"
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

//** PUBLIC FUNCTIONS
func Find(url_name string) (p *account.Account, err error) {
	return p, err
}

func Register(account account.Account) (p *account.Account, err error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("account")
 
 	err =  c.Insert(account)
 	if err != nil {
		fmt.Print(err)
	}
	return p, err
}

func RegisterSite(site site.Site) (p *site.Site, err error) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("site")
 
 	err =  c.Insert(site)
 	if err != nil {
		fmt.Print(err)
	}
	return p, err
}
