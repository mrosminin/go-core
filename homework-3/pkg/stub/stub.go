package stub

func Scan() (map[string]string, error) {
	m := make(map[string]string)
	m["http://www.transflow.ru/service"] = "Сервис ТРАНСФЛОУ"
	return m, nil
}
