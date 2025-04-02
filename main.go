package main

import (
	"bufio"   // Balíček pro bufferované čtení vstupu
	"fmt"     // Balíček pro formátovaný výstup
	"os"      // Balíček pro práci s operačním systémem (soubory, argumenty, atd.)
	"strconv" // Balíček pro konverzi mezi řetězci a číselnými typy
	"strings" // Balíček pro manipulaci s řetězci

	"7-nacviku/parser2" // Náš vygenerovaný ANTLR parser pro aritmetické výrazy

	"github.com/antlr/antlr4/runtime/Go/antlr" // ANTLR runtime knihovna pro Go
)

// ErrorListener je struktura pro zachycení chyb během parsování pomocí ANTLR
// Rozšiřuje standardní antlr.DefaultErrorListener pro zpracování syntaktických chyb
// Atributy:
// - DefaultErrorListener: pointer na výchozí ANTLR error listener
// - errors: seznam řetězců obsahující chybové zprávy
type ErrorListener struct {
	*antlr.DefaultErrorListener          // Vkládáme (embedding) výchozí error listener z ANTLR
	errors                      []string // Seznam pro ukládání chybových zpráv
}

// NewErrorListener vytvoří a inicializuje novou instanci ErrorListener
// Parametry: žádné
// Návratová hodnota:
// - *ErrorListener: pointer na novou instanci ErrorListener s inicializovanými atributy
func NewErrorListener() *ErrorListener {
	return &ErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(), // Vytvoří výchozí ANTLR error listener
		errors:               []string{},                      // Inicializuje prázdný seznam pro chybové zprávy
	}
}

// SyntaxError je metoda, která je automaticky volána ANTLRem při nalezení syntaktické chyby
// Je součástí rozhraní antlr.ErrorListener, které musíme implementovat
// Parametry:
// - recognizer: rozhraní antlr.Recognizer - objekt provádějící parsování (lexer nebo parser)
// - offendingSymbol: interface{} - token, který způsobil chybu
// - line: int - číslo řádku, kde byla chyba nalezena
// - column: int - pozice ve sloupci, kde byla chyba nalezena
// - msg: string - popis chyby
// - e: antlr.RecognitionException - výjimka spojená s chybou
// Návratová hodnota: žádná - metoda pouze přidá formátovanou chybovou zprávu do seznamu errors
func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line int, column int, msg string, e antlr.RecognitionException) {
	// Formátuje chybovou zprávu a přidá ji do seznamu
	l.errors = append(l.errors, fmt.Sprintf("řádek %d:%d %s", line, column, msg))
}

// HasErrors kontroluje, zda byly během parsování nalezeny nějaké chyby
// Parametry: žádné
// Návratová hodnota:
// - bool: true pokud byly nalezeny chyby (seznam errors není prázdný), jinak false
func (l *ErrorListener) HasErrors() bool {
	return len(l.errors) > 0 // Vrací true, pokud seznam obsahuje alespoň jednu chybovou zprávu
}

// GetErrors vrací seznam všech nalezených chybových zpráv
// Parametry: žádné
// Návratová hodnota:
// - []string: seznam chybových zpráv
func (l *ErrorListener) GetErrors() []string {
	return l.errors // Vrací celý seznam chybových zpráv
}

// ExpressionCalculator implementuje parser2.PLC_Lab7_expr_2Listener rozhraní pro vyhodnocení výrazů
// Používá vzor Listener, který je v ANTLR založen na volání callback metod při procházení syntaktického stromu
// Atributy:
// - BasePLC_Lab7_expr_2Listener: vygenerovaná základní implementace listeneru
// - stack: zásobník (slice) pro ukládání mezivýsledků během vyhodnocování výrazu
type ExpressionCalculator struct {
	parser2.BasePLC_Lab7_expr_2Listener         // Vložení (embedding) vygenerovaného listeneru
	stack                               []int64 // Zásobník pro ukládání hodnot během vyhodnocování výrazu
}

// push přidá novou hodnotu na vrchol zásobníku
// Tato pomocná metoda se používá v Exit* metodách pro ukládání výsledků výrazů
// Parametry:
// - i: int64 - číselná hodnota, která bude přidána na zásobník
// Návratová hodnota: žádná - metoda modifikuje zásobník v instanci ExpressionCalculator
func (c *ExpressionCalculator) push(i int64) {
	// Přidáme hodnotu na konec slice, což reprezentuje vrchol zásobníku
	c.stack = append(c.stack, i)
}

