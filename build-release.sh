#!/bin/bash

name='ip2region'

MD5='md5sum'
if [[ "$(uname)" == 'Darwin' ]]; then
	MD5='md5'
fi

LDFLAGS="-s -w"

# X86
OSES=(windows linux darwin freebsd)
ARCHS=(amd64 386)
rm -rf ./release
mkdir -p ./release
for os in ${OSES[@]}; do
	for arch in ${ARCHS[@]}; do
		suffix=""
		if [ "$os" == "windows" ]; then
			suffix=".exe"
		fi
		env CGO_ENABLED=0 GOOS=$os GOARCH=$arch go build -ldflags "$LDFLAGS" -o ./release/${name}_${os}_${arch}${suffix} .
		tar -C ./release -zcf ./release/${name}_${os}-${arch}.tar.gz ./${name}_${os}_${arch}${suffix}
		$MD5 ./release/${name}_${os}-${arch}.tar.gz
	done
done

# ARM64
env CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "$LDFLAGS" -o ./release/${name}_arm64 .
# ARM
ARMS=(5 6 7)
for v in ${ARMS[@]}; do
	env CGO_ENABLED=0 GOOS=linux GOARCH=arm GOARM=$v go build -ldflags "$LDFLAGS" -o ./release/${name}_arm$v .
done
tar -C ./release -zcf ./release/${name}_arm.tar.gz $(for v in ${ARMS[@]}; do echo -n "./${name}_arm$v ";done)
$MD5 ./release/${name}_arm.tar.gz

# MIPS/hardfloat
env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle go build -ldflags "$LDFLAGS" -o ./release/${name}_mipsle .
env CGO_ENABLED=0 GOOS=linux GOARCH=mips go build -ldflags "$LDFLAGS" -o ./release/${name}_mips .
# MIPS/softfloat
env CGO_ENABLED=0 GOOS=linux GOARCH=mipsle GOMIPS=softfloat go build -ldflags "$LDFLAGS" -o ./release/${name}_mipsle_sf .
env CGO_ENABLED=0 GOOS=linux GOARCH=mips GOMIPS=softfloat go build -ldflags "$LDFLAGS" -o ./release/${name}_mips_sf .
# if $UPX; then upx -9 ./release/${name}_mips**; fi
tar -C ./release -zcf ./release/${name}_mipsle.tar.gz ./${name}_mipsle
tar -C ./release -zcf ./release/${name}_mips.tar.gz ./${name}_mips
tar -C ./release -zcf ./release/${name}_mipsle-sf.tar.gz ./${name}_mipsle_sf
tar -C ./release -zcf ./release/${name}_mips-sf.tar.gz ./${name}_mips_sf
$MD5 ./release/${name}_mipsle.tar.gz