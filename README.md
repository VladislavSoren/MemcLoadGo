## MemcLoadGo

### Цель скрипта:
Заливает в мемкеш выгрузку логов трекера установленных приложений. 
Ключом является тип и идентификатор устройства через двоеточие, значением являет перечень кодов приложений установленных на устройство.

### Основная функциональность:
1. `main.go` - многопроцессная загрузка чанками

### Мониторинг:
1. Логи сновной программы пишутся в `app.log`

### Инструкции по запуску: 
1. `cd <Абсолютный путь к директории проекта>` 
3. `go run main.go`
