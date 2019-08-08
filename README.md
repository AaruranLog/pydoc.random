An recurrent neural network to artificially generate documentation for Python 3.
This project is NOT pure python (unfortunately for the purists) because I use golang
for web crawling and some parallelized data cleaning. That being said, it is
very possible to turn this into a 'pure' Python project using packages like BeautifulSoup.

# Plan:
- [x] Write a web crawler to fetch text from the python 3.7 docs
- [x] Create a data pipeline to turn that text into a learnable format.
- [x] Train a RNN (or LSTM or GRU) on the text.
- [x] Generate a tutorial using deep learning

Originally, the code uses a character-level approach. This method has terrible
performance, and training is still time-consuming on GPU.

Then, we applied the pretrained GPT-2 model. This approach was more successful.
Samples of the generated text can be found in the 'samples' folder. I've included
a sample of the best generated text below.
```
  The first line of Formatter class is defined by
  Formatter.setfield(value, type)
  If type is a sequence of
  strings that contains a field named value, the field is
  modified by the setfield() method.  If type is a single
  string, the field is modified by the setfield() method.  For example:
  >>> class _Field:
  ...     'fieldname' = '__class__'
  ...     'value' = 2


  This class is documented in section Field objects.
```
Impressively, the model was able to produce believeable Python code.

# Installation
To duplicate the results of this repository, I recommend pre-installing Go,
and Python 3.7+.

## The easiest way to duplicate my results
1. Upload the file 'notebook/gpt2_model.ipynb' to a Google Colab notebook.
2. Upload the file 'src/data/raw_corpus.tar.gz' to that notebook's workspace
3. Ensure the Colab runtime has GPU acceleration enabled*.
4. Run the notebook, and sit back for ~10-20 minutes as the model finetunes

## A Note on Environments
I use conda to manage my python packages, but because this
codebase is multi-lingual, there is no one-size-fits-all solution. However, the
data pipeline uses very few non-standard packages. As well, I recommend that the model training
and text generation be done on colab using 'notebook/gpt2_model.ipynb', which installs
one package into the cloud environment. In other words, this project is small enough
that I don't think it requires a dedicated virtual environment.

The go pipeline only requires you to install the package 'colly',
and its dependencies. The rest of the go packages used are standard.

However, if you wish to create your own python environment (with conda, or venv),
I have provided a requirements.txt file for your convenience.
Note: Model training is done using Google's colab, so although tensorflow and keras
are used in this codebase, they need not be installed locally.

I've included a compressed npz file which you can use to train your own models.
Alternatively, you can prepare the data yourself from scratch.

# To unpack the data from a compressed file

There is a file `src/data/raw_corpus.tar.gz` which you can unpack, either from
the command line or using the 'shutil' standard package in python. Make sure it
is unpacked into the directory `data/raw/corpus/` for the repo's compatibility.

# To download and prepare the data
## Automatically
After installing Go and python, run the script in src named 'prepare_data.sh',
from the root-directory level.
```
  chmod u+x src/prepare_data.sh
  ./src/prepare_data.sh
```
The last python script (txt-to-numpy.py) may fail, but if you run it a few times
on a device with sufficient memory, it should complete without too much trouble.

## Manually
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
