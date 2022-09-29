package item

type ItemFormatter struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Price string `json: "price"`
	Cost  string `json: "cost"`
}

func FormatItem(item Item) ItemFormatter {
	ItemFormatter := ItemFormatter{}
	ItemFormatter.ID = item.ID
	ItemFormatter.Name = item.Name
	ItemFormatter.Price = item.Price
	ItemFormatter.Cost = item.Cost

	return ItemFormatter
}

func FormatItems(item []Item) []ItemFormatter {
	ItemsFormatter := []ItemFormatter{}

	for _, item := range item {
		ItemFormatter := FormatItem(item)
		ItemsFormatter = append(ItemsFormatter, ItemFormatter)
	}

	return ItemsFormatter
}

type ItemDetailFotmatter struct {
	ID    int    `json: "id"`
	Name  string `json: "name"`
	Price string `json: "price"`
	Cost  string `json: "cost"`
}

func FormatItemDetail(item Item) ItemDetailFotmatter {
	ItemDetailFotmatter := ItemDetailFotmatter{}
	ItemDetailFotmatter.ID = item.ID
	ItemDetailFotmatter.Name = item.Name
	ItemDetailFotmatter.Price = item.Price
	ItemDetailFotmatter.Cost = item.Cost

	return ItemDetailFotmatter

}
