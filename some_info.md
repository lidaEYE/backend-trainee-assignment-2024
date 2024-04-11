Примеры запросов и ответов:
___1. Получение баннера для пользователя (/user_banner):___
  Метод: GET
  
  Параметры запроса:
  
    tag_id (идентификатор тега пользователя)
    
    feature_id (идентификатор функции)
    
    use_last_revision (использовать последнюю ревизию) - опционально
    
  Заголовок запроса:
  
    token (токен аутентификации администратора)
   
  Пример запроса:
```bash
curl -X GET \
  'http://localhost:8080/user_banner?tag_id=123&feature_id=456&use_last_revision=true' \
  -H 'token: your_admin_token'
```

  Пример успешного ответа:
```json
{
  "banner_id": 1,
  "tag_ids": [123],
  "feature_id": 456,
  "content": {
    "title": "Welcome!",
    "message": "Enjoy our latest feature!"
  },
  "is_active": true,
  "created_at": "2024-04-10T12:00:00Z",
  "updated_at": "2024-04-10T12:05:00Z"
}
```

  Пример ответа при возникновении ошибки:
```json
{
  "error": "Failed to get user banner"
}
```

___2. Получение списка баннеров (/banner):___

  Метод: GET
  
  Параметры запроса:
  
    feature_id (идентификатор функции) - опционально
    
    tag_id (идентификатор тега) - опционально
    
    limit (максимальное количество баннеров) - опционально
    
    offset (смещение для пагинации) - опционально
    
  Заголовок запроса:
  
    token (токен аутентификации администратора)
    
  Пример запроса:
```bash
curl -X GET \
  'http://localhost:8080/banner?feature_id=456&tag_id=123&limit=10&offset=0' \
  -H 'token: your_admin_token'
```

  Пример успешного ответа:
```json
[
  {
    "banner_id": 1,
    "tag_ids": [123],
    "feature_id": 456,
    "content": {
      "title": "New Feature!",
      "message": "Check out our latest update!"
    },
    "is_active": true,
    "created_at": "2024-04-10T12:00:00Z",
    "updated_at": "2024-04-10T12:05:00Z"
  },
  {
    "banner_id": 2,
    "tag_ids": [123, 456],
    "feature_id": 789,
    "content": {
      "title": "Special Offer!",
      "message": "Limited time discount!"
    },
    "is_active": true,
    "created_at": "2024-04-10T12:10:00Z",
    "updated_at": "2024-04-10T12:15:00Z"
  }
]
```

  Пример ответа при возникновении ошибки:
```json
{
  "error": "Failed to get banners"
}
```

