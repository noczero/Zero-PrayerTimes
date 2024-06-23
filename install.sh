#!/bin/bash

# Define variables
BINARY_FILE="main"  # Replace with your binary file name
SERVICE_NAME="zero-prayer"  # Replace with your desired service name

# compiled go application
echo "zero-prayer is compiling"
go build -o main ./cmd/main.go
echo "zero-prayer has been compiled, please don't remove main file!"

# Create systemd service file
sudo bash -c "cat > /etc/systemd/system/$SERVICE_NAME.service <<EOF
[Unit]
Description=$SERVICE_NAME service
After=network.target

[Service]
ExecStart=$PWD/$BINARY_FILE
WorkingDirectory=$PWD
EnvironmentFile=$PWD/.env
Restart=always
User=$USER
Group=nogroup

[Install]
WantedBy=multi-user.target
EOF"

# Reload systemd, enable and start the service
sudo systemctl daemon-reload
sudo systemctl enable $SERVICE_NAME
sudo systemctl start $SERVICE_NAME

echo "$SERVICE_NAME has been installed and started."
