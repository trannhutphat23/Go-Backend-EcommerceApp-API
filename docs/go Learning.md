0. fmt.Println(), fmt.Printf()

1. struct, struct pointer, json.Marshal()

2. interface, type assertion


## Keywords
*gin.Engine = toàn bộ web server Gin, instance của Gin

*gin.Context = toàn bộ thông tin của request + response + tiện ích để xử lý chúng trong Gin -> 1 request cụ thể đang được xử lý bởi server đó

gin.Default() = gin.New() + gin.Logger() + gin.Recovery() => tạo ra đối tượng *gin.Engine

gin.Logger() và gin.Recovery() là 2 middleware
+ gin.Logger(): Ghi log ra console cho mỗi request (method, path, status, thời gian xử lý, v.v.)
+ gin.Recovery(): Bắt lỗi panic trong quá trình xử lý request, ngăn server bị crash, và trả về lỗi 500

gin.H : kiểu dữ liệu tiện lợi để tạo JSON Object, alias cho kiểu map[string]interface{} (type H map[string]interface{}) => map có key là chuỗi và value là bất kỳ kiểu dữ liệu nào

### GIN vs Router
### GIN vs MVC
### GIN vs Error
### GIN vs ZAP

* ZAP:
Có 3 loại cấu hình logger mặc định chính:
+ zap.NewExample(): minh họa hoặc test đơn giản
+ zap.NewDevelopment(): phát triển (dev mode)
+ zap.NewProduction(): môi trường production / server thật
+ zap.Config{...}.Build(): tự cấu hình logger theo ý muốn

Encoder:
zap.NewProductionEncoderConfig()

``
func NewProductionEncoderConfig() zapcore.EncoderConfig {
    return zapcore.EncoderConfig{
        TimeKey:        "ts",
        LevelKey:       "level",
        NameKey:        "logger",
        CallerKey:      "caller",
        MessageKey:     "msg",
        StacktraceKey:  "stacktrace",
        LineEnding:     zapcore.DefaultLineEnding,
        EncodeLevel:    zapcore.LowercaseLevelEncoder,
        EncodeTime:     zapcore.EpochTimeEncoder,
        EncodeDuration: zapcore.SecondsDurationEncoder,
        EncodeCaller:   zapcore.ShortCallerEncoder,
    }
}
``

func NewJSONEncoder(cfg EncoderConfig) Encoder -> hàm tạo