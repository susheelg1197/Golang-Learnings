/*
Select statement is similar to switch case statement
the only difference in select is that each case will either be send
or recieve data whereas in switch each case is an expression

the select blocks until any of the case is ready.
If multiple case statements are ready then it randomly chooses any one

the select statement is useful when there are multiple go routines
sending data to multiple channels. So select is used to read from multiple channels
concurrently.

the select case is set to nil to ignore the case of that switch
*/

