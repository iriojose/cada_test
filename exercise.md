
# Cadana Backend Engineer : Coding Exercise

## Task 1: Currency Exchange Rate API

### Objective
Develop an API that receives a JSON request with a currency pair (e.g., `{ "currency-pair": "USD-EUR" }`) and retrieves the exchange rate for this pair.

### Workflow
1. Simultaneously request exchange rates from two distinct external services (Service A, Service B)
2. Return the first response and disregard the second.
3. Expect responses in the format: `{ "currency-pair": rate }`, e.g., `{ "USD-EUR": 0.92 }`.

### Assumptions/Constraints
- Use REST API for communication with external services.
- Simulate these external services within your solution.
- Do not store but fetch and return exchange rates in real-time.
- Each external service requires an API key. Implement a secure method for key management, assuming AWS services access.
- Include tests demonstrating your code's functionality.

### Environment Setup
- Use Go as the programming language.
- Local development environment setup is at your discretion.
- Any IDE or code editor can be used.


### API Key Management
- Utilize best practices for secure API key storage and retrieval with AWS services. Use mocks where applicable to simluate fetching the keys.


### Testing Requirements
- Include both unit and integration tests.
- Use testing frameworks compatible with Go.

### Documentation
- Document your solution with a README file and inline comments where necessary.

### Performance Considerations
- Aim for efficient and responsive API handling.

---

## Additional Task: Data Manipulation and Interfaces in Go

### Objective
Write a Go program that demonstrates data manipulation and interface usage, focusing on object-oriented practices.

### Tasks
1. **Unmarshal and Object Creation:**
   - Unmarshal a JSON file with 10 records in the format: `{ "id": "X", "personName": "Cadanaut X", "salary": { "value": "10", "currency": "USD" } }` into Go objects.
   - Design a `Person` object with appropriate fields and methods to encapsulate the data and operations.

2. **Methods for Data Operations:**
   - Attach methods to the `Person` object to perform the following operations:
     - Sort the array of `Person` objects by salary in ascending and descending order.
     - Group `Person` objects by salary currency using hash maps.
     - Filter `Person` objects by salary criteria in USD, using exchange rates obtained from the primary task. *


### Submission Format
- Submit your solution as a GitHub repository link or a zip file.
- Follow standard Go project directory structures and naming conventions.