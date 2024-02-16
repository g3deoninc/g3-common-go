# Módulo Común de Gedeón

Este repositorio es un módulo de Go donde se encuentra la lógica básica de Gedeón. Puede ser utilizado en cualquier servicio de Gedeón.

## Modelos

Conoce a los habitantes de este módulo: los modelos representan la esencia de los datos en Gedeón, así tenemos consistencia de datos en todos los servicios.

### User (Usuario 👤) 

Este modelo encarna a los usuarios, tanto su perfil público como datos privados para el manejo de métricas e información interna del servicio.

- `ID`: Identificador único del usuario, representado como una cadena hexadecimal de un [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) de MongoDB.

```json
    {
        "id": "5f4a5be4076a6055a7a4f7a0"
    }
```

- `Username` (Nombre del usuario): **Índice único** sin caracteres especiales, tiene una longitud entre 1 y 32 caracteres.

- `Email` (Correo electrónico): **Índice único**, longitud entre 4 y 256 caracteres.

- `DisplayName` (Nombre a mostrar): Admite caracteres especiales, tiene una longitud entre 1 y 32 caracteres.
  
- `Bio` (Biografía): Admite caracteres especiales, longitud máxima de 256 caracteres.

- `AvatarURL`  (URL del avatar): Debe ser una **URL de la [CDN](https://en.wikipedia.org/wiki/Content_delivery_network) de Gedeón**.

- `ProfileColor` (Color del perfil): Código Hex del color del perfil.
  
- `Country` (País): Código del país del usuario, **representado por el código [iso3166_1_alpha2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)**.
  
- `City` (Ciudad): Longitud máxima de 256 caracteres.
  
- `XP` (Experiencia): No *admite números negativos*, el **mínimo es 0**.

- `Badges` (Medallas): **Lista de cadenas hexadecimales** que representa el [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) de las medallas del usuario.

- `BirthDate` (Fecha de nacimiento)

- `CreatedDate` (Fecha de creación)

- `UpdatedDate` (Fecha de actualización): Representa la **ÚLTIMA ACTUALIZACIÓN** del usuario.
  
También tiene atributos privados, estos sirven para el manejo interno del usuario, aquí un ejemplo:

```go

    type User {
        // ...
        // Estos atributos son privados.
        password  string
        ipAddress string
    }
```

> [!IMPORTANT]
> Estos atributos privados se manejan únicamente de manera interna y no se deben proporcionar en ninguna interfaz.


#### ¿Cómo manejar los atributos privados?

Para manejar los atributos privados puedes utilizar los métodos de la estructura, aquí un ejemplo:

```go
    if err := myUser.SetPassword(myHashedPassword); err != nil {
        fmt.Println(err)
    }
```

> [!IMPORTANT]
> Antes de utilizar `user.SetPassword()`, debe **ENCRIPTAR** su contraseña a un hash.

Después crear una instancia de un usuario, se recomienda que se utilice el método `Validate`, para validar los campos, aquí un ejemplo:

```go
    myUser := &myPackage.User{
        //...
    }

    if err := myUser.Validate(); err != nil {
        fmt.Println(err)
    }
```  


### Badge (Medalla)

- `ID`: Identificador único de la medalla, representado como una cadena hexadecimal de un  [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/).
  
- `IconCode` (Código del icono): Código Gedeón del icono ejemplo: `GD_ICON_CODE`.
  
- `IconUrl` (URL de Icono): URL que respeta el ruta del icono, ejemplo:
`cdn.g3deon.com/image/icon/gd_icon.svg`

- `UserCount` (Cuenta de usuarios): Cantidad de usuarios con el icono.

- `CreatedDate` (Fecha de creación)

> [!NOTE]
> Antes de utilizar `user.SetPassword()`, debe **ENCRIPTAR** su contraseña a un hash.


### Métodos

```go
    func (b *Badge) Validate() error {
	    validate := validator.New()

	    if err := validate.RegisterValidation("prefix", validatePrefix); err != nil {
		    return err
	    }

	    return validate.Struct(b)
    }
```
> [!IMPORTANT]
> Después crear una instancia una medalla, debe utilizar el método anterior, `Validate()` para validar los campos.

## Utilidades

### Json Web Token (JWT)

Este paquete nos ayuda a crear tokens (JWT) de manera más fácil.

Antes de utilizar las funciones debe saber los errores que **pueden retornar** las funciones, estos errores son los escenarios conocidos, no siempre sabremos que errores puedan ocurrir, aquí un ejemplo:

```go
    // [ Errores estáticos ]

    var (
        // ErrInvalidSecret, se retorna cuando la clave secreta no fue la utilizada para la firma del JSON WEB token.
	    ErrInvalidSecret = errors.New("the secret key was not the one used to sign")
    
        // ErrInvalidToken, se retorna cuando el token generado no es valido, *muy posiblemente por un error del cliente*.
	    ErrInvalidToken = errors.New("token is not valid")
    )
```
No siempre sabremos cuál será el error, hay casos los cuales el error es desconocido *"Es mejor prevenir que lamentar"*. Manejaremos el error 
desconocido al igual que el conocido.

```go
    _, err := jwt.Tokenize(data, "short", expire)
    if err != nil {
        if jwt.isShortSecretError(err) {
            fmt.Println(err)
        }

        // Manejo de error desconocido.
        panic(err)
    }
```
> [!NOTE] 
> **NO DEBE** utilizar el `panic` a la hora de manejar un error, esto es un ejemplo.

### Encriptar con Hash

Este paquete es una manera fácil de encriptar información, utiliza `bcrypt` por debajo y proporciona funciones con los costos que se deben utilizar en Gedeón para la optimización y rendimiento de los servicios.

> [!NOTE] 
> Este paquete se utiliza para datos vulnerables como contraseñas, por eso la información a encriptar tiene una longitud minima de 2 caracteres y máxima de 256. (Sujeto a cambio)

Errores **conocidos** que puede retornar el paquete:

```go
    var (
        // ErrSameHashAndPassword, retornado cuando la contraseña y el hash proporcionado son iguales.
        ErrSameHashAndPassword = errors.New("password and the hashed password cannot be the same")
	    
        // ErrEmptyPassword, retornado cuando la contraseña es nil o la cadena esta vacía ("").
        ErrEmptyPassword       = errors.New("hashed password or password must be not empty")
    )
```