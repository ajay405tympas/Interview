'''You run a chai stall. Customers order multiple cups of chai, sometimes with add-ons (extra ginger, masala, biscuits). Write a function to calculate their bill. Your task: Write a function chai_bill(cups, *addons, **discount) that:

Charges ₹15 per cup of chai Adds ₹5 for each addon passed via *args (like "ginger", "masala", "biscuit") Applies a discount only if discount_percent is passed via **kwargs Returns the final bill amount

Test it with these calls:

print(chai_bill(2))
print(chai_bill(3, "ginger", "masala"))
print(chai_bill(4, "biscuit", discount_percent=10))
print(chai_bill(5, "ginger", "masala", "biscuit", discount_percent=20))

Expected output:

30
55
58.5
112.0'''

chai_price = 15
ginger_price = 5
masala_price = 5
biscuit_price = 5

def chai_bill (cups, *addons, **discount):
  total_price = 0
  total_chai_price = cups*chai_price
  

  total_price = total_price + total_chai_price
  #print("Chai price", total_price)
  
  total_ginger_price = 0
  total_masala_price = 0
  total_biscuit_price = 0

  for i in addons:
    if i == "ginger":
      total_ginger_price =  total_ginger_price + ginger_price
    if i == "masala":
      total_masala_price =  total_masala_price + masala_price
    if i == "biscuit":
      total_biscuit_price =  total_biscuit_price + biscuit_price

  total_price = total_price + total_ginger_price + total_masala_price + total_biscuit_price
  
  #print("Addon + Chai price", total_price)

  for key,value in discount.items():
    total_price = total_price-((total_price*value) / 100)

  print (total_price)
    

chai_bill(2)
chai_bill(3, "ginger", "masala")
chai_bill(4, "biscuit", discount_percent=10)
chai_bill(5, "ginger", "masala", "biscuit", discount_percent=20)