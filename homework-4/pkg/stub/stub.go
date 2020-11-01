package stub

func Scan() (map[string]string, error) {
	return map[string]string{
		"http://www.transflow.ru/service": "Сервис ТРАНСФЛОУ",
		"http://www.transflow.ru/about":   "О платформе ТРАНСФЛОУ",
	}, nil
}
