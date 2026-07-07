## Objective: Extract structured data from unstructured input.

## 1 . Choose an unstructured text (e.g., "John Doe, age 29, lives in Paris and works as a software engineer.").

### unstructured text
"samuel omafu, age 16, live in london and works as a software engineer."

### 2. Create a template prompt:

"Extract the following fields from the text: Name, Age, Location, Occupation. Return the output in JSON format."

### Validate the output and ensure consistency across multiple inputs. 

*1st input:*

"samuel omafu, age 16, live in london and works as a software engineer."

*1st Output*
```json
{
  "Name": "Samuel Omafu",
  "Age": 16,
  "Location": "London",
  "Occupation": "Software Engineer"
}
```

*2nd input:*
"my name is joseph, 15 , based in china, a pilot" 

*2nd output:*
```json
{
  "Name": "Joseph",
  "Age": 15,
  "Location": "China",
  "Occupation": "Pilot"
}
```
### consistency
Using multiple inputs,From The series of result i got shows that the prompt is well structured, consistent in it output,The AI model had no choice but to think out loud to solve the problem as instructed.