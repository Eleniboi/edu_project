## Objective: Understand how parameters affect output diversity.

### 1. Use the same prompt with temperature = 0.2 and then with temperature = 0.9.

### prompt used

"Explain programming in simple terms for beginners."

## Experiment 1: Temperature Comparison 
### Temperature = 0.2 
Response:
[Programming is the process of giving clear instructions to a computer so it can perform tasks for you.]

observation:
- The response was more focused 
- The explanation was direct and less creative
- The wording was simple and predictable

### Temperature = 0.9
Response:
[Programming is like teaching a computer how to solve problems by giving it instructions it can understand.
Imagine a computer as a very fast but very literal assistant.
It can do amazing things, but only if you tell it exactly what to do, step by step.]

Observation:
- The response was more creative and expressive 
- more example and analogy were use, for the purpose of understanding 
- Detailed explanation and focused points 

## Experiment 2: Top-p Comparison 
### Top-p = 0.5
Response:
[the art of giving a computer a set of instructions to follow.]

Observation:
- The response was more controlled 
- It stayed focused on the main topic
- No much difference in wording

### Top-p 1.0
Response:
[Think of programming as the art of writing a very detailed "To-Do" list for a brilliant but incredibly literal-minded robot.]

Observation:
- The output was more flexible
- More flexible explanations
- A bit creative output

## Method 
The experiment was conducted by using the same prompt and adjusting the temperature and top-p parameters separately to observe how they affect response diversity, creativity, and focus.
## Conclusion 
From this experiment, I observed that temperature determine and has a strong effect on creativity and randomness. lower temparature(0.2) produced more focused and predictable response while higher temparature(0.9) gave a more diverse answer with analogies.

I also observed that top-p affects the diversity of word choices. A lower top-p(0.5) will produce a more controlled and safer outputs, While top-p = 1 allowed more diverse and flexible responses.

## Reflection 
Through this experiment, I learnt that AI responses can change significantly even when the prompt is still the same , Tuning the temperature increases creativity and randomness, while adjusting top-p affect the diversity of word selection.