// pop odebere a vrátí hodnotu z vrcholu zásobníku
// Používá se v Exit* metodách pro získání hodnot operandů při vyhodnocování operací
// Parametry: žádné
// Návratová hodnota:
// - int64: hodnota z vrcholu zásobníku
// Panická situace: pokud je zásobník prázdný, metoda vyvolá paniku
func (c *ExpressionCalculator) pop() int64 {
	// Kontrola, zda zásobník není prázdný
	if len(c.stack) < 1 {
		panic("Zásobník je prázdný - chyba při vyhodnocování výrazu")
	}

	// Získáme hodnotu z vrcholu zásobníku (poslední prvek slice)
	result := c.stack[len(c.stack)-1]

	// Odebereme hodnotu ze zásobníku (zmenšíme slice o poslední prvek)
	c.stack = c.stack[:len(c.stack)-1]

	return result // Vrátíme odebranou hodnotu
}

// ExitInt je metoda volaná ANTLRem při opuštění uzlu s dekadickým číslem
// Tato metoda překrývá stejnojmennou metodu z BasePLC_Lab7_expr_2Listener
// Je automaticky volána, když walker opouští uzel typu IntContext v syntaktickém stromu
// Parametry:
// - ctx: *parser2.IntContext - kontext uzlu obsahující informace o zpracovávaném dekadickém čísle
// Návratová hodnota: žádná - metoda pouze ukládá převedenou hodnotu na zásobník
func (c *ExpressionCalculator) ExitInt(ctx *parser2.IntContext) {
	// Získáme text uzlu (např. "123") pomocí GetText() metody
	text := ctx.GetText()

	// Převedeme text na číslo v desítkové soustavě (základ 10)
	// Ignorujeme chybu (druhý návratový parametr), protože gramatika zajišťuje, že text je validní číslo
	value, _ := strconv.ParseInt(text, 10, 64)

	// Uložíme hodnotu na zásobník
	c.push(value)
}

// ExitHex je metoda volaná ANTLRem při opuštění uzlu s hexadecimálním číslem
// Je automaticky volána, když walker opouští uzel typu HexContext v syntaktickém stromu
// Parametry:
// - ctx: *parser2.HexContext - kontext uzlu obsahující informace o zpracovávaném hexadecimálním čísle
// Návratová hodnota: žádná - metoda pouze ukládá převedenou hodnotu na zásobník
func (c *ExpressionCalculator) ExitHex(ctx *parser2.HexContext) {
	// Získáme text uzlu (např. "0x1A") pomocí GetText()
	hexText := ctx.GetText()

	// Odstraníme "0x" prefix, protože strconv.ParseInt očekává pouze hexadecimální číslice
	text := hexText[2:]

	// Převedeme text na číslo v šestnáctkové soustavě (základ 16)
	// Ignorujeme chybu (druhý návratový parametr)
	value, _ := strconv.ParseInt(text, 16, 64)

	// Uložíme hodnotu na zásobník
	c.push(value)
}

// ExitOct je metoda volaná ANTLRem při opuštění uzlu s osmičkovým číslem
// Je automaticky volána, když walker opouští uzel typu OctContext v syntaktickém stromu
// Parametry:
// - ctx: *parser2.OctContext - kontext uzlu obsahující informace o zpracovávaném osmičkovém čísle
// Návratová hodnota: žádná - metoda pouze ukládá převedenou hodnotu na zásobník
func (c *ExpressionCalculator) ExitOct(ctx *parser2.OctContext) {
	// Získáme text uzlu (např. "012") pomocí GetText()
	octText := ctx.GetText()

	// Převedeme text na číslo v osmičkové soustavě (základ 8)
	// strconv.ParseInt automaticky rozpozná, že jde o osmičkové číslo, pokud má prefix "0"
	// Ignorujeme chybu (druhý návratový parametr)
	value, _ := strconv.ParseInt(octText, 8, 64)

	// Uložíme hodnotu na zásobník
	c.push(value)
}

