import pandas as pd
import numpy as np
import pandas_datareader as web
import math
import os
import time

def get_data(ticker: str, start_date: str, end_date: str):
    df = web.DataReader(ticker, data_source='yahoo', start=start_date, end=end_date)
    return df

def get_stock_price_on_year(ticker: str, year: int) -> float:
    # TODO: Check why is this not working
    wdf = web.DataReader(ticker, data_source='yahoo', start=year, end=year)
    return wdf['Close'].mean()


# If the stock was at a price x then one month later it was at a price y > x then it was a buy
# If the stock was at a price x then one month later it was at a price y < x then it was a Sell
# If the stock was at a price x then one month later it was at a price y ≃ x then it was a Hold

# TODO: Separete the data in train and test (buy, sell, hold) -> 70% train, 30% test
# TODO: Check how was the params like P/E Ratio when stock was a buy, sell or hold and so one for other params
# That will be the input of the neural network

def read_csv_data(path: str) -> pd.DataFrame:

    df = pd.DataFrame()
    files = os.listdir(path)
    for file in files:
        if os.path.isfile(os.path.join(path, file)):
            date = file.split('_')[0]
            current_df = pd.read_csv(os.path.join(path, file))
            current_df.rename(columns={'Unnamed: 0': 'Stock'}, inplace=True)
            current_df['Date'] = date
            current_df = current_df[['Stock','PE ratio', 'EPS', 'PB ratio', 'Debt to Equity', 'Dividend Yield', 'ROE', 'Revenue Growth', 'Profit Margin', 'Market Cap', 'EBITDA', 'Date']]
            current_df = current_df.dropna()
            df = pd.concat([df, current_df], ignore_index=True)
    return df

def group_data_by_stock(df: pd.DataFrame, stock: str) -> pd.DataFrame:
    return df[df['Stock'] == stock]

if __name__ == '__main__':
    path = "./trainingData/"
    df = read_csv_data(path)
    df = group_data_by_stock(df, 'AAPL')
    price = get_stock_price_on_year('AAPL', 2019)
    print(price)
    print(df)
