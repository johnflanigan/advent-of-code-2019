def main():
    with open('input/day1.txt') as file:
        lines = file.readlines()

        total_fuel = 0

        for line in lines:
            mass = int(line)
            total_fuel += calculate_fuel_required(mass)

        print(total_fuel)


def calculate_fuel_required(mass):
    fuel = (mass // 3) - 2

    if fuel <= 0:
        return 0
    else:
        return fuel + calculate_fuel_required(fuel)


if __name__ == "__main__":
    main()
