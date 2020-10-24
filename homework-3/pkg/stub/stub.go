package stub

func Scan() (map[string]string, error) {
	m := make(map[string]string)
	m["http://www.transflow.ru/service"] = "Сервис ТРАНСФЛОУ"
	m["http://www.transflow.ru/about"] = "О платформе ТРАНСФЛОУ"
	return m, nil
}
