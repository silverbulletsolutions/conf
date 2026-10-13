# conf

Biblioteka w języku Go przeznaczona do prostego i elastycznego zarządzania konfiguracją aplikacji. Zapewnia ujednolicony interfejs (`Configuration`) oraz implementacje pozwalające na odczyt parametrów konfiguracyjnych ze zmiennych środowiskowych, plików lub mapy w pamięci (przydatne np. w testach).

## Funkcjonalności

- **Pobieranie wartości różnych typów**: obsługa ciągów znaków (`string`), tablic ciągów (`[]string`), liczb całkowitych (`int`) oraz czasu trwania (`time.Duration`).
- **Obsługa wartości domyślnych**: bezpieczne pobieranie wartości z opcją zwrócenia zdefiniowanej wartości domyślnej w przypadku braku klucza lub błędów parsowania.
- **Wymagalność parametrów**: metoda `GetStringRequired` rzuca `panic`, jeśli klucz nie został podany.
- **Odczyt z plików / sekretów (`GetStringWithFile`)**: możliwość odczytu wartości ze wskazanego pliku (poprzez zmienną `<ID>_file`), a w przypadku jego braku – bezpośrednio ze zmiennej środowiskowej `<ID>`.
- **Sekwencje/tablice wartości (`GetStrings`)**: obsługa indeksowanych zmiennych w formacie `<ID>__<IDX>` (np. `SERVERS__0`, `SERVERS__1`).
- **Łatwe mockowanie/testowanie**: implementacja `DefaultConfiguration` pozwala przekazać konfigurację jako `map[string]string`.

## Instalacja

```bash
go get github.com/silverbulletsolutions/conf
```

## Interfejs `Configuration`

```go
type Configuration interface {
    GetStringWithFile(id string, defaultValue string) string
    GetString(id string, defaultValue string) string
    GetStrings(id string) []string
    GetStringRequired(id string) string
    GetInt(id string, defaultValue int) int
    GetDuration(id string, defaultValue int, unit time.Duration) time.Duration
    GetDurationS(id string, defaultValueS int) time.Duration // Deprecated: użyj GetDuration
}
```

## Implementacje

### 1. `EnvironmentConfiguration`

Odczytuje konfigurację ze zmiennych środowiskowych procesu:

- `GetString(id, default)`: pobiera wartość zmiennej środowiskowej `id`. Jeśli pusta/brak – zwraca `default`.
- `GetStringWithFile(id, default)`: sprawdza najpierw zmienną `id + "_file"`. Jeśli istnieje, odczytuje zawartość wskazanego pliku. Jeśli nie, pobiera wartość zmiennej `id`. W przeciwnym razie zwraca `default`.
- `GetStringRequired(id)`: pobiera wartość zmiennej `id`. Rzuca `panic`, jeśli zmienna nie jest ustawiona lub jest pusta.
- `GetInt(id, default)`: pobiera zmienną i parsuje do liczby całkowitej. W przypadku błędu parsowania lub braku zmiennej zwraca `default`.
- `GetDuration(id, default, unit)`: pobiera liczbę i mnoży przez jednostkę `unit` (np. `time.Second`, `time.Minute`).
- `GetStrings(id)`: odczytuje listę wartości ze zmiennych `id__0`, `id__1`, `id__2`, ... aż do pierwszej brakującej lub pustej (maksymalnie 1024 elementy).

### 2. `DefaultConfiguration`

Implementacja oparta na mapie `map[string]string`, idealna do testów jednostkowych lub konfiguracji statycznej.

## Przykłady użycia

### Konfiguracja ze środowiska produkcyjnego

```go
package main

import (
    "fmt"
    "time"

    "github.com/silverbulletsolutions/conf"
)

func main() {
    cfg := conf.NewEnvironmentConfiguration()

    // Odczyt z pliku lub zmiennej środowiskowej (np. hasła/sekrety w Docker/Kubernetes)
    dbPassword := cfg.GetStringWithFile("DB_PASSWORD", "secret")

    // Zwykły ciąg znaków z wartością domyślną
    appEnv := cfg.GetString("APP_ENV", "production")

    // Parametr wymagany
    apiKey := cfg.GetStringRequired("API_KEY")

    // Liczba całkowita
    port := cfg.GetInt("PORT", 8080)

    // Czas trwania
    timeout := cfg.GetDuration("TIMEOUT_SECONDS", 30, time.Second)

    // Lista wartości (np. HOSTS__0="host1", HOSTS__1="host2")
    hosts := cfg.GetStrings("HOSTS")

    fmt.Printf("Env: %s, Port: %d, Timeout: %v, Hosts: %v\n", appEnv, port, timeout, hosts)
}
```

### Konfiguracja w testach jednostkowych

```go
package mytest

import (
    "testing"
    "time"

    "github.com/silverbulletsolutions/conf"
)

func TestService(t *testing.T) {
    mockCfg := conf.NewDefaultConfiguration(map[string]string{
        "APP_ENV":   "test",
        "PORT":      "9000",
        "SERVERS__0": "127.0.0.1",
        "SERVERS__1": "127.0.0.2",
    })

    if mockCfg.GetInt("PORT", 8080) != 9000 {
        t.Errorf("niepoprawny port")
    }
}
```
