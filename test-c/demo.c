#include <stdio.h>
#include <stdlib.h>
#include <stdint.h>
#include <string.h>

#include "colors.h"
#include "color_funcs.h"

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
	/*
	printColor(colors[10]);
	printColor(colors[20]);
	printColor(colors[30]);
	*/
	/*
	printf("%d\n", chanDiffSq(10, 20));
	printf("%d\n", chanDiffSq(20, 10));
	printf("%d\n", chanDiffSq(255, 0));
	printf("%d\n", chanDiffSq(0, 255));
	*/
	/*
	uint8_t c1[4] = {0, 100, 0, 0};
	uint8_t c2[4] = {0, 0, 100, 0};
	printf("%f\n", distanceRGB(c1, c2));
	*/
	uint8_t c[4];
	for (uint8_t code = 0; code < 100; code += 13) {
		codeToRGB(code, c);
		printf("%d -> rgb(%d, %d, %d)\n", code, c[0], c[1], c[2]);
	}
}
