package main

import "fmt"

type Student struct {
    Name  string
    Grade string
}

func groupByGrade(students []Student) map[string][]Student {
    groups := make(map[string][]Student)
    
    for _, student := range students {
        groups[student.Grade] = append(groups[student.Grade], student)
    }
    
    return groups
}

func displayStudentGroupWise() {
    students := []Student{
        {"Alice", "A"},
        {"Bob", "B"},
        {"Charlie", "A"},
        {"David", "C"},
        {"Eve", "B"},
    }
    
    grouped := groupByGrade(students)
    
    for grade, students := range grouped {
        fmt.Printf("Grade %s: ", grade)
        for _, student := range students {
            fmt.Printf("%s ", student.Name)
        }
        fmt.Println()
    }
}