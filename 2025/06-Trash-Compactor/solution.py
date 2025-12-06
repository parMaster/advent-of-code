# Claude code translation from Go to Python, just to study how it works. And measure performance.
def solve(filename):
    with open(filename) as f:
        lines = f.read().split('\n')

    # Part 1: Read problems column-by-column (left-to-right)
    operators = lines[-1].split()  # Last line has operators
    number_lines = lines[:-1]      # All other lines have numbers

    # Parse each line into list of integers
    columns = [line.split() for line in number_lines]

    # Transpose to get columns, then calculate each problem
    problems = list(zip(*columns))  # zip(*) transposes!

    results = []
    for i, problem in enumerate(problems):
        nums = [int(n) for n in problem]  # Convert strings to ints
        op = operators[i]

        # Apply the operator to all numbers in the problem
        if op == '*':
            result = 1
            for n in nums:
                result *= n
        else:  # op == '+'
            result = sum(nums)

        results.append(result)

    part1 = sum(results)

    # Part 2: Read right-to-left, column by column
    # For each character position (right to left), read top to bottom
    max_width = max(len(line) for line in lines)
    part2 = 0

    current_problem_nums = []

    for col_idx in range(max_width - 1, -1, -1):  # Right to left
        # Get the operator for this column (from the last row)
        operator = lines[-1][col_idx] if col_idx < len(lines[-1]) else ' '

        # Get the digits from all other rows (top to bottom)
        num_chars = []
        for row_idx in range(len(lines) - 1):  # Exclude operator row
            line = lines[row_idx]
            if col_idx < len(line):
                num_chars.append(line[col_idx])
            else:
                num_chars.append(' ')

        # Form the number from this column
        num_str = ''.join(num_chars).strip().replace(' ', '')

        # Check if this column has an operator
        if operator in ['*', '+']:
            # Add the current column's number if it exists
            if num_str:
                current_problem_nums.append(int(num_str))

            # Calculate result for current problem
            if current_problem_nums:
                if operator == '*':
                    result = 1
                    for n in current_problem_nums:
                        result *= n
                else:
                    result = sum(current_problem_nums)
                part2 += result
                current_problem_nums = []

        elif operator.strip() == '':
            # Empty column separates problems, but save any number first
            if num_str:
                current_problem_nums.append(int(num_str))

        else:
            # Regular number column
            if num_str:
                current_problem_nums.append(int(num_str))

    return part1, part2


if __name__ == '__main__':
    import time
    start = time.time()

    print("Day 06: Trash Compactor")
    p1, p2 = solve("../aoc-inputs/2025/06/input.txt")
    # p1, p2 = solve("input-pub.txt")

    print(f"\tPart One: {p1}")  # 6757749566978
    print(f"\tPart Two: {p2}")  # 10603075273949
    print(f"Done in {time.time() - start:.3f} seconds")
