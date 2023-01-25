package main

import (
	"fmt"
  "flag"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string
var funfacts string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablere
// er initialisert.
func init() {

  /*
    Her er eksempler på hvordan man implementerer parsing av flagg.
    For eksempel, kommando
        funtemps -F -out C
    skal returnere output: 0°F er -17.78°C
    siden -F har standard verdi 0.0 i flagg-definisjonen under
  */

  flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
  // Du må selv definere flag-variablene for "C" og "K"
  flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
  flag.StringVar(&funfacts, "funfacts", "sun", "\"fun-facts\" om sun - Solen, luna - Månen og terra - Jorden")
  // Du må selv definere flag-variabelen for -t flagget, som bestemmer
  // hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

  flag.Parse()

  /**
    Her må logikken for flaggene og kall til funksjoner fra conv og funfacts
    pakkene implementeres.

    Det er anbefalt å sette opp en tabell med alle mulige kombinasjoner
    av flagg. flag-pakken har funksjoner som man kan bruke for å teste hvor mange flagg og argumenter er spesifisert på kommandolinje.
    fmt.Println("len(flag.Args())", len(flag.Args()))
		fmt.Println("flag.NFlag()", flag.NFlag())

    Enkelte kombinasjoner skal ikke være gyldige og da må kontrollstrukturer
    brukes for å utelukke ugyldige kombinasjoner:
    -F, -C, -K kan ikke brukes samtidig
    disse tre kan brukes med -out, men ikke med -funfacts
    -funfacts kan brukes kun med -t
    ...

  */

  // Her er noen eksempler du kan bruke i den manuelle testingen
  fmt.Println(fahr, out, funfacts)

  fmt.Println("len(flag.Args())", len(flag.Args()))
  fmt.Println("flag.NFlag()", flag.NFlag())

  fmt.Println(isFlagPassed("out"))


}

// Funksjonene sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
    found := false
    flag.Visit(func(f *flag.Flag) {
        if f.Name == name {
            found = true
        }
    })
    return found
}
