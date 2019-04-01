NAME = psphinx

SRCS = sphinx.c

SFLAGS = -DMODELDIR=\"`pkg-config --variable=modeldir pocketsphinx`\" \
    `pkg-config --cflags --libs pocketsphinx sphinxbase`

all: $(NAME)

$(NAME):
	gcc $(SRCS) $(SFLAGS) -o $(NAME)

re:
	rm $(NAME)
	make all

setup:
	brew tap watsonbox/cmu-sphinx
	brew install --HEAD watsonbox/cmu-sphinx/cmu-sphinxbase
	brew install --HEAD watsonbox/cmu-sphinx/cmu-sphinxtrain
	brew install --HEAD watsonbox/cmu-sphinx/cmu-pocketsphinx
