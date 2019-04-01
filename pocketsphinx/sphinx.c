/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   test.c                                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: dromansk <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/03/25 21:45:14 by dromansk          #+#    #+#             */
/*   Updated: 2019/03/31 18:59:02 by dromansk         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <pocketsphinx.h>

/*
** idk what to return yet, using strings for test purposes
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
	return(ps_get_hyp(reading, score));
}

char const	*pocketsphinx_string(char *path, cmd_ln_t *config)
{
	ps_decoder_t	*parsing;
	FILE			*file;
	char const		*utt;
	int				score;

	parsing = ps_init(config);
	file = fopen(path, "rb");
	utt = parse_input(parsing, file, &score);
	fclose(file);
	ps_free(parsing);
	return (utt);
}

cmd_ln_t	*psphinx_config(void)
{
	cmd_ln_t		*config;

	config = cmd_ln_init(NULL, ps_args(), TRUE,
		        "-hmm", MODELDIR "/en-us/en-us",
		        "-lm", MODELDIR "/en-us/en-us.lm.bin",
	    		"-dict", MODELDIR "/en-us/cmudict-en-us.dict",
				"-logfn", "/dev/null", NULL);
	return (config);
}

int			main(int ac, char **av)
{
	if (ac == 2)
	{
		cmd_ln_t *config = psphinx_config();
		char *s = (char *)pocketsphinx_string(av[1], config);
		printf("%s\n", s);
		cmd_ln_free_r(config);
	}
	return (0);
}
