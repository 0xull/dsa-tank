import numpy as np
from numpy.typing import NDArray

def is_stochastic(matrix: NDArray[np.float64]):
    """Checks if a matrix is a valid column-stochastic matrix."""
    return np.allclose(np.sum(matrix, axis=0), 1)

def is_square(matrix: NDArray[np.float64]) -> bool:
    """Checks if a NumPy array is a square matrix."""
    return matrix.ndim == 2 and matrix.shape[0] == matrix.shape[1]

