import numpy as np

# the transition matrix A from the example
A = np.array([
    [0.7, 0.1, 0.0],
    [0.2, 0.9, 0.2],
    [0.1, 0.0, 0.8]
])

# the initial state vector x_2004
x_2004 = np.array([
    [25],
    [20],
    [55]
])

# the 2009 state
x_2009 = A @ x_2004

print("Transition Matrix A:\n", A)
print("\nState Vector for 2004 (x_2004):\n", x_2004)
print("\nCalculated State Vector for 2009 (x_2009):\n", x_2009)