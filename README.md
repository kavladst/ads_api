### Запуск

```
docker-compose -f docker-compose.yml build
docker-compose -f docker-compose.yml up
```

#### Запуск тестов

```
docker-compose -f docker-compose.test.yml build
docker-compose -f docker-compose.test.yml up
```

### Использование

При вводе неправильных параметров API это обработает

```bash
curl --location --request GET 'http://localhost:8000/v1/ads/detail/?id=car'
```

```bash
{
    "error": "id must be UUID",
    "response": null
}
```

##### Создание объявления
Параметры:
* name - название объявления
* description - описание объявления
* price - цена
* photos_urls - URL на одно фото
```bash
curl --location --request POST 'http://localhost:8000/v1/ads/' \
--form 'name="car"' \
--form 'description="it'\''s Nissan"' \
--form 'price="10000"' \
--form 'photos_urls="url_nissan_1"' \
--form 'photos_urls="url_nissan_2"'
```
```bash
{
    "result": {
        "id": "user_id_1",
        "status": "ok"
    }
}
```

#### Получение объявления
Параметры:
* id - ID объявления
* fields - дополнительные параметры для вывода. Могут быть *urls* (вывод всех картинок)
  или *description* (вывод описания)
```bash
curl --location --request GET 'http://localhost:8000/v1/ads/detail/?id=459ef0a1-51de-451c-a2cb-b7a7f21441c2&fields=urls&fields=description'
```
```bash
{
    {
    "error": "",
    "response": {
        "description": "it's Nissan",
        "name": "car",
        "price": 10000,
        "urls": [
            "url_nissan_1",
            "url_nissan_2"
        ]
    }
}
}
```
* * *
```bash
curl --location --request GET 'http://localhost:8000/v1/ads/detail/?id=459ef0a1-51de-451c-a2cb-b7a7f21441c2'
```
```bash
{
    "error": "",
    "response": {
        "name": "car",
        "price": 10000,
        "url": "url_nissan_1"
    }
}
```

#### Получения объявлений
Параметры:
 * page - номер страницы (по умолчанию page = 1)
 * per_page - число объявлений на странице (по умолчанию per_page = 10)
 * sort_by - поля для сортировки объявлений. Или *price* или *created_at* (по умолчанию sort_by = created_at)
 * order - сортировка по возрастанию/убыванию. Или *asc* или *desc* (по умолчанию order = desc)
```bash
curl --location --request GET 'http://localhost:8000/v1/ads/?per_page=1'
```

```bash
{
    "error": "",
    "response": [
        {
            "id": "5807cce6-7a07-4c44-9307-adf9d37d53f0",
            "image_url": "ad has not image",
            "name": "car",
            "price": 20000
        }
    ]
}
```

```bash
curl --location --request GET 'http://localhost:8000/v1/ads/?page=1&per_page=2&sort_by=price'
```

```bash
{
    "error": "",
    "response": [
        {
            "id": "5807cce6-7a07-4c44-9307-adf9d37d53f0",
            "image_url": "ad has not image",
            "name": "car",
            "price": 20000
        },
        {
            "id": "459ef0a1-51de-451c-a2cb-b7a7f21441c2",
            "image_url": "url_nissan_1",
            "name": "car",
            "price": 10000
        }
    ]
}
```

### База данных
![alt text](./img/img.png)
