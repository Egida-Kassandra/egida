// Generated from d:\OneDrive\Universidad de Oviedo\Dr. & Dra F - General\TFMs\Egida\Code\egida\Aspida.g4 by ANTLR 4.8
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class AspidaParser extends Parser {
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
	public static final int
		RULE_program = 0, RULE_main = 1, RULE_hosts = 2, RULE_tasks = 3, RULE_variables = 4, 
		RULE_main_content = 5, RULE_main_prop = 6, RULE_name = 7, RULE_connection = 8, 
		RULE_description = 9, RULE_tasks_content = 10, RULE_tasks_prop = 11, RULE_sections = 12, 
		RULE_points = 13, RULE_controls = 14, RULE_exclusions = 15, RULE_vars_content = 16, 
		RULE_vars_prop = 17, RULE_ifStat = 18, RULE_expression = 19, RULE_comparison = 20, 
		RULE_comp_op = 21, RULE_value = 22, RULE_str_array = 23, RULE_cadena = 24, 
		RULE_array = 25;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "main", "hosts", "tasks", "variables", "main_content", "main_prop", 
			"name", "connection", "description", "tasks_content", "tasks_prop", "sections", 
			"points", "controls", "exclusions", "vars_content", "vars_prop", "ifStat", 
			"expression", "comparison", "comp_op", "value", "str_array", "cadena", 
			"array"
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

	@Override
	public String getGrammarFileName() { return "Aspida.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public AspidaParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class ProgramContext extends ParserRuleContext {
		public MainContext main() {
			return getRuleContext(MainContext.class,0);
		}
		public HostsContext hosts() {
			return getRuleContext(HostsContext.class,0);
		}
		public TasksContext tasks() {
			return getRuleContext(TasksContext.class,0);
		}
		public VariablesContext variables() {
			return getRuleContext(VariablesContext.class,0);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			main();
			setState(53);
			hosts();
			setState(55);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TASKS_KW) {
				{
				setState(54);
				tasks();
				}
			}

			setState(58);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==VARS_KW) {
				{
				setState(57);
				variables();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MainContext extends ParserRuleContext {
		public TerminalNode MAIN_KW() { return getToken(AspidaParser.MAIN_KW, 0); }
		public Main_contentContext main_content() {
			return getRuleContext(Main_contentContext.class,0);
		}
		public MainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main; }
	}

	public final MainContext main() throws RecognitionException {
		MainContext _localctx = new MainContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_main);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			match(MAIN_KW);
			setState(61);
			match(T__0);
			setState(62);
			match(T__1);
			setState(63);
			main_content();
			setState(64);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class HostsContext extends ParserRuleContext {
		public TerminalNode HOSTS_KW() { return getToken(AspidaParser.HOSTS_KW, 0); }
		public TerminalNode STRING() { return getToken(AspidaParser.STRING, 0); }
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public HostsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_hosts; }
	}

	public final HostsContext hosts() throws RecognitionException {
		HostsContext _localctx = new HostsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_hosts);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(HOSTS_KW);
			setState(67);
			match(T__0);
			setState(68);
			match(STRING);
			setState(69);
			match(NS);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TasksContext extends ParserRuleContext {
		public TerminalNode TASKS_KW() { return getToken(AspidaParser.TASKS_KW, 0); }
		public Tasks_contentContext tasks_content() {
			return getRuleContext(Tasks_contentContext.class,0);
		}
		public TasksContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tasks; }
	}

	public final TasksContext tasks() throws RecognitionException {
		TasksContext _localctx = new TasksContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_tasks);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(TASKS_KW);
			setState(72);
			match(T__0);
			setState(73);
			match(T__1);
			setState(74);
			tasks_content();
			setState(75);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariablesContext extends ParserRuleContext {
		public TerminalNode VARS_KW() { return getToken(AspidaParser.VARS_KW, 0); }
		public Vars_contentContext vars_content() {
			return getRuleContext(Vars_contentContext.class,0);
		}
		public VariablesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variables; }
	}

	public final VariablesContext variables() throws RecognitionException {
		VariablesContext _localctx = new VariablesContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_variables);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			match(VARS_KW);
			setState(78);
			match(T__0);
			setState(79);
			match(T__1);
			setState(80);
			vars_content();
			setState(81);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Main_contentContext extends ParserRuleContext {
		public List<Main_propContext> main_prop() {
			return getRuleContexts(Main_propContext.class);
		}
		public Main_propContext main_prop(int i) {
			return getRuleContext(Main_propContext.class,i);
		}
		public Main_contentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main_content; }
	}

	public final Main_contentContext main_content() throws RecognitionException {
		Main_contentContext _localctx = new Main_contentContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_main_content);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(83);
			main_prop();
			setState(87);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__3) | (1L << T__4) | (1L << T__5) | (1L << T__6) | (1L << T__7) | (1L << T__8))) != 0)) {
				{
				{
				setState(84);
				main_prop();
				}
				}
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Main_propContext extends ParserRuleContext {
		public Main_propContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main_prop; }
	 
		public Main_propContext() { }
		public void copyFrom(Main_propContext ctx) {
			super.copyFrom(ctx);
		}
	}
	public static class NameMainContext extends Main_propContext {
		public NameContext name() {
			return getRuleContext(NameContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public NameMainContext(Main_propContext ctx) { copyFrom(ctx); }
	}
	public static class ConnectionMainContext extends Main_propContext {
		public ConnectionContext connection() {
			return getRuleContext(ConnectionContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public ConnectionMainContext(Main_propContext ctx) { copyFrom(ctx); }
	}
	public static class DescriptionMainContext extends Main_propContext {
		public DescriptionContext description() {
			return getRuleContext(DescriptionContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public DescriptionMainContext(Main_propContext ctx) { copyFrom(ctx); }
	}

	public final Main_propContext main_prop() throws RecognitionException {
		Main_propContext _localctx = new Main_propContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_main_prop);
		try {
			setState(99);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
			case T__4:
				_localctx = new NameMainContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(90);
				name();
				setState(91);
				match(NS);
				}
				break;
			case T__5:
			case T__6:
				_localctx = new ConnectionMainContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(93);
				connection();
				setState(94);
				match(NS);
				}
				break;
			case T__7:
			case T__8:
				_localctx = new DescriptionMainContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(96);
				description();
				setState(97);
				match(NS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class NameContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(AspidaParser.STRING, 0); }
		public NameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_name; }
	}

	public final NameContext name() throws RecognitionException {
		NameContext _localctx = new NameContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_name);
		try {
			setState(107);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__3:
				enterOuterAlt(_localctx, 1);
				{
				setState(101);
				match(T__3);
				setState(102);
				match(T__0);
				setState(103);
				match(STRING);
				}
				break;
			case T__4:
				enterOuterAlt(_localctx, 2);
				{
				setState(104);
				match(T__4);
				setState(105);
				match(T__0);
				setState(106);
				match(STRING);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConnectionContext extends ParserRuleContext {
		public TerminalNode LOCAL_KW() { return getToken(AspidaParser.LOCAL_KW, 0); }
		public TerminalNode SSH_KW() { return getToken(AspidaParser.SSH_KW, 0); }
		public TerminalNode SSHPASS_KW() { return getToken(AspidaParser.SSHPASS_KW, 0); }
		public ConnectionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_connection; }
	}

	public final ConnectionContext connection() throws RecognitionException {
		ConnectionContext _localctx = new ConnectionContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_connection);
		int _la;
		try {
			setState(115);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__5:
				enterOuterAlt(_localctx, 1);
				{
				setState(109);
				match(T__5);
				setState(110);
				match(T__0);
				setState(111);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LOCAL_KW) | (1L << SSH_KW) | (1L << SSHPASS_KW))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case T__6:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				match(T__6);
				setState(113);
				match(T__0);
				setState(114);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << LOCAL_KW) | (1L << SSH_KW) | (1L << SSHPASS_KW))) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DescriptionContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(AspidaParser.STRING, 0); }
		public DescriptionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_description; }
	}

	public final DescriptionContext description() throws RecognitionException {
		DescriptionContext _localctx = new DescriptionContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_description);
		try {
			setState(123);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__7:
				enterOuterAlt(_localctx, 1);
				{
				setState(117);
				match(T__7);
				setState(118);
				match(T__0);
				setState(119);
				match(STRING);
				}
				break;
			case T__8:
				enterOuterAlt(_localctx, 2);
				{
				setState(120);
				match(T__8);
				setState(121);
				match(T__0);
				setState(122);
				match(STRING);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tasks_contentContext extends ParserRuleContext {
		public List<Tasks_propContext> tasks_prop() {
			return getRuleContexts(Tasks_propContext.class);
		}
		public Tasks_propContext tasks_prop(int i) {
			return getRuleContext(Tasks_propContext.class,i);
		}
		public IfStatContext ifStat() {
			return getRuleContext(IfStatContext.class,0);
		}
		public Tasks_contentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tasks_content; }
	}

	public final Tasks_contentContext tasks_content() throws RecognitionException {
		Tasks_contentContext _localctx = new Tasks_contentContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_tasks_content);
		int _la;
		try {
			setState(133);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
			case T__10:
			case T__11:
			case T__12:
			case T__13:
			case T__14:
			case T__15:
			case T__16:
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				tasks_prop();
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__9) | (1L << T__10) | (1L << T__11) | (1L << T__12) | (1L << T__13) | (1L << T__14) | (1L << T__15) | (1L << T__16))) != 0)) {
					{
					{
					setState(126);
					tasks_prop();
					}
					}
					setState(131);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case IF:
				enterOuterAlt(_localctx, 2);
				{
				setState(132);
				ifStat();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tasks_propContext extends ParserRuleContext {
		public SectionsContext sections() {
			return getRuleContext(SectionsContext.class,0);
		}
		public PointsContext points() {
			return getRuleContext(PointsContext.class,0);
		}
		public ControlsContext controls() {
			return getRuleContext(ControlsContext.class,0);
		}
		public ExclusionsContext exclusions() {
			return getRuleContext(ExclusionsContext.class,0);
		}
		public Tasks_propContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tasks_prop; }
	}

	public final Tasks_propContext tasks_prop() throws RecognitionException {
		Tasks_propContext _localctx = new Tasks_propContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_tasks_prop);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
			case T__10:
				{
				setState(135);
				sections();
				}
				break;
			case T__11:
			case T__12:
				{
				setState(136);
				points();
				}
				break;
			case T__13:
			case T__14:
				{
				setState(137);
				controls();
				}
				break;
			case T__15:
			case T__16:
				{
				setState(138);
				exclusions();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class SectionsContext extends ParserRuleContext {
		public Str_arrayContext str_array() {
			return getRuleContext(Str_arrayContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public SectionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sections; }
	}

	public final SectionsContext sections() throws RecognitionException {
		SectionsContext _localctx = new SectionsContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_sections);
		try {
			setState(151);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__9:
				enterOuterAlt(_localctx, 1);
				{
				setState(141);
				match(T__9);
				setState(142);
				match(T__0);
				setState(143);
				str_array();
				setState(144);
				match(NS);
				}
				break;
			case T__10:
				enterOuterAlt(_localctx, 2);
				{
				setState(146);
				match(T__10);
				setState(147);
				match(T__0);
				setState(148);
				str_array();
				setState(149);
				match(NS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class PointsContext extends ParserRuleContext {
		public Str_arrayContext str_array() {
			return getRuleContext(Str_arrayContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public PointsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_points; }
	}

	public final PointsContext points() throws RecognitionException {
		PointsContext _localctx = new PointsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_points);
		try {
			setState(163);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__11:
				enterOuterAlt(_localctx, 1);
				{
				setState(153);
				match(T__11);
				setState(154);
				match(T__0);
				setState(155);
				str_array();
				setState(156);
				match(NS);
				}
				break;
			case T__12:
				enterOuterAlt(_localctx, 2);
				{
				setState(158);
				match(T__12);
				setState(159);
				match(T__0);
				setState(160);
				str_array();
				setState(161);
				match(NS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ControlsContext extends ParserRuleContext {
		public Str_arrayContext str_array() {
			return getRuleContext(Str_arrayContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public ControlsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_controls; }
	}

	public final ControlsContext controls() throws RecognitionException {
		ControlsContext _localctx = new ControlsContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_controls);
		try {
			setState(175);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__13:
				enterOuterAlt(_localctx, 1);
				{
				setState(165);
				match(T__13);
				setState(166);
				match(T__0);
				setState(167);
				str_array();
				setState(168);
				match(NS);
				}
				break;
			case T__14:
				enterOuterAlt(_localctx, 2);
				{
				setState(170);
				match(T__14);
				setState(171);
				match(T__0);
				setState(172);
				str_array();
				setState(173);
				match(NS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExclusionsContext extends ParserRuleContext {
		public Str_arrayContext str_array() {
			return getRuleContext(Str_arrayContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public ExclusionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_exclusions; }
	}

	public final ExclusionsContext exclusions() throws RecognitionException {
		ExclusionsContext _localctx = new ExclusionsContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_exclusions);
		try {
			setState(187);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__15:
				enterOuterAlt(_localctx, 1);
				{
				setState(177);
				match(T__15);
				setState(178);
				match(T__0);
				setState(179);
				str_array();
				setState(180);
				match(NS);
				}
				break;
			case T__16:
				enterOuterAlt(_localctx, 2);
				{
				setState(182);
				match(T__16);
				setState(183);
				match(T__0);
				setState(184);
				str_array();
				setState(185);
				match(NS);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Vars_contentContext extends ParserRuleContext {
		public List<Vars_propContext> vars_prop() {
			return getRuleContexts(Vars_propContext.class);
		}
		public Vars_propContext vars_prop(int i) {
			return getRuleContext(Vars_propContext.class,i);
		}
		public Vars_contentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vars_content; }
	}

	public final Vars_contentContext vars_content() throws RecognitionException {
		Vars_contentContext _localctx = new Vars_contentContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_vars_content);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			vars_prop();
			setState(193);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==STRING) {
				{
				{
				setState(190);
				vars_prop();
				}
				}
				setState(195);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Vars_propContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(AspidaParser.STRING, 0); }
		public ValueContext value() {
			return getRuleContext(ValueContext.class,0);
		}
		public TerminalNode NS() { return getToken(AspidaParser.NS, 0); }
		public Vars_contentContext vars_content() {
			return getRuleContext(Vars_contentContext.class,0);
		}
		public Vars_propContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vars_prop; }
	}

	public final Vars_propContext vars_prop() throws RecognitionException {
		Vars_propContext _localctx = new Vars_propContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_vars_prop);
		try {
			setState(208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				match(STRING);
				setState(197);
				match(T__0);
				setState(198);
				value();
				setState(199);
				match(NS);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(201);
				match(STRING);
				setState(202);
				match(T__0);
				setState(203);
				match(T__1);
				setState(204);
				vars_content();
				setState(205);
				match(T__2);
				setState(206);
				match(NS);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfStatContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(AspidaParser.IF, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public List<Tasks_contentContext> tasks_content() {
			return getRuleContexts(Tasks_contentContext.class);
		}
		public Tasks_contentContext tasks_content(int i) {
			return getRuleContext(Tasks_contentContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(AspidaParser.ELSE, 0); }
		public List<TerminalNode> ELIF() { return getTokens(AspidaParser.ELIF); }
		public TerminalNode ELIF(int i) {
			return getToken(AspidaParser.ELIF, i);
		}
		public IfStatContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStat; }
	}

	public final IfStatContext ifStat() throws RecognitionException {
		IfStatContext _localctx = new IfStatContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_ifStat);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(210);
			match(IF);
			setState(211);
			expression();
			setState(212);
			match(T__1);
			setState(213);
			tasks_content();
			setState(214);
			match(T__2);
			setState(223);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==ELIF) {
				{
				{
				setState(215);
				match(ELIF);
				setState(216);
				expression();
				setState(217);
				match(T__1);
				setState(218);
				tasks_content();
				setState(219);
				match(T__2);
				}
				}
				setState(225);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(226);
			match(ELSE);
			setState(227);
			match(T__1);
			setState(228);
			tasks_content();
			setState(229);
			match(T__2);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public ComparisonContext comparison() {
			return getRuleContext(ComparisonContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			comparison();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ComparisonContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public List<Comp_opContext> comp_op() {
			return getRuleContexts(Comp_opContext.class);
		}
		public Comp_opContext comp_op(int i) {
			return getRuleContext(Comp_opContext.class,i);
		}
		public ComparisonContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparison; }
	}

	public final ComparisonContext comparison() throws RecognitionException {
		ComparisonContext _localctx = new ComparisonContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_comparison);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(233);
			value();
			setState(239);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << OR) | (1L << AND) | (1L << NOT) | (1L << IS))) != 0)) {
				{
				{
				setState(234);
				comp_op();
				setState(235);
				value();
				}
				}
				setState(241);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Comp_opContext extends ParserRuleContext {
		public TerminalNode IS() { return getToken(AspidaParser.IS, 0); }
		public TerminalNode AND() { return getToken(AspidaParser.AND, 0); }
		public TerminalNode NOT() { return getToken(AspidaParser.NOT, 0); }
		public TerminalNode OR() { return getToken(AspidaParser.OR, 0); }
		public Comp_opContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comp_op; }
	}

	public final Comp_opContext comp_op() throws RecognitionException {
		Comp_opContext _localctx = new Comp_opContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_comp_op);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << T__17) | (1L << T__18) | (1L << T__19) | (1L << T__20) | (1L << T__21) | (1L << T__22) | (1L << OR) | (1L << AND) | (1L << NOT) | (1L << IS))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ValueContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(AspidaParser.STRING, 0); }
		public TerminalNode NUMBER() { return getToken(AspidaParser.NUMBER, 0); }
		public ArrayContext array() {
			return getRuleContext(ArrayContext.class,0);
		}
		public ValueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_value; }
	}

	public final ValueContext value() throws RecognitionException {
		ValueContext _localctx = new ValueContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_value);
		try {
			setState(250);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(244);
				match(STRING);
				}
				break;
			case NUMBER:
				enterOuterAlt(_localctx, 2);
				{
				setState(245);
				match(NUMBER);
				}
				break;
			case T__23:
				enterOuterAlt(_localctx, 3);
				{
				setState(246);
				match(T__23);
				}
				break;
			case T__24:
				enterOuterAlt(_localctx, 4);
				{
				setState(247);
				match(T__24);
				}
				break;
			case T__25:
				enterOuterAlt(_localctx, 5);
				{
				setState(248);
				match(T__25);
				}
				break;
			case T__26:
				enterOuterAlt(_localctx, 6);
				{
				setState(249);
				array();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Str_arrayContext extends ParserRuleContext {
		public List<CadenaContext> cadena() {
			return getRuleContexts(CadenaContext.class);
		}
		public CadenaContext cadena(int i) {
			return getRuleContext(CadenaContext.class,i);
		}
		public Str_arrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_str_array; }
	}

	public final Str_arrayContext str_array() throws RecognitionException {
		Str_arrayContext _localctx = new Str_arrayContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_str_array);
		int _la;
		try {
			setState(265);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(252);
				match(T__26);
				setState(253);
				cadena();
				setState(258);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__27) {
					{
					{
					setState(254);
					match(T__27);
					setState(255);
					cadena();
					}
					}
					setState(260);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(261);
				match(T__28);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(263);
				match(T__26);
				setState(264);
				match(T__28);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CadenaContext extends ParserRuleContext {
		public TerminalNode STRING() { return getToken(AspidaParser.STRING, 0); }
		public CadenaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cadena; }
	}

	public final CadenaContext cadena() throws RecognitionException {
		CadenaContext _localctx = new CadenaContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_cadena);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(267);
			match(STRING);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ArrayContext extends ParserRuleContext {
		public List<ValueContext> value() {
			return getRuleContexts(ValueContext.class);
		}
		public ValueContext value(int i) {
			return getRuleContext(ValueContext.class,i);
		}
		public ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_array; }
	}

	public final ArrayContext array() throws RecognitionException {
		ArrayContext _localctx = new ArrayContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_array);
		int _la;
		try {
			setState(282);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(269);
				match(T__26);
				setState(270);
				value();
				setState(275);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__27) {
					{
					{
					setState(271);
					match(T__27);
					setState(272);
					value();
					}
					}
					setState(277);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(278);
				match(T__28);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(280);
				match(T__26);
				setState(281);
				match(T__28);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\63\u011f\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13"+
		"\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\3\2\3\2\3\2\5\2:\n\2\3\2\5\2=\n\2\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\4\3\4\3\4\3\4\3\4\3\5\3\5\3\5\3\5\3\5\3\5\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\7\3\7\7\7X\n\7\f\7\16\7[\13\7\3\b\3\b\3\b\3\b\3\b\3\b\3\b\3\b"+
		"\3\b\5\bf\n\b\3\t\3\t\3\t\3\t\3\t\3\t\5\tn\n\t\3\n\3\n\3\n\3\n\3\n\3\n"+
		"\5\nv\n\n\3\13\3\13\3\13\3\13\3\13\3\13\5\13~\n\13\3\f\3\f\7\f\u0082\n"+
		"\f\f\f\16\f\u0085\13\f\3\f\5\f\u0088\n\f\3\r\3\r\3\r\3\r\5\r\u008e\n\r"+
		"\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u009a\n\16\3\17"+
		"\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\3\17\5\17\u00a6\n\17\3\20\3\20"+
		"\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\5\20\u00b2\n\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\3\21\3\21\3\21\5\21\u00be\n\21\3\22\3\22\7\22\u00c2"+
		"\n\22\f\22\16\22\u00c5\13\22\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3"+
		"\23\3\23\3\23\3\23\5\23\u00d3\n\23\3\24\3\24\3\24\3\24\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\7\24\u00e0\n\24\f\24\16\24\u00e3\13\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\25\3\25\3\26\3\26\3\26\3\26\7\26\u00f0\n\26\f\26\16"+
		"\26\u00f3\13\26\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30\5\30\u00fd\n\30"+
		"\3\31\3\31\3\31\3\31\7\31\u0103\n\31\f\31\16\31\u0106\13\31\3\31\3\31"+
		"\3\31\3\31\5\31\u010c\n\31\3\32\3\32\3\33\3\33\3\33\3\33\7\33\u0114\n"+
		"\33\f\33\16\33\u0117\13\33\3\33\3\33\3\33\3\33\5\33\u011d\n\33\3\33\2"+
		"\2\34\2\4\6\b\n\f\16\20\22\24\26\30\32\34\36 \"$&(*,.\60\62\64\2\4\3\2"+
		"*,\4\2\24\31\60\63\2\u0122\2\66\3\2\2\2\4>\3\2\2\2\6D\3\2\2\2\bI\3\2\2"+
		"\2\nO\3\2\2\2\fU\3\2\2\2\16e\3\2\2\2\20m\3\2\2\2\22u\3\2\2\2\24}\3\2\2"+
		"\2\26\u0087\3\2\2\2\30\u008d\3\2\2\2\32\u0099\3\2\2\2\34\u00a5\3\2\2\2"+
		"\36\u00b1\3\2\2\2 \u00bd\3\2\2\2\"\u00bf\3\2\2\2$\u00d2\3\2\2\2&\u00d4"+
		"\3\2\2\2(\u00e9\3\2\2\2*\u00eb\3\2\2\2,\u00f4\3\2\2\2.\u00fc\3\2\2\2\60"+
		"\u010b\3\2\2\2\62\u010d\3\2\2\2\64\u011c\3\2\2\2\66\67\5\4\3\2\679\5\6"+
		"\4\28:\5\b\5\298\3\2\2\29:\3\2\2\2:<\3\2\2\2;=\5\n\6\2<;\3\2\2\2<=\3\2"+
		"\2\2=\3\3\2\2\2>?\7&\2\2?@\7\3\2\2@A\7\4\2\2AB\5\f\7\2BC\7\5\2\2C\5\3"+
		"\2\2\2DE\7\'\2\2EF\7\3\2\2FG\7#\2\2GH\7%\2\2H\7\3\2\2\2IJ\7(\2\2JK\7\3"+
		"\2\2KL\7\4\2\2LM\5\26\f\2MN\7\5\2\2N\t\3\2\2\2OP\7)\2\2PQ\7\3\2\2QR\7"+
		"\4\2\2RS\5\"\22\2ST\7\5\2\2T\13\3\2\2\2UY\5\16\b\2VX\5\16\b\2WV\3\2\2"+
		"\2X[\3\2\2\2YW\3\2\2\2YZ\3\2\2\2Z\r\3\2\2\2[Y\3\2\2\2\\]\5\20\t\2]^\7"+
		"%\2\2^f\3\2\2\2_`\5\22\n\2`a\7%\2\2af\3\2\2\2bc\5\24\13\2cd\7%\2\2df\3"+
		"\2\2\2e\\\3\2\2\2e_\3\2\2\2eb\3\2\2\2f\17\3\2\2\2gh\7\6\2\2hi\7\3\2\2"+
		"in\7#\2\2jk\7\7\2\2kl\7\3\2\2ln\7#\2\2mg\3\2\2\2mj\3\2\2\2n\21\3\2\2\2"+
		"op\7\b\2\2pq\7\3\2\2qv\t\2\2\2rs\7\t\2\2st\7\3\2\2tv\t\2\2\2uo\3\2\2\2"+
		"ur\3\2\2\2v\23\3\2\2\2wx\7\n\2\2xy\7\3\2\2y~\7#\2\2z{\7\13\2\2{|\7\3\2"+
		"\2|~\7#\2\2}w\3\2\2\2}z\3\2\2\2~\25\3\2\2\2\177\u0083\5\30\r\2\u0080\u0082"+
		"\5\30\r\2\u0081\u0080\3\2\2\2\u0082\u0085\3\2\2\2\u0083\u0081\3\2\2\2"+
		"\u0083\u0084\3\2\2\2\u0084\u0088\3\2\2\2\u0085\u0083\3\2\2\2\u0086\u0088"+
		"\5&\24\2\u0087\177\3\2\2\2\u0087\u0086\3\2\2\2\u0088\27\3\2\2\2\u0089"+
		"\u008e\5\32\16\2\u008a\u008e\5\34\17\2\u008b\u008e\5\36\20\2\u008c\u008e"+
		"\5 \21\2\u008d\u0089\3\2\2\2\u008d\u008a\3\2\2\2\u008d\u008b\3\2\2\2\u008d"+
		"\u008c\3\2\2\2\u008e\31\3\2\2\2\u008f\u0090\7\f\2\2\u0090\u0091\7\3\2"+
		"\2\u0091\u0092\5\60\31\2\u0092\u0093\7%\2\2\u0093\u009a\3\2\2\2\u0094"+
		"\u0095\7\r\2\2\u0095\u0096\7\3\2\2\u0096\u0097\5\60\31\2\u0097\u0098\7"+
		"%\2\2\u0098\u009a\3\2\2\2\u0099\u008f\3\2\2\2\u0099\u0094\3\2\2\2\u009a"+
		"\33\3\2\2\2\u009b\u009c\7\16\2\2\u009c\u009d\7\3\2\2\u009d\u009e\5\60"+
		"\31\2\u009e\u009f\7%\2\2\u009f\u00a6\3\2\2\2\u00a0\u00a1\7\17\2\2\u00a1"+
		"\u00a2\7\3\2\2\u00a2\u00a3\5\60\31\2\u00a3\u00a4\7%\2\2\u00a4\u00a6\3"+
		"\2\2\2\u00a5\u009b\3\2\2\2\u00a5\u00a0\3\2\2\2\u00a6\35\3\2\2\2\u00a7"+
		"\u00a8\7\20\2\2\u00a8\u00a9\7\3\2\2\u00a9\u00aa\5\60\31\2\u00aa\u00ab"+
		"\7%\2\2\u00ab\u00b2\3\2\2\2\u00ac\u00ad\7\21\2\2\u00ad\u00ae\7\3\2\2\u00ae"+
		"\u00af\5\60\31\2\u00af\u00b0\7%\2\2\u00b0\u00b2\3\2\2\2\u00b1\u00a7\3"+
		"\2\2\2\u00b1\u00ac\3\2\2\2\u00b2\37\3\2\2\2\u00b3\u00b4\7\22\2\2\u00b4"+
		"\u00b5\7\3\2\2\u00b5\u00b6\5\60\31\2\u00b6\u00b7\7%\2\2\u00b7\u00be\3"+
		"\2\2\2\u00b8\u00b9\7\23\2\2\u00b9\u00ba\7\3\2\2\u00ba\u00bb\5\60\31\2"+
		"\u00bb\u00bc\7%\2\2\u00bc\u00be\3\2\2\2\u00bd\u00b3\3\2\2\2\u00bd\u00b8"+
		"\3\2\2\2\u00be!\3\2\2\2\u00bf\u00c3\5$\23\2\u00c0\u00c2\5$\23\2\u00c1"+
		"\u00c0\3\2\2\2\u00c2\u00c5\3\2\2\2\u00c3\u00c1\3\2\2\2\u00c3\u00c4\3\2"+
		"\2\2\u00c4#\3\2\2\2\u00c5\u00c3\3\2\2\2\u00c6\u00c7\7#\2\2\u00c7\u00c8"+
		"\7\3\2\2\u00c8\u00c9\5.\30\2\u00c9\u00ca\7%\2\2\u00ca\u00d3\3\2\2\2\u00cb"+
		"\u00cc\7#\2\2\u00cc\u00cd\7\3\2\2\u00cd\u00ce\7\4\2\2\u00ce\u00cf\5\""+
		"\22\2\u00cf\u00d0\7\5\2\2\u00d0\u00d1\7%\2\2\u00d1\u00d3\3\2\2\2\u00d2"+
		"\u00c6\3\2\2\2\u00d2\u00cb\3\2\2\2\u00d3%\3\2\2\2\u00d4\u00d5\7-\2\2\u00d5"+
		"\u00d6\5(\25\2\u00d6\u00d7\7\4\2\2\u00d7\u00d8\5\26\f\2\u00d8\u00e1\7"+
		"\5\2\2\u00d9\u00da\7.\2\2\u00da\u00db\5(\25\2\u00db\u00dc\7\4\2\2\u00dc"+
		"\u00dd\5\26\f\2\u00dd\u00de\7\5\2\2\u00de\u00e0\3\2\2\2\u00df\u00d9\3"+
		"\2\2\2\u00e0\u00e3\3\2\2\2\u00e1\u00df\3\2\2\2\u00e1\u00e2\3\2\2\2\u00e2"+
		"\u00e4\3\2\2\2\u00e3\u00e1\3\2\2\2\u00e4\u00e5\7/\2\2\u00e5\u00e6\7\4"+
		"\2\2\u00e6\u00e7\5\26\f\2\u00e7\u00e8\7\5\2\2\u00e8\'\3\2\2\2\u00e9\u00ea"+
		"\5*\26\2\u00ea)\3\2\2\2\u00eb\u00f1\5.\30\2\u00ec\u00ed\5,\27\2\u00ed"+
		"\u00ee\5.\30\2\u00ee\u00f0\3\2\2\2\u00ef\u00ec\3\2\2\2\u00f0\u00f3\3\2"+
		"\2\2\u00f1\u00ef\3\2\2\2\u00f1\u00f2\3\2\2\2\u00f2+\3\2\2\2\u00f3\u00f1"+
		"\3\2\2\2\u00f4\u00f5\t\3\2\2\u00f5-\3\2\2\2\u00f6\u00fd\7#\2\2\u00f7\u00fd"+
		"\7$\2\2\u00f8\u00fd\7\32\2\2\u00f9\u00fd\7\33\2\2\u00fa\u00fd\7\34\2\2"+
		"\u00fb\u00fd\5\64\33\2\u00fc\u00f6\3\2\2\2\u00fc\u00f7\3\2\2\2\u00fc\u00f8"+
		"\3\2\2\2\u00fc\u00f9\3\2\2\2\u00fc\u00fa\3\2\2\2\u00fc\u00fb\3\2\2\2\u00fd"+
		"/\3\2\2\2\u00fe\u00ff\7\35\2\2\u00ff\u0104\5\62\32\2\u0100\u0101\7\36"+
		"\2\2\u0101\u0103\5\62\32\2\u0102\u0100\3\2\2\2\u0103\u0106\3\2\2\2\u0104"+
		"\u0102\3\2\2\2\u0104\u0105\3\2\2\2\u0105\u0107\3\2\2\2\u0106\u0104\3\2"+
		"\2\2\u0107\u0108\7\37\2\2\u0108\u010c\3\2\2\2\u0109\u010a\7\35\2\2\u010a"+
		"\u010c\7\37\2\2\u010b\u00fe\3\2\2\2\u010b\u0109\3\2\2\2\u010c\61\3\2\2"+
		"\2\u010d\u010e\7#\2\2\u010e\63\3\2\2\2\u010f\u0110\7\35\2\2\u0110\u0115"+
		"\5.\30\2\u0111\u0112\7\36\2\2\u0112\u0114\5.\30\2\u0113\u0111\3\2\2\2"+
		"\u0114\u0117\3\2\2\2\u0115\u0113\3\2\2\2\u0115\u0116\3\2\2\2\u0116\u0118"+
		"\3\2\2\2\u0117\u0115\3\2\2\2\u0118\u0119\7\37\2\2\u0119\u011d\3\2\2\2"+
		"\u011a\u011b\7\35\2\2\u011b\u011d\7\37\2\2\u011c\u010f\3\2\2\2\u011c\u011a"+
		"\3\2\2\2\u011d\65\3\2\2\2\319<Yemu}\u0083\u0087\u008d\u0099\u00a5\u00b1"+
		"\u00bd\u00c3\u00d2\u00e1\u00f1\u00fc\u0104\u010b\u0115\u011c";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}