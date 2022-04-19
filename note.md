# loop

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




# unused variables

in go, if you define a variable but don't use it it throws an error

if you want to define an unused variable: use _