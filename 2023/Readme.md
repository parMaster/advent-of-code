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

Also available in [JS](https://github.com/parMaster/advent-of-code/tree/main/2023/09-mirage-maintenance/main.js).
I like these .every and .reduce thingies.

## [Day 10 - Pipe Maze](https://github.com/parMaster/advent-of-code/tree/main/2023/10-pipe-maze)
Another one brutal punch in the groin. It took me too long to do it, even after spoiling part two with a subreddit meme, it took so much time and attention to get it right. The code is too long and ugly, I'm sure it can be solved in a much more elegant way, but I'll take it, it works with no cheating like painting the map in a paint program and counting the pixels, and it vomits a lot of visuals.

## [Day 11 - Cosmic Expansion](https://github.com/parMaster/advent-of-code/tree/main/2023/11-cosmic-expansion)
Took a couple hours to make "expansion" part right, used image.Point to do coordinates math, which gave correct distance value, so all it took is to correctly read and expand the map. Part two was an additional parameter to the read function, that's it. A couple of functions for coordinates math from previous days happened to be useful.

## [Day 12 - Hot Springs (p1 only)](https://github.com/parMaster/advent-of-code/tree/main/2023/12-hot-springs)
Part one solved with a recursive brute force, part two left for better days...

## [Day 13 - Point of Incidence](https://github.com/parMaster/advent-of-code/tree/main/2023/13-point-of-incidence)
Pretty straightforward solution, just a couple of well-tested functions, combined together. Nothing fancy, one run = one star.

## [Day 14 - Parabolic Reflector Dish](https://github.com/parMaster/advent-of-code/tree/main/2023/14-parabolic-reflector-dish)
After solving Part One with naive string manipulation, for Part Two it was necessary to bring big guns - DP. quickly refactored slice of string to the map of coordinates, didn't bother refactoring four practically identical copy-pasted pieces of code. It was obvious that there should be a cycle, so it was a matter of serializing the matrix and saving to cache, then detect cycle lenght and skip all the iterations in between. Pretty obvious, but requires attention and time to implement.

## [Day 15 - Lens Library](https://github.com/parMaster/advent-of-code/tree/main/2023/15-lens-library)
It's just a translation of the problem description into code. One run = one star.

Also available in [JS](https://github.com/parMaster/advent-of-code/tree/main/2023/15-lens-library/index.js). It's so unbelievably hard to write JS after Go - no types, always guessing what's the type of the variable. Still know very little even about the basics of (V8)JS, but it's fun to learn. Map() apparently is a linked hash map, so it's easy to iterate over it in the insertion order, which saved one for-loop compared to Go version.

## [Day 16 - The floor will be lava](https://github.com/parMaster/advent-of-code/tree/main/2023/16-the-floor-will-be-lava)
Coordinates math, again. I'm glad I had functions for that already. The rest is just a matter of reading the problem description and translating it into code.

## [Day 17 - Clumsy Crucible](https://github.com/parMaster/advent-of-code/tree/main/2023/17-clumsy-crucible)
First part traditionally - naive method (flood-fill and count everything), to see the second part, which was insane. Implemented Dijkstra pretty fast, but stuck for hours, unable to imagine the 4-D solution space and even think of a composite key for priority queue, to avoid overwriting the same cell with different values. Finally, I gave up and looked at the subreddit, where I found a discussion of a composite array, which was a good hint. Couldn't stop before got the second star.

## [Day 18 - Lavaduct Lagoon](https://github.com/parMaster/advent-of-code/tree/main/2023/18-lavaduct-lagoon)
This is really not fair. I've spent my day inventing the rendering algorithm for the second part, wrote a metric shitton of code and tapped out - got to the reddit, saw two keywords, read two wiki articles and got the solution which was a stupid formula, no iterations, no coordinate math, no debugging, no nothing, just "if you know, you know". Should've I spent some time researching first? At least googling something like "area of the polygon". Though, I'm surprised that the area of the polygon can be calculated with a simple formula, I thought it's a complex problem. [I'll just leave it here for later, maybe I'll need this ray-tracing algorithm someday.](https://www.reddit.com/r/adventofcode/comments/18l25ks/comment/kdv3lqz/?utm_source=share&utm_medium=web2x&context=3)

## [Day 19 - Aplenty](https://github.com/parMaster/advent-of-code/tree/main/2023/19-aplenty)
First part is just a naive solution - put all the parts through the workflows and count the results. Second part is pretty mush the same problem as Day 5 which I solved with brute force but this one is N>10^14... so, I guess I can't get away from solving it a proper way.

Thought of giving up part two after writing 95% working code, never understood my mistake, went to reddit, found [a comment](https://www.reddit.com/r/adventofcode/comments/18mau1e/comment/ke38kvc/?utm_source=reddit&utm_medium=web2x&context=3) - seems like I'm not the one with the same mistake. It only took 10 minutes to fix reading the input data correctly and get the second star. I'm glad I didn't give up.

~~That's it for now. I'm really pissed after getting day 20 wrong TWICE only to realize that problem description didn't match the example in the description. I'm not sure if I'll continue, it's getting difficult to understand 4 pages of incorrect problem description just to implement a glorified counter with cycle detection.~~

## [Day 20 - Pulse Propagation](https://github.com/parMaster/advent-of-code/tree/main/2023/20-pulse-propagation)
First part would be straightforwart if I didn't get the problem description wrong. Second part is identifying modules that we need to detect cycles on, counting cycles and finding the LCM of all the cycles. That "otherwise" word in the problem description screwed me up for days, _otherwise_ I would've solved it in a couple of hours.

## [Day 21 - Step Counter](https://github.com/parMaster/advent-of-code/tree/main/2023/21-step-counter)
Only raw-dog bruteforce, no optimization whatsoever. Takes 30 minutes on 10 year old i5 3.5Ghz, 20 minutes on 10 year old 8-core Xeon. Hardware must work for its money.
It was the plan from the very beginning to solve it with a brute force, then check for cycles. It just took very long to iron-out all the bugs, and it takes really long time to compute, but it's a correct solution.
Also, did Part 2 with raw slices, no image.Point, no fancy stuff, just a primitive coordinates math.

Again and again - it's a really messy code, because there were N hipothesis that I was testing for different parts of the strategy, and this one worked, so why bother refactoring it? The chances are - it won't run anymore, it did its thing, it's done. ~~I'd like to refactor it, because I know which hipothesis is correct, maybe I'll do it someday.~~ Got it down to 5 seconds with a couple of optimizations - maps instead of slices, prediction algorithm dramatically faster.

## [Day 22 - Sand Slabs](https://github.com/parMaster/advent-of-code/tree/main/2023/22-sand-slabs)
Bruteforced the first part, 5 minutes later bruteforced the second part. Took a real TDD approach with this one - failing tests first, then code. First test was to check coordinate math for intersecting rectangles, then reading input data into a stack, then dropping stack. 

