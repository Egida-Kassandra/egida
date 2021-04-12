// Code generated from d:\OneDrive\Universidad de Oviedo\Dr. & Dra F - General\TFMs\Egida\Code\egida\Aspida.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // Aspida

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 50, 289,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 61, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 7, 7, 88, 10, 7, 12, 7, 14,
	7, 91, 11, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5,
	8, 102, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 110, 10, 9, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 118, 10, 10, 3, 11, 3, 11,
	5, 11, 122, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 5, 12, 130,
	10, 12, 3, 13, 3, 13, 7, 13, 134, 10, 13, 12, 13, 14, 13, 137, 11, 13,
	3, 13, 3, 13, 7, 13, 141, 10, 13, 12, 13, 14, 13, 144, 11, 13, 3, 13, 3,
	13, 5, 13, 148, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 5, 14, 154, 10, 14,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5,
	15, 166, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 16, 5, 16, 178, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 190, 10, 17, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 202, 10, 18, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20,
	3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 7, 22, 224,
	10, 22, 12, 22, 14, 22, 227, 11, 22, 3, 23, 3, 23, 7, 23, 231, 10, 23,
	12, 23, 14, 23, 234, 11, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 5, 24, 248, 10, 24, 3, 25, 3,
	25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 258, 10, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 7, 27, 264, 10, 27, 12, 27, 14, 27, 267, 11, 27, 3,
	27, 3, 27, 3, 27, 5, 27, 272, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 7, 28,
	278, 10, 28, 12, 28, 14, 28, 281, 11, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5,
	28, 287, 10, 28, 3, 28, 2, 2, 29, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 2, 3, 3,
	2, 20, 25, 2, 291, 2, 56, 3, 2, 2, 2, 4, 62, 3, 2, 2, 2, 6, 68, 3, 2, 2,
	2, 8, 73, 3, 2, 2, 2, 10, 79, 3, 2, 2, 2, 12, 85, 3, 2, 2, 2, 14, 101,
	3, 2, 2, 2, 16, 109, 3, 2, 2, 2, 18, 117, 3, 2, 2, 2, 20, 121, 3, 2, 2,
	2, 22, 129, 3, 2, 2, 2, 24, 147, 3, 2, 2, 2, 26, 153, 3, 2, 2, 2, 28, 165,
	3, 2, 2, 2, 30, 177, 3, 2, 2, 2, 32, 189, 3, 2, 2, 2, 34, 201, 3, 2, 2,
	2, 36, 203, 3, 2, 2, 2, 38, 209, 3, 2, 2, 2, 40, 215, 3, 2, 2, 2, 42, 220,
	3, 2, 2, 2, 44, 228, 3, 2, 2, 2, 46, 247, 3, 2, 2, 2, 48, 249, 3, 2, 2,
	2, 50, 257, 3, 2, 2, 2, 52, 271, 3, 2, 2, 2, 54, 286, 3, 2, 2, 2, 56, 57,
	5, 4, 3, 2, 57, 58, 5, 6, 4, 2, 58, 60, 5, 8, 5, 2, 59, 61, 5, 10, 6, 2,
	60, 59, 3, 2, 2, 2, 60, 61, 3, 2, 2, 2, 61, 3, 3, 2, 2, 2, 62, 63, 7, 38,
	2, 2, 63, 64, 7, 3, 2, 2, 64, 65, 7, 4, 2, 2, 65, 66, 5, 12, 7, 2, 66,
	67, 7, 5, 2, 2, 67, 5, 3, 2, 2, 2, 68, 69, 7, 39, 2, 2, 69, 70, 7, 3, 2,
	2, 70, 71, 7, 35, 2, 2, 71, 72, 7, 37, 2, 2, 72, 7, 3, 2, 2, 2, 73, 74,
	7, 40, 2, 2, 74, 75, 7, 3, 2, 2, 75, 76, 7, 4, 2, 2, 76, 77, 5, 24, 13,
	2, 77, 78, 7, 5, 2, 2, 78, 9, 3, 2, 2, 2, 79, 80, 7, 41, 2, 2, 80, 81,
	7, 3, 2, 2, 81, 82, 7, 4, 2, 2, 82, 83, 5, 44, 23, 2, 83, 84, 7, 5, 2,
	2, 84, 11, 3, 2, 2, 2, 85, 89, 5, 14, 8, 2, 86, 88, 5, 14, 8, 2, 87, 86,
	3, 2, 2, 2, 88, 91, 3, 2, 2, 2, 89, 87, 3, 2, 2, 2, 89, 90, 3, 2, 2, 2,
	90, 13, 3, 2, 2, 2, 91, 89, 3, 2, 2, 2, 92, 93, 5, 16, 9, 2, 93, 94, 7,
	37, 2, 2, 94, 102, 3, 2, 2, 2, 95, 96, 5, 18, 10, 2, 96, 97, 7, 37, 2,
	2, 97, 102, 3, 2, 2, 2, 98, 99, 5, 22, 12, 2, 99, 100, 7, 37, 2, 2, 100,
	102, 3, 2, 2, 2, 101, 92, 3, 2, 2, 2, 101, 95, 3, 2, 2, 2, 101, 98, 3,
	2, 2, 2, 102, 15, 3, 2, 2, 2, 103, 104, 7, 6, 2, 2, 104, 105, 7, 3, 2,
	2, 105, 110, 5, 50, 26, 2, 106, 107, 7, 7, 2, 2, 107, 108, 7, 3, 2, 2,
	108, 110, 5, 50, 26, 2, 109, 103, 3, 2, 2, 2, 109, 106, 3, 2, 2, 2, 110,
	17, 3, 2, 2, 2, 111, 112, 7, 8, 2, 2, 112, 113, 7, 3, 2, 2, 113, 118, 5,
	20, 11, 2, 114, 115, 7, 9, 2, 2, 115, 116, 7, 3, 2, 2, 116, 118, 5, 20,
	11, 2, 117, 111, 3, 2, 2, 2, 117, 114, 3, 2, 2, 2, 118, 19, 3, 2, 2, 2,
	119, 122, 7, 42, 2, 2, 120, 122, 7, 43, 2, 2, 121, 119, 3, 2, 2, 2, 121,
	120, 3, 2, 2, 2, 122, 21, 3, 2, 2, 2, 123, 124, 7, 10, 2, 2, 124, 125,
	7, 3, 2, 2, 125, 130, 5, 50, 26, 2, 126, 127, 7, 11, 2, 2, 127, 128, 7,
	3, 2, 2, 128, 130, 5, 50, 26, 2, 129, 123, 3, 2, 2, 2, 129, 126, 3, 2,
	2, 2, 130, 23, 3, 2, 2, 2, 131, 135, 5, 26, 14, 2, 132, 134, 5, 26, 14,
	2, 133, 132, 3, 2, 2, 2, 134, 137, 3, 2, 2, 2, 135, 133, 3, 2, 2, 2, 135,
	136, 3, 2, 2, 2, 136, 148, 3, 2, 2, 2, 137, 135, 3, 2, 2, 2, 138, 142,
	5, 36, 19, 2, 139, 141, 5, 38, 20, 2, 140, 139, 3, 2, 2, 2, 141, 144, 3,
	2, 2, 2, 142, 140, 3, 2, 2, 2, 142, 143, 3, 2, 2, 2, 143, 145, 3, 2, 2,
	2, 144, 142, 3, 2, 2, 2, 145, 146, 5, 40, 21, 2, 146, 148, 3, 2, 2, 2,
	147, 131, 3, 2, 2, 2, 147, 138, 3, 2, 2, 2, 148, 25, 3, 2, 2, 2, 149, 154,
	5, 28, 15, 2, 150, 154, 5, 30, 16, 2, 151, 154, 5, 32, 17, 2, 152, 154,
	5, 34, 18, 2, 153, 149, 3, 2, 2, 2, 153, 150, 3, 2, 2, 2, 153, 151, 3,
	2, 2, 2, 153, 152, 3, 2, 2, 2, 154, 27, 3, 2, 2, 2, 155, 156, 7, 12, 2,
	2, 156, 157, 7, 3, 2, 2, 157, 158, 5, 52, 27, 2, 158, 159, 7, 37, 2, 2,
	159, 166, 3, 2, 2, 2, 160, 161, 7, 13, 2, 2, 161, 162, 7, 3, 2, 2, 162,
	163, 5, 52, 27, 2, 163, 164, 7, 37, 2, 2, 164, 166, 3, 2, 2, 2, 165, 155,
	3, 2, 2, 2, 165, 160, 3, 2, 2, 2, 166, 29, 3, 2, 2, 2, 167, 168, 7, 14,
	2, 2, 168, 169, 7, 3, 2, 2, 169, 170, 5, 52, 27, 2, 170, 171, 7, 37, 2,
	2, 171, 178, 3, 2, 2, 2, 172, 173, 7, 15, 2, 2, 173, 174, 7, 3, 2, 2, 174,
	175, 5, 52, 27, 2, 175, 176, 7, 37, 2, 2, 176, 178, 3, 2, 2, 2, 177, 167,
	3, 2, 2, 2, 177, 172, 3, 2, 2, 2, 178, 31, 3, 2, 2, 2, 179, 180, 7, 16,
	2, 2, 180, 181, 7, 3, 2, 2, 181, 182, 5, 52, 27, 2, 182, 183, 7, 37, 2,
	2, 183, 190, 3, 2, 2, 2, 184, 185, 7, 17, 2, 2, 185, 186, 7, 3, 2, 2, 186,
	187, 5, 52, 27, 2, 187, 188, 7, 37, 2, 2, 188, 190, 3, 2, 2, 2, 189, 179,
	3, 2, 2, 2, 189, 184, 3, 2, 2, 2, 190, 33, 3, 2, 2, 2, 191, 192, 7, 18,
	2, 2, 192, 193, 7, 3, 2, 2, 193, 194, 5, 52, 27, 2, 194, 195, 7, 37, 2,
	2, 195, 202, 3, 2, 2, 2, 196, 197, 7, 19, 2, 2, 197, 198, 7, 3, 2, 2, 198,
	199, 5, 52, 27, 2, 199, 200, 7, 37, 2, 2, 200, 202, 3, 2, 2, 2, 201, 191,
	3, 2, 2, 2, 201, 196, 3, 2, 2, 2, 202, 35, 3, 2, 2, 2, 203, 204, 7, 44,
	2, 2, 204, 205, 5, 42, 22, 2, 205, 206, 7, 4, 2, 2, 206, 207, 5, 24, 13,
	2, 207, 208, 7, 5, 2, 2, 208, 37, 3, 2, 2, 2, 209, 210, 7, 45, 2, 2, 210,
	211, 5, 42, 22, 2, 211, 212, 7, 4, 2, 2, 212, 213, 5, 24, 13, 2, 213, 214,
	7, 5, 2, 2, 214, 39, 3, 2, 2, 2, 215, 216, 7, 46, 2, 2, 216, 217, 7, 4,
	2, 2, 217, 218, 5, 24, 13, 2, 218, 219, 7, 5, 2, 2, 219, 41, 3, 2, 2, 2,
	220, 221, 5, 50, 26, 2, 221, 225, 5, 48, 25, 2, 222, 224, 5, 50, 26, 2,
	223, 222, 3, 2, 2, 2, 224, 227, 3, 2, 2, 2, 225, 223, 3, 2, 2, 2, 225,
	226, 3, 2, 2, 2, 226, 43, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 232, 5,
	46, 24, 2, 229, 231, 5, 46, 24, 2, 230, 229, 3, 2, 2, 2, 231, 234, 3, 2,
	2, 2, 232, 230, 3, 2, 2, 2, 232, 233, 3, 2, 2, 2, 233, 45, 3, 2, 2, 2,
	234, 232, 3, 2, 2, 2, 235, 236, 7, 35, 2, 2, 236, 237, 7, 3, 2, 2, 237,
	238, 5, 50, 26, 2, 238, 239, 7, 37, 2, 2, 239, 248, 3, 2, 2, 2, 240, 241,
	7, 35, 2, 2, 241, 242, 7, 3, 2, 2, 242, 243, 7, 4, 2, 2, 243, 244, 5, 44,
	23, 2, 244, 245, 7, 5, 2, 2, 245, 246, 7, 37, 2, 2, 246, 248, 3, 2, 2,
	2, 247, 235, 3, 2, 2, 2, 247, 240, 3, 2, 2, 2, 248, 47, 3, 2, 2, 2, 249,
	250, 9, 2, 2, 2, 250, 49, 3, 2, 2, 2, 251, 258, 7, 35, 2, 2, 252, 258,
	7, 36, 2, 2, 253, 258, 7, 26, 2, 2, 254, 258, 7, 27, 2, 2, 255, 258, 7,
	28, 2, 2, 256, 258, 5, 54, 28, 2, 257, 251, 3, 2, 2, 2, 257, 252, 3, 2,
	2, 2, 257, 253, 3, 2, 2, 2, 257, 254, 3, 2, 2, 2, 257, 255, 3, 2, 2, 2,
	257, 256, 3, 2, 2, 2, 258, 51, 3, 2, 2, 2, 259, 260, 7, 29, 2, 2, 260,
	265, 7, 35, 2, 2, 261, 262, 7, 30, 2, 2, 262, 264, 7, 35, 2, 2, 263, 261,
	3, 2, 2, 2, 264, 267, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265, 266, 3, 2,
	2, 2, 266, 268, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 268, 272, 7, 31, 2, 2,
	269, 270, 7, 29, 2, 2, 270, 272, 7, 31, 2, 2, 271, 259, 3, 2, 2, 2, 271,
	269, 3, 2, 2, 2, 272, 53, 3, 2, 2, 2, 273, 274, 7, 29, 2, 2, 274, 279,
	5, 50, 26, 2, 275, 276, 7, 30, 2, 2, 276, 278, 5, 50, 26, 2, 277, 275,
	3, 2, 2, 2, 278, 281, 3, 2, 2, 2, 279, 277, 3, 2, 2, 2, 279, 280, 3, 2,
	2, 2, 280, 282, 3, 2, 2, 2, 281, 279, 3, 2, 2, 2, 282, 283, 7, 31, 2, 2,
	283, 287, 3, 2, 2, 2, 284, 285, 7, 29, 2, 2, 285, 287, 7, 31, 2, 2, 286,
	273, 3, 2, 2, 2, 286, 284, 3, 2, 2, 2, 287, 55, 3, 2, 2, 2, 25, 60, 89,
	101, 109, 117, 121, 129, 135, 142, 147, 153, 165, 177, 189, 201, 225, 232,
	247, 257, 265, 271, 279, 286,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "':'", "'{'", "'}'", "'name'", "'NAME'", "'connection'", "'CONNECTION'",
	"'description'", "'DESCRIPTION'", "'sections'", "'SECTIONS'", "'points'",
	"'POINTS'", "'controls'", "'CONTROLS'", "'exclusions'", "'EXCLUSIONS'",
	"'<'", "'>'", "'=='", "'>='", "'<='", "'!='", "'true'", "'false'", "'null'",
	"'['", "','", "']'", "", "", "", "", "", "';'", "'MAIN'", "'HOST'", "'TASKS'",
	"'VARS'", "'LOCAL'", "'SSH'", "'IF'", "'ELIF'", "'ELSE'", "'OR'", "'AND'",
	"'NOT'", "'IS'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "COMMENT", "WHITESPACE",
	"NEWLINE", "STRING", "NUMBER", "NS", "MAIN_KW", "HOSTS_KW", "TASKS_KW",
	"VARS_KW", "LOCAL_KW", "SSH_KW", "IF", "ELIF", "ELSE", "OR", "AND", "NOT",
	"IS",
}

