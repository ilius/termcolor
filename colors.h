#include <stdio.h>
#include <stdint.h>


typedef struct ColorProp {
	uint8_t code;
	uint8_t rgba[4];
	double hsl[3];
	char *hex;
	char *names; // separated by ","
	uint8_t nameCount; 
} ColorProp;


ColorProp colors[] = {
	{
		.code = 0,
		.rgba = {0, 0, 0, 255},
		.hsl = {0, 0, 0},
		.hex = "#000000",
		.names = "black,css:black",
		.nameCount = 2,
	},
	{
		.code = 1,
		.rgba = {170, 0, 0, 255},
		.hsl = {0, 1, 0.3333333333333333},
		.hex = "#aa0000",
		.names = "dark red 1",
		.nameCount = 1,
	},
	{
		.code = 2,
		.rgba = {0, 170, 0, 255},
		.hsl = {120, 1, 0.3333333333333333},
		.hex = "#00aa00",
		.names = "mixed green 1",
		.nameCount = 1,
	},
	{
		.code = 3,
		.rgba = {170, 85, 0, 255},
		.hsl = {30, 1, 0.3333333333333333},
		.hex = "#aa5500",
		.names = "mixed orange 1",
		.nameCount = 1,
	},
	{
		.code = 4,
		.rgba = {0, 0, 170, 255},
		.hsl = {240, 1, 0.3333333333333333},
		.hex = "#0000aa",
		.names = "mixed blue 1",
		.nameCount = 1,
	},
	{
		.code = 5,
		.rgba = {170, 0, 170, 255},
		.hsl = {300, 1, 0.3333333333333333},
		.hex = "#aa00aa",
		.names = "purple 1",
		.nameCount = 1,
	},
	{
		.code = 6,
		.rgba = {0, 170, 170, 255},
		.hsl = {180, 1, 0.3333333333333333},
		.hex = "#00aaaa",
		.names = "cyan 1",
		.nameCount = 1,
	},
	{
		.code = 7,
		.rgba = {185, 185, 185, 255},
		.hsl = {0, 0, 0.7254901960784313},
		.hex = "#b9b9b9",
		.names = "light gray",
		.nameCount = 1,
	},
	{
		.code = 8,
		.rgba = {85, 85, 85, 255},
		.hsl = {0, 0, 0.3333333333333333},
		.hex = "#555555",
		.names = "dark gray",
		.nameCount = 1,
	},
	{
		.code = 9,
		.rgba = {255, 85, 85, 255},
		.hsl = {0, 1, 0.6666666666666666},
		.hex = "#ff5555",
		.names = "light red",
		.nameCount = 1,
	},
	{
		.code = 10,
		.rgba = {85, 255, 85, 255},
		.hsl = {120, 1, 0.6666666666666666},
		.hex = "#55ff55",
		.names = "light green",
		.nameCount = 1,
	},
	{
		.code = 11,
		.rgba = {255, 255, 85, 255},
		.hsl = {60, 1, 0.6666666666666666},
		.hex = "#ffff55",
		.names = "yellow",
		.nameCount = 1,
	},
	{
		.code = 12,
		.rgba = {85, 85, 255, 255},
		.hsl = {240, 1, 0.6666666666666666},
		.hex = "#5555ff",
		.names = "light blue",
		.nameCount = 1,
	},
	{
		.code = 13,
		.rgba = {255, 85, 255, 255},
		.hsl = {300, 1, 0.6666666666666666},
		.hex = "#ff55ff",
		.names = "light purple",
		.nameCount = 1,
	},
	{
		.code = 14,
		.rgba = {85, 255, 255, 255},
		.hsl = {180, 1, 0.6666666666666666},
		.hex = "#55ffff",
		.names = "light cyan",
		.nameCount = 1,
	},
	{
		.code = 15,
		.rgba = {255, 255, 255, 255},
		.hsl = {0, 0, 1},
		.hex = "#ffffff",
		.names = "white,css:white",
		.nameCount = 2,
	},
	{
		.code = 16,
		.rgba = {0, 0, 0, 255},
		.hsl = {0, 0, 0},
		.hex = "#000000",
		.names = "black,css:black",
		.nameCount = 2,
	},
	{
		.code = 17,
		.rgba = {0, 0, 95, 255},
		.hsl = {240, 1, 0.18627450980392157},
		.hex = "#00005f",
		.names = "blue 4",
		.nameCount = 1,
	},
	{
		.code = 18,
		.rgba = {0, 0, 135, 255},
		.hsl = {240, 1, 0.2647058823529412},
		.hex = "#000087",
		.names = "blue 3,css:darkblue",
		.nameCount = 2,
	},
	{
		.code = 19,
		.rgba = {0, 0, 175, 255},
		.hsl = {240, 1, 0.3431372549019608},
		.hex = "#0000af",
		.names = "blue 2",
		.nameCount = 1,
	},
	{
		.code = 20,
		.rgba = {0, 0, 215, 255},
		.hsl = {240, 1, 0.4215686274509804},
		.hex = "#0000d7",
		.names = "blue 1,css:mediumblue",
		.nameCount = 2,
	},
	{
		.code = 21,
		.rgba = {0, 0, 255, 255},
		.hsl = {240, 1, 0.5},
		.hex = "#0000ff",
		.names = "blue,css:blue",
		.nameCount = 2,
	},
	{
		.code = 22,
		.rgba = {0, 95, 0, 255},
		.hsl = {120, 1, 0.18627450980392157},
		.hex = "#005f00",
		.names = "green 4,css:darkgreen",
		.nameCount = 2,
	},
	{
		.code = 23,
		.rgba = {0, 95, 95, 255},
		.hsl = {180, 1, 0.18627450980392157},
		.hex = "#005f5f",
		.names = "blue stone",
		.nameCount = 1,
	},
	{
		.code = 24,
		.rgba = {0, 95, 135, 255},
		.hsl = {197.77777777777777, 1, 0.2647058823529412},
		.hex = "#005f87",
		.names = "orient",
		.nameCount = 1,
	},
	{
		.code = 25,
		.rgba = {0, 95, 175, 255},
		.hsl = {207.42857142857144, 1, 0.3431372549019608},
		.hex = "#005faf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 26,
		.rgba = {0, 95, 215, 255},
		.hsl = {213.48837209302326, 1, 0.4215686274509804},
		.hex = "#005fd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 27,
		.rgba = {0, 95, 255, 255},
		.hsl = {217.64705882352942, 1, 0.5},
		.hex = "#005fff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 28,
		.rgba = {0, 135, 0, 255},
		.hsl = {120, 1, 0.2647058823529412},
		.hex = "#008700",
		.names = "green 3,css:green",
		.nameCount = 2,
	},
	{
		.code = 29,
		.rgba = {0, 135, 95, 255},
		.hsl = {162.22222222222223, 1, 0.2647058823529412},
		.hex = "#00875f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 30,
		.rgba = {0, 135, 135, 255},
		.hsl = {180, 1, 0.2647058823529412},
		.hex = "#008787",
		.names = "css:darkcyan",
		.nameCount = 1,
	},
	{
		.code = 31,
		.rgba = {0, 135, 175, 255},
		.hsl = {193.71428571428572, 1, 0.3431372549019608},
		.hex = "#0087af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 32,
		.rgba = {0, 135, 215, 255},
		.hsl = {202.32558139534882, 1, 0.4215686274509804},
		.hex = "#0087d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 33,
		.rgba = {0, 135, 255, 255},
		.hsl = {208.23529411764707, 1, 0.5},
		.hex = "#0087ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 34,
		.rgba = {0, 175, 0, 255},
		.hsl = {120, 1, 0.3431372549019608},
		.hex = "#00af00",
		.names = "green 2",
		.nameCount = 1,
	},
	{
		.code = 35,
		.rgba = {0, 175, 95, 255},
		.hsl = {152.57142857142856, 1, 0.3431372549019608},
		.hex = "#00af5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 36,
		.rgba = {0, 175, 135, 255},
		.hsl = {166.28571428571428, 1, 0.3431372549019608},
		.hex = "#00af87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 37,
		.rgba = {0, 175, 175, 255},
		.hsl = {180, 1, 0.3431372549019608},
		.hex = "#00afaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 38,
		.rgba = {0, 175, 215, 255},
		.hsl = {191.1627906976744, 1, 0.4215686274509804},
		.hex = "#00afd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 39,
		.rgba = {0, 175, 255, 255},
		.hsl = {198.8235294117647, 1, 0.5},
		.hex = "#00afff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 40,
		.rgba = {0, 215, 0, 255},
		.hsl = {120, 1, 0.4215686274509804},
		.hex = "#00d700",
		.names = "green 1",
		.nameCount = 1,
	},
	{
		.code = 41,
		.rgba = {0, 215, 95, 255},
		.hsl = {146.51162790697674, 1, 0.4215686274509804},
		.hex = "#00d75f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 42,
		.rgba = {0, 215, 135, 255},
		.hsl = {157.67441860465118, 1, 0.4215686274509804},
		.hex = "#00d787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 43,
		.rgba = {0, 215, 175, 255},
		.hsl = {168.8372093023256, 1, 0.4215686274509804},
		.hex = "#00d7af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 44,
		.rgba = {0, 215, 215, 255},
		.hsl = {180, 1, 0.4215686274509804},
		.hex = "#00d7d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 45,
		.rgba = {0, 215, 255, 255},
		.hsl = {189.41176470588235, 1, 0.5},
		.hex = "#00d7ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 46,
		.rgba = {0, 255, 0, 255},
		.hsl = {120, 1, 0.5},
		.hex = "#00ff00",
		.names = "green,css:lime",
		.nameCount = 2,
	},
	{
		.code = 47,
		.rgba = {0, 255, 95, 255},
		.hsl = {142.35294117647058, 1, 0.5},
		.hex = "#00ff5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 48,
		.rgba = {0, 255, 135, 255},
		.hsl = {151.76470588235293, 1, 0.5},
		.hex = "#00ff87",
		.names = "css:springgreen",
		.nameCount = 1,
	},
	{
		.code = 49,
		.rgba = {0, 255, 175, 255},
		.hsl = {161.1764705882353, 1, 0.5},
		.hex = "#00ffaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 50,
		.rgba = {0, 255, 215, 255},
		.hsl = {170.58823529411765, 1, 0.5},
		.hex = "#00ffd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 51,
		.rgba = {0, 255, 255, 255},
		.hsl = {180, 1, 0.5},
		.hex = "#00ffff",
		.names = "cyan,css:cyan,css:aqua",
		.nameCount = 3,
	},
	{
		.code = 52,
		.rgba = {95, 0, 0, 255},
		.hsl = {0, 1, 0.18627450980392157},
		.hex = "#5f0000",
		.names = "red 5",
		.nameCount = 1,
	},
	{
		.code = 53,
		.rgba = {95, 0, 95, 255},
		.hsl = {300, 1, 0.18627450980392157},
		.hex = "#5f005f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 54,
		.rgba = {95, 0, 135, 255},
		.hsl = {282.22222222222223, 1, 0.2647058823529412},
		.hex = "#5f0087",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 55,
		.rgba = {95, 0, 175, 255},
		.hsl = {272.57142857142856, 1, 0.3431372549019608},
		.hex = "#5f00af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 56,
		.rgba = {95, 0, 215, 255},
		.hsl = {266.51162790697674, 1, 0.4215686274509804},
		.hex = "#5f00d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 57,
		.rgba = {95, 0, 255, 255},
		.hsl = {262.3529411764706, 1, 0.5},
		.hex = "#5f00ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 58,
		.rgba = {95, 95, 0, 255},
		.hsl = {60, 1, 0.18627450980392157},
		.hex = "#5f5f00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 59,
		.rgba = {95, 95, 95, 255},
		.hsl = {0, 0, 0.37254901960784315},
		.hex = "#5f5f5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 60,
		.rgba = {95, 95, 135, 255},
		.hsl = {240, 0.17391304347826086, 0.45098039215686275},
		.hex = "#5f5f87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 61,
		.rgba = {95, 95, 175, 255},
		.hsl = {240, 0.3333333333333333, 0.5294117647058824},
		.hex = "#5f5faf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 62,
		.rgba = {95, 95, 215, 255},
		.hsl = {240, 0.6000000000000001, 0.607843137254902},
		.hex = "#5f5fd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 63,
		.rgba = {95, 95, 255, 255},
		.hsl = {240, 1, 0.6862745098039216},
		.hex = "#5f5fff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 64,
		.rgba = {95, 135, 0, 255},
		.hsl = {77.77777777777777, 1, 0.2647058823529412},
		.hex = "#5f8700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 65,
		.rgba = {95, 135, 95, 255},
		.hsl = {120, 0.17391304347826086, 0.45098039215686275},
		.hex = "#5f875f",
		.names = "glade green",
		.nameCount = 1,
	},
	{
		.code = 66,
		.rgba = {95, 135, 135, 255},
		.hsl = {180, 0.17391304347826086, 0.45098039215686275},
		.hex = "#5f8787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 67,
		.rgba = {95, 135, 175, 255},
		.hsl = {210, 0.3333333333333333, 0.5294117647058824},
		.hex = "#5f87af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 68,
		.rgba = {95, 135, 215, 255},
		.hsl = {220, 0.6000000000000001, 0.607843137254902},
		.hex = "#5f87d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 69,
		.rgba = {95, 135, 255, 255},
		.hsl = {225, 1, 0.6862745098039216},
		.hex = "#5f87ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 70,
		.rgba = {95, 175, 0, 255},
		.hsl = {87.42857142857144, 1, 0.3431372549019608},
		.hex = "#5faf00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 71,
		.rgba = {95, 175, 95, 255},
		.hsl = {120, 0.3333333333333333, 0.5294117647058824},
		.hex = "#5faf5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 72,
		.rgba = {95, 175, 135, 255},
		.hsl = {150, 0.3333333333333333, 0.5294117647058824},
		.hex = "#5faf87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 73,
		.rgba = {95, 175, 175, 255},
		.hsl = {180, 0.3333333333333333, 0.5294117647058824},
		.hex = "#5fafaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 74,
		.rgba = {95, 175, 215, 255},
		.hsl = {200, 0.6000000000000001, 0.607843137254902},
		.hex = "#5fafd7",
		.names = "tradewind",
		.nameCount = 1,
	},
	{
		.code = 75,
		.rgba = {95, 175, 255, 255},
		.hsl = {210, 1, 0.6862745098039216},
		.hex = "#5fafff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 76,
		.rgba = {95, 215, 0, 255},
		.hsl = {93.48837209302326, 1, 0.4215686274509804},
		.hex = "#5fd700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 77,
		.rgba = {95, 215, 95, 255},
		.hsl = {120, 0.6000000000000001, 0.607843137254902},
		.hex = "#5fd75f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 78,
		.rgba = {95, 215, 135, 255},
		.hsl = {140, 0.6000000000000001, 0.607843137254902},
		.hex = "#5fd787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 79,
		.rgba = {95, 215, 175, 255},
		.hsl = {160, 0.6000000000000001, 0.607843137254902},
		.hex = "#5fd7af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 80,
		.rgba = {95, 215, 215, 255},
		.hsl = {180, 0.6000000000000001, 0.607843137254902},
		.hex = "#5fd7d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 81,
		.rgba = {95, 215, 255, 255},
		.hsl = {195, 1, 0.6862745098039216},
		.hex = "#5fd7ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 82,
		.rgba = {95, 255, 0, 255},
		.hsl = {97.6470588235294, 1, 0.5},
		.hex = "#5fff00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 83,
		.rgba = {95, 255, 95, 255},
		.hsl = {120, 1, 0.6862745098039216},
		.hex = "#5fff5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 84,
		.rgba = {95, 255, 135, 255},
		.hsl = {135, 1, 0.6862745098039216},
		.hex = "#5fff87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 85,
		.rgba = {95, 255, 175, 255},
		.hsl = {150, 1, 0.6862745098039216},
		.hex = "#5fffaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 86,
		.rgba = {95, 255, 215, 255},
		.hsl = {165, 1, 0.6862745098039216},
		.hex = "#5fffd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 87,
		.rgba = {95, 255, 255, 255},
		.hsl = {180, 1, 0.6862745098039216},
		.hex = "#5fffff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 88,
		.rgba = {135, 0, 0, 255},
		.hsl = {0, 1, 0.2647058823529412},
		.hex = "#870000",
		.names = "red 4,css:darkred",
		.nameCount = 2,
	},
	{
		.code = 89,
		.rgba = {135, 0, 95, 255},
		.hsl = {317.77777777777777, 1, 0.2647058823529412},
		.hex = "#87005f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 90,
		.rgba = {135, 0, 135, 255},
		.hsl = {300, 1, 0.2647058823529412},
		.hex = "#870087",
		.names = "css:darkmagenta",
		.nameCount = 1,
	},
	{
		.code = 91,
		.rgba = {135, 0, 175, 255},
		.hsl = {286.2857142857143, 1, 0.3431372549019608},
		.hex = "#8700af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 92,
		.rgba = {135, 0, 215, 255},
		.hsl = {277.6744186046512, 1, 0.4215686274509804},
		.hex = "#8700d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 93,
		.rgba = {135, 0, 255, 255},
		.hsl = {271.7647058823529, 1, 0.5},
		.hex = "#8700ff",
		.names = "electric violet",
		.nameCount = 1,
	},
	{
		.code = 94,
		.rgba = {135, 95, 0, 255},
		.hsl = {42.22222222222222, 1, 0.2647058823529412},
		.hex = "#875f00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 95,
		.rgba = {135, 95, 95, 255},
		.hsl = {0, 0.17391304347826086, 0.45098039215686275},
		.hex = "#875f5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 96,
		.rgba = {135, 95, 135, 255},
		.hsl = {300, 0.17391304347826086, 0.45098039215686275},
		.hex = "#875f87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 97,
		.rgba = {135, 95, 175, 255},
		.hsl = {270, 0.3333333333333333, 0.5294117647058824},
		.hex = "#875faf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 98,
		.rgba = {135, 95, 215, 255},
		.hsl = {260, 0.6000000000000001, 0.607843137254902},
		.hex = "#875fd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 99,
		.rgba = {135, 95, 255, 255},
		.hsl = {255, 1, 0.6862745098039216},
		.hex = "#875fff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 100,
		.rgba = {135, 135, 0, 255},
		.hsl = {60, 1, 0.2647058823529412},
		.hex = "#878700",
		.names = "css:olive",
		.nameCount = 1,
	},
	{
		.code = 101,
		.rgba = {135, 135, 95, 255},
		.hsl = {60, 0.17391304347826086, 0.45098039215686275},
		.hex = "#87875f",
		.names = "clay creek",
		.nameCount = 1,
	},
	{
		.code = 102,
		.rgba = {135, 135, 135, 255},
		.hsl = {0, 0, 0.5294117647058824},
		.hex = "#878787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 103,
		.rgba = {135, 135, 175, 255},
		.hsl = {240, 0.2, 0.607843137254902},
		.hex = "#8787af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 104,
		.rgba = {135, 135, 215, 255},
		.hsl = {240, 0.5000000000000001, 0.6862745098039216},
		.hex = "#8787d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 105,
		.rgba = {135, 135, 255, 255},
		.hsl = {240, 1, 0.7647058823529411},
		.hex = "#8787ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 106,
		.rgba = {135, 175, 0, 255},
		.hsl = {73.71428571428572, 1, 0.3431372549019608},
		.hex = "#87af00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 107,
		.rgba = {135, 175, 95, 255},
		.hsl = {90, 0.3333333333333333, 0.5294117647058824},
		.hex = "#87af5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 108,
		.rgba = {135, 175, 135, 255},
		.hsl = {120, 0.2, 0.607843137254902},
		.hex = "#87af87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 109,
		.rgba = {135, 175, 175, 255},
		.hsl = {180, 0.2, 0.607843137254902},
		.hex = "#87afaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 110,
		.rgba = {135, 175, 215, 255},
		.hsl = {210, 0.5000000000000001, 0.6862745098039216},
		.hex = "#87afd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 111,
		.rgba = {135, 175, 255, 255},
		.hsl = {220, 1, 0.7647058823529411},
		.hex = "#87afff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 112,
		.rgba = {135, 215, 0, 255},
		.hsl = {82.32558139534883, 1, 0.4215686274509804},
		.hex = "#87d700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 113,
		.rgba = {135, 215, 95, 255},
		.hsl = {100, 0.6000000000000001, 0.607843137254902},
		.hex = "#87d75f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 114,
		.rgba = {135, 215, 135, 255},
		.hsl = {120, 0.5000000000000001, 0.6862745098039216},
		.hex = "#87d787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 115,
		.rgba = {135, 215, 175, 255},
		.hsl = {150, 0.5000000000000001, 0.6862745098039216},
		.hex = "#87d7af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 116,
		.rgba = {135, 215, 215, 255},
		.hsl = {180, 0.5000000000000001, 0.6862745098039216},
		.hex = "#87d7d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 117,
		.rgba = {135, 215, 255, 255},
		.hsl = {200, 1, 0.7647058823529411},
		.hex = "#87d7ff",
		.names = "css:lightskyblue",
		.nameCount = 1,
	},
	{
		.code = 118,
		.rgba = {135, 255, 0, 255},
		.hsl = {88.23529411764707, 1, 0.5},
		.hex = "#87ff00",
		.names = "css:chartreuse",
		.nameCount = 1,
	},
	{
		.code = 119,
		.rgba = {135, 255, 95, 255},
		.hsl = {105, 1, 0.6862745098039216},
		.hex = "#87ff5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 120,
		.rgba = {135, 255, 135, 255},
		.hsl = {120, 1, 0.7647058823529411},
		.hex = "#87ff87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 121,
		.rgba = {135, 255, 175, 255},
		.hsl = {140, 1, 0.7647058823529411},
		.hex = "#87ffaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 122,
		.rgba = {135, 255, 215, 255},
		.hsl = {160, 1, 0.7647058823529411},
		.hex = "#87ffd7",
		.names = "css:aquamarine",
		.nameCount = 1,
	},
	{
		.code = 123,
		.rgba = {135, 255, 255, 255},
		.hsl = {180, 1, 0.7647058823529411},
		.hex = "#87ffff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 124,
		.rgba = {175, 0, 0, 255},
		.hsl = {0, 1, 0.3431372549019608},
		.hex = "#af0000",
		.names = "red 3,bright red",
		.nameCount = 2,
	},
	{
		.code = 125,
		.rgba = {175, 0, 95, 255},
		.hsl = {327.42857142857144, 1, 0.3431372549019608},
		.hex = "#af005f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 126,
		.rgba = {175, 0, 135, 255},
		.hsl = {313.7142857142857, 1, 0.3431372549019608},
		.hex = "#af0087",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 127,
		.rgba = {175, 0, 175, 255},
		.hsl = {300, 1, 0.3431372549019608},
		.hex = "#af00af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 128,
		.rgba = {175, 0, 215, 255},
		.hsl = {288.83720930232556, 1, 0.4215686274509804},
		.hex = "#af00d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 129,
		.rgba = {175, 0, 255, 255},
		.hsl = {281.1764705882353, 1, 0.5},
		.hex = "#af00ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 130,
		.rgba = {175, 95, 0, 255},
		.hsl = {32.57142857142857, 1, 0.3431372549019608},
		.hex = "#af5f00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 131,
		.rgba = {175, 95, 95, 255},
		.hsl = {0, 0.3333333333333333, 0.5294117647058824},
		.hex = "#af5f5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 132,
		.rgba = {175, 95, 135, 255},
		.hsl = {330, 0.3333333333333333, 0.5294117647058824},
		.hex = "#af5f87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 133,
		.rgba = {175, 95, 175, 255},
		.hsl = {300, 0.3333333333333333, 0.5294117647058824},
		.hex = "#af5faf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 134,
		.rgba = {175, 95, 215, 255},
		.hsl = {280, 0.6000000000000001, 0.607843137254902},
		.hex = "#af5fd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 135,
		.rgba = {175, 95, 255, 255},
		.hsl = {270, 1, 0.6862745098039216},
		.hex = "#af5fff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 136,
		.rgba = {175, 135, 0, 255},
		.hsl = {46.285714285714285, 1, 0.3431372549019608},
		.hex = "#af8700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 137,
		.rgba = {175, 135, 95, 255},
		.hsl = {30, 0.3333333333333333, 0.5294117647058824},
		.hex = "#af875f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 138,
		.rgba = {175, 135, 135, 255},
		.hsl = {0, 0.2, 0.607843137254902},
		.hex = "#af8787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 139,
		.rgba = {175, 135, 175, 255},
		.hsl = {300, 0.2, 0.607843137254902},
		.hex = "#af87af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 140,
		.rgba = {175, 135, 215, 255},
		.hsl = {270, 0.5000000000000001, 0.6862745098039216},
		.hex = "#af87d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 141,
		.rgba = {175, 135, 255, 255},
		.hsl = {260, 1, 0.7647058823529411},
		.hex = "#af87ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 142,
		.rgba = {175, 175, 0, 255},
		.hsl = {60, 1, 0.3431372549019608},
		.hex = "#afaf00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 143,
		.rgba = {175, 175, 95, 255},
		.hsl = {60, 0.3333333333333333, 0.5294117647058824},
		.hex = "#afaf5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 144,
		.rgba = {175, 175, 135, 255},
		.hsl = {60, 0.2, 0.607843137254902},
		.hex = "#afaf87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 145,
		.rgba = {175, 175, 175, 255},
		.hsl = {0, 0, 0.6862745098039216},
		.hex = "#afafaf",
		.names = "silver chalice",
		.nameCount = 1,
	},
	{
		.code = 146,
		.rgba = {175, 175, 215, 255},
		.hsl = {240, 0.3333333333333334, 0.7647058823529411},
		.hex = "#afafd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 147,
		.rgba = {175, 175, 255, 255},
		.hsl = {240, 1, 0.8431372549019608},
		.hex = "#afafff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 148,
		.rgba = {175, 215, 0, 255},
		.hsl = {71.16279069767441, 1, 0.4215686274509804},
		.hex = "#afd700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 149,
		.rgba = {175, 215, 95, 255},
		.hsl = {80.00000000000001, 0.6000000000000001, 0.607843137254902},
		.hex = "#afd75f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 150,
		.rgba = {175, 215, 135, 255},
		.hsl = {90, 0.5000000000000001, 0.6862745098039216},
		.hex = "#afd787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 151,
		.rgba = {175, 215, 175, 255},
		.hsl = {120, 0.3333333333333334, 0.7647058823529411},
		.hex = "#afd7af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 152,
		.rgba = {175, 215, 215, 255},
		.hsl = {180, 0.3333333333333334, 0.7647058823529411},
		.hex = "#afd7d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 153,
		.rgba = {175, 215, 255, 255},
		.hsl = {210, 1, 0.8431372549019608},
		.hex = "#afd7ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 154,
		.rgba = {175, 255, 0, 255},
		.hsl = {78.82352941176471, 1, 0.5},
		.hex = "#afff00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 155,
		.rgba = {175, 255, 95, 255},
		.hsl = {90, 1, 0.6862745098039216},
		.hex = "#afff5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 156,
		.rgba = {175, 255, 135, 255},
		.hsl = {100, 1, 0.7647058823529411},
		.hex = "#afff87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 157,
		.rgba = {175, 255, 175, 255},
		.hsl = {120, 1, 0.8431372549019608},
		.hex = "#afffaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 158,
		.rgba = {175, 255, 215, 255},
		.hsl = {150, 1, 0.8431372549019608},
		.hex = "#afffd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 159,
		.rgba = {175, 255, 255, 255},
		.hsl = {180, 1, 0.8431372549019608},
		.hex = "#afffff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 160,
		.rgba = {215, 0, 0, 255},
		.hsl = {0, 1, 0.4215686274509804},
		.hex = "#d70000",
		.names = "red 2",
		.nameCount = 1,
	},
	{
		.code = 161,
		.rgba = {215, 0, 95, 255},
		.hsl = {333.48837209302326, 1, 0.4215686274509804},
		.hex = "#d7005f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 162,
		.rgba = {215, 0, 135, 255},
		.hsl = {322.3255813953488, 1, 0.4215686274509804},
		.hex = "#d70087",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 163,
		.rgba = {215, 0, 175, 255},
		.hsl = {311.16279069767444, 1, 0.4215686274509804},
		.hex = "#d700af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 164,
		.rgba = {215, 0, 215, 255},
		.hsl = {300, 1, 0.4215686274509804},
		.hex = "#d700d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 165,
		.rgba = {215, 0, 255, 255},
		.hsl = {290.5882352941176, 1, 0.5},
		.hex = "#d700ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 166,
		.rgba = {215, 95, 0, 255},
		.hsl = {26.511627906976745, 1, 0.4215686274509804},
		.hex = "#d75f00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 167,
		.rgba = {215, 95, 95, 255},
		.hsl = {0, 0.6000000000000001, 0.607843137254902},
		.hex = "#d75f5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 168,
		.rgba = {215, 95, 135, 255},
		.hsl = {340, 0.6000000000000001, 0.607843137254902},
		.hex = "#d75f87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 169,
		.rgba = {215, 95, 175, 255},
		.hsl = {320, 0.6000000000000001, 0.607843137254902},
		.hex = "#d75faf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 170,
		.rgba = {215, 95, 215, 255},
		.hsl = {300, 0.6000000000000001, 0.607843137254902},
		.hex = "#d75fd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 171,
		.rgba = {215, 95, 255, 255},
		.hsl = {285, 1, 0.6862745098039216},
		.hex = "#d75fff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 172,
		.rgba = {215, 135, 0, 255},
		.hsl = {37.674418604651166, 1, 0.4215686274509804},
		.hex = "#d78700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 173,
		.rgba = {215, 135, 95, 255},
		.hsl = {20, 0.6000000000000001, 0.607843137254902},
		.hex = "#d7875f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 174,
		.rgba = {215, 135, 135, 255},
		.hsl = {0, 0.5000000000000001, 0.6862745098039216},
		.hex = "#d78787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 175,
		.rgba = {215, 135, 175, 255},
		.hsl = {330, 0.5000000000000001, 0.6862745098039216},
		.hex = "#d787af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 176,
		.rgba = {215, 135, 215, 255},
		.hsl = {300, 0.5000000000000001, 0.6862745098039216},
		.hex = "#d787d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 177,
		.rgba = {215, 135, 255, 255},
		.hsl = {280, 1, 0.7647058823529411},
		.hex = "#d787ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 178,
		.rgba = {215, 175, 0, 255},
		.hsl = {48.83720930232558, 1, 0.4215686274509804},
		.hex = "#d7af00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 179,
		.rgba = {215, 175, 95, 255},
		.hsl = {40, 0.6000000000000001, 0.607843137254902},
		.hex = "#d7af5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 180,
		.rgba = {215, 175, 135, 255},
		.hsl = {30, 0.5000000000000001, 0.6862745098039216},
		.hex = "#d7af87",
		.names = "css:tan",
		.nameCount = 1,
	},
	{
		.code = 181,
		.rgba = {215, 175, 175, 255},
		.hsl = {0, 0.3333333333333334, 0.7647058823529411},
		.hex = "#d7afaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 182,
		.rgba = {215, 175, 215, 255},
		.hsl = {300, 0.3333333333333334, 0.7647058823529411},
		.hex = "#d7afd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 183,
		.rgba = {215, 175, 255, 255},
		.hsl = {270, 1, 0.8431372549019608},
		.hex = "#d7afff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 184,
		.rgba = {215, 215, 0, 255},
		.hsl = {60, 1, 0.4215686274509804},
		.hex = "#d7d700",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 185,
		.rgba = {215, 215, 95, 255},
		.hsl = {60, 0.6000000000000001, 0.607843137254902},
		.hex = "#d7d75f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 186,
		.rgba = {215, 215, 135, 255},
		.hsl = {60, 0.5000000000000001, 0.6862745098039216},
		.hex = "#d7d787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 187,
		.rgba = {215, 215, 175, 255},
		.hsl = {60, 0.3333333333333334, 0.7647058823529411},
		.hex = "#d7d7af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 188,
		.rgba = {215, 215, 215, 255},
		.hsl = {0, 0, 0.8431372549019608},
		.hex = "#d7d7d7",
		.names = "gray 3",
		.nameCount = 1,
	},
	{
		.code = 189,
		.rgba = {215, 215, 255, 255},
		.hsl = {240, 1, 0.9215686274509804},
		.hex = "#d7d7ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 190,
		.rgba = {215, 255, 0, 255},
		.hsl = {69.41176470588235, 1, 0.5},
		.hex = "#d7ff00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 191,
		.rgba = {215, 255, 95, 255},
		.hsl = {75, 1, 0.6862745098039216},
		.hex = "#d7ff5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 192,
		.rgba = {215, 255, 135, 255},
		.hsl = {80.00000000000001, 1, 0.7647058823529411},
		.hex = "#d7ff87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 193,
		.rgba = {215, 255, 175, 255},
		.hsl = {90, 1, 0.8431372549019608},
		.hex = "#d7ffaf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 194,
		.rgba = {215, 255, 215, 255},
		.hsl = {120, 1, 0.9215686274509804},
		.hex = "#d7ffd7",
		.names = "snowy mint",
		.nameCount = 1,
	},
	{
		.code = 195,
		.rgba = {215, 255, 255, 255},
		.hsl = {180, 1, 0.9215686274509804},
		.hex = "#d7ffff",
		.names = "css:lightcyan",
		.nameCount = 1,
	},
	{
		.code = 196,
		.rgba = {255, 0, 0, 255},
		.hsl = {0, 1, 0.5},
		.hex = "#ff0000",
		.names = "red,css:red",
		.nameCount = 2,
	},
	{
		.code = 197,
		.rgba = {255, 0, 95, 255},
		.hsl = {337.6470588235294, 1, 0.5},
		.hex = "#ff005f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 198,
		.rgba = {255, 0, 135, 255},
		.hsl = {328.2352941176471, 1, 0.5},
		.hex = "#ff0087",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 199,
		.rgba = {255, 0, 175, 255},
		.hsl = {318.8235294117647, 1, 0.5},
		.hex = "#ff00af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 200,
		.rgba = {255, 0, 215, 255},
		.hsl = {309.4117647058824, 1, 0.5},
		.hex = "#ff00d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 201,
		.rgba = {255, 0, 255, 255},
		.hsl = {300, 1, 0.5},
		.hex = "#ff00ff",
		.names = "css:fuchsia,css:magenta",
		.nameCount = 2,
	},
	{
		.code = 202,
		.rgba = {255, 95, 0, 255},
		.hsl = {22.352941176470587, 1, 0.5},
		.hex = "#ff5f00",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 203,
		.rgba = {255, 95, 95, 255},
		.hsl = {0, 1, 0.6862745098039216},
		.hex = "#ff5f5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 204,
		.rgba = {255, 95, 135, 255},
		.hsl = {345, 1, 0.6862745098039216},
		.hex = "#ff5f87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 205,
		.rgba = {255, 95, 175, 255},
		.hsl = {330, 1, 0.6862745098039216},
		.hex = "#ff5faf",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 206,
		.rgba = {255, 95, 215, 255},
		.hsl = {315, 1, 0.6862745098039216},
		.hex = "#ff5fd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 207,
		.rgba = {255, 95, 255, 255},
		.hsl = {300, 1, 0.6862745098039216},
		.hex = "#ff5fff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 208,
		.rgba = {255, 135, 0, 255},
		.hsl = {31.764705882352942, 1, 0.5},
		.hex = "#ff8700",
		.names = "css:darkorange",
		.nameCount = 1,
	},
	{
		.code = 209,
		.rgba = {255, 135, 95, 255},
		.hsl = {15, 1, 0.6862745098039216},
		.hex = "#ff875f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 210,
		.rgba = {255, 135, 135, 255},
		.hsl = {0, 1, 0.7647058823529411},
		.hex = "#ff8787",
		.names = "geraldine",
		.nameCount = 1,
	},
	{
		.code = 211,
		.rgba = {255, 135, 175, 255},
		.hsl = {340, 1, 0.7647058823529411},
		.hex = "#ff87af",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 212,
		.rgba = {255, 135, 215, 255},
		.hsl = {320, 1, 0.7647058823529411},
		.hex = "#ff87d7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 213,
		.rgba = {255, 135, 255, 255},
		.hsl = {300, 1, 0.7647058823529411},
		.hex = "#ff87ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 214,
		.rgba = {255, 175, 0, 255},
		.hsl = {41.1764705882353, 1, 0.5},
		.hex = "#ffaf00",
		.names = "css:orange",
		.nameCount = 1,
	},
	{
		.code = 215,
		.rgba = {255, 175, 95, 255},
		.hsl = {30, 1, 0.6862745098039216},
		.hex = "#ffaf5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 216,
		.rgba = {255, 175, 135, 255},
		.hsl = {20, 1, 0.7647058823529411},
		.hex = "#ffaf87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 217,
		.rgba = {255, 175, 175, 255},
		.hsl = {0, 1, 0.8431372549019608},
		.hex = "#ffafaf",
		.names = "cornflower lilac",
		.nameCount = 1,
	},
	{
		.code = 218,
		.rgba = {255, 175, 215, 255},
		.hsl = {330, 1, 0.8431372549019608},
		.hex = "#ffafd7",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 219,
		.rgba = {255, 175, 255, 255},
		.hsl = {300, 1, 0.8431372549019608},
		.hex = "#ffafff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 220,
		.rgba = {255, 215, 0, 255},
		.hsl = {50.588235294117645, 1, 0.5},
		.hex = "#ffd700",
		.names = "css:gold",
		.nameCount = 1,
	},
	{
		.code = 221,
		.rgba = {255, 215, 95, 255},
		.hsl = {45, 1, 0.6862745098039216},
		.hex = "#ffd75f",
		.names = "dandelion",
		.nameCount = 1,
	},
	{
		.code = 222,
		.rgba = {255, 215, 135, 255},
		.hsl = {40, 1, 0.7647058823529411},
		.hex = "#ffd787",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 223,
		.rgba = {255, 215, 175, 255},
		.hsl = {30, 1, 0.8431372549019608},
		.hex = "#ffd7af",
		.names = "css:navajowhite,light apricot",
		.nameCount = 2,
	},
	{
		.code = 224,
		.rgba = {255, 215, 215, 255},
		.hsl = {0, 1, 0.9215686274509804},
		.hex = "#ffd7d7",
		.names = "cosmos",
		.nameCount = 1,
	},
	{
		.code = 225,
		.rgba = {255, 215, 255, 255},
		.hsl = {300, 1, 0.9215686274509804},
		.hex = "#ffd7ff",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 226,
		.rgba = {255, 255, 0, 255},
		.hsl = {60, 1, 0.5},
		.hex = "#ffff00",
		.names = "css:yellow",
		.nameCount = 1,
	},
	{
		.code = 227,
		.rgba = {255, 255, 95, 255},
		.hsl = {60, 1, 0.6862745098039216},
		.hex = "#ffff5f",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 228,
		.rgba = {255, 255, 135, 255},
		.hsl = {60, 1, 0.7647058823529411},
		.hex = "#ffff87",
		.names = "",
		.nameCount = 0,
	},
	{
		.code = 229,
		.rgba = {255, 255, 175, 255},
		.hsl = {60, 1, 0.8431372549019608},
		.hex = "#ffffaf",
		.names = "portafino",
		.nameCount = 1,
	},
	{
		.code = 230,
		.rgba = {255, 255, 215, 255},
		.hsl = {60, 1, 0.9215686274509804},
		.hex = "#ffffd7",
		.names = "css:cornsilk,cumulus",
		.nameCount = 2,
	},
	{
		.code = 231,
		.rgba = {255, 255, 255, 255},
		.hsl = {0, 0, 1},
		.hex = "#ffffff",
		.names = "white,css:white",
		.nameCount = 2,
	},
	{
		.code = 232,
		.rgba = {8, 8, 8, 255},
		.hsl = {0, 0, 0.03137254901960784},
		.hex = "#080808",
		.names = "gray 24,cod gray",
		.nameCount = 2,
	},
	{
		.code = 233,
		.rgba = {18, 18, 18, 255},
		.hsl = {0, 0, 0.07058823529411765},
		.hex = "#121212",
		.names = "gray 23",
		.nameCount = 1,
	},
	{
		.code = 234,
		.rgba = {28, 28, 28, 255},
		.hsl = {0, 0, 0.10980392156862745},
		.hex = "#1c1c1c",
		.names = "gray 22",
		.nameCount = 1,
	},
	{
		.code = 235,
		.rgba = {38, 38, 38, 255},
		.hsl = {0, 0, 0.14901960784313725},
		.hex = "#262626",
		.names = "gray 21",
		.nameCount = 1,
	},
	{
		.code = 236,
		.rgba = {48, 48, 48, 255},
		.hsl = {0, 0, 0.18823529411764706},
		.hex = "#303030",
		.names = "gray 20,mine shaft",
		.nameCount = 2,
	},
	{
		.code = 237,
		.rgba = {58, 58, 58, 255},
		.hsl = {0, 0, 0.22745098039215686},
		.hex = "#3a3a3a",
		.names = "gray 19",
		.nameCount = 1,
	},
	{
		.code = 238,
		.rgba = {68, 68, 68, 255},
		.hsl = {0, 0, 0.26666666666666666},
		.hex = "#444444",
		.names = "gray 18",
		.nameCount = 1,
	},
	{
		.code = 239,
		.rgba = {78, 78, 78, 255},
		.hsl = {0, 0, 0.3058823529411765},
		.hex = "#4e4e4e",
		.names = "gray 17",
		.nameCount = 1,
	},
	{
		.code = 240,
		.rgba = {88, 88, 88, 255},
		.hsl = {0, 0, 0.34509803921568627},
		.hex = "#585858",
		.names = "gray 16",
		.nameCount = 1,
	},
	{
		.code = 241,
		.rgba = {98, 98, 98, 255},
		.hsl = {0, 0, 0.3843137254901961},
		.hex = "#626262",
		.names = "gray 15",
		.nameCount = 1,
	},
	{
		.code = 242,
		.rgba = {108, 108, 108, 255},
		.hsl = {0, 0, 0.4235294117647059},
		.hex = "#6c6c6c",
		.names = "gray 14,css:dimgray,dove gray",
		.nameCount = 3,
	},
	{
		.code = 243,
		.rgba = {118, 118, 118, 255},
		.hsl = {0, 0, 0.4627450980392157},
		.hex = "#767676",
		.names = "gray 13",
		.nameCount = 1,
	},
	{
		.code = 244,
		.rgba = {128, 128, 128, 255},
		.hsl = {0, 0, 0.5019607843137255},
		.hex = "#808080",
		.names = "gray 12,css:gray,css:grey",
		.nameCount = 3,
	},
	{
		.code = 245,
		.rgba = {138, 138, 138, 255},
		.hsl = {0, 0, 0.5411764705882353},
		.hex = "#8a8a8a",
		.names = "gray 11",
		.nameCount = 1,
	},
	{
		.code = 246,
		.rgba = {148, 148, 148, 255},
		.hsl = {0, 0, 0.5803921568627451},
		.hex = "#949494",
		.names = "gray 10",
		.nameCount = 1,
	},
	{
		.code = 247,
		.rgba = {158, 158, 158, 255},
		.hsl = {0, 0, 0.6196078431372549},
		.hex = "#9e9e9e",
		.names = "gray 9",
		.nameCount = 1,
	},
	{
		.code = 248,
		.rgba = {168, 168, 168, 255},
		.hsl = {0, 0, 0.6588235294117647},
		.hex = "#a8a8a8",
		.names = "gray 8,css:darkgrey",
		.nameCount = 2,
	},
	{
		.code = 249,
		.rgba = {178, 178, 178, 255},
		.hsl = {0, 0, 0.6980392156862745},
		.hex = "#b2b2b2",
		.names = "gray 7",
		.nameCount = 1,
	},
	{
		.code = 250,
		.rgba = {188, 188, 188, 255},
		.hsl = {0, 0, 0.7372549019607844},
		.hex = "#bcbcbc",
		.names = "gray 6,css:silver",
		.nameCount = 2,
	},
	{
		.code = 251,
		.rgba = {198, 198, 198, 255},
		.hsl = {0, 0, 0.7764705882352941},
		.hex = "#c6c6c6",
		.names = "gray 5",
		.nameCount = 1,
	},
	{
		.code = 252,
		.rgba = {208, 208, 208, 255},
		.hsl = {0, 0, 0.8156862745098039},
		.hex = "#d0d0d0",
		.names = "gray 4,css:lightgray",
		.nameCount = 2,
	},
	{
		.code = 253,
		.rgba = {218, 218, 218, 255},
		.hsl = {0, 0, 0.8549019607843137},
		.hex = "#dadada",
		.names = "gray 2,css:gainsboro,alto",
		.nameCount = 3,
	},
	{
		.code = 254,
		.rgba = {228, 228, 228, 255},
		.hsl = {0, 0, 0.8941176470588236},
		.hex = "#e4e4e4",
		.names = "gray 1,mercury",
		.nameCount = 2,
	},
	{
		.code = 255,
		.rgba = {238, 238, 238, 255},
		.hsl = {0, 0, 0.9333333333333333},
		.hex = "#eeeeee",
		.names = "gray 0,gallery",
		.nameCount = 2,
	},
};