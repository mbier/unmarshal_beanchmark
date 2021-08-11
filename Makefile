CUR_DIR = $(CURDIR)

run-benchmark:
	go test -bench=. -run=. -benchmem -benchtime=1s

proto-gen:
	rm -rf gen/*
	docker run --rm -v "$(CUR_DIR):/work" uber/prototool prototool generate

avro-gen:
	go get github.com/actgardner/gogen-avro/v9/cmd/gogen-avro
	mkdir -p ./avro
	$(GOPATH)/bin/gogen-avro -containers ./avro model.avsc