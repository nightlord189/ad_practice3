package model

func GetInvests() []Invest {
	result := make([]Invest, 0)
	result = append(result, Invest{
		Stoimost: 291.39,
		Period:   "2018-2019",
		Product:  "Алакольский район",
		Iniciatr: "ТОО «Алакөл Flesh»",
		ID:       "1",
	})
	result = append(result, Invest{
		Stoimost: 40,
		Period:   "2018-2019",
		Product:  "Балхашский район",
		Iniciatr: "ИП \"Онербеков\"",
		ID:       "2",
	})
	result = append(result, Invest{
		Stoimost: 53.5,
		Period:   "2018-2019",
		Product:  "Енбекшиказахский район",
		Iniciatr: "ТОО «Алматинская дорожно-строительная компания (АДСК)»",
		ID:       "3",
	})
	result = append(result, Invest{
		Stoimost: 110.4,
		Period:   "2017-2019",
		Product:  "Енбекшиказахский район",
		Iniciatr: "КХ \"Усенов\"",
		ID:       "4",
	})
	result = append(result, Invest{
		Stoimost: 814,
		Period:   "2016-2019",
		Product:  "Енбекшиказахский район",
		Iniciatr: "ТОО \"Энергия Алемы\"",
		ID:       "5",
	})
	result = append(result, Invest{
		Stoimost: 150,
		Period:   "2018-2019",
		Product:  "Жамбылский район",
		Iniciatr: "ТОО «Viratec»",
		ID:       "6",
	})
	return result
}
