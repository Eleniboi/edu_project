
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