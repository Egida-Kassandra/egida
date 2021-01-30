// Generated from d:\OneDrive\Universidad de Oviedo\Dr. & Dra F - General\TFMs\Egida\Code\egida\Aspida.g4 by ANTLR 4.8
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class AspidaLexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, COMMENT=30, WHITESPACE=31, 
		NEWLINE=32, STRING=33, NUMBER=34, NS=35, MAIN_KW=36, HOSTS_KW=37, TASKS_KW=38, 
		VARS_KW=39, LOCAL_KW=40, SSH_KW=41, SSHPASS_KW=42, IF=43, ELIF=44, ELSE=45, 
		OR=46, AND=47, NOT=48, IS=49;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	private static String[] makeRuleNames() {
		return new String[] {
			"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8", 
			"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16", 
			"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24", 
			"T__25", "T__26", "T__27", "T__28", "COMMENT", "WHITESPACE", "NEWLINE", 
			"ESC", "UNICODE", "HEX", "SAFECODEPOINT", "INT", "EXP", "STRING", "NUMBER", 
			"NS", "MAIN_KW", "HOSTS_KW", "TASKS_KW", "VARS_KW", "LOCAL_KW", "SSH_KW", 
			"SSHPASS_KW", "IF", "ELIF", "ELSE", "OR", "AND", "NOT", "IS"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "':'", "'{'", "'}'", "'name'", "'NAME'", "'connection'", "'CONNECTION'", 
			"'description'", "'DESCRIPTION'", "'sections'", "'SECTIONS'", "'points'", 
			"'POINTS'", "'controls'", "'CONTROLS'", "'exclusions'", "'EXCLUSIONS'", 
			"'<'", "'>'", "'=='", "'>='", "'<='", "'!='", "'true'", "'false'", "'null'", 
			"'['", "','", "']'", null, null, null, null, null, "';'", "'MAIN'", "'HOSTS-GROUP'", 
			"'TASKS'", "'VARS'", "'LOCAL'", "'SSH'", "'SSH-PASSWORD'", "'IF'", "'ELIF'", 
			"'ELSE'", "'OR'", "'AND'", "'NOT'", "'IS'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, "COMMENT", "WHITESPACE", "NEWLINE", 
			"STRING", "NUMBER", "NS", "MAIN_KW", "HOSTS_KW", "TASKS_KW", "VARS_KW", 
			"LOCAL_KW", "SSH_KW", "SSHPASS_KW", "IF", "ELIF", "ELSE", "OR", "AND", 
			"NOT", "IS"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public AspidaLexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Aspida.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\63\u01bd\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\4#\t#\4$\t$\4%\t%\4&\t&\4\'\t\'\4(\t(\4)\t)\4*\t*\4+\t"+
		"+\4,\t,\4-\t-\4.\t.\4/\t/\4\60\t\60\4\61\t\61\4\62\t\62\4\63\t\63\4\64"+
		"\t\64\4\65\t\65\4\66\t\66\4\67\t\67\48\t8\3\2\3\2\3\3\3\3\3\4\3\4\3\5"+
		"\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\7\3\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\t\3\t\3\t\3\t"+
		"\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\r\3\r\3\r\3\r\3\r\3\r\3\r\3\16\3\16\3\16\3\16"+
		"\3\16\3\16\3\16\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\23\3\23\3\24\3\24\3\25\3\25\3\25\3\26\3\26\3\26\3\27\3\27\3\27"+
		"\3\30\3\30\3\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3\32\3\32\3\32"+
		"\3\33\3\33\3\33\3\33\3\33\3\34\3\34\3\35\3\35\3\36\3\36\3\37\3\37\7\37"+
		"\u0120\n\37\f\37\16\37\u0123\13\37\3\37\3\37\3 \3 \3 \3 \3!\5!\u012c\n"+
		"!\3!\3!\6!\u0130\n!\r!\16!\u0131\3!\3!\3\"\3\"\3\"\5\"\u0139\n\"\3#\3"+
		"#\3#\3#\3#\3#\3$\3$\3%\3%\3&\3&\3&\7&\u0148\n&\f&\16&\u014b\13&\5&\u014d"+
		"\n&\3\'\3\'\5\'\u0151\n\'\3\'\3\'\3(\3(\3(\7(\u0158\n(\f(\16(\u015b\13"+
		"(\3(\3(\3)\5)\u0160\n)\3)\3)\3)\6)\u0165\n)\r)\16)\u0166\5)\u0169\n)\3"+
		")\5)\u016c\n)\3*\3*\3+\3+\3+\3+\3+\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3,\3"+
		",\3-\3-\3-\3-\3-\3-\3.\3.\3.\3.\3.\3/\3/\3/\3/\3/\3/\3\60\3\60\3\60\3"+
		"\60\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3\61\3"+
		"\62\3\62\3\62\3\63\3\63\3\63\3\63\3\63\3\64\3\64\3\64\3\64\3\64\3\65\3"+
		"\65\3\65\3\66\3\66\3\66\3\66\3\67\3\67\3\67\3\67\38\38\38\2\29\3\3\5\4"+
		"\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27\r\31\16\33\17\35\20\37\21!\22"+
		"#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33\65\34\67\359\36;\37= ?!A\"C"+
		"\2E\2G\2I\2K\2M\2O#Q$S%U&W\'Y([)]*_+a,c-e.g/i\60k\61m\62o\63\3\2\13\4"+
		"\2\f\f\17\17\4\2\13\13\"\"\n\2$$\61\61^^ddhhppttvv\5\2\62;CHch\5\2\2!"+
		"$$^^\3\2\63;\3\2\62;\4\2GGgg\4\2--//\2\u01c4\2\3\3\2\2\2\2\5\3\2\2\2\2"+
		"\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2\2\2\17\3\2\2\2\2\21\3\2"+
		"\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2\31\3\2\2\2\2\33\3\2\2\2"+
		"\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2\2\2%\3\2\2\2\2\'\3\2\2"+
		"\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2\2\61\3\2\2\2\2\63\3\2\2"+
		"\2\2\65\3\2\2\2\2\67\3\2\2\2\29\3\2\2\2\2;\3\2\2\2\2=\3\2\2\2\2?\3\2\2"+
		"\2\2A\3\2\2\2\2O\3\2\2\2\2Q\3\2\2\2\2S\3\2\2\2\2U\3\2\2\2\2W\3\2\2\2\2"+
		"Y\3\2\2\2\2[\3\2\2\2\2]\3\2\2\2\2_\3\2\2\2\2a\3\2\2\2\2c\3\2\2\2\2e\3"+
		"\2\2\2\2g\3\2\2\2\2i\3\2\2\2\2k\3\2\2\2\2m\3\2\2\2\2o\3\2\2\2\3q\3\2\2"+
		"\2\5s\3\2\2\2\7u\3\2\2\2\tw\3\2\2\2\13|\3\2\2\2\r\u0081\3\2\2\2\17\u008c"+
		"\3\2\2\2\21\u0097\3\2\2\2\23\u00a3\3\2\2\2\25\u00af\3\2\2\2\27\u00b8\3"+
		"\2\2\2\31\u00c1\3\2\2\2\33\u00c8\3\2\2\2\35\u00cf\3\2\2\2\37\u00d8\3\2"+
		"\2\2!\u00e1\3\2\2\2#\u00ec\3\2\2\2%\u00f7\3\2\2\2\'\u00f9\3\2\2\2)\u00fb"+
		"\3\2\2\2+\u00fe\3\2\2\2-\u0101\3\2\2\2/\u0104\3\2\2\2\61\u0107\3\2\2\2"+
		"\63\u010c\3\2\2\2\65\u0112\3\2\2\2\67\u0117\3\2\2\29\u0119\3\2\2\2;\u011b"+
		"\3\2\2\2=\u011d\3\2\2\2?\u0126\3\2\2\2A\u012f\3\2\2\2C\u0135\3\2\2\2E"+
		"\u013a\3\2\2\2G\u0140\3\2\2\2I\u0142\3\2\2\2K\u014c\3\2\2\2M\u014e\3\2"+
		"\2\2O\u0154\3\2\2\2Q\u015f\3\2\2\2S\u016d\3\2\2\2U\u016f\3\2\2\2W\u0174"+
		"\3\2\2\2Y\u0180\3\2\2\2[\u0186\3\2\2\2]\u018b\3\2\2\2_\u0191\3\2\2\2a"+
		"\u0195\3\2\2\2c\u01a2\3\2\2\2e\u01a5\3\2\2\2g\u01aa\3\2\2\2i\u01af\3\2"+
		"\2\2k\u01b2\3\2\2\2m\u01b6\3\2\2\2o\u01ba\3\2\2\2qr\7<\2\2r\4\3\2\2\2"+
		"st\7}\2\2t\6\3\2\2\2uv\7\177\2\2v\b\3\2\2\2wx\7p\2\2xy\7c\2\2yz\7o\2\2"+
		"z{\7g\2\2{\n\3\2\2\2|}\7P\2\2}~\7C\2\2~\177\7O\2\2\177\u0080\7G\2\2\u0080"+
		"\f\3\2\2\2\u0081\u0082\7e\2\2\u0082\u0083\7q\2\2\u0083\u0084\7p\2\2\u0084"+
		"\u0085\7p\2\2\u0085\u0086\7g\2\2\u0086\u0087\7e\2\2\u0087\u0088\7v\2\2"+
		"\u0088\u0089\7k\2\2\u0089\u008a\7q\2\2\u008a\u008b\7p\2\2\u008b\16\3\2"+
		"\2\2\u008c\u008d\7E\2\2\u008d\u008e\7Q\2\2\u008e\u008f\7P\2\2\u008f\u0090"+
		"\7P\2\2\u0090\u0091\7G\2\2\u0091\u0092\7E\2\2\u0092\u0093\7V\2\2\u0093"+
		"\u0094\7K\2\2\u0094\u0095\7Q\2\2\u0095\u0096\7P\2\2\u0096\20\3\2\2\2\u0097"+
		"\u0098\7f\2\2\u0098\u0099\7g\2\2\u0099\u009a\7u\2\2\u009a\u009b\7e\2\2"+
		"\u009b\u009c\7t\2\2\u009c\u009d\7k\2\2\u009d\u009e\7r\2\2\u009e\u009f"+
		"\7v\2\2\u009f\u00a0\7k\2\2\u00a0\u00a1\7q\2\2\u00a1\u00a2\7p\2\2\u00a2"+
		"\22\3\2\2\2\u00a3\u00a4\7F\2\2\u00a4\u00a5\7G\2\2\u00a5\u00a6\7U\2\2\u00a6"+
		"\u00a7\7E\2\2\u00a7\u00a8\7T\2\2\u00a8\u00a9\7K\2\2\u00a9\u00aa\7R\2\2"+
		"\u00aa\u00ab\7V\2\2\u00ab\u00ac\7K\2\2\u00ac\u00ad\7Q\2\2\u00ad\u00ae"+
		"\7P\2\2\u00ae\24\3\2\2\2\u00af\u00b0\7u\2\2\u00b0\u00b1\7g\2\2\u00b1\u00b2"+
		"\7e\2\2\u00b2\u00b3\7v\2\2\u00b3\u00b4\7k\2\2\u00b4\u00b5\7q\2\2\u00b5"+
		"\u00b6\7p\2\2\u00b6\u00b7\7u\2\2\u00b7\26\3\2\2\2\u00b8\u00b9\7U\2\2\u00b9"+
		"\u00ba\7G\2\2\u00ba\u00bb\7E\2\2\u00bb\u00bc\7V\2\2\u00bc\u00bd\7K\2\2"+
		"\u00bd\u00be\7Q\2\2\u00be\u00bf\7P\2\2\u00bf\u00c0\7U\2\2\u00c0\30\3\2"+
		"\2\2\u00c1\u00c2\7r\2\2\u00c2\u00c3\7q\2\2\u00c3\u00c4\7k\2\2\u00c4\u00c5"+
		"\7p\2\2\u00c5\u00c6\7v\2\2\u00c6\u00c7\7u\2\2\u00c7\32\3\2\2\2\u00c8\u00c9"+
		"\7R\2\2\u00c9\u00ca\7Q\2\2\u00ca\u00cb\7K\2\2\u00cb\u00cc\7P\2\2\u00cc"+
		"\u00cd\7V\2\2\u00cd\u00ce\7U\2\2\u00ce\34\3\2\2\2\u00cf\u00d0\7e\2\2\u00d0"+
		"\u00d1\7q\2\2\u00d1\u00d2\7p\2\2\u00d2\u00d3\7v\2\2\u00d3\u00d4\7t\2\2"+
		"\u00d4\u00d5\7q\2\2\u00d5\u00d6\7n\2\2\u00d6\u00d7\7u\2\2\u00d7\36\3\2"+
		"\2\2\u00d8\u00d9\7E\2\2\u00d9\u00da\7Q\2\2\u00da\u00db\7P\2\2\u00db\u00dc"+
		"\7V\2\2\u00dc\u00dd\7T\2\2\u00dd\u00de\7Q\2\2\u00de\u00df\7N\2\2\u00df"+
		"\u00e0\7U\2\2\u00e0 \3\2\2\2\u00e1\u00e2\7g\2\2\u00e2\u00e3\7z\2\2\u00e3"+
		"\u00e4\7e\2\2\u00e4\u00e5\7n\2\2\u00e5\u00e6\7w\2\2\u00e6\u00e7\7u\2\2"+
		"\u00e7\u00e8\7k\2\2\u00e8\u00e9\7q\2\2\u00e9\u00ea\7p\2\2\u00ea\u00eb"+
		"\7u\2\2\u00eb\"\3\2\2\2\u00ec\u00ed\7G\2\2\u00ed\u00ee\7Z\2\2\u00ee\u00ef"+
		"\7E\2\2\u00ef\u00f0\7N\2\2\u00f0\u00f1\7W\2\2\u00f1\u00f2\7U\2\2\u00f2"+
		"\u00f3\7K\2\2\u00f3\u00f4\7Q\2\2\u00f4\u00f5\7P\2\2\u00f5\u00f6\7U\2\2"+
		"\u00f6$\3\2\2\2\u00f7\u00f8\7>\2\2\u00f8&\3\2\2\2\u00f9\u00fa\7@\2\2\u00fa"+
		"(\3\2\2\2\u00fb\u00fc\7?\2\2\u00fc\u00fd\7?\2\2\u00fd*\3\2\2\2\u00fe\u00ff"+
		"\7@\2\2\u00ff\u0100\7?\2\2\u0100,\3\2\2\2\u0101\u0102\7>\2\2\u0102\u0103"+
		"\7?\2\2\u0103.\3\2\2\2\u0104\u0105\7#\2\2\u0105\u0106\7?\2\2\u0106\60"+
		"\3\2\2\2\u0107\u0108\7v\2\2\u0108\u0109\7t\2\2\u0109\u010a\7w\2\2\u010a"+
		"\u010b\7g\2\2\u010b\62\3\2\2\2\u010c\u010d\7h\2\2\u010d\u010e\7c\2\2\u010e"+
		"\u010f\7n\2\2\u010f\u0110\7u\2\2\u0110\u0111\7g\2\2\u0111\64\3\2\2\2\u0112"+
		"\u0113\7p\2\2\u0113\u0114\7w\2\2\u0114\u0115\7n\2\2\u0115\u0116\7n\2\2"+
		"\u0116\66\3\2\2\2\u0117\u0118\7]\2\2\u01188\3\2\2\2\u0119\u011a\7.\2\2"+
		"\u011a:\3\2\2\2\u011b\u011c\7_\2\2\u011c<\3\2\2\2\u011d\u0121\7%\2\2\u011e"+
		"\u0120\n\2\2\2\u011f\u011e\3\2\2\2\u0120\u0123\3\2\2\2\u0121\u011f\3\2"+
		"\2\2\u0121\u0122\3\2\2\2\u0122\u0124\3\2\2\2\u0123\u0121\3\2\2\2\u0124"+
		"\u0125\b\37\2\2\u0125>\3\2\2\2\u0126\u0127\t\3\2\2\u0127\u0128\3\2\2\2"+
		"\u0128\u0129\b \2\2\u0129@\3\2\2\2\u012a\u012c\7\17\2\2\u012b\u012a\3"+
		"\2\2\2\u012b\u012c\3\2\2\2\u012c\u012d\3\2\2\2\u012d\u0130\7\f\2\2\u012e"+
		"\u0130\7\17\2\2\u012f\u012b\3\2\2\2\u012f\u012e\3\2\2\2\u0130\u0131\3"+
		"\2\2\2\u0131\u012f\3\2\2\2\u0131\u0132\3\2\2\2\u0132\u0133\3\2\2\2\u0133"+
		"\u0134\b!\2\2\u0134B\3\2\2\2\u0135\u0138\7^\2\2\u0136\u0139\t\4\2\2\u0137"+
		"\u0139\5E#\2\u0138\u0136\3\2\2\2\u0138\u0137\3\2\2\2\u0139D\3\2\2\2\u013a"+
		"\u013b\7w\2\2\u013b\u013c\5G$\2\u013c\u013d\5G$\2\u013d\u013e\5G$\2\u013e"+
		"\u013f\5G$\2\u013fF\3\2\2\2\u0140\u0141\t\5\2\2\u0141H\3\2\2\2\u0142\u0143"+
		"\n\6\2\2\u0143J\3\2\2\2\u0144\u014d\7\62\2\2\u0145\u0149\t\7\2\2\u0146"+
		"\u0148\t\b\2\2\u0147\u0146\3\2\2\2\u0148\u014b\3\2\2\2\u0149\u0147\3\2"+
		"\2\2\u0149\u014a\3\2\2\2\u014a\u014d\3\2\2\2\u014b\u0149\3\2\2\2\u014c"+
		"\u0144\3\2\2\2\u014c\u0145\3\2\2\2\u014dL\3\2\2\2\u014e\u0150\t\t\2\2"+
		"\u014f\u0151\t\n\2\2\u0150\u014f\3\2\2\2\u0150\u0151\3\2\2\2\u0151\u0152"+
		"\3\2\2\2\u0152\u0153\5K&\2\u0153N\3\2\2\2\u0154\u0159\7$\2\2\u0155\u0158"+
		"\5C\"\2\u0156\u0158\5I%\2\u0157\u0155\3\2\2\2\u0157\u0156\3\2\2\2\u0158"+
		"\u015b\3\2\2\2\u0159\u0157\3\2\2\2\u0159\u015a\3\2\2\2\u015a\u015c\3\2"+
		"\2\2\u015b\u0159\3\2\2\2\u015c\u015d\7$\2\2\u015dP\3\2\2\2\u015e\u0160"+
		"\7/\2\2\u015f\u015e\3\2\2\2\u015f\u0160\3\2\2\2\u0160\u0161\3\2\2\2\u0161"+
		"\u0168\5K&\2\u0162\u0164\7\60\2\2\u0163\u0165\t\b\2\2\u0164\u0163\3\2"+
		"\2\2\u0165\u0166\3\2\2\2\u0166\u0164\3\2\2\2\u0166\u0167\3\2\2\2\u0167"+
		"\u0169\3\2\2\2\u0168\u0162\3\2\2\2\u0168\u0169\3\2\2\2\u0169\u016b\3\2"+
		"\2\2\u016a\u016c\5M\'\2\u016b\u016a\3\2\2\2\u016b\u016c\3\2\2\2\u016c"+
		"R\3\2\2\2\u016d\u016e\7=\2\2\u016eT\3\2\2\2\u016f\u0170\7O\2\2\u0170\u0171"+
		"\7C\2\2\u0171\u0172\7K\2\2\u0172\u0173\7P\2\2\u0173V\3\2\2\2\u0174\u0175"+
		"\7J\2\2\u0175\u0176\7Q\2\2\u0176\u0177\7U\2\2\u0177\u0178\7V\2\2\u0178"+
		"\u0179\7U\2\2\u0179\u017a\7/\2\2\u017a\u017b\7I\2\2\u017b\u017c\7T\2\2"+
		"\u017c\u017d\7Q\2\2\u017d\u017e\7W\2\2\u017e\u017f\7R\2\2\u017fX\3\2\2"+
		"\2\u0180\u0181\7V\2\2\u0181\u0182\7C\2\2\u0182\u0183\7U\2\2\u0183\u0184"+
		"\7M\2\2\u0184\u0185\7U\2\2\u0185Z\3\2\2\2\u0186\u0187\7X\2\2\u0187\u0188"+
		"\7C\2\2\u0188\u0189\7T\2\2\u0189\u018a\7U\2\2\u018a\\\3\2\2\2\u018b\u018c"+
		"\7N\2\2\u018c\u018d\7Q\2\2\u018d\u018e\7E\2\2\u018e\u018f\7C\2\2\u018f"+
		"\u0190\7N\2\2\u0190^\3\2\2\2\u0191\u0192\7U\2\2\u0192\u0193\7U\2\2\u0193"+
		"\u0194\7J\2\2\u0194`\3\2\2\2\u0195\u0196\7U\2\2\u0196\u0197\7U\2\2\u0197"+
		"\u0198\7J\2\2\u0198\u0199\7/\2\2\u0199\u019a\7R\2\2\u019a\u019b\7C\2\2"+
		"\u019b\u019c\7U\2\2\u019c\u019d\7U\2\2\u019d\u019e\7Y\2\2\u019e\u019f"+
		"\7Q\2\2\u019f\u01a0\7T\2\2\u01a0\u01a1\7F\2\2\u01a1b\3\2\2\2\u01a2\u01a3"+
		"\7K\2\2\u01a3\u01a4\7H\2\2\u01a4d\3\2\2\2\u01a5\u01a6\7G\2\2\u01a6\u01a7"+
		"\7N\2\2\u01a7\u01a8\7K\2\2\u01a8\u01a9\7H\2\2\u01a9f\3\2\2\2\u01aa\u01ab"+
		"\7G\2\2\u01ab\u01ac\7N\2\2\u01ac\u01ad\7U\2\2\u01ad\u01ae\7G\2\2\u01ae"+
		"h\3\2\2\2\u01af\u01b0\7Q\2\2\u01b0\u01b1\7T\2\2\u01b1j\3\2\2\2\u01b2\u01b3"+
		"\7C\2\2\u01b3\u01b4\7P\2\2\u01b4\u01b5\7F\2\2\u01b5l\3\2\2\2\u01b6\u01b7"+
		"\7P\2\2\u01b7\u01b8\7Q\2\2\u01b8\u01b9\7V\2\2\u01b9n\3\2\2\2\u01ba\u01bb"+
		"\7K\2\2\u01bb\u01bc\7U\2\2\u01bcp\3\2\2\2\21\2\u0121\u012b\u012f\u0131"+
		"\u0138\u0149\u014c\u0150\u0157\u0159\u015f\u0166\u0168\u016b\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}