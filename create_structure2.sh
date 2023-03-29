#!/bin/bash

# Create directories
mkdir auth config container controllers models public repositories routers services views

# Create files
touch main.go go.mod go.sum config/app.yaml config/database.yaml container/container.go container/routes.go controllers/product_controller.go controllers/user_controller.go models/product.go models/user.go public/css/style.css public/css/bootstrap.min.css public/js/script.js public/js/jquery.min.js public/img/logo.png public/img/background.jpg repositories/product_repository.go repositories/user_repository.go routers/product_routes.go routers/user_routes.go services/product_service.go services/user_service.go views/auth/login.html views/auth/register.html views/layout/base.html views/partials/header.html views/partials/footer.html views/product/create.html views/product/edit.html views/product/index.html views/user/create.html views/user/edit.html views/user/index.html auth/auth_controller.go auth/auth_routes.go auth/auth_service.go auth/jwt_util.go routers/auth_routes.go auth/tests/auth_controller_test.go auth/tests/auth_routes_test.go auth/tests/auth_service_test.go

echo "Basic project structure generated successfully."
