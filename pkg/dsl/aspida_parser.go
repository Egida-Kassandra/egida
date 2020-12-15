// Code generated from aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package dsl // aspida

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 45, 450,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 3, 2, 3, 2, 3, 2,
	5, 2, 68, 10, 2, 3, 2, 5, 2, 71, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 97, 10, 7, 12, 7, 14, 7, 100,
	11, 7, 3, 8, 3, 8, 3, 8, 5, 8, 105, 10, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 5, 9, 115, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 5, 10, 123, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 131, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 137, 10, 12, 12, 12,
	14, 12, 140, 11, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 146, 10, 12, 3,
	13, 3, 13, 5, 13, 150, 10, 13, 3, 14, 3, 14, 7, 14, 154, 10, 14, 12, 14,
	14, 14, 157, 11, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 163, 10, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 173, 10, 16,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 181, 10, 17, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 189, 10, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 197, 10, 19, 3, 20, 3, 20, 7, 20, 201, 10,
	20, 12, 20, 14, 20, 204, 11, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 5, 21, 218, 10, 21, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 226, 10, 22, 3, 23, 3, 23, 3,
	23, 3, 23, 7, 23, 232, 10, 23, 12, 23, 14, 23, 235, 11, 23, 3, 23, 3, 23,
	3, 23, 3, 23, 5, 23, 241, 10, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3,
	25, 7, 25, 249, 10, 25, 12, 25, 14, 25, 252, 11, 25, 3, 25, 3, 25, 3, 25,
	3, 25, 5, 25, 258, 10, 25, 3, 26, 3, 26, 5, 26, 262, 10, 26, 3, 27, 3,
	27, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29,
	296, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 312, 10, 29, 3, 29, 5, 29,
	315, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 5, 29, 329, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29,
	334, 10, 29, 3, 29, 5, 29, 337, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5, 29, 349, 10, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 354, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 359, 10, 29, 3,
	29, 5, 29, 362, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 372, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 377, 10, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 382, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 387,
	10, 29, 3, 29, 5, 29, 390, 10, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5,
	29, 397, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 402, 10, 29, 3, 29, 3, 29,
	3, 29, 5, 29, 407, 10, 29, 3, 29, 3, 29, 3, 29, 5, 29, 412, 10, 29, 3,
	29, 3, 29, 3, 29, 5, 29, 417, 10, 29, 3, 29, 5, 29, 420, 10, 29, 3, 29,
	3, 29, 5, 29, 424, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 439, 10, 30, 3, 31,
	3, 31, 3, 31, 3, 31, 3, 31, 5, 31, 446, 10, 31, 3, 32, 3, 32, 3, 32, 2,
	2, 33, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34,
	36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 2, 4, 3, 2, 43,
	45, 3, 2, 32, 33, 2, 481, 2, 64, 3, 2, 2, 2, 4, 72, 3, 2, 2, 2, 6, 78,
	3, 2, 2, 2, 8, 82, 3, 2, 2, 2, 10, 88, 3, 2, 2, 2, 12, 94, 3, 2, 2, 2,
	14, 104, 3, 2, 2, 2, 16, 114, 3, 2, 2, 2, 18, 122, 3, 2, 2, 2, 20, 130,
	3, 2, 2, 2, 22, 145, 3, 2, 2, 2, 24, 149, 3, 2, 2, 2, 26, 151, 3, 2, 2,
	2, 28, 162, 3, 2, 2, 2, 30, 172, 3, 2, 2, 2, 32, 180, 3, 2, 2, 2, 34, 188,
	3, 2, 2, 2, 36, 196, 3, 2, 2, 2, 38, 198, 3, 2, 2, 2, 40, 217, 3, 2, 2,
	2, 42, 225, 3, 2, 2, 2, 44, 240, 3, 2, 2, 2, 46, 242, 3, 2, 2, 2, 48, 257,
	3, 2, 2, 2, 50, 261, 3, 2, 2, 2, 52, 263, 3, 2, 2, 2, 54, 265, 3, 2, 2,
	2, 56, 423, 3, 2, 2, 2, 58, 438, 3, 2, 2, 2, 60, 445, 3, 2, 2, 2, 62, 447,
	3, 2, 2, 2, 64, 65, 5, 4, 3, 2, 65, 67, 5, 6, 4, 2, 66, 68, 5, 8, 5, 2,
	67, 66, 3, 2, 2, 2, 67, 68, 3, 2, 2, 2, 68, 70, 3, 2, 2, 2, 69, 71, 5,
	10, 6, 2, 70, 69, 3, 2, 2, 2, 70, 71, 3, 2, 2, 2, 71, 3, 3, 2, 2, 2, 72,
	73, 7, 39, 2, 2, 73, 74, 7, 3, 2, 2, 74, 75, 7, 4, 2, 2, 75, 76, 5, 12,
	7, 2, 76, 77, 7, 5, 2, 2, 77, 5, 3, 2, 2, 2, 78, 79, 7, 40, 2, 2, 79, 80,
	7, 3, 2, 2, 80, 81, 5, 22, 12, 2, 81, 7, 3, 2, 2, 2, 82, 83, 7, 41, 2,
	2, 83, 84, 7, 3, 2, 2, 84, 85, 7, 4, 2, 2, 85, 86, 5, 26, 14, 2, 86, 87,
	7, 5, 2, 2, 87, 9, 3, 2, 2, 2, 88, 89, 7, 42, 2, 2, 89, 90, 7, 3, 2, 2,
	90, 91, 7, 4, 2, 2, 91, 92, 5, 38, 20, 2, 92, 93, 7, 5, 2, 2, 93, 11, 3,
	2, 2, 2, 94, 98, 5, 14, 8, 2, 95, 97, 5, 14, 8, 2, 96, 95, 3, 2, 2, 2,
	97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3, 2, 2, 2, 99, 13, 3,
	2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 105, 5, 16, 9, 2, 102, 105, 5, 18, 10,
	2, 103, 105, 5, 20, 11, 2, 104, 101, 3, 2, 2, 2, 104, 102, 3, 2, 2, 2,
	104, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 7, 38, 2, 2, 107,
	15, 3, 2, 2, 2, 108, 109, 7, 6, 2, 2, 109, 110, 7, 3, 2, 2, 110, 115, 7,
	30, 2, 2, 111, 112, 7, 7, 2, 2, 112, 113, 7, 3, 2, 2, 113, 115, 7, 30,
	2, 2, 114, 108, 3, 2, 2, 2, 114, 111, 3, 2, 2, 2, 115, 17, 3, 2, 2, 2,
	116, 117, 7, 8, 2, 2, 117, 118, 7, 3, 2, 2, 118, 123, 9, 2, 2, 2, 119,
	120, 7, 9, 2, 2, 120, 121, 7, 3, 2, 2, 121, 123, 9, 2, 2, 2, 122, 116,
	3, 2, 2, 2, 122, 119, 3, 2, 2, 2, 123, 19, 3, 2, 2, 2, 124, 125, 7, 10,
	2, 2, 125, 126, 7, 3, 2, 2, 126, 131, 7, 30, 2, 2, 127, 128, 7, 11, 2,
	2, 128, 129, 7, 3, 2, 2, 129, 131, 7, 30, 2, 2, 130, 124, 3, 2, 2, 2, 130,
	127, 3, 2, 2, 2, 131, 21, 3, 2, 2, 2, 132, 133, 7, 12, 2, 2, 133, 138,
	5, 24, 13, 2, 134, 135, 7, 13, 2, 2, 135, 137, 5, 24, 13, 2, 136, 134,
	3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2, 138, 139, 3, 2,
	2, 2, 139, 141, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141, 142, 7, 14, 2, 2,
	142, 146, 3, 2, 2, 2, 143, 144, 7, 12, 2, 2, 144, 146, 7, 14, 2, 2, 145,
	132, 3, 2, 2, 2, 145, 143, 3, 2, 2, 2, 146, 23, 3, 2, 2, 2, 147, 150, 5,
	50, 26, 2, 148, 150, 5, 56, 29, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2,
	2, 2, 150, 25, 3, 2, 2, 2, 151, 155, 5, 28, 15, 2, 152, 154, 5, 28, 15,
	2, 153, 152, 3, 2, 2, 2, 154, 157, 3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155,
	156, 3, 2, 2, 2, 156, 27, 3, 2, 2, 2, 157, 155, 3, 2, 2, 2, 158, 163, 5,
	30, 16, 2, 159, 163, 5, 32, 17, 2, 160, 163, 5, 34, 18, 2, 161, 163, 5,
	36, 19, 2, 162, 158, 3, 2, 2, 2, 162, 159, 3, 2, 2, 2, 162, 160, 3, 2,
	2, 2, 162, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 7, 38, 2, 2,
	165, 29, 3, 2, 2, 2, 166, 167, 7, 15, 2, 2, 167, 168, 7, 3, 2, 2, 168,
	173, 5, 44, 23, 2, 169, 170, 7, 16, 2, 2, 170, 171, 7, 3, 2, 2, 171, 173,
	5, 44, 23, 2, 172, 166, 3, 2, 2, 2, 172, 169, 3, 2, 2, 2, 173, 31, 3, 2,
	2, 2, 174, 175, 7, 17, 2, 2, 175, 176, 7, 3, 2, 2, 176, 181, 5, 44, 23,
	2, 177, 178, 7, 18, 2, 2, 178, 179, 7, 3, 2, 2, 179, 181, 5, 44, 23, 2,
	180, 174, 3, 2, 2, 2, 180, 177, 3, 2, 2, 2, 181, 33, 3, 2, 2, 2, 182, 183,
	7, 19, 2, 2, 183, 184, 7, 3, 2, 2, 184, 189, 5, 44, 23, 2, 185, 186, 7,
	20, 2, 2, 186, 187, 7, 3, 2, 2, 187, 189, 5, 44, 23, 2, 188, 182, 3, 2,
	2, 2, 188, 185, 3, 2, 2, 2, 189, 35, 3, 2, 2, 2, 190, 191, 7, 21, 2, 2,
	191, 192, 7, 3, 2, 2, 192, 197, 5, 44, 23, 2, 193, 194, 7, 22, 2, 2, 194,
	195, 7, 3, 2, 2, 195, 197, 5, 44, 23, 2, 196, 190, 3, 2, 2, 2, 196, 193,
	3, 2, 2, 2, 197, 37, 3, 2, 2, 2, 198, 202, 5, 40, 21, 2, 199, 201, 5, 40,
	21, 2, 200, 199, 3, 2, 2, 2, 201, 204, 3, 2, 2, 2, 202, 200, 3, 2, 2, 2,
	202, 203, 3, 2, 2, 2, 203, 39, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 205, 206,
	7, 30, 2, 2, 206, 207, 7, 3, 2, 2, 207, 208, 5, 42, 22, 2, 208, 209, 7,
	38, 2, 2, 209, 218, 3, 2, 2, 2, 210, 211, 7, 30, 2, 2, 211, 212, 7, 3,
	2, 2, 212, 213, 7, 4, 2, 2, 213, 214, 5, 38, 20, 2, 214, 215, 7, 5, 2,
	2, 215, 216, 7, 38, 2, 2, 216, 218, 3, 2, 2, 2, 217, 205, 3, 2, 2, 2, 217,
	210, 3, 2, 2, 2, 218, 41, 3, 2, 2, 2, 219, 226, 7, 30, 2, 2, 220, 226,
	7, 31, 2, 2, 221, 226, 7, 23, 2, 2, 222, 226, 7, 24, 2, 2, 223, 226, 7,
	25, 2, 2, 224, 226, 5, 48, 25, 2, 225, 219, 3, 2, 2, 2, 225, 220, 3, 2,
	2, 2, 225, 221, 3, 2, 2, 2, 225, 222, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2,
	225, 224, 3, 2, 2, 2, 226, 43, 3, 2, 2, 2, 227, 228, 7, 12, 2, 2, 228,
	233, 5, 46, 24, 2, 229, 230, 7, 13, 2, 2, 230, 232, 5, 46, 24, 2, 231,
	229, 3, 2, 2, 2, 232, 235, 3, 2, 2, 2, 233, 231, 3, 2, 2, 2, 233, 234,
	3, 2, 2, 2, 234, 236, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 236, 237, 7, 14,
	2, 2, 237, 241, 3, 2, 2, 2, 238, 239, 7, 12, 2, 2, 239, 241, 7, 14, 2,
	2, 240, 227, 3, 2, 2, 2, 240, 238, 3, 2, 2, 2, 241, 45, 3, 2, 2, 2, 242,
	243, 7, 30, 2, 2, 243, 47, 3, 2, 2, 2, 244, 245, 7, 12, 2, 2, 245, 250,
	5, 42, 22, 2, 246, 247, 7, 13, 2, 2, 247, 249, 5, 42, 22, 2, 248, 246,
	3, 2, 2, 2, 249, 252, 3, 2, 2, 2, 250, 248, 3, 2, 2, 2, 250, 251, 3, 2,
	2, 2, 251, 253, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 253, 254, 7, 14, 2, 2,
	254, 258, 3, 2, 2, 2, 255, 256, 7, 12, 2, 2, 256, 258, 7, 14, 2, 2, 257,
	244, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2, 258, 49, 3, 2, 2, 2, 259, 262, 5,
	52, 27, 2, 260, 262, 5, 54, 28, 2, 261, 259, 3, 2, 2, 2, 261, 260, 3, 2,
	2, 2, 262, 51, 3, 2, 2, 2, 263, 264, 7, 36, 2, 2, 264, 53, 3, 2, 2, 2,
	265, 266, 7, 37, 2, 2, 266, 55, 3, 2, 2, 2, 267, 268, 5, 58, 30, 2, 268,
	269, 7, 3, 2, 2, 269, 270, 5, 58, 30, 2, 270, 271, 7, 3, 2, 2, 271, 272,
	5, 58, 30, 2, 272, 273, 7, 3, 2, 2, 273, 274, 5, 58, 30, 2, 274, 275, 7,
	3, 2, 2, 275, 276, 5, 58, 30, 2, 276, 277, 7, 3, 2, 2, 277, 278, 5, 58,
	30, 2, 278, 279, 7, 3, 2, 2, 279, 280, 5, 60, 31, 2, 280, 424, 3, 2, 2,
	2, 281, 282, 7, 26, 2, 2, 282, 283, 5, 58, 30, 2, 283, 284, 7, 3, 2, 2,
	284, 285, 5, 58, 30, 2, 285, 286, 7, 3, 2, 2, 286, 287, 5, 58, 30, 2, 287,
	288, 7, 3, 2, 2, 288, 289, 5, 58, 30, 2, 289, 290, 7, 3, 2, 2, 290, 291,
	5, 58, 30, 2, 291, 292, 7, 3, 2, 2, 292, 293, 5, 60, 31, 2, 293, 424, 3,
	2, 2, 2, 294, 296, 5, 58, 30, 2, 295, 294, 3, 2, 2, 2, 295, 296, 3, 2,
	2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 7, 26, 2, 2, 298, 299, 5, 58, 30,
	2, 299, 300, 7, 3, 2, 2, 300, 301, 5, 58, 30, 2, 301, 302, 7, 3, 2, 2,
	302, 303, 5, 58, 30, 2, 303, 304, 7, 3, 2, 2, 304, 305, 5, 58, 30, 2, 305,
	306, 7, 3, 2, 2, 306, 307, 5, 60, 31, 2, 307, 424, 3, 2, 2, 2, 308, 309,
	5, 58, 30, 2, 309, 310, 7, 3, 2, 2, 310, 312, 3, 2, 2, 2, 311, 308, 3,
	2, 2, 2, 311, 312, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313, 315, 5, 58, 30,
	2, 314, 311, 3, 2, 2, 2, 314, 315, 3, 2, 2, 2, 315, 316, 3, 2, 2, 2, 316,
	317, 7, 26, 2, 2, 317, 318, 5, 58, 30, 2, 318, 319, 7, 3, 2, 2, 319, 320,
	5, 58, 30, 2, 320, 321, 7, 3, 2, 2, 321, 322, 5, 58, 30, 2, 322, 323, 7,
	3, 2, 2, 323, 324, 5, 60, 31, 2, 324, 424, 3, 2, 2, 2, 325, 326, 5, 58,
	30, 2, 326, 327, 7, 3, 2, 2, 327, 329, 3, 2, 2, 2, 328, 325, 3, 2, 2, 2,
	328, 329, 3, 2, 2, 2, 329, 330, 3, 2, 2, 2, 330, 331, 5, 58, 30, 2, 331,
	332, 7, 3, 2, 2, 332, 334, 3, 2, 2, 2, 333, 328, 3, 2, 2, 2, 333, 334,
	3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335, 337, 5, 58, 30, 2, 336, 333, 3,
	2, 2, 2, 336, 337, 3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 339, 7, 26, 2,
	2, 339, 340, 5, 58, 30, 2, 340, 341, 7, 3, 2, 2, 341, 342, 5, 58, 30, 2,
	342, 343, 7, 3, 2, 2, 343, 344, 5, 60, 31, 2, 344, 424, 3, 2, 2, 2, 345,
	346, 5, 58, 30, 2, 346, 347, 7, 3, 2, 2, 347, 349, 3, 2, 2, 2, 348, 345,
	3, 2, 2, 2, 348, 349, 3, 2, 2, 2, 349, 350, 3, 2, 2, 2, 350, 351, 5, 58,
	30, 2, 351, 352, 7, 3, 2, 2, 352, 354, 3, 2, 2, 2, 353, 348, 3, 2, 2, 2,
	353, 354, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 356, 5, 58, 30, 2, 356,
	357, 7, 3, 2, 2, 357, 359, 3, 2, 2, 2, 358, 353, 3, 2, 2, 2, 358, 359,
	3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 362, 5, 58, 30, 2, 361, 358, 3,
	2, 2, 2, 361, 362, 3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 364, 7, 26, 2,
	2, 364, 365, 5, 58, 30, 2, 365, 366, 7, 3, 2, 2, 366, 367, 5, 60, 31, 2,
	367, 424, 3, 2, 2, 2, 368, 369, 5, 58, 30, 2, 369, 370, 7, 3, 2, 2, 370,
	372, 3, 2, 2, 2, 371, 368, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 373,
	3, 2, 2, 2, 373, 374, 5, 58, 30, 2, 374, 375, 7, 3, 2, 2, 375, 377, 3,
	2, 2, 2, 376, 371, 3, 2, 2, 2, 376, 377, 3, 2, 2, 2, 377, 378, 3, 2, 2,
	2, 378, 379, 5, 58, 30, 2, 379, 380, 7, 3, 2, 2, 380, 382, 3, 2, 2, 2,
	381, 376, 3, 2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383,
	384, 5, 58, 30, 2, 384, 385, 7, 3, 2, 2, 385, 387, 3, 2, 2, 2, 386, 381,
	3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 388, 3, 2, 2, 2, 388, 390, 5, 58,
	30, 2, 389, 386, 3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 391, 3, 2, 2, 2,
	391, 392, 7, 26, 2, 2, 392, 424, 5, 60, 31, 2, 393, 394, 5, 58, 30, 2,
	394, 395, 7, 3, 2, 2, 395, 397, 3, 2, 2, 2, 396, 393, 3, 2, 2, 2, 396,
	397, 3, 2, 2, 2, 397, 398, 3, 2, 2, 2, 398, 399, 5, 58, 30, 2, 399, 400,
	7, 3, 2, 2, 400, 402, 3, 2, 2, 2, 401, 396, 3, 2, 2, 2, 401, 402, 3, 2,
	2, 2, 402, 403, 3, 2, 2, 2, 403, 404, 5, 58, 30, 2, 404, 405, 7, 3, 2,
	2, 405, 407, 3, 2, 2, 2, 406, 401, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407,
	408, 3, 2, 2, 2, 408, 409, 5, 58, 30, 2, 409, 410, 7, 3, 2, 2, 410, 412,
	3, 2, 2, 2, 411, 406, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 413, 3, 2,
	2, 2, 413, 414, 5, 58, 30, 2, 414, 415, 7, 3, 2, 2, 415, 417, 3, 2, 2,
	2, 416, 411, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 418, 3, 2, 2, 2, 418,
	420, 5, 58, 30, 2, 419, 416, 3, 2, 2, 2, 419, 420, 3, 2, 2, 2, 420, 421,
	3, 2, 2, 2, 421, 422, 7, 26, 2, 2, 422, 424, 5, 58, 30, 2, 423, 267, 3,
	2, 2, 2, 423, 281, 3, 2, 2, 2, 423, 295, 3, 2, 2, 2, 423, 314, 3, 2, 2,
	2, 423, 336, 3, 2, 2, 2, 423, 361, 3, 2, 2, 2, 423, 389, 3, 2, 2, 2, 423,
	419, 3, 2, 2, 2, 424, 57, 3, 2, 2, 2, 425, 426, 5, 62, 32, 2, 426, 427,
	5, 62, 32, 2, 427, 428, 5, 62, 32, 2, 428, 429, 5, 62, 32, 2, 429, 439,
	3, 2, 2, 2, 430, 431, 5, 62, 32, 2, 431, 432, 5, 62, 32, 2, 432, 433, 5,
	62, 32, 2, 433, 439, 3, 2, 2, 2, 434, 435, 5, 62, 32, 2, 435, 436, 5, 62,
	32, 2, 436, 439, 3, 2, 2, 2, 437, 439, 5, 62, 32, 2, 438, 425, 3, 2, 2,
	2, 438, 430, 3, 2, 2, 2, 438, 434, 3, 2, 2, 2, 438, 437, 3, 2, 2, 2, 439,
	59, 3, 2, 2, 2, 440, 441, 5, 58, 30, 2, 441, 442, 7, 3, 2, 2, 442, 443,
	5, 58, 30, 2, 443, 446, 3, 2, 2, 2, 444, 446, 5, 50, 26, 2, 445, 440, 3,
	2, 2, 2, 445, 444, 3, 2, 2, 2, 446, 61, 3, 2, 2, 2, 447, 448, 9, 3, 2,
	2, 448, 63, 3, 2, 2, 2, 50, 67, 70, 98, 104, 114, 122, 130, 138, 145, 149,
	155, 162, 172, 180, 188, 196, 202, 217, 225, 233, 240, 250, 257, 261, 295,
	311, 314, 328, 333, 336, 348, 353, 358, 361, 371, 376, 381, 386, 389, 396,
	401, 406, 411, 416, 419, 423, 438, 445,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'{'", "'}'", "'name'", "'NAME'", "'connection'", "'CONNECTION'",
	"'description'", "'DESCRIPTION'", "'['", "','", "']'", "'sections'", "'SECTIONS'",
	"'points'", "'POINTS'", "'controls'", "'CONTROLS'", "'exclusions'", "'EXCLUSIONS'",
	"'true'", "'false'", "'null'", "'::'", "", "", "", "", "", "", "", "",
	"", "", "", "';'", "'MAIN'", "'HOSTS'", "'TASKS'", "'VARS'", "'LOCAL'",
	"'SSH'", "'SSH-PASSWORD'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "COMMENT", "WHITESPACE", "NEWLINE", "STRING",
	"NUMBER", "DIGIT", "HEX_CHAR", "NUMBER_RANGE", "ARROW", "SINGLE_IP", "IP_RANGE",
	"NS", "MAIN_KW", "HOSTS_KW", "TASKS_KW", "VARS_KW", "LOCAL_KW", "SSH_KW",
	"SSHPASS_KW",
}

