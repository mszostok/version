set theApp to "iTerm"
set appHeight to 510
set appWidth to 880

tell application theApp
	activate
	reopen
	repeat with x from 1 to (count windows)
		set xAxis to (95 * x) as integer
		set yAxis to (100 * x) as integer
		set the bounds of the window x to {xAxis, yAxis, appWidth + xAxis, appHeight + yAxis}
	end repeat
end tell