var ruleNames = []string{
	"program", "main", "hosts", "tasks", "variables", "main_content", "main_prop",
	"name", "connection", "connection_type", "description", "tasks_content",
	"tasks_prop", "sections", "points", "controls", "exclusions", "ifStat",
	"elifStat", "elseStat", "comparison", "vars_content", "vars_prop", "comp_op",
	"value", "str_array", "array",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type AspidaParser struct {
	*antlr.BaseParser
}

func NewAspidaParser(input antlr.TokenStream) *AspidaParser {
	this := new(AspidaParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Aspida.g4"

	return this
}

// AspidaParser tokens.
const (
	AspidaParserEOF        = antlr.TokenEOF
	AspidaParserT__0       = 1
	AspidaParserT__1       = 2
	AspidaParserT__2       = 3
	AspidaParserT__3       = 4
	AspidaParserT__4       = 5
	AspidaParserT__5       = 6
	AspidaParserT__6       = 7
	AspidaParserT__7       = 8
	AspidaParserT__8       = 9
	AspidaParserT__9       = 10
	AspidaParserT__10      = 11
	AspidaParserT__11      = 12
	AspidaParserT__12      = 13
	AspidaParserT__13      = 14
	AspidaParserT__14      = 15
	AspidaParserT__15      = 16
	AspidaParserT__16      = 17
	AspidaParserT__17      = 18
	AspidaParserT__18      = 19
	AspidaParserT__19      = 20
	AspidaParserT__20      = 21
	AspidaParserT__21      = 22
	AspidaParserT__22      = 23
	AspidaParserT__23      = 24
	AspidaParserT__24      = 25
	AspidaParserT__25      = 26
	AspidaParserT__26      = 27
	AspidaParserT__27      = 28
	AspidaParserT__28      = 29
	AspidaParserCOMMENT    = 30
	AspidaParserWHITESPACE = 31
	AspidaParserNEWLINE    = 32
	AspidaParserSTRING     = 33
	AspidaParserNUMBER     = 34
	AspidaParserNS         = 35
	AspidaParserMAIN_KW    = 36
	AspidaParserHOSTS_KW   = 37
	AspidaParserTASKS_KW   = 38
	AspidaParserVARS_KW    = 39
	AspidaParserLOCAL_KW   = 40
	AspidaParserSSH_KW     = 41
	AspidaParserIF         = 42
	AspidaParserELIF       = 43
	AspidaParserELSE       = 44
	AspidaParserOR         = 45
	AspidaParserAND        = 46
	AspidaParserNOT        = 47
	AspidaParserIS         = 48
)

// AspidaParser rules.
const (
	AspidaParserRULE_program         = 0
	AspidaParserRULE_main            = 1
	AspidaParserRULE_hosts           = 2
	AspidaParserRULE_tasks           = 3
	AspidaParserRULE_variables       = 4
	AspidaParserRULE_main_content    = 5
	AspidaParserRULE_main_prop       = 6
	AspidaParserRULE_name            = 7
	AspidaParserRULE_connection      = 8
	AspidaParserRULE_connection_type = 9
	AspidaParserRULE_description     = 10
	AspidaParserRULE_tasks_content   = 11
	AspidaParserRULE_tasks_prop      = 12
	AspidaParserRULE_sections        = 13
	AspidaParserRULE_points          = 14
	AspidaParserRULE_controls        = 15
	AspidaParserRULE_exclusions      = 16
	AspidaParserRULE_ifStat          = 17
	AspidaParserRULE_elifStat        = 18
	AspidaParserRULE_elseStat        = 19
	AspidaParserRULE_comparison      = 20
	AspidaParserRULE_vars_content    = 21
	AspidaParserRULE_vars_prop       = 22
	AspidaParserRULE_comp_op         = 23
	AspidaParserRULE_value           = 24
	AspidaParserRULE_str_array       = 25
	AspidaParserRULE_array           = 26
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
	p.RuleIndex = AspidaParserRULE_program
	return p
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_program

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

func (p *AspidaParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, AspidaParserRULE_program)
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
		p.SetState(54)
		p.Main()
	}
	{
		p.SetState(55)
		p.Hosts()
	}
	{
		p.SetState(56)
		p.Tasks()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == AspidaParserVARS_KW {
		{
			p.SetState(57)
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
	p.RuleIndex = AspidaParserRULE_main
	return p
}

func (*MainContext) IsMainContext() {}

func NewMainContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MainContext {
	var p = new(MainContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_main

	return p
}

func (s *MainContext) GetParser() antlr.Parser { return s.parser }

func (s *MainContext) MAIN_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserMAIN_KW, 0)
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

func (p *AspidaParser) Main() (localctx IMainContext) {
	localctx = NewMainContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, AspidaParserRULE_main)

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
		p.SetState(60)
		p.Match(AspidaParserMAIN_KW)
	}
	{
		p.SetState(61)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(62)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(63)
		p.Main_content()
	}
	{
		p.SetState(64)
		p.Match(AspidaParserT__2)
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
	p.RuleIndex = AspidaParserRULE_hosts
	return p
}

func (*HostsContext) IsHostsContext() {}

func NewHostsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HostsContext {
	var p = new(HostsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_hosts

	return p
}

func (s *HostsContext) GetParser() antlr.Parser { return s.parser }

func (s *HostsContext) HOSTS_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserHOSTS_KW, 0)
}

func (s *HostsContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *HostsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *HostsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HostsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Hosts() (localctx IHostsContext) {
	localctx = NewHostsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, AspidaParserRULE_hosts)

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
		p.SetState(66)
		p.Match(AspidaParserHOSTS_KW)
	}
	{
		p.SetState(67)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(68)
		p.Match(AspidaParserSTRING)
	}
	{
		p.SetState(69)
		p.Match(AspidaParserNS)
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
	p.RuleIndex = AspidaParserRULE_tasks
	return p
}

func (*TasksContext) IsTasksContext() {}

func NewTasksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TasksContext {
	var p = new(TasksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tasks

	return p
}

func (s *TasksContext) GetParser() antlr.Parser { return s.parser }

func (s *TasksContext) TASKS_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserTASKS_KW, 0)
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

func (p *AspidaParser) Tasks() (localctx ITasksContext) {
	localctx = NewTasksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, AspidaParserRULE_tasks)

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
		p.SetState(71)
		p.Match(AspidaParserTASKS_KW)
	}
	{
		p.SetState(72)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(73)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(74)
		p.Tasks_content()
	}
	{
		p.SetState(75)
		p.Match(AspidaParserT__2)
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
	p.RuleIndex = AspidaParserRULE_variables
	return p
}

func (*VariablesContext) IsVariablesContext() {}

func NewVariablesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariablesContext {
	var p = new(VariablesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_variables

	return p
}

func (s *VariablesContext) GetParser() antlr.Parser { return s.parser }

func (s *VariablesContext) VARS_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserVARS_KW, 0)
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

func (p *AspidaParser) Variables() (localctx IVariablesContext) {
	localctx = NewVariablesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, AspidaParserRULE_variables)

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
		p.SetState(77)
		p.Match(AspidaParserVARS_KW)
	}
	{
		p.SetState(78)
		p.Match(AspidaParserT__0)
	}
	{
		p.SetState(79)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(80)
		p.Vars_content()
	}
	{
		p.SetState(81)
		p.Match(AspidaParserT__2)
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
	p.RuleIndex = AspidaParserRULE_main_content
	return p
}

