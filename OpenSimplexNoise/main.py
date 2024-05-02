import random
import opensimplex
import numpy as np
import matplotlib.pyplot as plt

size = 1000
scale = 100

noise_map = np.zeros((size, size))
noise = opensimplex.OpenSimplex(seed=random.randrange(1))
for i in range(size):
    for j in range(size):
        noise_map[i, j] = noise.noise2(i / scale, j / scale)

plt.imshow(noise_map, cmap="gray", origin="lower")
plt.title("OpenSimplex Noise Map")
plt.tight_layout()
plt.savefig("OpenSimplex_Noise_Map.png")
plt.close()
