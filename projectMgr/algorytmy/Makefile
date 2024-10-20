bench :
	go test -bench=. ./regular -benchtime=1x -benchmem

benchCount :
	go test -test.bench=. -benchmem ./regular -benchtime=1x -count=10

benchOnce  :
	go test -bench=MainWord ./regular -benchtime=1x -benchmem
  
benchstat :
	benchstat test_results/10_run_main.txt test_results/20_run_main.txt 