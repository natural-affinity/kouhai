# Kouhai (後輩)

Periodic command execution at specific intervals (e.g. gnu watch).

## Example Usage

```bash
# help and usage options
kouhai -h
```

```bash
# invoke make test every 2s and stop on first failure
kouhai -s -i 2s "make test"
```

```Makefile
# watch task to invoke make test via kouhai
.PHONY: watch
watch: 
    @kouhai -s --interval 2s "make test"

.PHONY: test
test:
    @go test ./...
```
