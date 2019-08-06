#!/usr/bin/env python
# coding: utf-8

import pandas as pd

print('Reading counts file ...')
chars = pd.read_csv('src/features/counts.csv')
# Again, by inspection, these characters are (generally) rarely occurring.
# We might want to keep the capital 'K', and the hyphen ('-').
# For now, we will simply reject everything with frequency less than 14.
relevant_chars = chars[chars['count'] >= 14]
relevant_chars = relevant_chars.sort_values('count', ascending=False)
relevant_chars['indices'] = [i+1 for i in range(len(relevant_chars))]
relevant_chars = relevant_chars.drop('count', axis=1)
print('Saving map to file...')
relevant_chars.to_csv("src/features/char_int_map.csv", index=False)
print('Done.')
