NAME = psphinx

SRCS = sphinx.c helpers.c

SFLAGS = -DMODELDIR=\"`pkg-config --variable=modeldir pocketsphinx`\" \
    `pkg-config --cflags --libs pocketsphinx sphinxbase`

I = -I ./

all: $(NAME)

$(NAME):
	gcc $(I) $(SRCS) $(SFLAGS) -o $(NAME)

debug:
	gcc $(I) $(SRCS) $(SFLAGS) -o $(NAME) -fsanitize=address -g

re:
	rm $(NAME)
	make all

setup:
	brew tap watsonbox/cmu-sphinx
	brew install --HEAD watsonbox/cmu-sphinx/cmu-sphinxbase
	brew install --HEAD watsonbox/cmu-sphinx/cmu-sphinxtrain
	brew install --HEAD watsonbox/cmu-sphinx/cmu-pocketsphinx

fclean:
	rm $(NAME)
