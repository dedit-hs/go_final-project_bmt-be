# BMT API Documentation

## Register Anggota API

Endpoint: POST /api/users

Request Body:

```json
{
  "nama": "Dedit",
  "nik": 2171031908889013,
  "alamat": "Desa Ngroto RT. 9 RW. 4 Pujon",
  "no_hp": "+6285264865959",
  "email": "deditherys@gmail.com",
  "password": "rahasia"
}
```

Response Body Success:

```json
{
  "message": "registration success",
  "data": {
    "id": 1,
    "nama": "Dedit",
    "nik": 2171031908889013,
    "alamat": "Desa Ngroto RT. 9 RW. 4 Pujon",
    "no_hp": "+6285264865959",
    "email": "deditherys@gmail.com"
  }
}
```

Response Body Failed:

```json
{
  "message": "registration failed",
  "error": "nik already registered"
}
```

## Login Anggota API

Endpoint: POST /api/login

Request Body:

```json
{
  "nik": "2171031908889013",
  "password": "rahasia"
}
```

Response Body Success:

```json
{
  "token": "xYcc12hkASG6hasjd7POilaz238has87d"
}
```

Response Body Error:

```json
{
  "message": "login failed",
  "error": "invalid credentials"
}
```

## Update Profile Anggota API

Endpoint: PUT /api/users/1/profile

Request Body:

```json
{
  "nama": "Dedit Hery Suprastyo",
  "alamat": "Desa Ngroto RT. 9 RW. 4 Pujon",
  "no_hp": "+6285264865959",
  "email": "deditherys@gmail.com"
}
```

Response Body Success:

```json
{
  "message": "update profile success",
  "data": {
      "nama": "Dedit Hery Suprastyo",
      "nik": 21713001708459013,
      "alamat": "Desa Ngroto RT. 9 RW. 4 Pujon",
      "no_hp": "+6285264865959",
      "email": "deditherys@gmail.com"
  }
}
```

## Update Password Anggota API

Endpoint: PUT /api/users/1/password

Request Body:

```json
{
  "password": "dirahasiakan"
}
```

Response Body Success:

```json
{
  "message": "update password success",
}
```