// ExitAddSub je metoda volaná ANTLRem při opuštění uzlu se sčítáním nebo odčítáním
// Je automaticky volána, když walker opouští uzel typu AddSubContext v syntaktickém stromu
// Parametry:
// - ctx: *parser2.AddSubContext - kontext uzlu obsahující informace o operaci sčítání/odčítání
// Návratová hodnota: žádná - metoda provede operaci a výsledek uloží na zásobník
func (c *ExpressionCalculator) ExitAddSub(ctx *parser2.AddSubContext) {
	// Zásobník obsahuje hodnoty obou operandů, které již byly vyhodnoceny
	// Protože jsme nejprve navštívili pravou stranu, pak levou stranu, je pořadí na zásobníku:
	// vrchol zásobníku -> pravý operand
	// druhý prvek -> levý operand

	// Odebíráme hodnoty ze zásobníku v opačném pořadí, než byly vloženy (LIFO)
	right := c.pop() // Nejprve získáme pravý operand
	left := c.pop()  // Poté levý operand

	// Zjistíme, jaký operátor byl použit (+ nebo -)
	// ctx.GetChild(1) vrací druhé dítě uzlu (indexováno od 0):
	// - GetChild(0) je levý výraz
	// - GetChild(1) je operátor
	// - GetChild(2) je pravý výraz
	// Přetypujeme na TerminalNode, abychom mohli získat text tokenu
	operator := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	// Provedeme odpovídající operaci podle operátoru
	var result int64
	switch operator {
	case "+":
		result = left + right // Sčítání
	case "-":
		result = left - right // Odčítání
	}

	// Výsledek operace uložíme zpět na zásobník
	c.push(result)
}

// ExitMulDiv je metoda volaná ANTLRem při opuštění uzlu s násobením nebo dělením
// Je automaticky volána, když walker opouští uzel typu MulDivContext v syntaktickém stromu
// Parametry:
// - ctx: *parser2.MulDivContext - kontext uzlu obsahující informace o operaci násobení/dělení
// Návratová hodnota: žádná - metoda provede operaci a výsledek uloží na zásobník
// Panická situace: pokud je dělitel nula, metoda vyvolá paniku
func (c *ExpressionCalculator) ExitMulDiv(ctx *parser2.MulDivContext) {
	// Stejně jako u ExitAddSub, odebíráme hodnoty ze zásobníku
	right := c.pop() // Pravý operand (dělitel nebo druhý činitel)
	left := c.pop()  // Levý operand (dělenec nebo první činitel)

	// Zjistíme operátor (* nebo /)
	operator := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	// Provedeme odpovídající operaci
	var result int64
	switch operator {
	case "*":
		result = left * right // Násobení
	case "/":
		// Ošetření dělení nulou - vyvoláme paniku s popisným textem
		if right == 0 {
			panic("Dělení nulou - nelze dělit číslem 0")
		}
		result = left / right // Celočíselné dělení
	}

	// Výsledek operace uložíme zpět na zásobník
	c.push(result)
}

// ExitParens je metoda volaná ANTLRem při opuštění uzlu se závorkami
// Je automaticky volána, když walker opouští uzel typu ParensContext v syntaktickém stromu
// Parametry:
// - ctx: *parser2.ParensContext - kontext uzlu obsahující informace o výrazu v závorkách
// Návratová hodnota: žádná - metoda nemusí nic dělat, protože hodnota výrazu v závorkách je již na zásobníku
func (c *ExpressionCalculator) ExitParens(ctx *parser2.ParensContext) {
	// Výraz uvnitř závorek již byl vyhodnocen a jeho hodnota je na zásobníku
	// Proto zde nemusíme nic dělat - tato metoda je zde uvedena pro úplnost a dokumentaci
}

