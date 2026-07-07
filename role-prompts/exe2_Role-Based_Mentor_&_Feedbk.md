## Objective: Use role-based prompts for personalized feedback.
### 1.Ask the model to act as a coding mentor.
"Act as a mentor reviewing this code. Provide corrections and explain improvements."

### 2. Provide a simple piece of code (e.g., a Python function with a bug).

```py
a = 3 
b := 4
c = b-a

Prin("a*b")
print(c,+b)


```

### 3. Prompt the model: "Act as a mentor reviewing this code. Provide corrections and explain improvements." 

**Correction from AI:**

```py
a = 3
b = 4 
c = b-a 
print(a*b)
print(c+b)
```
**Improvement:**

the model did explicitly well like a real coding mentor, listing out bugs ,thing that are not valid in python language.
it gave me insight on some line i did mistakes like line 2, was corrected to be "b = 4" not "b := 4" this is invalid in python, it also spotted a typo + logic "print(a* b)" instead of this "prin("a*b")"

### 4. Review the response and check if the role influences feedback style.

The mentor persona influenced the feedback style by making it structured, supportive, and educational. instead of just going straight in correcting the code it first list out the bugs in my code why some syntax were not working and steps by step manner to do it right. The mentor role reflected great in it response not just the corrections provided, but also the teaching style and communication approach.