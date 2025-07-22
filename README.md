Este proyecto es una aplicación CRUD desarrollada en Go, utilizando MySQL como base de datos con XAMPP y Tailwind CSS para el diseño de la interfaz. Permite gestionar registros de personas (crear, leer, actualizar y eliminar) de manera sencilla y eficiente.

## Características

- Backend en Go
- Base de datos MySQL
- Interfaz moderna con Tailwind CSS
- Operaciones CRUD completas

## Instalación

1. Clona el repositorio.
2. Configura la base de datos MySQL en XAMPP y actualiza las credenciales en `main.go`.
3. Instala las dependencias de Go:
   ```
   go get -u github.com/go-sql-driver/mysql
   ```
4. Compila y ejecuta el servidor:
   ```
   go run main.go
   ```
5. Asegúrate de que Tailwind esté generando el CSS en `static/css/output.css`.

## Uso

- Accede a la aplicación en [http://localhost:8080](http://localhost:8080)
- Utiliza los botones para agregar, editar o eliminar personas.

## Información académica

**Proyecto realizado para la materia Paradigmas de Programación I**  
**Sección 1301, 3er semestre, IUTEPAL**  
**TSU en Informática**

---
Desarrollado por Rainer
