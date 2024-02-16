# Módulo Comum de Gedeón

Este repositório é um módulo em Go onde está localizada a lógica básica do Gedeón. Pode ser utilizado em qualquer serviço do Gedeón.

## Modelos

Conheça os habitantes deste módulo: os modelos representam a essência dos dados no Gedeón, garantindo consistência de dados em todos os serviços.

### User (Usuário 👤)

Este modelo incorpora os usuários, tanto seu perfil público quanto dados privados para o gerenciamento de métricas e informações internas do serviço.

- `ID`: Identificador único do usuário, representado como uma cadeia hexadecimal de um [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/).

```json
    {
        "id": "5f4a5be4076a6055a7a4f7a0"
    }
```

- `Username` (Nome do usuário): **Índice único** sem caracteres especiais, com comprimento entre 1 e 32 caracteres.

- `Email` (E-mail): **Índice único**, com comprimento entre 4 e 256 caracteres.

- `DisplayName` (Nome de exibição): Aceita caracteres especiais, com comprimento entre 1 e 32 caracteres.

- `Bio` (Biografia): Aceita caracteres especiais, com comprimento máximo de 256 caracteres.

- `AvatarURL` (URL do avatar): Deve ser uma **URL do [CDN](https://en.wikipedia.org/wiki/Content_delivery_network) do Gedeón**.

- `ProfileColor` (Cor do perfil): Código Hex da cor do perfil.

- `Country` (País): Código do país do usuário, **representado pelo código [iso3166_1_alpha2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)**.

- `City` (Cidade): Comprimento máximo de 256 caracteres.

- `XP` (Experiência): Não *aceita números negativos*, com o **mínimo sendo 0**.

- `Badges` (Medalhas): **Lista de cadeias hexadecimais** que representam o [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/) das medalhas do usuário.

- `BirthDate` (Data de nascimento)

- `CreatedDate` (Data de criação)

- `UpdatedDate` (Data de atualização): Representa a **ÚLTIMA ATUALIZAÇÃO** do usuário.

Também possui atributos privados, que servem para o gerenciamento interno do usuário, aqui está um exemplo:
****
```go
    type User {
        // ...
        // Estes atributos são privados.
        password  string
        ipAddress string
    }
```

> [!IMPORTANT]
> Estes atributos privados são manipulados exclusivamente internamente e não devem ser fornecidos em nenhuma interface.


#### Como gerenciar os atributos privados?

Para gerenciar os atributos privados, você pode utilizar os métodos da estrutura, aqui está um exemplo:

```go
    if err := myUser.SetPassword(myHashedPassword); err != nil {
        fmt.Println(err)
    }
```

> [!IMPORTANT]
> Antes de utilizar `user.SetPassword()`, você deve **ENCRYPTAR** sua senha para um hash.

Após criar uma instância de um usuário, recomenda-se utilizar o método `Validate`, para validar os campos, aqui está um exemplo:

```go
    myUser := &myPackage.User{
        //...
    }

    if err := myUser.Validate(); err != nil {
        fmt.Println(err)
    }
```


### Badge (Medalha)

- `ID`: Identificador único da medalha, representado como uma cadeia hexadecimal de um [**ObjectId**](https://www.mongodb.com/docs/manual/reference/method/ObjectId/).

- `IconCode` (Código do ícone): Código Gedeón do ícone, exemplo: `GD_ICON_CODE`.

- `IconUrl` (URL do ícone): URL que respeita o caminho do ícone, exemplo: `cdn.g3deon.com/image/icon/gd_icon.svg`

- `UserCount` (Contagem de usuários): Quantidade de usuários com o ícone.

- `CreatedDate` (Data de criação)

> [!NOTE]
> Antes de utilizar `user.SetPassword()`, você deve **ENCRYPTAR** sua senha para um hash.


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
> Após criar uma instância de uma medalha, você deve utilizar o método acima, `Validate()`, para validar os campos.

## Utilitários

### Json Web Token (JWT)

Este pacote nos ajuda a criar tokens (JWT) de forma mais fácil.

Antes de utilizar as funções, você deve estar ciente dos erros que **podem ser retornados** pelas funções, esses erros são cenários conhecidos, nem sempre saberemos quais erros podem ocorrer, aqui está um exemplo:

```go
    // [ Erros estáticos ]

    var (
        // ErrInvalidSecret, é retornado quando a chave secreta não foi a utilizada para assinar o JSON WEB token.
	    ErrInvalidSecret = errors.New("a chave secreta não foi a utilizada para assinar")
    
        // ErrInvalidToken, é retornado quando o token gerado não é válido, *muito provavelmente por um erro do cliente*.
	    ErrInvalidToken = errors.New("o token não é válido")
    )
```
Nem sempre saberemos qual será o erro, há casos em que o erro é desconhecido *"É melhor prevenir do que remediar"*. Lidaremos com o erro 
desconhecido da mesma forma que com o conhecido.

```go
    _, err := jwt.Tokenize(data, "curto", expire)
    if err != nil {
        if jwt.isShortSecretError(err) {
            fmt.Println(err)
        }

        // Lidar com erro desconhecido.
        panic(err)
    }
```
> [!NOTE] 
> **NÃO DEVE** usar `panic` para lidar com um erro, isto é apenas um exemplo.

### Encriptar com Hash

Este pacote é uma maneira fácil de encriptar informações, ele usa o `bcrypt` por baixo e fornece funções com os custos que devem ser usados no Gedeón para otimização e desempenho dos serviços.

> [!NOTE] 
> Este pacote é usado para dados sensíveis como senhas, portanto, a informação a ser encriptada tem um comprimento mínimo de 8 caracteres e máximo de 50. (Sujeito a alteração)

Erros **conhecidos**

 que o pacote pode retornar:

```go
    var (
        // ErrSameHashAndPassword, retornado quando a senha e o hash fornecido são iguais.
        ErrSameHashAndPassword = errors.New("a senha e a senha encriptada não podem ser iguais")
	    
        // ErrEmptyPassword, retornado quando a senha é nula ou a string está vazia ("").
        ErrEmptyPassword       = errors.New("a senha encriptada ou a senha não podem estar vazias")
    )
```