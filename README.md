Список всех команд - `/help`..

Идея этого бота возникла супер давно.
В частности раньше он служил для связи GTA SA:MP lua скрипта и чата сотрудников старшего состава Министерства Здравоохранения (в игре, очевидно).
Не прям что-бы бот - связующий элемент. Скорее клиент взаимодействия с базой данных.
В результате многих изменений часть функций была упущена и переписана (изначально он был на php, потом переписан под python, сейчас go)

Сейчас в боте реализованы в основном развлекательные команды (получить анекдот, шар предсказаний и т.д), так же реализована работа с OpenAI и языковой моделью gpt-3.5-turbo и gpt-4, а так же свой личный супер пупер уникальный неповторимый промпт, который придумывает шутки, оскорбления, обзывалки на имена (всё с матами и всё может реально оскорбить человека. gpt-4 творит чудеса).
Из взаимодействий с базой данных реализована проверка пользователя на нахождения в БД и получения финда из базы данных.
В проекте используется только получения сотруников онлайн из базы данных.
P.S. изначально планировал переписать и базу данных под postgres, но со временем принял решение оставить mysql и вся логика которая сейчас есть написана для mysql баз данных, но остался задел под postgres на будущее.

Проект развёрнут у меня на сервере и доступен по ссылке `https://vk.com/hmm_senior_bot`. Так же там можно пригласить бота в беседу.

Так же проект доступен и на hub.docker.com. `docker pull dnevsky/hmm-bot`

Так же написан docker-compose файл, который запускает приложение и базу данных mysql.

.env переменные:
```
DB_PASSWORD=<password>
VK_TOKEN=<token>
OPENAI_TOKEN=<token>

MYSQL_ROOT_PASSWORD=<password>
MYSQL_DATABASE=<db>
MYSQL_USER=<user>
MYSQL_PASSWORD=<password>

POSTGRES_DB=<db>
POSTGRES_USER=<user>
POSTGRES_PASSWORD=<password>
```

`make postgres` - запускает postgres базу данных через docker-compose. Необходимы `env` переменные в файле `.env`.

`make mysql` - запускает mysql базу данных через docker-compose. Необходимы `env` переменные в файле `.env`.

`make create-migrate` - создает новый файл миграции. Перед использованием нужно изменить название новой миграции.

`make migrate` - запускает миграцию.

`make build` - собирает образ приложения.

`make run` - запускает приложение и базу данных через docker-compose.

`make shutdown` - останавливает выполнение приложения и базы данных. Так же через docker-compose

Будете деплоить у себя, передавайте их при запуске приложения в файле `.env`

