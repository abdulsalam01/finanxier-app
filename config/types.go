package config

import (
	"github.com/api-sekejap/internal/constant"
	"github.com/api-sekejap/internal/entity"
)

type Config struct {
	App      `json:"app" yaml:"app"`
	Database `json:"database" yaml:"database"`
	Service  `json:"service" yaml:"services"`
}

type App struct {
	Name        string       `json:"name" yaml:"name"`
	Port        string       `json:"port" yaml:"port"`
	Environment constant.Env `json:"env" yaml:"env"`
}

type Database struct {
	// Generic info.
	Driver string `json:"driver" yaml:"driver"`
	Port   string `json:"port" yaml:"port"`
	Host   string `json:"host" yaml:"host"`

	// Data info.
	Name     string                 `json:"name" yaml:"name"`
	Username string                 `json:"user" yaml:"user"`
	Password string                 `json:"password" yaml:"password"`
	Extras   map[string]interface{} `json:"extras" yaml:"extras"`
}

type Service struct {
	Storage storageService `json:"storage" yaml:"storage"`
	Auth    authService    `json:"oauth" yaml:"oauth"`
}

type storageService struct {
	Firebase entity.FirebaseStorage
}

type authService struct {
	Google   entity.GoogleAuth   `json:"google" yaml:"google"`
	Facebook entity.FacebookAuth `json:"facebook" yaml:"facebook"`
}

// To determine testing and actual environment, setup proper env to do a RnD and enable after-research-implementer.
// Development mode.
func (c *Config) IsDevelopmentMode() bool {
	if c.App.Environment == constant.EnvDevelopment {
		return true
	}

	return false
}

// Staging mode.
func (c *Config) IsStagingMode() bool {
	if c.App.Environment == constant.EnvStaging {
		return true
	}

	return false
}
