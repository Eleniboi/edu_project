# Part A: The Critical Distinction 

## Write down your honest answers: 

● I have always use AI for learning new skill build critical thinking  

● No, I always take my time trying some possible solution on my own  

● Yes, because I participated actively doing the solution  

● Engage myself in critical thinking as I will always do 

**2. Identify your current pattern:** Which learner are you now? 

**Learner B:** "AI is my learning amplifier" 

**3. Write a brief paragraph:** Where are you now, and where do you want to be?  

● Am still a beginner programmer trying to finish tutorial, memorizing code, join any possible code class , but in the nearest future i want to be in a situation where i will be solving real problem, writing code that add position value in the tech world.

## Part B: The Wrong Way vs. The Right Way 
### Track B — The Right Way (DO THIS): Step 1: Attempt independently 
**1. Write pseudocode for a palindrome check**
```
 FUNCTION palindrome(text):
    creat an empty container "clean"
    check each character in text: 
               Add the character to "clean"
                reverse the clean version
                Compare if the reverse varsion is == to the clean version
```
### my palindrome code 

```PY


def palindrome(text): #this is the function define
    clean = ""  # creat an empty string "clean"
    for ch in text: # LOOP through the text
        if ch.isalnum(): # if the characters are alphabet or number
            clean += ch.lower() # Turn the selected character to lower case
    reverse = clean[::-1] # reverse the clean version 
    if reverse == clean: # compare if the reverse is == to the clean
        return True 
    else:
        return False 
print(palindrome("racecar"))
print(palindrome("hello"))
print(palindrome("A man a plan a canal Panama"))
```
### Step 2: Strategic AI use After you have a working solution, ask AI:** 

"I wrote a function to check palindromes. Here's my approach: [paste YOUR code] 

● What's the time complexity? 

    O(n²) (because of repeated string concatenation) 

● What edge cases am I missing? 

    - Edge cases like empty string. Example “palindrome (“”) will always return true” 

    - Only punctuation will always return true e.g. “palindrome (“!!!”)” 

●  Alternatives and trade-offs? 
```py
   def palindrome(text):
        clean = []
        for ch in text:
           if ch.isalnum():
              clean.append(ch.lower())
        clean = ''.join(clean)
        return clean == clean[::-1]
```
●  Time Complexity: 

    O(n) total 

    ''. join() is linear 
●  Performance on Very Long Strings

 **Using +=**

- May take dramatically longer

- Memory fragmentation increases

- Can cause noticeable slowdown

**Using list + join**
- True O(n)

- Much faster

- Memory allocated once

- Industry best practice

### Step 3: Reflection 

●  What did you learn by struggling first? 
```
I learnt how to think logically before attempting a question, trying problem and solution yourself make you think wildly, make you understand in your own way 
```
● How is your understanding different than if you'd just asked for the solution? 
```
If I would’ve just asked for the solution, I won’t have gotten deep understanding of the concept  
```
● Can you now implement similar functions (reverse a string, find duplicates) without AI? 
```
Yes, I can, using various way to arrive at the same solution  
```
● What mental model did you build? 
```
Alway try yourself first, verify later 
```
### Part C: Testing Your Understanding 

Without using AI, complete these variations: 

1. Ignore spaces and punctuation 

2. Make it case-insensitive 

3. Return the position where the string stops being a palindrome (if not a palindrome) 

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
### Then use AI strategically:
```py
def palindrome(text):
    left, right = 0, len(text) -1

    while left < right:
        while left < right and not text[left].isalnum():
            left += 1
        while left < right and not text[right].isalnum():
            right -= 1
    
        if text[left].lower() != text[right].lower():
            return f"Not a palindrome (failed at index {left})"
        left += 1
        right -= 1
    return "palindrome"
```
For more efficient pattern use ;   clean = ‘ ’.join(ch.lower( ) ) instead of using clean += ch.lower() 

 why because strings in python are immutable in python using clean += ch will always,
 create new string
 copy old contents 
 add new char 
 reassign
 That’s why repeated concatenation can become O(n²).

 ## Part D: The Fairness Contract 

 

Create ex00_ai_ethics_coder.md with a personal code of ethics for AI use in learning: 

 

1. I will use AI when: 

- After I have observed my 15-minute rule of critical thinking 

- I want to know why and how a particular concept works 

- To try other possible ways after my solution works properly 

- [I can boldly explain why a concept work] 

2. I will NOT use AI when: 

- I have not tried to find possible solutions myself  

- I’m in an interview 

- Am still learning basics on a particular concept 

- [I have no mental model of my own] 

3. I know I'm using AI fairly when: 

- I can explain my code with ideas that work without the help of AI 

- I could solve similar problems without AI  

- I feel more affirm in my own understanding 

- [I am using it as a learning tool but not an answer generator] 

**Sign: omafu oloche samuel – 17th February 2026**


## Part E: Real-World Scenario Analysis 

1. Interview: You're asked to explain a caching system design. What happens if you've always relied on AI to design systems? 
 

- *ANSWER 1:* if i have always relied on AI, would want to ask AI that moment  
 

2. Production bug at 2 AM: You must debug code generated by AI months ago - but AI is offline. Can  
you fix it? 
 

- *ANSWER 2:* Yes, because I will not be so dependent on AI, but if i was i won't be able to fix that  
paticular problem at that moment 
 

3. New tech with little documentation: A new library is released, and AI hasn't been trained on it yet. How do you learn it? 
 

- *ANSWER 3:* I will learn through documentation and learning aid like watching video online apply pear to pear learning 
 

- *Reflection:* when I use AI fairly prepare me for unknown bugs, make me think critically, bring idea that really work and solve problem without the use of AI 
 

 ## Part F: Building Irreplaceable Skills 

 **Rate yourself 1–5 and write an improvement plan for your lowest area:**



| **Skill**                | **Description**                           | **Rating** | **Improvement Plan**                                                                                                                               |
| ------------------------ | ----------------------------------------- | ---------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| Problem Decomposition    | Breaking down problems logically          | 2/5        | Spend more time understanding the problem step by step before attempting to solve it. Break large problems into smaller sub-problems first.        |
| Systems Thinking         | Understanding how components interact     | 2/5        | Take time to understand how components interact and what could happen if one part fails. Think about cause and effect within the system.           |
| Critical Evaluation      | Knowing when code is wrong or inefficient | 2/5        | Identify potential bugs early and analyze inefficiencies to prevent crashes or poor performance. Review code more critically before finalizing it. |
| Debugging Mindset        | Investigating unexpected behavior         | 4/5        | Continue improving error-tracking skills and strengthen systematic debugging when unexpected outputs occur.                                        |
| Conceptual Understanding | Knowing WHY, not just HOW                 | 3/5        | Focus on understanding what happens behind the scenes and what could go wrong if certain logic or structures are not used correctly.               |


### Action plan: 3 specific actions this week to improve it without outsourcing thinking to AI. 

1. **problem Decomposition:**

I will make sure I break down questions into pieces for more understanding, what the problem is all about  and why  the problem occurs, and also think of an ethical way to solve the problem

2. **Systems Thinking:**

Learn to know what is happening behind the scenes, how anything i do affects something else, writing code that can build but not to distroy 

3. **Critical Evaluation:** 

Considering edge cases in my programs and time complexity, find possible potential bugs to prevent program from crashing