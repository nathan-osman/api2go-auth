## api2go-auth

[![Build Status](https://travis-ci.org/nathan-osman/api2go-auth.svg?branch=master)](https://travis-ci.org/nathan-osman/api2go-auth)
[![GoDoc](https://godoc.org/github.com/nathan-osman/api2go-auth?status.svg)](https://godoc.org/github.com/nathan-osman/api2go-auth)
[![MIT License](http://img.shields.io/badge/license-MIT-9370d8.svg?style=flat)](http://opensource.org/licenses/MIT)

This package simplifies the task of adding authentication to an application using [api2go](https://github.com/manyminds/api2go).

### Features

Here are some of the features that api2go-auth provides:

- Provides methods for login and logout
- Ensures all API methods are authenticated
- Enables full customization of the authentication process

### Server Usage

To use api2g-auth, you must first create a type that implements [`Authenticator`](https://godoc.org/github.com/nathan-osman/api2go-auth#Authenticator). In the following example, user credentials are stored in a database:

```go
type UserAuth struct {}

func (u *UserAuth) Authenticate(r *http.Request) (interface{}, interface{}, error) {
    u, err := isValidUser(r)
    if err != nil {
        return nil, nil, err
    }
    return u.ID, u, err
}

func (u *UserAuth) Initialize(r *http.Request, i interface{}) (*http.Request, error) {
    u, err := fetchUser(i)
    if err != nil {
        return nil, err
    }
    return r.WithContext(
        context.WithValue(r.Context(), "user", u)
    ), nil
}
```

The `Authenticate()` method is invoked when the client attempts to login. Assuming valid credentials are supplied, the method returns both a unique identifier for the user as well as the user object itself (which will be sent to the client).

The `Initialize()` method is invoked before each API request. It loads the user object from the unique identifier (which was returned in `Authenticate()`) and adds a variable to the request context so that it can be used by data sources.

The next step is to simply create an [`Auth`](https://godoc.org/github.com/nathan-osman/api2go-auth#Auth) instance:

```go
var (
    api = api2go.NewAPI("api")
    h   = auth.New(api, &UserAuth{}, nil)
)
```

`h` can then be used as an HTTP handler.

### Client Usage

Clients must log in my sending a POST request to the `/login` endpoint and including the data expected by `Authenticate` (a username and password, for example). If successful, the data returned by `Authenticate` will be send to the client in JSONAPI format. A cookie will be set that authenticates future requests.

When a session is ready to be ended, the client may send a POST request to `/logout` to destroy the session.