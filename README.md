# Mercadolibre Job Application test "Fuego de Quasar"

Service for MercadoLibre jon application

## Challenge description

🇨🇱 Han Solo ha sido recientemente nombrado General de la Alianza Rebelde y busca dar un gran golpe contra el Imperio Galáctico para reavivar la llama de la resistencia. 
El servicio de inteligencia rebelde ha detectado un llamado de auxilio de una nave portacarga imperial a la deriva en un campo de asteroides. El manifiesto de la nave es ultra clasificado, pero se rumorea que transporta raciones y armamento para una legión entera.

Como jefe de comunicaciones rebelde, tu misión es crear un programa en Golang que retorne la fuente y contenido del mensaje de auxilio. Para esto, cuentas con tres satélites que te permitirán triangular la posición, ¡pero cuidado! el mensaje puede no llegar completo a cada satélite debido al campo de asteroides frente a la nave. 


🇺🇸 Han Solo has recently been appointed General of the Rebel Alliance and seeks to strike a major blow against the Galactic Empire to rekindle the flame of resistance.
The rebel intelligence service has detected a call for help from an Imperial cargo ship adrift in an asteroid field. The ship's manifest is ultra classified, but is rumored to carry rations and weaponry for an entire legion.

As the rebel communications chief, your mission is to create a Golang program that returns the source and content of the distress message. For this, you have three satellites that will allow you to triangulate the position, but be careful! the message may not reach each satellite in full due to the asteroid field in front of the spacecraft.

## Installation

First, install all dependencies:
```bash
go mod download
```
Then you can start the server:

```bash
go run main.go
```

## Usage

When Go server is running, you can go to the `http://localhost:8080/` in your favorite web browser

### Run tests
To run all tests suites, excecute this command at projects home directory:

```bash
go test ./... -v
```

## API Documentation

### 1.- Embeded Swagger
If you deployed to any cloud service this service, you can check all API endpoints documentation in `/api-docs/index.html` 

### 2.- Endpoints

#### - Healthcheck `GET /`
##### Response:
```javascript
{
    "time":"2021-03-07T19:01:52.974065388Z"
}
```

#### - Get message and coodinates `POST /topsecret/`
##### Request:
```javascript
{ 
    "satellites": [ 
        { 
            “name”: "kenobi", 
            “distance”: 100.0, 
            “message”: ["este", "", "", "mensaje", ""] 
        }, 
        { 
            “name”: "skywalker", 
            “distance”: 115.5 
            “message”: ["", "es", "", "", "secreto"] 
        }, 
        { 
            “name”: "sato", 
            “distance”: 142.7 
            “message”: ["este", "", "un", "", ""] 
        } 
    ] 
}

```
##### Response:
```javascript
200 - OK
{ 
    "position": { 
        "x": -100.0, 
        "y": 75.5 
    }, 
    "message": "este es un mensaje secreto" 
} 
```

##### Response (If we do not provide the information of the 3 satellites or the message or position cannot be determined):
```javascript
404 - Not Found 
```

#### - Inform message fragment and distance by satellite `POST /topsecret_split/:satellite_name`
##### Request:
```javascript
{ 
    "distance": 100.0, 
    "message": ["este", "", "", "mensaje", ""] 
}
```
##### Response:
```javascript
HTTP 201 - Created
```

#### - Get message and coordinates informed `GET /topsecret_split/`
##### Response (if we have the information of the 3 satellites):
```javascript
200 - OK
{ 
    "position": { 
        "x": -100.0, 
        "y": 75.5 
    }, 
    "message": "este es un mensaje secreto" 
} 
```

##### Response (if we do not have the information of the 3 satellites):
```javascript
400 - Bad request
{ 
    "error": "Could not determine the position and message, we need more information" 
} 
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## Bibliography
- https://www.udemy.com/course/golang-how-to-design-and-build-rest-microservices-in-go
- https://www.udemy.com/course/building-modern-web-applications-with-go
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
- https://www.youtube.com/watch?v=ktPuxq3UVX4
- https://tutorialedge.net/golang/parsing-json-with-golang/

## License
[MIT](https://choosealicense.com/licenses/mit/)
