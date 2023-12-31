# chat-gpt-rasperry-pi-assistant
ChatGPT Raspberry Pi Assistant is an open-source project that enables voice-based artificial intelligence (AI) interactions on Raspberry Pi devices. It leverages OpenAI's ChatGPT model to provide natural language understanding and generation capabilities, making it a versatile voice assistant.

https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/c6bdf42a-68d5-4903-87a5-be4ee5d2cc9c

# Bill of Materials
### Parts
1. Raspberry Pi 4/5
2. Qualcomm Quick Charge Power supply with two outputs capable of 5v/3 Amps output (https://www.amazon.com/dp/B074Q3TN8L?psc=1&ref=ppx_yo2ov_dt_b_product_details)
3. USB speaker phone (https://www.amazon.com/dp/B08THGFBTV?psc=1&ref=ppx_yo2ov_dt_b_product_details)
4. H-Bridge Motor Contol (https://www.amazon.com/dp/B09B36LCXQ?psc=1&ref=ppx_yo2ov_dt_b_product_details)
5. Sound Sensor (https://www.amazon.com/dp/B00XT0PH10?psc=1&ref=ppx_yo2ov_dt_b_product_details)
6. 6" USB A to USB Cable (https://www.amazon.com/gp/product/B08LL1SVZD/ref=ppx_yo_dt_b_search_asin_title?ie=UTF8&psc=1)
7. USB A to DC Female Barrel (https://www.amazon.com/gp/product/B08MX8NR2Z/ref=ppx_yo_dt_b_search_asin_title?ie=UTF8&psc=1)
8. Voltage Regulator (https://www.amazon.com/dp/B082XQC2DS?psc=1&ref=ppx_yo2ov_dt_b_product_details)
9. Micro HDMI adapter (https://www.amazon.com/dp/B0BLZM5CQL?psc=1&ref=ppx_yo2ov_dt_b_product_details)
10. Barrel adpaters with wire leads (https://www.amazon.com/dp/B09XQZLM1Q?psc=1&ref=ppx_yo2ov_dt_b_product_details)
11. Barrel adpaters (https://www.amazon.com/dp/B09XQZ5L6G?ref=ppx_yo2ov_dt_b_product_details&th=1)
12. Bread Board leads (https://www.amazon.com/EDGELEC-Breadboard-Optional-Assorted-Multicolored/dp/B07GD2BWPY/ref=sr_1_3?crid=1JYT2UF3EZS2E&th=1)
13. Standoffs and mounting hardware (https://www.amazon.com/dp/B07PHBTTGV?psc=1&ref=ppx_yo2ov_dt_b_product_details)  
15. Bluetooth Keyboard (https://www.amazon.com/Bluetooth-Wireless-Keyboard-Multi-Device-Desktop-Black/dp/B0C4JSGJQ5/ref=sr_1_5?crid=2EJBEYSFHJUFY)
16. A Billy Bass signing fish (https://www.amazon.com/dp/B09NT16S23?psc=1&ref=ppx_yo2ov_dt_b_product_details)

### Tools needed
1. Dremel or similar small cutting tool
2. Drill
3. Hot glue gun
4. Wire strippers
5. Small Label Printer (optional)

![billy_bass(1)](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/1f2cb0f3-a0f6-4364-8005-36807af46830)

# Schematics
![schematics](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/49c2fc3b-7f3d-4afe-a5f9-1bd1bb857400)

# Hardware Install
![IMG_4343](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/9fa4195b-b738-4483-b64d-e30c6ea61e86)

### Step 1
Remove the screws on the back side of the Billy Bass fish and remove the AA batteries.  Open the fish cautiously and be mindful of the attached cables.

### Step 2
Disconnect the cables attached to the circuit board.  Remove the screws that hold the circuit board in place and discard it.  After removing the ciruit board, use a small dremel or similar cutting wheel to remove the plastic standoffs to create more space inside of the fish.

### Step 3
Remove the motion sensor and activation button by unscrewing the screws that hold them into place.  

### Step 4 
Allong the the top side there are plastic flares that will need to be removed in order to fit the raspberry pi.  With a cutting tool, carefully trim the flares down making sure to not cut or damage the edges of the case.

### Step 5 
With wire strippers, trim the motor wires off of the electrical wire adapters.  There will be a pair of wires for each motor.  The red and gray wire pair controls the mouth.  Strip the wires and connect to a DC female plug (Part 11) via the screw clamp bindings.  The red wire should be attached to the negative terminal and gray to the positive terminal.  The black and orange pair control the body and tail.  Connect the black wire to the positive termianl of a DC female plug and the orange to the negative terminal.

### Step 6
![IMG_4344](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/419d5224-40ca-42ae-9f9d-48b67bbe7ec9)
MAKE SURE TO REMOVE THE ALL AA BATTERIES BEFORE THIS STEP!!! FAILURE TO DO SO WILL RESULT IN A NOT SO GOOD TIME!!!
With a cutting tool, cut out the backside/inside of the battery compartment from the inside.  You will want to leave the walls alone.  Be warned that you will be hitting some small metal parts of the battery connection system.  Once you have the backside/inside taken off, slowly trim out the upper wall to allow for cables to be routed through.

### Step 7
![pi_placement](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/0d30dee0-7e69-4a2d-bd24-ff6ae3e15ebd)
Mount the raspberry pi by drilling small holes on the backside of the Fish plaque.  Try to have about a 1/4 inch of spacing from the top and try to make it flush with the raised standing mount container.  Mount the pi with 1/8 inch standoffs.

### Step 8
![audio_sensor](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/ac1bf77d-090f-46d8-b181-e8d594535c04)
![adjustment](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/9ee03e3f-7876-406a-a150-9230d45ce8ff)
Mount the audio sensor (Part 5) below the raspberry pi and inside the corners of the standing mount container.  Drill a small to medium size hole at the location of the audio sensors adjusment screw.  The screw will set the sensitivty of the mouth movement.  The hole will allow for adjustments once completed.

### Step 9
![motor_shield](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/722d29af-87e0-407b-a420-2a24e3d159a4)
Mount the H bridge (Part 4) on the right side and try to make it flush with the bottom of the standing mount bracket.  Drill out the holes and mount the H bridge with 1/8 inch standoffs. Connect DC barrel adapters (Part 10) to the output termianls.  Cable 1: Red OUT1, Black OUT2 and this will be for the mouth motor.  Cable 2: Red OUT4, Black OUT3 and this will be for the tail and head motor.  Remove the enable jumpers and connect the control pins to the rapsberry pi as shown in the schematics diagram.  Connect the VCC and Gnd for the voltage regulator with plain wires. Also attach a gnd pin from the raspberry pi to the gnd of the H bridge.  Connect the logical pins on the H bridge to the corresponding pins on the raspberry pi in the diagram.

### Step 10
Mount the voltage regulator (Part 8, note: not shown in the pictures).  There should be room to mount the voltage regulator right next to the H bridge.  Wire a pair of cables from the H bridge 5v and Gnd inputs and connect them the to voltage regulator output.  Wire a male DC power plug to the input of the voltage regulator.  Make sure black goes to all the ground/negative and red to all the positive terminals.

### Step 11
With a hot glue gun, glue the Qualcomm Quick Charge power supply (Part 2) to the left side of the board

### Step 12
With a hot glue gun, glue the USB Speaker Phone (Part 2) above the H Birdge and to the right of the raspberry pi.

### Step 13
Wire up the parts according to the diagram if you have not already done so.

# Software Install

### Step 1
Install a Raspberry Pi OS following this guide https://www.raspberrypi.com/software/

### Step 2
In a terminal on the Raspberry Pi, clone this repository with
```
git clone https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant.git
``` 
or 
```
git clone git@github.com:GoBig87/chat-gpt-rasperry-pi-assistant.git
```

### Step 3
Start the installation process with
```
cd chat-gpt-rasperry-pi-assistant 
./installer.sh
```
Throughout the installation process you will be prompted to enter certain API keys.
You will need three total API keys

1. Porcupine API Key
2. A Google Cloud Service Account
3. Open AI API Key

The Porcupine API key is used for the wake word detection.  The Google service account is 
used for translating speech to text and text to speech.  Finally, the Open API key is used 
for creating Chat GPT prompting.  Note these are paid services.  You will be billed for heavy
usage.  Light usage should result in very minimal billing (it can range from free to a few 
dollars a month ).  Be warned that you should place limits on these accounts and also fully
vet this repo for security concerns.  You should always be cautious when giving unknown 
applications access to your API keys.

# Configuration  
The installer will create a few files in the `/var/lib/gpt` directory that will be consumed 
by the Chat GPT application.  

## The config.env file
This file contains the values that are used with accessing the various APIs that
are needed by the application

### Values
`PORCUPINE_ACCESS_KEY` This gives the application access to the Porcupine Wake Word servers.

`CHAT_GPT_API_ENDPOINT` End point where the ChatGPT endpoints should be directed to.

`CHAT_GPT_SYS_PROMPT` This is used to change how you would wish Chat GPT to respond.  
This prompt gives the AI an initial state so it knows how it should handle prompts

`CHAT_GPT_API_KEY` Open AI API key to access servers.  Obtained from Opens AIs website.

`CHAT_GPT_ORG_ID` Open AI orginazatioon ID.  Obtained from Opens AIs website.

`GOOGLE_APPLICATION_CREDS` Absolute path to the Google application credentials json file 
for the speech to text and text to speech service account.

## The gpio.env file
This file contains the location for the Raspberry Pi's GPIO pins.  These numbers need to 
be the physical pin location on the Raspberry Pi. These ARE NOT the GPIO pin numbers.  By
default they are autogenerated to match the diagram.  If for some reason you need to adjust
GPIO pins for some reason they can be adjusted here.

### Values 

`MOTOR_MOUTH_ENA` Turns the H bridge motor A on or off

`MOTOR_MOUTH_IN1` Sets the voltage for motor A's positive input

`MOTOR_MOUTH_IN2` Sets the voltage for motor A's negative input

`MOTOR_BODY_IN3`  Sets the voltage for motor B's positive input

`MOTOR_BODY_IN4` Sets the voltage for motor B's negative input

`MOTOR_BODY_ENB` Turns the H bridge motor B on or off

`AUDIO_DETECTOR` Pin for reading input from the audio sensor

## The wake-word directory
This is where you can place custom wake words that will be picked up
by the application.

## Applying Changes
After changing any of these configs users should restart the app with

`sudo systemctl restart gpt-app.service && systemctl restart gpt-api.service`

# Testing & Troubleshooting
Lots can go wrong when attempting trying to get the application to run and  
communicate with all the physical components.  The installation process also installs
some command line tools that are useful for testing each individual pieces.

### gpt-ctl
`gpt-ctl` can be invoked from a terminal and available commands can be viewed
with 

`gpt-ctl --help`

### Speech Testing
To test Googles Speech APIs there are two test commands that you can
use to make sure everything is setup properly

`gpt-ctl s2t` Will take audio input and convert it to a text response.

`gpt-ctl t2s <your text input here>` takes a text input and converts it to speech.

### Motor Testing
To test the motors there a few commands to run

`gpt-ctl raise-tail` This command will raise the tail

`gpt-ctl lower-tail` This command will lower the tail

`gpt-ctl raise-head` This command will raise the head

`gpt-ctl lower-head` This command will lower the head

`gpt-ctl open-mouth` This command will open the mouth

`gpt-ctl close-mouth` This command will close the mouth

`gpt-ctl reset-all` This command will disable a shut off all motors

`gpt-ctl speech-to-movement <duration (0 for continuous)>`  This will enable the audio
sensor and it will open and close the mouth based on if audio is detected over a period.
 This command can be paired with `gpt-ctl t2s <long phrase>` and is very usefull when adjusting
the audio sensors sensitivity. 

### Chat GPT testing
To test to make sure Chat GPT is functioning correctly you can use
the following commands

`gpt-ctl prompt <prompt for chat gpt>`

This will take aa prompt input and output a response from chat GPT.

### Wake Word testing
When testing to make sure the wake word functionality is correctly working you 
can use

`gpt-ctl wake`

This command will run continuously until a wake word is detected.

# Custom Integration
The Chat GPT Raspberry Pi Assistant is powered by a local GRPC server that can be
called from any language that supports GRPC.  All the functionality of the app is 
called through this GRPC API which allows users to make custom apps from the api.  
For instance if a user wants to build a custom python app the proto files can be used to 
generate python proto messages that can be used to make GRPC calls to the Golang GRPC 
server.  This could be useful if you wanted to handle special prompts differently such as 
"What time is it" or "What's the weather".  Current tasks that ChatGPT can't handle but 
could easily be coded up in any language.  Or possibly wifi controlled lights could be
controlled by saying special commands like "turn on all the lights".  This repo is meant 
to be a jumping off point for getting the base assistant off the ground.

TODO add common program language GRPC messages and clients inside a src dir. 
