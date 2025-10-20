#expected run from repo root

#!/bin/bash

repo-analyzer --remote-url https://github.com/gokh4nozturk/repo-analyzer --output-format json
mkdir -p ./data
mv report.json ./data/report.json