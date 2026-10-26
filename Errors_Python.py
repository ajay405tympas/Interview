print("Hello, World!")

# Your solution here
chai_price = 15

def safe_chai_bill(cups, price):
  try:
    if not isinstance(cups, int):
        raise TypeError("cups must be an integer")
    elif (cups == 0):
        raise ValueError("cups must be a value > 0")
    
    return cups * price

  except TypeError as e:
  	print("error is", e)
  except ValueError as e:
  	print("error is", e)

print(safe_chai_bill("3", "15"))     # normal text input from a form
print(safe_chai_bill(2, 15))          # normal numbers
print(safe_chai_bill("two", 15))      # bad input -> caught
print(safe_chai_bill(0, 15)) 