var ruleNames = []string{
	"program", "main", "hosts", "tasks", "variables", "main_content", "main_prop",
	"name", "connection", "description", "hosts_list", "host", "tasks_content",
	"tasks_prop", "sections", "points", "controls", "exclusions", "vars_content",
	"vars_prop", "value", "str_array", "cadena", "array", "ip_v4", "single_ip",
	"ip_range", "ip_v6", "h16", "ls32", "hexdigit",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type aspidaParser struct {
	*antlr.BaseParser
}

func NewaspidaParser(input antlr.TokenStream) *aspidaParser {
	this := new(aspidaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "aspida.g4"

	return this
}

// aspidaParser tokens.
const (
	aspidaParserEOF          = antlr.TokenEOF
	aspidaParserT__0         = 1
	aspidaParserT__1         = 2
	aspidaParserT__2         = 3
	aspidaParserT__3         = 4
	aspidaParserT__4         = 5
	aspidaParserT__5         = 6
	aspidaParserT__6         = 7
	aspidaParserT__7         = 8
	aspidaParserT__8         = 9
	aspidaParserT__9         = 10
	aspidaParserT__10        = 11
	aspidaParserT__11        = 12
	aspidaParserT__12        = 13
	aspidaParserT__13        = 14
	aspidaParserT__14        = 15
	aspidaParserT__15        = 16
	aspidaParserT__16        = 17
	aspidaParserT__17        = 18
	aspidaParserT__18        = 19
	aspidaParserT__19        = 20
	aspidaParserT__20        = 21
	aspidaParserT__21        = 22
	aspidaParserT__22        = 23
	aspidaParserT__23        = 24
	aspidaParserCOMMENT      = 25
	aspidaParserWHITESPACE   = 26
	aspidaParserNEWLINE      = 27
	aspidaParserSTRING       = 28
	aspidaParserNUMBER       = 29
	aspidaParserDIGIT        = 30
	aspidaParserHEX_CHAR     = 31
	aspidaParserNUMBER_RANGE = 32
	aspidaParserARROW        = 33
	aspidaParserSINGLE_IP    = 34
	aspidaParserIP_RANGE     = 35
	aspidaParserNS           = 36
	aspidaParserMAIN_KW      = 37
	aspidaParserHOSTS_KW     = 38
	aspidaParserTASKS_KW     = 39
	aspidaParserVARS_KW      = 40
	aspidaParserLOCAL_KW     = 41
	aspidaParserSSH_KW       = 42
	aspidaParserSSHPASS_KW   = 43
)

// aspidaParser rules.
const (
	aspidaParserRULE_program       = 0
	aspidaParserRULE_main          = 1
	aspidaParserRULE_hosts         = 2
	aspidaParserRULE_tasks         = 3
	aspidaParserRULE_variables     = 4
	aspidaParserRULE_main_content  = 5
	aspidaParserRULE_main_prop     = 6
	aspidaParserRULE_name          = 7
	aspidaParserRULE_connection    = 8
	aspidaParserRULE_description   = 9
	aspidaParserRULE_hosts_list    = 10
	aspidaParserRULE_host          = 11
	aspidaParserRULE_tasks_content = 12
	aspidaParserRULE_tasks_prop    = 13
	aspidaParserRULE_sections      = 14
	aspidaParserRULE_points        = 15
	aspidaParserRULE_controls      = 16
	aspidaParserRULE_exclusions    = 17
	aspidaParserRULE_vars_content  = 18
	aspidaParserRULE_vars_prop     = 19
	aspidaParserRULE_value         = 20
	aspidaParserRULE_str_array     = 21
	aspidaParserRULE_cadena        = 22
	aspidaParserRULE_array         = 23
	aspidaParserRULE_ip_v4         = 24
	aspidaParserRULE_single_ip     = 25
	aspidaParserRULE_ip_range      = 26
	aspidaParserRULE_ip_v6         = 27
	aspidaParserRULE_h16           = 28
	aspidaParserRULE_ls32          = 29
	aspidaParserRULE_hexdigit      = 30
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) Main() IMainContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMainContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMainContext)
}

