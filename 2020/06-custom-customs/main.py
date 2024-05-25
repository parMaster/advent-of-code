import os, time

def solve(file="input.txt"):
	p1, p2 = 0,0

	groups = open(file, "r").read().split("\n\n")
	groups = [group.strip() for group in groups]

	for group in groups:
		answers = dict()
		for line in group.split("\n"):
			for a in line:
				answers[a] = answers.get(a, 0) + 1
		p1 += len(answers)

		# Part 2
		for answer in answers:
			if answers[answer] == len(group.split("\n")): # everyone in the group answered yes
				p2 += 1

	return p1, p2


os.chdir(__file__.rsplit("/", 1)[0])
print("Day 06: Custom Customs")
start = time.time()
p1, p2 = solve("input0.txt")
assert p1 == 11
assert p2 == 6

# os.abort()
p1, p2 = solve()
print("\tPart 1: ", p1)
print("\tPart 2: ", p2)
assert p1 == 6532
assert p2 == 3427
print("Done in %s seconds" % round(time.time() - start, 3))
