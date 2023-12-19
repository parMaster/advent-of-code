const { assert } = require('node:console');
const fs = require('node:fs');

let hash = (s) => {
	let h = 0;
	for (let i=0; i<s.length; i++) {
			h += s.charCodeAt(i)
		h *= 17
		h %= 256
	}
	return h;
}

assert(hash('HASH') == 52)

// part 1
let input = fs.readFileSync('../aoc-inputs/2023/15/input.txt', 'utf8');
input = input.trimEnd();
const data = input.split(',');
const p1 = data.reduce(
	(acc, curr) => acc + hash(curr), 0);


// part 2
let p = new Array(256)
for (let i=0; i<data.length; i++) {
	if (data[i].indexOf('-') > -1) {
		let label = data[i].slice(0,data[i].indexOf('-'))
		let box = hash(label)
		if (p[box]){
			p[box].delete(label)
		}
	} else {
		let parts = data[i].split('=')
		let label = parts[0]
		let focal = parts[1]
		let box = hash(label)
		if (!p[box]) {
			p[box] = new Map()
		}
		p[box].set(label, focal)
	}
}

let sum = 0
for (let i=0; i<p.length; i++) {
	if (p[i]) {
		let slot = 0
		for (let [_, focal] of p[i]) {
			slot++
			sum += (i+1)*slot*focal
		}
	}
}

console.log("Day 15: Lens Library")
console.log("\tPart One:",p1) // 513643
assert(p1 == 513643)
console.log("\tPart Two:",sum) // 265345
assert(sum == 265345)