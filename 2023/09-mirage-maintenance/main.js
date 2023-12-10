const fs = require('node:fs');

const data = fs.readFileSync('../aoc-inputs/2023/09/input.txt', 'utf8');

let predict = function(a) {
	if (a.every(function(v){return v==0})) {
		return 0
	}

	next = new Array(a.length-1)
	for (let i=0; i<a.length-1; i++) {
		next[i] = a[i+1]-a[i]
	}

	return a[a.length-1] + predict(next)
}

console.log("Day 9: Mirage Maintenance")

const numbers = data.split('\n');

let next = 0
next = numbers.reduce(
	(acc, curr) => acc + predict(curr.split(' ').map(Number)), next
)
console.log("\tPart One:",next) // 1921197370

let prev = 0
prev = numbers.reduce(
	(acc, curr) => acc + predict(curr.split(' ').map(Number).reverse()), prev
)
console.log("\tPart Two:",prev) // 1124