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


func (p Period) Seconds() (uint64, Period) {
	return uint64(p / Second), p % Second
}


func (p Period) Minutes() (uint64, Period) {
	return uint64(p / Minute), p % Minute
}


// Slice out as many whole hours as possible,
// returning both the hours and the remainder
// as a fresh Period for more slicing.
func (p Period) Hours() (uint32, Period) {
	return uint32(p / Hour), p % Hour
}


func (p Period) Days() (uint32, Period) {
	return uint32(p / Day), p % Day
}


func (p Period) Weeks() (uint32, Period) {
	return uint32(p / Week), p % Week
}


func (p Period) Years() (uint32, Period) {
	return uint32(p / Year), p % Year
}


func (p Period) Decades() (uint16, Period) {
	return uint16(p / Decade), p % Decade
}

