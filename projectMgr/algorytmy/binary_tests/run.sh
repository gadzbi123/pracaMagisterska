set -euo pipefail
IFS=$'\n\t'

DIR=/run/media/gadzbi/GryIFilmy/baza_mgr
cd /home/gadzbi/proj/pracaMagisterska/projectMgr/algorytmy
go build -ldflags "-s -w" -o _build/gsearch cmd/gsearch/main.go

hyperfine --warmup 1 --runs 1 '_build/gsearch main $DIR' '_build/rg -lcUz main $DIR/*' '_build/zgrep ' --export-csv binary_tests/test-res.csv