func (*Main_contentContext) IsMain_contentContext() {}

func NewMain_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_contentContext {
	var p = new(Main_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_main_content

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

func (p *AspidaParser) Main_content() (localctx IMain_contentContext) {
	localctx = NewMain_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, AspidaParserRULE_main_content)
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
		p.SetState(83)
		p.Main_prop()
	}
	p.SetState(87)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AspidaParserT__3)|(1<<AspidaParserT__4)|(1<<AspidaParserT__5)|(1<<AspidaParserT__6)|(1<<AspidaParserT__7)|(1<<AspidaParserT__8))) != 0 {
		{
			p.SetState(84)
			p.Main_prop()
		}

		p.SetState(89)
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
	p.RuleIndex = AspidaParserRULE_main_prop
	return p
}

func (*Main_propContext) IsMain_propContext() {}

func NewMain_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Main_propContext {
	var p = new(Main_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_main_prop

	return p
}

func (s *Main_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Main_propContext) CopyFrom(ctx *Main_propContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Main_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Main_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type NameMainContext struct {
	*Main_propContext
}

func NewNameMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NameMainContext {
	var p = new(NameMainContext)

	p.Main_propContext = NewEmptyMain_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Main_propContext))

	return p
}

func (s *NameMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameMainContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NameMainContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

type ConnectionMainContext struct {
	*Main_propContext
}

func NewConnectionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConnectionMainContext {
	var p = new(ConnectionMainContext)

	p.Main_propContext = NewEmptyMain_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Main_propContext))

	return p
}

