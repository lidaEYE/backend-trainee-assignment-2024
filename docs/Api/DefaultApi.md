# Swagger\Client\DefaultApi

All URIs are relative to */*

Method | HTTP request | Description
------------- | ------------- | -------------
[**bannerGet**](DefaultApi.md#bannerget) | **GET** /banner | Получение всех баннеров c фильтрацией по фиче и/или тегу
[**bannerIdDelete**](DefaultApi.md#banneriddelete) | **DELETE** /banner/{id} | Удаление баннера по идентификатору
[**bannerIdPatch**](DefaultApi.md#banneridpatch) | **PATCH** /banner/{id} | Обновление содержимого баннера
[**bannerPost**](DefaultApi.md#bannerpost) | **POST** /banner | Создание нового баннера
[**userBannerGet**](DefaultApi.md#userbannerget) | **GET** /user_banner | Получение баннера для пользователя

# **bannerGet**
> \Swagger\Client\Model\InlineResponse200[] bannerGet($token, $feature_id, $tag_id, $limit, $offset)

Получение всех баннеров c фильтрацией по фиче и/или тегу

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$token = "token_example"; // string | Токен админа
$feature_id = 56; // int | 
$tag_id = 56; // int | 
$limit = 56; // int | 
$offset = 56; // int | 

try {
    $result = $apiInstance->bannerGet($token, $feature_id, $tag_id, $limit, $offset);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->bannerGet: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string**| Токен админа | [optional]
 **feature_id** | **int**|  | [optional]
 **tag_id** | **int**|  | [optional]
 **limit** | **int**|  | [optional]
 **offset** | **int**|  | [optional]

### Return type

[**\Swagger\Client\Model\InlineResponse200[]**](../Model/InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **bannerIdDelete**
> bannerIdDelete($id, $token)

Удаление баннера по идентификатору

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$id = 56; // int | 
$token = "token_example"; // string | Токен админа

try {
    $apiInstance->bannerIdDelete($id, $token);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->bannerIdDelete: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **int**|  |
 **token** | **string**| Токен админа | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **bannerIdPatch**
> bannerIdPatch($body, $id, $token)

Обновление содержимого баннера

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$body = new \Swagger\Client\Model\BannerIdBody(); // \Swagger\Client\Model\BannerIdBody | 
$id = 56; // int | 
$token = "token_example"; // string | Токен админа

try {
    $apiInstance->bannerIdPatch($body, $id, $token);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->bannerIdPatch: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\Swagger\Client\Model\BannerIdBody**](../Model/BannerIdBody.md)|  |
 **id** | **int**|  |
 **token** | **string**| Токен админа | [optional]

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **bannerPost**
> \Swagger\Client\Model\InlineResponse201 bannerPost($body, $token)

Создание нового баннера

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$body = new \Swagger\Client\Model\BannerBody(); // \Swagger\Client\Model\BannerBody | 
$token = "token_example"; // string | Токен админа

try {
    $result = $apiInstance->bannerPost($body, $token);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->bannerPost: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**\Swagger\Client\Model\BannerBody**](../Model/BannerBody.md)|  |
 **token** | **string**| Токен админа | [optional]

### Return type

[**\Swagger\Client\Model\InlineResponse201**](../Model/InlineResponse201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

# **userBannerGet**
> map[string,object] userBannerGet($tag_id, $feature_id, $use_last_revision, $token)

Получение баннера для пользователя

### Example
```php
<?php
require_once(__DIR__ . '/vendor/autoload.php');

$apiInstance = new Swagger\Client\Api\DefaultApi(
    // If you want use custom http client, pass your client which implements `GuzzleHttp\ClientInterface`.
    // This is optional, `GuzzleHttp\Client` will be used as default.
    new GuzzleHttp\Client()
);
$tag_id = 56; // int | 
$feature_id = 56; // int | 
$use_last_revision = false; // bool | 
$token = "token_example"; // string | Токен пользователя

try {
    $result = $apiInstance->userBannerGet($tag_id, $feature_id, $use_last_revision, $token);
    print_r($result);
} catch (Exception $e) {
    echo 'Exception when calling DefaultApi->userBannerGet: ', $e->getMessage(), PHP_EOL;
}
?>
```

### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tag_id** | **int**|  |
 **feature_id** | **int**|  |
 **use_last_revision** | **bool**|  | [optional] [default to false]
 **token** | **string**| Токен пользователя | [optional]

### Return type

**map[string,object]**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../../README.md#documentation-for-api-endpoints) [[Back to Model list]](../../README.md#documentation-for-models) [[Back to README]](../../README.md)

