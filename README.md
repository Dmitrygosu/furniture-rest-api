# REST API для управления сущностями мебели

#Автор: Игнатьев Дмитрий
#Вуз: Улгту
#Группа:Истбд-32


## Описание проекта

Этот проект представляет собой REST API, написанный на языке Go с использованием стандартной библиотеки и маршрутизатора Gorilla Mux. 
API предоставляет возможность управления сущностями мебели, хранящимися в JSON файле на сервере.

### Сущность Мебель

Каждая сущность мебели содержит следующие поля:

- Название мебели
- Производитель
- Высота
- Ширина
- Длина

## Возможности API

API предоставляет следующие операции с мебелью:

- **POST /furniture**: Создание новой сущности мебели.
- **GET /furniture**: Получение всех сущностей мебели.
- **GET /furniture/{id}**: Получение сущности мебели по её ID.
- **PUT /furniture/{id}**: Полное обновление сущности мебели по её ID.
- **PATCH /furniture/{id}**: Частичное обновление сущности мебели по её ID.
- **DELETE /furniture/{id}**: Удаление сущности мебели по её ID.

## Установка необходимых пакетов

Для работы проекта требуется установка пакета Gorilla Mux. Выполните следующие шаги:

1. Откройте терминал 
2. Установите Gorilla Mux, выполнив команду:
" go get -u github.com/gorilla/mux "
3. При запуске проекта в GoLand он сам установит пакет при нажатии на ошибку появившуюся в терминале.

### Запуск
### Используя Visual Studio Code

1. Откройте проект в Visual Studio Code.
2. Убедитесь, что у вас установлено Go расширение и библиотека Gorilla Mux .
3. Настройте файл `launch.json` для запуска приложения при необходимости.
4. Запустите отладку для запуска сервера на порту 8080.

### Используя терминал

1. Перейдите в папку `cmd/furniserver`.
2. Выполните команду:
" go run .\furniserver.go "

### Используя GoLand 
1. Откройте проект в GoLand.
2. Убедитесь, что у вас установлено Go расширение и библиотека Gorilla Mux .
3. Нажмите зеленую кнопку запуска в правом верхнем углу выбрав компилятор GO 1.22.4
4. Если это не помогло - проверьте работоспособность  выбранной конфигурации запуска или нажмите пкм по папке furniserver и Run.

###   Тестирование
1. #  Настоятельно рекомендуется использование программы Postman для удобной проверки моей работы
   

 
