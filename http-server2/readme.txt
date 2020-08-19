ДЗ :
Создать клиента, который будет отправлять запросы на ваш сервер, на эндпоинты:
Получение, создание, вашего воркера , с помощью функций пакета "net/http" :
func Post(url string, contentType string, body io.Reader) (resp *Response, err error)
func PostForm(url string, data url.Values) (resp *Response, err error)

Посмотреть реализацию этих функций в пакете.