func (s *ConnectionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionMainContext) Connection() IConnectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnectionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnectionContext)
}

func (s *ConnectionMainContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

type DescriptionMainContext struct {
	*Main_propContext
}

func NewDescriptionMainContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DescriptionMainContext {
	var p = new(DescriptionMainContext)

	p.Main_propContext = NewEmptyMain_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Main_propContext))

	return p
}

func (s *DescriptionMainContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionMainContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DescriptionMainContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (p *AspidaParser) Main_prop() (localctx IMain_propContext) {
	localctx = NewMain_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, AspidaParserRULE_main_prop)

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

	p.SetState(99)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__3, AspidaParserT__4:
		localctx = NewNameMainContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Name()
		}
		{
			p.SetState(91)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__5, AspidaParserT__6:
		localctx = NewConnectionMainContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(93)
			p.Connection()
		}
		{
			p.SetState(94)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__7, AspidaParserT__8:
		localctx = NewDescriptionMainContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(96)
			p.Description()
		}
		{
			p.SetState(97)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = AspidaParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, AspidaParserRULE_name)

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

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(101)
			p.Match(AspidaParserT__3)
		}
		{
			p.SetState(102)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(103)
			p.Value()
		}

	case AspidaParserT__4:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(104)
			p.Match(AspidaParserT__4)
		}
		{
			p.SetState(105)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(106)
			p.Value()
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
	p.RuleIndex = AspidaParserRULE_connection
	return p
}

