# teleprompter
Tool to show text in the terminal to be read looking at the camera.

## Usage

    $ ./teleprompter -script path/to/script-file.txt
    
For demonstration purpose, the teleprompter runs the `example.txt` by default by running the command without the `script` flag:

    $ ./teleprompter
    
## Script

The script file contains the content to be read out loud during the presentation. Each line of the script contains a time tag and a sentence to be read.

    [[00h00m10s000ms 06s000ms]] Hi! My name is Hildeberto Mendonca. I'm a software engineer, working and living in Toronto, Canada.

The beginning of each line of the script comes with a time tag that describes when and for how long the sentence is shown. The time tag follows the format below:

    [[00h00m10s000ms 06s000ms]]
    
The first part is the point in time when the sentence appears, starting from the beginning of the streaming. It is defined in a single word without spaces. The second part is the duration in which the text stays visible. By default, the sentences stay visible, with new sentences printed in new lines.