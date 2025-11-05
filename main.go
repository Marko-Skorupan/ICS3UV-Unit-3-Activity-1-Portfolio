/*
 * @author Marko Skorupan
 * @version 1.0.0
 * @date 2025-11-05
 * @fileoverview "This program sums 3 constants"
 */

package main

import "fmt"

func main(){
//Set the constants
const NUMBER1 int = 10
const NUMBER2 int = -20
const NUMBER3 int = 85

//INPUT - none

//PROCESS
//calculate the sum
answer := NUMBER1 + NUMBER2 + NUMBER3

//OUTPUT
//display the result
fmt.Print("The sum of ", NUMBER1," & " , NUMBER2," & ", NUMBER3," is ", answer)

fmt.Println("\nDone.")
}