#!/usr/bin/env python
# coding: utf-8

# source/inspiration: https://github.com/keras-team/keras/blob/master/examples/lstm_text_generation.py
import numpy as np
import os
import pandas as pd
from tqdm import tqdm

maxlen = 40 # tunable parameter, but fine for now
number_of_characters = 104
phrases = []
next_chars = []
source_directory = 'data/interim/cleaned_corpus'
print('Converting interim data to numpy arrays...')
for f in tqdm(os.listdir(source_directory)):
    file_name = source_directory + '/' + f
    with open(file_name, 'r') as src:
        body = src.read()
        str_tokens = body.split()
        tokens = [int(i) for i in str_tokens]
        for i in range(len(tokens) - maxlen):
            phrases.append(tokens[i:i + maxlen])
            next_chars.append(tokens[i+maxlen])
print('Creating numpy arrays...')
x_data = np.array(phrases)
y_data = np.array(next_chars)
print('Saving to npz file...')
target_file = 'data/processed/compressed_data.npz'
np.savez_compressed(target_file, x_data=x_data, y_data=y_data)
print('Done.')
