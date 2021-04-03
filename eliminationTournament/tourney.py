def tourney(inp):
    plays = [inp[:]]
    while len(inp) > 1:
        next_round = []
        if len(inp) % 2 != 0:
            next_round.append(inp[-1])
            inp.pop(-1)

        while len(inp) > 1:
            next_round.append(max(inp[0], inp[1]))
            inp.pop(0)
            inp.pop(0)

        plays.append(next_round[:])
        inp = next_round

    return plays
