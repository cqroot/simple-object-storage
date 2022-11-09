.PHONY: test
test:
	go test -v ./...

.PHONY: test-object-put
test-object-put:
	curl -i -XPUT http://127.0.0.1:6003/v1/account/bucket/object -d "123"
	@echo

.PHONY: test-object-get
test-object-get:
	curl -i -XGET http://127.0.0.1:6003/v1/account/bucket/object
	@echo

.PHONY: test-object-delete
test-object-delete:
	curl -i -XDELETE http://127.0.0.1:6003/v1/account/bucket/object
	@echo
