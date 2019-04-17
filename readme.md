# goinsert
goinsert is tool for insert text to file  
known another solution as used `sed` command, but  
can't get i expected result in mac  
  
you can use to `echo` command if you want insert  to last line  

# example
hoge.txt
````
text
````

command
````
goinsert -file=hoge.txt -text=aaa
````

hoge.txt
````
aaa
text
````

# example 2
````
find . -iname *.hoge | xargs goinsert -text=aaa
````