#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

#include "colors.h"

char *colorString(ColorProp *c) {
	char *str = malloc(70);
	sprintf(
		str,
		"code=%d, hex=%s, rgba={%d, %d, %d}, hsl={%.3f, %.3f, %.3f}",
		c->code,
		c->hex,
		c->rgba[0], c->rgba[1], c->rgba[2], c->rgba[3],
		c->hsl[0], c->hsl[1], c->hsl[2]
	);
	// printf("colorString: size=%d\n", strlen(str));
	return str;
}

void printColor(ColorProp c) {
	printf("Color: %s\n", colorString(&c));
	printf("    Names: %s\n\n", c.names);
}

int main(int argc, char **argv) {
	printColor(colors[10]);
	printColor(colors[20]);
	printColor(colors[30]);
}
