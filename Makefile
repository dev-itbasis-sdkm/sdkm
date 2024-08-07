REPORTS_DIR = build/reports

all: go-update go-generate go-all-tests go-build-bin go-multiple-os-distributions
go-all-tests: go-lint go-unit-tests

before-reports:
	mkdir -p "${REPORTS_DIR}"
	rm -Rf "${REPORTS_DIR}"/*

go-dependencies:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/onsi/ginkgo/v2/ginkgo@latest
	go install go.uber.org/mock/mockgen@latest

go-update: go-dependencies
	go mod tidy && go get -t -v -u ./...

go-generate:
	go generate ./...
	$(MAKE) go-update

go-lint:
	golangci-lint run

go-unit-tests: before-reports
	ginkgo -race --cover --coverprofile="${REPORTS_DIR}/ginkgo-coverage-unit.out" --junit-report="${REPORTS_DIR}/junit-report.xml" ./...
	grep -v ".mock.go" "${REPORTS_DIR}/ginkgo-coverage-unit.out" >"${REPORTS_DIR}/ginkgo-coverage-unit.clean.out"
	go tool cover -func "${REPORTS_DIR}/ginkgo-coverage-unit.clean.out" -o "${REPORTS_DIR}/coverage-unit.out"
	go tool cover -html "${REPORTS_DIR}/ginkgo-coverage-unit.clean.out" -o "${REPORTS_DIR}/coverage-unit.html"

go-prebuild:
	echo "make target dir for build: ${targetDir}"
	rm -Rf "${targetDir}" 2>/dev/null
	mkdir -p "${targetDir}"

go-build:
	$(MAKE) go-prebuild targetDir="${targetDir}"
	echo "build version='${version}' to target dir: ${targetDir}"
	GOOS="${targetOS}" GOARCH="${targetARCH}" go build -trimpath -pgo auto -ldflags "-w -extldflags '-static' -X 'github.com/dev.itbasis.sdkm/internal/version.Version=${version}'" -tags musl -o "${targetDir}/sdkm" "./cmd/main.go"
	$(MAKE) copy-docs targetDir="${targetDir}"

copy-docs:
	cp ./README.adoc ./CHANGELOG.adoc "${targetDir}/"
	cp -R ./changelog "${targetDir}/"

go-build-bin:
	$(eval targetOS=$(uname -s | tr '[:upper:]' '[:lower:]'))
	$(eval targetARCH=$(uname -m))
	$(MAKE) go-build targetDir="bin" targetOS="${targetOS}" targetARCH="${targetARCH}"

go-multiple-os-distributions:
	$(MAKE) go-make-distribution targetDir="build/darwin-amd64" targetOS="darwin" targetARCH="amd64"
	$(MAKE) go-make-distribution targetDir="build/darwin-arm64" targetOS="darwin" targetARCH="arm64"
	$(MAKE) go-make-distribution targetDir="build/linux-ppc64" targetOS="linux" targetARCH="ppc64"
	$(MAKE) go-make-distribution targetDir="build/linux-arm64" targetOS="linux" targetARCH="arm64"
	#$(MAKE) go-make-distribution targetDir="build/windows-amd64" targetOS="windows" targetARCH="amd64"
	#$(MAKE) go-make-distribution targetDir="build/windows-arm64" targetOS="windows" targetARCH="arm64"

go-make-distribution:
	mkdir -p "distributions"
	$(MAKE) go-build targetDir="${targetDir}" targetOS="${targetOS}" targetARCH="${targetARCH}"
	tar -zcf "distributions/${targetOS}-${targetARCH}.tar.gz" -C "${targetDir}/" .
