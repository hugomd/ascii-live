package frames

import "time"

var DancingCoin = FrameType{GetFrame: DefaultGetFrame(dancingCoinPR135Frames), GetLength: DefaultGetLength(dancingCoinPR135Frames), GetSleep: func() time.Duration { return 100 * time.Millisecond }}

var dancingCoinPR135Frames = []string{
	"\u001b[36m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n ffffffffttft\n ffffffffttff\n fffffffttttt\n\u001b[0m",
	"\u001b[35m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n ffffffffttft\n  ffffffffttf\n   fffffffttt\n\u001b[0m",
	"\u001b[36m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n  ffffffffttf\n   fffffffftt\n    ffffffftt\n\u001b[0m",
	"\u001b[35m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n   fffffffftt\n    fffffffft\n     fffffftt\n\u001b[0m",
	"\u001b[36m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n    fffffffft\n     ffffffff\n      fffffft\n\u001b[0m",
	"\u001b[35m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n     ffffffft\n      fffffff\n       ffffff\n\u001b[0m",
	"\u001b[36m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n    fffffffft\n     ffffffff\n      fffffft\n\u001b[0m",
	"\u001b[35m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n   fffffffftt\n    fffffffft\n     fffffftt\n\u001b[0m",
	"\u001b[36m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n  ffffffffttf\n   fffffffftt\n    ffffffftt\n\u001b[0m",
	"\u001b[35m      .---.\n     / | \\ \\\n    /  |  \\ \\\n   |   |   | |\n    \\  |  / /\n     '---'\n    LLfftfLL\n   fttttffff\n  ttttttttttt\n ffffffffttft\n  ffffffffttf\n   fffffffttt\n\u001b[0m",
}
