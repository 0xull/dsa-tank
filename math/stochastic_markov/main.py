import numpy as np
from numpy.typing import NDArray

from markov.chain import MarkovChain

def get_matrix_from_user(prompt: str) -> NDArray[np.float64]:
    """Interactively prompts the user to enter a square matrix."""
    print(f"\n--- {prompt} ---")
    first_row_str = input("Enter the first row (comma-separated values): ")
    first_row = [float(val) for val in first_row_str.split(',')]
    n = len(first_row)
    
    matrix_data = [first_row]
    print(f"Detected a {n}x{n} matrix. Please enter the {n-1} rows.")
    
    for i in range(1, n):
        row_str = input(f"Enter row {i+1}: ")
        row = [float(val) for val in row_str.split(',')]
        if len(row) != n:
            raise ValueError(f"Error: Row {i+1} has {len(row)} elements, but expected {n}.")
        matrix_data.append(row)
    return np.array(matrix_data)

def get_vector_from_user(size: int):
    """Interactively prompts the user for a column vector"""
    print("\n--- Initial State Vector ---")
    vec_str = input(f"Enter the {size} values for the initial state vector (comma-separated): ")
    vec = [float(val) for val in vec_str.split(',')]
    if len(vec) != size:
        raise ValueError(f"Error: Vector has {len(vec)} elements, but expected {size}.")
    return np.array(vec).reshape(size, 1)
    
def main():
    try:
        transition_matrix = get_matrix_from_user("Transition Matrix A")
        initial_state = get_vector_from_user(transition_matrix.shape[0])
        num_steps = int(input("\nEnter the number of steps to calculate: "))
        
        chain = MarkovChain(transition_matrix=transition_matrix, initial_state=initial_state)
        chain.run(num_steps=num_steps)
        
        print("\n--- Calculation Result ---")
        print("Initial state: ")
        print(chain.history[0])
        print(f"\nState after {num_steps} steps: ")
        print(chain.current_state)
    except ValueError as e:
        print(f"\nInput Error: {e}")
    except Exception as e:
        print(f"\nAn unexpected error occurred: {e}")

if __name__ == "__main__":
    main()