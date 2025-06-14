# to-do

- Mejorar el parser para que no utilice interface de puente, esto rompe el tipado y hay que usar assert para recuperar

- Mejorar el actualizar usuario para que no tenga que solicitar el usuario sino que condicionalmente actulice solo los campos necesarios

- Mejorar el crear usuario para no validar en solicitudes independientes si existe o no usuario y email

- acomodar el parser para no tener que tener separado el paginador

- Me gusto la estructura del parser documentar para usar como core utils o replicar en los otros paquetes

- getuserby no trae el status

- auth no se ha mirado desde el inicio del proyecto asi que hay que actualizar el formato

- agregar comentarios en las funciones claves

- revisar los triggers de la base de datos

- chequear en que casos enviar datos de usuario de creacion, actualizacion

- implementar la funcion de los repo a la linea de comando

- empezar a documentar

- Mejorar el bulk delete, agregar un repo specifico para eso que maneje transacciones y acepte un arreglo en lugar de hacerlo secuencialmente (es posible que tenga que soltarse sqlc para esto)
# users
