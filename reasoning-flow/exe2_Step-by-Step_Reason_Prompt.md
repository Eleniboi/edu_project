## Objective: Use structured prompting to encourage reasoning.

### 1. Choose a math or coding problem (e.g., "Calculate the factorial of 5").

**problem**

 "find the factors of 20"

### 2. Prompt the model: "Let's think step by step before answering."

**prompt**

"Let's think step by step before answering."

### 3. Compare outputs with and without the reasoning instruction.

**with the reasoning instruction.**

Step 1: Start with 1
Every number is divisible by 1 and itself. Since 1×20=20, our first pair is 1 and 20.

Step 2: Check 2
Since 20 is an even number, it is divisible by 2. 2×10=20. So, our next pair is 2 and 10.

Step 3: Check 3
If you add the digits of 20 (2+0=2), the result isn't divisible by 3. Also, 3×6=18 and 3×7=21. So, 3 is not a factor.

Step 4: Check 4
We know that 4×5=20. This gives us the pair 4 and 5.

Step 5: Check 5
We just used 5 in the previous step. Since we have reached a number we already listed, we know we have found all possible factors!

The Final List
The factors of 20, in order from least to greatest, are:
1, 2, 4, 5, 10, and 20.

**without the reasoning instruction.**

1×20=20: 1 and 20 are factors.

2×10=20: 2 and 10 are factors (20 is even, so 2 always works).

Check 3: 20÷3=6.66..., so 3 is not a factor.

4×5=20: 4 and 5 are factors.

Check 5: Since we already listed 5 in the previous step, we have found all possible factors.

### Comparison between outputs with and without the reasoning instruction.
with the reasoning instruction ("Let think step by step before answering") is included, the response becomes more structured and instructional. The Model assume no previous knowledge on that topic, Solve problem with beginner friendly words, solved the problem in steps starting from the easiest to ensure the learner build knowledge, and gave the final answer with reason.

Without the reasoning instruction, the response is more direct and compressed. Although it still checks the numbers, the explanation is less guided and assume the learner has previous knowledge. The steps are present, but they are not framed as a logical progression, and the explanation is slightly less detailed.