func (s *ProgramContext) Hosts() IHostsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHostsContext)
}

func (s *ProgramContext) Tasks() ITasksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasksContext)
}

func (s *ProgramContext) Variables() IVariablesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariablesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariablesContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, aspidaParserRULE_program)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.Main()
	}
	{
		p.SetState(63)
		p.Hosts()
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == aspidaParserTASKS_KW {
		{
			p.SetState(64)
			p.Tasks()
		}

	}
	p.SetState(68)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == aspidaParserVARS_KW {
		{
			p.SetState(67)
			p.Variables()
		}

	}

	return localctx
}

// IMainContext is an interface to support dynamic dispatch.
type IMainContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMainContext differentiates from other interfaces.
	IsMainContext()
}

type MainContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMainContext() *MainContext {
	var p = new(MainContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) MAIN_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserMAIN_KW, 0)
}

func (s *MainContext) Main_content() IMain_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMain_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMain_contentContext)
}

func (s *MainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MainContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MainContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterMain(s)
	}
}

func (s *MainContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitMain(s)
	}
}

func (s *MainContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitMain(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, aspidaParserRULE_main)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(aspidaParserMAIN_KW)
	}
	{
		p.SetState(71)
		p.Match(aspidaParserT__0)
	}
	{
		p.SetState(72)
		p.Match(aspidaParserT__1)
	}
	{
		p.SetState(73)
		p.Main_content()
	}
	{
		p.SetState(74)
		p.Match(aspidaParserT__2)
	}

	return localctx
}

