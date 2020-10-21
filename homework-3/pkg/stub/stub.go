package stub

func Scan() (data map[string]string, err error) {
	return map[string]string{
		"http://www.transflow.ru/service": "Сервис ТРАНСФЛОУ",
	}, nil
}
