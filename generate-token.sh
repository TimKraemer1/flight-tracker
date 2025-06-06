#!/bin/bash

CREDENTIALS_JSON="credentials.json"
ENV_FILE=".env"

# Check dependencies
if ! command -v jq &> /dev/null; then
    echo "'jq' is not installed. Install it with: sudo apt install jq (or brew install jq)"
    exit 1
fi

# Use an array (without commas)
KEYS=("clientId" "clientSecret")

export CLIENT_ID=""
export CLIENT_SECRET=""

# Check if the JSON file exists
if [[ ! -f "$CREDENTIALS_JSON" ]]; then
    echo "JSON file not found: $CREDENTIALS_JSON"
    exit 1
fi

# Loop through keys and extract values
for KEY in "${KEYS[@]}"; do
    VALUE=$(jq -r --arg key "$KEY" '.[$key]' "$CREDENTIALS_JSON")

    if [[ "$VALUE" == "null" ]]; then
        echo "Key '$KEY' not found in $CREDENTIALS_JSON — skipping"
        continue
    fi

    # Assign to exported variables
    if [[ "$KEY" == "clientId" ]]; then
        export CLIENT_ID="$VALUE"
    elif [[ "$KEY" == "clientSecret" ]]; then
        export CLIENT_SECRET="$VALUE"
    fi
done

export TOKEN=$(curl -X POST "https://auth.opensky-network.org/auth/realms/opensky-network/protocol/openid-connect/token" \
    -H "Content-Type: application/x-www-form-urlencoded" \
    -d "grant_type=client_credentials" \
    -d "client_id=$CLIENT_ID" \
    -d "client_secret=$CLIENT_SECRET" | jq -r .access_token)

echo "CLIENT_ID=$CLIENT_ID" >> .env
echo "CLIENT_SECRET=$CLIENT_SECRET" >> .env
echo "TOKEN=$TOKEN" >> .env