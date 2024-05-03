import matplotlib.pyplot as plt


def get_times(filename):
    with open(filename, "r") as file:
        return [float(line.strip()) for line in file.readlines()]


def get_average(list):
    return "{:.2f}".format(sum(list) / len(list))


simplex_time_list = get_times("Simplex.txt")
open_simplex_time_list = get_times("OpenSimplex.txt")

print(f"Simplex Noise Average Time: {get_average(simplex_time_list)}")
print(f"OpenSimplex Noise Average Time: {get_average(open_simplex_time_list)}")

plt.plot(simplex_time_list)
plt.plot(open_simplex_time_list)
plt.xlabel("Iteration", fontsize=16)
plt.ylabel("Time", fontsize=16)
plt.title("Speed comparison", fontsize=18)
plt.legend(["Simplex Noise", "OpenSimplex Noise"], fontsize=18)
plt.tight_layout()
plt.show()
