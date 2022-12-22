/* SYNTAX////need to declare mapname tha only we can use above syntax
var <map_name> map[<key_data_type>]<value_data_type

<map_name> := map[,key_data_type]<value_data_type>{<key-value-pairs>}
ex: 
codes :=map[string]string{"en":"english","fr": "french"}
var codes map[string]string

//need to declare mapname tha only we can use above syntax

ex:  var my_map map[string]int

#############################################################################################
INITILIZING MAKE FUNCTION

<map_name> := make(map[<key_data_type>]<value_data_type>,<initial_capacity>)
#############################################################################################

