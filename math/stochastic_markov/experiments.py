import numpy as np
import matplotlib.pyplot as plt
from markov.chain import MarkovChain

def plot_land_use_evolution(history):
    """Plots the specific land-use history."""
    labels = ["Commercial", "Industrial", "Residentials"]
    history_array = np.array(history).squeeze()

    plt.figure(figsize=(10, 6))
    for i in range(history_array.shape[1]):
        plt.plot(history_array[:, i], label=labels[i], marker='.', linestyle='-')

    plt.title('Evolution of Land Use Over 5-Year Intervals')
    plt.xlabel('Step (5-Year Interval)')
    plt.ylabel('Percentage (%)')
    plt.grid(True)
    plt.legend()
    plt.show()
    

def run_example_13():
    "Reproduces and visualizes "