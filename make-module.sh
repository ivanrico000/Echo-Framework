#!/bin/bash

MODULE_NAME="$1"

if [ -z "$MODULE_NAME" ]; then
  echo "Debes especificar un nombre de módulo. Ejemplo:"
  echo "./make-module.sh rooms"
  exit 1
fi

BASE_DIR="internal/modules/$MODULE_NAME"

# Crear estructura
mkdir -p "$BASE_DIR/core"
mkdir -p "$BASE_DIR/service"
mkdir -p "$BASE_DIR/web"

# Crear entity.go
cat <<EOF > "$BASE_DIR/core/entity.go"
package core

type ${MODULE_NAME^} struct {
    ID string
}
EOF

# Crear repository.go
cat <<EOF > "$BASE_DIR/core/repository.go"
package core

type ${MODULE_NAME^}Repository interface {
    Save(entity *${MODULE_NAME^}) error
    FindByID(id string) (*${MODULE_NAME^}, error)
}
EOF

# Crear dto.go
cat <<EOF > "$BASE_DIR/service/dto.go"
package service

type ${MODULE_NAME^}DTO struct {
    ID string
}
EOF

# Crear usecases.go
cat <<EOF > "$BASE_DIR/service/usecases.go"
package service

import "internal/modules/$MODULE_NAME/core"

type ${MODULE_NAME^}Service struct {
    Repo core.${MODULE_NAME^}Repository
}

func New${MODULE_NAME^}Service(repo core.${MODULE_NAME^}Repository) *${MODULE_NAME^}Service {
    return &${MODULE_NAME^}Service{Repo: repo}
}
EOF

# Crear handler.go
cat <<EOF > "$BASE_DIR/web/handler.go"
package web

import "github.com/labstack/echo/v4"

type ${MODULE_NAME^}Handler struct{}

func New${MODULE_NAME^}Handler() *${MODULE_NAME^}Handler {
    return &${MODULE_NAME^}Handler{}
}

func (h *${MODULE_NAME^}Handler) Get(c echo.Context) error {
    return c.JSON(200, map[string]string{"module": "$MODULE_NAME"})
}
EOF

# Crear routes.go
cat <<EOF > "$BASE_DIR/web/routes.go"
package web

import "github.com/labstack/echo/v4"

func Register${MODULE_NAME^}Routes(g *echo.Group) {
    handler := New${MODULE_NAME^}Handler()

    g.GET("/$MODULE_NAME", handler.Get)
}
EOF

echo "Módulo '$MODULE_NAME' creado correctamente con packages incluidos."
