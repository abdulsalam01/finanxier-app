package config

import (
	"github.com/finanxier-app/internal/constant"
	"github.com/finanxier-app/internal/entity"
	"github.com/finanxier-app/pkg/database"
	"github.com/finanxier-app/pkg/redis"
)

type Config struct {
	App               `json:"app" yaml:"app"`
	database.Database `json:"database" yaml:"database"`
	redis.MemoryCache `json:"redis" yaml:"redis"`
	Service           `json:"service" yaml:"services"`
	JWTConfig         `json:"jwt" yaml:"jwt"`
}

type App struct {
	Name        string       `json:"name" yaml:"name"`
	Port        string       `json:"port" yaml:"port"`
	Environment constant.Env `json:"env" yaml:"env"`
}

type JWTConfig struct {
	SecretKey string `json:"secret_key" yaml:"secret_key"`
}

type Service struct {
	Auth authService `json:"oauth" yaml:"oauth"`
}

type authService struct {
	Google   entity.GoogleAuth   `json:"google" yaml:"google"`
	Facebook entity.FacebookAuth `json:"facebook" yaml:"facebook"`
}

// To determine testing and actual environment, setup proper env to do a RnD and enable after-research-implementer.
// Development mode.
func (c *Config) IsDevelopmentMode() bool {
	return c.App.Environment == constant.EnvDevelopment
}

// Staging mode.
func (c *Config) IsStagingMode() bool {
	return c.App.Environment == constant.EnvStaging
}
