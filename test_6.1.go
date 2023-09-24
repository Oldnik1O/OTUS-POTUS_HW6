// Тест TestMovableAdapter проверяет, правильно ли адаптер делегирует вызовы методов своему внутреннему объекту
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

// Тестирование адаптера для интерфейса IMovable
func TestMovableAdapter(t *testing.T) {
  ioc := NewIoC()
  // Тип который реализует IMovable регистрации в IoC
  ioc.Register("IMovable", &MovableConcreteType{Position: Vector{1, 2, 3}})

  obj, err := ioc.Resolve("IMovable")
  if err != nil {
    t.Fatalf("Resolve failed: %v", err)
  }

  adapter := &MovableAdapter{Obj: obj.(IMovable)}

  pos := adapter.GetPosition()
  if pos.X != 1 || pos.Y != 2 || pos.Z != 3 {
    t.Errorf("GetPosition failed, got: %v", pos)
  }

  newPos := adapter.SetPosition(Vector{4, 5, 6})
  if newPos.X != 4 || newPos.Y != 5 || newPos.Z != 6 {
    t.Errorf("SetPosition failed, got: %v", newPos)
  }
}

// MovableConcreteType тип который реализует IMovable
type MovableConcreteType struct {
  Position Vector
}

func (m *MovableConcreteType) GetPosition() Vector {
  return m.Position
}

func (m *MovableConcreteType) SetPosition(v Vector) Vector {
  m.Position = v
  return m.Position
}

func (m *MovableConcreteType) GetVelocity() Vector {
  // Dummy velocity
  return Vector{0, 0, 0}
}