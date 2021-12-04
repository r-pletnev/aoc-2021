def main(file_path: str) -> int:
    a = 0
    b = 0
    c = 0
    summary = 0

    with open(file_path, mode="r") as fin:
        while fin:
            try:
                sum_prev = a + b + c
                if sum_prev == 0:
                    a = int(fin.readline())
                    b = int(fin.readline())
                    c = int(fin.readline())
                    sum_curr = a + b + c
                else:
                    a1 = b
                    b1 = c
                    c1 = int(fin.readline())

                    sum_curr = a1 + b1 + c1

                    a = a1
                    b = b1
                    c = c1

                if sum_prev > 0 and sum_curr > sum_prev:
                    summary += 1

            except ValueError:
                break

    print(summary)


if __name__ == "__main__":
    main("./input.txt")
