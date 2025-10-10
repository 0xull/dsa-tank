import numpy as np
from numpy.typing import NDArray
from . import utils

class MarkovChain:
    """A class to represent and run a discrete-time Markov process."""

    def __init__(self, transition_matrix: NDArray[np.float64], initial_state: NDArray[np.float64]):
        """
        Initializes the Markov chain.

        Args:
            transition_matrix (np.ndarray): The stochastic transition matrix.
            initial_state (np.ndarray): The initial state vector.
        """

        if not utils.is_square(transition_matrix):
            raise ValueError("The transition matrix must be square.")

        if not utils.is_stochastic(transition_matrix):
            raise ValueError("The provided matrix is not stochastic. All columns must sum to 1.")

        if transition_matrix.shape[1] != initial_state.shape[0]:
            raise ValueError("Matrix and state vector dimensions are not compatible.")

        self.transition_matrix: NDArray[np.float64] = transition_matrix
        self.current_state: NDArray[np.float64] = initial_state
        self.history: list[NDArray[np.float64]] = [initial_state]

    def step(self):
        """Advances the chain by one step."""
        self.current_state = self.transition_matrix @ self.current_state
        self.history.append(self.current_state)
        return self.current_state

    def run(self, num_steps: int):
        """Runs the chain for a given number of steps"""
        for _ in range(num_steps):
            _ = self.step()
        return self.history

    def find_steady_state(self):
        """Finds the steady-state vector (eigenvector for eigenvalue 1)."""
        eigenvalues, eigenvectors = np.linalg.eig(self.transition_matrix)

        # Find the index of the eigenvalue that is close to 1
        steady_state_index = np.isclose(eigenvalues, 1).nonzero()[0][0]
        steady_state_vector = eigenvectors[:, steady_state_index].real

        # Normalize to ensure it sums to 1 (making it a probability vector)
        return steady_state_vector / np.sum(steady_state_vector)
