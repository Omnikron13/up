package main


// Define the (mostly) fixed units of time in nanoseconds
const (
	Nanosecond  =    1
	Microsecond = 1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Centisecond =   10 * Millisecond
	Decisecond  =   10 * Centisecond
	Second      =   10 * Decisecond
	Minute      =   60 * Second
	Hour        =   60 * Minute
	Day         =   24 * Hour
	Week        =    7 * Day
	Year        =  365 * Day
	Decade      =   10 * Year
	Century     =   10 * Decade
	Millennia   =   10 * Century  // Only about half a millennia actually fits...
)


// Stores a length of time in nanoseconds
type Period uint64

