with open("2/commands") as fp:
    horizontal = 0
    depth = 0

    for line in fp.read().splitlines():
        action, n = line.split(" ")

        if action == "forward":
            horizontal += int(n)
        elif action == "down":
            depth += int(n)
        elif action == "up":
            depth -= int(n)

    print(f"Day 2.1: {horizontal * depth}")

with open("2/commands") as fp:
    horizontal = 0
    depth = 0
    aim = 0

    for line in fp.read().splitlines():
        action, n = line.split(" ")

        if action == "forward":
            horizontal += int(n)
            depth += aim * int(n)
        elif action == "down":
            aim += int(n)
        elif action == "up":
            aim -= int(n)

    print(f"Day 2.2: {horizontal * depth}")
