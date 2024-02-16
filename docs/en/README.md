# Gedeon Common Module

This repository is a Go module where the basic logic of Gedeon resides. It can be utilized in any Gedeon service.

## Models

Meet the inhabitants of this module: models represent the essence of data in Gedeon, ensuring data consistency across all services.

### User (User 👤)

This model embodies users, including both their public profile and private data for handling metrics and internal service information.

- `ID`: Unique identifier of the user, represented as a hexadecimal string of a [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) from MongoDB.

```json
    {
        "id": "5f4a5be4076a6055a7a4f7a0"
    }
```

- `Username`: **Unique index** without special characters, with a length between 1 and 32 characters.

- `Email`: **Unique index**, with a length between 4 and 256 characters.

- `DisplayName`: Accepts special characters, with a length between 1 and 32 characters.

- `Bio`: Accepts special characters, with a maximum length of 256 characters.

- `AvatarURL`: Must be a **URL of the [Gedeon CDN](https://en.wikipedia.org/wiki/Content_delivery_network)**.

- `ProfileColor`: Hex code of the profile color.

- `Country`: User's country code, **represented by the [iso3166_1_alpha2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) code**.

- `City`: Maximum length of 256 characters.

- `XP`: No *negative numbers accepted*, with a **minimum of 0**.

- `Badges`: **List of hexadecimal strings** representing the [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) of the user's badges.

- `BirthDate`

- `CreatedDate`

- `UpdatedDate`: Represents the **LAST UPDATED** date of the user.

It also has private attributes, which serve for internal user management, here's an example:

```go
    type User {
        // ...
        // These attributes are private.
        password  string
        ipAddress string
    }
```

> [!IMPORTANT]
> These private attributes are handled exclusively internally and should not be provided in any interface.


#### How to handle private attributes?

To handle private attributes, you can utilize the methods of the structure, here's an example:

```go
    if err := myUser.SetPassword(myHashedPassword); err != nil {
        fmt.Println(err)
    }
```

> [!IMPORTANT]
> Before using `user.SetPassword()`, you must **ENCRYPT** your password to a hash.

After creating a user instance, it's recommended to use the `Validate` method to validate the fields, here's an example:

```go
    myUser := &myPackage.User{
        //...
    }

    if err := myUser.Validate(); err != nil {
        fmt.Println(err)
    }
```


### Badge

- `ID`: Unique identifier of the badge, represented as a hexadecimal string of an [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/).

- `IconCode`: Gedeon code of the icon, for example: `GD_ICON_CODE`.

- `IconUrl`: Icon URL respecting the icon path, for example: `cdn.g3deon.com/image/icon/gd_icon.svg`

- `UserCount`: Number of users with the icon.

- `CreatedDate`

> [!NOTE]
> Before using `user.SetPassword()`, you must **ENCRYPT** your password to a hash.


### Methods

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
> After creating an instance of a badge, you should use the above method, `Validate()`, to validate the fields.

## Utilities

### Json Web Token (JWT)

This package helps us create tokens (JWT) more easily.

Before using the functions, you should be aware of the errors that **can be returned** by the functions; these errors are known scenarios, we may not always know what errors might occur, here's an example:

```go
    // [ Static errors ]

    var (
        // ErrInvalidSecret, returned when the secret key used for signing the JSON WEB token was not the one used.
	    ErrInvalidSecret = errors.New("the secret key was not the one used to sign")
    
        // ErrInvalidToken, returned when the generated token is not valid, *most likely due to a client error*.
	    ErrInvalidToken = errors.New("token is not valid")
    )
```
We may not always know what the error will be, there are cases where the error is unknown *"It's better to be safe than sorry"*. We'll handle the unknown error 
the same way we handle the known one.

```go
    _, err := jwt.Tokenize(data, "short", expire)
    if err != nil {
        if jwt.isShortSecretError(err) {
            fmt.Println(err)
        }

        // Handling unknown error.
        panic(err)
    }
```
> [!NOTE] 
> **DO NOT** use `panic` to handle an error, this is just an example.

### Hash Encryption

This package is an easy way to encrypt information, it uses `bcrypt` underneath and provides functions with the costs to be used in Gedeon for optimization and service performance.

> [!NOTE] 
> This package is used for sensitive data such as passwords, so the information to be encrypted has a minimum length of 8 characters and a maximum of 50. (Subject to change)

Known errors that the package may return:

```go
    var (
        // ErrSameHashAndPassword, returned when the password and provided hash are the same.
        ErrSameHashAndPassword = errors.New("password and the hashed password cannot be the same")
	    
        // ErrEmptyPassword, returned when the password is nil or the string is empty ("").
        ErrEmptyPassword       = errors.New("hashed password or password must be not empty")
    )
```