# Tutorial GQLGEN - GORM - MySQL
GraphQL api, con gqlgen y gorm

## Tabla de contenidos

* [Información general](#información-general)
* [Tecnologías](#tecnologías)


## Información general

El proyecto consiste en un API Graphql para guardar, consultar, modificar y eliminar datos sobre usuarios y posts. Estableceremos una relación "one to many" de forma que un usuario
puede tener muchos posts, y un post pertenece a un usuario. Generaremos las queries de forma que al consultar los usuarios podamos acceder a los posts que tienen y, al consultar los
posts poder ver los usuarios al que pertenecen. Las mutations nos permitirán añadir usuarios y posts, actualizarlos y eliminarlos.


## Tecnologías

* [GQLGEN](https://gqlgen.com/)
* [GORM](https://gorm.io/)

## Authors

- [AlegreCode](https://github.com/AlegreCode)


## License

[MIT](https://choosealicense.com/licenses/mit/)
