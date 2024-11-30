## 1. DESCRIPTION

`go-config-extender` is a utility that helps to work with config files: it loads a JSON config file and ***overrides values by setting environment variables***. This is useful when we are having, i.e., `Dockerfile` or `docker-compose`, consider that we are having application 
config file `config.json` and we have to override log level or db `username+passord` based on values in an `.env` file.

## 2. Requirements

To make this thing work, we have to do the following:
1. We have to separate *"normal"* environment variables from technical (those using for such override);
2. We should have some relation between `JSON` properties in config and env variable.

Considering these two restrictions, we decided to have a variable's name pattern that allows it to meet all these requirements.

consider following example of config:
```json
{
    "server": {
        "schema": "http",
        "address": "0.0.0.0",
        "port": 8182,
        "secret_file": "./keyfile"
    },
    "logging": {
        "level": "debug",
        "appenders": [
            {
                "type": "rolling_file",
                "enabled": true,
                "level": "debug",
                "destination": {
                    "file": "./logs/ferrum.log",
                    "max_size": 100,
                    "max_age": 5,
                    "max_backups": 5,
                    "local_time": true
                }
            },
            {
                "type": "console",
                "enabled": true,
                "level": "debug"
            }
        ],
        "http_log": true,
        "http_console_out": true
    },
    "data_source": {
        "type": "redis",
        "source": "redis:6379",
        "credentials": {
            "username": "ferrum_db",
            "password": "FeRRuM000"
        },
        "options": {
            "namespace": "ferrum_1",
            "db_number": "0"
        }
    }
}
```

And we would like to override:
1. `logging.level` -> `env` variable `__logging.level`
2. `data_source.username` -> `env` variable `__data_source.username`
3. `data_source.password` -> `env` variable `__data_source.password`

What methods this package have:
1. `LoadConfig [T any] (configFile string) (T, error)` for loading without overriding
2. `LoadConfigWithEnv [T any] (configFile string) (T, error)` for loading with override

## 3. Full example

A full example could be found in the test `TestLoadJSONConfigWithEnvOverride` [test_file](./config_loader_test.go)
In This test were set 3 variables of different types: `bool`, `int` and `float`

Real App usage will be shown later after this package Release && usage in `Ferrum Community Authorization Server`