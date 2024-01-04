INPUT=$(wildcard $(PWD)/schemas/v2.3/*.xsd)

%.go: $(PWD)/schemas/v2.3/%.xsd
	go run github.com/xuri/xgen/cmd/xgen@latest -i $< -o $(basename $@) -l Go -p gofx
	go fmt $@

generate: $(patsubst %.xsd,%.go,$(notdir $(INPUT)))
	@echo "Generated!"

clean:
	-rm $(patsubst %.xsd,%.go,$(notdir $(INPUT)))
