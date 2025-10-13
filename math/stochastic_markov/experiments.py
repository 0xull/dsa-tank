import numpy as np
import matplotlib.pyplot as plt
from numpy.typing import NDArray
from markov.chain import MarkovChain

def plot_land_use_evolution(history: list[NDArray[np.float64]]):
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
    plt.savefig('land_use_evolution.png')
    plt.show()
    

def run_example_13():
    """
    Reproduces and visualizes the results from the Advanced Engineering 
    mathematics textbook's Example 13
    """
    print("Run calculation and plot the resulting graph for Example 13...")
    
    A = np.array([[0.7,0.1,0.0], [0.2,0.9,0.2], [0.1,0.0,0.8]])
    x0 = np.array([[25], [20], [55]])
    
    chain = MarkovChain(transition_matrix=A, initial_state=x0)
    history = chain.run(num_steps=15) # This is equivalent to 75years of land usage
    
    print("\nInitial State (2004):")
    print(history[0])
    print(f"\nFinal State after {len(history)-1} steps:")
    print(history[-1])
    
    plot_land_use_evolution(history=history)

if __name__ == "__main__":
    run_example_13()
    