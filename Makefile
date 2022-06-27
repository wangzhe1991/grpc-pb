.PHONY: gitag all error copy
# 注意: 新的protoc-gen-go插件已经不支持plugins选项
# error-proto路径
ERRORX_INPUT_DIR  = ./errorx
ERRORX_OUTPUT_DIR = ./pb


THIRD_PARTY_DIR = ./vendor

export TagName := v1.0.5
gitag:
	git add . && git commit -m "$(TagName)"
	git push
	git tag -a $(TagName) -m "$(TagName)" # 创建带标签的Tag
	git push origin $(TagName)  # 推送Tag到远程


all :
	protoc --proto_path=proto/common \
	--proto_path=proto/user \
	--proto_path=proto/region \
	--go_out=. \
	--go-grpc_out=. \
	proto/common/*.proto \
	proto/region/*.proto \
	proto/user/*.proto \

copy:
	cp -rf github.com/wangzhe1991/grpc-sdk/pb/* ./pb

error:
	@echo ">>> errors......"
	protoc --proto_path=. \
     --proto_path=$(THIRD_PARTY_DIR) \
     --go_out=paths=source_relative:$(ERRORX_OUTPUT_DIR) \
     --go-errors_out=paths=source_relative:$(ERRORX_OUTPUT_DIR) \
     $(ERRORX_INPUT_DIR)/*.proto
