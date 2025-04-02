# Interpret aritmetických výrazů pomocí ANTLR

## Zadání

Pomocí ANTLR implementujte interpret aritmetických výrazů. Tyto výrazy obsahují operátory +, -, \*, / (s běžnými prioritami a levou asociativitou) a závorky. Pro zjednodušení úlohy uvažujte pouze binární operátory. V našem jazyce nejsou žádné unární operátory.

### Vstupní specifikace

Ve vstupu jsou výrazy, které jsou formátovány. Každý výraz končí středníkem. Čísla mohou být zapsána podobně jako konstanty v jazyce C. Může se jednat buď o:

-   decimální číslo
-   osmičkové číslo (začínající nulou)
-   hexadecimální číslo (začínající znaky 0x)

### Výstupní specifikace

Pro každý výraz vypište jeden řádek obsahující výsledek – vypočtenou hodnotu výrazu. Pokud se ve vstupu vyskytne jakákoliv chyba, můžete výpočet zastavit.

### Příklad

**Vstup:**

```
012-10; 2 * (0xff+5);
0x23e5-0x201;
```

**Výstup:**

```
0
520
8676
```

## O programu

Tento program je implementací interpretu aritmetických výrazů pomocí ANTLR (ANother Tool for Language Recognition). Program dokáže:

-   Parseovat a vyhodnocovat aritmetické výrazy s operátory +, -, \*, /
-   Respektovat závorky a priority operátorů
-   Podporovat čísla v různých formátech (desítkové, osmičkové, šestnáctkové)
-   Zpracovávat více výrazů oddělených středníkem

Implementace využívá vzor "Listener" z ANTLR, který umožňuje procházet syntaktický strom výrazu a provádět akce při opuštění jednotlivých uzlů. Pro vyhodnocení výrazů je použit zásobník, na který se ukládají mezivýsledky.

## Použité technologie

-   **Go** (verze 1.22 nebo novější)
-   **ANTLR4** (verze 4.9.2)
-   **ANTLR4 Go runtime**

## Struktura projektu

-   `PLC_Lab7_expr_2.g4` - ANTLR gramatika pro aritmetické výrazy
-   `main.go` - Hlavní implementace interpretu
-   `parser2/` - Generované soubory parseru
-   `go.mod`, `go.sum` - Definice Go modulu

## Jak spustit program

### Instalace ANTLR

1. Stáhněte ANTLR:

```
Invoke-WebRequest -Uri https://www.antlr.org/download/antlr-4.9.2-complete.jar -OutFile antlr-4.9.2-complete.jar
```

2. Vytvořte dávkový soubor `antlr4.bat` s obsahem:

```
@echo off
java -jar %~dp0antlr-4.9.2-complete.jar %*
```

### Generování parseru

```
java -jar antlr-4.9.2-complete.jar -Dlanguage=Go -package parser2 -o parser2 PLC_Lab7_expr_2.g4
```

### Instalace závislostí

```
go mod tidy
```

### Spuštění programu

Pro interaktivní režim:

```
go run main.go
```

Pro zpracování souboru:

```
go run main.go input.txt
```

## Příklady použití

V interaktivním režimu lze zadat výrazy přímo z klávesnice:

```
Interpret aritmetických výrazů
================================
Zadejte výrazy končící středníkem (Ctrl+D pro ukončení):
42;
42
2+2*3;
8
(2+2)*3;
12
```

Nebo lze spustit program se souborem obsahujícím výrazy:

```
go run main.go input.txt
```

kde `input.txt` obsahuje:

```
012-10; 2 * (0xff+5);
0x23e5-0x201;
```
