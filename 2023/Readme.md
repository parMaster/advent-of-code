# [Advent of Code 2023](https://adventofcode.com/2023)
No Copilot or any AI is allowed here.

## [Day 01 - Trebuchet?!](https://github.com/parMaster/advent-of-code/tree/main/2023/01-calibration-values)
Part 1 is solved without regex. Part 2 has a nasty gotcha with matching words like "oneight" or "twone".

## [Day 02 - Cube Conundrum](https://github.com/parMaster/advent-of-code/tree/main/2023/02-cube-conundrum)
Typical dynamic programming problems. Parsing input data was the longest part.

## [Day 03 - Gear Ratios](https://github.com/parMaster/advent-of-code/tree/main/2023/03-gear-ratios)
This one was brutal :D I'm glad I tested functions in the first part. I wonder what will be the stats for this day ))

## [Day 04 - Scratchcards](https://github.com/parMaster/advent-of-code/tree/main/2023/04-scratchcards)
10KB of problem description, 33 lines of solution. One hour wasted on a classic `m := range memo` (`m` is the key, not the value) in the unnecessary code that I deleted afterwards anyway.

## [Day 05 - If You Give A Seed A Fertilizer ](https://github.com/parMaster/advent-of-code/tree/main/2023/04-scratchcards)
Part One was like 15 minutes. Part Two computing for at least 15 minutes...

Brute force it baby! I'd like to, if I had time someday: 
1. ~~make it bruteforce in parallel, because now it's using only one core~~ faning out every seed range to a separate goroutine got every core 100% loaded, cutting the time by 60%
1. ~~think of better solution (collect all stages combinations and solve a problem "what kind of seed can get through?".~~ idk if there is any..
