# ➤ Exercise 1: The Palindrome Challenge

# Step 1 - Do it yourself
##  1. Write pseudocode for a function that checks if a string is a palindrome. 
```
 FUNCTION palindrome(text):
            creat an empty container "clean"
                check each character in text: 
        Add the character to "clean"
    reverse the clean version
    compare if there are the same 
 ```               
## Implement your solution in Python.
## Test with examples like "racecar", "hello", and "A man a plan a canal Panama".
## Add comments explaining your reasoning.
```py
def palindrome(text): #function definition
    clean = "" # an empty string was created named "clean"
    for ch in text: # ch take each character from the string one at a time
        clean += ch # add all the character it took from text
        
    reverse = clean[::-1] # reverse the clean version of the text
    return reverse == clean# compare if the are the same
    
print(palindrome("racecar"))                       #True
print(palindrome("hello"))                         #False
print(palindrome("A man a plan a canal Panama"))   # False
```

# Step 2 - Use AI to learn
## Now that your function works, use AI to go deeper:
### What's the time complexity?
Overall Complexity

Time: O(n²) (dominated by +=)

Space: O(n) (store clean and reverse)
### What edge cases might I miss?
Empty string → returns True (usually okay)

All non-letters → "!!!" → returns True

Digits are ignored → "A1A" → becomes "aa"

Might be unexpected if you want numbers counted

Unicode letters → works if isalpha() supports them

Case differences → handled with .lower()

Large strings → O(n²) will be slow
### Are there better approaches?"
yes there are


# Step 3 - Reflection

## What did you learn from solving it before asking AI?
 > I learnt that there are better way of approaching without the use of AI, and it also improve critical thinking 

## How is your understanding different now? 
> Now i think more logically than before 

## Could you now write similar functions (e.g., reverse a string) without help?
> yes i can because i learnt a lot from the previous exercises
