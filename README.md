# Flavorful Command Line Interfaces
This repo contains a command line tool, `flavorful-cli`, implemented in different flavors of programming languages. Each flavour of `flavorful-cli` lives in each its individual branch.

### List of CLI tool flavors:
- [x] [golang](https://github.com/amandazhuyilan/flavorful-cli/tree/golang) ([link to study notes](https://github.com/amandazhuyilan/flavorful-cli/blob/main/golang/NOTES.md))
- [ ] Rust
- [ ] C++14
- [ ] Python

### `hello-cli`
A simple command line tool with the following arguments:
- `flavorful-cli name` - displays the username of the current session
   - e.g: `echo "$USER"`
- `flavorful-cli time` - displays the current system time
   - e.g: `date '+%A %W %Y %X'`
- `flavorful-cli ip` - gets and displays the user's current device IP address
   - e.g: via `curl ipinfo`.
- `flavorful-cli weather --city` - gets and displays a provided city's current weather condition
   - [OpenWeather API](https://openweathermap.org/current)

### Acknowledgements

Shout out to Sebastian, whose [100-day-UI-challenge and wonderful UI designs](https://www.instagram.com/p/Bsf6xuwBfJv/), inspired me to kick-start this amazing project.
