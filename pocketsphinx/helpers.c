/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   helpers.c                                          :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: dromansk <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2019/03/25 21:45:14 by dromansk          #+#    #+#             */
/*   Updated: 2019/03/31 21:26:58 by dromansk         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "sphinx.h"
#include "stdio.h"

size_t	ft_strlen(char const *str)
{
	size_t	i;

	i = 0;
	while (str[i])
		i++;
	return (i);
}

char	*ft_strdup(char const *str)
{
	size_t	len;
	char	*new;

	len = ft_strlen(str);
	if (!(new = (char *)malloc(sizeof(char) * ++len)))
		return (NULL);
	while (len--)
		new[len] = str[len];
	return (new);
}
