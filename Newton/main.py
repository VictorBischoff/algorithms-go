
guess = 7
xn = float(guess)

def f(x):
  return x**3-x-1/51
  
def fm(x):
  return 3*x**2-1
  
for omgang in range(10):
    xn = xn -f(xn)/fm(xn)
    print(f"Rod: {xn} in omgang: {omgang}")