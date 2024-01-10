package fmemory

import (
  "github.com/gofiber/storage/memory/v2"
)

// Config defines the config for storage.
type Config = memory.Config

func New(config ...Config) (memoryStorage *memory.Storage) {
  return memory.New(config...)
}
