# smofa
/comment, only can use in first line
$name  prm1,prm2,...  //exec func name
!name  prm          //exec func name and all string after name as param
@name1=val1,name2=val2,...  //set name=value in current space
>name        //tail another command, the lines save to name.([]string)
?name       //print name, name.Value in console, for debug
%name       //string buffer is string[]
:name num         //a tag , for goto command. nu means add num line default = 0
~name      //goto tag:name
<name file       //output buffer to file.
)               //is a command,but no use.

