#!/bin/bash

if [ "$(id -u)" -eq 0 ] || [[ $EUID -eq 0 ]]; then
  echo "This script should not be run as root!"
  exit 1
fi

# Install Golang
go_desired_version="1.21.5"
go_installed_version=$(go version | awk '{print $3}' | cut -c 3-)

if [ "$go_installed_version" == "$go_desired_version" ]; then
  echo "Golang is already installed with version $go_desired_version"
else
  echo "Golang version $go_desired_version not found, installing..."

  # Download and install Golang
  wget https://golang.org/dl/go$go_desired_version.linux-arm64.tar.gz
  sudo tar -C /usr/local -xzf go$go_desired_version.linux-arm64.tar.gz
  echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
  echo 'export GOPATH=$HOME/go' >> ~/.bashrc
  echo 'export GOBIN=$GOPATH/bin' >> ~/.bashrc
  echo 'export PATH=$GOBIN:$GOROOT/bin:$PATH' >> ~/.bashrc

  # Clean up downloaded archive
  rm go$go_desired_version.linux-arm64.tar.gz
  echo "Golang version $go_desired_version installed successfully"

  source ~/.bashrc
fi

# install jq
if which jq &> /dev/null; then
  echo "jq is already installed. Skipping installation."
else
  echo "Installing jq"
  sudo apt-get install jq
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
declare -a targets=("gpt-api" "gpt-app" "gpt-ctl")

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
sudo chown -R "$USER:$USER" "/var/lib/gpt/"

# Check if the file exists
if [ ! -f "$env_file" ]; then
  echo "File $env_file does not exist, creating it..."
  touch "$env_file"
fi

# Build /var/lib/gpt/gpio.env file
gpio_env_file="/var/lib/gpt/gpio.env"

# Check if the file exists
if [ ! -f "$gpio_env_file" ]; then
  echo "File $gpio_env_file does not exist, creating it..."
  touch "$gpio_env_file"
fi


# Check if each environment variable exists in the file
check_variable() {
  local env_file=$1
  local variable_name=$2
  local variable_value=$3

  if ! grep -q "^$variable_name=" "$env_file"; then
    echo "$variable_name=$variable_value" >> "$env_file"
    echo "Added $variable_name to $env_file"
  else
    echo "$variable_name already exists in $env_file"
  fi
}

# Check and add each environment variable
check_variable $gpio_env_file "MOTOR_MOUTH_ENA" "29"
check_variable $gpio_env_file "MOTOR_MOUTH_IN1" "31"
check_variable $gpio_env_file "MOTOR_MOUTH_IN2" "33"
check_variable $gpio_env_file "MOTOR_BODY_IN3" "35"
check_variable $gpio_env_file "MOTOR_BODY_IN4" "37"
check_variable $gpio_env_file "MOTOR_BODY_ENB" "32"
check_variable $gpio_env_file "AUDIO_DETECTOR" "36"

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

chat_gpt_sys_prompt_variable="CHAT_GPT_SYS_PROMPT"
default_chat_gpt_sys_prompt='"Hello, I am physical wall mounted animatronic singing fish named Billy Bass that can assist as a sentient all-knowing AI."'
# Check if CHAT_GPT_API_SYS_PROMPT is present in the file
if ! grep -q "^$chat_gpt_sys_prompt_variable=" "$env_file"; then
  echo "The $chat_gpt_sys_prompt_variable variable is not set in $env_file."
  echo "$chat_gpt_sys_prompt_variable=$default_chat_gpt_sys_prompt" >> "$env_file"
  echo "Added $chat_gpt_sys_prompt_variable to $env_file with default value."
fi

# Check if the open ai access key variable is present in the file
openai_key="CHAT_GPT_API_KEY"
if ! grep -q "^$openai_key=" "$env_file" || [ -z "$(grep "^$openai_key=" "$env_file" | cut -d'=' -f2)" ]; then
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
  if jq -e '.type == "service_account"' "$json_file" >/dev/null; then
    selected_json="$json_file"
    break
  fi
done

# after finding the latest json file, copy it to the gpt directory
if [ -n "$selected_json" ]; then
  filename="$gpt_dir/$(basename "$selected_json")"
  echo "Found service account file: $filename"
  cp "$selected_json" "$filename"
  echo "Copied $filename to $gpt_dir"
  gcloud auth activate-service-account --key-file=$filename
  export GOOGLE_APPLICATION_CREDENTIALS=$filename
else
  echo "service account file not found. exiting..."
  exit 1
fi

# add the location of the service account file to the config.env file
google_application_creds="GOOGLE_APPLICATION_CREDS"
if ! grep -q "^$google_application_creds=" "$env_file" || [ -z "$(grep "^$google_application_creds=" "$env_file" | cut -d'=' -f2)" ]; then
  echo "$google_application_creds=$filename" >> "$env_file"
  echo "Added google application creds to $env_file"
fi

# make the wake word dir
wake_word_dir="/var/lib/gpt/wake-words"
if [ ! -d "$wake_word_dir" ]; then
  echo "Creating wake word directory $wake_word_dir"
  mkdir -p "$wake_word_dir"
fi

cp wake-words/* "$wake_word_dir"

# make the wake word dir
asound_file="/etc/asound.conf"
if [ ! -d "$asound_file" ]; then
  cat /proc/asound/cards
  # Prompt user for the access key
  read -p "Enter the number of the sound card to use: " user_input

  sed "s/<sound_card>/$user_input/g" packaging/asound.conf > asound.conf.tmp
  echo "Creating $asound_file"
  sudo cp asound.conf.tmp /etc/asound.conf
fi

# Get the current username
USERNAME=$USER

# Replace <user> with $USERNAME in the service files
sed "s/<user>/$USERNAME/g" packaging/gpt-api.service > gpt-api.service.tmp
sed "s/<user>/$USERNAME/g" packaging/gpt-app.service > gpt-app.service.tmp

# Copy the modified service files to /etc/systemd/system/
sudo cp gpt-api.service.tmp /etc/systemd/system/gpt-api.service
sudo cp gpt-app.service.tmp /etc/systemd/system/gpt-app.service
rm gpt-api.service.tmp
rm gpt-app.service.tmp

# Reload systemd to pick up the changes
sudo systemctl daemon-reload

# Enable and start the services
sudo systemctl enable gpt-api.service
sudo systemctl start gpt-api.service
sudo systemctl enable gpt-app.service
sudo systemctl start gpt-app.service
