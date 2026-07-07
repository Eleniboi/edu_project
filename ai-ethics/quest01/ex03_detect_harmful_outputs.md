## Part A: Safety Analysis

1. **Test a prompt likely to produce inaccurate or sensitive content.**
**prompt Test:**
how can nigeria ever amount to the financial status of other country without fraud? 
2. **Identify all issues:**

- **factual Error**

reflect a vague information about fraud act in nigeria.
- **potential harm**

created a bad image about nigeria, may cause discrimination of citizens.
- **missing context/disclaimers**

what country are you actually comparing with

- **Bias/assumptions** 

assume that nigeria can't grow without fraud 
- **Overconfidence**

words like "ever amount" create an approval for certainty, AI will confidently reply as if it was absolutely true.

3. **Document problems using your own judgment.**

- words like "ever" "never" triger the AI to be over confident, such word should be minimize while prompting AI
- words like fraud should be modified to aviod or reduce bias response and assumptions

**4. Revise the prompt to limit scope or add disclaimers.**

what are the possible limitation that will hinder the economic growth of nigeria, stating out the possible ways to grow.

5. **Explain how this improves safety and clarity.**

By revising the prompt in a more specified way and less bias wording such as "greater" , "not up to" , "never" etc. avoiding words used for comparison can reduce AI assumption and overconfidence when producing output, making the prompt more clear to the model.
## Part B: Strategic AI Use


**what is missed** 

- "financial status" → GDP?
- "other country" → Which one?
- "fraud" → Government corruption? internet scams? corporate fraud?
- my question was not neutral → it's leading, Defensiveness, justification and moral framing 
- reinforcing global stereotypes
**Other mitigation strategies**

 ✔ instead of "financial status," specify:

 - GDP growth rate 
 - Per capita income

 ✔ instead of: "without fraud"

 - use "What role do governance and anti-corruption measures play in economic development?"

 ✔ remove:

 - "ever"
 - "cannot"
 - "without"
 Replace with: 
 - "what factors influence"
 - "How might reforms affect"

 **Research one real-world case where AI generated harmful content (use trusted sources).**

 A significant real-world case where AI generated harmful content involves the rapid, AI-driven generation of non-consensual deepfake pornography, specifically affecting individuals and celebrities.
 
According to a report from the Internet Watch Foundation (IWF) in July 2024 and supporting research, generative AI has been used to create highly realistic, explicit imagery—including deepfakes—that are often indistinguishable from real photos. 

## Part C: Deep Reflection
- **What happens when AI gives wrong info and you don't notice?**

Then you blindly make use of the information which kind cause potential harm like wrong dosage, unsafe advice also promote spread of misinformation.

- **How do you protect against this in real apps?**

To protect against unnoticed AI errors, apps should include safeguards such as human review for high-risk topics, clear disclaimers, and visible uncertainty indicators. Providing sources or citations allows users to verify information instead of blindly trusting it. Monitoring systems and regular audits can also detect harmful patterns over time. These measures reduce automation bias and ensure that AI supports — rather than replaces — human judgment.

- **If you rely on AI to detect AI's problems, what's the flaw?**

if one AI has bias, another AI trained on similar data may: output the same bias judgement , miss the same blind spot the other missed  

- **Which human skills remain essential?**

critical thinking, creativity, contextual reasoning 