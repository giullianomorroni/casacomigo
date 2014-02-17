package account

import (
	"fmt"
	"casacomigo/app/models/account"
	"casacomigo/app/models/site"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//** PUBLIC FUNCTIONS
func Find(url_name string) (p *account.Account, err error) {
	return p, err
}

func FindCouple(noivo, noiva string) ([]account.Account) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("account")

	var results []account.Account
	//err = c.Find( bson.M{"$or": [] bson.M{ bson.M{"noivo":{noivo} , bson.M{"noiva":{noiva}} } } } ).One(&results)
	err = c.Find(bson.M{"noivo":noivo,"noiva":noiva}).All(&results)
    if err != nil {
    	panic(err)
    }
    return results; 
}

func UpdateAccountStatus(apelido ,status string) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("account")

	colQuerier := bson.M{"apelido": apelido}
	change := bson.M{"$set": bson.M{"status": status}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}
}

func FindAccount(apelido string) (account.Account) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("account")

	result := account.Account{}
	err = c.Find(bson.M{"apelido": apelido}).One(&result)
    if err != nil {
    	panic(err)
    }
    return result; 
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
