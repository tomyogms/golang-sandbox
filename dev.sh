#!/bin/bash
# Development script with live reload
# Automatically uses air if available

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${YELLOW}🚀 Starting development server with live reload...${NC}"
echo ""

# Check if air is available
if command -v air &> /dev/null; then
    echo -e "${GREEN}Using air from PATH${NC}"
    air
elif [ -f ~/go/bin/air ]; then
    echo -e "${GREEN}Using air from ~/go/bin/air${NC}"
    ~/go/bin/air
else
    echo -e "${RED}Error: air is not installed${NC}"
    echo ""
    echo "Install air with:"
    echo "  go install github.com/air-verse/air@latest"
    echo ""
    echo "Or use 'go run' without live reload:"
    echo "  go run cmd/api/main.go"
    exit 1
fi
