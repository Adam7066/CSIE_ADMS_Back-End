# CSIE_ADMS_Back-End

## Setup
1. `cp ./manifest/config/config.yaml.example ./manifest/config/config.yaml`
2. 編輯 `./manifest/config/config.yaml`，修改相關資訊
3. 執行 `go run main.go --init`
   - 初始化 Database
   - 新增預設 Admin 帳號

## Run
- `gf run main.go`

## 預設 Admin 帳號
- Email：`admin@csie.adms`
- Password：`1234567890`

> [!IMPORTANT]
> 實際使用時，登入後請記得修改帳號、密碼

## API
- Swagger: http://127.0.0.1:8000/swagger/
