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
		T__31=32, T__32=33, T__33=34, INT=35, FLOAT=36, STRING=37, BOOL=38, CHARACTER=39, 
		NIL=40, VAR=41, LET=42, FUNC=43, INOUT=44, REFERENCE=45, PRINT=46, IF=47, 
		ELSE=48, SWITCH=49, CASE=50, DEFAULT=51, WHILE=52, FOR=53, IN=54, GUARD=55, 
		CONTINUE=56, BREAK=57, RETURN=58, SEMICOLON=59, COLON=60, EQUAL=61, QUESTION_MARK=62, 
		INCREMENT=63, DECREMENT=64, UNDERSCORE=65, DIGIT=66, STR=67, CHAR=68, 
		ID=69, WHITESPACE=70, MULTI_COMMENT=71, LINE_COMMENT=72;
	public static final int
		RULE_start = 0, RULE_block = 1, RULE_stmts = 2, RULE_varDeclaration = 3, 
		RULE_varAssignment = 4, RULE_letDeclaration = 5, RULE_vectorDeclaration = 6, 
		RULE_valuesVectorDeclaration = 7, RULE_vectorType = 8, RULE_vectorAppend = 9, 
		RULE_vectorRemoveLast = 10, RULE_vectorRemoveAt = 11, RULE_matrixDeclaration = 12, 
		RULE_valuesMatrixDeclaration = 13, RULE_structDeclaration = 14, RULE_structAttribute = 15, 
		RULE_structFunction = 16, RULE_structAccess = 17, RULE_structAssignment = 18, 
		RULE_funcDeclaration = 19, RULE_paramsFuncDeclaration = 20, RULE_funcCall = 21, 
		RULE_paramsFuncCall = 22, RULE_printStmt = 23, RULE_ifStmt = 24, RULE_switchStmt = 25, 
		RULE_cases = 26, RULE_caseStmt = 27, RULE_defaultStmt = 28, RULE_whileStmt = 29, 
		RULE_forStmt = 30, RULE_range = 31, RULE_guardStmt = 32, RULE_breakStmt = 33, 
		RULE_continueStmt = 34, RULE_returnStmt = 35, RULE_expr = 36, RULE_type = 37;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "block", "stmts", "varDeclaration", "varAssignment", "letDeclaration", 
			"vectorDeclaration", "valuesVectorDeclaration", "vectorType", "vectorAppend", 
			"vectorRemoveLast", "vectorRemoveAt", "matrixDeclaration", "valuesMatrixDeclaration", 
			"structDeclaration", "structAttribute", "structFunction", "structAccess", 
			"structAssignment", "funcDeclaration", "paramsFuncDeclaration", "funcCall", 
			"paramsFuncCall", "printStmt", "ifStmt", "switchStmt", "cases", "caseStmt", 
			"defaultStmt", "whileStmt", "forStmt", "range", "guardStmt", "breakStmt", 
			"continueStmt", "returnStmt", "expr", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'['", "']'", "','", "'.append'", "'('", "')'", "'.removeLast'", 
			"'.remove'", "'at'", "'struct'", "'{'", "'}'", "'mutating'", "'.'", "'->'", 
			"'...'", "'!'", "'-'", "'*'", "'/'", "'%'", "'+'", "'>='", "'>'", "'<='", 
			"'<'", "'=='", "'!='", "'&&'", "'||'", "'true'", "'false'", "'.count'", 
			"'.isEmpty'", "'Int'", "'Float'", "'String'", "'Bool'", "'Character'", 
			"'nil'", "'var'", "'let'", "'func'", "'inout'", "'&'", "'print'", "'if'", 
			"'else'", "'switch'", "'case'", "'default'", "'while'", "'for'", "'in'", 
			"'guard'", "'continue'", "'break'", "'return'", "';'", "':'", "'='", 
			"'?'", "'+='", "'-='", "'_'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, "INT", 
			"FLOAT", "STRING", "BOOL", "CHARACTER", "NIL", "VAR", "LET", "FUNC", 
			"INOUT", "REFERENCE", "PRINT", "IF", "ELSE", "SWITCH", "CASE", "DEFAULT", 
			"WHILE", "FOR", "IN", "GUARD", "CONTINUE", "BREAK", "RETURN", "SEMICOLON", 
			"COLON", "EQUAL", "QUESTION_MARK", "INCREMENT", "DECREMENT", "UNDERSCORE", 
			"DIGIT", "STR", "CHAR", "ID", "WHITESPACE", "MULTI_COMMENT", "LINE_COMMENT"
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
			setState(76);
			block();
			setState(77);
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
			setState(82);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(79);
					stmts();
					}
					} 
				}
				setState(84);
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
			setState(157);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(85);
				varDeclaration();
				setState(87);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(86);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(89);
				varAssignment();
				setState(91);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(90);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(93);
				letDeclaration();
				setState(95);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(94);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(97);
				vectorDeclaration();
				setState(99);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(98);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(101);
				vectorAppend();
				setState(103);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(102);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(105);
				vectorRemoveLast();
				setState(107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(106);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(109);
				vectorRemoveAt();
				setState(111);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(110);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(113);
				funcDeclaration();
				setState(115);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(114);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(117);
				funcCall();
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(118);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(121);
				printStmt();
				setState(123);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(122);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(125);
				ifStmt();
				setState(127);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(126);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(129);
				switchStmt();
				setState(131);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(130);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(133);
				whileStmt();
				setState(135);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(134);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(137);
				forStmt();
				setState(139);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(138);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(141);
				guardStmt();
				setState(143);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(142);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(145);
				breakStmt();
				setState(147);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(146);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(149);
				continueStmt();
				setState(151);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(150);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(153);
				returnStmt();
				setState(155);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(154);
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
			setState(176);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				_localctx = new TypeValueVarDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(159);
				match(VAR);
				setState(160);
				match(ID);
				setState(161);
				match(COLON);
				setState(162);
				type();
				setState(163);
				match(EQUAL);
				setState(164);
				expr(0);
				}
				break;
			case 2:
				_localctx = new ValueVarDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(166);
				match(VAR);
				setState(167);
				match(ID);
				setState(168);
				match(EQUAL);
				setState(169);
				expr(0);
				}
				break;
			case 3:
				_localctx = new TypeVarDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(170);
				match(VAR);
				setState(171);
				match(ID);
				setState(172);
				match(COLON);
				setState(173);
				type();
				setState(174);
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
			setState(208);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,24,_ctx) ) {
			case 1:
				_localctx = new ValueVarAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(184);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
				case 1:
					{
					setState(178);
					match(ID);
					}
					break;
				case 2:
					{
					setState(179);
					match(ID);
					setState(180);
					match(T__0);
					setState(181);
					expr(0);
					setState(182);
					match(T__1);
					}
					break;
				}
				setState(186);
				match(EQUAL);
				setState(187);
				expr(0);
				}
				break;
			case 2:
				_localctx = new IncrementVarAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(194);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
				case 1:
					{
					setState(188);
					match(ID);
					}
					break;
				case 2:
					{
					setState(189);
					match(ID);
					setState(190);
					match(T__0);
					setState(191);
					expr(0);
					setState(192);
					match(T__1);
					}
					break;
				}
				setState(196);
				match(INCREMENT);
				setState(197);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DecrementVarAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(204);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,23,_ctx) ) {
				case 1:
					{
					setState(198);
					match(ID);
					}
					break;
				case 2:
					{
					setState(199);
					match(ID);
					setState(200);
					match(T__0);
					setState(201);
					expr(0);
					setState(202);
					match(T__1);
					}
					break;
				}
				setState(206);
				match(DECREMENT);
				setState(207);
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
			setState(221);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				_localctx = new TypeValueLetDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(210);
				match(LET);
				setState(211);
				match(ID);
				setState(212);
				match(COLON);
				setState(213);
				type();
				setState(214);
				match(EQUAL);
				setState(215);
				expr(0);
				}
				break;
			case 2:
				_localctx = new ValueLetDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(217);
				match(LET);
				setState(218);
				match(ID);
				setState(219);
				match(EQUAL);
				setState(220);
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
			setState(245);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,27,_ctx) ) {
			case 1:
				_localctx = new VectorValuesDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(223);
				match(VAR);
				setState(224);
				match(ID);
				setState(225);
				match(COLON);
				setState(226);
				match(T__0);
				setState(227);
				type();
				setState(228);
				match(T__1);
				setState(229);
				match(EQUAL);
				setState(230);
				match(T__0);
				setState(232);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1346472640544L) != 0) || ((((_la - 66)) & ~0x3f) == 0 && ((1L << (_la - 66)) & 15L) != 0)) {
					{
					setState(231);
					valuesVectorDeclaration();
					}
				}

				setState(234);
				match(T__1);
				}
				break;
			case 2:
				_localctx = new VectorIdDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(236);
				match(VAR);
				setState(237);
				match(ID);
				setState(238);
				match(COLON);
				setState(239);
				match(T__0);
				setState(240);
				type();
				setState(241);
				match(T__1);
				setState(242);
				match(EQUAL);
				setState(243);
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
			setState(247);
			expr(0);
			setState(252);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(248);
				match(T__2);
				setState(249);
				expr(0);
				}
				}
				setState(254);
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
			setState(255);
			match(T__0);
			setState(256);
			type();
			setState(257);
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
			setState(259);
			match(ID);
			setState(260);
			match(T__3);
			setState(261);
			match(T__4);
			setState(262);
			expr(0);
			setState(263);
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
			setState(265);
			match(ID);
			setState(266);
			match(T__6);
			setState(267);
			match(T__4);
			setState(268);
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
			setState(270);
			match(ID);
			setState(271);
			match(T__7);
			setState(272);
			match(T__4);
			setState(273);
			match(T__8);
			setState(274);
			match(COLON);
			setState(275);
			expr(0);
			setState(276);
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
			setState(304);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,30,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(278);
				match(VAR);
				setState(279);
				match(ID);
				setState(280);
				match(COLON);
				setState(281);
				match(T__0);
				setState(282);
				match(T__0);
				setState(283);
				type();
				setState(284);
				match(T__1);
				setState(285);
				match(T__1);
				setState(286);
				match(EQUAL);
				setState(287);
				match(T__0);
				setState(289);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__0) {
					{
					setState(288);
					valuesMatrixDeclaration();
					}
				}

				setState(291);
				match(T__1);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(293);
				match(VAR);
				setState(294);
				match(ID);
				setState(295);
				match(COLON);
				setState(296);
				match(T__0);
				setState(297);
				match(T__0);
				setState(298);
				type();
				setState(299);
				match(T__1);
				setState(300);
				match(T__1);
				setState(301);
				match(EQUAL);
				setState(302);
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
			setState(306);
			match(T__0);
			setState(308);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1346472640544L) != 0) || ((((_la - 66)) & ~0x3f) == 0 && ((1L << (_la - 66)) & 15L) != 0)) {
				{
				setState(307);
				valuesVectorDeclaration();
				}
			}

			setState(310);
			match(T__1);
			setState(319);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__2) {
				{
				{
				setState(311);
				match(T__2);
				setState(312);
				match(T__0);
				setState(314);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 1346472640544L) != 0) || ((((_la - 66)) & ~0x3f) == 0 && ((1L << (_la - 66)) & 15L) != 0)) {
					{
					setState(313);
					valuesVectorDeclaration();
					}
				}

				setState(316);
				match(T__1);
				}
				}
				setState(321);
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
		public List<StructAttributeContext> structAttribute() {
			return getRuleContexts(StructAttributeContext.class);
		}
		public StructAttributeContext structAttribute(int i) {
			return getRuleContext(StructAttributeContext.class,i);
		}
		public List<StructFunctionContext> structFunction() {
			return getRuleContexts(StructFunctionContext.class);
		}
		public StructFunctionContext structFunction(int i) {
			return getRuleContext(StructFunctionContext.class,i);
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
			setState(322);
			match(T__9);
			setState(323);
			match(ID);
			setState(324);
			match(T__10);
			setState(329);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 15393162797056L) != 0)) {
				{
				setState(327);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case VAR:
				case LET:
					{
					setState(325);
					structAttribute();
					}
					break;
				case T__12:
				case FUNC:
					{
					setState(326);
					structFunction();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(331);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(332);
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
	public static class StructAttributeContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(GrammarParser.ID, 0); }
		public TerminalNode VAR() { return getToken(GrammarParser.VAR, 0); }
		public TerminalNode LET() { return getToken(GrammarParser.LET, 0); }
		public TerminalNode COLON() { return getToken(GrammarParser.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUAL() { return getToken(GrammarParser.EQUAL, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(GrammarParser.SEMICOLON, 0); }
		public StructAttributeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structAttribute; }
	}

	public final StructAttributeContext structAttribute() throws RecognitionException {
		StructAttributeContext _localctx = new StructAttributeContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_structAttribute);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(334);
			_la = _input.LA(1);
			if ( !(_la==VAR || _la==LET) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(335);
			match(ID);
			setState(338);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COLON) {
				{
				setState(336);
				match(COLON);
				setState(337);
				type();
				}
			}

			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==EQUAL) {
				{
				setState(340);
				match(EQUAL);
				setState(341);
				expr(0);
				}
			}

			setState(345);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==SEMICOLON) {
				{
				setState(344);
				match(SEMICOLON);
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
	public static class StructFunctionContext extends ParserRuleContext {
		public FuncDeclarationContext funcDeclaration() {
			return getRuleContext(FuncDeclarationContext.class,0);
		}
		public StructFunctionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structFunction; }
	}

	public final StructFunctionContext structFunction() throws RecognitionException {
		StructFunctionContext _localctx = new StructFunctionContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_structFunction);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(348);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__12) {
				{
				setState(347);
				match(T__12);
				}
			}

			setState(350);
			funcDeclaration();
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
		enterRule(_localctx, 34, RULE_structAccess);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(352);
			match(ID);
			setState(355); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(353);
				match(T__13);
				setState(354);
				match(ID);
				}
				}
				setState(357); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==T__13 );
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
		enterRule(_localctx, 36, RULE_structAssignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(ID);
			setState(360);
			match(T__13);
			setState(361);
			match(ID);
			setState(362);
			match(EQUAL);
			setState(363);
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
		enterRule(_localctx, 38, RULE_funcDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(365);
			match(FUNC);
			setState(366);
			match(ID);
			setState(367);
			match(T__4);
			setState(369);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==UNDERSCORE || _la==ID) {
				{
				setState(368);
				paramsFuncDeclaration();
				}
			}

			setState(371);
			match(T__5);
			setState(374);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__14) {
				{
				setState(372);
				match(T__14);
				setState(373);
				type();
				}
			}

			setState(376);
			match(T__10);
			setState(377);
			block();
			setState(378);
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
		enterRule(_localctx, 40, RULE_paramsFuncDeclaration);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(381);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				{
				setState(380);
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
			setState(383);
			match(ID);
			setState(384);
			match(COLON);
			setState(386);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INOUT) {
				{
				setState(385);
				match(INOUT);
				}
			}

			setState(390);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
			case FLOAT:
			case STRING:
			case BOOL:
			case CHARACTER:
				{
				setState(388);
				type();
				}
				break;
			case T__0:
				{
				setState(389);
				vectorType();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			setState(394);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__2) {
				{
				setState(392);
				match(T__2);
				setState(393);
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
		enterRule(_localctx, 42, RULE_funcCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(396);
			match(ID);
			setState(397);
			match(T__4);
			setState(399);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 36530844729376L) != 0) || ((((_la - 66)) & ~0x3f) == 0 && ((1L << (_la - 66)) & 15L) != 0)) {
				{
				setState(398);
				paramsFuncCall();
				}
			}

			setState(401);
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
		enterRule(_localctx, 44, RULE_paramsFuncCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				{
				setState(403);
				match(ID);
				setState(404);
				match(COLON);
				}
				break;
			}
			setState(408);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==REFERENCE) {
				{
				setState(407);
				match(REFERENCE);
				}
			}

			setState(410);
			expr(0);
			setState(413);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__2) {
				{
				setState(411);
				match(T__2);
				setState(412);
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
		enterRule(_localctx, 46, RULE_printStmt);
		int _la;
		try {
			setState(432);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(415);
				match(PRINT);
				setState(416);
				match(T__4);
				setState(417);
				expr(0);
				setState(418);
				match(T__5);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(420);
				match(PRINT);
				setState(421);
				match(T__4);
				setState(422);
				expr(0);
				setState(427);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__2) {
					{
					{
					setState(423);
					match(T__2);
					setState(424);
					expr(0);
					}
					}
					setState(429);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(430);
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
		enterRule(_localctx, 48, RULE_ifStmt);
		try {
			setState(458);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,53,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(434);
				match(IF);
				setState(435);
				expr(0);
				setState(436);
				match(T__10);
				setState(437);
				block();
				setState(438);
				match(T__11);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(440);
				match(IF);
				setState(441);
				expr(0);
				setState(442);
				match(T__10);
				setState(443);
				block();
				setState(444);
				match(T__11);
				setState(445);
				match(ELSE);
				setState(446);
				match(T__10);
				setState(447);
				block();
				setState(448);
				match(T__11);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(450);
				match(IF);
				setState(451);
				expr(0);
				setState(452);
				match(T__10);
				setState(453);
				block();
				setState(454);
				match(T__11);
				setState(455);
				match(ELSE);
				setState(456);
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
		enterRule(_localctx, 50, RULE_switchStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(460);
			match(SWITCH);
			setState(461);
			expr(0);
			setState(462);
			match(T__10);
			setState(463);
			cases();
			setState(464);
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
		enterRule(_localctx, 52, RULE_cases);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(470);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE || _la==DEFAULT) {
				{
				setState(468);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case CASE:
					{
					setState(466);
					caseStmt();
					}
					break;
				case DEFAULT:
					{
					setState(467);
					defaultStmt();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(472);
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
		enterRule(_localctx, 54, RULE_caseStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(473);
			match(CASE);
			setState(474);
			expr(0);
			setState(475);
			match(COLON);
			setState(476);
			block();
			setState(478);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(477);
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
		enterRule(_localctx, 56, RULE_defaultStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(480);
			match(DEFAULT);
			setState(481);
			match(COLON);
			setState(482);
			block();
			setState(484);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==BREAK) {
				{
				setState(483);
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
		enterRule(_localctx, 58, RULE_whileStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(486);
			match(WHILE);
			setState(487);
			expr(0);
			setState(488);
			match(T__10);
			setState(489);
			block();
			setState(490);
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
		enterRule(_localctx, 60, RULE_forStmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(492);
			match(FOR);
			setState(493);
			_la = _input.LA(1);
			if ( !(_la==UNDERSCORE || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(494);
			match(IN);
			setState(497);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,58,_ctx) ) {
			case 1:
				{
				setState(495);
				expr(0);
				}
				break;
			case 2:
				{
				setState(496);
				range();
				}
				break;
			}
			setState(499);
			match(T__10);
			setState(500);
			block();
			setState(501);
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
		enterRule(_localctx, 62, RULE_range);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(503);
			expr(0);
			setState(504);
			match(T__15);
			setState(505);
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
		enterRule(_localctx, 64, RULE_guardStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(507);
			match(GUARD);
			setState(508);
			expr(0);
			setState(509);
			match(ELSE);
			setState(510);
			match(T__10);
			setState(511);
			block();
			setState(512);
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
		enterRule(_localctx, 66, RULE_breakStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(514);
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
		enterRule(_localctx, 68, RULE_continueStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(516);
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
		enterRule(_localctx, 70, RULE_returnStmt);
		try {
			setState(521);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,59,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(518);
				match(RETURN);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(519);
				match(RETURN);
				setState(520);
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
	@SuppressWarnings("CheckReturnValue")
	public static class FuncCallExprContext extends ExprContext {
		public FuncCallContext funcCall() {
			return getRuleContext(FuncCallContext.class,0);
		}
		public FuncCallExprContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 72;
		enterRecursionRule(_localctx, 72, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(561);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,60,_ctx) ) {
			case 1:
				{
				_localctx = new GroupExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(524);
				match(T__4);
				setState(525);
				expr(0);
				setState(526);
				match(T__5);
				}
				break;
			case 2:
				{
				_localctx = new UnaryOperationExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(528);
				((UnaryOperationExprContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__16 || _la==T__17) ) {
					((UnaryOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(529);
				((UnaryOperationExprContext)_localctx).right = expr(20);
				}
				break;
			case 3:
				{
				_localctx = new BooleanExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(530);
				_la = _input.LA(1);
				if ( !(_la==T__30 || _la==T__31) ) {
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
				setState(531);
				match(NIL);
				}
				break;
			case 5:
				{
				_localctx = new DigitExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(532);
				match(DIGIT);
				}
				break;
			case 6:
				{
				_localctx = new StringExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(533);
				match(STR);
				}
				break;
			case 7:
				{
				_localctx = new CharacterExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(534);
				match(CHAR);
				}
				break;
			case 8:
				{
				_localctx = new FuncCallExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(535);
				funcCall();
				}
				break;
			case 9:
				{
				_localctx = new CountExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(536);
				match(ID);
				setState(537);
				match(T__32);
				}
				break;
			case 10:
				{
				_localctx = new IsEmptyExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(538);
				match(ID);
				setState(539);
				match(T__33);
				}
				break;
			case 11:
				{
				_localctx = new VectorAccessExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(540);
				match(ID);
				setState(541);
				match(T__0);
				setState(542);
				expr(0);
				setState(543);
				match(T__1);
				}
				break;
			case 12:
				{
				_localctx = new IdExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(545);
				match(ID);
				}
				break;
			case 13:
				{
				_localctx = new IntEmbededExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(546);
				match(INT);
				setState(547);
				match(T__4);
				setState(548);
				expr(0);
				setState(549);
				match(T__5);
				}
				break;
			case 14:
				{
				_localctx = new FloatEmbededExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(551);
				match(FLOAT);
				setState(552);
				match(T__4);
				setState(553);
				expr(0);
				setState(554);
				match(T__5);
				}
				break;
			case 15:
				{
				_localctx = new StringEmbededExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(556);
				match(STRING);
				setState(557);
				match(T__4);
				setState(558);
				expr(0);
				setState(559);
				match(T__5);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(583);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,62,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(581);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,61,_ctx) ) {
					case 1:
						{
						_localctx = new ArithmeticOperationExprContext(new ExprContext(_parentctx, _parentState));
						((ArithmeticOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(563);
						if (!(precpred(_ctx, 19))) throw new FailedPredicateException(this, "precpred(_ctx, 19)");
						setState(564);
						((ArithmeticOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3670016L) != 0)) ) {
							((ArithmeticOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(565);
						((ArithmeticOperationExprContext)_localctx).right = expr(20);
						}
						break;
					case 2:
						{
						_localctx = new ArithmeticOperationExprContext(new ExprContext(_parentctx, _parentState));
						((ArithmeticOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(566);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(567);
						((ArithmeticOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__17 || _la==T__21) ) {
							((ArithmeticOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(568);
						((ArithmeticOperationExprContext)_localctx).right = expr(19);
						}
						break;
					case 3:
						{
						_localctx = new RelationalOperationExprContext(new ExprContext(_parentctx, _parentState));
						((RelationalOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(569);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(570);
						((RelationalOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__22 || _la==T__23) ) {
							((RelationalOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(571);
						((RelationalOperationExprContext)_localctx).right = expr(18);
						}
						break;
					case 4:
						{
						_localctx = new RelationalOperationExprContext(new ExprContext(_parentctx, _parentState));
						((RelationalOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(572);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(573);
						((RelationalOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__24 || _la==T__25) ) {
							((RelationalOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(574);
						((RelationalOperationExprContext)_localctx).right = expr(17);
						}
						break;
					case 5:
						{
						_localctx = new ComparisonOperationExprContext(new ExprContext(_parentctx, _parentState));
						((ComparisonOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(575);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(576);
						((ComparisonOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__26 || _la==T__27) ) {
							((ComparisonOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(577);
						((ComparisonOperationExprContext)_localctx).right = expr(16);
						}
						break;
					case 6:
						{
						_localctx = new LogicalOperationExprContext(new ExprContext(_parentctx, _parentState));
						((LogicalOperationExprContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(578);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(579);
						((LogicalOperationExprContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__28 || _la==T__29) ) {
							((LogicalOperationExprContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(580);
						((LogicalOperationExprContext)_localctx).right = expr(15);
						}
						break;
					}
					} 
				}
				setState(585);
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
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(586);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 1065151889408L) != 0)) ) {
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
		case 36:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 19);
		case 1:
			return precpred(_ctx, 18);
		case 2:
			return precpred(_ctx, 17);
		case 3:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 15);
		case 5:
			return precpred(_ctx, 14);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001H\u024d\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0001\u0000\u0001\u0000\u0001\u0000"+
		"\u0001\u0001\u0005\u0001Q\b\u0001\n\u0001\f\u0001T\t\u0001\u0001\u0002"+
		"\u0001\u0002\u0003\u0002X\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002"+
		"\\\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002`\b\u0002\u0001\u0002\u0001"+
		"\u0002\u0003\u0002d\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002h\b\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002l\b\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002p\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002t\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002x\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002|\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0080\b\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002\u0084\b\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002\u0088\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u008c\b\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002\u0090\b\u0002\u0001\u0002\u0001\u0002"+
		"\u0003\u0002\u0094\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u0098\b"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002\u009c\b\u0002\u0003\u0002\u009e"+
		"\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003"+
		"\u0003\u00b1\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u00b9\b\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u00c3\b\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u00cd\b\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004\u00d1\b\u0004\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u00de\b\u0005\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006\u00e9\b\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0003\u0006\u00f6\b\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0005\u0007\u00fb\b\u0007\n\u0007\f\u0007\u00fe\t\u0007"+
		"\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0003\f\u0122\b\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0003"+
		"\f\u0131\b\f\u0001\r\u0001\r\u0003\r\u0135\b\r\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0003\r\u013b\b\r\u0001\r\u0005\r\u013e\b\r\n\r\f\r\u0141\t\r"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e"+
		"\u0148\b\u000e\n\u000e\f\u000e\u014b\t\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0153\b\u000f\u0001"+
		"\u000f\u0001\u000f\u0003\u000f\u0157\b\u000f\u0001\u000f\u0003\u000f\u015a"+
		"\b\u000f\u0001\u0010\u0003\u0010\u015d\b\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0011\u0001\u0011\u0001\u0011\u0004\u0011\u0164\b\u0011\u000b\u0011"+
		"\f\u0011\u0165\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013"+
		"\u0172\b\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u0177\b"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0014\u0003"+
		"\u0014\u017e\b\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u0183"+
		"\b\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u0187\b\u0014\u0001\u0014"+
		"\u0001\u0014\u0003\u0014\u018b\b\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0003\u0015\u0190\b\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001\u0016"+
		"\u0003\u0016\u0196\b\u0016\u0001\u0016\u0003\u0016\u0199\b\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0003\u0016\u019e\b\u0016\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001\u0017\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u01aa\b\u0017\n\u0017\f\u0017"+
		"\u01ad\t\u0017\u0001\u0017\u0001\u0017\u0003\u0017\u01b1\b\u0017\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0003"+
		"\u0018\u01cb\b\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0005\u001a\u01d5\b\u001a\n"+
		"\u001a\f\u001a\u01d8\t\u001a\u0001\u001b\u0001\u001b\u0001\u001b\u0001"+
		"\u001b\u0001\u001b\u0003\u001b\u01df\b\u001b\u0001\u001c\u0001\u001c\u0001"+
		"\u001c\u0001\u001c\u0003\u001c\u01e5\b\u001c\u0001\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u01f2\b\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001f\u0001\u001f\u0001\u001f\u0001"+
		"\u001f\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0001!\u0001!"+
		"\u0001\"\u0001\"\u0001#\u0001#\u0001#\u0003#\u020a\b#\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0003$\u0232\b$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0001$\u0005$\u0246\b$\n$\f$\u0249\t$\u0001"+
		"%\u0001%\u0001%\u0000\u0001H&\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010"+
		"\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJ\u0000"+
		"\u000b\u0001\u0000)*\u0002\u0000AAEE\u0001\u0000\u0011\u0012\u0001\u0000"+
		"\u001f \u0001\u0000\u0013\u0015\u0002\u0000\u0012\u0012\u0016\u0016\u0001"+
		"\u0000\u0017\u0018\u0001\u0000\u0019\u001a\u0001\u0000\u001b\u001c\u0001"+
		"\u0000\u001d\u001e\u0001\u0000#\'\u0289\u0000L\u0001\u0000\u0000\u0000"+
		"\u0002R\u0001\u0000\u0000\u0000\u0004\u009d\u0001\u0000\u0000\u0000\u0006"+
		"\u00b0\u0001\u0000\u0000\u0000\b\u00d0\u0001\u0000\u0000\u0000\n\u00dd"+
		"\u0001\u0000\u0000\u0000\f\u00f5\u0001\u0000\u0000\u0000\u000e\u00f7\u0001"+
		"\u0000\u0000\u0000\u0010\u00ff\u0001\u0000\u0000\u0000\u0012\u0103\u0001"+
		"\u0000\u0000\u0000\u0014\u0109\u0001\u0000\u0000\u0000\u0016\u010e\u0001"+
		"\u0000\u0000\u0000\u0018\u0130\u0001\u0000\u0000\u0000\u001a\u0132\u0001"+
		"\u0000\u0000\u0000\u001c\u0142\u0001\u0000\u0000\u0000\u001e\u014e\u0001"+
		"\u0000\u0000\u0000 \u015c\u0001\u0000\u0000\u0000\"\u0160\u0001\u0000"+
		"\u0000\u0000$\u0167\u0001\u0000\u0000\u0000&\u016d\u0001\u0000\u0000\u0000"+
		"(\u017d\u0001\u0000\u0000\u0000*\u018c\u0001\u0000\u0000\u0000,\u0195"+
		"\u0001\u0000\u0000\u0000.\u01b0\u0001\u0000\u0000\u00000\u01ca\u0001\u0000"+
		"\u0000\u00002\u01cc\u0001\u0000\u0000\u00004\u01d6\u0001\u0000\u0000\u0000"+
		"6\u01d9\u0001\u0000\u0000\u00008\u01e0\u0001\u0000\u0000\u0000:\u01e6"+
		"\u0001\u0000\u0000\u0000<\u01ec\u0001\u0000\u0000\u0000>\u01f7\u0001\u0000"+
		"\u0000\u0000@\u01fb\u0001\u0000\u0000\u0000B\u0202\u0001\u0000\u0000\u0000"+
		"D\u0204\u0001\u0000\u0000\u0000F\u0209\u0001\u0000\u0000\u0000H\u0231"+
		"\u0001\u0000\u0000\u0000J\u024a\u0001\u0000\u0000\u0000LM\u0003\u0002"+
		"\u0001\u0000MN\u0005\u0000\u0000\u0001N\u0001\u0001\u0000\u0000\u0000"+
		"OQ\u0003\u0004\u0002\u0000PO\u0001\u0000\u0000\u0000QT\u0001\u0000\u0000"+
		"\u0000RP\u0001\u0000\u0000\u0000RS\u0001\u0000\u0000\u0000S\u0003\u0001"+
		"\u0000\u0000\u0000TR\u0001\u0000\u0000\u0000UW\u0003\u0006\u0003\u0000"+
		"VX\u0005;\u0000\u0000WV\u0001\u0000\u0000\u0000WX\u0001\u0000\u0000\u0000"+
		"X\u009e\u0001\u0000\u0000\u0000Y[\u0003\b\u0004\u0000Z\\\u0005;\u0000"+
		"\u0000[Z\u0001\u0000\u0000\u0000[\\\u0001\u0000\u0000\u0000\\\u009e\u0001"+
		"\u0000\u0000\u0000]_\u0003\n\u0005\u0000^`\u0005;\u0000\u0000_^\u0001"+
		"\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`\u009e\u0001\u0000\u0000"+
		"\u0000ac\u0003\f\u0006\u0000bd\u0005;\u0000\u0000cb\u0001\u0000\u0000"+
		"\u0000cd\u0001\u0000\u0000\u0000d\u009e\u0001\u0000\u0000\u0000eg\u0003"+
		"\u0012\t\u0000fh\u0005;\u0000\u0000gf\u0001\u0000\u0000\u0000gh\u0001"+
		"\u0000\u0000\u0000h\u009e\u0001\u0000\u0000\u0000ik\u0003\u0014\n\u0000"+
		"jl\u0005;\u0000\u0000kj\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000"+
		"l\u009e\u0001\u0000\u0000\u0000mo\u0003\u0016\u000b\u0000np\u0005;\u0000"+
		"\u0000on\u0001\u0000\u0000\u0000op\u0001\u0000\u0000\u0000p\u009e\u0001"+
		"\u0000\u0000\u0000qs\u0003&\u0013\u0000rt\u0005;\u0000\u0000sr\u0001\u0000"+
		"\u0000\u0000st\u0001\u0000\u0000\u0000t\u009e\u0001\u0000\u0000\u0000"+
		"uw\u0003*\u0015\u0000vx\u0005;\u0000\u0000wv\u0001\u0000\u0000\u0000w"+
		"x\u0001\u0000\u0000\u0000x\u009e\u0001\u0000\u0000\u0000y{\u0003.\u0017"+
		"\u0000z|\u0005;\u0000\u0000{z\u0001\u0000\u0000\u0000{|\u0001\u0000\u0000"+
		"\u0000|\u009e\u0001\u0000\u0000\u0000}\u007f\u00030\u0018\u0000~\u0080"+
		"\u0005;\u0000\u0000\u007f~\u0001\u0000\u0000\u0000\u007f\u0080\u0001\u0000"+
		"\u0000\u0000\u0080\u009e\u0001\u0000\u0000\u0000\u0081\u0083\u00032\u0019"+
		"\u0000\u0082\u0084\u0005;\u0000\u0000\u0083\u0082\u0001\u0000\u0000\u0000"+
		"\u0083\u0084\u0001\u0000\u0000\u0000\u0084\u009e\u0001\u0000\u0000\u0000"+
		"\u0085\u0087\u0003:\u001d\u0000\u0086\u0088\u0005;\u0000\u0000\u0087\u0086"+
		"\u0001\u0000\u0000\u0000\u0087\u0088\u0001\u0000\u0000\u0000\u0088\u009e"+
		"\u0001\u0000\u0000\u0000\u0089\u008b\u0003<\u001e\u0000\u008a\u008c\u0005"+
		";\u0000\u0000\u008b\u008a\u0001\u0000\u0000\u0000\u008b\u008c\u0001\u0000"+
		"\u0000\u0000\u008c\u009e\u0001\u0000\u0000\u0000\u008d\u008f\u0003@ \u0000"+
		"\u008e\u0090\u0005;\u0000\u0000\u008f\u008e\u0001\u0000\u0000\u0000\u008f"+
		"\u0090\u0001\u0000\u0000\u0000\u0090\u009e\u0001\u0000\u0000\u0000\u0091"+
		"\u0093\u0003B!\u0000\u0092\u0094\u0005;\u0000\u0000\u0093\u0092\u0001"+
		"\u0000\u0000\u0000\u0093\u0094\u0001\u0000\u0000\u0000\u0094\u009e\u0001"+
		"\u0000\u0000\u0000\u0095\u0097\u0003D\"\u0000\u0096\u0098\u0005;\u0000"+
		"\u0000\u0097\u0096\u0001\u0000\u0000\u0000\u0097\u0098\u0001\u0000\u0000"+
		"\u0000\u0098\u009e\u0001\u0000\u0000\u0000\u0099\u009b\u0003F#\u0000\u009a"+
		"\u009c\u0005;\u0000\u0000\u009b\u009a\u0001\u0000\u0000\u0000\u009b\u009c"+
		"\u0001\u0000\u0000\u0000\u009c\u009e\u0001\u0000\u0000\u0000\u009dU\u0001"+
		"\u0000\u0000\u0000\u009dY\u0001\u0000\u0000\u0000\u009d]\u0001\u0000\u0000"+
		"\u0000\u009da\u0001\u0000\u0000\u0000\u009de\u0001\u0000\u0000\u0000\u009d"+
		"i\u0001\u0000\u0000\u0000\u009dm\u0001\u0000\u0000\u0000\u009dq\u0001"+
		"\u0000\u0000\u0000\u009du\u0001\u0000\u0000\u0000\u009dy\u0001\u0000\u0000"+
		"\u0000\u009d}\u0001\u0000\u0000\u0000\u009d\u0081\u0001\u0000\u0000\u0000"+
		"\u009d\u0085\u0001\u0000\u0000\u0000\u009d\u0089\u0001\u0000\u0000\u0000"+
		"\u009d\u008d\u0001\u0000\u0000\u0000\u009d\u0091\u0001\u0000\u0000\u0000"+
		"\u009d\u0095\u0001\u0000\u0000\u0000\u009d\u0099\u0001\u0000\u0000\u0000"+
		"\u009e\u0005\u0001\u0000\u0000\u0000\u009f\u00a0\u0005)\u0000\u0000\u00a0"+
		"\u00a1\u0005E\u0000\u0000\u00a1\u00a2\u0005<\u0000\u0000\u00a2\u00a3\u0003"+
		"J%\u0000\u00a3\u00a4\u0005=\u0000\u0000\u00a4\u00a5\u0003H$\u0000\u00a5"+
		"\u00b1\u0001\u0000\u0000\u0000\u00a6\u00a7\u0005)\u0000\u0000\u00a7\u00a8"+
		"\u0005E\u0000\u0000\u00a8\u00a9\u0005=\u0000\u0000\u00a9\u00b1\u0003H"+
		"$\u0000\u00aa\u00ab\u0005)\u0000\u0000\u00ab\u00ac\u0005E\u0000\u0000"+
		"\u00ac\u00ad\u0005<\u0000\u0000\u00ad\u00ae\u0003J%\u0000\u00ae\u00af"+
		"\u0005>\u0000\u0000\u00af\u00b1\u0001\u0000\u0000\u0000\u00b0\u009f\u0001"+
		"\u0000\u0000\u0000\u00b0\u00a6\u0001\u0000\u0000\u0000\u00b0\u00aa\u0001"+
		"\u0000\u0000\u0000\u00b1\u0007\u0001\u0000\u0000\u0000\u00b2\u00b9\u0005"+
		"E\u0000\u0000\u00b3\u00b4\u0005E\u0000\u0000\u00b4\u00b5\u0005\u0001\u0000"+
		"\u0000\u00b5\u00b6\u0003H$\u0000\u00b6\u00b7\u0005\u0002\u0000\u0000\u00b7"+
		"\u00b9\u0001\u0000\u0000\u0000\u00b8\u00b2\u0001\u0000\u0000\u0000\u00b8"+
		"\u00b3\u0001\u0000\u0000\u0000\u00b9\u00ba\u0001\u0000\u0000\u0000\u00ba"+
		"\u00bb\u0005=\u0000\u0000\u00bb\u00d1\u0003H$\u0000\u00bc\u00c3\u0005"+
		"E\u0000\u0000\u00bd\u00be\u0005E\u0000\u0000\u00be\u00bf\u0005\u0001\u0000"+
		"\u0000\u00bf\u00c0\u0003H$\u0000\u00c0\u00c1\u0005\u0002\u0000\u0000\u00c1"+
		"\u00c3\u0001\u0000\u0000\u0000\u00c2\u00bc\u0001\u0000\u0000\u0000\u00c2"+
		"\u00bd\u0001\u0000\u0000\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4"+
		"\u00c5\u0005?\u0000\u0000\u00c5\u00d1\u0003H$\u0000\u00c6\u00cd\u0005"+
		"E\u0000\u0000\u00c7\u00c8\u0005E\u0000\u0000\u00c8\u00c9\u0005\u0001\u0000"+
		"\u0000\u00c9\u00ca\u0003H$\u0000\u00ca\u00cb\u0005\u0002\u0000\u0000\u00cb"+
		"\u00cd\u0001\u0000\u0000\u0000\u00cc\u00c6\u0001\u0000\u0000\u0000\u00cc"+
		"\u00c7\u0001\u0000\u0000\u0000\u00cd\u00ce\u0001\u0000\u0000\u0000\u00ce"+
		"\u00cf\u0005@\u0000\u0000\u00cf\u00d1\u0003H$\u0000\u00d0\u00b8\u0001"+
		"\u0000\u0000\u0000\u00d0\u00c2\u0001\u0000\u0000\u0000\u00d0\u00cc\u0001"+
		"\u0000\u0000\u0000\u00d1\t\u0001\u0000\u0000\u0000\u00d2\u00d3\u0005*"+
		"\u0000\u0000\u00d3\u00d4\u0005E\u0000\u0000\u00d4\u00d5\u0005<\u0000\u0000"+
		"\u00d5\u00d6\u0003J%\u0000\u00d6\u00d7\u0005=\u0000\u0000\u00d7\u00d8"+
		"\u0003H$\u0000\u00d8\u00de\u0001\u0000\u0000\u0000\u00d9\u00da\u0005*"+
		"\u0000\u0000\u00da\u00db\u0005E\u0000\u0000\u00db\u00dc\u0005=\u0000\u0000"+
		"\u00dc\u00de\u0003H$\u0000\u00dd\u00d2\u0001\u0000\u0000\u0000\u00dd\u00d9"+
		"\u0001\u0000\u0000\u0000\u00de\u000b\u0001\u0000\u0000\u0000\u00df\u00e0"+
		"\u0005)\u0000\u0000\u00e0\u00e1\u0005E\u0000\u0000\u00e1\u00e2\u0005<"+
		"\u0000\u0000\u00e2\u00e3\u0005\u0001\u0000\u0000\u00e3\u00e4\u0003J%\u0000"+
		"\u00e4\u00e5\u0005\u0002\u0000\u0000\u00e5\u00e6\u0005=\u0000\u0000\u00e6"+
		"\u00e8\u0005\u0001\u0000\u0000\u00e7\u00e9\u0003\u000e\u0007\u0000\u00e8"+
		"\u00e7\u0001\u0000\u0000\u0000\u00e8\u00e9\u0001\u0000\u0000\u0000\u00e9"+
		"\u00ea\u0001\u0000\u0000\u0000\u00ea\u00eb\u0005\u0002\u0000\u0000\u00eb"+
		"\u00f6\u0001\u0000\u0000\u0000\u00ec\u00ed\u0005)\u0000\u0000\u00ed\u00ee"+
		"\u0005E\u0000\u0000\u00ee\u00ef\u0005<\u0000\u0000\u00ef\u00f0\u0005\u0001"+
		"\u0000\u0000\u00f0\u00f1\u0003J%\u0000\u00f1\u00f2\u0005\u0002\u0000\u0000"+
		"\u00f2\u00f3\u0005=\u0000\u0000\u00f3\u00f4\u0005E\u0000\u0000\u00f4\u00f6"+
		"\u0001\u0000\u0000\u0000\u00f5\u00df\u0001\u0000\u0000\u0000\u00f5\u00ec"+
		"\u0001\u0000\u0000\u0000\u00f6\r\u0001\u0000\u0000\u0000\u00f7\u00fc\u0003"+
		"H$\u0000\u00f8\u00f9\u0005\u0003\u0000\u0000\u00f9\u00fb\u0003H$\u0000"+
		"\u00fa\u00f8\u0001\u0000\u0000\u0000\u00fb\u00fe\u0001\u0000\u0000\u0000"+
		"\u00fc\u00fa\u0001\u0000\u0000\u0000\u00fc\u00fd\u0001\u0000\u0000\u0000"+
		"\u00fd\u000f\u0001\u0000\u0000\u0000\u00fe\u00fc\u0001\u0000\u0000\u0000"+
		"\u00ff\u0100\u0005\u0001\u0000\u0000\u0100\u0101\u0003J%\u0000\u0101\u0102"+
		"\u0005\u0002\u0000\u0000\u0102\u0011\u0001\u0000\u0000\u0000\u0103\u0104"+
		"\u0005E\u0000\u0000\u0104\u0105\u0005\u0004\u0000\u0000\u0105\u0106\u0005"+
		"\u0005\u0000\u0000\u0106\u0107\u0003H$\u0000\u0107\u0108\u0005\u0006\u0000"+
		"\u0000\u0108\u0013\u0001\u0000\u0000\u0000\u0109\u010a\u0005E\u0000\u0000"+
		"\u010a\u010b\u0005\u0007\u0000\u0000\u010b\u010c\u0005\u0005\u0000\u0000"+
		"\u010c\u010d\u0005\u0006\u0000\u0000\u010d\u0015\u0001\u0000\u0000\u0000"+
		"\u010e\u010f\u0005E\u0000\u0000\u010f\u0110\u0005\b\u0000\u0000\u0110"+
		"\u0111\u0005\u0005\u0000\u0000\u0111\u0112\u0005\t\u0000\u0000\u0112\u0113"+
		"\u0005<\u0000\u0000\u0113\u0114\u0003H$\u0000\u0114\u0115\u0005\u0006"+
		"\u0000\u0000\u0115\u0017\u0001\u0000\u0000\u0000\u0116\u0117\u0005)\u0000"+
		"\u0000\u0117\u0118\u0005E\u0000\u0000\u0118\u0119\u0005<\u0000\u0000\u0119"+
		"\u011a\u0005\u0001\u0000\u0000\u011a\u011b\u0005\u0001\u0000\u0000\u011b"+
		"\u011c\u0003J%\u0000\u011c\u011d\u0005\u0002\u0000\u0000\u011d\u011e\u0005"+
		"\u0002\u0000\u0000\u011e\u011f\u0005=\u0000\u0000\u011f\u0121\u0005\u0001"+
		"\u0000\u0000\u0120\u0122\u0003\u001a\r\u0000\u0121\u0120\u0001\u0000\u0000"+
		"\u0000\u0121\u0122\u0001\u0000\u0000\u0000\u0122\u0123\u0001\u0000\u0000"+
		"\u0000\u0123\u0124\u0005\u0002\u0000\u0000\u0124\u0131\u0001\u0000\u0000"+
		"\u0000\u0125\u0126\u0005)\u0000\u0000\u0126\u0127\u0005E\u0000\u0000\u0127"+
		"\u0128\u0005<\u0000\u0000\u0128\u0129\u0005\u0001\u0000\u0000\u0129\u012a"+
		"\u0005\u0001\u0000\u0000\u012a\u012b\u0003J%\u0000\u012b\u012c\u0005\u0002"+
		"\u0000\u0000\u012c\u012d\u0005\u0002\u0000\u0000\u012d\u012e\u0005=\u0000"+
		"\u0000\u012e\u012f\u0005E\u0000\u0000\u012f\u0131\u0001\u0000\u0000\u0000"+
		"\u0130\u0116\u0001\u0000\u0000\u0000\u0130\u0125\u0001\u0000\u0000\u0000"+
		"\u0131\u0019\u0001\u0000\u0000\u0000\u0132\u0134\u0005\u0001\u0000\u0000"+
		"\u0133\u0135\u0003\u000e\u0007\u0000\u0134\u0133\u0001\u0000\u0000\u0000"+
		"\u0134\u0135\u0001\u0000\u0000\u0000\u0135\u0136\u0001\u0000\u0000\u0000"+
		"\u0136\u013f\u0005\u0002\u0000\u0000\u0137\u0138\u0005\u0003\u0000\u0000"+
		"\u0138\u013a\u0005\u0001\u0000\u0000\u0139\u013b\u0003\u000e\u0007\u0000"+
		"\u013a\u0139\u0001\u0000\u0000\u0000\u013a\u013b\u0001\u0000\u0000\u0000"+
		"\u013b\u013c\u0001\u0000\u0000\u0000\u013c\u013e\u0005\u0002\u0000\u0000"+
		"\u013d\u0137\u0001\u0000\u0000\u0000\u013e\u0141\u0001\u0000\u0000\u0000"+
		"\u013f\u013d\u0001\u0000\u0000\u0000\u013f\u0140\u0001\u0000\u0000\u0000"+
		"\u0140\u001b\u0001\u0000\u0000\u0000\u0141\u013f\u0001\u0000\u0000\u0000"+
		"\u0142\u0143\u0005\n\u0000\u0000\u0143\u0144\u0005E\u0000\u0000\u0144"+
		"\u0149\u0005\u000b\u0000\u0000\u0145\u0148\u0003\u001e\u000f\u0000\u0146"+
		"\u0148\u0003 \u0010\u0000\u0147\u0145\u0001\u0000\u0000\u0000\u0147\u0146"+
		"\u0001\u0000\u0000\u0000\u0148\u014b\u0001\u0000\u0000\u0000\u0149\u0147"+
		"\u0001\u0000\u0000\u0000\u0149\u014a\u0001\u0000\u0000\u0000\u014a\u014c"+
		"\u0001\u0000\u0000\u0000\u014b\u0149\u0001\u0000\u0000\u0000\u014c\u014d"+
		"\u0005\f\u0000\u0000\u014d\u001d\u0001\u0000\u0000\u0000\u014e\u014f\u0007"+
		"\u0000\u0000\u0000\u014f\u0152\u0005E\u0000\u0000\u0150\u0151\u0005<\u0000"+
		"\u0000\u0151\u0153\u0003J%\u0000\u0152\u0150\u0001\u0000\u0000\u0000\u0152"+
		"\u0153\u0001\u0000\u0000\u0000\u0153\u0156\u0001\u0000\u0000\u0000\u0154"+
		"\u0155\u0005=\u0000\u0000\u0155\u0157\u0003H$\u0000\u0156\u0154\u0001"+
		"\u0000\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000\u0157\u0159\u0001"+
		"\u0000\u0000\u0000\u0158\u015a\u0005;\u0000\u0000\u0159\u0158\u0001\u0000"+
		"\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000\u015a\u001f\u0001\u0000"+
		"\u0000\u0000\u015b\u015d\u0005\r\u0000\u0000\u015c\u015b\u0001\u0000\u0000"+
		"\u0000\u015c\u015d\u0001\u0000\u0000\u0000\u015d\u015e\u0001\u0000\u0000"+
		"\u0000\u015e\u015f\u0003&\u0013\u0000\u015f!\u0001\u0000\u0000\u0000\u0160"+
		"\u0163\u0005E\u0000\u0000\u0161\u0162\u0005\u000e\u0000\u0000\u0162\u0164"+
		"\u0005E\u0000\u0000\u0163\u0161\u0001\u0000\u0000\u0000\u0164\u0165\u0001"+
		"\u0000\u0000\u0000\u0165\u0163\u0001\u0000\u0000\u0000\u0165\u0166\u0001"+
		"\u0000\u0000\u0000\u0166#\u0001\u0000\u0000\u0000\u0167\u0168\u0005E\u0000"+
		"\u0000\u0168\u0169\u0005\u000e\u0000\u0000\u0169\u016a\u0005E\u0000\u0000"+
		"\u016a\u016b\u0005=\u0000\u0000\u016b\u016c\u0003H$\u0000\u016c%\u0001"+
		"\u0000\u0000\u0000\u016d\u016e\u0005+\u0000\u0000\u016e\u016f\u0005E\u0000"+
		"\u0000\u016f\u0171\u0005\u0005\u0000\u0000\u0170\u0172\u0003(\u0014\u0000"+
		"\u0171\u0170\u0001\u0000\u0000\u0000\u0171\u0172\u0001\u0000\u0000\u0000"+
		"\u0172\u0173\u0001\u0000\u0000\u0000\u0173\u0176\u0005\u0006\u0000\u0000"+
		"\u0174\u0175\u0005\u000f\u0000\u0000\u0175\u0177\u0003J%\u0000\u0176\u0174"+
		"\u0001\u0000\u0000\u0000\u0176\u0177\u0001\u0000\u0000\u0000\u0177\u0178"+
		"\u0001\u0000\u0000\u0000\u0178\u0179\u0005\u000b\u0000\u0000\u0179\u017a"+
		"\u0003\u0002\u0001\u0000\u017a\u017b\u0005\f\u0000\u0000\u017b\'\u0001"+
		"\u0000\u0000\u0000\u017c\u017e\u0007\u0001\u0000\u0000\u017d\u017c\u0001"+
		"\u0000\u0000\u0000\u017d\u017e\u0001\u0000\u0000\u0000\u017e\u017f\u0001"+
		"\u0000\u0000\u0000\u017f\u0180\u0005E\u0000\u0000\u0180\u0182\u0005<\u0000"+
		"\u0000\u0181\u0183\u0005,\u0000\u0000\u0182\u0181\u0001\u0000\u0000\u0000"+
		"\u0182\u0183\u0001\u0000\u0000\u0000\u0183\u0186\u0001\u0000\u0000\u0000"+
		"\u0184\u0187\u0003J%\u0000\u0185\u0187\u0003\u0010\b\u0000\u0186\u0184"+
		"\u0001\u0000\u0000\u0000\u0186\u0185\u0001\u0000\u0000\u0000\u0187\u018a"+
		"\u0001\u0000\u0000\u0000\u0188\u0189\u0005\u0003\u0000\u0000\u0189\u018b"+
		"\u0003(\u0014\u0000\u018a\u0188\u0001\u0000\u0000\u0000\u018a\u018b\u0001"+
		"\u0000\u0000\u0000\u018b)\u0001\u0000\u0000\u0000\u018c\u018d\u0005E\u0000"+
		"\u0000\u018d\u018f\u0005\u0005\u0000\u0000\u018e\u0190\u0003,\u0016\u0000"+
		"\u018f\u018e\u0001\u0000\u0000\u0000\u018f\u0190\u0001\u0000\u0000\u0000"+
		"\u0190\u0191\u0001\u0000\u0000\u0000\u0191\u0192\u0005\u0006\u0000\u0000"+
		"\u0192+\u0001\u0000\u0000\u0000\u0193\u0194\u0005E\u0000\u0000\u0194\u0196"+
		"\u0005<\u0000\u0000\u0195\u0193\u0001\u0000\u0000\u0000\u0195\u0196\u0001"+
		"\u0000\u0000\u0000\u0196\u0198\u0001\u0000\u0000\u0000\u0197\u0199\u0005"+
		"-\u0000\u0000\u0198\u0197\u0001\u0000\u0000\u0000\u0198\u0199\u0001\u0000"+
		"\u0000\u0000\u0199\u019a\u0001\u0000\u0000\u0000\u019a\u019d\u0003H$\u0000"+
		"\u019b\u019c\u0005\u0003\u0000\u0000\u019c\u019e\u0003,\u0016\u0000\u019d"+
		"\u019b\u0001\u0000\u0000\u0000\u019d\u019e\u0001\u0000\u0000\u0000\u019e"+
		"-\u0001\u0000\u0000\u0000\u019f\u01a0\u0005.\u0000\u0000\u01a0\u01a1\u0005"+
		"\u0005\u0000\u0000\u01a1\u01a2\u0003H$\u0000\u01a2\u01a3\u0005\u0006\u0000"+
		"\u0000\u01a3\u01b1\u0001\u0000\u0000\u0000\u01a4\u01a5\u0005.\u0000\u0000"+
		"\u01a5\u01a6\u0005\u0005\u0000\u0000\u01a6\u01ab\u0003H$\u0000\u01a7\u01a8"+
		"\u0005\u0003\u0000\u0000\u01a8\u01aa\u0003H$\u0000\u01a9\u01a7\u0001\u0000"+
		"\u0000\u0000\u01aa\u01ad\u0001\u0000\u0000\u0000\u01ab\u01a9\u0001\u0000"+
		"\u0000\u0000\u01ab\u01ac\u0001\u0000\u0000\u0000\u01ac\u01ae\u0001\u0000"+
		"\u0000\u0000\u01ad\u01ab\u0001\u0000\u0000\u0000\u01ae\u01af\u0005\u0006"+
		"\u0000\u0000\u01af\u01b1\u0001\u0000\u0000\u0000\u01b0\u019f\u0001\u0000"+
		"\u0000\u0000\u01b0\u01a4\u0001\u0000\u0000\u0000\u01b1/\u0001\u0000\u0000"+
		"\u0000\u01b2\u01b3\u0005/\u0000\u0000\u01b3\u01b4\u0003H$\u0000\u01b4"+
		"\u01b5\u0005\u000b\u0000\u0000\u01b5\u01b6\u0003\u0002\u0001\u0000\u01b6"+
		"\u01b7\u0005\f\u0000\u0000\u01b7\u01cb\u0001\u0000\u0000\u0000\u01b8\u01b9"+
		"\u0005/\u0000\u0000\u01b9\u01ba\u0003H$\u0000\u01ba\u01bb\u0005\u000b"+
		"\u0000\u0000\u01bb\u01bc\u0003\u0002\u0001\u0000\u01bc\u01bd\u0005\f\u0000"+
		"\u0000\u01bd\u01be\u00050\u0000\u0000\u01be\u01bf\u0005\u000b\u0000\u0000"+
		"\u01bf\u01c0\u0003\u0002\u0001\u0000\u01c0\u01c1\u0005\f\u0000\u0000\u01c1"+
		"\u01cb\u0001\u0000\u0000\u0000\u01c2\u01c3\u0005/\u0000\u0000\u01c3\u01c4"+
		"\u0003H$\u0000\u01c4\u01c5\u0005\u000b\u0000\u0000\u01c5\u01c6\u0003\u0002"+
		"\u0001\u0000\u01c6\u01c7\u0005\f\u0000\u0000\u01c7\u01c8\u00050\u0000"+
		"\u0000\u01c8\u01c9\u00030\u0018\u0000\u01c9\u01cb\u0001\u0000\u0000\u0000"+
		"\u01ca\u01b2\u0001\u0000\u0000\u0000\u01ca\u01b8\u0001\u0000\u0000\u0000"+
		"\u01ca\u01c2\u0001\u0000\u0000\u0000\u01cb1\u0001\u0000\u0000\u0000\u01cc"+
		"\u01cd\u00051\u0000\u0000\u01cd\u01ce\u0003H$\u0000\u01ce\u01cf\u0005"+
		"\u000b\u0000\u0000\u01cf\u01d0\u00034\u001a\u0000\u01d0\u01d1\u0005\f"+
		"\u0000\u0000\u01d13\u0001\u0000\u0000\u0000\u01d2\u01d5\u00036\u001b\u0000"+
		"\u01d3\u01d5\u00038\u001c\u0000\u01d4\u01d2\u0001\u0000\u0000\u0000\u01d4"+
		"\u01d3\u0001\u0000\u0000\u0000\u01d5\u01d8\u0001\u0000\u0000\u0000\u01d6"+
		"\u01d4\u0001\u0000\u0000\u0000\u01d6\u01d7\u0001\u0000\u0000\u0000\u01d7"+
		"5\u0001\u0000\u0000\u0000\u01d8\u01d6\u0001\u0000\u0000\u0000\u01d9\u01da"+
		"\u00052\u0000\u0000\u01da\u01db\u0003H$\u0000\u01db\u01dc\u0005<\u0000"+
		"\u0000\u01dc\u01de\u0003\u0002\u0001\u0000\u01dd\u01df\u00059\u0000\u0000"+
		"\u01de\u01dd\u0001\u0000\u0000\u0000\u01de\u01df\u0001\u0000\u0000\u0000"+
		"\u01df7\u0001\u0000\u0000\u0000\u01e0\u01e1\u00053\u0000\u0000\u01e1\u01e2"+
		"\u0005<\u0000\u0000\u01e2\u01e4\u0003\u0002\u0001\u0000\u01e3\u01e5\u0005"+
		"9\u0000\u0000\u01e4\u01e3\u0001\u0000\u0000\u0000\u01e4\u01e5\u0001\u0000"+
		"\u0000\u0000\u01e59\u0001\u0000\u0000\u0000\u01e6\u01e7\u00054\u0000\u0000"+
		"\u01e7\u01e8\u0003H$\u0000\u01e8\u01e9\u0005\u000b\u0000\u0000\u01e9\u01ea"+
		"\u0003\u0002\u0001\u0000\u01ea\u01eb\u0005\f\u0000\u0000\u01eb;\u0001"+
		"\u0000\u0000\u0000\u01ec\u01ed\u00055\u0000\u0000\u01ed\u01ee\u0007\u0001"+
		"\u0000\u0000\u01ee\u01f1\u00056\u0000\u0000\u01ef\u01f2\u0003H$\u0000"+
		"\u01f0\u01f2\u0003>\u001f\u0000\u01f1\u01ef\u0001\u0000\u0000\u0000\u01f1"+
		"\u01f0\u0001\u0000\u0000\u0000\u01f2\u01f3\u0001\u0000\u0000\u0000\u01f3"+
		"\u01f4\u0005\u000b\u0000\u0000\u01f4\u01f5\u0003\u0002\u0001\u0000\u01f5"+
		"\u01f6\u0005\f\u0000\u0000\u01f6=\u0001\u0000\u0000\u0000\u01f7\u01f8"+
		"\u0003H$\u0000\u01f8\u01f9\u0005\u0010\u0000\u0000\u01f9\u01fa\u0003H"+
		"$\u0000\u01fa?\u0001\u0000\u0000\u0000\u01fb\u01fc\u00057\u0000\u0000"+
		"\u01fc\u01fd\u0003H$\u0000\u01fd\u01fe\u00050\u0000\u0000\u01fe\u01ff"+
		"\u0005\u000b\u0000\u0000\u01ff\u0200\u0003\u0002\u0001\u0000\u0200\u0201"+
		"\u0005\f\u0000\u0000\u0201A\u0001\u0000\u0000\u0000\u0202\u0203\u0005"+
		"9\u0000\u0000\u0203C\u0001\u0000\u0000\u0000\u0204\u0205\u00058\u0000"+
		"\u0000\u0205E\u0001\u0000\u0000\u0000\u0206\u020a\u0005:\u0000\u0000\u0207"+
		"\u0208\u0005:\u0000\u0000\u0208\u020a\u0003H$\u0000\u0209\u0206\u0001"+
		"\u0000\u0000\u0000\u0209\u0207\u0001\u0000\u0000\u0000\u020aG\u0001\u0000"+
		"\u0000\u0000\u020b\u020c\u0006$\uffff\uffff\u0000\u020c\u020d\u0005\u0005"+
		"\u0000\u0000\u020d\u020e\u0003H$\u0000\u020e\u020f\u0005\u0006\u0000\u0000"+
		"\u020f\u0232\u0001\u0000\u0000\u0000\u0210\u0211\u0007\u0002\u0000\u0000"+
		"\u0211\u0232\u0003H$\u0014\u0212\u0232\u0007\u0003\u0000\u0000\u0213\u0232"+
		"\u0005(\u0000\u0000\u0214\u0232\u0005B\u0000\u0000\u0215\u0232\u0005C"+
		"\u0000\u0000\u0216\u0232\u0005D\u0000\u0000\u0217\u0232\u0003*\u0015\u0000"+
		"\u0218\u0219\u0005E\u0000\u0000\u0219\u0232\u0005!\u0000\u0000\u021a\u021b"+
		"\u0005E\u0000\u0000\u021b\u0232\u0005\"\u0000\u0000\u021c\u021d\u0005"+
		"E\u0000\u0000\u021d\u021e\u0005\u0001\u0000\u0000\u021e\u021f\u0003H$"+
		"\u0000\u021f\u0220\u0005\u0002\u0000\u0000\u0220\u0232\u0001\u0000\u0000"+
		"\u0000\u0221\u0232\u0005E\u0000\u0000\u0222\u0223\u0005#\u0000\u0000\u0223"+
		"\u0224\u0005\u0005\u0000\u0000\u0224\u0225\u0003H$\u0000\u0225\u0226\u0005"+
		"\u0006\u0000\u0000\u0226\u0232\u0001\u0000\u0000\u0000\u0227\u0228\u0005"+
		"$\u0000\u0000\u0228\u0229\u0005\u0005\u0000\u0000\u0229\u022a\u0003H$"+
		"\u0000\u022a\u022b\u0005\u0006\u0000\u0000\u022b\u0232\u0001\u0000\u0000"+
		"\u0000\u022c\u022d\u0005%\u0000\u0000\u022d\u022e\u0005\u0005\u0000\u0000"+
		"\u022e\u022f\u0003H$\u0000\u022f\u0230\u0005\u0006\u0000\u0000\u0230\u0232"+
		"\u0001\u0000\u0000\u0000\u0231\u020b\u0001\u0000\u0000\u0000\u0231\u0210"+
		"\u0001\u0000\u0000\u0000\u0231\u0212\u0001\u0000\u0000\u0000\u0231\u0213"+
		"\u0001\u0000\u0000\u0000\u0231\u0214\u0001\u0000\u0000\u0000\u0231\u0215"+
		"\u0001\u0000\u0000\u0000\u0231\u0216\u0001\u0000\u0000\u0000\u0231\u0217"+
		"\u0001\u0000\u0000\u0000\u0231\u0218\u0001\u0000\u0000\u0000\u0231\u021a"+
		"\u0001\u0000\u0000\u0000\u0231\u021c\u0001\u0000\u0000\u0000\u0231\u0221"+
		"\u0001\u0000\u0000\u0000\u0231\u0222\u0001\u0000\u0000\u0000\u0231\u0227"+
		"\u0001\u0000\u0000\u0000\u0231\u022c\u0001\u0000\u0000\u0000\u0232\u0247"+
		"\u0001\u0000\u0000\u0000\u0233\u0234\n\u0013\u0000\u0000\u0234\u0235\u0007"+
		"\u0004\u0000\u0000\u0235\u0246\u0003H$\u0014\u0236\u0237\n\u0012\u0000"+
		"\u0000\u0237\u0238\u0007\u0005\u0000\u0000\u0238\u0246\u0003H$\u0013\u0239"+
		"\u023a\n\u0011\u0000\u0000\u023a\u023b\u0007\u0006\u0000\u0000\u023b\u0246"+
		"\u0003H$\u0012\u023c\u023d\n\u0010\u0000\u0000\u023d\u023e\u0007\u0007"+
		"\u0000\u0000\u023e\u0246\u0003H$\u0011\u023f\u0240\n\u000f\u0000\u0000"+
		"\u0240\u0241\u0007\b\u0000\u0000\u0241\u0246\u0003H$\u0010\u0242\u0243"+
		"\n\u000e\u0000\u0000\u0243\u0244\u0007\t\u0000\u0000\u0244\u0246\u0003"+
		"H$\u000f\u0245\u0233\u0001\u0000\u0000\u0000\u0245\u0236\u0001\u0000\u0000"+
		"\u0000\u0245\u0239\u0001\u0000\u0000\u0000\u0245\u023c\u0001\u0000\u0000"+
		"\u0000\u0245\u023f\u0001\u0000\u0000\u0000\u0245\u0242\u0001\u0000\u0000"+
		"\u0000\u0246\u0249\u0001\u0000\u0000\u0000\u0247\u0245\u0001\u0000\u0000"+
		"\u0000\u0247\u0248\u0001\u0000\u0000\u0000\u0248I\u0001\u0000\u0000\u0000"+
		"\u0249\u0247\u0001\u0000\u0000\u0000\u024a\u024b\u0007\n\u0000\u0000\u024b"+
		"K\u0001\u0000\u0000\u0000?RW[_cgkosw{\u007f\u0083\u0087\u008b\u008f\u0093"+
		"\u0097\u009b\u009d\u00b0\u00b8\u00c2\u00cc\u00d0\u00dd\u00e8\u00f5\u00fc"+
		"\u0121\u0130\u0134\u013a\u013f\u0147\u0149\u0152\u0156\u0159\u015c\u0165"+
		"\u0171\u0176\u017d\u0182\u0186\u018a\u018f\u0195\u0198\u019d\u01ab\u01b0"+
		"\u01ca\u01d4\u01d6\u01de\u01e4\u01f1\u0209\u0231\u0245\u0247";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}