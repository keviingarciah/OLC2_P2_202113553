## Create Backemd Project

```bash
go get github.com/antlr4-go/antlr/v4
go get github.com/gofiber/fiber/v2
antlr4 -Dlanguage=Go -o parser -package parser -visitor -no-listener *.g4
```

## Create Frontend Project

```bash
npm create vite
Project name: frontend
Select a framework: react
Select a variant: js
```

```bash
cd frontend
npm i
npm run dev
```

## Dependencies

```bash
npm i monaco edit
npm i react-bootstrap bootstrap axios
npm i react-monaco-editor
```
```bash
npm i d3
npm i graphviz react-graphviz
npm i svg
```

## ANTLR4

```bash
antlr4 -Dlanguage=Go -o parser -package parser -visitor -no-listener *.g4
```