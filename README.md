# Kouhai (後輩)

Periodic command execution at specific intervals (i.e. `watch`).

## Usage

```bash
# view help and usage options
kouhai -h
```

## Examples
Use it on the `CLI`

```bash
# invoke every 2s and stop on first failure
kouhai -s -i 2s "make test"
```

Use it within a `makefile`

```makefile
# define watch task to invoke make test via kouhai
.PHONY: watch
watch: 
    @kouhai -s --interval 2s "make test"

.PHONY: test
test:
    @go test ./...
```

## Links
* [watch](https://en.wikipedia.org/wiki/Watch_(Unix))

## License
Released under the MIT License. See [LICENSE.md](./LICENSE.md) for details.
