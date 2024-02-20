# paramM

Finds param in url

## Installtion

```bash
go install github.com/SpeedyQweku/paramM@latest
```

## Usage

```bash
paramM -p < xss.txt > -l < url.txt >| qsreplace FUZZ | tee result.txt
```

## NOTE

Install qsreplace from tomnomnom

```bash
go install github.com/tomnomnom/qsreplace@latest
```
