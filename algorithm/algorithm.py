import pandas as pd
import numpy as np
import pandas_datareader as web
from sklearn.preprocessing import MinMaxScaler
from keras.models import Sequential
import math

def get_data(ticker: str, start_date: str, end_date: str):
    df = web.DataReader(ticker, data_source='yahoo', start=start_date, end=end_date)
    return df



