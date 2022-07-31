(*

From http://www.labnol.org/software/resize-mac-windows-to-specific-size/28345/
This Apple script will resize any program window to an exact size and the
window is then moved to the center of your screen.  Specify the program name,
height and width below and run the script.

Written by Amit Agarwal on December 10, 2013

*)

set theApp to "iTerm"
set appHeight to 510
set appWidth to 1095

tell application theApp
	activate
	reopen
	repeat with x from 1 to (count windows)
		set xAxis to (95 * x) as integer
		set yAxis to (100 * x) as integer
		set the bounds of the window x to {xAxis, yAxis, appWidth + xAxis, appHeight + yAxis}
	end repeat
end tell

--do shell script "ls /Applications/"
--do shell script "export KUBECONFIG=''"
--do shell script "clear"
--do shell script "gimme version"
--do shell script "screencapture -x -R0,25,1285,650 file.png"
