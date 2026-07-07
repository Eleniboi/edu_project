## Goal: Practice the "right way" to use AI for learning on a technical topic (wireless routing protocols) — attempt first, then use AI to probe and deepen.

**self explanation:**

WRP is a routing protocol that uses multiple tables to select the best path for data transmission, communicate with neighboring nodes, receive updates about path changes, and ensure reliable and safe packet delivery in a wireless network.

**Phase 1: Build Foundation (you first)**
**Scenario Design:**

Node A is connected to B
Node B is connected to C and D
Node D is connected to E 
Node E is connected to F

This forms a multi-hop wireless network of six devices where nodes communicate with their immediate neighbors. 

**Justification:**

WRP is suitable for this network because it operates efficiently in multi-hop wireless environments where nodes rely on neighbors to forward data. since A cannot directly reach F, WRP uses its multiple routing tables to determine the best path(A-B-D-E-F), track path changes, and quickly update routes

## Phase 2: Strategic AI Use
1. **understanding Test Questions:**

- **what if D link fail?**

*failure occurs*
failure occurs , router B tries to send packet to D, D does not respond (no acknowledgement). WRP immediately notices something is wrong.

*failure dection (very fast)*
WRP uses:

update messages, acknowledgements(ACK), message Retransmission List(MRL)

if acknowledgements stop coming from D: 
WRP marks the B-D link as failed.

once failure is detected:
WRP updates its distance table, updates routing table, Removes the broken path (B-D-E-F)

now WRP checks:
is there another possible path to F?
for example, if C can reach E:
new route could be:
A-B-C-E-F
if no alternative path exists:
packet transmission is temporarily stopped , Network searches for new valid route 


**Does WRP store backup paths?**

*WRP stores detailed routing information in multiple tables, which allows it to:*

quickly recompute alternative paths, Not blindly to a stored backup


**How does WRP detect failure?**

WRP detect failure using three mechanisms:

- *Acknowledgement system(ACK)*

Every update message expects a reply.
if no reply from node D - Failure suspected.

- *message Retransmission List (MRL)*

WRP keeps track of:
sent messages 
received acknowledgement 
If a message is unacknowledged for too long:
the link is considered broken.

*Neighbor updates stop:* 

WRP constantly communicates with neighboring nodes (like B-D).
if neighbor updates stop:
Node is marked unreachable.



2. **Edge Cases Exploration:**

i checked for the edge case where the link between B and D fails during transfer in a multiple wireless network. then i understand how WRP works in different environments where node connection change rapidly and routes may suddenly change to avoid lost of path.

3. **validation of Concepts:**

Through AI probing(inquiry), i comfirmed that WRP does not simply store backup routes but uses multiple routing tables and acknowledgements to detect failures and recalculate the best paths. this corrected my earlier assumption that WRP instantly switches to a stored backup route.

## Phase 3: Real Application

## 1. Network Overview (Scenario Design)
**Smart-city Traffic Hub Description** 
The network represents a smart traffic control system consisting of :

- 1,000 IoT sensors (road sensors, cameras, pollution sensors)

- 50 smart traffic lights
- 10 emergency vehicles (ambulance, fire trucks, police)
- 1 central traffic control hub

these devices form multi-hop wireless network where each device acts as a node and communicates with nearby nodes.

## protocol choice and justification(WRP) 
**selected Protocols: WRP (wireless Routing Protocol)**
WRP is suitable for the traffic hub network because 
- 1. it supports multi-hop communication 

since most sensors cannot directly reach the central hub, WRP allows data to pass through neighboring nodes (e.g sensor - Traffic light- Hub)

- 2. it uses multiple routing tables
WRP maintains:
- Distance Table
- Routing Table
- link cost Table
- message retransmission list(MRL)
this improves route accuracy and reliability.
 3. Fast failure detection 

in a smart city, traffic must be real - time. 
WRP detects broken links quickly using acknowledgement (ACK). 
 4. Reliable packet delivery

Traffic systems cannot afford data loss (e.g accident alerts). WRP ensures safe delivery through message tracking and retransmission. 
 5.  Neighbor awareness 

 Each node constanly communicates with neighbors, making it ideal for dynamic environment like moving emergency vehicles.

 ## 3. How the Network Works in Real Life (Flow)

- **Real-world data flow example:**

- A road sensor detects traffic congestion

- Sensor sends data to nearest traffic light

- Traffic light forwards data to Traffic Hub

- Hub adjusts signal timing

- Emergency vehicles receive priority route updates

## 4.Failure Points (Critical Analysis)
**Possible Network Failures**

- Link failure between nodes (e.g., Traffic Light B disconnects)

- Node failure (sensor battery dies)

- Network tangle (too many sensor updates)

- Mobility problem (emergency vehicles moving fast)

- Wireless hindrances (buildings, weather)

## 5. How WRP Handles Failures (Refinement)
**Case 1: Traffic Light Link Failure**

*If Traffic Light B fails:*

- Neighbor nodes stop receiving Acknowledgement(ACK)

- WRP  notice missing acknowledgement

- Broken link is removed from routing table and an Alternative route is been recomputed automatically

**Case 2: Emergency Vehicle Mobility**

*Since vehicles move frequently:*

- WRP constantly updates neighbor information

- Routing tables adapt differently

- New best path is selected in real time

**Case 3: No Available Backup Path**

*If all routes fail:*

- WRP partially stops packet forwarding

- Network searches for a new valid path

- Once discovered, routing resumes

## 6. Edge Case Exploration (Advanced Thinking)

**Edge Case 1:** several Node Failures
If multiple sensors fail at once, WRP recalculates routes using available neighbors without breaking the whole network.

**Edge Case 2:** High Traffic Data Load
WRP manages congestion using structured routing tables instead of random flooding, improving efficiency.

**Edge Case 3:** Partial Network separate
If part of the city network disconnects, WRP isolates the failure zone while maintaining communication in the rest of the network.

## 7. AI Feedback Integration (Learning Refinement)

AI helped to:

- Test my understanding of WRP behavior during link failures

- Help me check for edge cases like mobility and congestion

- verify protocol suitability for smart-city environments

- strengthen conceptual clarity beyond textbook reading

But the initial explanation, scenario design, and justification were self developed before AI validation 

## Reflection:

**1. % Human Judgment vs AI Contribution**

- Human judgement: 70%
I first attempted the , Scenario design, explanation, reasoning 
- AI Contribution: 30% 
AI helped me with validation, deeper probing, edge case exploration 

**2. Could you defend decisions without AI?**

yes, i understand how WRP uses multi-hop routing , to detect fast failure occurances, select the best path for data, and reliable packet delivery, which are important for real-time smart city traffic systems.

**3. What will you still remember in 6 months?**

In 6 months i will still remember that WRP has four(4) routing table, which are :
- Distance Table: which store details of distance from each destination 
- Routing Table: stores the best path to each destination or routes
- link cost Table: it store how expensive a path is 
- message retransmmision list(MRL): it tracks message sent, acknowledgement recieved and prevent lost of path.
- it recomputes routes instead of blindly storing backups. 

**Did AI make you sharper, or think for you?**
yes, it actually made me sharper and help me also think critically in a logical way, using AI strategically brought out the best in me , help me try myself first and later go deeper in understanding using AI as a learning amplifier.