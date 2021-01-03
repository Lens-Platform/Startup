doc-site: # Spin us documentation site locally
	cd docs && yarn run serve --build --port 9999 --host 0.0.0.0

signoff: ## Signsoff all previous commits since branch creation
	scripts/signoff.sh

release: # Invokes a script to automate the creation of a release

benchmarks: # Invokes a script to run various benchmarks (E2E)

run-tests: # Invokes a script able to run many type of tests. Reads args. and deciphers wether to run e2e, unit, stress testing ....etc

spin-up-kube: # Spins up local mini kube cluster
