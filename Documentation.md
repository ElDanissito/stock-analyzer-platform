# Stock Analyzer Platform - Documentation

## 📋 Problem Definition

Your task is to design and implement a system that retrieves stock information from a given API. The system should be able to handle different types of stock and display the relevant information to the user in a user-friendly interface.

---

## 🛠️ Technology Stack

### Backend
- **Language:** Golang

### Frontend
- **Framework:** Vue 3
- **Language:** TypeScript
- **State Management:** Pinia
- **Styling:** Tailwind CSS

### Database
- **Database:** CockroachDB

---

## 🔌 API Documentation

### Endpoint

```
GET https://api.karenai.click/swechallenge/list
```

### Query Parameters

| Parameter | Description |
|-----------|-------------|
| `next_page` | The key to start the next page |

### Authentication

Include API key in the `Authorization` header:

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6MjQ3LCJlbWFpbCI6ImRhbmkucjM0M0BnbWFpbC5jb20iLCJleHAiOjE3NjAyNTA5OTEsImlkIjoiIiwicGFzc3dvcmQiOiJ1c2VybmFtZS8qKi8gRlJPTS8qKi8gdXNlcnMgLS0ifQ.Zo-JdRJJMJO7kCaLvgDW2hN05_gVKc5_9zWLuIR1T0o
```

---

## 📝 Implementation Steps

### 1. Connect to the API and Store the Data

To begin, you will need to connect to the API that provides the stock information. This will involve:

- Making HTTP requests to the API endpoints to retrieve the necessary data
- Handling errors appropriately
- Ensuring that the data is properly formatted for use in the UI
- Storing the retrieved data in CockroachDB

### 2. Create a Simple API & UI

Once you have successfully connected to the API and stored the stock information, you will need to create a user interface to display this information. 

**Requirements:**
- The UI should be intuitive and easy to use
- Allow users to easily navigate the stock information
- Implement search functionality
- Implement sorting capabilities
- Display detailed views of individual stocks

### 3. Recommend the Best Stocks to Invest Today

For the final part of the challenge, you will need to develop an algorithm that analyzes the stock data and recommends the best stock to invest in.

**Note:** Feel free to enrich the data with external sources if you want to improve the recommendation algorithm.

### 4. Write Unit Tests for Your Code

We highly encourage you to write unit tests for your code. This will ensure the reliability and stability of your application.

**Testing should cover:**
- API integration
- Data processing logic
- Recommendation algorithm
- UI components

---

## 🚀 Getting Started

_To be added: Setup and installation instructions_
