## Objective: Practice designing prompts that mimic function calls.

### 1. Create a prompt: “Simulate a weather API response for Paris today.”

“Simulate a weather API response for Paris today.”

### 2. Define an expected JSON schema, e.g.:
```json
{
    "location": "Paris",
    "temperature_celsius": 18,
    "condition": "Cloudy"
}
```
DEFINING EXPECTION:

"Simulate a weather API response for Paris today. Return the response in JSON format like the example below:"
```json
{
  "location": "string",
  "temperature_celsius": "number",
  "condition": "string"
}
```
### 3. Test the model with different cities and compare outputs.

### Test 1 : Paris 
### Prompt:

"Simulate a weather API response for Paris today. Return the response only in JSON format like the example below."
```json
{
  "location": "string",
  "temperature_celsius": "number",
  "condition": "string"
}
```

**Output 1: Paris**
```json
{
  "location": "Paris",
  "temperature_celsius": "18",
  "condition": "cloudy"
}
```
### Test 2: Tokyo
### Prompt:
"Simulate a weather API response for Benue today. Return the response only in JSON format like the example below."

```json
{
  "location": "string",
  "temperature_celsius": "number",
  "condition": "string"
}
```
**Output 2: Benue**
```json
{
    "location":"Benue",
    "temperature_celsius": 22,
    "condition": "sunny"
}
```

COMPARISON:

All the JSON output are valid and follow the expected format. it provided the schema in such a way the mechine can read it.