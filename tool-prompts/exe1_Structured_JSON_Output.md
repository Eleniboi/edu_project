## Objective: Create prompts that produce structured machine-readable output.

### 1. Write a prompt that asks for unstructured information (e.g., “List three cities with their populations.”).
UNSTRUCTURED INFORMATION:

“List three cities with their populations."

**Response**

Tokyo - 14 million
New York - 8.5 million
Paris - 2.1 million

### 2. Rewrite the prompt to enforce structured JSON output:

    Example: “Return the following in JSON format: [{city: string, population: number}].”

### 3. Test with multiple inputs and verify the JSON stays valid.

**1st input**

List three European cities with their populations. Return valid JSON only.
```json
[
    {"city": "Berlin", "population": 3645000},
    {"city":"Madrid","population":3223000},
    {"city":"Rome","population":2873000}
]
```
**2nd input**

“Provide three Nigerian state capitals, including Makurdi.” beoutputing one of them. return in valid JSON only
```json
[
    {"city":"Makurdi","population":370000},
    {"city":"Abuja","population":1235880},
    {"city":"Lagos","population": 14160000} 
]
```
The two schema were maintained in the same order that shows a valid JSON