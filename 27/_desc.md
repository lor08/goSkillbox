Цель задания

Научиться работать с композитными типами данных: структурами и картами


Что нужно сделать

Напишите программу, которая считывает ввод с stdin, создаёт структуру student и записывает указатель на структуру в хранилище map[studentName] *Student.

```go
type Student struct {

    name string
    
    age int
    
    grade int

}
```

Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent, далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d) вывести на экран имена всех студентов из хранилища. Также необходимо реализовать методы put, get. 