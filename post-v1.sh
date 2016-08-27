echo aceapi version:
../aceapi/aceapi -v

read -p "deploy pblog?"
curl \
	-H "Content-Type: application/octet-stream" \
	--cacert ../aceapi/acenet.crt \
	--data-binary @pblog.linux \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../pblog/pblog

curl \
	-H "Content-Type: application/octet-stream" \
	--cacert ../aceapi/acenet.crt \
	--data-binary @templates/toc.html \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../pblog/templates/toc.html

curl \
	-H "Content-Type: application/octet-stream" \
	--cacert ../aceapi/acenet.crt \
	--data-binary @templates/main.html \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../pblog/templates/main.html

../aceapi/apiexec "chmod +x ../../pblog/pblog"
../aceapi/apiexec "../../pblog/pblog -v"
../aceapi/apiexec "md5sum ../../pblog/pblog"

md5 pblog.linux
