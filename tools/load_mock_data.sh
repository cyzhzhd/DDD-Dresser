#!/bin/bash

# Default server URL
SERVER_URL="http://localhost:8080"

# Parse command line arguments
while [[ $# -gt 0 ]]; do
  case $1 in
    -s|--server)
      SERVER_URL="$2"
      shift 2
      ;;
    *)
      echo "Unknown option: $1"
      echo "Usage: $0 [-s|--server SERVER_URL]"
      exit 1
      ;;
  esac
done

echo "Loading mock data to server: $SERVER_URL"

# Location of the mock data JSON file
MOCK_DATA_FILE="$(dirname "$0")/mock_data.json"

if [ ! -f "$MOCK_DATA_FILE" ]; then
  echo "Error: Mock data file not found at $MOCK_DATA_FILE"
  exit 1
fi

# Function to check if the request was successful
check_response() {
  local status_code=$1
  local entity_type=$2
  local entity_id=$3
  
  if [[ $status_code -ge 200 && $status_code -lt 300 ]]; then
    echo "Added $entity_type with ID $entity_id successfully"
    return 0
  else
    echo "Failed to add $entity_type with ID $entity_id (Status: $status_code)"
    return 1
  fi
}

# Add brands
echo "Adding brands..."
BRAND_COUNT=$(jq '.brands | length' "$MOCK_DATA_FILE")
for (( i=0; i<$BRAND_COUNT; i++ )); do
  BRAND=$(jq -c ".brands[$i]" "$MOCK_DATA_FILE")
  NAME=$(jq -r ".brands[$i].name" "$MOCK_DATA_FILE")
  
  echo "Adding brand $NAME..."
  
  STATUS_CODE=$(curl -s -o /dev/null -w "%{http_code}" \
    -X POST \
    -H "Content-Type: application/json" \
    -d "$BRAND" \
    "$SERVER_URL/api/brands")
  
  check_response "$STATUS_CODE" "brand" "$NAME"
done

# Add products
echo "Adding products..."
PRODUCT_COUNT=$(jq '.products | length' "$MOCK_DATA_FILE")
for (( i=0; i<$PRODUCT_COUNT; i++ )); do
  PRODUCT=$(jq -c ".products[$i]" "$MOCK_DATA_FILE")
  NAME=$(jq -r ".products[$i].name" "$MOCK_DATA_FILE")
  
  echo "Adding product $NAME..."
  
  STATUS_CODE=$(curl -s -o /dev/null -w "%{http_code}" \
    -X POST \
    -H "Content-Type: application/json" \
    -d "$PRODUCT" \
    "$SERVER_URL/api/products")
  
  check_response "$STATUS_CODE" "product" "$NAME"
done

echo "Mock data loading completed!" 