// IHostsContext is an interface to support dynamic dispatch.
type IHostsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostsContext differentiates from other interfaces.
	IsHostsContext()
}

type HostsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostsContext() *HostsContext {
	var p = new(HostsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_hosts
	return p
}

func (*HostsContext) IsHostsContext() {}

func NewHostsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostsContext {
	var p = new(HostsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_hosts

	return p
}

func (s *HostsContext) GetParser() antlr.Parser { return s.parser }

func (s *HostsContext) HOSTS_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserHOSTS_KW, 0)
}

func (s *HostsContext) Hosts_list() IHosts_listContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHosts_listContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHosts_listContext)
}

func (s *HostsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterHosts(s)
	}
}

func (s *HostsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitHosts(s)
	}
}

func (s *HostsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitHosts(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Hosts() (localctx IHostsContext) {
	localctx = NewHostsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, aspidaParserRULE_hosts)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(aspidaParserHOSTS_KW)
	}
	{
		p.SetState(77)
		p.Match(aspidaParserT__0)
	}
	{
		p.SetState(78)
		p.Hosts_list()
	}

	return localctx
}

// ITasksContext is an interface to support dynamic dispatch.
type ITasksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasksContext differentiates from other interfaces.
	IsTasksContext()
}

type TasksContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasksContext() *TasksContext {
	var p = new(TasksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_tasks
	return p
}

func (*TasksContext) IsTasksContext() {}

func NewTasksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TasksContext {
	var p = new(TasksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_tasks

	return p
}

func (s *TasksContext) GetParser() antlr.Parser { return s.parser }

func (s *TasksContext) TASKS_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserTASKS_KW, 0)
}

func (s *TasksContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *TasksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TasksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TasksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterTasks(s)
	}
}

func (s *TasksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitTasks(s)
	}
}

func (s *TasksContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTasks(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Tasks() (localctx ITasksContext) {
	localctx = NewTasksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, aspidaParserRULE_tasks)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(80)
		p.Match(aspidaParserTASKS_KW)
	}
	{
		p.SetState(81)
		p.Match(aspidaParserT__0)
	}
	{
		p.SetState(82)
		p.Match(aspidaParserT__1)
	}
	{
		p.SetState(83)
		p.Tasks_content()
	}
	{
		p.SetState(84)
		p.Match(aspidaParserT__2)
	}

	return localctx
}

// IVariablesContext is an interface to support dynamic dispatch.
type IVariablesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariablesContext differentiates from other interfaces.
	IsVariablesContext()
}

type VariablesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariablesContext() *VariablesContext {
	var p = new(VariablesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_variables
	return p
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) VARS_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserVARS_KW, 0)
}

func (s *VariablesContext) Vars_content() IVars_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_contentContext)
}

func (s *VariablesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariablesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariablesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterVariables(s)
	}
}

func (s *VariablesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitVariables(s)
	}
}

func (s *VariablesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVariables(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Variables() (localctx IVariablesContext) {
	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, aspidaParserRULE_variables)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(aspidaParserVARS_KW)
	}
	{
		p.SetState(87)
		p.Match(aspidaParserT__0)
	}
	{
		p.SetState(88)
		p.Match(aspidaParserT__1)
	}
	{
		p.SetState(89)
		p.Vars_content()
	}
	{
		p.SetState(90)
		p.Match(aspidaParserT__2)
	}

	return localctx
}

// IMain_contentContext is an interface to support dynamic dispatch.
type IMain_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMain_contentContext differentiates from other interfaces.
	IsMain_contentContext()
}

type Main_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_contentContext() *Main_contentContext {
	var p = new(Main_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_main_content
	return p
}

func (*Main_contentContext) IsMain_contentContext() {}

func NewMain_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_contentContext {
	var p = new(Main_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_main_content

	return p
}

func (s *Main_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_contentContext) AllMain_prop() []IMain_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMain_propContext)(nil)).Elem())
	var tst = make([]IMain_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMain_propContext)
		}
	}

	return tst
}

func (s *Main_contentContext) Main_prop(i int) IMain_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMain_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMain_propContext)
}

func (s *Main_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Main_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterMain_content(s)
	}
}

func (s *Main_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitMain_content(s)
	}
}

func (s *Main_contentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitMain_content(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Main_content() (localctx IMain_contentContext) {
	localctx = NewMain_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, aspidaParserRULE_main_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Main_prop()
	}
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<aspidaParserT__3)|(1<<aspidaParserT__4)|(1<<aspidaParserT__5)|(1<<aspidaParserT__6)|(1<<aspidaParserT__7)|(1<<aspidaParserT__8))) != 0 {
		{
			p.SetState(93)
			p.Main_prop()
		}

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IMain_propContext is an interface to support dynamic dispatch.
type IMain_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMain_propContext differentiates from other interfaces.
	IsMain_propContext()
}

type Main_propContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMain_propContext() *Main_propContext {
	var p = new(Main_propContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_main_prop
	return p
}

func (*Main_propContext) IsMain_propContext() {}

func NewMain_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_propContext {
	var p = new(Main_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_main_prop

	return p
}

func (s *Main_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_propContext) NS() antlr.TerminalNode {
	return s.GetToken(aspidaParserNS, 0)
}

func (s *Main_propContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *Main_propContext) Connection() IConnectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectionContext)
}

func (s *Main_propContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *Main_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Main_propContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterMain_prop(s)
	}
}

func (s *Main_propContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitMain_prop(s)
	}
}

