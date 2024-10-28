API hecha con go y el framework gin <br>
Instrucciones para ejecutar la API

1.- descargar e instalar el compilador de go (version 1.21.0)<br>
2.- situarse en la carpeta del proyecto y ejecutar "go mod download" para descargar las dependencias del proyecto<br>
3.- configurar las credenciales de la base de datos local de postgreSQL para que la API se pueda conectar<br>
4.- levantar el servidor local con el comando "go run main.go"

Se crearon 5 endpoints

1 endpoint de tipo GET("/personas") que trae las personas con su nombre, email y edad
éste endpoint trae los registros paginados de 15 en 15, recibe un parámetro que se llama "page" para especificar el numeró de página
que se desea consultar, si no se proporciona por default te devuelve la página 1

1 endpoint de tipo POST("/personas") para crear una nueva persona, se manda un objeto JSON en el body
con los 3 datos de la persona, en la API se llaman "name", "email" y "age"

1 endpoint de tipo PUT("/personas/:ID") para editar una persona existente donde "ID" es un número, se manda un objeto JSON en el body
con los 3 datos de la persona ("name", "email" y "age"), se valida en la API que  se actualicen solo los datos que cambiaron, 
te devuelve el objeto actualizado

1 endpoint de tipo DELETE("/personas/:ID") para eliminar la persona con el ID que se mande como parámetro, te devuelve el objeto eliminado

1 endpoint de tipo GET("/personas/search") recibe un parámetro llamado "query" para buscar una persona por nombre o email, 
te devuelve una lista con las personas con ese nombre o email, si no se encuentra nada te devuelve una lista vacía
