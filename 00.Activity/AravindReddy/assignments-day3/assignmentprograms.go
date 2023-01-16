1.Golang Program to Print Hello World
package main
import "fmt"
func main(){
    fmt.Println("hello world)
}
2.Go Program to add Two Numbers
package main
func main(){
    a:=1
    b:=2
    c:=a+b
    fmt.Printf("the sum of %d and %d is %d",a,b,c)
}
3.Go Program to Find the Compound Interest
package main

import (
    "fmt"
    "math"
)
func main() {

    var Pamount, InterestRate, timePeriod, ciFuture, compoundI float64

    fmt.Print("Enter the Total or Principal Amount = ")
    fmt.Scanln(&Pamount)

    fmt.Print("Enter the Interest Rate = ")
    fmt.Scanln(&InterestRate)

    fmt.Print("Enter the Total number of Years = ")
    fmt.Scanln(&timePeriod)

    ciFuture = Pamount * (math.Pow((1 + InterestRate/100), timePeriod))
    compoundI = ciFuture - Pamount

    fmt.Println("\nThe Compound Interest  = ", compoundI)
    fmt.Println("The Future Compound Interest  = ", ciFuture)
}
4.Go Program to Count Digits in a Number
package main
func main(){
    v:=123
    for v{
        v=v/10
        count=count+1
    }
    fmt.println(count)
}
5.Go program to count total notes in an Amount
package main

import "fmt"

func main() {
    notes := [8]int{500, 100, 50, 20, 10, 5, 2, 1}
    var amount int
    fmt.Print("Enter the Total Amount of Cash = ")
    fmt.Scanln(&amount)

    temp := amount
    for i := 0; i < 8; i++ {
        fmt.Println(notes[i], " Notes = ", temp/notes[i])
        temp = temp % notes[i]
    }
}
6.Go program to find the cube of a number
package main

import "fmt"

func main() {

    var num int

    fmt.Print("Enter the Number to find Cube = ")
    fmt.Scanln(&num)

    cube := num * num * num

    fmt.Println("\nThe Cube of a Given Number  = ", cube)
}
7.Go program to calculate the employee salary
mport "fmt"

func main() {

    var basicSal, hra, da, grossSal float64

    fmt.Print("Enter the Employee Basic Salary = ")
    fmt.Scanln(&basicSal)

    if basicSal <= 10000 {
        hra = (basicSal * 8) / 100
        da = (basicSal * 10) / 100
    } else if basicSal <= 20000 {
        hra = (basicSal * 16) / 100
        da = (basicSal * 20) / 100
    } else {
        hra = (basicSal * 24) / 100
        da = (basicSal * 30) / 100
    }

    grossSal = basicSal + hra + da
    fmt.Println("The Gross Salary of this Employee = ", grossSal)
}
8.Go program to calculate the electricity bill
package main

import "fmt"

func main() {

    var units, surCharge, amount, totAmount float64

    fmt.Print("Enter the Consumed Units = ")
    fmt.Scanln(&units)

    if units < 50 {
        amount = units * 2.60
        surCharge = 25
    } else if units <= 100 {
        amount = 130 + ((units - 50) * 3.25)
        surCharge = 35
    } else if units <= 100 {
        amount = 130 + 162.50 + ((units - 100) * 5.26)
        surCharge = 45
    } else {
        amount = 130 + +162.50 + 526 + ((units - 200) * 7.75)
        surCharge = 55
    }
    totAmount = amount + surCharge
    fmt.Println("Electricity Bill = ", totAmount)
}
9.Go program to check even or odd
package main
import "fmt"
func main(){
	var num int
	fmt.Println("enter the number)
	fmt.Scanln(&num)
	if num%2==0{
		fmt.Println("The number is even")
	}else{
		fmt.Println("The number is odd")
	}
}
10. Go program to print even numbers from 1 to N
package main
import "fmt"
func main(){
	var n int
	fmt.Print("enter the number")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		ifi%2==0{
			fmt.Println(i)
		}
	}
}
11. Go program to find the factors of a number
package main
import fmt
func main(){
	var n int
	fmt.Println("enter the number")
	fmt.Scanln(&n)
	for i:=1;i<=n;i++{
		if(n%i==0){
			fmt.Println(i)
		}
	}
}
12. Go program to find the factorial of a number
package main
import "fmt"
func factorial(num int) int{
   if num == 1 || num == 0{
      return num
   }
   return num * factorial(num-1)
}
func main(){
   fmt.Println(factorial(3))
   fmt.Println(factorial(4))
   fmt.Println(factorial(5))
}
13. Go program to find the first digit of a number
package main
import "fmt"
func main(){
	var n int
	fmt.Println("enter the number")
	fmt.Scanln(&n)
	a:=n/10
	fmt.Println(a)
}
14. Go program to find the generic root of a number
package main
import "fmt"
func main() {
var genNumber, sum, remainder int
fmt.Print("Enter the Number to find Generic Root = ")
    fmt.Scanln(&genNumber)

    for genNumber >= 10 {
        for sum = 0; genNumber > 0; genNumber = genNumber / 10 {
            remainder = genNumber % 10
            sum = sum + remainder
        }
        if sum >= 10 {
            genNumber = sum
        } else {
            fmt.Println("The Generic Root of this Number = ", sum)
        }
    }
}
15. Go program to find the largest of two numbers
package main
import "fmt"
func main(){
	var a,b int
	fmt.Println("enter the numbers")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if(a>b){
		fmt.Println(a,"is bigger")
	}else if(b>a){
		fmt.Println(b,"is bigger")
	}else{
		fmt.Println("both are equal")
	}
}
16. Go program to find largest of three numbers
ackage main

import "fmt"

func main() {

    var a, b, c int

    fmt.Print("\nEnter the Three Numbers to find Largest = ")
    fmt.Scanln(&a, &b, &c)

    if a > b && a > c {
        fmt.Println(a, " is Greater Than ", b, " and ", c)
    } else if b > a && b > c {
        fmt.Println(b, " is Greater Than ", a, " and ", c)
    } else if c > a && c > b {
        fmt.Println(c, " is Greater Than ", a, " and ", b)
    } else {
        fmt.Println("Either of them are Equal ")
    }
}
17. Go program to check leap Years
package main
import "fmt"
func main() {
   var year int
   fmt.Print("Enter the year to be checked:")
   fmt.Scanf("%d", &year)
   if year%4==0 && year%100!=0 || year%400==0{
      fmt.Println("The year is a leap year!")
   }else{
      fmt.Println("The year isn't a leap year!")
   }
}
18. Go program to print the multiplication table
package main

import "fmt"

func main(){
    var n int
    fmt.Print("Enter any Integer Number : ")
    fmt.Scan(&n)
    i:=1
    /*     For loop as a Go's While     */
    for {
        if(i>10){
            break;
        }
        fmt.Println(n," X ",i," = ",n*i)
        i++
    }
}
19. Go program to print natural numbers from 1 to N
package main

import "fmt"

func main() {

    var num, i int

    fmt.Print("\nEnter the Maximum Natural Number Limit = ")
    fmt.Scanln(&num)

    fmt.Println("\nNatural Numbers from 1 to ", num, " are = ")
    for i = 1; i <= num; i++ {
        fmt.Print(i, "\t")
    }
    fmt.Println()
}
20. Go program to natural numbers in reverse order
package main

import "fmt"

func main() {

    var num, i int

    fmt.Print("\nEnter the Last Natural Number = ")
    fmt.Scanln(&num)

    fmt.Println("\nNatural Numbers from ", num, " to 1 are = ")
    for i = num; i >= 1; i = i - 1 {
        fmt.Print(i, "\t")
    }
    fmt.Println()
}
21. Go program to find ncr factorial of a number
package main

import "fmt"

func factorialCal(number int) int {
    factorial := 1
    for i := 1; i <= number; i++ {
        factorial = factorial * i
    }
    return factorial
}
func main() {

    var ncr, n, r int

    fmt.Print("Enter any N and R Values = ")
    fmt.Scanln(&n, &r)

    ncr = factorialCal(n) / (factorialCal(r) * factorialCal(n-r))

    fmt.Println("The NCR Factorial = ", ncr)
}
22. Go program to find a number divisible by 5 and 11
package main

import "fmt"

func main() {

    var num int

    fmt.Print("Enter the Number that may Divisible by 5 and 11 = ")
    fmt.Scanln(&num)

    if num%5 == 0 && num%11 == 0 {
        fmt.Println(num, " is Divisible by 5 and 11")
    } else {
        fmt.Println(num, " is Not Divisible by 5 and 11")
    }
}
23. Odd numbers from 1 to n
package main

import "fmt"

func main() {

    var odnum int

    fmt.Print("Enter the Number to Print Odd's = ")
    fmt.Scanln(&odnum)

    for x := 1; x <= odnum; x++ {
        if x%2 != 0 {
            fmt.Print(x, "\t")
        }
    }
}
24. go program to find the product of digits of  a number
package main

import "fmt"

func main() {

    var prodNum, product, prodReminder int

    fmt.Print("Enter the Number to find the Digits Product = ")
    fmt.Scanln(&prodNum)

    for product = 1; prodNum > 0; prodNum = prodNum / 10 {
        prodReminder = prodNum % 10
        product = product * prodReminder
    }
    fmt.Println("The Product of Digits in this Number = ", product)
}
25. go program to check palindrome number
package main

import "fmt"

func main() {

    var palNum, remainder int

    fmt.Print("Enter the Number to check Palindrome = ")
    fmt.Scanln(&palNum)

    reverse := 0

    for temp := palNum; temp > 0; temp = temp / 10 {
        remainder = temp % 10
        reverse = reverse*10 + remainder
    }

    fmt.Println("The Reverse of the Given Number = ", reverse)
    if palNum == reverse {
        fmt.Println(palNum, " is a Palindrome Number")
    } else {
        fmt.Println(palNum, " is Not a Palindrome Number")
    }
}
26. go program for perfect number
package main

import "fmt"

func main() {

    var perfNum, perfsum int
    perfsum = 0

    fmt.Print("Enter the Number to find the Perfect = ")
    fmt.Scanln(&perfNum)

    for i := 1; i < perfNum; i++ {
        if perfNum%i == 0 {
            perfsum = perfsum + i
        }
    }

    if perfNum == perfsum {
        fmt.Println(perfNum, " is a Perfect Number")
    } else {
        fmt.Println(perfNum, " is Not a Perfect Number")
    }
}
27. go program to check prime number
package main
import (
   "fmt"
   "math"
)

func checkPrimeNumber(num int) {
   if num < 2 {
      fmt.Println("Number must be greater than 2.")
      return
   }
   sq_root := int(math.Sqrt(float64(num)))
   for i:=2; i<=sq_root; i++{
      if num % i == 0 {
         fmt.Println("Non Prime Number")
         return
      }
   }
   fmt.Println("Prime Number")
   return
}

func main(){
   checkPrimeNumber(0)
   checkPrimeNumber(2)
   checkPrimeNumber(13)
   checkPrimeNumber(152)
}
28. go program to check positive or negative
package main

import "fmt"

func main() {

    var num int

    fmt.Print("Please enter any Integer = ")
    fmt.Scanf("%d", &num)

    if num >= 0 {
        fmt.Println("Positive Integer")
    } else {
        fmt.Println("Negative Integer")
    }
}
29. go program to calculate power of a number
func POWER(num int, power int) int {
	var result int = 1
	if power != 0 {
	
	   // Recursive function call to itself
	   result = (num * POWER(num, power-1))
	}
	return result
 }
 func main() {
	fmt.Println("Golang Program to calculate the power using recursion")
	// declare and initialize the integer variables
	var base int = 4
	var power int = 2
	var result int
	
	// calling the POWER() function
	result = POWER(base, power)
	
	// Print the result using in-built function fmt.Printf()
	fmt.Printf("%d to the power of %d is: %d\n", base, power, result)
 }
 30. go program for profit or loss
 package main

import "fmt"

func main() {

    var pcost, sa, amount int

    fmt.Print("\nEnter the Actual Product Cost = ")
    fmt.Scanln(&pcost)

    fmt.Print("\nEnter the Sale Price = ")
    fmt.Scanln(&sa)

    if sa > pcost {
        amount = sa - pcost
        fmt.Println("Total Profit = ", amount)
    } else if pcost > sa {
        amount = pcost - sa
        fmt.Println("Total Loss = ", amount)
    } else {
        fmt.Println("No Profit No Loss")
    }
}
31. print 1 t0 100
package main
import "fmt"
func main(){
	for i:=1;i<=100;i++{
		fmt.Println(i)
	}
}
32. print 1 to 100 without loop
package main

import "fmt"

func printNumbers(num int) {
    if num <= 100 {
        fmt.Print(num, "\t")
        printNumbers(num + 1)
    }
}
func main() {

    number := 1
    printNumbers(number)
}
32. reverse a number
package main
import "fmt"
func reverseNumber(num int) int {

   res := 0
   for num>0 {
      remainder := num % 10
      res = (res * 10) + remainder
      num /= 10
   }
   return res
}

func main(){
   fmt.Println(reverseNumber(168))
   fmt.Println(reverseNumber(576))
   fmt.Println(reverseNumber(12345))
}
33. go program for root of quadratic equation
package main

import (
    "fmt"
    "math"
)

func main() {

    var a, b, c, root1, root2, imaginary, discriminant float64

    fmt.Print("Enter the a, b, c of Quadratic equation = ")
    fmt.Scanln(&a, &b, &c)

    discriminant = (b * b) - (4 * a * c)

    if discriminant > 0 {
        root1 = (-b + math.Sqrt(discriminant)/(2*a))
        root2 = (-b - math.Sqrt(discriminant)/(2*a))
        fmt.Println("Two Distinct Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
    } else if discriminant == 0 {
        root1 = -b / (2 * a)
        root2 = -b / (2 * a)
        fmt.Println("Two Equal and Real Roots Exist: root1 = ", root1, " and root2 = ", root2)
    } else if discriminant < 0 {
        root1 = -b / (2 * a)
        root2 = -b / (2 * a)
        imaginary = math.Sqrt(-discriminant) / (2 * a)
        fmt.Println("Two Distinct Complex Roots Exist: root1 = ", root1, "+", imaginary, " and root2 = ", root2, "-", imaginary)
    }
}
