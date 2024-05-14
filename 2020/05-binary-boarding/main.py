import time, os, math

# binary partitioning of interval
def decode(interval, line, chr):
	for c in line:
		if c == chr:
			interval[1] -= math.floor((interval[1] - interval[0])/2)
			interval[1] -=1
		else:
			interval[0] += math.floor((interval[1] - interval[0])/2)
			interval[0] +=1
	return interval[0]

def solve(file="input.txt"):
	p1,p2 = 0,0
	lines = open(file, "r").read().strip().split("\n")
	lines = [line.strip() for line in lines]

	seats = set()
	for line in lines:
		row = decode([0,127], line[:7], 'F')
		col = decode([0,7], line[7:], 'L')
		seats.add(row*8 + col)

	# Part 2
	for seat in range(min(seats), max(seats)):
		if seat not in seats:
			return max(seats),seat

os.chdir(__file__.rsplit("/", 1)[0])
print("Day 05: Binary Boarding")
start = time.time()
p1, p2 = solve("input.txt")
print("\tPart 1: ", p1)
print("\tPart 2: ", p2)
# assert p1 == 978
# assert p2 == 727
print("Done in %s seconds" % round(time.time() - start, 3))
