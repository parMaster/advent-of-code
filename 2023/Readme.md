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

## [Day 05 - If You Give A Seed A Fertilizer ](https://github.com/parMaster/advent-of-code/tree/main/2023/05-if-you-give-a-seed-a-fertilizer)
Part One was like 15 minutes. Part Two computing for at least 15 minutes...

Brute force it baby! I'd like to, if I had time someday: 
1. ~~make it bruteforce in parallel, because now it's using only one core~~ faning out every seed range to a separate goroutine got every core 100% loaded, cutting the time by 60%
1. ~~think of better solution (collect all stages combinations and solve a problem "what kind of seed can get through?".~~ idk if there is any..

## [Day 06 - Wait for it](https://github.com/parMaster/advent-of-code/tree/main/2023/06-wait-for-it)
The easiest day so far

## [Day 07 - Camel Cards](https://github.com/parMaster/advent-of-code/tree/main/2023/07-camel-cards)
Is this a Go way? Trying not to use Go as a C++ with garbage collector.

## [Day 08 - Haunted Wasteland](https://github.com/parMaster/advent-of-code/tree/main/2023/08-haunted-wasteland)
LCM and prime generator form scratch. First solution I got with an online LCM calculator and realized it's not brute forceable, so, in a way, if I was ok with cheating, I didn't even need to write the code, but it was a good exercise.

## [Day 09 - Mirage Maintenance](https://github.com/parMaster/advent-of-code/tree/main/2023/09-mirage-maintenance)
Is that it? Ok, I'll take it.

## [Day 10 - Pipe Maze](https://github.com/parMaster/advent-of-code/tree/main/2023/10-pipe-maze)
Another one brutal punch in the groin. It took me too long to do it, even after spoiling part two with a subreddit meme, it took so much time and attention to get it right. The code is too long and ugly, I'm sure it can be solved in a much more elegant way, but I'll take it, it works with no cheating like painting the map in a paint program and counting the pixels, and it vomits a lot of visuals.

## [Day 11 - Cosmic Expansion](https://github.com/parMaster/advent-of-code/tree/main/2023/10-cosmic-expansion)
Took a couple hours to make "expansion" part right, used image.Point to do coordinates math, which gave correct distance value, so all it took is to correctly read and expand the map. Part two was an additional parameter to the read function, that's it. A couple of functions for coordinates math from previous days happened to be useful.