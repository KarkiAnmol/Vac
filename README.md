# Vac

**An AI-driven application providing real-time commentary for Esports events.**

---

## Overview

Vac is an Artificial Intelligence (AI)-powered application that generates commentary for Esports events like CS:GO, Dota 2, and Call of Duty in real-time. It integrates cutting-edge technologies such as GPT-3.5, Go to deliver dynamic, insightful, and engaging commentary. The backend efficiently processes Esports data for generating commentary.

---

## Features

- **Real-Time Commentary Generation**: Generates engaging and contextually relevant commentary during live Esports matches.

---

## Technologies Used

- **Backend**: Go (Golang)
- **AI Model**: GPT-3.5
- **API Communication**: RESTful APIs and WebSocket (via Socket.IO)

---

## Commentary Output Example

![Commentary Generated Output](https://github.com/user-attachments/assets/9b449fe3-676e-49d6-b337-4a4b0e11f4f5)  
*Figure 1: Example of generated commentary output.*

---

## Architecture

The Vac system consists of three major components:
1. **Server1 (Data Preprocessor)**: Handles data ingestion, filtering, and structuring.
2. **Server3 (AI Commentary Generator)**: Integrates with GPT-3.5 for generating commentary based on processed data.

### Flowchart of Data Processing

![Flowchart of Data Processing](https://github.com/user-attachments/assets/1fcc83fb-3ba5-4eb2-b940-c7cd23ae1eab)  
*Figure 2: Data processing flowchart.*

### System Architecture

![System Architecture](https://github.com/user-attachments/assets/65644797-91ab-4593-9ef2-3e12a18337ab)  
*Figure 3: System architecture.*

---

## Prerequisites

### Software
- Go 1.20+

### Hardware
- Minimum Requirements:
  - **Processor**: Intel Core i5
  - **RAM**: 8GB
  - **Storage**: 256GB SSD
- Recommended:
  - **Processor**: Intel Core i7 or equivalent
  - **RAM**: 16GB
  - **Storage**: 512GB SSD

---

## Installation and Setup

1. Clone this repository:
   ```bash
   git clone https://github.com/karkianmol/vac.git
   cd vac
   ```

2. Install dependencies for the servers:
   - **Server1**:
     ```bash
     cd server1
     go mod tidy
     ```
   - **Server3**:
     ```bash
     cd server3
     go mod tidy
     ```

3. Configure the environment variables:
   - Create `.env` files for both servers with necessary configurations:
     ```
     OPENAI_API_KEY=<Your OpenAI Key>
     ```

4. Start the servers:
   - Server1:
     ```bash
     go run server1.go
     ```
   - Server3:
     ```bash
     go run server3.go
     ```

---

## API Endpoints

### Server1
- **POST /process**: Sends processed data to Server3.

### Server3
- **POST /commentary**: Generates and returns AI commentary.

### Commentary API Workflow

![Commentary API Workflow](https://github.com/user-attachments/assets/f028b9ab-8636-447f-a957-474b1a1d60a8)  
*Figure 4: API workflow for commentary generation.*

---

## Data Source

The data used for this project is from the **Grid Esports Data Jam 2023**. You can access the data files [here](https://github.com/grid-esports/datajam-2023/tree/master/data_files). 

---

## How It Works

1. **Data Ingestion**: Server1 preprocesses Esports data in JSONL format, categorizes events, and sends them to Server3.
2. **AI Integration**: Server3 processes the data using GPT-3.5 and generates context-rich commentary.

---

## Future Enhancements

- **Audio Commentary**: Generate AI-driven voice commentary with emotional tone.
- **Game Support Expansion**: Adapt functionality for more Esports titles.
- **Enhanced Personalization**: Provide user-specific commentary settings.

---

## License

This project is licensed under the MIT License.

---
