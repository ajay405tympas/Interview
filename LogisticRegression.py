
import pandas as pd
import numpy as np

from sklearn.model_selection import train_test_split
from sklearn.linear_model import LogisticRegression
from sklearn.metrics import (
		accuracy_score,
        precision_score,
        confusion_matrix,
        classification_report,
        recall_score,
        f1_score
        )
 
data = {
    "Hours": [1,2,3,4,5,6,7,8,9,10],
    "Pass":  [0,0,0,0,1,1,1,1,0,0]
}

df = pd.DataFrame(data)



#df["Hours"] = df["Hours"].fillna(df["Hours"].mean())
#df["Pass"] = df["Pass"].fillna(df["Pass"].mean())

df.dropna(inplace = True)
print(df)

X = df[["Hours"]]
y = df["Pass"]



print(X)

print(y)

X_train,X_test, y_train, y_test = train_test_split( X, y, test_size = 0.4, random_state =42)


model = LogisticRegression()

model.fit(X_train,y_train)

y_pred = model.predict(X_test)

print("y_pred",y_pred)

ac_score = accuracy_score(y_test,y_pred)

print("ac_score:", ac_score)

Prec_score = precision_score(y_test,y_pred)

print("Prec_score:", Prec_score)


conf_mat = confusion_matrix(y_test,y_pred)

print("conf_mat:", conf_mat)

classification_report = classification_report(y_test,y_pred)

print("classification_report:", classification_report)

rec_score = recall_score(y_test,y_pred)

print("rec_score", rec_score)

f1_score = f1_score(y_test,y_pred)

print("f1_score", f1_score)












print("Hello, World!")

x = "Python"
y = "is"
z = "awesome"
print(x, y, z)



