<h1>DSA</h1>

04/ 08/2025
---------------
🎯 আজকের টাস্ক:
    1. নিচের তথ্যগুলো ইনপুট নাও ও প্রিন্ট করো:
        ◦ নাম (string)
        ◦ বয়স (int)
        ◦ পছন্দের সংখ্যা (float)
    2. নিচের মতো প্রিন্ট করো:

Hello Palash! You are 25 years old. Your lucky number is 3.14

✨ আজকের মূল শিক্ষা:
    • Go-এর প্রোগ্রাম কিভাবে শুরু হয়
    • Variable কী
    • Input/output কিভাবে কাজ করে
    • Short-hand variable syntax

তুমি কোডটা চালিয়ে দেখো, কোনো সমস্যা হলে আমি পাশে আছি।
তুমি করলে, আমি আগামীকাল Day 2: Loops & Conditions শুরু করব। 🔄


MY Code
=======

package main
import "fmt"
func main() {

var name string = "PALASH"
var age int = 30
var lucky_num float64 = 1

fmt.Print("Enter your Name : - ")
fmt.Scanln(&name)
fmt.Println(name)
fmt.Print("Enter Your Age")
fmt.Scanln(&age)
fmt.Println(age)
fmt.Print("Enter Lucky Number :- ")
fmt.Scan(&lucky_num)
fmt.Println(lucky_num)

}


তোমার কোডটা একদম ঠিক কাজ করছে, তবে আমি এটাকে একটু clean, efficient, and professional করে সাজিয়ে দিচ্ছি — যেন ভবিষ্যতে তুমি বড় প্রজেক্টেও এমন কোড লিখতে পারো।

package main

import (
"fmt"
)

func main() {
var (
	name string
	age int
	luckyNum float64
)

fmt.Print("Enter your Name: ")
fmt.Scanln(&name)

fmt.Print("Enter your Age: ")
fmt.Scanln(&age)

fmt.Print("Enter your Lucky Number: ")
fmt.Scanln(&luckyNum)

fmt.Println("\n--- Output ---")
fmt.Printf("Hello, %s!\n", name)
fmt.Printf("You are %d years old.\n", age)
fmt.Printf("Your lucky number is %.2f\n", luckyNum)
}

amar error hoche .
ChatGPT said:
ঠিক আছে — তুমি একদম ঠিক জায়গায় এসেছো।
আমি তোমার error দেখে ঠিক করে দেবো। শুধু আমাকে নিচের দুটি জিনিস দাও:

✅ আমি কী চাই:
    1. ✅ তুমি যে কোড লিখেছো সেটার পুরোটা এখানে paste করো
    2. ✅ Error message টা হুবহু লিখে দাও (যেমন: undefined: something, cannot use ..., ইত্যাদি)

এই দুটো পেলেই আমি:
🔧 ভুল ধরব
🛠️ ঠিক করে দেব
🧠 ব্যাখ্যা করব কেন ভুল হলো

তাহলে, তোমার কোড আর error-msg টা এখনই পাঠাও — আমি রেডি! 🧑‍💻💬
You said:
code
package main

import "fmt"

func main() {
var (
name string
age int
lucky_num float64
)

fmt.Print("Enter Youe Name:- ")
fmt.Scanln(&name)

fmt.Print("Enter Your Age:- ")
fmt.Scanln(&age)

fmt.Print("Enter Your Lucky Number:- ")
fmt.Scanln(&lucky_num)

fmt.Println("\n-- OUTPUT--")
fmt.Printf("Hello %s \n", name)
fmt.Printf("Your Age %d \n", age)
fmt.Printf("Your Lucky Number %.2f \n", lucky_num)

}

 error  
Enter Youe Name:- PALASH KAYAL
Enter Your Age:- 
Enter Your Lucky Number:- -- 
OUTPUT-- 
Hello PALASH :- 
Your Age 0:-
Your Lucky Number 0.000000 :- 
ChatGPT said:
খুব ভালো, তুমি দারুণ চেষ্টা করছো! 👏
তোমার কোডে একটা ছোট্ট ভুল হয়েছে — যা প্রায় সবাই করে। সেটা আমি স্পষ্টভাবে বুঝিয়ে বলছি:

🛠️ সমস্যার মূল কারণ: fmt.Scanln(&name) একটাই শব্দ নেয়
তুমি লিখেছো:
ruby
CopyEdit
Enter Youe Name:- PALASH KAYAL
কিন্তু fmt.Scanln(&name) কেবল PALASH নেয়।
KAYAL তখন buffer-এ থেকে যায়, এবং Age ইনপুট নেবার সময় কিছুই নেয় না।
তাই age ও lucky_num = 0 হয়ে যায়।

✅ সমাধান ১: পুরো লাইন নাও bufio দিয়ে
এটা সবচেয়ে সঠিক ও পেশাদার সমাধান।
go
CopyEdit



package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your Name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)

    fmt.Print("Enter your Age: ")
    ageInput, _ := reader.ReadString('\n')
    ageInput = strings.TrimSpace(ageInput)
    age, _ := strconv.Atoi(ageInput)

    fmt.Print("Enter your Lucky Number: ")
    luckyInput, _ := reader.ReadString('\n')
    luckyInput = strings.TrimSpace(luckyInput)
    luckyNum, _ := strconv.ParseFloat(luckyInput, 64)

    fmt.Println("\n--- OUTPUT ---")
    fmt.Printf("Hello %s\n", name)
    fmt.Printf("Your Age: %d\n", age)
    fmt.Printf("Your Lucky Number: %.2f\n", luckyNum)
}

✅ সমাধান ২ (সরল): শুধু এক শব্দ নাও
তুমি যদি এখনই bufio না শিখতে চাও, তবে:
👉 নাম ইনপুট দেবার সময় একটাই শব্দ দাও (যেমন PALASH)
না হলে name-এর পরের শব্দটা age হিসেবে ধরবে — ভুল করবে।

✍️ প্র্যাকটিস টিপস:
    • একাধিক শব্দের ইনপুট → bufio + strings + strconv combo
    • এক শব্দ → fmt.Scanln() যথেষ্ট

✅ তোমার শেখা এগোচ্ছে — অনেক ভালোভাবে!
তুমি কি চাইলে আমি তোমার এই কোডটা দিয়ে নতুন একটা টাস্ক দিই?
নাকি তুমি চাইবে এখন Condition বা If/Else শেখা শুরু করতে?


package main

import (
"bufio"
"fmt"
"os"
"strconv"
"strings"
)

func main() {
reader := bufio.NewReader(os.Stdin)

fmt.Print("Enter Your Name :-")
name, _ := reader.ReadString('\n')
name = strings.TrimSpace(name)

fmt.Print("Enter Your AGE :-")
ageInput, _ := reader.ReadString('\n')
ageInput = strings.TrimSpace(ageInput)
age, _ := strconv.Atoi(ageInput)

fmt.Print("Enter Your Lucky Number :-")
luckyInput, _ := reader.ReadString('\n')
luckyInput = strings.TrimSpace(luckyInput)
luckyNum, _ := strconv.ParseFloat(luckyInput, 64)

fmt.Println("\n--OUTPUT--")
fmt.Printf("Hello %s\n", name)
fmt.Printf("your AGE : %d\n", age)
fmt.Printf("Your Lucky Number: %.2f\n", luckyNum)
# Golang_DSA
DEPTH DSA
