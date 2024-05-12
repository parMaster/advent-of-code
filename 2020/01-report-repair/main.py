import time

def solve():
	f = open("input.txt", "r")
	lines = f.readlines()
	lines = [int(line.strip()) for line in lines]

	for i in range(len(lines)):
		for j in range(i+1, len(lines)):
			# Part 1
			if lines[i] + lines[j] == 2020:
				p1 = (lines[i] * lines[j])
			# Part 2
			for k in range(j+1, len(lines)):
				if lines[i] + lines[j] + lines[k] == 2020:
					p2 = (lines[i] * lines[j] * lines[k])
	f.close()
	return p1, p2

print("Day 01: Report Repair")
start = time.time()
p1, p2 = solve()
print("\tPart 1: ", p1)
print("\tPart 2: ", p2)
print("Done in %s seconds" % round(time.time() - start, 3))
