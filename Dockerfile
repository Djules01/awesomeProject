# =============================================================================
# Build multi-étapes : on compile dans une image "builder", puis on ne garde
# que le binaire dans une image Alpine minimale
# =============================================================================

# --- Étape 1 : compilation ---
# Image officielle Go sur Alpine (petite, adaptée au build).
FROM golang:1.26-alpine AS builder

# Répertoire de travail dans le conteneur de build.
WORKDIR /app

# Copie des fichiers de dépendances en premier : Docker met en cache cette
# couche tant que go.mod / go.sum ne changent pas, ce qui accélère les rebuilds.
COPY go.mod go.sum ./
RUN go mod download

# Copie du reste du code source et compilation du binaire.
COPY . .
RUN go build -o server .

# --- Étape 2 : image finale légère ---
# Alpine (~5 Mo) : suffisant pour exécuter le binaire statique, sans Go ni gcc.
FROM alpine:latest

WORKDIR /app

# On ne récupère que le binaire compilé depuis l'étape builder (pas les sources).
COPY --from=builder /app/server .

# Port HTTP exposé par l'application (documentation ; le mapping se fait au run).
EXPOSE 8080

# Variables lues par configuration.LoadConfig() au démarrage.
# MONGO_URI pointe vers le service "mongo" défini dans docker-compose.yml.
ENV PORT=8080
ENV MONGO_URI=mongodb://mongo:27017
ENV MONGO_DB=todolist
ENV API_KEY=humancraft

# Lance le serveur HTTP (forme exec : PID 1 = le binaire, signaux bien reçus).
CMD ["./server"]
