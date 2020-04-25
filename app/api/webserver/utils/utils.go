package utils

//func GetGRPCService() micro.Service {
//
//	//配置注册中心
//	consulReg := consul.NewRegistry(func(options *registry.Options) {
//		options.Addrs = []string{Consul_ADDR + ":" + Consul_PORT}
//	})
//	//创建 grpc 服务
//	grpcService := grpc.NewService(
//		micro.Registry(consulReg),
//	)
//
//	return grpcService
//}

//获取前端传过来的图片
//func GetImageBuf(ctx *gin.Context, fileName string) ([]byte, string, error) {
//	resp := new(postHousesImageSrv.Response)
//	//参数, 前端传递的key值
//	fileHeader, err := ctx.FormFile(fileName)
//	if err != nil {
//		resp.Errno = RECODE_NODATA
//		resp.Errmsg = RecodeText(RECODE_NODATA)
//		//422 Unprocessable Entity ：客户端上传的附件无法处理，导致请求失败。
//		ctx.JSON(422, resp)
//		errLog.Error.Println("获取前端文件失败: ", err)
//		return nil, "", err
//	}
//
//	//check
//	if fileHeader.Size > 1024*1024*10 {
//		resp.Errno = RECODE_DATAERR
//		resp.Errmsg = RecodeText(RECODE_DATAERR)
//		//422 Unprocessable Entity ：客户端上传的附件无法处理，导致请求失败。
//		ctx.JSON(422, resp)
//		errLog.Error.Println("图片超过10M,请重新上传")
//		return nil, "", err
//	}
//
//	fileExt := path.Ext(fileHeader.Filename)
//	if fileExt != ".jpg" && fileExt != ".png" && fileExt != ".jepg" {
//		resp.Errno = RECODE_DATAERR
//		resp.Errmsg = RecodeText(RECODE_DATAERR)
//		ctx.JSON(422, resp)
//		errLog.Error.Println("图片格式不正确")
//		return nil, "", err
//	}
//
//	file, err := fileHeader.Open()
//	if err != nil {
//		resp.Errno = RECODE_IOERR
//		resp.Errmsg = RecodeText(RECODE_IOERR)
//		ctx.JSON(500, resp)
//		errLog.Error.Println("fileHeader.Open err: ", err)
//		return nil, "", err
//	}
//	defer file.Close()
//
//	fileBuf := make([]byte, fileHeader.Size)
//	_, err = file.Read(fileBuf)
//	if err != nil {
//		resp.Errno = RECODE_IOERR
//		resp.Errmsg = RecodeText(RECODE_IOERR)
//		ctx.JSON(500, resp)
//		errLog.Error.Println("file.Read err: ", err)
//		return nil, "", err
//	}
//
//	return fileBuf, fileExt, nil
//}