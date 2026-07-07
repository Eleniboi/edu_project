## Objective: Combine multiple prompts to simulate a tool-using agent.

### 1. Step 1 prompt: Ask for a user’s query (e.g., “What’s the temperature in Paris?”).

prompt:
"what would you like to know?"

User Input: 
"What's the temperature in paris" 


### 2. Step 2 prompt: Simulate a call to the weather API.
API CALL:

"WeatherAPI.getCurrentWeather(location= "Paris")"

SIMULATED API RESPONSE:
this output is not shown to the user.
```json
{
    "location": "paris",
    "temperature":"8°c",
    "condition":"cloudy"
}
```
### 3. Step 3 prompt: Format the final answer back to the user.
the agent converts structured data into natural language.

FORMATTED RESPONSE:
The current temperature in Paris is 8°C with cloudy conditions.

### 4. Document the chained process and outputs.

1. user intent: User provides a natural language query. 
2. Entity Extraction: Agent identifies "Intent: Get_weather" and "Parameter: Paris"
3. Tool Selection: Agent selects WeatherAPI from its toolkit.
4. Observation: Agent receives a JSON payload. 
5. Synthesis: Agent merges the payload with conversational context to produce a friendly response.