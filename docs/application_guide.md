# Application Guide

This guide provides an overview of the features available in the Barcode Reader Emulator application.

## Main Application Window

[Screenshot of the main application window]
*This screenshot should depict the main window of the application, showing the "Value" input field, "Scan from the screen" button, "Hotkey" display, "Input Key Delay" slider, "Send ENTER at the end" checkbox, and the "Scan Delay" slider.*

## Features

### 1. Barcode Value Input

You can manually type or paste barcode data directly into the "Value" text field. This is useful for testing or when a barcode scanner is not available.

### 2. Screen Barcode Scanning

The application allows you to scan barcodes directly from your screen.
- Click the "Scan from the screen" button.
- The application will attempt to detect and decode a barcode visible on your screen.
- If successful, the decoded barcode value will populate the "Value" text field.
- If scanning fails, a message box will appear with an error. This might happen if the barcode is too small, blurry, or not clearly visible. Try zooming in on the barcode image or ensuring it's unobstructed.

### 3. Configurable Hotkey

The application uses a global hotkey to trigger the emulation of typing the barcode value. You can customize this hotkey.
- The currently configured hotkey is displayed next to the "Hotkey" label.
- To change the hotkey, click the "Change" button.
- This will open the hotkey selection dialog.

[Screenshot of the hotkey selection dialog]
*This screenshot should show the dialog box that appears when the "Change" button for the hotkey is clicked. It should display the list of modifier keys (e.g., Ctrl, Shift, Alt) and the dropdown list of keys (e.g., F1, A, B, 1, 2).*

- In the dialog:
    - Select one or more modifier keys (e.g., Ctrl, Shift) from the list on the left.
    - Select a primary key (e.g., F10, V) from the dropdown menu on the right.
    - Click "OK" to save the new hotkey.
    - If the selected hotkey is already in use by another application, an error message may appear. Try a different combination.

### 4. Adjustable Typing Speed

The "Input Key Delay" slider controls the speed at which the application emulates typing.
- The slider ranges from a minimum (faster typing) to a maximum (slower typing) delay between keystrokes.
- Adjust this slider if the target application requires a specific typing speed or if characters are being missed.

[Screenshot of the Input Key Delay slider]
*This screenshot should highlight the "Input Key Delay" slider in the main application window.*

### 5. Optional "Enter" Key Simulation

The 'Send ENTER at the end' checkbox determines whether the application simulates pressing the "Enter" key after typing the barcode value.
- **Checked**: An "Enter" key press will be simulated after the full barcode value (and each segment, if using semicolons) is typed.
- **Unchecked**: No "Enter" key press will be simulated.
This is useful for applications that automatically move to the next field after input or require a manual "Enter" press.

[Screenshot of the 'Send ENTER at the end' checkbox]
*This screenshot should highlight the 'Send ENTER at the end' checkbox in the main application window.*

### 6. Segmented Typing with Delays

You can input multiple barcode segments at once by separating them with a semicolon (`;`) in the "Value" field. The application will type each segment individually.
- **Example**: `VALUE1;VALUE2;VALUE3`
- The application will type `VALUE1`, then pause, then type `VALUE2`, pause, and then type `VALUE3`.
- The "Scan Delay" slider controls the duration of the pause *between* these segments. **Note:** This slider might be more accurately named 'Segment Delay'.
- Adjust this slider to control the timing between the typing of each segment.

[Screenshot of the 'Scan Delay' slider, perhaps with a note about its actual function as 'Segment Delay']
*This screenshot should highlight the "Scan Delay" slider, and ideally, include a textual callout in the image or caption clarifying it controls the delay between segments when using semicolons.*

## Suggested Improvements

1.  **Clarify Slider Label:** The "Scan Delay" slider should be renamed to "Segment Delay" or "Delay Between Segments" as it controls the delay between typing segments, not scanning.
2.  **Improve Error Handling for Scanning:** Provide more specific error messages when barcode scanning fails. For example, instead of a generic "Could not find or parse a barcode," messages like "No barcode detected in the selected area," "Barcode detected but could not be decoded (unsupported type or poor quality)," or "Screen capture failed" would be more helpful.
3.  **Add Support for More Barcode Types:** Currently, the application primarily supports QR codes and some linear barcodes. Consider adding robust support for other common barcode symbologies like Data Matrix, PDF417, Aztec, Code 39, Code 128, EAN, and UPC.
4.  **Visual Feedback for Active Hotkey:** Implement a visual indicator to show that the hotkey listener is active and ready to receive input. This could be a change in the application's tray icon, or a small status indicator in the main window. This would help users confirm the application is ready without having to test the hotkey.
5.  **Configuration Persistence:** Save user settings (selected hotkey, input key delay, "Send ENTER at the end" preference, and segment delay) so they persist across application restarts. This would prevent users from having to reconfigure the application every time they open it. Settings could be stored in a configuration file (e.g., JSON, INI) in the user's application data directory.
6.  **Allow Configuration of Beep Sound:** Add an option in the settings to disable or customize the scanner beep sound. Some users may find the beep distracting or unnecessary.
7.  **Progress Indication for Multi-Segment Typing:** For inputs with many segments separated by semicolons, consider showing a visual progress indicator (e.g., a small progress bar or a text update like "Typing segment 2 of 5..."). This would be helpful for long sequences to let the user know the application is still working.
