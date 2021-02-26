UNAME=$( command -v uname)
IMAGENAME="gobuild"

case $( "${UNAME}" | tr '[:upper:]' '[:lower:]') in
  linux*)
    printf 'linux\n'
    env GOOS="${GOOS}" GOARCH="${GOARCH}"
    go get -d -v ./...
    go build "${FLAGS}" -o ~/temp/VectorworksUtility
    ;;
  darwin*)
    printf 'darwin\n'
    ;;
  msys*|cygwin*|mingw*)
    # or possible 'bash on windows'
    printf 'windows\n'
    docker build -t ${IMAGENAME} .
    docker run ${IMAGENAME} -e flags='-ldflags="-H windowsgui"' -e os=windows -v ./:/go/src/app
    ;;
  nt|win*)
    printf 'windows\n'
    docker build -t ${IMAGENAME} .
    docker run ${IMAGENAME} \
      -e flags='-ldflags="-H windowsgui"' \
      -e os=windows \
      -v ./:/go/src/app
    ;;
  *)
    printf 'unknown\n'
    ;;
esac

echo "Press any key to exit"
while read -r -n 1 ; do
  break
done