package product

import (
	"casacomigo/app/models/product"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

func Find(titulo string) ([]product.Product) {
    session, err := mgo.Dial("127.0.0.1")
    if err != nil {
	    panic(err)
    }

    defer session.Close()
    session.SetMode(mgo.Monotonic, true)

    c := session.DB("casacomigo").C("product")

    var results []product.Product
    err = c.Find(bson.M{"titulo":bson.RegEx{".*"+titulo+"*", "i"}}).All(&results)
	if err != nil {
    	panic(err)
    }
    return results; 
}

func FindProduct(code int) (product.Product) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("product")

	result := product.Product{}
	err = c.Find(bson.M{"codigo": code}).One(&result)
    if err != nil {
    	panic(err)
    }
    return result; 
}

func List(limite, pular int) ([]product.Product) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("product")

	var results []product.Product
	err = c.Find(bson.M{}).Skip(pular).Limit(limite).All(&results)
    if err != nil {
    	panic(err)
    }
    return results;
}

func ListOffers(limite, pular int) ([]product.Product) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
 
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
 
 	c := session.DB("casacomigo").C("product")

	var results []product.Product
	err = c.Find(bson.M{"oferta": true}).Skip(pular).Limit(limite).All(&results)
    if err != nil {
    	panic(err)
    }
    return results;
}