[[2017-01-04] Challenge #298 [Intermediate] Too many or too few Parentheses](https://www.reddit.com/r/dailyprogrammer/comments/5m034l/20170104_challenge_298_intermediate_too_many_or/)

this challenge is about spotting where parentheses are unbalanced.

**find the index in string where the first extra `)` or `(` parentheses is, or return the length of string if parentheses are balanced.**

for flashiness, insert `**` around any index that is found. (will bold in markdown)

**inputs:**  

)(asdf)))  
((((asdf)))   
((((asdf))  
(ab)((cd)(asdf)))  
(ab)((cd)(asdf)())  
(ab)(((cd)(asdf)  
(ab)(((cd)(asdf  
(ab)(((cd)(asdf)))))  

**outputs:**

**)**(asdf)))  
**(**(((asdf)))  
(**(**((asdf))  
(ab)((cd)(asdf))**)**  
(ab)((cd)(asdf)())  
(ab)(**(**(cd)(asdf)  
(ab)((**(**cd)(asdf  NB. may argue it should be 1 left.  
(ab)(((cd)(asdf)))**)**)
