## Objective: Break complex questions into smaller steps.

### 1. Select a multi-step question (e.g., "What is 25% of the sum of 40 and 60?").

**Selected Question**

"If a number is doubled and then increased by 10 to get 50, what is the number?”

### 2. Create a prompt that asks the model to:

    Restate the problem

    Break it into smaller steps

    Solve each step before giving the final answer

**prompt**

"If a number is doubled and then increased by 10 to get 50, what is the number?”
Restate the problem, break it into smaller steps, solve each step, and then provide the final answer.

### 3. Review the reasoning process and check the accuracy of the final solution.

*Restate the Problem*

We are told that a number is first doubled, then 10 is added to it, and the final result becomes 50. We need to find the original number.

*Break the Problem into Smaller Steps*

Represent the unknown number with a variable (let it be x).

Translate the words into a mathematical equation.

Solve the equation step by step.

Verify the answer to ensure it satisfies the condition.

*Step-by-Step Solution*

*Step 1: Represent the number*
Let the number be x.

*Step 2: Translate the statement into an equation*

“If a number is doubled” → 2x
“and then increased by 10” → 2x+10
“to get 50” → 2x+10=50

*Step 3: Solve the equation*
*Subtract 10 from both sides:*

2x+10−10=50−10

2x=40

Now divide both sides by 2:

x=40÷2

x=20

*Step 4: Verify the answer*

Double the number: 20×2=40
Add 10: 40+10=50
The result matches the condition, so the solution is correct.

*Final Answer*

The number is 20.

### Review of the Reasoning Process

The reasoning process is clear, logical, and well arranged. The problem was first restated in simple terms, which gave clear understanding of what is being asked. Then, the unknown value was correctly represented using a variable (x), which is a proper algebraic approach.

Each solving step followed correct algebraic rules:
Subtracting 10 from both sides to remove the term with x
Dividing both sides by 2 to find the value of x
The reasoning flows step-by-step without skipping operations, making it easy to understand and verify.

### Accuracy Check of the Final Solution
Final answer was:

x = 20

To confirm accuracy, the solution was switch back into the original condition:

Double 20 → 20×2=40

Add 10 → 40+10=50

Since the result equals 50 as required, the final answer is correct.

### Reflection 

This exercise taught me a lot of thing like thinking logically , structuring question in such a way that the AI Model will answer in steps not just going straight to the final answer, using this method of prompt gives you the desired response you want in step ranging from the less difficult step to the final Answer