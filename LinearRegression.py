import pandas as pd
import numpy as np

from sklearn.linear_model import LinearRegression
from sklearn.model_selection import train_test_split
from sklearn.metrics import (
	mean_absolute_error,
    mean_squared_error,
    r2_score
    )
    
data = {
    "Area": [1000, 1200, 1500, 1800, 2000,
             2200, 2500, 2800, 3000, 3500],
    "Price": [200000,np.nan, 280000, 320000, 350000,
              390000, 450000, 500000, 540000, 620000]
}

df = pd.DataFrame(data)

print(data)


df["Area"] = df["Area"].fillna(df["Area"].mean())
df["Price"] = df["Price"].fillna(df["Price"].mean())

print(df)


X = df[["Area"]]
y = df[["Price"]]

#split data set 

X_train, X_test, y_train,y_test = train_test_split(X,y, test_size = 0.3, random_state = 42)

print("X_TrainDataSet", X_train)
print("X_TestDataSet",X_test)
print("y_TrainDataSet",y_train)
print("y_TestDataSet",y_test)

model = LinearRegression()
model.fit(X_train,y_train)

print("intercept", model.intercept_)
print("Coeff", model.coef_)

y_pred = model.predict(X_test)

print("y_Pred", y_pred)

mae = mean_absolute_error(y_test, y_pred)

mse = mean_squared_error(y_test, y_pred)

rmse = np.sqrt(mse)

r2s = r2_score(y_test,y_pred)

print("MAE",mae)

print("MSE",mse)

print("RMSE",rmse)

print("r2_score",r2s)







print("Ajay")

