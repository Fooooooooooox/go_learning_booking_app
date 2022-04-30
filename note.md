## loop

there are while, do while, for each in other languages

but in go 

you only have "for" for loop

1. infinite loop

for{
    
}

2. finite loop

 
// range is for making iterable stuff

for index, booking := range bookings {

}

//but we don't need index actually, so we can write in this way:

for _, booking := range bookings {

}




## unused variables

in go, if you define a variable but don't use it it throws an error

if you want to define an unused variable: use _


## export function

if you want your function to be used in other file, you can capitalize the first letter of the function name, then it'll be automatically exported.

and you can do the same with variables

you capitalize the first letter and it'll be automatically exported