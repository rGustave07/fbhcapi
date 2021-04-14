.PHONY: create-keypair

PWD = $(shell pwd)
APIFOLDERPATH = $(PWD)/api

create-keypair:
	@echo "Creating an rsa 256 key pair"
	openssl genpkey -algorithm RSA -out $(APIFOLDERPATH)/rsa_private_$(ENV).pem -pkeyopt rsa_keygen_bits:2048
	openssl rsa -in $(APIFOLDERPATH)/rsa_private_$(ENV).pem -pubout -out $(APIFOLDERPATH)/rsa_public_$(ENV).pem