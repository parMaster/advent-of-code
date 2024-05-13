import time

def solve(file="input.txt"):
	p1, p2 = 0, 0
	lines = open(file, "r").read().splitlines()
	for i in range(len(lines)):
		parts = lines[i].split(":")
		pwd = parts[1].strip(" ")
		min, max = parts[0].split(" ")[0].split("-", 2)
		# Part 1:
		if  int(min) <= pwd.count(parts[0][-1]) <= int(max):
			p1+=1
		# Part 2:
		if (pwd[int(min)-1] == parts[0][-1]) ^ (pwd[int(max)-1] == parts[0][-1]): # XOR
			p2+=1

	return p1, p2


print("Day 02: Password Philosophy")
start = time.time()
p1, p2 = solve()
print("\tPart 1: ", p1)
print("\tPart 2: ", p2)
print("Done in %s seconds" % round(time.time() - start, 3))


