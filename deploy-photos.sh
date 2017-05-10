read -p "deploy photos/pblog?"
shasum -a 256 pblog.linux

curl \
	-H "Content-Type: application/octet-stream" \
	--data-binary @pblog.linux \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url "https://api.voilokov.com/v1/file?dst=../../photos/pblog&mode=0700"

