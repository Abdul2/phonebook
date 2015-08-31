# Go lang and postgres 


## Intro 

I've started the day trying to understand how GO works with databases but for some reason i hit a time out connection error with MySql which i could not resolve. so to stop wasting more time, i decided to switch to postgres and voila it worked.

I am new to GO programming but very much impressed by its neatness. GO concepts makes perfect sense although at times it takes a bit of getting used to. I also like the idea of building from the ground up. There is talk about speed but I will leave that to the experts. 

Like almost every one I prefer to work with examples and straight diving in but i found reading a book on the grammer of the GO makes the learning experaince  more enjoyable.

code of this example is based on golang-nuts discussion. 

Below is description of how I went about this simple example.

 
## Steps  

 1. Installed postgres.app on Mac machine.  postgres instaltion  creates an posgress account for you.  
 2. created a  database and table
	```
	* $ psql
	* # create database phonebook ;
	* # \c phonebook
	* # create table contact (
   	* id serial primary key ,
   	name varchar(255),
   	phone_number varchar(255));
	* # \q
   ```
  3. created a Go project on Eclipse 
  4. imported pq driver `go get github.com/lib/pq`
  5. looked for sample code and played with it until it worked.
  6. wrote README.md to remind me of what i did (it maybe of use to others)
   
## To do 

  1. Comment on the code
  2. Extend code to cover full db operations 
  3. Organise the code into packages      
  
  
  