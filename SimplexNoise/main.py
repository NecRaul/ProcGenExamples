import random
import noise
import numpy as np
import matplotlib.pyplot as plt

size = 1024
scale = 256
octaves = 6
persistence = 0.5
lacunarity = 1.0

noise_map = np.zeros((size, size))
for i in range(size):
    for j in range(size):
        noise_map[i][j] = noise.snoise2(
            i / scale,
            j / scale,
            octaves=octaves,
            persistence=persistence,
            lacunarity=lacunarity,
            repeatx=1024,
            repeaty=1024,
            base=0,
        )

plt.imshow(noise_map, cmap="gray", origin="lower")
plt.title("Simplex Noise Map")
plt.tight_layout()
plt.savefig("Simplex_Noise_Map.png")
plt.close()
