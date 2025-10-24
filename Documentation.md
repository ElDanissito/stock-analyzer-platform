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

## � Data Model Overview

The platform ingests analyst actions and price targets for tickers. Core fields:

- Ticker, Company
- Brokerage (free text)
- Action (e.g., raised, upgraded, reiterated, initiated)
- Rating From / Rating To (e.g., Sell, Hold/Neutral, Buy/Strong Buy)
- Target From / Target To (string amounts as provided by the source)
- Time (event timestamp)

---

## 📝 Implementation Steps

### 1. Connect to the API and Store the Data

To begin, you will need to connect to the API that provides the stock information. This will involve:

- Implementing a pull-based data ingestion from a provider (credentials managed via local environment variables)
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

See the project READMEs for setup flow. High level:
- Start database services (CockroachDB)
- Initialize the backend schema and run the backend service
- Start the frontend
- Use the UI to trigger data sync. You can choose how many pages to load (default 20; max 100). No provider endpoints or keys are stored in this repository.

Security note: never commit endpoints or API keys. Configure any credentials locally via environment variables and .env files that you keep outside of version control.
