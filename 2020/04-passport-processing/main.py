import time
import os
import re

def solve(file="input.txt"):
	p1, p2 = 0, 0
	passports = open(file, "r").read().strip().split("\n\n")
	passports = [p.replace("\n", " ") for p in passports]

	for p in passports:
		fields = p.split(" ")

		props = dict()
		for f in fields:
			props[f.split(":")[0]] = f.split(":")[1]

		if all(e in props.keys() for e in ("byr", "iyr", "eyr", "hgt", "hcl", "pid", "ecl")) and \
			1920 <= int(props["byr"]) <= 2002 and \
			2010 <= int(props["iyr"]) <= 2020 and \
			2020 <= int(props["eyr"]) <= 2030 and \
			((props["hgt"].endswith("cm") and 150 <= int(props["hgt"].rstrip("cm")) <= 193) or \
			(props["hgt"].endswith("in") and 59 <= int(props["hgt"].rstrip("in")) <= 76)) and \
			re.search("#[0-9a-f]{6}", props["hcl"]) != "None" and len(props["hcl"]) == 7 and \
			re.search("[0-9]{9}", props["pid"]) != "None" and len(props["pid"]) == 9 and \
			props["ecl"] in ("amb", "blu", "brn", "gry", "grn", "hzl", "oth"):
				p2 +=1

		if len(fields) == 8 or (len(fields) == 7 and "cid" not in props.keys()):
			p1+=1

	return p1, p2

os.chdir(__file__.rsplit("/", 1)[0]) # hacky way to change directory to input file location
print("Day 04: Passport Processing")
start = time.time()
p1, p2 = solve("input.txt")
print("\tPart 1: ", p1)
print("\tPart 2: ", p2)
assert p1 == 228
assert p2 == 175
print("Done in %s seconds" % round(time.time() - start, 3))
