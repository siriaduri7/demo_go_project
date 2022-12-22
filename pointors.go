/* 
SYNTAX
var <pointer_name> *<data_type>

########### INITIALIZING A POINTER###########
var <pointer_name> *<data_type> = &<variable_name>

EX: i :=10
var ptr_i *int =&i
EX:
var ptr_i *int
var ptr_s *string

Zero value of a pointer is nil.

A pointer is a variable that containing memory address of another variable.
 pointers considered as special variables, and this is because they do not only hold memory addresses
but they also point where this memory is located and provide ways to find or even change the value that's located
at that memory location.