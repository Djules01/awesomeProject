# awesomeProject

API Go de gestion de todo list avec Chi et SQLite.

## Prerequis

- Go 1.26.4 ou compatible

## Configuration

L'application peut etre configuree avec ces variables d'environnement :

- `API_KEY` : cle attendue par le middleware API key, valeur par defaut `humancraft`
- `DB_PATH` : chemin de la base SQLite, valeur par defaut `todos.db`
- `PORT` : port HTTP, valeur par defaut `8080`

## Lancer le projet

```powershell
go run .
```

Le serveur ecoute par defaut sur `http://localhost:8080`.

## Tests

```powershell
go test ./...
```

## Routes

- `GET /ToDoList`
- `POST /Creation`
- `GET /AfficherParDate`
- `PUT /Modifier/{id}`
- `DELETE /Delete/{id}`
