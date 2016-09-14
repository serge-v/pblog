echo aceapi version:
../aceapi/aceapi -v

read -p "deploy photos/pblog?"

curl \
	-H "Content-Type: application/octet-stream" \
	--data-binary @pblog.linux \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../photos/pblog

curl \
	-H "Content-Type: application/octet-stream" \
	--data-binary @templates/toc.html \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../photos/templates/toc.html

curl \
	-H "Content-Type: application/octet-stream" \
	--data-binary @templates/main.html \
	-H "Token: `cat ~/.config/aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../photos/templates/main.html

../aceapi/apiexec "chmod +x ../../photos/pblog"
../aceapi/apiexec "../../photos/pblog -v"
../aceapi/apiexec "md5sum ../../photos/pblog"

md5 pblog.linux