// Calculate je hlavní funkce pro vyhodnocení aritmetického výrazu
// Vytvoří a inicializuje všechny potřebné ANTLR komponenty, spustí parsing a vrátí výsledek
// Parametry:
// - input: string - textový řetězec obsahující aritmetický výraz k vyhodnocení
// Návratové hodnoty:
// - int64: číselný výsledek vyhodnoceného výrazu
// - error: chyba, pokud k nějaké došlo během parsování nebo vyhodnocování
func Calculate(input string) (int64, error) {
	// 1. VYTVOŘENÍ VSTUPNÍHO PROUDU (INPUT STREAM)
	// Vytvoříme vstupní proud znaků (CharStream) z textového řetězce
	// Tento proud poskytuje znaky lexeru
	inputStream := antlr.NewInputStream(input)

	// 2. VYTVOŘENÍ LEXERU
	// Lexer převádí proud znaků na proud tokenů (např. čísla, operátory, závorky)
	// Používáme vygenerovaný lexer z naší gramatiky
	lexer := parser2.NewPLC_Lab7_expr_2Lexer(inputStream)

	// 3. PŘIDÁNÍ ERROR LISTENERU K LEXERU
	// Vytvoříme vlastní error listener pro zachycení lexikálních chyb
	errorListener := NewErrorListener()
	lexer.RemoveErrorListeners()          // Odstraníme výchozí error listenery
	lexer.AddErrorListener(errorListener) // Přidáme náš vlastní error listener

	// 4. VYTVOŘENÍ PROUDU TOKENŮ
	// CommonTokenStream poskytuje tokeny z lexeru parseru
	// Používáme výchozí kanál (DEFAULT_TOKEN_CHANNEL)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// 5. VYTVOŘENÍ PARSERU
	// Parser používá proud tokenů k sestavení syntaktického stromu
	// Používáme vygenerovaný parser z naší gramatiky
	parser := parser2.NewPLC_Lab7_expr_2Parser(tokens)

	// 6. PŘIDÁNÍ ERROR LISTENERU K PARSERU
	// Přidáme stejný error listener i k parseru pro zachycení syntaktických chyb
	parser.RemoveErrorListeners()
	parser.AddErrorListener(errorListener)

	// 7. SPUŠTĚNÍ PARSOVÁNÍ
	// Začínáme na pravidle 'expr' definovaném v gramatice
	// Parser vytvoří syntaktický strom reprezentující strukturu výrazu
	tree := parser.Expr()

	// 8. KONTROLA CHYB PARSOVÁNÍ
	// Zkontrolujeme, zda listener nezachytil žádné chyby
	if errorListener.HasErrors() {
		// Pokud byly nalezeny chyby, vrátíme chybový objekt s popisem
		return 0, fmt.Errorf("chyba při parsování: %s", strings.Join(errorListener.GetErrors(), "; "))
	}

	// 9. VYTVOŘENÍ KALKULÁTORU (LISTENERU)
	// Vytvoříme náš listener pro vyhodnocení výrazu
	calculator := &ExpressionCalculator{}

	// 10. ZACHYCENÍ RUNTIME CHYB
	// Nastavíme deferred funkci pro zachycení případných panik (např. dělení nulou)
	defer func() {
		// Tato funkce se spustí při opuštění funkce Calculate (normálně nebo při panice)
		if r := recover(); r != nil {
			// Pokud došlo k panice, vypíšeme chybovou zprávu a ukončíme program
			fmt.Println("Chyba při vyhodnocování:", r)
			os.Exit(1)
		}
	}()

	// 11. PROCHÁZENÍ SYNTAKTICKÉHO STROMU
	// Použijeme výchozí walker z ANTLR runtime, který prochází stromem a volá metody listeneru
	// Walker postupuje zdola nahoru - nejprve volá Enter* metody, pak rekurzivně zpracuje děti,
	// pak volá Exit* metody. Pro vyhodnocení výrazů používáme pouze Exit* metody.
	antlr.ParseTreeWalkerDefault.Walk(calculator, tree)

	// 12. KONTROLA VÝSLEDKU
	// Po zpracování celého výrazu by měl být na zásobníku právě jeden prvek - výsledek
	if len(calculator.stack) != 1 {
		return 0, fmt.Errorf("neplatný výraz - na zásobníku zůstalo %d hodnot místo očekávané jedné", len(calculator.stack))
	}

	// 13. VRÁCENÍ VÝSLEDKU
	// Vrátíme hodnotu z vrcholu zásobníku jako výsledek výrazu
	return calculator.stack[0], nil
}

