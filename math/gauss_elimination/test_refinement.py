import numpy as np
from gelim.solver import GaussianSolver

def run_refinement_test(name, A, b):
    """Test iterative refinement on a system"""
    print(f"\n{'='*70}")
    print(f"Testing: {name}")
    print(f"{'='*70}")
    
    solver = GaussianSolver(A, b)
    
    print("\n1. Solving system with Gaussian elimination...")
    ref_matrix = solver.eliminate(use_pivoting=True)
    x_initial = solver.back_substitution(ref_matrix)
    
    if isinstance(x_initial, str):
        print(f"   System is {x_initial}")
        return
    
    print(f"   Initial solution: {np.round(x_initial, 6)}")
    
    # initial residual
    initial_residual = b - A @ x_initial
    initial_norm = np.linalg.norm(initial_residual)
    print(f"   Initial residual norm: {initial_norm:.2e}")
    
    print("\n2. Applying iterative refinement...")
    x_refined = solver.iterative_refinement(x_initial, max_iterations=3)
    
    # final residual
    final_residual = b - A @ x_refined
    final_norm = np.linalg.norm(final_residual)
    
    print(f"\n3. Results:")
    print(f"   Refined solution: {np.round(x_refined, 6)}")
    print(f"   Final residual norm: {final_norm:.2e}")
    
    if final_norm > 0:
        improvement = initial_norm / final_norm
        print(f"   Improvement factor: {improvement:.2f}x")
    else:
        print(f"   Improvement: Perfect solution achieved!")

def create_poorly_scaled_system():
    """
    Create a system where initial solution has visible errors
    due to poor scaling, so refinement can demonstrate improvement
    """
    # Create a system with vastly different scales
    A = np.array([
        [1e10, 1.0, 1.0],
        [1.0, 1e10, 1.0],
        [1.0, 1.0, 1e10]
    ])
    b = np.array([1e10 + 2, 1e10 + 3, 1e10 + 4])
    return A, b

def create_vandermonde_system():
    """
    Vandermonde matrices are notoriously ill-conditioned
    """
    n = 5
    points = np.linspace(0, 1, n)
    A = np.vander(points, increasing=True)
    b = np.ones(n)
    return A, b

if __name__ == "__main__":
    print("\n" + "="*70)
    print("Sample Iterative Refinement")
    print("="*70)
    
    # Well-conditioned system (diagonal dominant)
    print("\n--- TEST 1: Well-conditioned system ---")
    A1 = np.array([
        [10.0, -1.0, 2.0],
        [-1.0, 11.0, -1.0],
        [2.0, -1.0, 10.0]
    ])
    b1 = np.array([6.0, 25.0, -11.0])
    run_refinement_test("Diagonal dominant matrix", A1, b1)
    
    # Moderately ill-conditioned (Hilbert matrix 3x3)
    print("\n--- TEST 2: Moderately ill-conditioned system ---")
    n = 3
    A2 = np.array([[1.0/(i+j+1) for j in range(n)] for i in range(n)])
    b2 = np.ones(n)
    run_refinement_test("3x3 Hilbert matrix", A2, b2)
    
    # More ill-conditioned (Hilbert matrix 4x4)
    print("\n--- TEST 3: Severely ill-conditioned system ---")
    n = 4
    A3 = np.array([[1.0/(i+j+1) for j in range(n)] for i in range(n)])
    b3 = np.ones(n)
    run_refinement_test("4x4 Hilbert matrix", A3, b3)
    
    # Problem 14 from Problem set 7.3 of Advanced Engineering Mathematics (Erwin) questions
    print("\n--- TEST 4: Problem from Kreyszig textbook ---")
    A4 = np.array([
        [2, 3, 1, -11],
        [5, -2, 5, -4],
        [1, -1, 3, -3],
        [3, 4, -7, 2]
    ], dtype=np.float64)
    b4 = np.array([1, 5, 3, -7], dtype=np.float64)
    run_refinement_test("Kreyszig Problem 14", A4, b4)
    
    # Poorly scaled system
    print("\n--- TEST 5: Poorly scaled system ---")
    A5, b5 = create_poorly_scaled_system()
    run_refinement_test("Poorly scaled diagonal matrix", A5, b5)
    
    # Vandermonde matrix
    print("\n--- TEST 6: Vandermonde matrix (severely ill-conditioned) ---")
    A6, b6 = create_vandermonde_system()
    run_refinement_test("5x5 Vandermonde matrix", A6, b6)
    