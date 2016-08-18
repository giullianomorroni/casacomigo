package account

import (
	"casacomigo/models/account"
	"casacomigo/models/site"
	"casacomigo/models/shopper"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
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
    err = c.Find( bson.M{"$or": [] bson.M{ bson.M{"noivo":bson.RegEx{".*"+noivo+"*", "i"}}, bson.M{"noiva":bson.RegEx{".*"+noiva+"*", "i"}} } }).All(&results)
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

func UpdateCoupleAmmount(casal, comprador string, presente float64) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close();
	session.SetMode(mgo.Monotonic, true)

 	c := session.DB("casacomigo").C("account")

	//find account
	result := account.Account{}
	colQuerier := bson.M{"apelido": casal}
	err = c.Find(colQuerier).One(&result)
	if err != nil {
    	panic(err)
    }
    lucro := result.Lucro;
	lucro += presente;

	//revel.TRACE.Printf("UpdateCoupleAmmount - presente: %s", presente);
	//revel.TRACE.Printf("UpdateCoupleAmmount - lucro total: %s", lucro);

	//update value
	change := bson.M{"$set": bson.M{"lucro": lucro}}
	err = c.Update(colQuerier, change)
	if err != nil {
		panic(err)
	}

	c = session.DB("casacomigo").C("bank")
	bank := account.Bank{};
	bank.Apelido = casal;
	const layout = "2 Jan, 2006 - 15:04"
	bank.Data = time.Now().Format(layout);
	bank.Comprador = comprador;
	bank.Valor = presente;
	c.Insert(bank);
	//revel.TRACE.Printf("Bank: %s", c);

	c = session.DB("casacomigo").C("notification")
	ntf := account.Notification{};
	ntf.Apelido = casal;
	ntf.Tipo = "Presente";
	ntf.Descricao = "Você ganhou mais um presente :) ";
	c.Insert(ntf);
	//revel.TRACE.Printf("Notification: %s", ntf);
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

func Login(apelido, senha string) (account.Account) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

 	c := session.DB("casacomigo").C("account")

	result := account.Account{}
	err = c.Find(bson.M{"apelido": apelido, "senha": senha}).One(&result)
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

	index := mgo.Index{
	    Key: []string{"apelido"},
	    Unique: true,
	    DropDups: false,
	    Background: false,
	    Sparse: true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		//revel.TRACE.Printf("Error %s", err)
		return p, err
	}

 	err =  c.Insert(account)
 	if err != nil {
		//revel.TRACE.Printf("Error %s", err)
		return p, err
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
		//revel.TRACE.Printf("Error %s", err)
	}
	return p, err
}

func RegisterShopper(shopper shopper.Shopper) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

 	c := session.DB("casacomigo").C("shopper")

 	err =  c.Insert(shopper)
 	if err != nil {
		//revel.TRACE.Printf("Error %s", err)
	}
}

func FindNotifications(apelido string) ([]account.Notification) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

 	c := session.DB("casacomigo").C("notification")

	var results []account.Notification
	err = c.Find(bson.M{"apelido": apelido}).All(&results)
    if err != nil {
    	panic(err)
    }
    return results;
}

func FindBanks(apelido string) ([]account.Bank) {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

 	c := session.DB("casacomigo").C("bank")

	var results []account.Bank
	err = c.Find(bson.M{"apelido": apelido}).All(&results)
    if err != nil {
    	panic(err)
    }
    return results;
}
