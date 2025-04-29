# RecapCards
A simple memory training app with a React frontend and a Go backend.
Users can create their own cards and practice remembering their contents.

# Getting started
To enable auto-reloading install air and run it in zsh: `air`
https://github.com/air-verse/air

To work with environment files install direnv
https://github.com/direnv/direnv/blob/master/docs/installation.md
If variables are not visible, you need to run 
```bash 
  direnv allow
```

# Database migrations
Migrations are made with golang-migrate tool. Install like a cli:
```bash
  brew install golang-migrate
```
More insttructions
https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

# Validation
https://github.com/go-playground/validator