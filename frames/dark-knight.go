package frames

// DarkKnight is a cinematic Batman-inspired ASCII animation.
var DarkKnight = DefaultFrameType(darkknightPR135Frames)

var darkknightPR135Frames = []string{

	// FRAME 1: Gotham City - Bat-Signal begins to appear
	`
                         *
                        / \
                       /   \
                      /     \
                     /       \
                    /         \
                   /           \
                  /             \
                 /               \
                /                 \
               /                   \
              /                     \
             /                       \
            /                         \
           /                           \
          /_____________________________\
                    
                 GOTHAM CITY
	`,

	// FRAME 2: The Bat-Signal emerges
	`
                   .       .
                .             .
             .                   .
          .                         .
       .                               .
     .          /\          /\          .
    .          /  \________/  \           .
   .          /                \          .
  .          /    /\      /\     \         .
 .          /____/  \____/  \_____\         .
  .          \                  /          .
   .          \________________/           .
    .                                  .
       .                            .
          .                      .
             .                .

                 THE DARK KNIGHT
	`,

	// FRAME 3: Batman emblem becomes visible
	`
                  .------------.
              .-'                '-.
           .-'       /\      /\      '-.
         .'         /  \____/  \         '.
        /          /            \          \
       /      ____/              \____       \
      |      /                         \      |
      |     /     /\              /\     \     |
      |    /_____/  \____________/  \_____\    |
      |          \                  /          |
      |           \________________/           |
       \                                       /
        '.                                   .'
          '-.                             .-'
             '---------------------------'

                   BATMAN
	`,

	// FRAME 4: The Dark Knight - Full emblem
	`
             __________________________________
          .-'                                  '-.
        .'       /\                    /\         '.
       /        /  \__________________/  \          \
      /        /                        \           \
     |    ____/                          \____       |
     |   /                                    \      |
     |  /       /\                    /\        \     |
     | /_______/  \__________________/  \________\    |
     |             \                  /               |
     |              \________________/                |
      \                                               /
       '.                                           .'
         '-.                                     .-'
            '-----------------------------------'

                  THE DARK KNIGHT
	`,

	// FRAME 5: Spotlight pulse
	`
                    *       *
                 *               *
              *                     *
           *                           *
        *          /\          /\         *
      *           /  \________/  \          *
     *           /                \          *
    *       ____/                  \____      *
   *       /                            \      *
  *       /       /\            /\        \     *
 *       /_______/  \__________/  \________\     *
 *                \              /               *
  *                \____________/               *
    *                                        *
       *                                  *
          *                            *
              *                    *

                GOTHAM NEEDS YOU

                  BATMAN
	`,

	// FRAME 6: Final cinematic Bat-Signal
	`
                .----------------.
             .-'                  '-.
           .'       /\      /\        '.
          /        /  \____/  \         \
         /        /            \         \
        |    ____/              \____     |
        |   /                        \    |
        |  /      /\          /\      \   |
        | /______/  \________/  \______\  |
        |          \          /           |
         \          \________/           /
          '.                          .'
            '-.                    .-'
               '----------------'

              THE DARK KNIGHT

             I AM THE NIGHT.
	`,

	// FRAME 7: Fade back to the Bat-Signal
	`
                         *
                       *   *
                     *       *
                   *           *
                 *               *
               *                   *
             *                       *
           *                           *
         *                               *
       *                                   *
     *                                       *
   *                                           *
 *                                               *

                  BATMAN
	`,
}
