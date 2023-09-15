package main

type Result int64

const (
	WON Result = iota
	LOST
	TIE
)

func (s Result) String() string {
	switch s {
	case WON:
		return "won"
	case LOST:
		return "lost"
	case TIE:
		return "tie"
	}
	return "unknown"
}

type Take int64

const (
	ROCK Take = iota
	PAPER
	SCISSORS
)

func (t Take) String() string {
	switch t {
	case ROCK:
		return "rock"
	case PAPER:
		return "paper"
	case SCISSORS:
		return "scissors"
	}
	return "unknown"
}

var outcomeMap = map[Take]map[Take]Result{
	ROCK: {
		PAPER:    LOST,
		SCISSORS: WON,
	},
	PAPER: {
		ROCK:     WON,
		SCISSORS: LOST,
	},
	SCISSORS: {
		ROCK:  LOST,
		PAPER: WON,
	},
}

func (t Take) Match(o Take) Result {
	if t == o {
		return TIE
	} else {
		return outcomeMap[t][o]
	}
}

type GameRequest struct {
	UserTake Take `json:"take"`
}

type GameResponse struct {
	RobotTake   Take   `json:"robotTake"`
	RobotResult Result `json:"robotResult"`
	UserTake    Take   `json:"userTake"`
	UserResult  Result `json:"userResult"`
}
