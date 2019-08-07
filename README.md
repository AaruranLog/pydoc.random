An recurrent neural network to artificially generate documentation for Python 3.
This project is NOT pure python (unfortunately for the purists) because I use golang
for web crawling and some parallelized data cleaning. That being said, it is
very possible to turn this into a 'pure' Python project using packages like BeautifulSoup.

# Plan:
- [x] Write a web crawler to fetch text from the python 3.7 docs
- [x] Create a data pipeline to turn that text into a learnable format.
- [ ] Train a RNN (or LSTM or GRU) on the text.
- [ ] Generate a tutorial using deep learning

Currently, the code uses a character-level approach. This method has terrible
performance, and training is still time-consuming. More model-tuning to be done.

# Installation
To duplicate the results of this repository, I recommend pre-installing Go,
and Python 3.7+.

I use conda to manage my python packages, but because this
codebase is multi-lingual, there is no one-size-fits-all solution.

The go pipeline only requires you to install the package 'colly',
and its dependencies. The rest of the go packages used are standard.

However, if you wish to create your own python environment (with conda, or venv),
I have provided a requirements.txt file for your convenience.
Note: Model training is done using Google's colab, so although tensorflow and keras
are used in this codebase, they need not be installed locally.

I've included a compressed npz file which you can use to train your own models.
Alternatively, you can prepare the data yourself from scratch.

## To prepare the data automatically
After installing Go and python, run the script in src named 'prepare_data.sh',
from the root-directory level.
```
  chmod u+x src/prepare_data.sh
  ./src/prepare_data.sh
```
The last python script (txt-to-numpy.py) may fail, but if you run it a few times
on a device with sufficient memory, it should complete without too much trouble.

## To prepare the data manually
1. Install the following go package to replicate the web crawling.
```
  go get github.com/gocolly/colly
```
2. Run the web scraping tool
```
  go run src/data/crawl.go
```
3. Process the data in the raw corpus into character tokens
```
  go run src/features/unique_chars.go -dir=data/raw/corpus
  go run src/features/encode_text.go -dir=data/raw/corpus -target=data/interim/cleaned_corpus
```
4. Transform the integerized character tokens into a compressed numpy matrix
```
  python src/features/count-chars.py
  python src/features/txt-to-numpy.py
```
