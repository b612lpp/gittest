package data

//ItemTemplate declares items metadata
type ItemTemplate struct {
	Name string
}

//ItemsDBstruct emulates DB
type ItemsDBstruct struct {
	Items []ItemTemplate
}

//ShopDB DB'sinstance
var ShopDB ItemsDBstruct

//AddItem adds record to emulated DB
func (db *ItemsDBstruct) AddItem(i ItemTemplate) {
	db.Items = append(db.Items, i)

}