func (s *Main_propContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitMain_prop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Main_prop() (localctx IMain_propContext) {
	localctx = NewMain_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, aspidaParserRULE_main_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(102)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__3, aspidaParserT__4:
		{
			p.SetState(99)
			p.Name()
		}

	case aspidaParserT__5, aspidaParserT__6:
		{
			p.SetState(100)
			p.Connection()
		}

	case aspidaParserT__7, aspidaParserT__8:
		{
			p.SetState(101)
			p.Description()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(104)
		p.Match(aspidaParserNS)
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) STRING() antlr.TerminalNode {
	return s.GetToken(aspidaParserSTRING, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, aspidaParserRULE_name)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(112)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(106)
			p.Match(aspidaParserT__3)
		}
		{
			p.SetState(107)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(108)
			p.Match(aspidaParserSTRING)
		}

	case aspidaParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(109)
			p.Match(aspidaParserT__4)
		}
		{
			p.SetState(110)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(111)
			p.Match(aspidaParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConnectionContext is an interface to support dynamic dispatch.
type IConnectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnectionContext differentiates from other interfaces.
	IsConnectionContext()
}

type ConnectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnectionContext() *ConnectionContext {
	var p = new(ConnectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_connection
	return p
}

func (*ConnectionContext) IsConnectionContext() {}

func NewConnectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectionContext {
	var p = new(ConnectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_connection

	return p
}

func (s *ConnectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectionContext) LOCAL_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserLOCAL_KW, 0)
}

func (s *ConnectionContext) SSH_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserSSH_KW, 0)
}

func (s *ConnectionContext) SSHPASS_KW() antlr.TerminalNode {
	return s.GetToken(aspidaParserSSHPASS_KW, 0)
}

func (s *ConnectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConnectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterConnection(s)
	}
}

func (s *ConnectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitConnection(s)
	}
}

func (s *ConnectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitConnection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Connection() (localctx IConnectionContext) {
	localctx = NewConnectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, aspidaParserRULE_connection)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(120)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(114)
			p.Match(aspidaParserT__5)
		}
		{
			p.SetState(115)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(116)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(aspidaParserLOCAL_KW-41))|(1<<(aspidaParserSSH_KW-41))|(1<<(aspidaParserSSHPASS_KW-41)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case aspidaParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(117)
			p.Match(aspidaParserT__6)
		}
		{
			p.SetState(118)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(119)
			_la = p.GetTokenStream().LA(1)

			if !(((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(aspidaParserLOCAL_KW-41))|(1<<(aspidaParserSSH_KW-41))|(1<<(aspidaParserSSHPASS_KW-41)))) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) STRING() antlr.TerminalNode {
	return s.GetToken(aspidaParserSTRING, 0)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, aspidaParserRULE_description)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(128)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(122)
			p.Match(aspidaParserT__7)
		}
		{
			p.SetState(123)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(124)
			p.Match(aspidaParserSTRING)
		}

	case aspidaParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(125)
			p.Match(aspidaParserT__8)
		}
		{
			p.SetState(126)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(127)
			p.Match(aspidaParserSTRING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHosts_listContext is an interface to support dynamic dispatch.
type IHosts_listContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHosts_listContext differentiates from other interfaces.
	IsHosts_listContext()
}

type Hosts_listContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHosts_listContext() *Hosts_listContext {
	var p = new(Hosts_listContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_hosts_list
	return p
}

func (*Hosts_listContext) IsHosts_listContext() {}

func NewHosts_listContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Hosts_listContext {
	var p = new(Hosts_listContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_hosts_list

	return p
}

func (s *Hosts_listContext) GetParser() antlr.Parser { return s.parser }

func (s *Hosts_listContext) AllHost() []IHostContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHostContext)(nil)).Elem())
	var tst = make([]IHostContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHostContext)
		}
	}

	return tst
}

func (s *Hosts_listContext) Host(i int) IHostContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHostContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHostContext)
}

func (s *Hosts_listContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Hosts_listContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Hosts_listContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterHosts_list(s)
	}
}

func (s *Hosts_listContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitHosts_list(s)
	}
}

func (s *Hosts_listContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitHosts_list(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Hosts_list() (localctx IHosts_listContext) {
	localctx = NewHosts_listContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, aspidaParserRULE_hosts_list)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(143)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(aspidaParserT__9)
		}
		{
			p.SetState(131)
			p.Host()
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == aspidaParserT__10 {
			{
				p.SetState(132)
				p.Match(aspidaParserT__10)
			}
			{
				p.SetState(133)
				p.Host()
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(139)
			p.Match(aspidaParserT__11)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Match(aspidaParserT__9)
		}
		{
			p.SetState(142)
			p.Match(aspidaParserT__11)
		}

	}

	return localctx
}

// IHostContext is an interface to support dynamic dispatch.
type IHostContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHostContext differentiates from other interfaces.
	IsHostContext()
}

type HostContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHostContext() *HostContext {
	var p = new(HostContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_host
	return p
}

func (*HostContext) IsHostContext() {}

func NewHostContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostContext {
	var p = new(HostContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_host

	return p
}

func (s *HostContext) GetParser() antlr.Parser { return s.parser }

func (s *HostContext) Ip_v4() IIp_v4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v4Context)
}

func (s *HostContext) Ip_v6() IIp_v6Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v6Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v6Context)
}

func (s *HostContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HostContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterHost(s)
	}
}

func (s *HostContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitHost(s)
	}
}

func (s *HostContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitHost(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Host() (localctx IHostContext) {
	localctx = NewHostContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, aspidaParserRULE_host)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserSINGLE_IP, aspidaParserIP_RANGE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.Ip_v4()
		}

	case aspidaParserT__23, aspidaParserDIGIT, aspidaParserHEX_CHAR:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.Ip_v6()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITasks_contentContext is an interface to support dynamic dispatch.
type ITasks_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasks_contentContext differentiates from other interfaces.
	IsTasks_contentContext()
}

type Tasks_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasks_contentContext() *Tasks_contentContext {
	var p = new(Tasks_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_tasks_content
	return p
}

func (*Tasks_contentContext) IsTasks_contentContext() {}

func NewTasks_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tasks_contentContext {
	var p = new(Tasks_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_tasks_content

	return p
}

func (s *Tasks_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Tasks_contentContext) AllTasks_prop() []ITasks_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITasks_propContext)(nil)).Elem())
	var tst = make([]ITasks_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITasks_propContext)
		}
	}

	return tst
}

func (s *Tasks_contentContext) Tasks_prop(i int) ITasks_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITasks_propContext)
}

func (s *Tasks_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tasks_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tasks_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterTasks_content(s)
	}
}

func (s *Tasks_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitTasks_content(s)
	}
}

func (s *Tasks_contentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTasks_content(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Tasks_content() (localctx ITasks_contentContext) {
	localctx = NewTasks_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, aspidaParserRULE_tasks_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Tasks_prop()
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<aspidaParserT__12)|(1<<aspidaParserT__13)|(1<<aspidaParserT__14)|(1<<aspidaParserT__15)|(1<<aspidaParserT__16)|(1<<aspidaParserT__17)|(1<<aspidaParserT__18)|(1<<aspidaParserT__19))) != 0 {
		{
			p.SetState(150)
			p.Tasks_prop()
		}

		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITasks_propContext is an interface to support dynamic dispatch.
type ITasks_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTasks_propContext differentiates from other interfaces.
	IsTasks_propContext()
}

type Tasks_propContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTasks_propContext() *Tasks_propContext {
	var p = new(Tasks_propContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_tasks_prop
	return p
}

func (*Tasks_propContext) IsTasks_propContext() {}

func NewTasks_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tasks_propContext {
	var p = new(Tasks_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_tasks_prop

	return p
}

func (s *Tasks_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Tasks_propContext) NS() antlr.TerminalNode {
	return s.GetToken(aspidaParserNS, 0)
}

func (s *Tasks_propContext) Sections() ISectionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionsContext)
}

func (s *Tasks_propContext) Points() IPointsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointsContext)
}

func (s *Tasks_propContext) Controls() IControlsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlsContext)
}

func (s *Tasks_propContext) Exclusions() IExclusionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExclusionsContext)
}

func (s *Tasks_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tasks_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tasks_propContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterTasks_prop(s)
	}
}

func (s *Tasks_propContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitTasks_prop(s)
	}
}

