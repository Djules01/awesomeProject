# awesomeProject

API Go de gestion de todo list avec Chi et MongoDB.

## Prerequis

- Go 1.26.4 ou compatible
- MongoDB 7 (local ou via Docker)

## Configuration

L'application peut etre configuree avec ces variables d'environnement :

- `API_KEY` : cle attendue par le middleware API key, valeur par defaut `humancraft`
- `MONGO_URI` : URI de connexion MongoDB, valeur par defaut `mongodb://localhost:27017`
- `MONGO_DB` : nom de la base de donnees, valeur par defaut `todolist`
- `PORT` : port HTTP, valeur par defaut `8080`

## Lancer le projet en local

Demarrer MongoDB (ex. via Docker) :

```powershell
docker run -d --name todolist-mongo -p 27017:27017 mongo:7
```

Puis lancer l'API :

```powershell
go run .
```

Le serveur ecoute par defaut sur `http://localhost:8080`.

## Lancer avec Docker Compose

Lance l'API et MongoDB ensemble. Les donnees sont persistees dans le volume Docker `mongo-data`.

```powershell
docker compose up --build
```

- API : `http://localhost:8080`
- MongoDB : `localhost:27017`

Arreter les services :

```powershell
docker compose down
```

## Tests

```powershell
go test ./...
```

Le test d'integration de `main` necessite MongoDB sur `localhost:27017` (sinon il est ignore).

## Routes

- `GET /ToDoList`
- `POST /Creation`
- `GET /AfficherParDate`
- `PUT /Modifier/{id}`
- `DELETE /Delete/{id}`
