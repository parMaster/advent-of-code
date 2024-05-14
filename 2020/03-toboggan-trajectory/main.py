import time

def solve(file="input.txt"):
	p1, p2 = 0, 1
	
	slopes = [
		(1, 1),
		(3, 1), # Part 1 slope (slopes[1])
		(5, 1),
		(7, 1),
		(1, 2)
	]

	lines = open(file, "r").read().splitlines()
	grid = [[0]*len(lines[0])]*len(lines)
	for i in range(len(lines)):
		grid[i] = list(lines[i])
	
	for i in range(len(slopes)):
		print(slopes[i][0], slopes[i][1])
		x, y = 0, 0
		p2trees = 0
		while y < len(grid):
			
			if x >= len(grid[y]):
				x = x - len(grid[y])

			if grid[y][x] == "#" and i == 1:
				p1+=1
			
			if grid[y][x] == "#":
				p2trees+=1
			
			x+=slopes[i][0]
			y+=slopes[i][1]

		p2*=p2trees
	return p1, p2

print("Day 03: Toboggan Trajectory")
start = time.time()
p1, p2 = solve("input.txt")
print("\tPart 1: ", p1) # 240
print("\tPart 2: ", p2) # 2832009600
print("Done in %s seconds" % round(time.time() - start, 3))