func (s *Tasks_propContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitTasks_prop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Tasks_prop() (localctx ITasks_propContext) {
	localctx = NewTasks_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, aspidaParserRULE_tasks_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(160)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__12, aspidaParserT__13:
		{
			p.SetState(156)
			p.Sections()
		}

	case aspidaParserT__14, aspidaParserT__15:
		{
			p.SetState(157)
			p.Points()
		}

	case aspidaParserT__16, aspidaParserT__17:
		{
			p.SetState(158)
			p.Controls()
		}

	case aspidaParserT__18, aspidaParserT__19:
		{
			p.SetState(159)
			p.Exclusions()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	{
		p.SetState(162)
		p.Match(aspidaParserNS)
	}

	return localctx
}

// ISectionsContext is an interface to support dynamic dispatch.
type ISectionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSectionsContext differentiates from other interfaces.
	IsSectionsContext()
}

type SectionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySectionsContext() *SectionsContext {
	var p = new(SectionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_sections
	return p
}

func (*SectionsContext) IsSectionsContext() {}

func NewSectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionsContext {
	var p = new(SectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_sections

	return p
}

func (s *SectionsContext) GetParser() antlr.Parser { return s.parser }

func (s *SectionsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *SectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SectionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterSections(s)
	}
}

func (s *SectionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitSections(s)
	}
}

func (s *SectionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitSections(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Sections() (localctx ISectionsContext) {
	localctx = NewSectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, aspidaParserRULE_sections)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(170)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__12:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(164)
			p.Match(aspidaParserT__12)
		}
		{
			p.SetState(165)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(166)
			p.Str_array()
		}

	case aspidaParserT__13:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(167)
			p.Match(aspidaParserT__13)
		}
		{
			p.SetState(168)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(169)
			p.Str_array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPointsContext is an interface to support dynamic dispatch.
type IPointsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPointsContext differentiates from other interfaces.
	IsPointsContext()
}

type PointsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPointsContext() *PointsContext {
	var p = new(PointsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_points
	return p
}

func (*PointsContext) IsPointsContext() {}

func NewPointsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointsContext {
	var p = new(PointsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_points

	return p
}

func (s *PointsContext) GetParser() antlr.Parser { return s.parser }

func (s *PointsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *PointsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PointsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterPoints(s)
	}
}

func (s *PointsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitPoints(s)
	}
}

func (s *PointsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitPoints(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Points() (localctx IPointsContext) {
	localctx = NewPointsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, aspidaParserRULE_points)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(178)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(172)
			p.Match(aspidaParserT__14)
		}
		{
			p.SetState(173)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(174)
			p.Str_array()
		}

	case aspidaParserT__15:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(aspidaParserT__15)
		}
		{
			p.SetState(176)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(177)
			p.Str_array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IControlsContext is an interface to support dynamic dispatch.
type IControlsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsControlsContext differentiates from other interfaces.
	IsControlsContext()
}

type ControlsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyControlsContext() *ControlsContext {
	var p = new(ControlsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_controls
	return p
}

func (*ControlsContext) IsControlsContext() {}

func NewControlsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlsContext {
	var p = new(ControlsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_controls

	return p
}

func (s *ControlsContext) GetParser() antlr.Parser { return s.parser }

func (s *ControlsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *ControlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ControlsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterControls(s)
	}
}

func (s *ControlsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitControls(s)
	}
}

func (s *ControlsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitControls(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Controls() (localctx IControlsContext) {
	localctx = NewControlsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, aspidaParserRULE_controls)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(186)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(180)
			p.Match(aspidaParserT__16)
		}
		{
			p.SetState(181)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(182)
			p.Str_array()
		}

	case aspidaParserT__17:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(aspidaParserT__17)
		}
		{
			p.SetState(184)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(185)
			p.Str_array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExclusionsContext is an interface to support dynamic dispatch.
type IExclusionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExclusionsContext differentiates from other interfaces.
	IsExclusionsContext()
}

type ExclusionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExclusionsContext() *ExclusionsContext {
	var p = new(ExclusionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_exclusions
	return p
}

func (*ExclusionsContext) IsExclusionsContext() {}

func NewExclusionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusionsContext {
	var p = new(ExclusionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_exclusions

	return p
}

func (s *ExclusionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ExclusionsContext) Str_array() IStr_arrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStr_arrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStr_arrayContext)
}

func (s *ExclusionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExclusionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterExclusions(s)
	}
}

func (s *ExclusionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitExclusions(s)
	}
}

func (s *ExclusionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitExclusions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Exclusions() (localctx IExclusionsContext) {
	localctx = NewExclusionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, aspidaParserRULE_exclusions)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(194)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserT__18:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(188)
			p.Match(aspidaParserT__18)
		}
		{
			p.SetState(189)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(190)
			p.Str_array()
		}

	case aspidaParserT__19:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(191)
			p.Match(aspidaParserT__19)
		}
		{
			p.SetState(192)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(193)
			p.Str_array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IVars_contentContext is an interface to support dynamic dispatch.
type IVars_contentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_contentContext differentiates from other interfaces.
	IsVars_contentContext()
}

type Vars_contentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_contentContext() *Vars_contentContext {
	var p = new(Vars_contentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_vars_content
	return p
}

func (*Vars_contentContext) IsVars_contentContext() {}

func NewVars_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_contentContext {
	var p = new(Vars_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_vars_content

	return p
}

func (s *Vars_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_contentContext) AllVars_prop() []IVars_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVars_propContext)(nil)).Elem())
	var tst = make([]IVars_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVars_propContext)
		}
	}

	return tst
}

func (s *Vars_contentContext) Vars_prop(i int) IVars_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVars_propContext)
}

func (s *Vars_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_contentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterVars_content(s)
	}
}

func (s *Vars_contentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitVars_content(s)
	}
}

func (s *Vars_contentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVars_content(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Vars_content() (localctx IVars_contentContext) {
	localctx = NewVars_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, aspidaParserRULE_vars_content)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Vars_prop()
	}
	p.SetState(200)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == aspidaParserSTRING {
		{
			p.SetState(197)
			p.Vars_prop()
		}

		p.SetState(202)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IVars_propContext is an interface to support dynamic dispatch.
type IVars_propContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVars_propContext differentiates from other interfaces.
	IsVars_propContext()
}

type Vars_propContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVars_propContext() *Vars_propContext {
	var p = new(Vars_propContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_vars_prop
	return p
}

func (*Vars_propContext) IsVars_propContext() {}

func NewVars_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_propContext {
	var p = new(Vars_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_vars_prop

	return p
}

func (s *Vars_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_propContext) STRING() antlr.TerminalNode {
	return s.GetToken(aspidaParserSTRING, 0)
}

func (s *Vars_propContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *Vars_propContext) NS() antlr.TerminalNode {
	return s.GetToken(aspidaParserNS, 0)
}

func (s *Vars_propContext) Vars_content() IVars_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_contentContext)
}

func (s *Vars_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Vars_propContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterVars_prop(s)
	}
}

func (s *Vars_propContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitVars_prop(s)
	}
}

func (s *Vars_propContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitVars_prop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Vars_prop() (localctx IVars_propContext) {
	localctx = NewVars_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, aspidaParserRULE_vars_prop)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(215)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(203)
			p.Match(aspidaParserSTRING)
		}
		{
			p.SetState(204)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(205)
			p.Value()
		}
		{
			p.SetState(206)
			p.Match(aspidaParserNS)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(aspidaParserSTRING)
		}
		{
			p.SetState(209)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(210)
			p.Match(aspidaParserT__1)
		}
		{
			p.SetState(211)
			p.Vars_content()
		}
		{
			p.SetState(212)
			p.Match(aspidaParserT__2)
		}
		{
			p.SetState(213)
			p.Match(aspidaParserNS)
		}

	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) STRING() antlr.TerminalNode {
	return s.GetToken(aspidaParserSTRING, 0)
}

func (s *ValueContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(aspidaParserNUMBER, 0)
}

func (s *ValueContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, aspidaParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(223)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserSTRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)
			p.Match(aspidaParserSTRING)
		}

	case aspidaParserNUMBER:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(aspidaParserNUMBER)
		}

	case aspidaParserT__20:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(219)
			p.Match(aspidaParserT__20)
		}

	case aspidaParserT__21:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(220)
			p.Match(aspidaParserT__21)
		}

	case aspidaParserT__22:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(221)
			p.Match(aspidaParserT__22)
		}

	case aspidaParserT__9:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(222)
			p.Array()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IStr_arrayContext is an interface to support dynamic dispatch.
type IStr_arrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStr_arrayContext differentiates from other interfaces.
	IsStr_arrayContext()
}

type Str_arrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStr_arrayContext() *Str_arrayContext {
	var p = new(Str_arrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_str_array
	return p
}

func (*Str_arrayContext) IsStr_arrayContext() {}

func NewStr_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Str_arrayContext {
	var p = new(Str_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_str_array

	return p
}

func (s *Str_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Str_arrayContext) AllCadena() []ICadenaContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICadenaContext)(nil)).Elem())
	var tst = make([]ICadenaContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICadenaContext)
		}
	}

	return tst
}

func (s *Str_arrayContext) Cadena(i int) ICadenaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICadenaContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICadenaContext)
}

func (s *Str_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Str_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Str_arrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterStr_array(s)
	}
}

func (s *Str_arrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitStr_array(s)
	}
}

