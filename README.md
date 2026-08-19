# ccat v2.0

Moderní, univerzální a přenosný nástroj typu `cat` s podporou barevného zvýraznění syntaxe, různých výstupních formátů a optimalizací pro desktop i embedded zařízení.

## Vlastnosti
- Syntax highlighting pro mnoho jazyků (Go, Python, Bash, JSON, YAML, Rust, C, C++, Markdown, TOML, XML, JavaScript, TypeScript, Lua, PHP, SQL)
- Automatická detekce jazyka podle přípony, shebang nebo obsahu
- Výstup s ANSI barvami, HTML nebo prostý text
- Podpora motivů (vlastní i předdefinované)
- Čísla řádků
- Pager (less -R) pro dlouhé výstupy
- Bezpečnostní kontroly (binární soubory, velikost, dlouhé řádky)
- Přenositelný – běží na Linuxu, Windows, macOS, ARM (Raspberry Pi, NAS), Android (Termux)

## Instalace
```

go build -o ccat ./cmd/ccat

```

## Použití
```

ccat [přepínače] [soubor...]

```

Přepínače:
- `-n` – zobrazit čísla řádků
- `--color auto|never|always` – ovládání barev
- `--language <jazyk>` – vynutit jazyk
- `--theme <jméno>` – vybrat motiv
- `--html` – HTML výstup
- `--version` – verze

## Konfigurace
Soubor `~/.config/ccat/config.toml`:
```toml
theme = "dracula"
color = true
line_numbers = false
pager = true
max_size = 10485760
max_line_length = 4096
```

Motivy

Umístěte vlastní motivy do ~/.config/ccat/themes/ ve formátu TOML.

Licence

MIT
