#### [ 注意: 新的protoc-gen-go插件已经不支持plugins选项 ]

# sdk-pb目录
SdkPbDir = pb
# error配置
ErrorxInputDir = errorx
# 第三方proto目录
ThirdPartyDir = third_party
# sdk的git版本号
TagName = v1.0.7
# sdk业务proto地址
ProtoCommonInput = proto/common
ProtoUserInput   = proto/user
ProtoRegionInput = proto/region

.PHONY: gitag all pb error copypb
gitag:
	git add . && git commit -m "$(TagName)"
	git push
	git tag -a $(TagName) -m "$(TagName)" # 创建带标签的Tag
	git push origin $(TagName)  # 推送Tag到远程

all:
	@make pb
	@make error
	@make copypb

pb:
	@echo ">>>>>>>>>>>> [ gen sdk ] >>>>>>>>>>>>"
	protoc --proto_path=$(ProtoCommonInput) \
	--proto_path=$(ProtoUserInput) \
	--proto_path=$(ProtoRegionInput) \
	--go_out=. \
	--go-grpc_out=. \
	$(ProtoCommonInput)/*.proto \
	$(ProtoUserInput)/*.proto \
	$(ProtoRegionInput)/*.proto

copypb:
	@echo ">>>>>>>>>>>> [ copy pb ] >>>>>>>>>>>>"
	cp -rf github.com/wangzhe1991/grpc-sdk/pb/* ./pb

error:
	@echo ">>>>>>>>>>>> [ gen errors ] >>>>>>>>>>>>"
	protoc --proto_path=. \
     --proto_path=$(ThirdPartyDir) \
     --go_out=paths=source_relative:$(SdkPbDir) \
     --go-errorx_out=paths=source_relative:$(SdkPbDir) \
     $(ErrorxInputDir)/*.proto