func (s *Str_arrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitStr_array(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Str_array() (localctx IStr_arrayContext) {
	localctx = NewStr_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, aspidaParserRULE_str_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(225)
			p.Match(aspidaParserT__9)
		}
		{
			p.SetState(226)
			p.Cadena()
		}
		p.SetState(231)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == aspidaParserT__10 {
			{
				p.SetState(227)
				p.Match(aspidaParserT__10)
			}
			{
				p.SetState(228)
				p.Cadena()
			}

			p.SetState(233)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(234)
			p.Match(aspidaParserT__11)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(236)
			p.Match(aspidaParserT__9)
		}
		{
			p.SetState(237)
			p.Match(aspidaParserT__11)
		}

	}

	return localctx
}

// ICadenaContext is an interface to support dynamic dispatch.
type ICadenaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsCadenaContext differentiates from other interfaces.
	IsCadenaContext()
}

type CadenaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCadenaContext() *CadenaContext {
	var p = new(CadenaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_cadena
	return p
}

func (*CadenaContext) IsCadenaContext() {}

func NewCadenaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CadenaContext {
	var p = new(CadenaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_cadena

	return p
}

func (s *CadenaContext) GetParser() antlr.Parser { return s.parser }

func (s *CadenaContext) STRING() antlr.TerminalNode {
	return s.GetToken(aspidaParserSTRING, 0)
}

func (s *CadenaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CadenaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CadenaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterCadena(s)
	}
}

func (s *CadenaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitCadena(s)
	}
}

func (s *CadenaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitCadena(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Cadena() (localctx ICadenaContext) {
	localctx = NewCadenaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, aspidaParserRULE_cadena)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(240)
		p.Match(aspidaParserSTRING)
	}

	return localctx
}

// IArrayContext is an interface to support dynamic dispatch.
type IArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArrayContext differentiates from other interfaces.
	IsArrayContext()
}

type ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayContext() *ArrayContext {
	var p = new(ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_array

	return p
}

func (s *ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ArrayContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterArray(s)
	}
}

func (s *ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitArray(s)
	}
}

func (s *ArrayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitArray(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, aspidaParserRULE_array)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(255)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(242)
			p.Match(aspidaParserT__9)
		}
		{
			p.SetState(243)
			p.Value()
		}
		p.SetState(248)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == aspidaParserT__10 {
			{
				p.SetState(244)
				p.Match(aspidaParserT__10)
			}
			{
				p.SetState(245)
				p.Value()
			}

			p.SetState(250)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(251)
			p.Match(aspidaParserT__11)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(253)
			p.Match(aspidaParserT__9)
		}
		{
			p.SetState(254)
			p.Match(aspidaParserT__11)
		}

	}

	return localctx
}

// IIp_v4Context is an interface to support dynamic dispatch.
type IIp_v4Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_v4Context differentiates from other interfaces.
	IsIp_v4Context()
}

type Ip_v4Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_v4Context() *Ip_v4Context {
	var p = new(Ip_v4Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_ip_v4
	return p
}

func (*Ip_v4Context) IsIp_v4Context() {}

func NewIp_v4Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_v4Context {
	var p = new(Ip_v4Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_ip_v4

	return p
}

func (s *Ip_v4Context) GetParser() antlr.Parser { return s.parser }

func (s *Ip_v4Context) Single_ip() ISingle_ipContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISingle_ipContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISingle_ipContext)
}

func (s *Ip_v4Context) Ip_range() IIp_rangeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_rangeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_rangeContext)
}

func (s *Ip_v4Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_v4Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_v4Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterIp_v4(s)
	}
}

func (s *Ip_v4Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitIp_v4(s)
	}
}

func (s *Ip_v4Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitIp_v4(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Ip_v4() (localctx IIp_v4Context) {
	localctx = NewIp_v4Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, aspidaParserRULE_ip_v4)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(259)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserSINGLE_IP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Single_ip()
		}

	case aspidaParserIP_RANGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Ip_range()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISingle_ipContext is an interface to support dynamic dispatch.
type ISingle_ipContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSingle_ipContext differentiates from other interfaces.
	IsSingle_ipContext()
}

type Single_ipContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySingle_ipContext() *Single_ipContext {
	var p = new(Single_ipContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_single_ip
	return p
}

func (*Single_ipContext) IsSingle_ipContext() {}

func NewSingle_ipContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Single_ipContext {
	var p = new(Single_ipContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_single_ip

	return p
}

func (s *Single_ipContext) GetParser() antlr.Parser { return s.parser }

func (s *Single_ipContext) SINGLE_IP() antlr.TerminalNode {
	return s.GetToken(aspidaParserSINGLE_IP, 0)
}

func (s *Single_ipContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Single_ipContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Single_ipContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterSingle_ip(s)
	}
}

func (s *Single_ipContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitSingle_ip(s)
	}
}

func (s *Single_ipContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitSingle_ip(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Single_ip() (localctx ISingle_ipContext) {
	localctx = NewSingle_ipContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, aspidaParserRULE_single_ip)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(261)
		p.Match(aspidaParserSINGLE_IP)
	}

	return localctx
}

// IIp_rangeContext is an interface to support dynamic dispatch.
type IIp_rangeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_rangeContext differentiates from other interfaces.
	IsIp_rangeContext()
}

type Ip_rangeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_rangeContext() *Ip_rangeContext {
	var p = new(Ip_rangeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_ip_range
	return p
}

func (*Ip_rangeContext) IsIp_rangeContext() {}

func NewIp_rangeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_rangeContext {
	var p = new(Ip_rangeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_ip_range

	return p
}

func (s *Ip_rangeContext) GetParser() antlr.Parser { return s.parser }

func (s *Ip_rangeContext) IP_RANGE() antlr.TerminalNode {
	return s.GetToken(aspidaParserIP_RANGE, 0)
}

func (s *Ip_rangeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_rangeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_rangeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterIp_range(s)
	}
}

func (s *Ip_rangeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitIp_range(s)
	}
}

func (s *Ip_rangeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitIp_range(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Ip_range() (localctx IIp_rangeContext) {
	localctx = NewIp_rangeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, aspidaParserRULE_ip_range)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(aspidaParserIP_RANGE)
	}

	return localctx
}

// IIp_v6Context is an interface to support dynamic dispatch.
type IIp_v6Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIp_v6Context differentiates from other interfaces.
	IsIp_v6Context()
}

type Ip_v6Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIp_v6Context() *Ip_v6Context {
	var p = new(Ip_v6Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_ip_v6
	return p
}

func (*Ip_v6Context) IsIp_v6Context() {}

func NewIp_v6Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ip_v6Context {
	var p = new(Ip_v6Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_ip_v6

	return p
}

func (s *Ip_v6Context) GetParser() antlr.Parser { return s.parser }

func (s *Ip_v6Context) AllH16() []IH16Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IH16Context)(nil)).Elem())
	var tst = make([]IH16Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IH16Context)
		}
	}

	return tst
}

func (s *Ip_v6Context) H16(i int) IH16Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IH16Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IH16Context)
}

func (s *Ip_v6Context) Ls32() ILs32Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILs32Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILs32Context)
}

func (s *Ip_v6Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ip_v6Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ip_v6Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterIp_v6(s)
	}
}

func (s *Ip_v6Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitIp_v6(s)
	}
}

