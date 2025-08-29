package frames

// Frame data (each string = one frame)
var as = []string{
	`
 /\_/\
( o.o )
 > ^ < 
`,
	`
 /\_/\
( -.- )
 > ^ < 
`,
	`
 /\_/\
( o.o )
 >   < 
`,
	`
 /\_/\
( ^.^ )
 > ^ < 
`,
}

// Exported variable with default frame type
var AStrend = DefaultFrameType(as)
