.PHONY: gitag all common user region
# 注意: 新的protoc-gen-go插件已经不支持plugins选项

export TagName := v1.0.2
.PHONY: gitag
gitag:
	git add . && git commit -m "$(TagName)"
	git push
	git tag -a $(TagName) -m "$(TagName)" # 创建带标签的Tag
	git push origin $(TagName)  # 推送Tag到远程

all: 
	make common
	make user
	make region

common: 
	protoc --proto_path=proto/common \
	--go_out=. \
	--go-grpc_out=. \
	proto/common/*.proto

user:
	protoc --proto_path=proto/user \
	--proto_path=proto/common \
	--go_out=. \
	--go-grpc_out=. \
	proto/user/*.proto

region :
	protoc --proto_path=proto/region \
	--proto_path=proto/common \
	--go_out=. \
	--go-grpc_out=. \
	proto/region/*.proto
