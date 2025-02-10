package songs

func ConvertToMap(slice []string) map[int]string { // Самописная функция преобразования среза в мапу, создающая ключи в виде порядкового номера каждому элементу
	resultingMap := make(map[int]string) // Инициализируем мапу

	for i := 0; i < len(slice); i++ { // Цикл, длящийся столько, сколько элементов в срезе
		resultingMap[i+1] = slice[i] // И записывающий каждый i элемент массива в мапу под ключом i+1
	}

	return resultingMap // Возвращаем мапу
}
