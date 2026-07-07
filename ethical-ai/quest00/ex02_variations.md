# ➤ Exercise 2: Strengthening Through Variations

## **1. Modify your palindrome function to:**

###     Ignore spaces and punctuation.

###     Be case-insensitive.

###     Return the position where the string stops being a palindrome (if not one).      

```py
def palindrome(text):
    clean = ""
    for ch in text:
        if ch.isalnum():
            clean += ch.lower()
    left = 0
    right = len(clean)-1
    
    while left < right:
        if clean[left] != clean[right]:
            return f"it is not a palindrome. failed at index {left}"
        left += 1
        right -= 1
        
    return "this is a palindrome"
    
print(palindrome("racecar"))
print(palindrome("hello"))
print(palindrome("A man a plan a canal Panama"))
```
## 2. After your first attempt, ask AI:  
> "I modified my palindrome function to handle more cases.
Did I miss anything? Can it be more efficient?"

### strategic AI usage
```
In this particular project is used AI a learning amplifier 
I used AI to check if they are possible edge cases i need to consider
I used "clean += ch.lower" to add the cleaner version of the text to the container, but since strings are immutable in python, this makes my cleaning step O(n2)
so AI gave me a more efficient pattern which is "clean_chars.append(ch.lower())" it is  better because list append is O(1)
```
### >AI MODEL
         
   ```py     
     def palindrome(text):
    clean_chars = []
    for ch in text:
        if ch.isalnum():
            clean_chars.append(ch.lower())
    clean = "".join(clean_chars)

    left, right = 0, len(clean) - 1

    while left < right:
        if clean[left] != clean[right]:
            return f"It is not a palindrome. Failed at index {left}"
        left += 1
        right -= 1

    return "This is a palindrome"
```
### Reflect on what AI added that you didn't consider initially.

#### How Python Handles Strings Internally
```
I didn’t think about:

Strings being immutable

+= creating a new string each time

The hidden O(n²) cost
```
