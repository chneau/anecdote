# anecdote
Ca donne une french anecdote !

# Install
You will need go installed.  (`sudo apt install golang-go` or `choco install go` ...)  
The following command will install the binary.  
You can invoke the program by typing `anecdote`. (or `anecdote.exe` for windows)
```bash
go get -u -v github.com/chneau/anecdote
```

# Usage
To get an anecdote:
```bash
$ anecdote
En France, il est interdit de s'embrasser sur des rails.
```
The default source is SI which is "Savoir Inutile".  
You can specify the source as SI or SCMB. ("Se Coucher Moins Bete")
```bash
$ anecdote -source SI
Dans le premier film Rambo, sorti en 1982, il n'y a qu'un seul mort, alors que dans le film Titanic de James Cameron sorti en 1997, on en dénombre 307.
```
SCMB shows the title followed by the content of the anecdote.
```bash
$ anecdote -source SCMB
Etre éboueur à New York est un job de rêve
Etre éboueur dans la ville de New-York est un bon poste, à cause du salaire et l'image positive qu'ils suscitent dans la population. En effet, ils peuvent gagner jusqu'à 70 000 dollars par an et sont considérés comme des « héros » du quotidien. Un métier tellement prisé qu'il existe une liste d'attente de 8 ans pour y postuler.
``` 