func (*ConnectionContext) IsConnectionContext() {}

func NewConnectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConnectionContext {
	var p = new(ConnectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_connection

	return p
}

func (s *ConnectionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConnectionContext) Connection_type() IConnection_typeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConnection_typeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConnection_typeContext)
}

func (s *ConnectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Connection() (localctx IConnectionContext) {
	localctx = NewConnectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, AspidaParserRULE_connection)

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

	p.SetState(115)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__5:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.Match(AspidaParserT__5)
		}
		{
			p.SetState(110)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(111)
			p.Connection_type()
		}

	case AspidaParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(112)
			p.Match(AspidaParserT__6)
		}
		{
			p.SetState(113)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(114)
			p.Connection_type()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IConnection_typeContext is an interface to support dynamic dispatch.
type IConnection_typeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConnection_typeContext differentiates from other interfaces.
	IsConnection_typeContext()
}

type Connection_typeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConnection_typeContext() *Connection_typeContext {
	var p = new(Connection_typeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_connection_type
	return p
}

func (*Connection_typeContext) IsConnection_typeContext() {}

func NewConnection_typeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Connection_typeContext {
	var p = new(Connection_typeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_connection_type

	return p
}

func (s *Connection_typeContext) GetParser() antlr.Parser { return s.parser }

func (s *Connection_typeContext) CopyFrom(ctx *Connection_typeContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Connection_typeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Connection_typeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ConnectionSSHContext struct {
	*Connection_typeContext
}

func NewConnectionSSHContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConnectionSSHContext {
	var p = new(ConnectionSSHContext)

	p.Connection_typeContext = NewEmptyConnection_typeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Connection_typeContext))

	return p
}

func (s *ConnectionSSHContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionSSHContext) SSH_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserSSH_KW, 0)
}

type ConnectionLocalContext struct {
	*Connection_typeContext
}

func NewConnectionLocalContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConnectionLocalContext {
	var p = new(ConnectionLocalContext)

	p.Connection_typeContext = NewEmptyConnection_typeContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Connection_typeContext))

	return p
}

func (s *ConnectionLocalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConnectionLocalContext) LOCAL_KW() antlr.TerminalNode {
	return s.GetToken(AspidaParserLOCAL_KW, 0)
}

func (p *AspidaParser) Connection_type() (localctx IConnection_typeContext) {
	localctx = NewConnection_typeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, AspidaParserRULE_connection_type)

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

	p.SetState(119)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserLOCAL_KW:
		localctx = NewConnectionLocalContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(117)
			p.Match(AspidaParserLOCAL_KW)
		}

	case AspidaParserSSH_KW:
		localctx = NewConnectionSSHContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(118)
			p.Match(AspidaParserSSH_KW)
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
	p.RuleIndex = AspidaParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, AspidaParserRULE_description)

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

	p.SetState(127)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__7:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(121)
			p.Match(AspidaParserT__7)
		}
		{
			p.SetState(122)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(123)
			p.Value()
		}

	case AspidaParserT__8:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(124)
			p.Match(AspidaParserT__8)
		}
		{
			p.SetState(125)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(126)
			p.Value()
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
	p.RuleIndex = AspidaParserRULE_tasks_content
	return p
}

func (*Tasks_contentContext) IsTasks_contentContext() {}

func NewTasks_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tasks_contentContext {
	var p = new(Tasks_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tasks_content

	return p
}

func (s *Tasks_contentContext) GetParser() antlr.Parser { return s.parser }

func (s *Tasks_contentContext) CopyFrom(ctx *Tasks_contentContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Tasks_contentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tasks_contentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type IfStatementContext struct {
	*Tasks_contentContext
}

func NewIfStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfStatementContext {
	var p = new(IfStatementContext)

	p.Tasks_contentContext = NewEmptyTasks_contentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_contentContext))

	return p
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) IfStat() IIfStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfStatContext)
}

func (s *IfStatementContext) ElseStat() IElseStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElseStatContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElseStatContext)
}

func (s *IfStatementContext) AllElifStat() []IElifStatContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IElifStatContext)(nil)).Elem())
	var tst = make([]IElifStatContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IElifStatContext)
		}
	}

	return tst
}

func (s *IfStatementContext) ElifStat(i int) IElifStatContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElifStatContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IElifStatContext)
}

type TContentContext struct {
	*Tasks_contentContext
}

func NewTContentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TContentContext {
	var p = new(TContentContext)

	p.Tasks_contentContext = NewEmptyTasks_contentContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_contentContext))

	return p
}

func (s *TContentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TContentContext) AllTasks_prop() []ITasks_propContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITasks_propContext)(nil)).Elem())
	var tst = make([]ITasks_propContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITasks_propContext)
		}
	}

	return tst
}

func (s *TContentContext) Tasks_prop(i int) ITasks_propContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_propContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITasks_propContext)
}

