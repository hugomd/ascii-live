package frames

var Luv = FrameType{
	GetFrame: func(i int) string {
		return "\033[31m" + heartFrames[i%len(heartFrames)] + "\033[0m"
	},
	GetLength: DefaultGetLength(heartFrames),
	GetSleep:  DefaultGetSleep(),
}

var heartFrames = []string{
	// 1. PEAK EXPANSION — Maximum width & brilliant '@' core
	`
    =+**+=     =+**+=
  =+*####*+= =+*####*+=
=+*###################*+=
=+*#####@@@@@@@@@#####*+=
 =+*#################*+=
   =+*#############*+=
      =+*#######*+=
         =+*#*+=
           =+=
            =
`,
	// 2. SUSTAINED GLOW — Holds full size while core gently dims to '#'
	`
    =+**+=     =+**+=
  =+*####*+= =+*####*+=
=+*###################*+=
=+*###################*+=
 =+*#################*+=
   =+*#############*+=
      =+*#######*+=
         =+*#*+=
           =+=
            =
`,
	// 3. GENTLE CONTRACTION — Softening to mid-large (width 21)
	`
    =+**+=     =+**+=
   =+*###*+= =+*###*+=
  =+*###############*+=
  =+*###############*+=
   =+*#############*+=
     =+*#########*+=
       =+*#####*+=
         =+*#*+=
           =+=
            =
`,
	// 4. MEDIUM CONTRACTION — Softening to medium (width 19)
	`
     -=+=-     -=+=-
    -=+**+=- -=+**+=-
   -=+*###########*+=-
   -=+*####***####*+=-
    -=+*#########*+=-
      -=+*#####*+=-
        -=+***+=-
          -=+=-
            -
`,
	// 5. RESTING MINIMUM — Smallest size, dimmest ':' glow
	`
       :-=:   :-=:
      :-=++: :++=-:
      :-=+++++++=-:
      :-=+++++++=-:
        :-=+++=-:
          :-=-:
            :
`,
	// 6. RESTING PAUSE — Holds at minimum for a natural heartbeat beat
	`
       :-=:   :-=:
      :-=++: :++=-:
      :-=+++++++=-:
      :-=+++++++=-:
        :-=+++=-:
          :-=-:
            :
`,
	// 7. EXPANDING — Waking back up to medium
	`
     -=+=-     -=+=-
    -=+**+=- -=+**+=-
   -=+*###########*+=-
   -=+*####***####*+=-
    -=+*#########*+=-
      -=+*#####*+=-
        -=+***+=-
          -=+=-
            -
`,
	// 8. IGNITING — Surging to mid-large before the peak burst
	`
    =+**+=     =+**+=
   =+*###*+= =+*###*+=
  =+*###############*+=
  =+*###############*+=
   =+*#############*+=
     =+*#########*+=
       =+*#####*+=
         =+*#*+=
           =+=
            =
`,
}
