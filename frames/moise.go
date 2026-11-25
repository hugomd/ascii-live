package frames

import "time"

var Moise = FrameType{
	GetFrame:  DefaultGetFrame(moiseFrames),
	GetLength: DefaultGetLength(moiseFrames),
	GetSleep:  func() time.Duration { return 900 * time.Millisecond },
}

var moiseFrames = []string{
	`   /\    /\   
  /  \  /  \  
 /    \/    \ `,
	`  ___________ 

 |           |
 |___________|`,
	`              

      | |     
      |_|     `,
	`    _____     
   / ____|    
   \___ \     
   ____) |    
  |_____/     `,
	`    ______    

   |  ____|   
   | |__      
   |  __|     
   | |____    
   |______|   `,
}
