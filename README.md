# ✈️ Flight Tracker

Welcome to **flight-tracker** - a _terminal-based_ air traffic control center built with Go. This project lets you fetch information on airports and track historical arrival and departure flights from specified airports worldwide. Utilizing the OpenSky Network API, you can access comprehensive flight data with a beautiful retro-terminal interface.

---

## Features

- **🌍 Global Airport Search**: Search for airports worldwide using ICAO codes
- **✈️ Flight Tracking**: View historical arrivals and departures from any airport
- **⚡ Smart Caching**: SQLite-based caching system to minimize API calls and improve performance
- **🎨 Terminal UI**: Beautiful, retro-style terminal interface built with tview
- **🔒 Secure Authentication**: OAuth2 token-based authentication with OpenSky Network
- **📊 Real-time Data**: Access to flight data from the previous 24 hours
- **🚀 Fast Performance**: Asynchronous data loading for smooth user experience

## Architecture

**flight-tracker** is built with a modular architecture:

- **API Layer** (`api/`): Handles OpenSky Network API communication and authentication
- **Models** (`models/`): Data structures for airports, flights, and API responses
- **UI Layer** (`ui/`): Terminal-based user interface components using tview
- **Utils** (`utils/`): Caching, data formatting, and utility functions
- **Core** (`main.go`): Application orchestration and state management

## Quick Start

### Prerequisites

- Go **1.24.2 or higher**
- OpenSky Network account and API credentials
- `jq` command-line JSON processor (for token generation)

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/TimKraemer1/go-flight-tracker.git
   cd flight-tracker
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Set up OpenSky Network credentials**

   Create a `credentials.json` file in the project root:

   ```json
   {
     "clientId": "your_opensky_client_id",
     "clientSecret": "your_opensky_client_secret"
   }
   ```

4. **Generate authentication token**

   ```bash
   chmod +x generate-token.sh
   ./generate-token.sh
   ```

5. **Set up environment variables**

   Create a `.env` file (or it will be created automatically by the token script):

   ```env
   TOKEN=your_generated_token
   CLIENT_ID=your_client_id
   CLIENT_SECRET=your_client_secret
   CACHE_PATH=./cache
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

## Usage

### Navigation

- **Search**: Enter an airport ICAO code (e.g., `KSFO` for San Francisco International)
- **Menu Navigation**: Use arrow keys to navigate between menu options
- **View Data**: Select "Airport Information", "Arrivals", or "Departures"
- **Quit**: Press `Ctrl+C` or select quit option

### Features Walkthrough

1. **Airport Search**: Enter a valid ICAO airport code
2. **Airport Information**: View detailed airport information including location, elevation, and coordinates
3. **Arrivals**: See all flights that arrived at the airport in the last 24 hours
4. **Departures**: View all flights that departed from the airport in the last 24 hours

## Project Structure

```
flight-tracker/
├── api/
│   ├── auth.go          # OAuth2 authentication with OpenSky
│   └── opensky.go       # OpenSky Network API client
├── models/
│   ├── airport.go       # Airport data structure
│   ├── flight.go        # Flight data structure
│   └── states.go        # API response structures
├── ui/
│   ├── handlers.go      # UI event handlers
│   └── layout.go        # UI layout and components
├── utils/
│   ├── airports.go      # Airport lookup utilities
│   ├── cache.go         # SQLite caching system
│   └── flights.go       # Flight data formatting
├── main.go              # Application entry point
├── airports.json        # Airport database
├── generate-token.sh    # Token generation script
└── credentials.json     # API credentials (create manually)
```

## 🔧 Configuration

### Environment Variables

| Variable        | Description                   | Required |
| --------------- | ----------------------------- | -------- |
| `TOKEN`         | OpenSky Network access token  | Yes      |
| `CLIENT_ID`     | OpenSky Network client ID     | Yes      |
| `CLIENT_SECRET` | OpenSky Network client secret | Yes      |
| `CACHE_PATH`    | Path to SQLite cache database | Yes      |

### Cache System

The application uses SQLite for intelligent caching:

- **Arrivals Cache**: Stores arrival data per airport
- **Departures Cache**: Stores departure data per airport
- **Automatic Expiry**: Cache expires after 24 hours
- **Rate Limiting**: Prevents excessive API calls

## Technologies Used

- **[Go](https://golang.org/)** - Primary programming language
- **[tview](https://github.com/rivo/tview)** - Terminal UI framework
- **[SQLite](https://www.sqlite.org/)** - Embedded database for caching
- **[OpenSky Network API](https://opensky-network.org/)** - Flight data source
- **[godotenv](https://github.com/joho/godotenv)** - Environment variable management

## 🔑 OpenSky Network Setup

1. Visit [OpenSky Network](https://opensky-network.org/)
2. Create an account
3. Generate API credentials in your account settings
4. Add credentials to `credentials.json`
5. Run the token generation script

## Acknowledgments

- [OpenSky Network](https://opensky-network.org/) for providing comprehensive flight data
- [rivo/tview](https://github.com/rivo/tview) for the excellent terminal UI framework
- The Go community for excellent tooling and libraries

---

_Happy flight tracking! ✈️_
