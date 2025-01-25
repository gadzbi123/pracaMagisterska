set -euo pipefail
IFS=$'\n\t'

DIR=/run/media/gadzbi/GryIFilmy/baza_mgr_reduced
cd /home/gadzbi/proj/pracaMagisterska/projectMgr/algorytmy
go build -ldflags "-s -w" -o _build/gsearch cmd/gsearch/main.go

hyperfine --warmup 1 --runs 10 "_build/gsearch wan $DIR" "_build/rg -lcUz wan $DIR/*" --export-csv binary_tests/wan.csv
hyperfine --runs 10 "_build/gsearch main $DIR" "_build/rg -lcUz main $DIR/*" --export-csv binary_tests/main.csv
hyperfine --runs 10 "_build/gsearch window $DIR" "_build/rg -lcUz window $DIR/*" --export-csv binary_tests/window.csv
hyperfine --runs 10 "_build/gsearch analysis $DIR" "_build/rg -lcUz analysis $DIR/*" --export-csv binary_tests/analysis.csv
hyperfine --runs 10 "_build/gsearch 'book desc' $DIR" "_build/rg -lcUz 'book desc' $DIR/*" --export-csv binary_tests/book_desc.csv
hyperfine --runs 10 "_build/gsearch informatyka $DIR" "_build/rg -lcUz informatyka $DIR/*" --export-csv binary_tests/informatyka.csv
hyperfine --runs 10 "_build/gsearch wInDoW $DIR" "_build/rg -lcUzi wInDoW $DIR/*" --export-csv binary_tests/window-case.csv