[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE.md)

`trimv`
======

### Usage:
```
trimv [flags]
```
```
trimv [command]
```

### Available Commands:
  - **`all`:** Trim all videos matching pattern
  - **`help`:** Help about any command

### Flags:
  - **`-h`, `--help`:** help for `trimv`
  - **`-i`, `--input string`:** Input file path
  - **`-l`, `--intro-duration` _`float32`_:** Specifies the intro duration _(default: 15)_
  - **`-o`, `--output string`:** Output file path
  - **`-r`, `--outro-duration` _`float32`_:** Specifies the outro duration _(default: 15)_

Command `all`
---------------

### Usage:
```
trimv all <input-pattern> [flags]
```

### Flags:
  - **`-h`, `--help`:** help for `all`
  - **`-l`, `--intro-duration` _`float32`_:** Specifies the intro duration _(default: 15)_
  - **`-o`, `--output string`:** Output folder
  - **`-r`, `--outro-duration` _`float32`_:** Specifies the outro duration _(default: 15)_