func (s *Ip_v6Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitIp_v6(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Ip_v6() (localctx IIp_v6Context) {
	localctx = NewIp_v6Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, aspidaParserRULE_ip_v6)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(421)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(265)
			p.H16()
		}
		{
			p.SetState(266)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(267)
			p.H16()
		}
		{
			p.SetState(268)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(269)
			p.H16()
		}
		{
			p.SetState(270)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(271)
			p.H16()
		}
		{
			p.SetState(272)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(273)
			p.H16()
		}
		{
			p.SetState(274)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(275)
			p.H16()
		}
		{
			p.SetState(276)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(277)
			p.Ls32()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(279)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(280)
			p.H16()
		}
		{
			p.SetState(281)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(282)
			p.H16()
		}
		{
			p.SetState(283)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(284)
			p.H16()
		}
		{
			p.SetState(285)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(286)
			p.H16()
		}
		{
			p.SetState(287)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(288)
			p.H16()
		}
		{
			p.SetState(289)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(290)
			p.Ls32()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		p.SetState(293)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR {
			{
				p.SetState(292)
				p.H16()
			}

		}
		{
			p.SetState(295)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(296)
			p.H16()
		}
		{
			p.SetState(297)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(298)
			p.H16()
		}
		{
			p.SetState(299)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(300)
			p.H16()
		}
		{
			p.SetState(301)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(302)
			p.H16()
		}
		{
			p.SetState(303)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(304)
			p.Ls32()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		p.SetState(312)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR {
			p.SetState(309)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
				{
					p.SetState(306)
					p.H16()
				}
				{
					p.SetState(307)
					p.Match(aspidaParserT__0)
				}

			}
			{
				p.SetState(311)
				p.H16()
			}

		}
		{
			p.SetState(314)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(315)
			p.H16()
		}
		{
			p.SetState(316)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(317)
			p.H16()
		}
		{
			p.SetState(318)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(319)
			p.H16()
		}
		{
			p.SetState(320)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(321)
			p.Ls32()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		p.SetState(334)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR {
			p.SetState(331)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 28, p.GetParserRuleContext()) == 1 {
				p.SetState(326)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) == 1 {
					{
						p.SetState(323)
						p.H16()
					}
					{
						p.SetState(324)
						p.Match(aspidaParserT__0)
					}

				}
				{
					p.SetState(328)
					p.H16()
				}
				{
					p.SetState(329)
					p.Match(aspidaParserT__0)
				}

			}
			{
				p.SetState(333)
				p.H16()
			}

		}
		{
			p.SetState(336)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(337)
			p.H16()
		}
		{
			p.SetState(338)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(339)
			p.H16()
		}
		{
			p.SetState(340)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(341)
			p.Ls32()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		p.SetState(359)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR {
			p.SetState(356)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) == 1 {
				p.SetState(351)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext()) == 1 {
					p.SetState(346)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
						{
							p.SetState(343)
							p.H16()
						}
						{
							p.SetState(344)
							p.Match(aspidaParserT__0)
						}

					}
					{
						p.SetState(348)
						p.H16()
					}
					{
						p.SetState(349)
						p.Match(aspidaParserT__0)
					}

				}
				{
					p.SetState(353)
					p.H16()
				}
				{
					p.SetState(354)
					p.Match(aspidaParserT__0)
				}

			}
			{
				p.SetState(358)
				p.H16()
			}

		}
		{
			p.SetState(361)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(362)
			p.H16()
		}
		{
			p.SetState(363)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(364)
			p.Ls32()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR {
			p.SetState(384)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
				p.SetState(379)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 36, p.GetParserRuleContext()) == 1 {
					p.SetState(374)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) == 1 {
						p.SetState(369)
						p.GetErrorHandler().Sync(p)

						if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) == 1 {
							{
								p.SetState(366)
								p.H16()
							}
							{
								p.SetState(367)
								p.Match(aspidaParserT__0)
							}

						}
						{
							p.SetState(371)
							p.H16()
						}
						{
							p.SetState(372)
							p.Match(aspidaParserT__0)
						}

					}
					{
						p.SetState(376)
						p.H16()
					}
					{
						p.SetState(377)
						p.Match(aspidaParserT__0)
					}

				}
				{
					p.SetState(381)
					p.H16()
				}
				{
					p.SetState(382)
					p.Match(aspidaParserT__0)
				}

			}
			{
				p.SetState(386)
				p.H16()
			}

		}
		{
			p.SetState(389)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(390)
			p.Ls32()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		p.SetState(417)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR {
			p.SetState(414)
			p.GetErrorHandler().Sync(p)

			if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) == 1 {
				p.SetState(409)
				p.GetErrorHandler().Sync(p)

				if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) == 1 {
					p.SetState(404)
					p.GetErrorHandler().Sync(p)

					if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
						p.SetState(399)
						p.GetErrorHandler().Sync(p)

						if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) == 1 {
							p.SetState(394)
							p.GetErrorHandler().Sync(p)

							if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) == 1 {
								{
									p.SetState(391)
									p.H16()
								}
								{
									p.SetState(392)
									p.Match(aspidaParserT__0)
								}

							}
							{
								p.SetState(396)
								p.H16()
							}
							{
								p.SetState(397)
								p.Match(aspidaParserT__0)
							}

						}
						{
							p.SetState(401)
							p.H16()
						}
						{
							p.SetState(402)
							p.Match(aspidaParserT__0)
						}

					}
					{
						p.SetState(406)
						p.H16()
					}
					{
						p.SetState(407)
						p.Match(aspidaParserT__0)
					}

				}
				{
					p.SetState(411)
					p.H16()
				}
				{
					p.SetState(412)
					p.Match(aspidaParserT__0)
				}

			}
			{
				p.SetState(416)
				p.H16()
			}

		}
		{
			p.SetState(419)
			p.Match(aspidaParserT__23)
		}
		{
			p.SetState(420)
			p.H16()
		}

	}

	return localctx
}

// IH16Context is an interface to support dynamic dispatch.
type IH16Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsH16Context differentiates from other interfaces.
	IsH16Context()
}

type H16Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyH16Context() *H16Context {
	var p = new(H16Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_h16
	return p
}

func (*H16Context) IsH16Context() {}

func NewH16Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *H16Context {
	var p = new(H16Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_h16

	return p
}

func (s *H16Context) GetParser() antlr.Parser { return s.parser }

func (s *H16Context) AllHexdigit() []IHexdigitContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IHexdigitContext)(nil)).Elem())
	var tst = make([]IHexdigitContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IHexdigitContext)
		}
	}

	return tst
}

func (s *H16Context) Hexdigit(i int) IHexdigitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHexdigitContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IHexdigitContext)
}

func (s *H16Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *H16Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *H16Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterH16(s)
	}
}

func (s *H16Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitH16(s)
	}
}

func (s *H16Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitH16(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) H16() (localctx IH16Context) {
	localctx = NewH16Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, aspidaParserRULE_h16)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(436)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(423)
			p.Hexdigit()
		}
		{
			p.SetState(424)
			p.Hexdigit()
		}
		{
			p.SetState(425)
			p.Hexdigit()
		}
		{
			p.SetState(426)
			p.Hexdigit()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(428)
			p.Hexdigit()
		}
		{
			p.SetState(429)
			p.Hexdigit()
		}
		{
			p.SetState(430)
			p.Hexdigit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(432)
			p.Hexdigit()
		}
		{
			p.SetState(433)
			p.Hexdigit()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(435)
			p.Hexdigit()
		}

	}

	return localctx
}

// ILs32Context is an interface to support dynamic dispatch.
type ILs32Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLs32Context differentiates from other interfaces.
	IsLs32Context()
}

type Ls32Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLs32Context() *Ls32Context {
	var p = new(Ls32Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_ls32
	return p
}

func (*Ls32Context) IsLs32Context() {}

func NewLs32Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Ls32Context {
	var p = new(Ls32Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_ls32

	return p
}

func (s *Ls32Context) GetParser() antlr.Parser { return s.parser }

func (s *Ls32Context) AllH16() []IH16Context {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IH16Context)(nil)).Elem())
	var tst = make([]IH16Context, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IH16Context)
		}
	}

	return tst
}

func (s *Ls32Context) H16(i int) IH16Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IH16Context)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IH16Context)
}

func (s *Ls32Context) Ip_v4() IIp_v4Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIp_v4Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIp_v4Context)
}

func (s *Ls32Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Ls32Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Ls32Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterLs32(s)
	}
}

func (s *Ls32Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitLs32(s)
	}
}

func (s *Ls32Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitLs32(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Ls32() (localctx ILs32Context) {
	localctx = NewLs32Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, aspidaParserRULE_ls32)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(443)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case aspidaParserDIGIT, aspidaParserHEX_CHAR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(438)
			p.H16()
		}
		{
			p.SetState(439)
			p.Match(aspidaParserT__0)
		}
		{
			p.SetState(440)
			p.H16()
		}

	case aspidaParserSINGLE_IP, aspidaParserIP_RANGE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(442)
			p.Ip_v4()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IHexdigitContext is an interface to support dynamic dispatch.
type IHexdigitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsHexdigitContext differentiates from other interfaces.
	IsHexdigitContext()
}

type HexdigitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHexdigitContext() *HexdigitContext {
	var p = new(HexdigitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = aspidaParserRULE_hexdigit
	return p
}

func (*HexdigitContext) IsHexdigitContext() {}

func NewHexdigitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HexdigitContext {
	var p = new(HexdigitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = aspidaParserRULE_hexdigit

	return p
}

func (s *HexdigitContext) GetParser() antlr.Parser { return s.parser }

func (s *HexdigitContext) DIGIT() antlr.TerminalNode {
	return s.GetToken(aspidaParserDIGIT, 0)
}

func (s *HexdigitContext) HEX_CHAR() antlr.TerminalNode {
	return s.GetToken(aspidaParserHEX_CHAR, 0)
}

func (s *HexdigitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HexdigitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HexdigitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.EnterHexdigit(s)
	}
}

func (s *HexdigitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(aspidaListener); ok {
		listenerT.ExitHexdigit(s)
	}
}

func (s *HexdigitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case AspidaVisitor:
		return t.VisitHexdigit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *aspidaParser) Hexdigit() (localctx IHexdigitContext) {
	localctx = NewHexdigitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, aspidaParserRULE_hexdigit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(445)
		_la = p.GetTokenStream().LA(1)

		if !(_la == aspidaParserDIGIT || _la == aspidaParserHEX_CHAR) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}
