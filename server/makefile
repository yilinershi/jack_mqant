start:
	go run main.go

build_proto:
	@echo "step1:先清掉pb文件夹"
	@rm -rf pb
	@mkdir pb
	@echo "step2:再生成pb文件"
	@protoc --go_out=../ -I=../proto  ../proto/*.proto
	@echo "生成pb文件成功"

build_proto_cs:
	@echo "step1:先清掉pb文件夹"
	@rm -rf ../client/Assets/Pb
	@mkdir ../client/Assets/Pb
	@echo "step2:再生成pb文件"
	@protoc --csharp_out=../client/Assets/Pb -I=../proto  ../proto/*.proto
	@echo "生成pb文件成功"