// main je vstupní bod programu
// Zpracovává argumenty příkazové řádky a řídí zpracování vstupu
// Parametry: žádné (argumenty jsou získány z os.Args)
// Návratová hodnota: žádná (program končí s kódem 0 při úspěchu, 1 při chybě)
func main() {
	// Získáme argumenty příkazové řádky bez názvu programu
	// os.Args[0] je název programu, os.Args[1:] jsou skutečné argumenty
	args := os.Args[1:]

	// Rozhodneme, zda číst vstup ze souboru nebo ze standardního vstupu
	if len(args) > 0 {
		// Pokud byl zadán alespoň jeden argument, považujeme ho za název vstupního souboru

		// Otevřeme soubor pro čtení
		file, err := os.Open(args[0])
		if err != nil {
			// Pokud dojde k chybě při otevírání souboru, vypíšeme ji a ukončíme program
			fmt.Fprintf(os.Stderr, "Chyba při otevírání souboru %s: %v\n", args[0], err)
			os.Exit(1)
		}
		// Zajistíme zavření souboru po ukončení funkce main
		defer file.Close()

		// Zpracujeme vstup ze souboru pomocí pomocné funkce
		processInput(bufio.NewScanner(file))
	} else {
		// Interaktivní režim - čtení ze standardního vstupu (klávesnice)
		fmt.Println("Interpret aritmetických výrazů")
		fmt.Println("================================")
		fmt.Println("Zadejte výrazy končící středníkem (Ctrl+D pro ukončení):")

		// Zpracujeme vstup ze standardního vstupu pomocí stejné pomocné funkce
		processInput(bufio.NewScanner(os.Stdin))
	}
}

// processInput zpracovává vstup řádek po řádku a vyhodnocuje kompletní výrazy
// Parametry:
// - scanner: *bufio.Scanner - scanner pro čtení vstupu (buď ze souboru nebo ze std. vstupu)
// Návratová hodnota: žádná - výsledky jsou vypisovány na standardní výstup
// Ukončení: funkce ukončí program s kódem 1, pokud dojde k chybě
func processInput(scanner *bufio.Scanner) {
	// Vytvoříme buffer pro postupné ukládání vstupních řádků
	// Používáme strings.Builder místo string pro efektivnější spojování řetězců
	buffer := strings.Builder{}

	// Čteme vstup řádek po řádku pomocí scanneru
	for scanner.Scan() {
		// Získáme aktuální řádek ze scanneru
		line := scanner.Text()

		// Přidáme řádek do bufferu
		buffer.WriteString(line)

		// Kontrolujeme, zda řádek obsahuje středník
		// Středník slouží jako oddělovač výrazů - podle zadání každý výraz končí středníkem
		if strings.Contains(line, ";") {
			// Pokud řádek obsahuje středník, máme alespoň jeden kompletní výraz

			// Získáme veškerý dosud načtený text z bufferu
			input := buffer.String()

			// Vyčistíme buffer pro další vstup
			buffer.Reset()

			// Rozdělíme vstup podle středníků na jednotlivé výrazy
			// Např. "1+2; 3*4;" se rozdělí na ["1+2", " 3*4", ""]
			expressions := strings.Split(input, ";")

			// Zpracujeme každý výraz
			for _, expr := range expressions {
				// Odstraníme bílé znaky na začátku a konci výrazu
				expr = strings.TrimSpace(expr)

				// Přeskočíme prázdné výrazy (např. poslední prvek po dělení)
				if expr == "" {
					continue
				}

				// Přidáme zpět středník, který byl odstraněn při dělení
				// ANTLR gramatika očekává středník na konci výrazu
				expr += ";"

				// Vyhodnotíme výraz pomocí naší hlavní funkce Calculate
				result, err := Calculate(expr)

				// Zkontrolujeme, zda nedošlo k chybě
				if err != nil {
					// Při chybě vypíšeme chybovou zprávu a ukončíme program
					fmt.Fprintf(os.Stderr, "Chyba při vyhodnocování výrazu '%s': %v\n", expr, err)
					os.Exit(1)
				}

				// Vypíšeme výsledek výrazu na standardní výstup
				// Podle zadání má být na každém řádku jeden výsledek
				fmt.Println(result)
			}
		}
	}

	// Zkontrolujeme, zda nedošlo k chybě při čtení vstupu
	if err := scanner.Err(); err != nil {
		// Při chybě vypíšeme chybovou zprávu a ukončíme program
		fmt.Fprintf(os.Stderr, "Chyba při čtení vstupu: %v\n", err)
		os.Exit(1)
	}
}
