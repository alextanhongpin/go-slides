# go-slides

A collection of slides on different go topics.


## Pre-commit hook


You can add this pre-commit hook to ensure the slides are build upon commit.
```bash
#!/bin/sh

make gen
git add docs/*
```
