## Objective: Restrict responses to a specific block of text.

### 1. Place a passage inside delimiters (### or triple backticks).

###
The Amazon rainforest is the largest tropical rainforest in the world currently. 
It is located primarily in Brazil and plays a crucial role in regulating 
the Earth's climate by absorbing carbon dioxide.
###

### 2. Add instructions: “Answer using only the text inside the delimiters.”

###
The Amazon rainforest is the largest tropical rainforest in the world currently. 
It is located primarily in Brazil and plays a crucial role in regulating 
the Earth's climate by absorbing carbon dioxide.
###
Question: where is the Amazon rainforest located?

Answer using only the text inside the delimiters.

### 3. Compare the output with and without delimiters to see the effect.

#### OutPut With Delimiters:

primarily in Brazil

#### OutPut Without Delimiters:
primarily in Brazil

### COMPARISON:
WITH DELIMITERS and instruction("Answer using only the text inside the delimiters"): the response is focus and limited to the provided passage and does not include any outside knowledge or additional details.

WITHOUT DELIMITERS the response may include external information beyond the given passage, resulting in a broader but less controlled answer.