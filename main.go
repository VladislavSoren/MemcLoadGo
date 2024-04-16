package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
	"time"
	"encoding/json"
	"github.com/bradfitz/gomemcache/memcache"
)


// AppsInstalled представляет установленные приложения и их атрибуты.
type AppsInstalled struct {
	DevType string
	DevID   string
	Lat     float64
	Lon     float64
	Apps    []int
}

// parseAppsInstalled парсит строку и возвращает структуру AppsInstalled.
func parseAppsInstalled(line string) *AppsInstalled {
	lineParts := strings.Split(strings.TrimSpace(line), "\t")
	if len(lineParts) < 5 {
		return nil
	}
	devType, devID, latStr, lonStr, rawApps := lineParts[0], lineParts[1], lineParts[2], lineParts[3], lineParts[4]
	if devType == "" || devID == "" {
		return nil
	}

	apps := make([]int, 0)
	for _, appStr := range strings.Split(rawApps, ",") {
		appStr = strings.TrimSpace(appStr)
		if app, err := strconv.Atoi(appStr); err == nil {
			apps = append(apps, app)
		} else {
			log.Printf("Not all user apps are digits: `%s`\n", line)
		}
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		log.Printf("Invalid geo coords: `%s`\n", line)
		return nil
	}
	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		log.Printf("Invalid geo coords: `%s`\n", line)
		return nil
	}

	return &AppsInstalled{
		DevType: devType,
		DevID:   devID,
		Lat:     lat,
		Lon:     lon,
		Apps:    apps,
	}
}


// Set in memcache
func setInMemcache(appsInstalled AppsInstalled, clientsMemc map[string]*memcache.Client) error {
	key := appsInstalled.DevType + ":" + appsInstalled.DevID
	clientMemc := clientsMemc[appsInstalled.DevType]

	// Сериализация среза в строку байтов
	data, err := json.Marshal(appsInstalled.Apps)
	if err != nil {
		log.Fatalf("Ошибка сериализации: %v", err)
		return nil
	}

	// Сохранение данных в Memcached
	err = clientMemc.Set(&memcache.Item{Key: key, Value: data})
	if err != nil {
		log.Printf("Ошибка сохранения данных в Memcached: %v", err)
		return err
	}
	fmt.Printf("Сохранено в Memcached: ключ=%s, значение=%s\n", key, data)
	
	return nil
}

func main() {

	const maxAttempts = 3
	const retryDelay = 1 * time.Second

	clientsMemc := map[string]*memcache.Client{
		"idfa": memcache.New("localhost:11211"),
		"gaid": memcache.New("localhost:11212"),
		"adid": memcache.New("localhost:11213"),
		"dvid": memcache.New("localhost:11214"),
	}

	// Путь к файлу для чтения данных
	filePath := "/home/soren/Projects/MemcLoadGo/data/appsinstalled/20170929000000_100.tsv.gz"

	// Открытие сжатого файла
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Ошибка открытия файла: %v", err)
	}
	defer file.Close()

	var reader *bufio.Scanner

	// Если файл с расширением .gz, создаем новый сканер для чтения сжатого содержимого
	if strings.HasSuffix(filePath, ".gz") {
		gzipReader, err := gzip.NewReader(file)
		if err != nil {
			log.Fatalf("Ошибка создания Gzip reader: %v", err)
		}
		defer gzipReader.Close()
		reader = bufio.NewScanner(gzipReader)
	} else {
		// Если не сжатый файл, создаем сканер для чтения файла
		reader = bufio.NewScanner(file)
	}

	// Итерация по строкам файла
	for reader.Scan() {
		line := reader.Text()

		// Line parsing
		appsInstalled := parseAppsInstalled(line)
		if appsInstalled != nil {
			fmt.Printf("Parsed AppsInstalled: %+v\n", *appsInstalled)
		} else {
			fmt.Println("Failed to parse AppsInstalled")
		}


		for attempt := 1; attempt <= maxAttempts; attempt++ {
			// Set data in memcache
			err := setInMemcache(*appsInstalled, clientsMemc)
			if err == nil {
				break // Если операция завершилась успешно, выходим из цикла
			}
			log.Printf("Ошибка при выполнении операции с Memcached (попытка %d): %v", attempt, err)
			time.Sleep(retryDelay) // Подождать перед повторной попыткой
		}
	}


	if err := reader.Err(); err != nil {
		log.Fatalf("Ошибка чтения файла: %v", err)
	}
}
