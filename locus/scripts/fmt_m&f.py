OldData = []

NewData = []

i = 0
currentClass = ""
for item in OldData:
  # Get the class name and check if NewData already has it
  if "Male and Female at" in item["Male & Female Item"]:
    fetchedClassName = item["Male & Female Item"].split(" - ")[-1]
    # Make sure this class doesnt already exist in NewData
    classAlreadyExists = False
    for classes in NewData:
      if classes["class_name"] == fetchedClassName:
        classAlreadyExists = True
        break

    if classAlreadyExists == False:
      print(fetchedClassName)
      NewData.append({
        "class_name": fetchedClassName,
        "class_items": []
      })
    currentClass = fetchedClassName
  else:
    # Create item object
    itemObject = {
      "id": i,
       "item": item["Male & Female Item"],
       "item_class": currentClass,
       "item_id": item["ID/Texture"].split("/")[0],
       "texture_id": item["ID/Texture"].split("/")[1]
    }
    
    # Add item object to the class once
    for classes in NewData:
      if classes["class_name"] == currentClass:
        classes["class_items"].append(itemObject)
        break

  i += 1

with open("fmt_m&f.json", "w") as file:
  file.write(str(NewData))