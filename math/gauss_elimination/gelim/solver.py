import numpy as np
import random
from numpy.typing import NDArray

class GaussianSolver:
    def __init__(self, A: NDArray[np.float64], b: NDArray[np.float64]) -> None:
        """
        Initialize with coefficient matrix A and constant vector b.
        Constructs the Augmented Matrix [A|b]
        """
        
        self.A = A.astype(np.float64) # M x N shape
        self.b = b.astype(np.float64) # N x 1 shape
        
        if self.A.shape[0] != self.b.shape[0]:
            raise ValueError("Dimensions of A and b do not match.")
            
        self.augmented = np.hstack((self.A, self.b.reshape(-1, 1))) # M x N+1 shape
        self.rows, self.cols = self.augmented.shape
        # The number of unknowns which is cols - 1 (since last column accounts for 'b')
        self.n = self.cols-1 
       
    def eliminate(self, use_pivoting: bool = True) -> NDArray[np.float64]:
        """
        Performs Gaussian Elimination to transform the matrix into Row Echelon Form.
        
        Args:
            use_pivoting (bool): If True, swaps rows to put largest absolute value on diagonal.
                                 If False, proceeds blindly
        """
        
        M = self.augmented.copy()
        
        # min accounts for pivot columns for undetermined and overdetermined systems
        for k in range(min(self.rows, self.n)):
            if use_pivoting:
                max_index = np.argmax(np.abs(M[k:, k])) + k
                
                if k != max_index:
                    M[[k, max_index]] = M[[max_index, k]]
            
            pivot = M[k, k]
            if np.isclose(pivot, 0):
                continue
            
            for i in range(k+1, self.rows):
                factor = M[i, k] / pivot
                M[i, k:] -= factor * M[k, k:]
                # To account for float ops quirkness
                if abs(M[i, k]) < 1e-10:
                    M[i, k] = 0.0
            
        return M
    
    def back_substitution(self, ref_matrix: NDArray[np.float64]):
        """
        Perform Back Substitution on a Row Echelon Form Matrix.
        Returns the solution vector or raise error on Infinite/No solution.
        """
        
        x = np.zeros(self.n)
        rows = self.rows
        
        for i in range(self.n-1, -1, -1):
            rows-= 1
            if np.allclose(ref_matrix[rows, :-1], 0) and not np.isclose(ref_matrix[rows, -1], 0):
                return "Inconsistent system (No Solution)"
            
            pivot = ref_matrix[rows, i]
            if np.isclose(pivot, 0):
                # Dependent system (Infinitely Many Solution), meaning x_i is a free variable.
                # pick a constant value for it, say 0.0 or 1.0
                x[i] = random.choice([0.0, 1.0, 2.0])
                print(f"free variable, x[{i}]: {x[i]}")
                continue
            
            sum_ax = np.dot(ref_matrix[rows, i+1:self.n], x[i+1:])
            x[i] = (ref_matrix[rows, -1] - sum_ax) / pivot
        
        return x

    def iterative_refinement(self, x: NDArray[np.float64], max_iterations: int = 5, tol: float = 1e-10) -> NDArray[np.float64]:
        """
        Improves solution accuracy using iterative refinement.
        
        Process:
        1. Compute residual: r = b - Ax
        2. Solve correction: A * delta_x = r
        3. Update solution: x_new = x + delta_x
        4. Repeat until residual is small enough or max iterations reached
        
        Args:
            x: Initial solution from back_substitution
            max_iterations: Maximum number of refinement iterations
            tol: Convergence tolerance for residual norm
            
        Returns:
            Refined solution vector
        """
        
        x_refined = x.copy()
        
        for iteration in range(max_iterations):
            residual = self.b - self.A @ x_refined
            
            # Check convergence
            residual_norm = np.linalg.norm(residual)
            print(f"  Iteration {iteration}: residual norm = {residual_norm:.2e}")
            
            if residual_norm < tol:
                print(f"  Converged after {iteration} iterations")
                break
            
            # Solve A * delta_x = residual for the correction
            correction_solver = GaussianSolver(self.A, residual)
            ref_matrix = correction_solver.eliminate(use_pivoting=True)
            delta_x = correction_solver.back_substitution(ref_matrix)
            
            if isinstance(delta_x, str):
                print(f"  Refinement failed: {delta_x}")
                break
            
            x_refined = x_refined + delta_x
        
        return x_refined