func (p *AspidaParser) Tasks_content() (localctx ITasks_contentContext) {
	localctx = NewTasks_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, AspidaParserRULE_tasks_content)
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

	p.SetState(145)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__9, AspidaParserT__10, AspidaParserT__11, AspidaParserT__12, AspidaParserT__13, AspidaParserT__14, AspidaParserT__15, AspidaParserT__16:
		localctx = NewTContentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(129)
			p.Tasks_prop()
		}
		p.SetState(133)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AspidaParserT__9)|(1<<AspidaParserT__10)|(1<<AspidaParserT__11)|(1<<AspidaParserT__12)|(1<<AspidaParserT__13)|(1<<AspidaParserT__14)|(1<<AspidaParserT__15)|(1<<AspidaParserT__16))) != 0 {
			{
				p.SetState(130)
				p.Tasks_prop()
			}

			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case AspidaParserIF:
		localctx = NewIfStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(136)
			p.IfStat()
		}
		p.SetState(140)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AspidaParserELIF {
			{
				p.SetState(137)
				p.ElifStat()
			}

			p.SetState(142)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(143)
			p.ElseStat()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = AspidaParserRULE_tasks_prop
	return p
}

func (*Tasks_propContext) IsTasks_propContext() {}

func NewTasks_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tasks_propContext {
	var p = new(Tasks_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_tasks_prop

	return p
}

func (s *Tasks_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Tasks_propContext) CopyFrom(ctx *Tasks_propContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Tasks_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tasks_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TPointsContext struct {
	*Tasks_propContext
}

func NewTPointsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TPointsContext {
	var p = new(TPointsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TPointsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TPointsContext) Points() IPointsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPointsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPointsContext)
}

type TControlsContext struct {
	*Tasks_propContext
}

func NewTControlsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TControlsContext {
	var p = new(TControlsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TControlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TControlsContext) Controls() IControlsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IControlsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IControlsContext)
}

type TExclusionsContext struct {
	*Tasks_propContext
}

func NewTExclusionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TExclusionsContext {
	var p = new(TExclusionsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TExclusionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TExclusionsContext) Exclusions() IExclusionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExclusionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExclusionsContext)
}

type TSectionsContext struct {
	*Tasks_propContext
}

func NewTSectionsContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TSectionsContext {
	var p = new(TSectionsContext)

	p.Tasks_propContext = NewEmptyTasks_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Tasks_propContext))

	return p
}

func (s *TSectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TSectionsContext) Sections() ISectionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISectionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISectionsContext)
}

func (p *AspidaParser) Tasks_prop() (localctx ITasks_propContext) {
	localctx = NewTasks_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, AspidaParserRULE_tasks_prop)

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

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__9, AspidaParserT__10:
		localctx = NewTSectionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(147)
			p.Sections()
		}

	case AspidaParserT__11, AspidaParserT__12:
		localctx = NewTPointsContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(148)
			p.Points()
		}

	case AspidaParserT__13, AspidaParserT__14:
		localctx = NewTControlsContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(149)
			p.Controls()
		}

	case AspidaParserT__15, AspidaParserT__16:
		localctx = NewTExclusionsContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(150)
			p.Exclusions()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	p.RuleIndex = AspidaParserRULE_sections
	return p
}

func (*SectionsContext) IsSectionsContext() {}

func NewSectionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SectionsContext {
	var p = new(SectionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_sections

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

func (s *SectionsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *SectionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SectionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Sections() (localctx ISectionsContext) {
	localctx = NewSectionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, AspidaParserRULE_sections)

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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__9:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.Match(AspidaParserT__9)
		}
		{
			p.SetState(154)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(155)
			p.Str_array()
		}
		{
			p.SetState(156)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__10:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(158)
			p.Match(AspidaParserT__10)
		}
		{
			p.SetState(159)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(160)
			p.Str_array()
		}
		{
			p.SetState(161)
			p.Match(AspidaParserNS)
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
	p.RuleIndex = AspidaParserRULE_points
	return p
}

func (*PointsContext) IsPointsContext() {}

func NewPointsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PointsContext {
	var p = new(PointsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_points

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

func (s *PointsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *PointsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PointsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Points() (localctx IPointsContext) {
	localctx = NewPointsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, AspidaParserRULE_points)

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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__11:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.Match(AspidaParserT__11)
		}
		{
			p.SetState(166)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(167)
			p.Str_array()
		}
		{
			p.SetState(168)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__12:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(170)
			p.Match(AspidaParserT__12)
		}
		{
			p.SetState(171)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(172)
			p.Str_array()
		}
		{
			p.SetState(173)
			p.Match(AspidaParserNS)
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
	p.RuleIndex = AspidaParserRULE_controls
	return p
}

func (*ControlsContext) IsControlsContext() {}

func NewControlsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ControlsContext {
	var p = new(ControlsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_controls

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

func (s *ControlsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *ControlsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ControlsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Controls() (localctx IControlsContext) {
	localctx = NewControlsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, AspidaParserRULE_controls)

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

	p.SetState(187)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__13:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(177)
			p.Match(AspidaParserT__13)
		}
		{
			p.SetState(178)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(179)
			p.Str_array()
		}
		{
			p.SetState(180)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__14:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(182)
			p.Match(AspidaParserT__14)
		}
		{
			p.SetState(183)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(184)
			p.Str_array()
		}
		{
			p.SetState(185)
			p.Match(AspidaParserNS)
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
	p.RuleIndex = AspidaParserRULE_exclusions
	return p
}

func (*ExclusionsContext) IsExclusionsContext() {}

func NewExclusionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExclusionsContext {
	var p = new(ExclusionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_exclusions

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

func (s *ExclusionsContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (s *ExclusionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExclusionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Exclusions() (localctx IExclusionsContext) {
	localctx = NewExclusionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, AspidaParserRULE_exclusions)

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

	p.SetState(199)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case AspidaParserT__15:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(189)
			p.Match(AspidaParserT__15)
		}
		{
			p.SetState(190)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(191)
			p.Str_array()
		}
		{
			p.SetState(192)
			p.Match(AspidaParserNS)
		}

	case AspidaParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(194)
			p.Match(AspidaParserT__16)
		}
		{
			p.SetState(195)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(196)
			p.Str_array()
		}
		{
			p.SetState(197)
			p.Match(AspidaParserNS)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IIfStatContext is an interface to support dynamic dispatch.
type IIfStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIfStatContext differentiates from other interfaces.
	IsIfStatContext()
}

type IfStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatContext() *IfStatContext {
	var p = new(IfStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_ifStat
	return p
}

func (*IfStatContext) IsIfStatContext() {}

func NewIfStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatContext {
	var p = new(IfStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_ifStat

	return p
}

func (s *IfStatContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatContext) IF() antlr.TerminalNode {
	return s.GetToken(AspidaParserIF, 0)
}

func (s *IfStatContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *IfStatContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *IfStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) IfStat() (localctx IIfStatContext) {
	localctx = NewIfStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, AspidaParserRULE_ifStat)

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
		p.SetState(201)
		p.Match(AspidaParserIF)
	}
	{
		p.SetState(202)
		p.Comparison()
	}
	{
		p.SetState(203)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(204)
		p.Tasks_content()
	}
	{
		p.SetState(205)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IElifStatContext is an interface to support dynamic dispatch.
type IElifStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElifStatContext differentiates from other interfaces.
	IsElifStatContext()
}

type ElifStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElifStatContext() *ElifStatContext {
	var p = new(ElifStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_elifStat
	return p
}

func (*ElifStatContext) IsElifStatContext() {}

func NewElifStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElifStatContext {
	var p = new(ElifStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_elifStat

	return p
}

func (s *ElifStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElifStatContext) ELIF() antlr.TerminalNode {
	return s.GetToken(AspidaParserELIF, 0)
}

func (s *ElifStatContext) Comparison() IComparisonContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonContext)
}

func (s *ElifStatContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *ElifStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElifStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) ElifStat() (localctx IElifStatContext) {
	localctx = NewElifStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, AspidaParserRULE_elifStat)

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
		p.SetState(207)
		p.Match(AspidaParserELIF)
	}
	{
		p.SetState(208)
		p.Comparison()
	}
	{
		p.SetState(209)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(210)
		p.Tasks_content()
	}
	{
		p.SetState(211)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IElseStatContext is an interface to support dynamic dispatch.
type IElseStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsElseStatContext differentiates from other interfaces.
	IsElseStatContext()
}

type ElseStatContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyElseStatContext() *ElseStatContext {
	var p = new(ElseStatContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_elseStat
	return p
}

func (*ElseStatContext) IsElseStatContext() {}

func NewElseStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElseStatContext {
	var p = new(ElseStatContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_elseStat

	return p
}

func (s *ElseStatContext) GetParser() antlr.Parser { return s.parser }

func (s *ElseStatContext) ELSE() antlr.TerminalNode {
	return s.GetToken(AspidaParserELSE, 0)
}

func (s *ElseStatContext) Tasks_content() ITasks_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITasks_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITasks_contentContext)
}

func (s *ElseStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElseStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) ElseStat() (localctx IElseStatContext) {
	localctx = NewElseStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, AspidaParserRULE_elseStat)

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
		p.SetState(213)
		p.Match(AspidaParserELSE)
	}
	{
		p.SetState(214)
		p.Match(AspidaParserT__1)
	}
	{
		p.SetState(215)
		p.Tasks_content()
	}
	{
		p.SetState(216)
		p.Match(AspidaParserT__2)
	}

	return localctx
}

// IComparisonContext is an interface to support dynamic dispatch.
type IComparisonContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonContext differentiates from other interfaces.
	IsComparisonContext()
}

type ComparisonContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonContext() *ComparisonContext {
	var p = new(ComparisonContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_comparison
	return p
}

func (*ComparisonContext) IsComparisonContext() {}

func NewComparisonContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonContext {
	var p = new(ComparisonContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_comparison

	return p
}

func (s *ComparisonContext) GetParser() antlr.Parser { return s.parser }

func (s *ComparisonContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ComparisonContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ComparisonContext) Comp_op() IComp_opContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComp_opContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComp_opContext)
}

func (s *ComparisonContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Comparison() (localctx IComparisonContext) {
	localctx = NewComparisonContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, AspidaParserRULE_comparison)
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
		p.SetState(218)
		p.Value()
	}
	{
		p.SetState(219)
		p.Comp_op()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la-24)&-(0x1f+1)) == 0 && ((1<<uint((_la-24)))&((1<<(AspidaParserT__23-24))|(1<<(AspidaParserT__24-24))|(1<<(AspidaParserT__25-24))|(1<<(AspidaParserT__26-24))|(1<<(AspidaParserSTRING-24))|(1<<(AspidaParserNUMBER-24)))) != 0 {
		{
			p.SetState(220)
			p.Value()
		}

		p.SetState(225)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = AspidaParserRULE_vars_content
	return p
}

func (*Vars_contentContext) IsVars_contentContext() {}

func NewVars_contentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_contentContext {
	var p = new(Vars_contentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_vars_content

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

func (p *AspidaParser) Vars_content() (localctx IVars_contentContext) {
	localctx = NewVars_contentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, AspidaParserRULE_vars_content)
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
		p.SetState(226)
		p.Vars_prop()
	}
	p.SetState(230)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == AspidaParserSTRING {
		{
			p.SetState(227)
			p.Vars_prop()
		}

		p.SetState(232)
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
	p.RuleIndex = AspidaParserRULE_vars_prop
	return p
}

func (*Vars_propContext) IsVars_propContext() {}

func NewVars_propContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Vars_propContext {
	var p = new(Vars_propContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_vars_prop

	return p
}

func (s *Vars_propContext) GetParser() antlr.Parser { return s.parser }

func (s *Vars_propContext) CopyFrom(ctx *Vars_propContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *Vars_propContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Vars_propContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarDefContext struct {
	*Vars_propContext
}

func NewVarDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarDefContext {
	var p = new(VarDefContext)

	p.Vars_propContext = NewEmptyVars_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Vars_propContext))

	return p
}

func (s *VarDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDefContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *VarDefContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *VarDefContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

type VarObjDefContext struct {
	*Vars_propContext
}

func NewVarObjDefContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarObjDefContext {
	var p = new(VarObjDefContext)

	p.Vars_propContext = NewEmptyVars_propContext()
	p.parser = parser
	p.CopyFrom(ctx.(*Vars_propContext))

	return p
}

func (s *VarObjDefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarObjDefContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

func (s *VarObjDefContext) Vars_content() IVars_contentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVars_contentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVars_contentContext)
}

func (s *VarObjDefContext) NS() antlr.TerminalNode {
	return s.GetToken(AspidaParserNS, 0)
}

func (p *AspidaParser) Vars_prop() (localctx IVars_propContext) {
	localctx = NewVars_propContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, AspidaParserRULE_vars_prop)

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

	p.SetState(245)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(233)
			p.Match(AspidaParserSTRING)
		}
		{
			p.SetState(234)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(235)
			p.Value()
		}
		{
			p.SetState(236)
			p.Match(AspidaParserNS)
		}

	case 2:
		localctx = NewVarObjDefContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(238)
			p.Match(AspidaParserSTRING)
		}
		{
			p.SetState(239)
			p.Match(AspidaParserT__0)
		}
		{
			p.SetState(240)
			p.Match(AspidaParserT__1)
		}
		{
			p.SetState(241)
			p.Vars_content()
		}
		{
			p.SetState(242)
			p.Match(AspidaParserT__2)
		}
		{
			p.SetState(243)
			p.Match(AspidaParserNS)
		}

	}

	return localctx
}

// IComp_opContext is an interface to support dynamic dispatch.
type IComp_opContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComp_opContext differentiates from other interfaces.
	IsComp_opContext()
}

