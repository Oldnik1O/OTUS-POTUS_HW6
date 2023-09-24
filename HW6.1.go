// Реализация IoC контейнера и адаптера для интерфейса IMovable. 
package main

import (
  "errors"
  "fmt"
  "sync"
  "testing"
)

type Vector struct {
  X, Y, Z float64
}

type IMovable interface {
  GetPosition() Vector
  SetPosition(Vector) Vector
  GetVelocity() Vector
}

type MovableAdapter struct {
  Obj IMovable
}

func (m *MovableAdapter) GetPosition() Vector {
  return m.Obj.GetPosition()
}

func (m *MovableAdapter) SetPosition(v Vector) Vector {
  return m.Obj.SetPosition(v)
}

func (m *MovableAdapter) GetVelocity() Vector {
  return m.Obj.GetVelocity()
}

type IoC struct {
  mu       sync.Mutex
  registry map[string]interface{}
}

func NewIoC() *IoC {
  return &IoC{
    registry: make(map[string]interface{}),
  }
}

func (ioc *IoC) Register(key string, instance interface{}) {
  ioc.mu.Lock()
  defer ioc.mu.Unlock()
  ioc.registry[key] = instance
}

func (ioc *IoC) Resolve(key string) (interface{}, error) {
  ioc.mu.Lock()
  defer ioc.mu.Unlock()

  if instance, exists := ioc.registry[key]; exists {
    return instance, nil
  }

  return nil, errors.New("key not found")
}

