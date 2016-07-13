../aceapi/aceapi -v
read -p "deploy pblog?"
curl \
	-H "Content-Type: application/octet-stream" \
	--cacert ../aceapi/acenet.crt \
	--data-binary @pblog.linux \
	-H "Token: `cat ../aceapi/token.txt`" \
	--url https://api.voilokov.com/v1/upload?dst=../../pblog/pblog

cd ../aceapi
./apiexec "chmod +x ../../pblog/pblog"
./apiexec "ls -l ../../pblog/pblog"

md5 pblog.linux
