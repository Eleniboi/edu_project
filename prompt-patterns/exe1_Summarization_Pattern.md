## Objective: Create a reusable prompt for text summarization.

### 1. write a weak prompt: 

"summarize this text."
### 2. Refine it into a reusable pattern:
 
"Summarize the following text in 3 bullet points, focusing on the main arguments and ignoring examples."
### 3. Test it with at least two different inputs (e.g., a news article and a technical blog).


## Example 1: News Article 
("The city council is proposing a 'Congestion Fee' for drivers entering the downtown core during rush hour. Proponents argue that this will significantly reduce carbon emissions and decrease traffic gridlock, making the city more walkable. They point to London and Stockholm as cities where similar fees have successfully cleared the streets. However, local business owners are concerned that the fee will deter shoppers. For instance, a local bakery owner, Maria Gomez, claims her morning sales dropped by 20% during a temporary road closure last year. Opponents also argue the tax is 'regressive,' unfairly hitting low-income workers who have no choice but to drive, such as construction workers or night-shift janitors who cannot rely on the current bus schedule.")

**prompt** "Summarize the following text in 3 bullet points, focusing on main arguments and ignoring examples "

### Summary 1

Environmental and Urban Benefits: Proponents argue the fee will lower carbon emissions, alleviate traffic congestion, and improve the city's overall walkability.

Economic Impact on Business: Critics express concern that the additional cost will discourage consumers from visiting the downtown core, potentially harming local commerce.

Equity and Accessibility Concerns: Opponents claim the fee is a regressive tax that disproportionately burdens low-income workers who lack flexible schedules or viable public transit alternatives.

### Example 2:The Technical Blog 
("Adopting 'Edge Computing' is becoming essential for modern IoT networks. The primary argument for this shift is the drastic reduction in latency; by processing data closer to the source, devices can react in real-time without waiting for a distant cloud server. For example, a self-driving car needs to process a 'stop' command in milliseconds, not seconds. Secondly, Edge Computing improves data privacy by keeping sensitive information on the local device. A smart thermostat, for instance, doesn't need to send your daily schedule to a central database to function. Finally, it reduces bandwidth costs for companies. Instead of streaming 24/7 video footage from a security camera to the cloud, the 'Edge' device only sends a clip when it detects actual movement.")

### Summary 2 
Minimized Latency: Processing data near its source allows for near-instantaneous response times, eliminating the delays associated with communicating with remote cloud servers.

Enhanced Data Privacy: By keeping data localized, Edge Computing reduces the amount of sensitive information transmitted to central databases, narrowing the window for potential data exposure.

Optimized Bandwidth Usage: The technology lowers operational costs by filtering data locally and only transmitting essential information, significantly reducing the volume of traffic sent over the network.