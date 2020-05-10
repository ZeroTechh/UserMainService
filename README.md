# UserMainService

A velocity which will be handling user main data.

## Functions

### Add

Takes in user data, then add it into database. Returns user ID.

```go
addResponse, err := client.Add(ctx, &UserData)
```

### Get

Takes in either user ID, username or email, then returns user data.

```go
getResponse, err := client.Get(ctx, &proto.GetRequest{
    Username: "username",
    UserID: "userid",
    Email: "email", // Either can be given
})
```

### Update

Takes in user ID and data to update.

```go
updateResponse, err := client.Update(ctx, &proto.UpdateRequest{
    UserID: "Userid",
    Update: UpdateData,
})
```

### Auth

Takes in either username or email and password. Then returns if credentials are valid.

```go
authResponse, err := client.Auth(ctx, &proto.AuthRequest{
    Username: "Username",
    Email: "Email",
    Password: "password",
})
```

### Validate

Takes in data and returns if it is valid.

```go
valid, _ := client.Validate(ctx, &UserData)
```
