## `flavorful-cli` in `golang`
A simple command line tool implemented with `golang` with the following arguments:
- `flavorful --name` - displays the username of the current session
   - e.g: `echo "$USER"`
- `flavorful --time` - displays the current system time
   - e.g: `date '+%A %W %Y %X'`
- `flavorful --location` - gets and displays the user's current GPS coordinates
   - e.g: via [`curl ipinfo`](https://www.howtogeek.com/405088/how-to-get-your-systems-geographic-location-from-a-bash-script/)
- `flavorful --fortune-cookie` - gets and displays a random piece of fortune
   - [`fortune`](https://wiki.archlinux.org/index.php/Fortune)
- `flavorful --weather --today` - gets and displays today's weather
   - [OpenWeather API](https://openweathermap.org/current)
- `flavorful --weather --tomorrow` - gets and displays tomorrow's weather
   - [OpenWeather API](https://openweathermap.org/current)

## Organization
The `flavorful-cli` go package is organized based on [`How to write Go code`](https://golang.org/doc/code.html#Organization).

## Install and Build
- Install `cobra`
```bash
cd .. && go get -u github.com/spf13/cobra/cobra
```

- Install `flavorful-cli`
```bash
cd flavorful-cli/
go install github.com/amandazhuyilan/flavorful-cli
```

- Build `flavorful-cli`
```bash
cd flavorful-cli/
go build
```
A `flavorful-cli` binary should be built in the current directory.

## Run
```bash
./flavorful-cli
```

## Inital Set up Instructions
1. Initialize go module:
```bash
go mod init github.com/amandazhuyilan/flavorful-cli
```

2. Install and initialize cobra module
```bash
cd .. && go get -u github.com/spf13/cobra/cobra
cobra init --pkg-name github.com/amandazhuyilan/flavorful-cli
```

3. Initialize basic cli arguments with [`cobra` generator](https://github.com/spf13/cobra/blob/master/cobra/README.md#cobra-add)
```bash
cobra add name
cobra add time
cobra add location
cobra add fortune-cookie
cobra add weather
cobra add create -p 'configCmd'
```

These commands will generate corresponding `*.go` files under `cmd/`.