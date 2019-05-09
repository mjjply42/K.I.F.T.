/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   sphinx.c                                           :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: dromansk <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/03/25 21:45:14 by dromansk          #+#    #+#             */
/*   Updated: 2019/05/08 19:38:08 by dromansk         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <pocketsphinx.h>
#include "sphinx.h"

/*
** Main included so we can compile it and use it as an outside program.
** However I have structured the code so you can just call psphinx_string
** to get your desired output string from the function directly.
*/

char const	*parse_input(ps_decoder_t *reading, FILE *file, int *score)
{
	int			stream;
	size_t		sample;
	short		buf[512];

	stream = ps_start_utt(reading);
	while (!feof(file))
	{
		sample = fread(buf, 2, 512, file);
		stream = ps_process_raw(reading, buf, sample, FALSE, FALSE);
	}
	ps_end_utt(reading);
	return (ps_get_hyp(reading, score));
}

char const	*pocketsphinx_string(char *path, cmd_ln_t *config)
{
	ps_decoder_t	*parsing;
	FILE			*file;
	char const		*utt;
	int				score;

	parsing = ps_init(config);
	if (!(file = fopen(path, "rb")))
		return (NULL);
	utt = ft_strdup(parse_input(parsing, file, &score));
	fclose(file);
	ps_free(parsing);
	return (utt);
}

cmd_ln_t	*psphinx_config(void)
{
	cmd_ln_t		*config;

	config = cmd_ln_init(NULL, ps_args(), TRUE,
				"-hmm", "./en-us/en-us-adapt",
				"-lm", "./en-us/en-us.lm.bin",
				"-dict", "./en-us/cmudict-en-us.dict",
				"-logfn", "/dev/null", NULL);
	return (config);
}

char const	*psphinx_string(char *path)
{
	char const		*utt;
	cmd_ln_t		*config;

	config = psphinx_config();
	utt = pocketsphinx_string(path, config);
	cmd_ln_free_r(config);
	return (utt);
}

int			main(int ac, char **av)
{
	char	*utt;

	utt = NULL;
	if (ac == 2)
	{
		utt = (char *)psphinx_string(av[1]);
		printf("%s", utt);
		free(utt);
	}
	return (0);
}
