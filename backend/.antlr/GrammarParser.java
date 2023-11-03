// Generated from /media/keviin/Almacenamiento/SEGUNDO SEMESTRE/COMPI 2/LAB/PROYECTOS/PROYECTO 2/OLC2_P2_202113553/backend/Grammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class GrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, INT=34, FLOAT=35, STRING=36, BOOL=37, CHARACTER=38, 
		NIL=39, VAR=40, LET=41, FUNC=42, INOUT=43, REFERENCE=44, PRINT=45, IF=46, 
		ELSE=47, SWITCH=48, CASE=49, DEFAULT=50, WHILE=51, FOR=52, IN=53, GUARD=54, 
		CONTINUE=55, BREAK=56, RETURN=57, SEMICOLON=58, COLON=59, EQUAL=60, QUESTION_MARK=61, 
		INCREMENT=62, DECREMENT=63, UNDERSCORE=64, DIGIT=65, STR=66, CHAR=67, 
		ID=68, WHITESPACE=69, MULTI_COMMENT=70, LINE_COMMENT=71;
	public static final int
		RULE_start = 0, RULE_block = 1, RULE_stmts = 2, RULE_varDeclaration = 3, 
		RULE_varAssignment = 4, RULE_letDeclaration = 5, RULE_vectorDeclaration = 6, 
		RULE_valuesVectorDeclaration = 7, RULE_vectorType = 8, RULE_vectorAppend = 9, 
		RULE_vectorRemoveLast = 10, RULE_vectorRemoveAt = 11, RULE_matrixDeclaration = 12, 
		RULE_valuesMatrixDeclaration = 13, RULE_structDeclaration = 14, RULE_structAttributes = 15, 
		RULE_structInstance = 16, RULE_attributesCall = 17, RULE_structAccess = 18, 
		RULE_structAssignment = 19, RULE_funcDeclaration = 20, RULE_paramsFuncDeclaration = 21, 
		RULE_funcCall = 22, RULE_paramsFuncCall = 23, RULE_printStmt = 24, RULE_ifStmt = 25, 
		RULE_switchStmt = 26, RULE_cases = 27, RULE_caseStmt = 28, RULE_defaultStmt = 29, 
		RULE_whileStmt = 30, RULE_forStmt = 31, RULE_range = 32, RULE_guardStmt = 33, 
		RULE_breakStmt = 34, RULE_continueStmt = 35, RULE_returnStmt = 36, RULE_expr = 37, 
		RULE_type = 38;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "block", "stmts", "varDeclaration", "varAssignment", "letDeclaration", 
			"vectorDeclaration", "valuesVectorDeclaration", "vectorType", "vectorAppend", 
			"vectorRemoveLast", "vectorRemoveAt", "matrixDeclaration", "valuesMatrixDeclaration", 
			"structDeclaration", "structAttributes", "structInstance", "attributesCall", 
			"structAccess", "structAssignment", "funcDeclaration", "paramsFuncDeclaration", 
			"funcCall", "paramsFuncCall", "printStmt", "ifStmt", "switchStmt", "cases", 
			"caseStmt", "defaultStmt", "whileStmt", "forStmt", "range", "guardStmt", 
			"breakStmt", "continueStmt", "returnStmt", "expr", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'['", "']'", "','", "'.append'", "'('", "')'", "'.removeLast'", 
			"'.remove'", "'at'", "'struct'", "'{'", "'}'", "'.'", "'->'", "'...'", 
			"'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='", "'>'", "'<='", "'<'", 
			"'=='", "'!='", "'&&'", "'||'", "'true'", "'false'", "'.count'", "'.isEmpty'", 
			"'Int'", "'Float'", "'String'", "'Bool'", "'Character'", "'nil'", "'var'", 
			"'let'", "'func'", "'inout'", "'&'", "'print'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'while'", "'for'", "'in'", "'guard'", "'continue'", 
			"'break'", "'return'", "';'", "':'", "'='", "'?'", "'+='", "'-='", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, "INT", "FLOAT", 
			"STRING", "BOOL", "CHARACTER", "NIL", "VAR", "LET", "FUNC", "INOUT", 
			"REFERENCE", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", "WHILE", 
			"FOR", "IN", "GUARD", "CONTINUE", "BREAK", "RETURN", "SEMICOLON", "COLON", 
			"EQUAL", "QUESTION_MARK", "INCREMENT", "DECREMENT", "UNDERSCORE", "DIGIT", 
			"STR", "CHAR", "ID", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT"
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
	public String getGrammarFileName() { return "Grammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public GrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StartContext extends ParserRuleContext {
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode EOF() { return getToken(GrammarParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			block();
			setState(79);
			match(EOF);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public List<StmtsContext> stmts() {
			return getRuleContexts(StmtsContext.class);
		}
		public StmtsContext stmts(int i) {
			return getRuleContext(StmtsContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_block);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(84);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(81);
					stmts();
					}
					} 
				}
				setState(86);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StmtsContext extends ParserRuleContext {
		public VarDeclarationContext varDeclaration() {
			return getRuleContext(VarDeclarationContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GrammarParser.SEMICOLON, 0); }
		public VarAssignmentContext varAssignment() {
			return getRuleContext(VarAssignmentContext.class,0);
		}
		public LetDeclarationContext letDeclaration() {
			return getRuleContext(LetDeclarationContext.class,0);
		}
		public VectorDeclarationContext vectorDeclaration() {
			return getRuleContext(VectorDeclarationContext.class,0);
		}
		public VectorAppendContext vectorAppend() {
			return getRuleContext(VectorAppendContext.class,0);
		}
		public VectorRemoveLastContext vectorRemoveLast() {
			return getRuleContext(VectorRemoveLastContext.class,0);
		}
		public VectorRemoveAtContext vectorRemoveAt() {
			return getRuleContext(VectorRemoveAtContext.class,0);
		}
		public FuncDeclarationContext funcDeclaration() {
			return getRuleContext(FuncDeclarationContext.class,0);
		}
		public FuncCallContext funcCall() {
			return getRuleContext(FuncCallContext.class,0);
		}
		public PrintStmtContext printStmt() {
			return getRuleContext(PrintStmtContext.class,0);
		}
		public IfStmtContext ifStmt() {
			return getRuleContext(IfStmtContext.class,0);
		}
		public SwitchStmtContext switchStmt() {
			return getRuleContext(SwitchStmtContext.class,0);
		}
		public WhileStmtContext whileStmt() {
			return getRuleContext(WhileStmtContext.class,0);
		}
		public ForStmtContext forStmt() {
			return getRuleContext(ForStmtContext.class,0);
		}
		public GuardStmtContext guardStmt() {
			return getRuleContext(GuardStmtContext.class,0);
		}
		public BreakStmtContext breakStmt() {
			return getRuleContext(BreakStmtContext.class,0);
		}
		public ContinueStmtContext continueStmt() {
			return getRuleContext(ContinueStmtContext.class,0);
		}
		public ReturnStmtContext returnStmt() {
			return getRuleContext(ReturnStmtContext.class,0);
		}
		public StructDeclarationContext structDeclaration() {
			return getRuleContext(StructDeclarationContext.class,0);
		}
		public StructInstanceContext structInstance() {
			return getRuleContext(StructInstanceContext.class,0);
		}
		public StructAccessContext structAccess() {
			return getRuleContext(StructAccessContext.class,0);
		}
		public StructAssignmentContext structAssignment() {
			return getRuleContext(StructAssignmentContext.class,0);
		}
		public StmtsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmts; }
	}

	public final StmtsContext stmts() throws RecognitionException {
		StmtsContext _localctx = new StmtsContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_stmts);
		int _la;
		try {
			setState(175);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(87);
				varDeclaration();
				setState(89);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(88);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(91);
				varAssignment();
				setState(93);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(92);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(95);
				letDeclaration();
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(96);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(99);
				vectorDeclaration();
				setState(101);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(100);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(103);
				vectorAppend();
				setState(105);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(104);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(107);
				vectorRemoveLast();
				setState(109);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(108);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(111);
				vectorRemoveAt();
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(112);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(115);
				funcDeclaration();
				setState(117);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(116);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(119);
				funcCall();
				setState(121);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(120);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(123);
				printStmt();
				setState(125);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(124);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(127);
				ifStmt();
				setState(129);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(128);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(131);
				switchStmt();
				setState(133);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(132);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(135);
				whileStmt();
				setState(137);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(136);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(139);
				forStmt();
				setState(141);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(140);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(143);
				guardStmt();
				setState(145);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(144);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(147);
				breakStmt();
				setState(149);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(148);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(151);
				continueStmt();
				setState(153);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(152);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(155);
				returnStmt();
				setState(157);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(156);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(159);
				structDeclaration();
				setState(161);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(160);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(163);
				structInstance();
				setState(165);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(164);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(167);
				structAccess();
				setState(169);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(168);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(171);
				structAssignment();
				setState(173);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(172);
					match(SEMICOLON);
					}
				}

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

	@SuppressWarnings("CheckReturnValue")
	public static class VarDeclarationContext extends ParserRuleContext {
		public VarDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDeclaration; }
	 
		public VarDeclarationContext() { }
		public void copyFrom(VarDeclarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeVarDeclarationContext extends VarDeclarationContext {
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode QUESTION_MARK() { return getToken(GrammarParser.QUESTION_MARK, 0); }
		public TypeVarDeclarationContext(VarDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeValueVarDeclarationContext extends VarDeclarationContext {
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TypeValueVarDeclarationContext(VarDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValueVarDeclarationContext extends VarDeclarationContext {
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ValueVarDeclarationContext(VarDeclarationContext ctx) { copyFrom(ctx); }
	}

	public final VarDeclarationContext varDeclaration() throws RecognitionException {
		VarDeclarationContext _localctx = new VarDeclarationContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_varDeclaration);
		try {
			setState(194);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new TypeValueVarDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(177);
				match(VAR);
				setState(178);
				match(ID);
				setState(179);
				match(COLON);
				setState(180);
				type();
				setState(181);
				match(EQUAL);
				setState(182);
				expr(0);
				}
				break;
			case 2:
				_localctx = new ValueVarDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(184);
				match(VAR);
				setState(185);
				match(ID);
				setState(186);
				match(EQUAL);
				setState(187);
				expr(0);
				}
				break;
			case 3:
				_localctx = new TypeVarDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(188);
				match(VAR);
				setState(189);
				match(ID);
				setState(190);
				match(COLON);
				setState(191);
				type();
				setState(192);
				match(QUESTION_MARK);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VarAssignmentContext extends ParserRuleContext {
		public VarAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varAssignment; }
	 
		public VarAssignmentContext() { }
		public void copyFrom(VarAssignmentContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DecrementVarAssignmentContext extends VarAssignmentContext {
		public TerminalNode DECREMENT() { return getToken(GrammarParser.DECREMENT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public DecrementVarAssignmentContext(VarAssignmentContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValueVarAssignmentContext extends VarAssignmentContext {
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ValueVarAssignmentContext(VarAssignmentContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncrementVarAssignmentContext extends VarAssignmentContext {
		public TerminalNode INCREMENT() { return getToken(GrammarParser.INCREMENT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public IncrementVarAssignmentContext(VarAssignmentContext ctx) { copyFrom(ctx); }
	}

	public final VarAssignmentContext varAssignment() throws RecognitionException {
		VarAssignmentContext _localctx = new VarAssignmentContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_varAssignment);
		try {
			setState(226);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
			case 1:
				_localctx = new ValueVarAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(202);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
				case 1:
					{
					setState(196);
					match(ID);
					}
					break;
				case 2:
					{
					setState(197);
					match(ID);
					setState(198);
					match(T__0);
					setState(199);
					expr(0);
					setState(200);
					match(T__1);
					}
					break;
				}
				setState(204);
				match(EQUAL);
				setState(205);
				expr(0);
				}
				break;
			case 2:
				_localctx = new IncrementVarAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(212);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,26,_ctx) ) {
				case 1:
					{
					setState(206);
					match(ID);
					}
					break;
				case 2:
					{
					setState(207);
					match(ID);
					setState(208);
					match(T__0);
					setState(209);
					expr(0);
					setState(210);
					match(T__1);
					}
					break;
				}
				setState(214);
				match(INCREMENT);
				setState(215);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DecrementVarAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(222);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
				case 1:
					{
					setState(216);
					match(ID);
					}
					break;
				case 2:
					{
					setState(217);
					match(ID);
					setState(218);
					match(T__0);
					setState(219);
					expr(0);
					setState(220);
					match(T__1);
					}
					break;
				}
				setState(224);
				match(DECREMENT);
				setState(225);
				expr(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class LetDeclarationContext extends ParserRuleContext {
		public LetDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_letDeclaration; }
	 
		public LetDeclarationContext() { }
		public void copyFrom(LetDeclarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValueLetDeclarationContext extends LetDeclarationContext {
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ValueLetDeclarationContext(LetDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeValueLetDeclarationContext extends LetDeclarationContext {
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TypeValueLetDeclarationContext(LetDeclarationContext ctx) { copyFrom(ctx); }
	}

	public final LetDeclarationContext letDeclaration() throws RecognitionException {
		LetDeclarationContext _localctx = new LetDeclarationContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_letDeclaration);
		try {
			setState(239);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,29,_ctx) ) {
			case 1:
				_localctx = new TypeValueLetDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(228);
				match(LET);
				setState(229);
				match(ID);
				setState(230);
				match(COLON);
				setState(231);
				type();
				setState(232);
				match(EQUAL);
				setState(233);
				expr(0);
				}
				break;
			case 2:
				_localctx = new ValueLetDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(235);
				match(LET);
				setState(236);
				match(ID);
				setState(237);
				match(EQUAL);
				setState(238);
				expr(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VectorDeclarationContext extends ParserRuleContext {
		public VectorDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorDeclaration; }
	 
		public VectorDeclarationContext() { }
		public void copyFrom(VectorDeclarationContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorIdDeclarationContext extends VectorDeclarationContext {
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public VectorIdDeclarationContext(VectorDeclarationContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorValuesDeclarationContext extends VectorDeclarationContext {
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ValuesVectorDeclarationContext valuesVectorDeclaration() {
			return getRuleContext(ValuesVectorDeclarationContext.class,0);
		}
		public VectorValuesDeclarationContext(VectorDeclarationContext ctx) { copyFrom(ctx); }
	}

	public final VectorDeclarationContext vectorDeclaration() throws RecognitionException {
		VectorDeclarationContext _localctx = new VectorDeclarationContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_vectorDeclaration);
		int _la;
		try {
			setState(263);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				_localctx = new VectorValuesDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(241);
				match(VAR);
				setState(242);
				match(ID);
				setState(243);
				match(COLON);
				setState(244);
				match(T__0);
				setState(245);
				type();
				setState(246);
				match(T__1);
				setState(247);
				match(EQUAL);
				setState(248);
				match(T__0);
				setState(250);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & -1152921483568211967L) != 0)) {
					{
					setState(249);
					valuesVectorDeclaration();
					}
				}

				setState(252);
				match(T__1);
				}
				break;
			case 2:
				_localctx = new VectorIdDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(254);
				match(VAR);
				setState(255);
				match(ID);
				setState(256);
				match(COLON);
				setState(257);
				match(T__0);
				setState(258);
				type();
				setState(259);
				match(T__1);
				setState(260);
				match(EQUAL);
				setState(261);
				match(ID);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ValuesVectorDeclarationContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ValuesVectorDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valuesVectorDeclaration; }
	}

	public final ValuesVectorDeclarationContext valuesVectorDeclaration() throws RecognitionException {
		ValuesVectorDeclarationContext _localctx = new ValuesVectorDeclarationContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_valuesVectorDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			expr(0);
			setState(270);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(266);
				match(T__2);
				setState(267);
				expr(0);
				}
				}
				setState(272);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VectorTypeContext extends ParserRuleContext {
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public VectorTypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorType; }
	}

	public final VectorTypeContext vectorType() throws RecognitionException {
		VectorTypeContext _localctx = new VectorTypeContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_vectorType);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(273);
			match(T__0);
			setState(274);
			type();
			setState(275);
			match(T__1);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VectorAppendContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VectorAppendContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorAppend; }
	}

	public final VectorAppendContext vectorAppend() throws RecognitionException {
		VectorAppendContext _localctx = new VectorAppendContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_vectorAppend);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(277);
			match(ID);
			setState(278);
			match(T__3);
			setState(279);
			match(T__4);
			setState(280);
			expr(0);
			setState(281);
			match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VectorRemoveLastContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public VectorRemoveLastContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorRemoveLast; }
	}

	public final VectorRemoveLastContext vectorRemoveLast() throws RecognitionException {
		VectorRemoveLastContext _localctx = new VectorRemoveLastContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_vectorRemoveLast);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(283);
			match(ID);
			setState(284);
			match(T__6);
			setState(285);
			match(T__4);
			setState(286);
			match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class VectorRemoveAtContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VectorRemoveAtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectorRemoveAt; }
	}

	public final VectorRemoveAtContext vectorRemoveAt() throws RecognitionException {
		VectorRemoveAtContext _localctx = new VectorRemoveAtContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_vectorRemoveAt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			match(ID);
			setState(289);
			match(T__7);
			setState(290);
			match(T__4);
			setState(291);
			match(T__8);
			setState(292);
			match(COLON);
			setState(293);
			expr(0);
			setState(294);
			match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class MatrixDeclarationContext extends ParserRuleContext {
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ValuesMatrixDeclarationContext valuesMatrixDeclaration() {
			return getRuleContext(ValuesMatrixDeclarationContext.class,0);
		}
		public MatrixDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrixDeclaration; }
	}

	public final MatrixDeclarationContext matrixDeclaration() throws RecognitionException {
		MatrixDeclarationContext _localctx = new MatrixDeclarationContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_matrixDeclaration);
		int _la;
		try {
			setState(322);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(296);
				match(VAR);
				setState(297);
				match(ID);
				setState(298);
				match(COLON);
				setState(299);
				match(T__0);
				setState(300);
				match(T__0);
				setState(301);
				type();
				setState(302);
				match(T__1);
				setState(303);
				match(T__1);
				setState(304);
				match(EQUAL);
				setState(305);
				match(T__0);
				setState(307);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(306);
					valuesMatrixDeclaration();
					}
				}

				setState(309);
				match(T__1);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(311);
				match(VAR);
				setState(312);
				match(ID);
				setState(313);
				match(COLON);
				setState(314);
				match(T__0);
				setState(315);
				match(T__0);
				setState(316);
				type();
				setState(317);
				match(T__1);
				setState(318);
				match(T__1);
				setState(319);
				match(EQUAL);
				setState(320);
				match(ID);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ValuesMatrixDeclarationContext extends ParserRuleContext {
		public List<ValuesVectorDeclarationContext> valuesVectorDeclaration() {
			return getRuleContexts(ValuesVectorDeclarationContext.class);
		}
		public ValuesVectorDeclarationContext valuesVectorDeclaration(int i) {
			return getRuleContext(ValuesVectorDeclarationContext.class,i);
		}
		public ValuesMatrixDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valuesMatrixDeclaration; }
	}

	public final ValuesMatrixDeclarationContext valuesMatrixDeclaration() throws RecognitionException {
		ValuesMatrixDeclarationContext _localctx = new ValuesMatrixDeclarationContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_valuesMatrixDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(324);
			match(T__0);
			setState(326);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & -1152921483568211967L) != 0)) {
				{
				setState(325);
				valuesVectorDeclaration();
				}
			}

			setState(328);
			match(T__1);
			setState(337);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(329);
				match(T__2);
				setState(330);
				match(T__0);
				setState(332);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & -1152921483568211967L) != 0)) {
					{
					setState(331);
					valuesVectorDeclaration();
					}
				}

				setState(334);
				match(T__1);
				}
				}
				setState(339);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructDeclarationContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public List<StructAttributesContext> structAttributes() {
			return getRuleContexts(StructAttributesContext.class);
		}
		public StructAttributesContext structAttributes(int i) {
			return getRuleContext(StructAttributesContext.class,i);
		}
		public StructDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDeclaration; }
	}

	public final StructDeclarationContext structDeclaration() throws RecognitionException {
		StructDeclarationContext _localctx = new StructDeclarationContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_structDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(340);
			match(T__9);
			setState(341);
			match(ID);
			setState(342);
			match(T__10);
			setState(346);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==VAR || _la==LET) {
				{
				{
				setState(343);
				structAttributes();
				}
				}
				setState(348);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(349);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructAttributesContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public StructAttributesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structAttributes; }
	}

	public final StructAttributesContext structAttributes() throws RecognitionException {
		StructAttributesContext _localctx = new StructAttributesContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_structAttributes);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(351);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(352);
			match(ID);
			setState(353);
			match(COLON);
			setState(354);
			type();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructInstanceContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public AttributesCallContext attributesCall() {
			return getRuleContext(AttributesCallContext.class,0);
		}
		public StructInstanceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structInstance; }
	}

	public final StructInstanceContext structInstance() throws RecognitionException {
		StructInstanceContext _localctx = new StructInstanceContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_structInstance);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(356);
			match(ID);
			setState(357);
			match(T__4);
			setState(359);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(358);
				attributesCall();
				}
			}

			setState(361);
			match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class AttributesCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AttributesCallContext attributesCall() {
			return getRuleContext(AttributesCallContext.class,0);
		}
		public AttributesCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_attributesCall; }
	}

	public final AttributesCallContext attributesCall() throws RecognitionException {
		AttributesCallContext _localctx = new AttributesCallContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_attributesCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(ID);
			setState(364);
			match(COLON);
			setState(365);
			expr(0);
			setState(368);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__2) {
				{
				setState(366);
				match(T__2);
				setState(367);
				attributesCall();
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructAccessContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public StructAccessContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structAccess; }
	}

	public final StructAccessContext structAccess() throws RecognitionException {
		StructAccessContext _localctx = new StructAccessContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_structAccess);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			match(ID);
			setState(371);
			match(T__12);
			setState(372);
			match(ID);
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

	@SuppressWarnings("CheckReturnValue")
	public static class StructAssignmentContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StructAssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structAssignment; }
	}

	public final StructAssignmentContext structAssignment() throws RecognitionException {
		StructAssignmentContext _localctx = new StructAssignmentContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_structAssignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(374);
			match(ID);
			setState(375);
			match(T__12);
			setState(376);
			match(ID);
			setState(377);
			match(EQUAL);
			setState(378);
			expr(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncDeclarationContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(GrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ParamsFuncDeclarationContext paramsFuncDeclaration() {
			return getRuleContext(ParamsFuncDeclarationContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public FuncDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDeclaration; }
	}

	public final FuncDeclarationContext funcDeclaration() throws RecognitionException {
		FuncDeclarationContext _localctx = new FuncDeclarationContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_funcDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(380);
			match(FUNC);
			setState(381);
			match(ID);
			setState(382);
			match(T__4);
			setState(384);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==UNDERSCORE || _la==ID) {
				{
				setState(383);
				paramsFuncDeclaration();
				}
			}

			setState(386);
			match(T__5);
			setState(389);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__13) {
				{
				setState(387);
				match(T__13);
				setState(388);
				type();
				}
			}

			setState(391);
			match(T__10);
			setState(392);
			block();
			setState(393);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParamsFuncDeclarationContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(GrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(GrammarParser.ID, i);
		}
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public VectorTypeContext vectorType() {
			return getRuleContext(VectorTypeContext.class,0);
		}
		public TerminalNode INOUT() { return getToken(GrammarParser.INOUT, 0); }
		public ParamsFuncDeclarationContext paramsFuncDeclaration() {
			return getRuleContext(ParamsFuncDeclarationContext.class,0);
		}
		public TerminalNode UNDERSCORE() { return getToken(GrammarParser.UNDERSCORE, 0); }
		public ParamsFuncDeclarationContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramsFuncDeclaration; }
	}

	public final ParamsFuncDeclarationContext paramsFuncDeclaration() throws RecognitionException {
		ParamsFuncDeclarationContext _localctx = new ParamsFuncDeclarationContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_paramsFuncDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(396);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				{
				setState(395);
				_la = _input.LA(1);
				if ( !(_la==UNDERSCORE || _la==ID) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			}
			setState(398);
			match(ID);
			setState(399);
			match(COLON);
			setState(401);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INOUT) {
				{
				setState(400);
				match(INOUT);
				}
			}

			setState(405);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
			case FLOAT:
			case STRING:
			case BOOL:
			case CHARACTER:
			case ID:
				{
				setState(403);
				type();
				}
				break;
			case T__0:
				{
				setState(404);
				vectorType();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(409);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__2) {
				{
				setState(407);
				match(T__2);
				setState(408);
				paramsFuncDeclaration();
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

	@SuppressWarnings("CheckReturnValue")
	public static class FuncCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ParamsFuncCallContext paramsFuncCall() {
			return getRuleContext(ParamsFuncCallContext.class,0);
		}
		public FuncCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcCall; }
	}

	public final FuncCallContext funcCall() throws RecognitionException {
		FuncCallContext _localctx = new FuncCallContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_funcCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(411);
			match(ID);
			setState(412);
			match(T__4);
			setState(414);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (((((_la - 5)) & ~0x3f) == 0 && ((1L << (_la - 5)) & -1152920933812398079L) != 0)) {
				{
				setState(413);
				paramsFuncCall();
				}
			}

			setState(416);
			match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ParamsFuncCallContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TerminalNode REFERENCE() { return getToken(GrammarParser.REFERENCE, 0); }
		public ParamsFuncCallContext paramsFuncCall() {
			return getRuleContext(ParamsFuncCallContext.class,0);
		}
		public ParamsFuncCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_paramsFuncCall; }
	}

	public final ParamsFuncCallContext paramsFuncCall() throws RecognitionException {
		ParamsFuncCallContext _localctx = new ParamsFuncCallContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_paramsFuncCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(420);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				{
				setState(418);
				match(ID);
				setState(419);
				match(COLON);
				}
				break;
			}
			setState(423);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==REFERENCE) {
				{
				setState(422);
				match(REFERENCE);
				}
			}

			setState(425);
			expr(0);
			setState(428);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__2) {
				{
				setState(426);
				match(T__2);
				setState(427);
				paramsFuncCall();
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

	@SuppressWarnings("CheckReturnValue")
	public static class PrintStmtContext extends ParserRuleContext {
		public TerminalNode PRINT() { return getToken(GrammarParser.PRINT, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public PrintStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_printStmt; }
	}

	public final PrintStmtContext printStmt() throws RecognitionException {
		PrintStmtContext _localctx = new PrintStmtContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_printStmt);
		int _la;
		try {
			setState(447);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(430);
				match(PRINT);
				setState(431);
				match(T__4);
				setState(432);
				expr(0);
				setState(433);
				match(T__5);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(435);
				match(PRINT);
				setState(436);
				match(T__4);
				setState(437);
				expr(0);
				setState(442);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__2) {
					{
					{
					setState(438);
					match(T__2);
					setState(439);
					expr(0);
					}
					}
					setState(444);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(445);
				match(T__5);
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

	@SuppressWarnings("CheckReturnValue")
	public static class IfStmtContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(GrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<BlockContext> block() {
			return getRuleContexts(BlockContext.class);
		}
		public BlockContext block(int i) {
			return getRuleContext(BlockContext.class,i);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public IfStmtContext ifStmt() {
			return getRuleContext(IfStmtContext.class,0);
		}
		public IfStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStmt; }
	}

	public final IfStmtContext ifStmt() throws RecognitionException {
		IfStmtContext _localctx = new IfStmtContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_ifStmt);
		try {
			setState(473);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,53,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(449);
				match(IF);
				setState(450);
				expr(0);
				setState(451);
				match(T__10);
				setState(452);
				block();
				setState(453);
				match(T__11);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(455);
				match(IF);
				setState(456);
				expr(0);
				setState(457);
				match(T__10);
				setState(458);
				block();
				setState(459);
				match(T__11);
				setState(460);
				match(ELSE);
				setState(461);
				match(T__10);
				setState(462);
				block();
				setState(463);
				match(T__11);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(465);
				match(IF);
				setState(466);
				expr(0);
				setState(467);
				match(T__10);
				setState(468);
				block();
				setState(469);
				match(T__11);
				setState(470);
				match(ELSE);
				setState(471);
				ifStmt();
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

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(GrammarParser.SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public CasesContext cases() {
			return getRuleContext(CasesContext.class,0);
		}
		public SwitchStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchStmt; }
	}

	public final SwitchStmtContext switchStmt() throws RecognitionException {
		SwitchStmtContext _localctx = new SwitchStmtContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_switchStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(475);
			match(SWITCH);
			setState(476);
			expr(0);
			setState(477);
			match(T__10);
			setState(478);
			cases();
			setState(479);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CasesContext extends ParserRuleContext {
		public List<CaseStmtContext> caseStmt() {
			return getRuleContexts(CaseStmtContext.class);
		}
		public CaseStmtContext caseStmt(int i) {
			return getRuleContext(CaseStmtContext.class,i);
		}
		public List<DefaultStmtContext> defaultStmt() {
			return getRuleContexts(DefaultStmtContext.class);
		}
		public DefaultStmtContext defaultStmt(int i) {
			return getRuleContext(DefaultStmtContext.class,i);
		}
		public CasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cases; }
	}

	public final CasesContext cases() throws RecognitionException {
		CasesContext _localctx = new CasesContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_cases);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(485);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE || _la==DEFAULT) {
				{
				setState(483);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case CASE:
					{
					setState(481);
					caseStmt();
					}
					break;
				case DEFAULT:
					{
					setState(482);
					defaultStmt();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(487);
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

	@SuppressWarnings("CheckReturnValue")
	public static class CaseStmtContext extends ParserRuleContext {
		public TerminalNode CASE() { return getToken(GrammarParser.CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(GrammarParser.BREAK, 0); }
		public CaseStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseStmt; }
	}

	public final CaseStmtContext caseStmt() throws RecognitionException {
		CaseStmtContext _localctx = new CaseStmtContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_caseStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(488);
			match(CASE);
			setState(489);
			expr(0);
			setState(490);
			match(COLON);
			setState(491);
			block();
			setState(493);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(492);
				match(BREAK);
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

	@SuppressWarnings("CheckReturnValue")
	public static class DefaultStmtContext extends ParserRuleContext {
		public TerminalNode DEFAULT() { return getToken(GrammarParser.DEFAULT, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode BREAK() { return getToken(GrammarParser.BREAK, 0); }
		public DefaultStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultStmt; }
	}

	public final DefaultStmtContext defaultStmt() throws RecognitionException {
		DefaultStmtContext _localctx = new DefaultStmtContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_defaultStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(495);
			match(DEFAULT);
			setState(496);
			match(COLON);
			setState(497);
			block();
			setState(499);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(498);
				match(BREAK);
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

	@SuppressWarnings("CheckReturnValue")
	public static class WhileStmtContext extends ParserRuleContext {
		public TerminalNode WHILE() { return getToken(GrammarParser.WHILE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public WhileStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileStmt; }
	}

	public final WhileStmtContext whileStmt() throws RecognitionException {
		WhileStmtContext _localctx = new WhileStmtContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_whileStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(501);
			match(WHILE);
			setState(502);
			expr(0);
			setState(503);
			match(T__10);
			setState(504);
			block();
			setState(505);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends ParserRuleContext {
		public TerminalNode FOR() { return getToken(GrammarParser.FOR, 0); }
		public TerminalNode IN() { return getToken(GrammarParser.IN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode UNDERSCORE() { return getToken(GrammarParser.UNDERSCORE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public RangeContext range() {
			return getRuleContext(RangeContext.class,0);
		}
		public ForStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forStmt; }
	}

	public final ForStmtContext forStmt() throws RecognitionException {
		ForStmtContext _localctx = new ForStmtContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_forStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(507);
			match(FOR);
			setState(508);
			_la = _input.LA(1);
			if ( !(_la==UNDERSCORE || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(509);
			match(IN);
			setState(512);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
			case 1:
				{
				setState(510);
				expr(0);
				}
				break;
			case 2:
				{
				setState(511);
				range();
				}
				break;
			}
			setState(514);
			match(T__10);
			setState(515);
			block();
			setState(516);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class RangeContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_range);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(518);
			expr(0);
			setState(519);
			match(T__14);
			setState(520);
			expr(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class GuardStmtContext extends ParserRuleContext {
		public TerminalNode GUARD() { return getToken(GrammarParser.GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE() { return getToken(GrammarParser.ELSE, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public GuardStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guardStmt; }
	}

	public final GuardStmtContext guardStmt() throws RecognitionException {
		GuardStmtContext _localctx = new GuardStmtContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_guardStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(522);
			match(GUARD);
			setState(523);
			expr(0);
			setState(524);
			match(ELSE);
			setState(525);
			match(T__10);
			setState(526);
			block();
			setState(527);
			match(T__11);
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

	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends ParserRuleContext {
		public TerminalNode BREAK() { return getToken(GrammarParser.BREAK, 0); }
		public BreakStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breakStmt; }
	}

	public final BreakStmtContext breakStmt() throws RecognitionException {
		BreakStmtContext _localctx = new BreakStmtContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_breakStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(529);
			match(BREAK);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends ParserRuleContext {
		public TerminalNode CONTINUE() { return getToken(GrammarParser.CONTINUE, 0); }
		public ContinueStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continueStmt; }
	}

	public final ContinueStmtContext continueStmt() throws RecognitionException {
		ContinueStmtContext _localctx = new ContinueStmtContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_continueStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(531);
			match(CONTINUE);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends ParserRuleContext {
		public TerminalNode RETURN() { return getToken(GrammarParser.RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_returnStmt; }
	}

	public final ReturnStmtContext returnStmt() throws RecognitionException {
		ReturnStmtContext _localctx = new ReturnStmtContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_returnStmt);
		try {
			setState(536);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(533);
				match(RETURN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(534);
				match(RETURN);
				setState(535);
				expr(0);
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

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringEmbededExprContext extends ExprContext {
		public TerminalNode STRING() { return getToken(GrammarParser.STRING, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StringEmbededExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringExprContext extends ExprContext {
		public TerminalNode STR() { return getToken(GrammarParser.STR, 0); }
		public StringExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilExprContext extends ExprContext {
		public TerminalNode NIL() { return getToken(GrammarParser.NIL, 0); }
		public NilExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAccessExprContext extends ExprContext {
		public StructAccessContext structAccess() {
			return getRuleContext(StructAccessContext.class,0);
		}
		public StructAccessExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public IdExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IsEmptyExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public IsEmptyExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LogicalOperationExprContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LogicalOperationExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInstanceExprContext extends ExprContext {
		public StructInstanceContext structInstance() {
			return getRuleContext(StructInstanceContext.class,0);
		}
		public StructInstanceExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatEmbededExprContext extends ExprContext {
		public TerminalNode FLOAT() { return getToken(GrammarParser.FLOAT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public FloatEmbededExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorAccessExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public VectorAccessExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GroupExprContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public GroupExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArithmeticOperationExprContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ArithmeticOperationExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ComparisonOperationExprContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ComparisonOperationExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RelationalOperationExprContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public RelationalOperationExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CharacterExprContext extends ExprContext {
		public TerminalNode CHAR() { return getToken(GrammarParser.CHAR, 0); }
		public CharacterExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DigitExprContext extends ExprContext {
		public TerminalNode DIGIT() { return getToken(GrammarParser.DIGIT, 0); }
		public DigitExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CountExprContext extends ExprContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public CountExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntEmbededExprContext extends ExprContext {
		public TerminalNode INT() { return getToken(GrammarParser.INT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public IntEmbededExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnaryOperationExprContext extends ExprContext {
		public Token op;
		public ExprContext right;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public UnaryOperationExprContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanExprContext extends ExprContext {
		public BooleanExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 74;
		enterRecursionRule(_localctx, 74, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(577);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,60,_ctx) ) {
			case 1:
				{
				_localctx = new GroupExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(539);
				match(T__4);
				setState(540);
				expr(0);
				setState(541);
				match(T__5);
				}
				break;
			case 2:
				{
				_localctx = new UnaryOperationExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(543);
				((UnaryOperationExprContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__15 || _la==T__16) ) {
					((UnaryOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(544);
				((UnaryOperationExprContext)_localctx).right = expr(21);
				}
				break;
			case 3:
				{
				_localctx = new BooleanExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(545);
				_la = _input.LA(1);
				if ( !(_la==T__29 || _la==T__30) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 4:
				{
				_localctx = new NilExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(546);
				match(NIL);
				}
				break;
			case 5:
				{
				_localctx = new DigitExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(547);
				match(DIGIT);
				}
				break;
			case 6:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(548);
				match(STR);
				}
				break;
			case 7:
				{
				_localctx = new CharacterExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(549);
				match(CHAR);
				}
				break;
			case 8:
				{
				_localctx = new CountExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(550);
				match(ID);
				setState(551);
				match(T__31);
				}
				break;
			case 9:
				{
				_localctx = new IsEmptyExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(552);
				match(ID);
				setState(553);
				match(T__32);
				}
				break;
			case 10:
				{
				_localctx = new VectorAccessExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(554);
				match(ID);
				setState(555);
				match(T__0);
				setState(556);
				expr(0);
				setState(557);
				match(T__1);
				}
				break;
			case 11:
				{
				_localctx = new StructInstanceExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(559);
				structInstance();
				}
				break;
			case 12:
				{
				_localctx = new StructAccessExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(560);
				structAccess();
				}
				break;
			case 13:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(561);
				match(ID);
				}
				break;
			case 14:
				{
				_localctx = new IntEmbededExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(562);
				match(INT);
				setState(563);
				match(T__4);
				setState(564);
				expr(0);
				setState(565);
				match(T__5);
				}
				break;
			case 15:
				{
				_localctx = new FloatEmbededExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(567);
				match(FLOAT);
				setState(568);
				match(T__4);
				setState(569);
				expr(0);
				setState(570);
				match(T__5);
				}
				break;
			case 16:
				{
				_localctx = new StringEmbededExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(572);
				match(STRING);
				setState(573);
				match(T__4);
				setState(574);
				expr(0);
				setState(575);
				match(T__5);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(599);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(597);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,61,_ctx) ) {
					case 1:
						{
						_localctx = new ArithmeticOperationExprContext(new ExprContext(_parentctx, _parentState));
						((ArithmeticOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(579);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(580);
						((ArithmeticOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1835008L) != 0)) ) {
							((ArithmeticOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(581);
						((ArithmeticOperationExprContext)_localctx).right = expr(21);
						}
						break;
					case 2:
						{
						_localctx = new ArithmeticOperationExprContext(new ExprContext(_parentctx, _parentState));
						((ArithmeticOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(582);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(583);
						((ArithmeticOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__16 || _la==T__20) ) {
							((ArithmeticOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(584);
						((ArithmeticOperationExprContext)_localctx).right = expr(20);
						}
						break;
					case 3:
						{
						_localctx = new RelationalOperationExprContext(new ExprContext(_parentctx, _parentState));
						((RelationalOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(585);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(586);
						((RelationalOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__21 || _la==T__22) ) {
							((RelationalOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(587);
						((RelationalOperationExprContext)_localctx).right = expr(19);
						}
						break;
					case 4:
						{
						_localctx = new RelationalOperationExprContext(new ExprContext(_parentctx, _parentState));
						((RelationalOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(588);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(589);
						((RelationalOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__23 || _la==T__24) ) {
							((RelationalOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(590);
						((RelationalOperationExprContext)_localctx).right = expr(18);
						}
						break;
					case 5:
						{
						_localctx = new ComparisonOperationExprContext(new ExprContext(_parentctx, _parentState));
						((ComparisonOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(591);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(592);
						((ComparisonOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__25 || _la==T__26) ) {
							((ComparisonOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(593);
						((ComparisonOperationExprContext)_localctx).right = expr(17);
						}
						break;
					case 6:
						{
						_localctx = new LogicalOperationExprContext(new ExprContext(_parentctx, _parentState));
						((LogicalOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(594);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(595);
						((LogicalOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__27 || _la==T__28) ) {
							((LogicalOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(596);
						((LogicalOperationExprContext)_localctx).right = expr(16);
						}
						break;
					}
					} 
				}
				setState(601);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode INT() { return getToken(GrammarParser.INT, 0); }
		public TerminalNode FLOAT() { return getToken(GrammarParser.FLOAT, 0); }
		public TerminalNode STRING() { return getToken(GrammarParser.STRING, 0); }
		public TerminalNode BOOL() { return getToken(GrammarParser.BOOL, 0); }
		public TerminalNode CHARACTER() { return getToken(GrammarParser.CHARACTER, 0); }
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(602);
			_la = _input.LA(1);
			if ( !(((((_la - 34)) & ~0x3f) == 0 && ((1L << (_la - 34)) & 17179869215L) != 0)) ) {
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 37:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 20);
		case 1:
			return precpred(_ctx, 19);
		case 2:
			return precpred(_ctx, 18);
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 16);
		case 5:
			return precpred(_ctx, 15);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001G\u025d\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0001\u0000\u0001\u0000"+
		"\u0001\u0000\u0001\u0001\u0005\u0001S\b\u0001\n\u0001\f\u0001V\t\u0001"+
		"\u0001\u0002\u0001\u0002\u0003\u0002Z\b\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002^\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002b\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002f\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002j\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002n\b\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002r\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"v\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002z\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002~\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0082"+
		"\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0086\b\u0002\u0001\u0002"+
		"\u0001\u0002\u0003\u0002\u008a\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"\u008e\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0092\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002\u0096\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002\u009a\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u009e\b\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002\u00a2\b\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002\u00a6\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u00aa\b"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u00ae\b\u0002\u0003\u0002\u00b0"+
		"\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003"+
		"\u0003\u00c3\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u00cb\b\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u00d5\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u00df\b\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u00e3\b\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00f0\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006\u00fb\b\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u0108\b\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0005\u0007\u010d\b\u0007\n\u0007\f\u0007\u0110\t\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0003\f\u0134\b\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003"+
		"\f\u0143\b\f\u0001\r\u0001\r\u0003\r\u0147\b\r\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0003\r\u014d\b\r\u0001\r\u0005\r\u0150\b\r\n\r\f\r\u0153\t\r"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e\u0159\b\u000e"+
		"\n\u000e\f\u000e\u015c\t\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0003\u0010\u0168\b\u0010\u0001\u0010\u0001\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0003\u0011\u0171\b\u0011\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0014\u0003\u0014\u0181\b\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0003\u0014\u0186\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0001"+
		"\u0014\u0001\u0015\u0003\u0015\u018d\b\u0015\u0001\u0015\u0001\u0015\u0001"+
		"\u0015\u0003\u0015\u0192\b\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u0196"+
		"\b\u0015\u0001\u0015\u0001\u0015\u0003\u0015\u019a\b\u0015\u0001\u0016"+
		"\u0001\u0016\u0001\u0016\u0003\u0016\u019f\b\u0016\u0001\u0016\u0001\u0016"+
		"\u0001\u0017\u0001\u0017\u0003\u0017\u01a5\b\u0017\u0001\u0017\u0003\u0017"+
		"\u01a8\b\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u01ad\b"+
		"\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0005\u0018\u01b9"+
		"\b\u0018\n\u0018\f\u0018\u01bc\t\u0018\u0001\u0018\u0001\u0018\u0003\u0018"+
		"\u01c0\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019"+
		"\u0001\u0019\u0003\u0019\u01da\b\u0019\u0001\u001a\u0001\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b\u0001\u001b\u0005\u001b"+
		"\u01e4\b\u001b\n\u001b\f\u001b\u01e7\t\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u01ee\b\u001c\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0003\u001d\u01f4\b\u001d\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u0201\b\u001f\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0001\u001f\u0001 \u0001 \u0001 \u0001"+
		" \u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001!\u0001\"\u0001\"\u0001"+
		"#\u0001#\u0001$\u0001$\u0001$\u0003$\u0219\b$\u0001%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0003%\u0242\b%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001%\u0001"+
		"%\u0001%\u0001%\u0001%\u0001%\u0001%\u0005%\u0256\b%\n%\f%\u0259\t%\u0001"+
		"&\u0001&\u0001&\u0000\u0001J\'\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJL\u0000"+
		"\u000b\u0001\u0000()\u0002\u0000@@DD\u0001\u0000\u0010\u0011\u0001\u0000"+
		"\u001e\u001f\u0001\u0000\u0012\u0014\u0002\u0000\u0011\u0011\u0015\u0015"+
		"\u0001\u0000\u0016\u0017\u0001\u0000\u0018\u0019\u0001\u0000\u001a\u001b"+
		"\u0001\u0000\u001c\u001d\u0002\u0000\"&DD\u029d\u0000N\u0001\u0000\u0000"+
		"\u0000\u0002T\u0001\u0000\u0000\u0000\u0004\u00af\u0001\u0000\u0000\u0000"+
		"\u0006\u00c2\u0001\u0000\u0000\u0000\b\u00e2\u0001\u0000\u0000\u0000\n"+
		"\u00ef\u0001\u0000\u0000\u0000\f\u0107\u0001\u0000\u0000\u0000\u000e\u0109"+
		"\u0001\u0000\u0000\u0000\u0010\u0111\u0001\u0000\u0000\u0000\u0012\u0115"+
		"\u0001\u0000\u0000\u0000\u0014\u011b\u0001\u0000\u0000\u0000\u0016\u0120"+
		"\u0001\u0000\u0000\u0000\u0018\u0142\u0001\u0000\u0000\u0000\u001a\u0144"+
		"\u0001\u0000\u0000\u0000\u001c\u0154\u0001\u0000\u0000\u0000\u001e\u015f"+
		"\u0001\u0000\u0000\u0000 \u0164\u0001\u0000\u0000\u0000\"\u016b\u0001"+
		"\u0000\u0000\u0000$\u0172\u0001\u0000\u0000\u0000&\u0176\u0001\u0000\u0000"+
		"\u0000(\u017c\u0001\u0000\u0000\u0000*\u018c\u0001\u0000\u0000\u0000,"+
		"\u019b\u0001\u0000\u0000\u0000.\u01a4\u0001\u0000\u0000\u00000\u01bf\u0001"+
		"\u0000\u0000\u00002\u01d9\u0001\u0000\u0000\u00004\u01db\u0001\u0000\u0000"+
		"\u00006\u01e5\u0001\u0000\u0000\u00008\u01e8\u0001\u0000\u0000\u0000:"+
		"\u01ef\u0001\u0000\u0000\u0000<\u01f5\u0001\u0000\u0000\u0000>\u01fb\u0001"+
		"\u0000\u0000\u0000@\u0206\u0001\u0000\u0000\u0000B\u020a\u0001\u0000\u0000"+
		"\u0000D\u0211\u0001\u0000\u0000\u0000F\u0213\u0001\u0000\u0000\u0000H"+
		"\u0218\u0001\u0000\u0000\u0000J\u0241\u0001\u0000\u0000\u0000L\u025a\u0001"+
		"\u0000\u0000\u0000NO\u0003\u0002\u0001\u0000OP\u0005\u0000\u0000\u0001"+
		"P\u0001\u0001\u0000\u0000\u0000QS\u0003\u0004\u0002\u0000RQ\u0001\u0000"+
		"\u0000\u0000SV\u0001\u0000\u0000\u0000TR\u0001\u0000\u0000\u0000TU\u0001"+
		"\u0000\u0000\u0000U\u0003\u0001\u0000\u0000\u0000VT\u0001\u0000\u0000"+
		"\u0000WY\u0003\u0006\u0003\u0000XZ\u0005:\u0000\u0000YX\u0001\u0000\u0000"+
		"\u0000YZ\u0001\u0000\u0000\u0000Z\u00b0\u0001\u0000\u0000\u0000[]\u0003"+
		"\b\u0004\u0000\\^\u0005:\u0000\u0000]\\\u0001\u0000\u0000\u0000]^\u0001"+
		"\u0000\u0000\u0000^\u00b0\u0001\u0000\u0000\u0000_a\u0003\n\u0005\u0000"+
		"`b\u0005:\u0000\u0000a`\u0001\u0000\u0000\u0000ab\u0001\u0000\u0000\u0000"+
		"b\u00b0\u0001\u0000\u0000\u0000ce\u0003\f\u0006\u0000df\u0005:\u0000\u0000"+
		"ed\u0001\u0000\u0000\u0000ef\u0001\u0000\u0000\u0000f\u00b0\u0001\u0000"+
		"\u0000\u0000gi\u0003\u0012\t\u0000hj\u0005:\u0000\u0000ih\u0001\u0000"+
		"\u0000\u0000ij\u0001\u0000\u0000\u0000j\u00b0\u0001\u0000\u0000\u0000"+
		"km\u0003\u0014\n\u0000ln\u0005:\u0000\u0000ml\u0001\u0000\u0000\u0000"+
		"mn\u0001\u0000\u0000\u0000n\u00b0\u0001\u0000\u0000\u0000oq\u0003\u0016"+
		"\u000b\u0000pr\u0005:\u0000\u0000qp\u0001\u0000\u0000\u0000qr\u0001\u0000"+
		"\u0000\u0000r\u00b0\u0001\u0000\u0000\u0000su\u0003(\u0014\u0000tv\u0005"+
		":\u0000\u0000ut\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000v\u00b0"+
		"\u0001\u0000\u0000\u0000wy\u0003,\u0016\u0000xz\u0005:\u0000\u0000yx\u0001"+
		"\u0000\u0000\u0000yz\u0001\u0000\u0000\u0000z\u00b0\u0001\u0000\u0000"+
		"\u0000{}\u00030\u0018\u0000|~\u0005:\u0000\u0000}|\u0001\u0000\u0000\u0000"+
		"}~\u0001\u0000\u0000\u0000~\u00b0\u0001\u0000\u0000\u0000\u007f\u0081"+
		"\u00032\u0019\u0000\u0080\u0082\u0005:\u0000\u0000\u0081\u0080\u0001\u0000"+
		"\u0000\u0000\u0081\u0082\u0001\u0000\u0000\u0000\u0082\u00b0\u0001\u0000"+
		"\u0000\u0000\u0083\u0085\u00034\u001a\u0000\u0084\u0086\u0005:\u0000\u0000"+
		"\u0085\u0084\u0001\u0000\u0000\u0000\u0085\u0086\u0001\u0000\u0000\u0000"+
		"\u0086\u00b0\u0001\u0000\u0000\u0000\u0087\u0089\u0003<\u001e\u0000\u0088"+
		"\u008a\u0005:\u0000\u0000\u0089\u0088\u0001\u0000\u0000\u0000\u0089\u008a"+
		"\u0001\u0000\u0000\u0000\u008a\u00b0\u0001\u0000\u0000\u0000\u008b\u008d"+
		"\u0003>\u001f\u0000\u008c\u008e\u0005:\u0000\u0000\u008d\u008c\u0001\u0000"+
		"\u0000\u0000\u008d\u008e\u0001\u0000\u0000\u0000\u008e\u00b0\u0001\u0000"+
		"\u0000\u0000\u008f\u0091\u0003B!\u0000\u0090\u0092\u0005:\u0000\u0000"+
		"\u0091\u0090\u0001\u0000\u0000\u0000\u0091\u0092\u0001\u0000\u0000\u0000"+
		"\u0092\u00b0\u0001\u0000\u0000\u0000\u0093\u0095\u0003D\"\u0000\u0094"+
		"\u0096\u0005:\u0000\u0000\u0095\u0094\u0001\u0000\u0000\u0000\u0095\u0096"+
		"\u0001\u0000\u0000\u0000\u0096\u00b0\u0001\u0000\u0000\u0000\u0097\u0099"+
		"\u0003F#\u0000\u0098\u009a\u0005:\u0000\u0000\u0099\u0098\u0001\u0000"+
		"\u0000\u0000\u0099\u009a\u0001\u0000\u0000\u0000\u009a\u00b0\u0001\u0000"+
		"\u0000\u0000\u009b\u009d\u0003H$\u0000\u009c\u009e\u0005:\u0000\u0000"+
		"\u009d\u009c\u0001\u0000\u0000\u0000\u009d\u009e\u0001\u0000\u0000\u0000"+
		"\u009e\u00b0\u0001\u0000\u0000\u0000\u009f\u00a1\u0003\u001c\u000e\u0000"+
		"\u00a0\u00a2\u0005:\u0000\u0000\u00a1\u00a0\u0001\u0000\u0000\u0000\u00a1"+
		"\u00a2\u0001\u0000\u0000\u0000\u00a2\u00b0\u0001\u0000\u0000\u0000\u00a3"+
		"\u00a5\u0003 \u0010\u0000\u00a4\u00a6\u0005:\u0000\u0000\u00a5\u00a4\u0001"+
		"\u0000\u0000\u0000\u00a5\u00a6\u0001\u0000\u0000\u0000\u00a6\u00b0\u0001"+
		"\u0000\u0000\u0000\u00a7\u00a9\u0003$\u0012\u0000\u00a8\u00aa\u0005:\u0000"+
		"\u0000\u00a9\u00a8\u0001\u0000\u0000\u0000\u00a9\u00aa\u0001\u0000\u0000"+
		"\u0000\u00aa\u00b0\u0001\u0000\u0000\u0000\u00ab\u00ad\u0003&\u0013\u0000"+
		"\u00ac\u00ae\u0005:\u0000\u0000\u00ad\u00ac\u0001\u0000\u0000\u0000\u00ad"+
		"\u00ae\u0001\u0000\u0000\u0000\u00ae\u00b0\u0001\u0000\u0000\u0000\u00af"+
		"W\u0001\u0000\u0000\u0000\u00af[\u0001\u0000\u0000\u0000\u00af_\u0001"+
		"\u0000\u0000\u0000\u00afc\u0001\u0000\u0000\u0000\u00afg\u0001\u0000\u0000"+
		"\u0000\u00afk\u0001\u0000\u0000\u0000\u00afo\u0001\u0000\u0000\u0000\u00af"+
		"s\u0001\u0000\u0000\u0000\u00afw\u0001\u0000\u0000\u0000\u00af{\u0001"+
		"\u0000\u0000\u0000\u00af\u007f\u0001\u0000\u0000\u0000\u00af\u0083\u0001"+
		"\u0000\u0000\u0000\u00af\u0087\u0001\u0000\u0000\u0000\u00af\u008b\u0001"+
		"\u0000\u0000\u0000\u00af\u008f\u0001\u0000\u0000\u0000\u00af\u0093\u0001"+
		"\u0000\u0000\u0000\u00af\u0097\u0001\u0000\u0000\u0000\u00af\u009b\u0001"+
		"\u0000\u0000\u0000\u00af\u009f\u0001\u0000\u0000\u0000\u00af\u00a3\u0001"+
		"\u0000\u0000\u0000\u00af\u00a7\u0001\u0000\u0000\u0000\u00af\u00ab\u0001"+
		"\u0000\u0000\u0000\u00b0\u0005\u0001\u0000\u0000\u0000\u00b1\u00b2\u0005"+
		"(\u0000\u0000\u00b2\u00b3\u0005D\u0000\u0000\u00b3\u00b4\u0005;\u0000"+
		"\u0000\u00b4\u00b5\u0003L&\u0000\u00b5\u00b6\u0005<\u0000\u0000\u00b6"+
		"\u00b7\u0003J%\u0000\u00b7\u00c3\u0001\u0000\u0000\u0000\u00b8\u00b9\u0005"+
		"(\u0000\u0000\u00b9\u00ba\u0005D\u0000\u0000\u00ba\u00bb\u0005<\u0000"+
		"\u0000\u00bb\u00c3\u0003J%\u0000\u00bc\u00bd\u0005(\u0000\u0000\u00bd"+
		"\u00be\u0005D\u0000\u0000\u00be\u00bf\u0005;\u0000\u0000\u00bf\u00c0\u0003"+
		"L&\u0000\u00c0\u00c1\u0005=\u0000\u0000\u00c1\u00c3\u0001\u0000\u0000"+
		"\u0000\u00c2\u00b1\u0001\u0000\u0000\u0000\u00c2\u00b8\u0001\u0000\u0000"+
		"\u0000\u00c2\u00bc\u0001\u0000\u0000\u0000\u00c3\u0007\u0001\u0000\u0000"+
		"\u0000\u00c4\u00cb\u0005D\u0000\u0000\u00c5\u00c6\u0005D\u0000\u0000\u00c6"+
		"\u00c7\u0005\u0001\u0000\u0000\u00c7\u00c8\u0003J%\u0000\u00c8\u00c9\u0005"+
		"\u0002\u0000\u0000\u00c9\u00cb\u0001\u0000\u0000\u0000\u00ca\u00c4\u0001"+
		"\u0000\u0000\u0000\u00ca\u00c5\u0001\u0000\u0000\u0000\u00cb\u00cc\u0001"+
		"\u0000\u0000\u0000\u00cc\u00cd\u0005<\u0000\u0000\u00cd\u00e3\u0003J%"+
		"\u0000\u00ce\u00d5\u0005D\u0000\u0000\u00cf\u00d0\u0005D\u0000\u0000\u00d0"+
		"\u00d1\u0005\u0001\u0000\u0000\u00d1\u00d2\u0003J%\u0000\u00d2\u00d3\u0005"+
		"\u0002\u0000\u0000\u00d3\u00d5\u0001\u0000\u0000\u0000\u00d4\u00ce\u0001"+
		"\u0000\u0000\u0000\u00d4\u00cf\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d6\u00d7\u0005>\u0000\u0000\u00d7\u00e3\u0003J%"+
		"\u0000\u00d8\u00df\u0005D\u0000\u0000\u00d9\u00da\u0005D\u0000\u0000\u00da"+
		"\u00db\u0005\u0001\u0000\u0000\u00db\u00dc\u0003J%\u0000\u00dc\u00dd\u0005"+
		"\u0002\u0000\u0000\u00dd\u00df\u0001\u0000\u0000\u0000\u00de\u00d8\u0001"+
		"\u0000\u0000\u0000\u00de\u00d9\u0001\u0000\u0000\u0000\u00df\u00e0\u0001"+
		"\u0000\u0000\u0000\u00e0\u00e1\u0005?\u0000\u0000\u00e1\u00e3\u0003J%"+
		"\u0000\u00e2\u00ca\u0001\u0000\u0000\u0000\u00e2\u00d4\u0001\u0000\u0000"+
		"\u0000\u00e2\u00de\u0001\u0000\u0000\u0000\u00e3\t\u0001\u0000\u0000\u0000"+
		"\u00e4\u00e5\u0005)\u0000\u0000\u00e5\u00e6\u0005D\u0000\u0000\u00e6\u00e7"+
		"\u0005;\u0000\u0000\u00e7\u00e8\u0003L&\u0000\u00e8\u00e9\u0005<\u0000"+
		"\u0000\u00e9\u00ea\u0003J%\u0000\u00ea\u00f0\u0001\u0000\u0000\u0000\u00eb"+
		"\u00ec\u0005)\u0000\u0000\u00ec\u00ed\u0005D\u0000\u0000\u00ed\u00ee\u0005"+
		"<\u0000\u0000\u00ee\u00f0\u0003J%\u0000\u00ef\u00e4\u0001\u0000\u0000"+
		"\u0000\u00ef\u00eb\u0001\u0000\u0000\u0000\u00f0\u000b\u0001\u0000\u0000"+
		"\u0000\u00f1\u00f2\u0005(\u0000\u0000\u00f2\u00f3\u0005D\u0000\u0000\u00f3"+
		"\u00f4\u0005;\u0000\u0000\u00f4\u00f5\u0005\u0001\u0000\u0000\u00f5\u00f6"+
		"\u0003L&\u0000\u00f6\u00f7\u0005\u0002\u0000\u0000\u00f7\u00f8\u0005<"+
		"\u0000\u0000\u00f8\u00fa\u0005\u0001\u0000\u0000\u00f9\u00fb\u0003\u000e"+
		"\u0007\u0000\u00fa\u00f9\u0001\u0000\u0000\u0000\u00fa\u00fb\u0001\u0000"+
		"\u0000\u0000\u00fb\u00fc\u0001\u0000\u0000\u0000\u00fc\u00fd\u0005\u0002"+
		"\u0000\u0000\u00fd\u0108\u0001\u0000\u0000\u0000\u00fe\u00ff\u0005(\u0000"+
		"\u0000\u00ff\u0100\u0005D\u0000\u0000\u0100\u0101\u0005;\u0000\u0000\u0101"+
		"\u0102\u0005\u0001\u0000\u0000\u0102\u0103\u0003L&\u0000\u0103\u0104\u0005"+
		"\u0002\u0000\u0000\u0104\u0105\u0005<\u0000\u0000\u0105\u0106\u0005D\u0000"+
		"\u0000\u0106\u0108\u0001\u0000\u0000\u0000\u0107\u00f1\u0001\u0000\u0000"+
		"\u0000\u0107\u00fe\u0001\u0000\u0000\u0000\u0108\r\u0001\u0000\u0000\u0000"+
		"\u0109\u010e\u0003J%\u0000\u010a\u010b\u0005\u0003\u0000\u0000\u010b\u010d"+
		"\u0003J%\u0000\u010c\u010a\u0001\u0000\u0000\u0000\u010d\u0110\u0001\u0000"+
		"\u0000\u0000\u010e\u010c\u0001\u0000\u0000\u0000\u010e\u010f\u0001\u0000"+
		"\u0000\u0000\u010f\u000f\u0001\u0000\u0000\u0000\u0110\u010e\u0001\u0000"+
		"\u0000\u0000\u0111\u0112\u0005\u0001\u0000\u0000\u0112\u0113\u0003L&\u0000"+
		"\u0113\u0114\u0005\u0002\u0000\u0000\u0114\u0011\u0001\u0000\u0000\u0000"+
		"\u0115\u0116\u0005D\u0000\u0000\u0116\u0117\u0005\u0004\u0000\u0000\u0117"+
		"\u0118\u0005\u0005\u0000\u0000\u0118\u0119\u0003J%\u0000\u0119\u011a\u0005"+
		"\u0006\u0000\u0000\u011a\u0013\u0001\u0000\u0000\u0000\u011b\u011c\u0005"+
		"D\u0000\u0000\u011c\u011d\u0005\u0007\u0000\u0000\u011d\u011e\u0005\u0005"+
		"\u0000\u0000\u011e\u011f\u0005\u0006\u0000\u0000\u011f\u0015\u0001\u0000"+
		"\u0000\u0000\u0120\u0121\u0005D\u0000\u0000\u0121\u0122\u0005\b\u0000"+
		"\u0000\u0122\u0123\u0005\u0005\u0000\u0000\u0123\u0124\u0005\t\u0000\u0000"+
		"\u0124\u0125\u0005;\u0000\u0000\u0125\u0126\u0003J%\u0000\u0126\u0127"+
		"\u0005\u0006\u0000\u0000\u0127\u0017\u0001\u0000\u0000\u0000\u0128\u0129"+
		"\u0005(\u0000\u0000\u0129\u012a\u0005D\u0000\u0000\u012a\u012b\u0005;"+
		"\u0000\u0000\u012b\u012c\u0005\u0001\u0000\u0000\u012c\u012d\u0005\u0001"+
		"\u0000\u0000\u012d\u012e\u0003L&\u0000\u012e\u012f\u0005\u0002\u0000\u0000"+
		"\u012f\u0130\u0005\u0002\u0000\u0000\u0130\u0131\u0005<\u0000\u0000\u0131"+
		"\u0133\u0005\u0001\u0000\u0000\u0132\u0134\u0003\u001a\r\u0000\u0133\u0132"+
		"\u0001\u0000\u0000\u0000\u0133\u0134\u0001\u0000\u0000\u0000\u0134\u0135"+
		"\u0001\u0000\u0000\u0000\u0135\u0136\u0005\u0002\u0000\u0000\u0136\u0143"+
		"\u0001\u0000\u0000\u0000\u0137\u0138\u0005(\u0000\u0000\u0138\u0139\u0005"+
		"D\u0000\u0000\u0139\u013a\u0005;\u0000\u0000\u013a\u013b\u0005\u0001\u0000"+
		"\u0000\u013b\u013c\u0005\u0001\u0000\u0000\u013c\u013d\u0003L&\u0000\u013d"+
		"\u013e\u0005\u0002\u0000\u0000\u013e\u013f\u0005\u0002\u0000\u0000\u013f"+
		"\u0140\u0005<\u0000\u0000\u0140\u0141\u0005D\u0000\u0000\u0141\u0143\u0001"+
		"\u0000\u0000\u0000\u0142\u0128\u0001\u0000\u0000\u0000\u0142\u0137\u0001"+
		"\u0000\u0000\u0000\u0143\u0019\u0001\u0000\u0000\u0000\u0144\u0146\u0005"+
		"\u0001\u0000\u0000\u0145\u0147\u0003\u000e\u0007\u0000\u0146\u0145\u0001"+
		"\u0000\u0000\u0000\u0146\u0147\u0001\u0000\u0000\u0000\u0147\u0148\u0001"+
		"\u0000\u0000\u0000\u0148\u0151\u0005\u0002\u0000\u0000\u0149\u014a\u0005"+
		"\u0003\u0000\u0000\u014a\u014c\u0005\u0001\u0000\u0000\u014b\u014d\u0003"+
		"\u000e\u0007\u0000\u014c\u014b\u0001\u0000\u0000\u0000\u014c\u014d\u0001"+
		"\u0000\u0000\u0000\u014d\u014e\u0001\u0000\u0000\u0000\u014e\u0150\u0005"+
		"\u0002\u0000\u0000\u014f\u0149\u0001\u0000\u0000\u0000\u0150\u0153\u0001"+
		"\u0000\u0000\u0000\u0151\u014f\u0001\u0000\u0000\u0000\u0151\u0152\u0001"+
		"\u0000\u0000\u0000\u0152\u001b\u0001\u0000\u0000\u0000\u0153\u0151\u0001"+
		"\u0000\u0000\u0000\u0154\u0155\u0005\n\u0000\u0000\u0155\u0156\u0005D"+
		"\u0000\u0000\u0156\u015a\u0005\u000b\u0000\u0000\u0157\u0159\u0003\u001e"+
		"\u000f\u0000\u0158\u0157\u0001\u0000\u0000\u0000\u0159\u015c\u0001\u0000"+
		"\u0000\u0000\u015a\u0158\u0001\u0000\u0000\u0000\u015a\u015b\u0001\u0000"+
		"\u0000\u0000\u015b\u015d\u0001\u0000\u0000\u0000\u015c\u015a\u0001\u0000"+
		"\u0000\u0000\u015d\u015e\u0005\f\u0000\u0000\u015e\u001d\u0001\u0000\u0000"+
		"\u0000\u015f\u0160\u0007\u0000\u0000\u0000\u0160\u0161\u0005D\u0000\u0000"+
		"\u0161\u0162\u0005;\u0000\u0000\u0162\u0163\u0003L&\u0000\u0163\u001f"+
		"\u0001\u0000\u0000\u0000\u0164\u0165\u0005D\u0000\u0000\u0165\u0167\u0005"+
		"\u0005\u0000\u0000\u0166\u0168\u0003\"\u0011\u0000\u0167\u0166\u0001\u0000"+
		"\u0000\u0000\u0167\u0168\u0001\u0000\u0000\u0000\u0168\u0169\u0001\u0000"+
		"\u0000\u0000\u0169\u016a\u0005\u0006\u0000\u0000\u016a!\u0001\u0000\u0000"+
		"\u0000\u016b\u016c\u0005D\u0000\u0000\u016c\u016d\u0005;\u0000\u0000\u016d"+
		"\u0170\u0003J%\u0000\u016e\u016f\u0005\u0003\u0000\u0000\u016f\u0171\u0003"+
		"\"\u0011\u0000\u0170\u016e\u0001\u0000\u0000\u0000\u0170\u0171\u0001\u0000"+
		"\u0000\u0000\u0171#\u0001\u0000\u0000\u0000\u0172\u0173\u0005D\u0000\u0000"+
		"\u0173\u0174\u0005\r\u0000\u0000\u0174\u0175\u0005D\u0000\u0000\u0175"+
		"%\u0001\u0000\u0000\u0000\u0176\u0177\u0005D\u0000\u0000\u0177\u0178\u0005"+
		"\r\u0000\u0000\u0178\u0179\u0005D\u0000\u0000\u0179\u017a\u0005<\u0000"+
		"\u0000\u017a\u017b\u0003J%\u0000\u017b\'\u0001\u0000\u0000\u0000\u017c"+
		"\u017d\u0005*\u0000\u0000\u017d\u017e\u0005D\u0000\u0000\u017e\u0180\u0005"+
		"\u0005\u0000\u0000\u017f\u0181\u0003*\u0015\u0000\u0180\u017f\u0001\u0000"+
		"\u0000\u0000\u0180\u0181\u0001\u0000\u0000\u0000\u0181\u0182\u0001\u0000"+
		"\u0000\u0000\u0182\u0185\u0005\u0006\u0000\u0000\u0183\u0184\u0005\u000e"+
		"\u0000\u0000\u0184\u0186\u0003L&\u0000\u0185\u0183\u0001\u0000\u0000\u0000"+
		"\u0185\u0186\u0001\u0000\u0000\u0000\u0186\u0187\u0001\u0000\u0000\u0000"+
		"\u0187\u0188\u0005\u000b\u0000\u0000\u0188\u0189\u0003\u0002\u0001\u0000"+
		"\u0189\u018a\u0005\f\u0000\u0000\u018a)\u0001\u0000\u0000\u0000\u018b"+
		"\u018d\u0007\u0001\u0000\u0000\u018c\u018b\u0001\u0000\u0000\u0000\u018c"+
		"\u018d\u0001\u0000\u0000\u0000\u018d\u018e\u0001\u0000\u0000\u0000\u018e"+
		"\u018f\u0005D\u0000\u0000\u018f\u0191\u0005;\u0000\u0000\u0190\u0192\u0005"+
		"+\u0000\u0000\u0191\u0190\u0001\u0000\u0000\u0000\u0191\u0192\u0001\u0000"+
		"\u0000\u0000\u0192\u0195\u0001\u0000\u0000\u0000\u0193\u0196\u0003L&\u0000"+
		"\u0194\u0196\u0003\u0010\b\u0000\u0195\u0193\u0001\u0000\u0000\u0000\u0195"+
		"\u0194\u0001\u0000\u0000\u0000\u0196\u0199\u0001\u0000\u0000\u0000\u0197"+
		"\u0198\u0005\u0003\u0000\u0000\u0198\u019a\u0003*\u0015\u0000\u0199\u0197"+
		"\u0001\u0000\u0000\u0000\u0199\u019a\u0001\u0000\u0000\u0000\u019a+\u0001"+
		"\u0000\u0000\u0000\u019b\u019c\u0005D\u0000\u0000\u019c\u019e\u0005\u0005"+
		"\u0000\u0000\u019d\u019f\u0003.\u0017\u0000\u019e\u019d\u0001\u0000\u0000"+
		"\u0000\u019e\u019f\u0001\u0000\u0000\u0000\u019f\u01a0\u0001\u0000\u0000"+
		"\u0000\u01a0\u01a1\u0005\u0006\u0000\u0000\u01a1-\u0001\u0000\u0000\u0000"+
		"\u01a2\u01a3\u0005D\u0000\u0000\u01a3\u01a5\u0005;\u0000\u0000\u01a4\u01a2"+
		"\u0001\u0000\u0000\u0000\u01a4\u01a5\u0001\u0000\u0000\u0000\u01a5\u01a7"+
		"\u0001\u0000\u0000\u0000\u01a6\u01a8\u0005,\u0000\u0000\u01a7\u01a6\u0001"+
		"\u0000\u0000\u0000\u01a7\u01a8\u0001\u0000\u0000\u0000\u01a8\u01a9\u0001"+
		"\u0000\u0000\u0000\u01a9\u01ac\u0003J%\u0000\u01aa\u01ab\u0005\u0003\u0000"+
		"\u0000\u01ab\u01ad\u0003.\u0017\u0000\u01ac\u01aa\u0001\u0000\u0000\u0000"+
		"\u01ac\u01ad\u0001\u0000\u0000\u0000\u01ad/\u0001\u0000\u0000\u0000\u01ae"+
		"\u01af\u0005-\u0000\u0000\u01af\u01b0\u0005\u0005\u0000\u0000\u01b0\u01b1"+
		"\u0003J%\u0000\u01b1\u01b2\u0005\u0006\u0000\u0000\u01b2\u01c0\u0001\u0000"+
		"\u0000\u0000\u01b3\u01b4\u0005-\u0000\u0000\u01b4\u01b5\u0005\u0005\u0000"+
		"\u0000\u01b5\u01ba\u0003J%\u0000\u01b6\u01b7\u0005\u0003\u0000\u0000\u01b7"+
		"\u01b9\u0003J%\u0000\u01b8\u01b6\u0001\u0000\u0000\u0000\u01b9\u01bc\u0001"+
		"\u0000\u0000\u0000\u01ba\u01b8\u0001\u0000\u0000\u0000\u01ba\u01bb\u0001"+
		"\u0000\u0000\u0000\u01bb\u01bd\u0001\u0000\u0000\u0000\u01bc\u01ba\u0001"+
		"\u0000\u0000\u0000\u01bd\u01be\u0005\u0006\u0000\u0000\u01be\u01c0\u0001"+
		"\u0000\u0000\u0000\u01bf\u01ae\u0001\u0000\u0000\u0000\u01bf\u01b3\u0001"+
		"\u0000\u0000\u0000\u01c01\u0001\u0000\u0000\u0000\u01c1\u01c2\u0005.\u0000"+
		"\u0000\u01c2\u01c3\u0003J%\u0000\u01c3\u01c4\u0005\u000b\u0000\u0000\u01c4"+
		"\u01c5\u0003\u0002\u0001\u0000\u01c5\u01c6\u0005\f\u0000\u0000\u01c6\u01da"+
		"\u0001\u0000\u0000\u0000\u01c7\u01c8\u0005.\u0000\u0000\u01c8\u01c9\u0003"+
		"J%\u0000\u01c9\u01ca\u0005\u000b\u0000\u0000\u01ca\u01cb\u0003\u0002\u0001"+
		"\u0000\u01cb\u01cc\u0005\f\u0000\u0000\u01cc\u01cd\u0005/\u0000\u0000"+
		"\u01cd\u01ce\u0005\u000b\u0000\u0000\u01ce\u01cf\u0003\u0002\u0001\u0000"+
		"\u01cf\u01d0\u0005\f\u0000\u0000\u01d0\u01da\u0001\u0000\u0000\u0000\u01d1"+
		"\u01d2\u0005.\u0000\u0000\u01d2\u01d3\u0003J%\u0000\u01d3\u01d4\u0005"+
		"\u000b\u0000\u0000\u01d4\u01d5\u0003\u0002\u0001\u0000\u01d5\u01d6\u0005"+
		"\f\u0000\u0000\u01d6\u01d7\u0005/\u0000\u0000\u01d7\u01d8\u00032\u0019"+
		"\u0000\u01d8\u01da\u0001\u0000\u0000\u0000\u01d9\u01c1\u0001\u0000\u0000"+
		"\u0000\u01d9\u01c7\u0001\u0000\u0000\u0000\u01d9\u01d1\u0001\u0000\u0000"+
		"\u0000\u01da3\u0001\u0000\u0000\u0000\u01db\u01dc\u00050\u0000\u0000\u01dc"+
		"\u01dd\u0003J%\u0000\u01dd\u01de\u0005\u000b\u0000\u0000\u01de\u01df\u0003"+
		"6\u001b\u0000\u01df\u01e0\u0005\f\u0000\u0000\u01e05\u0001\u0000\u0000"+
		"\u0000\u01e1\u01e4\u00038\u001c\u0000\u01e2\u01e4\u0003:\u001d\u0000\u01e3"+
		"\u01e1\u0001\u0000\u0000\u0000\u01e3\u01e2\u0001\u0000\u0000\u0000\u01e4"+
		"\u01e7\u0001\u0000\u0000\u0000\u01e5\u01e3\u0001\u0000\u0000\u0000\u01e5"+
		"\u01e6\u0001\u0000\u0000\u0000\u01e67\u0001\u0000\u0000\u0000\u01e7\u01e5"+
		"\u0001\u0000\u0000\u0000\u01e8\u01e9\u00051\u0000\u0000\u01e9\u01ea\u0003"+
		"J%\u0000\u01ea\u01eb\u0005;\u0000\u0000\u01eb\u01ed\u0003\u0002\u0001"+
		"\u0000\u01ec\u01ee\u00058\u0000\u0000\u01ed\u01ec\u0001\u0000\u0000\u0000"+
		"\u01ed\u01ee\u0001\u0000\u0000\u0000\u01ee9\u0001\u0000\u0000\u0000\u01ef"+
		"\u01f0\u00052\u0000\u0000\u01f0\u01f1\u0005;\u0000\u0000\u01f1\u01f3\u0003"+
		"\u0002\u0001\u0000\u01f2\u01f4\u00058\u0000\u0000\u01f3\u01f2\u0001\u0000"+
		"\u0000\u0000\u01f3\u01f4\u0001\u0000\u0000\u0000\u01f4;\u0001\u0000\u0000"+
		"\u0000\u01f5\u01f6\u00053\u0000\u0000\u01f6\u01f7\u0003J%\u0000\u01f7"+
		"\u01f8\u0005\u000b\u0000\u0000\u01f8\u01f9\u0003\u0002\u0001\u0000\u01f9"+
		"\u01fa\u0005\f\u0000\u0000\u01fa=\u0001\u0000\u0000\u0000\u01fb\u01fc"+
		"\u00054\u0000\u0000\u01fc\u01fd\u0007\u0001\u0000\u0000\u01fd\u0200\u0005"+
		"5\u0000\u0000\u01fe\u0201\u0003J%\u0000\u01ff\u0201\u0003@ \u0000\u0200"+
		"\u01fe\u0001\u0000\u0000\u0000\u0200\u01ff\u0001\u0000\u0000\u0000\u0201"+
		"\u0202\u0001\u0000\u0000\u0000\u0202\u0203\u0005\u000b\u0000\u0000\u0203"+
		"\u0204\u0003\u0002\u0001\u0000\u0204\u0205\u0005\f\u0000\u0000\u0205?"+
		"\u0001\u0000\u0000\u0000\u0206\u0207\u0003J%\u0000\u0207\u0208\u0005\u000f"+
		"\u0000\u0000\u0208\u0209\u0003J%\u0000\u0209A\u0001\u0000\u0000\u0000"+
		"\u020a\u020b\u00056\u0000\u0000\u020b\u020c\u0003J%\u0000\u020c\u020d"+
		"\u0005/\u0000\u0000\u020d\u020e\u0005\u000b\u0000\u0000\u020e\u020f\u0003"+
		"\u0002\u0001\u0000\u020f\u0210\u0005\f\u0000\u0000\u0210C\u0001\u0000"+
		"\u0000\u0000\u0211\u0212\u00058\u0000\u0000\u0212E\u0001\u0000\u0000\u0000"+
		"\u0213\u0214\u00057\u0000\u0000\u0214G\u0001\u0000\u0000\u0000\u0215\u0219"+
		"\u00059\u0000\u0000\u0216\u0217\u00059\u0000\u0000\u0217\u0219\u0003J"+
		"%\u0000\u0218\u0215\u0001\u0000\u0000\u0000\u0218\u0216\u0001\u0000\u0000"+
		"\u0000\u0219I\u0001\u0000\u0000\u0000\u021a\u021b\u0006%\uffff\uffff\u0000"+
		"\u021b\u021c\u0005\u0005\u0000\u0000\u021c\u021d\u0003J%\u0000\u021d\u021e"+
		"\u0005\u0006\u0000\u0000\u021e\u0242\u0001\u0000\u0000\u0000\u021f\u0220"+
		"\u0007\u0002\u0000\u0000\u0220\u0242\u0003J%\u0015\u0221\u0242\u0007\u0003"+
		"\u0000\u0000\u0222\u0242\u0005\'\u0000\u0000\u0223\u0242\u0005A\u0000"+
		"\u0000\u0224\u0242\u0005B\u0000\u0000\u0225\u0242\u0005C\u0000\u0000\u0226"+
		"\u0227\u0005D\u0000\u0000\u0227\u0242\u0005 \u0000\u0000\u0228\u0229\u0005"+
		"D\u0000\u0000\u0229\u0242\u0005!\u0000\u0000\u022a\u022b\u0005D\u0000"+
		"\u0000\u022b\u022c\u0005\u0001\u0000\u0000\u022c\u022d\u0003J%\u0000\u022d"+
		"\u022e\u0005\u0002\u0000\u0000\u022e\u0242\u0001\u0000\u0000\u0000\u022f"+
		"\u0242\u0003 \u0010\u0000\u0230\u0242\u0003$\u0012\u0000\u0231\u0242\u0005"+
		"D\u0000\u0000\u0232\u0233\u0005\"\u0000\u0000\u0233\u0234\u0005\u0005"+
		"\u0000\u0000\u0234\u0235\u0003J%\u0000\u0235\u0236\u0005\u0006\u0000\u0000"+
		"\u0236\u0242\u0001\u0000\u0000\u0000\u0237\u0238\u0005#\u0000\u0000\u0238"+
		"\u0239\u0005\u0005\u0000\u0000\u0239\u023a\u0003J%\u0000\u023a\u023b\u0005"+
		"\u0006\u0000\u0000\u023b\u0242\u0001\u0000\u0000\u0000\u023c\u023d\u0005"+
		"$\u0000\u0000\u023d\u023e\u0005\u0005\u0000\u0000\u023e\u023f\u0003J%"+
		"\u0000\u023f\u0240\u0005\u0006\u0000\u0000\u0240\u0242\u0001\u0000\u0000"+
		"\u0000\u0241\u021a\u0001\u0000\u0000\u0000\u0241\u021f\u0001\u0000\u0000"+
		"\u0000\u0241\u0221\u0001\u0000\u0000\u0000\u0241\u0222\u0001\u0000\u0000"+
		"\u0000\u0241\u0223\u0001\u0000\u0000\u0000\u0241\u0224\u0001\u0000\u0000"+
		"\u0000\u0241\u0225\u0001\u0000\u0000\u0000\u0241\u0226\u0001\u0000\u0000"+
		"\u0000\u0241\u0228\u0001\u0000\u0000\u0000\u0241\u022a\u0001\u0000\u0000"+
		"\u0000\u0241\u022f\u0001\u0000\u0000\u0000\u0241\u0230\u0001\u0000\u0000"+
		"\u0000\u0241\u0231\u0001\u0000\u0000\u0000\u0241\u0232\u0001\u0000\u0000"+
		"\u0000\u0241\u0237\u0001\u0000\u0000\u0000\u0241\u023c\u0001\u0000\u0000"+
		"\u0000\u0242\u0257\u0001\u0000\u0000\u0000\u0243\u0244\n\u0014\u0000\u0000"+
		"\u0244\u0245\u0007\u0004\u0000\u0000\u0245\u0256\u0003J%\u0015\u0246\u0247"+
		"\n\u0013\u0000\u0000\u0247\u0248\u0007\u0005\u0000\u0000\u0248\u0256\u0003"+
		"J%\u0014\u0249\u024a\n\u0012\u0000\u0000\u024a\u024b\u0007\u0006\u0000"+
		"\u0000\u024b\u0256\u0003J%\u0013\u024c\u024d\n\u0011\u0000\u0000\u024d"+
		"\u024e\u0007\u0007\u0000\u0000\u024e\u0256\u0003J%\u0012\u024f\u0250\n"+
		"\u0010\u0000\u0000\u0250\u0251\u0007\b\u0000\u0000\u0251\u0256\u0003J"+
		"%\u0011\u0252\u0253\n\u000f\u0000\u0000\u0253\u0254\u0007\t\u0000\u0000"+
		"\u0254\u0256\u0003J%\u0010\u0255\u0243\u0001\u0000\u0000\u0000\u0255\u0246"+
		"\u0001\u0000\u0000\u0000\u0255\u0249\u0001\u0000\u0000\u0000\u0255\u024c"+
		"\u0001\u0000\u0000\u0000\u0255\u024f\u0001\u0000\u0000\u0000\u0255\u0252"+
		"\u0001\u0000\u0000\u0000\u0256\u0259\u0001\u0000\u0000\u0000\u0257\u0255"+
		"\u0001\u0000\u0000\u0000\u0257\u0258\u0001\u0000\u0000\u0000\u0258K\u0001"+
		"\u0000\u0000\u0000\u0259\u0257\u0001\u0000\u0000\u0000\u025a\u025b\u0007"+
		"\n\u0000\u0000\u025bM\u0001\u0000\u0000\u0000?TY]aeimquy}\u0081\u0085"+
		"\u0089\u008d\u0091\u0095\u0099\u009d\u00a1\u00a5\u00a9\u00ad\u00af\u00c2"+
		"\u00ca\u00d4\u00de\u00e2\u00ef\u00fa\u0107\u010e\u0133\u0142\u0146\u014c"+
		"\u0151\u015a\u0167\u0170\u0180\u0185\u018c\u0191\u0195\u0199\u019e\u01a4"+
		"\u01a7\u01ac\u01ba\u01bf\u01d9\u01e3\u01e5\u01ed\u01f3\u0200\u0218\u0241"+
		"\u0255\u0257";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}