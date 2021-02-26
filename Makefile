buildmac:
	go build -o ./build/AppCleaner.app/Contents/MacOS/
	echo "[OK] App Binary build Successful"

buildwin:
	go build -ldflags="-H windowsgui" -o VectorworksUtility.exe
	echo "[OK] App Binary Build Successful"