type Comp_opContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComp_opContext() *Comp_opContext {
	var p = new(Comp_opContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = AspidaParserRULE_comp_op
	return p
}

func (*Comp_opContext) IsComp_opContext() {}

func NewComp_opContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Comp_opContext {
	var p = new(Comp_opContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_comp_op

	return p
}

func (s *Comp_opContext) GetParser() antlr.Parser { return s.parser }
func (s *Comp_opContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Comp_opContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Comp_op() (localctx IComp_opContext) {
	localctx = NewComp_opContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, AspidaParserRULE_comp_op)
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
		p.SetState(247)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<AspidaParserT__17)|(1<<AspidaParserT__18)|(1<<AspidaParserT__19)|(1<<AspidaParserT__20)|(1<<AspidaParserT__21)|(1<<AspidaParserT__22))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	p.RuleIndex = AspidaParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) CopyFrom(ctx *ValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TrueValContext struct {
	*ValueContext
}

func NewTrueValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TrueValContext {
	var p = new(TrueValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *TrueValContext) GetRuleContext() antlr.RuleContext {
	return s
}

type NumberValContext struct {
	*ValueContext
}

func NewNumberValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberValContext {
	var p = new(NumberValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NumberValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberValContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(AspidaParserNUMBER, 0)
}

type ArrayValContext struct {
	*ValueContext
}

func NewArrayValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ArrayValContext {
	var p = new(ArrayValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *ArrayValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayValContext) Array() IArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayContext)
}

type FalseValContext struct {
	*ValueContext
}

func NewFalseValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FalseValContext {
	var p = new(FalseValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *FalseValContext) GetRuleContext() antlr.RuleContext {
	return s
}

type StringValContext struct {
	*ValueContext
}

func NewStringValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringValContext {
	var p = new(StringValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *StringValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValContext) STRING() antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, 0)
}

type NullValContext struct {
	*ValueContext
}

func NewNullValContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullValContext {
	var p = new(NullValContext)

	p.ValueContext = NewEmptyValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueContext))

	return p
}

func (s *NullValContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (p *AspidaParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, AspidaParserRULE_value)

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

	switch p.GetTokenStream().LA(1) {
	case AspidaParserSTRING:
		localctx = NewStringValContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(249)
			p.Match(AspidaParserSTRING)
		}

	case AspidaParserNUMBER:
		localctx = NewNumberValContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(250)
			p.Match(AspidaParserNUMBER)
		}

	case AspidaParserT__23:
		localctx = NewTrueValContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(251)
			p.Match(AspidaParserT__23)
		}

	case AspidaParserT__24:
		localctx = NewFalseValContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(252)
			p.Match(AspidaParserT__24)
		}

	case AspidaParserT__25:
		localctx = NewNullValContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)
			p.Match(AspidaParserT__25)
		}

	case AspidaParserT__26:
		localctx = NewArrayValContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(254)
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
	p.RuleIndex = AspidaParserRULE_str_array
	return p
}

func (*Str_arrayContext) IsStr_arrayContext() {}

func NewStr_arrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Str_arrayContext {
	var p = new(Str_arrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_str_array

	return p
}

func (s *Str_arrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Str_arrayContext) AllSTRING() []antlr.TerminalNode {
	return s.GetTokens(AspidaParserSTRING)
}

func (s *Str_arrayContext) STRING(i int) antlr.TerminalNode {
	return s.GetToken(AspidaParserSTRING, i)
}

func (s *Str_arrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Str_arrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (p *AspidaParser) Str_array() (localctx IStr_arrayContext) {
	localctx = NewStr_arrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, AspidaParserRULE_str_array)
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

	p.SetState(269)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(257)
			p.Match(AspidaParserT__26)
		}
		{
			p.SetState(258)
			p.Match(AspidaParserSTRING)
		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AspidaParserT__27 {
			{
				p.SetState(259)
				p.Match(AspidaParserT__27)
			}
			{
				p.SetState(260)
				p.Match(AspidaParserSTRING)
			}

			p.SetState(265)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(266)
			p.Match(AspidaParserT__28)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(267)
			p.Match(AspidaParserT__26)
		}
		{
			p.SetState(268)
			p.Match(AspidaParserT__28)
		}

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
	p.RuleIndex = AspidaParserRULE_array
	return p
}

func (*ArrayContext) IsArrayContext() {}

func NewArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayContext {
	var p = new(ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = AspidaParserRULE_array

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

func (p *AspidaParser) Array() (localctx IArrayContext) {
	localctx = NewArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, AspidaParserRULE_array)
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

	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(271)
			p.Match(AspidaParserT__26)
		}
		{
			p.SetState(272)
			p.Value()
		}
		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == AspidaParserT__27 {
			{
				p.SetState(273)
				p.Match(AspidaParserT__27)
			}
			{
				p.SetState(274)
				p.Value()
			}

			p.SetState(279)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(280)
			p.Match(AspidaParserT__28)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Match(AspidaParserT__26)
		}
		{
			p.SetState(283)
			p.Match(AspidaParserT__28)
		}

	}

	return localctx
}
