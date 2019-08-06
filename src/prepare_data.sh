#!/bin/bash

go get github.com/gocolly/colly
go run src/data/crawl.go
go run src/features/unique_chars.go -dir=data/raw/corpus
go run src/features/encode_text.go -dir=data/raw/corpus -target=data/interim/cleaned_corpus
python count-chars.py
python txt-to-numpy.py
