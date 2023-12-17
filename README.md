# chat-gpt-rasperry-pi-assistant
ChatGPT Raspberry Pi Assistant is an open-source project that enables voice-based artificial intelligence (AI) interactions on Raspberry Pi devices. It leverages OpenAI's ChatGPT model to provide natural language understanding and generation capabilities, making it a versatile voice assistant.

![billy_bass(1)](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/1f2cb0f3-a0f6-4364-8005-36807af46830)

# Bill of Materials
Parts
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

Tools needed
1. Dremel or similar small cutting tool
2. Drill
3. Hot glue gun
4. Wire strippers
5. Small Label Printer (optional)

# Schematics
![schematics](https://github.com/GoBig87/chat-gpt-rasperry-pi-assistant/assets/39137894/49c2fc3b-7f3d-4afe-a5f9-1bd1bb857400)

# Preparation
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
Mount the H bridge (Part 4) on the right side and try to make it flush with the bottom of the standing mount bracket.  Drill out the holes and mount the H bridge with 1/8 inch standoffs. 

### Step 10
Mount the voltage regulator

### Step 11
With a hot glue gun, glue the Qualcomm Quick Charge power supply (Part 2) to the left side of the board

### Step 12
With a hot glue gun, glue the USB Speaker Phone (Part 2) above the H Birdge and to the right of the raspberry pi.
