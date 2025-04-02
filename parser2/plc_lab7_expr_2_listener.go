// Code generated from PLC_Lab7_expr_2.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser2 // PLC_Lab7_expr_2
import "github.com/antlr/antlr4/runtime/Go/antlr"

// PLC_Lab7_expr_2Listener is a complete listener for a parse tree produced by PLC_Lab7_expr_2Parser.
type PLC_Lab7_expr_2Listener interface {
	antlr.ParseTreeListener

	// EnterProg is called when entering the prog production.
	EnterProg(c *ProgContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterOct is called when entering the Oct production.
	EnterOct(c *OctContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterParens is called when entering the Parens production.
	EnterParens(c *ParensContext)

	// EnterHex is called when entering the Hex production.
	EnterHex(c *HexContext)

	// EnterInt is called when entering the Int production.
	EnterInt(c *IntContext)

	// ExitProg is called when exiting the prog production.
	ExitProg(c *ProgContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitOct is called when exiting the Oct production.
	ExitOct(c *OctContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitParens is called when exiting the Parens production.
	ExitParens(c *ParensContext)

	// ExitHex is called when exiting the Hex production.
	ExitHex(c *HexContext)

	// ExitInt is called when exiting the Int production.
	ExitInt(c *IntContext)
}
