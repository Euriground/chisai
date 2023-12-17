package config

import (
    "encoding/json"
    "os"
    "fmt"
)

type Config struct {
    Server struct {
        Port int `json:"port"`
    } `json:"server"`
    Database struct {
        ConnectionString string `json:"connectionString"`
    } `json:"database"`
}

func LoadConfig() (*Config, error) {
    configFile, err := os.Open("config.json")
    if err != nil {
        return nil, fmt.Errorf("failed to open config file: %s", err)
    }
    defer configFile.Close()

    var config Config
    err = json.NewDecoder(configFile).Decode(&config)
    if err != nil {
        return nil, fmt.Errorf("failed to decode config file: %s", err)
    }

    return &config, nil
}
