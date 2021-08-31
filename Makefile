base:
	zip module.zip *.tf README.md doc/*.md

base_tests:
	zip module_and_tests.zip *.tf README.md doc/*.md test/*_test.go test/*.md

all: base base_tests
