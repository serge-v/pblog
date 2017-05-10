read -p "deploy pblog/pblog?"

md5 pblog.linux

curl \
	-H "Content-Type: application/octet-stream" \
	--data-binary @pblog.linux \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url "https://api.voilokov.com/v1/upload?dst=../../pblog/pblog&mode=0700&md5=1"

