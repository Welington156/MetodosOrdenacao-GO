package main

import "fmt"
func selection(v []int){
    indice_atual := 0
    menor_indice := 0
    i:=1
    for indice_atual < len(v)-1 {
        menor_indice = indice_atual
        for i := indice_atual + 1; i < len(v); i++ {
            if v[menor_indice] > v[i] {
                menor_indice = i
            }
        }
        inverter(indice_atual, menor_indice, v)
        indice_atual++
        fmt.Println("Passo",i,":",v)
        i++
    }
    fmt.Println("Vetor Ordenado :",v)
}

func inverter(i, j int, v []int) {
    v[i], v[j] = v[j], v[i]
}
func insertion(v []int) {
    var n = len(v)
    for i := 1; i < n; i++ {
        j := i
        for j > 0 {
            if v[j-1] > v[j] {
                v[j-1], v[j] = v[j], v[j-1]
            }
            j = j - 1
        }
           fmt.Println("Passo",i,":",v)
    }
    fmt.Println("Vetor Ordenado :",v)
}
func buble(x []int,end int){
      i:=1
    for {
        if end == 0 {
            break
        }
        for i := 0; i < len(x)-1; i++ {
            if x[i] > x[i+1] {
                x[i], x[i+1] = x[i+1], x[i]
            }
        }
        end -= 1
        fmt.Println("Passo",i,":",x)
        i++
}
    fmt.Println("Vetor Ordenado :",x)
}
func main(){

    var a,d int

    fmt.Println(" ==== 4º Seminário - Paradigmas ==== ")
    fmt.Print(" Qnt de números:")
    fmt.Scan(&a)
    v:=make([]int,a)
    fmt.Print("Digite o vetor:")
        for i:=0;i<len(v);i++{
        fmt.Scan(&v[i])
        }

    fmt.Println("Vetor Escrito:",v)
    fmt.Println("Escolha o Metodo de Ordenaçao desejado:\n1-Bubble Sort\n2-Insertion Sort\n3-Selection Sort\n")
    fmt.Print("Metodo Desejado:")
    fmt.Scan(&d)
    switch d {
    case 1: 
    buble(v,a)
    case 2:
    insertion(v)
    case 3:
    selection(v)
    }

}