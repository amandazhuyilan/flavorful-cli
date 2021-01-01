## `flavorful-cli`
A simple command line tool implemeted in `python3` with the following arguments:
- `flavorful-cli name` - displays the username of the current session
   - e.g: `echo "$USER"`
- `flavorful-cli time` - displays the current system time
   - e.g: `date '+%A %W %Y %X'`
- `flavorful-cli ip` - gets and displays the user's current GPS coordinates
   - e.g: via `curl ipinfo`
- `flavorful-cli weather [city name]` - gets and displays a given a city's current weather
   - [OpenWeather API](https://openweathermap.org/current)

## Package Organization
The `flavorful-cli` go package is organized based on [`How to write Go code`](https://golang.org/doc/code.html#Organization).

### Set up
```bash
source flavorful-cli/bin/activate
(flavorful-cli) pip install -r requirements.txt
```

### Run
```bash
python3 main.py --greeting Hello Ann    # This should print "Hello, Ann!'
```

## Test
```bash
INSTRUCTIONS TO RUN TESTS HERE
```

## Inital Set up Procedures

### `venv`
This projects uses [`venv`|https://docs.python.org/3/library/venv.html], `python3`'s own virtual environment tool for dependency management and isolation, ease of installing and using Python packages without system-administrator access, and automated testing of Python software across multiple Python versions, etc.

To set up `flavorful-cli`'s new virtual environement,
```bash
cd flavorful-cli/
python3 -m venv .
```

To activate virtual environment,
```bash
source flavorful-cli/bin/activate
```

To keep the environment consistent, and “freeze” the current state of the packages,
```bash
(flavorful-cli) pip freeze > requirements.txt
```

To re-create the environment to install the same packages using the same versions:
```bash
(flavorful-cli) pip install -r requirements.txt
```

To deactivate the activated virtual environment,
```bash
deactivate
```

4. Tests
```bash
INSTRUCTIONS TO RUN TESTS HERE
```

5. `git-secret`
The API keys, `open_weather.txt` protected as `open_weather.txt.secret`, used in this project is protected by [`git-secret`](https://git-secret.io/).

To decrypt, do `git secret reveal` or `git secret cat`.