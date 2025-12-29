import numpy as np
from gelim.solver import GaussianSolver

def run_test_case(name, A, b):
    print(f"\n--- Solving {name} ---")
    solver = GaussianSolver(A, b)
    
    # Run without pivoting
    print("--- (a) No pivoting ---")
    ref_no_pivot = solver.eliminate(use_pivoting=False)
    print("Row Echelon Form:")
    print(np.round(ref_no_pivot, 2))
    result = solver.back_substitution(ref_no_pivot)
    print(f"Result: {result}")
    
    # Run with pivoting
    print("\n--- (b) With pivoting ---")
    ref_pivot = solver.eliminate(use_pivoting=True)
    print("Row Echelon Form:")
    print(np.round(ref_pivot, 2))
    result = solver.back_substitution(ref_pivot)
    print(f"Result: {result}")

if __name__ == "__main__":
    """
    Part of No. 16 of the Problem set 7.3 of Advanced Engineering Mathematics (Erwin) question
    is solve No. 11-14 questions utilizing the program written.
    """
    
    # Problem 11
    A11 = np.array([
        [0, 5, 5, -10],
        [2, -3, -3, 6],
        [4, 1, 1, -2]
    ])
    b11 = np.array([0, 2, 4])
    run_test_case("Problem 11", A11, b11)
    
    # Problem 12
    A12 = np.array([
        [2, -2, 4, 0],
        [-3, 3, -6, 5],
        [1, -1, 2, 0]
    ])
    b12 = np.array([0, 15, 0])
    run_test_case("Problem 12", A12, b12)
    
    # Problem 13
    #      10x  +4y  -2z = -4
    # -3w -17x   +y  +2z =  2
    #   w   +x   +y      =  6
    #  8w -34x +16y -10z =  4
    A13 = np.array([
        [0, 10, 4, -2],
        [-3, -17, 1, 2],
        [1, 1, 1, 0],
        [8, -34, 16, -10]
    ])
    b13 = np.array([-4, 2, 6, 4])
    run_test_case("Problem 13", A13, b13)
    
    # Problem 14
    A14 = np.array([
        [2, 3, 1, -11],
        [5, -2, 5, -4],
        [1, -1, 3, -3],
        [3, 4, -7, 2]
    ])
    b14 = np.array([1, 5, 3, -7])
    run_test_case("Problem 14", A14, b14)