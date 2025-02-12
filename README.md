# Automaton-Signals
Educational project developed for the final test of the "Algorithm and data structures" course at the University of Milan.

The goal of this project, is to model a bidimensional plane in which a series of automaton and obstacles are positioned.
Every automaton moves in the plane, and react to specific recalls.

## Project directory structure:
```
41461A_Ingenito_Emiddio/
┣ input-output/
┃ ┣ automi_limite_piano/
┃ ┃ ┣ input.txt
┃ ┃ ┗ output.txt
┃ ┣ automi_sovrapposti/
┃ ┃ ┣ input.txt
┃ ┃ ┗ output.txt
┃ ┣ grande_ostacolo/
┃ ┃ ┣ input.txt
┃ ┃ ┗ output.txt
┃ ┣ numerosi_ostacoli/
┃ ┃ ┣ input.txt
┃ ┃ ┗ output.txt
┃ ┣ piano_molto_grande/
┃ ┃ ┣ input.txt
┃ ┃ ┗ output.txt
┃ ┗ richiami/
┃   ┣ input.txt
┃   ┗ output.txt
┣ Soluzione/
┃ ┣ automa.go #Defines the automaton entity
┃ ┣ coordinate.go #Defines the coordinate structure (X,Y)
┃ ┣ esegui.go #Contains the "esegui" function, which is in charge of executing the user commands
┃ ┣ ostacoli.go #Defines the structure of an obstacle
┃ ┣ piano.go #Defines the plane.
┃ ┣ solution.go #Contains main function
┃ ┣ solution.exe
┃ ┣ #TEST FILES
┃ ┣ base-in
┃ ┣ base-out
┃ ┣ formato_test.go
┃ ┣ lib_test.go
┃ ┗ utils_test.go
┃ ┣ go.mod
┗ Laboratorio “Algoritmi e Strutture Dati” Relazione Ingenito Emiddio.pdf # Formal report about this project.
```
The ```Soluzione``` directory, contains all the files that define my solution.
The ```input-output``` directory, contains a serie of execution examples of my code.


In order to compile the code, ```cd``` into the ```Soluzione``` directory, and launch the ```go build .``` command. You can then execute the code by running ```go run .```


Also a formal report has been attached to the code, which widely explains the implementation decisions i took, and precisely analizes every function written, along with an analisys of their computational and spacial complexity.

---
## Contact

Coded with ❤️ by **@emikodes** - feel free to contact me!
