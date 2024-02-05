# 🐜 lem-in 
**lem-in** is a project that meant to make a digital version of an ant farm. 

![](https://cdn.lospec.com/gallery/ant-hill-671934.gif)

## How it works?

The goal of this project is to find the quickest way to get n ants across the farm. Obviously, there are some basic constraints. To be the first to arrive, n ants will need to take the shortest path (and that isn’t necessarily the simplest). They will also need to avoid traffic jams as well as walking all over their fellow ants. So we use DFS (Depth First Search) algorithm to find the shortest path and then we simulate the ants to move through the path.

## ✒️ How to audit the project?

Almost for all test cases I already written a shell scripts, so you can just run it and see the result. For example:
```shell
$ ./runNormalTestCases.sh

🚀 Run normal test cases!
📑 Example 00:
4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1

L1-2
L1-3 L2-2
L1-1 L2-3 L3-2
L2-1 L3-3 L4-2
L3-1 L4-3
L4-1
➡️ Took steps: 6, expected <=6

📑 Example 01:
...
```
