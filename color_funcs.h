#include <math.h>

#define DEFAULT_GRAY_MAX_DELTA 20

uint8_t values[6] = {
	0,
	95,  // prev + 95
	135, // prev + 40
	175, // prev + 40
	215, // prev + 40
	255, // prev + 40
};

// these colors are compatible with Xfce terminal and LXDE terminal
// but Tango palette (in Gnome Terminal) has brighter colors

uint8_t first16Colors[16][3] = {
	{0, 0, 0},       // 0 -> 000000, tango:2e3436, group 1
	{170, 0, 0},     // 1 -> aa0000, tango:cc0000, group 2
	{0, 170, 0},     // 2 -> 00aa00, tango:4e9a06, group 2
	{170, 85, 0},    // 3 -> aa5500, tango:c4a000, group 3
	{0, 0, 170},     // 4 -> 0000aa, tango:3465a4, group 2
	{170, 0, 170},   // 5 -> aa00aa, tango:75507b, group 4
	{0, 170, 170},   // 6 -> 00aaaa, tango:06989a, group 4
	{185, 185, 185}, // 7 -> b9b9b9, tango:d3d7cf, group 1
	{85, 85, 85},    // 8 -> 555555, tango:555753, group 1
	{255, 85, 85},   // 9 -> ff5555, tango:ef2929, group 5
	{85, 255, 85},   // 10 -> 55ff55, tango:8ae234, group 5
	{255, 255, 85},  // 11 -> ffff55, tango:fce94f, group 6
	{85, 85, 255},   // 12 -> 5555ff, tango:729fcf, group 5
	{255, 85, 255},  // 13 -> ff55ff, tango:ad7fa8, group 6
	{85, 255, 255},  // 14 -> 55ffff, tango:34e2e2, group 6
	{255, 255, 255}, // 15 -> ffffff, tango:eeeeec, group 1
};

enum ROUND_MODE {ROUND_CLOSER = 0, ROUND_DOWN = 1, ROUND_UP = 2};


uint8_t divideRoundUp(uint8_t a, uint8_t b) {
	return (a + b - 1) / b
}

uint8_t divideRoundCloser(uint8_t a, uint8_t b) {
	if a%b > b/2 {
		return (a + b - 1) / b
	}
	return a / b
}

// returns round(a/b) based on given RoundMode
uint8_t divideRound(uint8_t a, uint8_t b, ROUND_MODE mode) {
	switch mode {
	case RoundDown:
		return a / b
	case RoundUp:
		return divideRoundUp(a, b)
	}
	return divideRoundCloser(a, b)
}



uint32_t chanDiffSq(uint8_t x1, uint8_t x2) {
	int32_t d = (int32_t)x1 - (int32_t)x2;
	return (uint32_t)(d * d);
}

double distanceRGB(uint8_t c1[4], uint8_t c2[4]) {
	uint32_t d = chanDiffSq(c1[0], c2[0]) +
		chanDiffSq(c1[1], c2[1]) +
		chanDiffSq(c1[2], c2[2]);
	return sqrt((double)d);
}

void codeToRGB(uint8_t n, uint8_t c[3]) {
	if (n < 16) {
		uint8_t c2[3];
		c[0] = first16Colors[n][0];
		c[1] = first16Colors[n][1];
		c[2] = first16Colors[n][2];
		return;
	}
	if (n >= 232) {
		uint8_t v = 8 + 10*(n-232);
		c[0] = v; c[1] = v; c[2] = v;
		return;
	}
	uint8_t m = n - 16;

	c[0] = values[m/36];
	c[1] = values[(m%36)/6];
	c[2] = values[m%6];
}

