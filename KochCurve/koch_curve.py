import math
import matplotlib.pyplot as plt


def forward(x, y, angle):
    new_x = x + math.cos(math.radians(angle))
    new_y = y + math.sin(math.radians(angle))
    plt.plot([x, new_x], [y, new_y], color="black")
    return new_x, new_y


def turn(initial_angle, direction):
    turn_angle = 90 if direction == "+" else -90
    return (initial_angle + turn_angle) % 360


def generate_koch_curve(axiom, rules):
    new_axiom = ""
    for char in axiom:
        new_axiom += rules.get(char, char)
    return new_axiom


axiom = "F"
rules = {"F": "F+F-F-F+F"}
iterations = 5
x, y = 0, 0
angle = 0


for iteration in range(iterations + 1):
    points = [(x, y)]
    for char in axiom:
        if char == "F":
            x, y = forward(x, y, angle)
            points.append((x, y))
        else:
            angle = turn(angle, char)

    plt.plot(*zip(*points), color="black")
    plt.title(f"Iteration {iteration}")
    plt.axis("off")
    plt.savefig(f"Iteration_{iteration}.png")
    plt.close()

    axiom = generate_koch_curve(axiom, rules)
