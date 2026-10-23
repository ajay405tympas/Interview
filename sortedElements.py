from collections import Counter

nums = [1,1,1,2,2,3]

freq = Counter(nums)

print("Result:", sorted(nums,key=lambda x:(freq[x], -x)))


##Sorting strings

str = [ "ajay", "Ajay", "X", "x-mas", "--jiju"]

print ( "Result sorted strings:", sorted(str))

##Sorting Tuples 

t = ( 10 ,42 ,30)
print( "Sort tuples", sorted(t))

lst = [10 ,42 , "ajay", 30]
#print( "Sort list", sorted(lst)) error

data = [("jay", 20), ("Fox", 34), ("Bear", 87)]
print ("sort list of tuples", sorted(data))

print ("sort list of tuples with value", sorted(data, key = lambda x: x[1]))

print ("sort list of tuples with reverse", sorted(data, key = lambda x: x[1], reverse=True))

lst = [1,1,1,1,3,4,4,5,5,5,2]
freq = Counter(lst)
print (freq)

print("lstArray", sorted(lst, key = lambda x:(freq[x], -x)))

fruits = ["banana", "apple", "kiwi", "orange", "ber"]
print( sorted(fruits))

print(sorted (fruits, key = len))




print("Hello, World!")
