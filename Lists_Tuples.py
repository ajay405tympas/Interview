#print("Hello, World!")
#print("Hello, World!")

students = {
    "Aarav":   [85, 90, 78],
    "Priya":   [72, 68, 75],
    "Rohan":   [45, 52, 48],
    "Sneha":   [95, 92, 98],
    "Manish":  [60, 65, 70],
}

top = 0
min = 0
averages = []

for name, marks in students.items():
    #print(name)
    #print(len(marks))

    total = 0

    for i in range(len(marks)):
      total = total + marks[i]

    Avg = total/(len(marks))
    averages.append(Avg)

    students[name] = {
        "marks": marks,
        "average": Avg
    }

    if Avg > 80:
      title = "topper"
    elif Avg > 60 and Avg < 80:
      title = "pass"
    else:
      title = "NeedsImprovement"

    if top < Avg:
      top = Avg

    print(f"{name}: {Avg:.2f}--{title}")

#print(students)
#print(top)
#print(min)
print("===========Class Topper========")
for name, data in students.items():
  #print(f"{data["average"]:.2f}")
  #print(data["average"])
  if data["average"] == top:
      print(f"{name} with average {top}")

average_total = 0
print("===========Class Average========")
for name, data in students.items():
  average_total = average_total + data["average"]

print(average_total/len(students))


