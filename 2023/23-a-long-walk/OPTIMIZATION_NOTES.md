# Performance Optimization Notes for Day 23 Part 2

## Current Status
- Fixed: `slices.Index` bottleneck (O(n) -> O(1) with map)
- Still too slow: Working on raw maze grid instead of compressed graph

## Graph Contraction Strategy

### The Problem
The maze has lots of long corridors where there's only one way to go. Processing every cell individually is wasteful.

### The Solution
1. **Find Junctions**: Cells where paths split (>2 walkable neighbors), plus start/end
2. **Compress Corridors**: Treat entire corridors between junctions as single weighted edges
3. **Build Compressed Graph**: Nodes = junctions, Edges = corridor distances
4. **Run DFS on Small Graph**: Instead of 6000+ cells, maybe 30-40 junction nodes

### Implementation Steps

```go
type Junction struct {
    pos   [2]int
    edges map[[2]int]int  // neighbor junction -> distance to it
}

// Step 1: Find all junctions
func findJunctions(maze Maze) map[[2]int]*Junction {
    junctions := map[[2]int]*Junction{}

    for y := 1; y < len(maze)-1; y++ {
        for x := 1; x < len(maze[0])-1; x++ {
            if maze[y][x] == '#' {
                continue
            }
            neighbors := countWalkableNeighbors(maze, x, y)
            if neighbors > 2 {  // It's a junction!
                junctions[[2]int{x,y}] = &Junction{
                    pos: [2]int{x,y},
                    edges: map[[2]int]int{},
                }
            }
        }
    }

    // Also add start and end as special junctions
    start := [2]int{1, 0}
    end := [2]int{len(maze[0])-2, len(maze)-1}
    junctions[start] = &Junction{pos: start, edges: map[[2]int]int{}}
    junctions[end] = &Junction{pos: end, edges: map[[2]int]int{}}

    return junctions
}

// Step 2: For each junction, walk corridors to find connected junctions
func buildGraph(maze Maze, junctions map[[2]int]*Junction) {
    for jPos, junction := range junctions {
        // For each direction from this junction
        for _, dir := range [][2]int{{0,1}, {1,0}, {0,-1}, {-1,0}} {
            x, y := jPos[0] + dir[0], jPos[1] + dir[1]

            // Skip walls or out of bounds
            if !isWalkable(maze, x, y) {
                continue
            }

            // Walk the corridor until we hit another junction
            distance := 1
            prevX, prevY := jPos[0], jPos[1]

            for {
                // Found another junction?
                if otherJunction, exists := junctions[[2]int{x,y}]; exists {
                    junction.edges[otherJunction.pos] = distance
                    break
                }

                // Continue down the corridor
                neighbors := getWalkableNeighbors(maze, x, y)

                // Remove the direction we came from
                var next [2]int
                found := false
                for _, n := range neighbors {
                    if n[0] != prevX || n[1] != prevY {
                        next = n
                        found = true
                        break
                    }
                }

                if !found {
                    break // Dead end
                }

                prevX, prevY = x, y
                x, y = next[0], next[1]
                distance++
            }
        }
    }
}

// Step 3: Run DFS on the compressed graph
func longestPath(junctions map[[2]int]*Junction, start, end [2]int) int {
    visited := map[[2]int]bool{}

    var dfs func([2]int, int) int
    dfs = func(pos [2]int, dist int) int {
        if pos == end {
            return dist
        }

        visited[pos] = true
        defer delete(visited, pos)  // Backtrack

        maxDist := 0
        for neighbor, edgeWeight := range junctions[pos].edges {
            if !visited[neighbor] {
                maxDist = max(maxDist, dfs(neighbor, dist + edgeWeight))
            }
        }

        return maxDist
    }

    return dfs(start, 0)
}
```

### Expected Performance
- Raw maze: ~6000+ cells to explore
- Compressed graph: ~30-40 junctions
- Runtime: Hours -> Seconds

### Additional Optimizations (if still needed)
- Use bitsets instead of map[bool] for visited tracking
- Early pruning if current_path + max_remaining < best_so_far
- Parallel exploration of top-level branches

## Notes from Dec 7, 2025 debugging session
- Original bottleneck: `slices.Index` for visited checking (O(n) per check)
- Fixed with map[[2]int]bool for O(1) lookups
- Next step: Implement graph contraction above
