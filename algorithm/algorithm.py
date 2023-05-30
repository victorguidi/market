import pandas as pd
import numpy as np
import pandas_datareader as web
from sklearn.preprocessing import MinMaxScaler
from keras.models import Sequential
import math

def get_data(ticker: str, start_date: str, end_date: str):
    df = web.DataReader(ticker, data_source='yahoo', start=start_date, end=end_date)
    return df

# TODO: Divide the data in Buy, Sell, Hold -> 1, 0, 0.5
# If the stock was at a price x then one month later it was at a price y > x then it was a buy
# If the stock was at a price x then one month later it was at a price y < x then it was a Sell
# If the stock was at a price x then one month later it was at a price y ≃ x then it was a Hold

# TODO: Separete the data in train and test (buy, sell, hold) -> 70% train, 30% test
# TODO: Check how was the params like P/E Ratio when stock was a buy, sell or hold and so one for other params
# That will be the input of the neural network

