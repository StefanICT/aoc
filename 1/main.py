with open("1/measurements1.1") as fp:
    measurements = map(int, fp.read().splitlines())

    increased = 0
    previous: int = None
    for measurement in measurements:
        if previous and previous < measurement:
            increased += 1

        previous = measurement

    print(f"Day 1.1: {increased}")

with open("1/measurements1.2") as fp:
    measurements = list(map(int, fp.read().splitlines()))

    increased = 0
    previous: int = None
    for i in range(len(measurements)):
        if i + 1 >= len(measurements) or i + 2 >= len(measurements):
            break

        measurement = measurements[i] + measurements[i + 1] + measurements[i + 2]
        if previous and previous < measurement:
            increased += 1

        previous = measurement

    print(f"Day 1.2: {increased}")
