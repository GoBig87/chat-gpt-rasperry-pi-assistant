#!/bin/bash

# Install Golang
go_desired_version="1.21.5"
go_installed_version=$(go version | awk '{print $3}' | cut -c 3-)

if [ "$go_installed_version" == "$go_desired_version" ]; then
  echo "Golang is already installed with version $go_desired_version"
else
  echo "Golang version $go_desired_version not found, installing..."

  # Download and install Golang
  wget https://golang.org/dl/go$go_desired_version.linux-armv6l.tar.gz
  sudo tar -C /usr/local -xzf go$go_desired_version.linux-armv6l.tar.gz
  echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
  echo 'export GOPATH=$HOME/go' >> ~/.bashrc
  echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
  echo 'export PATH=$GOBIN:$GOROOT/bin:$PATH' >> ~/.bashrc

  # Clean up downloaded archive
  rm go$go_desired_version.linux-armv6l.tar.gz
  echo "Golang version $go_desired_version installed successfully"

  source ~/.bashrc
fi

# Install gcloud
if which gcloud &> /dev/null; then
  echo "gcloud is already installed. Skipping installation."
else
  echo "Installing gcloud"
  echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] https://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list
  sudo apt-get install apt-transport-https ca-certificates gnupg
  sudo apt-get update && sudo apt-get install google-cloud-sdk
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key --keyring /usr/share/keyrings/cloud.google.gpg add -
  sudo apt-get update && sudo apt-get install google-cloud-sdk
fi

# Install dependencies
echo "Installing golang dependencies"
go mod vendor

# Build targets
declare -a targets=("gptapi" "gptapp" "gptctl")

for target in "${targets[@]}"; do
  if which $target &> /dev/null; then
    echo "$target is already built. Skipping build."
  else
    echo "Installing $target"

    # Change to cmd/$target directory and run make build
    cd "cmd/$target"
    make build

    # Check if make build was successful
    if [ $? -ne 0 ]; then
      echo "Error: failed build $target. Exiting."
      exit 1
    fi

    # Return to the top-level directory
    cd ../..
  fi
done

# Build /var/lib/gpt/config.env file
env_file="/var/lib/gpt/config.env"
sudo mkdir -p "/var/lib/gpt/"
chown -R "$USER:$USER" "/var/lib/gpt/"

# Check if the file exists
if [ ! -f "$env_file" ]; then
  echo "File $env_file does not exist, creating it..."
  touch "$env_file"
fi

# Check if the porcupine access key variable is present in the file
porcupine_key="PORCUPINE_ACCESS_KEY"
if ! grep -q "^$porcupine_key=" "$env_file" || [ -z "$(grep "^$porcupine_key=" "$env_file" | cut -d'=' -f2)" ]; then
  echo "No Picovoice access key found in $env_file.  Please create one at https://picovoice.ai/console/ and enter it below."
  chromium-browser "https://console.picovoice.ai/"  2>/dev/null &

  # Prompt user for the access key
  read -p "Enter Picovoice key value: " user_input

  echo "$porcupine_key=$user_input" >> "$env_file"
  echo "Added key to $env_file"
fi

chat_gpt_api_endpoint_variable="CHAT_GPT_API_ENDPOINT"
default_chat_gpt_api_endpoint="https://api.openai.com/v1/chat/completions"
# Check if CHAT_GPT_API_ENDPOINT is present in the file
if ! grep -q "^$chat_gpt_api_endpoint_variable=" "$env_file"; then
  echo "The $chat_gpt_api_endpoint_variable variable is not set in $env_file."
  echo "$chat_gpt_api_endpoint_variable=$default_chat_gpt_api_endpoint" >> "$env_file"
  echo "Added $chat_gpt_api_endpoint_variable to $env_file with default value."
fi

# Check if the open ai access key variable is present in the file
openai_key="CHAT_GPT_API_KEY"
if ! grep -q "^$openai_keyy=" "$env_file" || [ -z "$(grep "^$openai_keyy=" "$env_file" | cut -d'=' -f2)" ]; then
  echo "No Open AI access key found in $env_file.  Please create one at https://platform.openai.com/api-keys and enter it below."
  chromium-browser "https://platform.openai.com/api-keys"  2>/dev/null &

  # Prompt user for the access key
  read -p "Enter Open AI key value: " user_input

  echo "$openai_key=$user_input" >> "$env_file"
  echo "Added key to $env_file"
fi

# check if the open ai organization id variable is present in the file
openai_org="CHAT_GPT_ORG_ID"
if ! grep -q "^$openai_org=" "$env_file" || [ -z "$(grep "^$openai_org=" "$env_file" | cut -d'=' -f2)" ]; then
  echo "No Open AI organization id found in $env_file.  Please create one at https://platform.openai.com/account/organization and enter it below."
  chromium-browser "https://platform.openai.com/account/organization"  2>/dev/null &

  # Prompt user for the access key
  read -p "Enter Open AI organization id value: " user_input

  echo "$openai_org=$user_input" >> "$env_file"
  echo "Added key to $env_file"
fi

# next sign up for google cloud and create a project
echo "Next, sign up for google cloud and create a project at https://cloud.google.com/speech-to-text"
chromium-browser "https://cloud.google.com/speech-to-text"  2>/dev/null &

# Prompt the user to press Enter
read -p "Press Enter once you have signed up for google cloud platform."

# next create a service account and download the json file
echo "Next, create a service account and download the json file at https://console.cloud.google.com/iam-admin/serviceaccounts"
chromium-browser "https://console.cloud.google.com/iam-admin/serviceaccounts"  2>/dev/null &

echo "Looking for the service account json file in ~/Downloads"
downloads_dir="$HOME/Downloads"
gpt_dir="/var/lib/gpt"
json_files=($(find "$downloads_dir" -type f -name "*.json" -exec stat -c "%Y %n" {} + | sort -nr | cut -d ' ' -f2-))
selected_json=""

for json_file in "${json_files[@]}"; do
  if jq -e '.type == "" or .project_id == ""' "$json_file" >/dev/null; then
    selected_json="$json_file"
    break
  fi
done

if [ -n "$selected_json" ]; then
  echo "Found a JSON file with missing keys: $selected_json"
  cp "$selected_json" "$gpt_dir"
  echo "Copied $selected_json to $gpt_dir"
else
  echo "No JSON file with missing keys found in $downloads_dir"
fi