// Code generated from PLC_Lab7_expr_2.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser2 // PLC_Lab7_expr_2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// BasePLC_Lab7_expr_2Listener is a complete listener for a parse tree produced by PLC_Lab7_expr_2Parser.
type BasePLC_Lab7_expr_2Listener struct{}

var _ PLC_Lab7_expr_2Listener = &BasePLC_Lab7_expr_2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePLC_Lab7_expr_2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePLC_Lab7_expr_2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitProg(ctx *ProgContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterOct is called when production Oct is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterOct(ctx *OctContext) {}

// ExitOct is called when production Oct is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitOct(ctx *OctContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitParens(ctx *ParensContext) {}

// EnterHex is called when production Hex is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterHex(ctx *HexContext) {}

// ExitHex is called when production Hex is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitHex(ctx *HexContext) {}

// EnterInt is called when production Int is entered.
func (s *BasePLC_Lab7_expr_2Listener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BasePLC_Lab7_expr_2Listener) ExitInt(ctx *IntContext) {}
