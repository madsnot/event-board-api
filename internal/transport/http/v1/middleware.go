package v1

const headerAuthorize = "Authorization"

//
//func AuthMiddleware()
//func checkAuthorization(ctx context.Context, expectedScheme string) (string, error) {
//	value := ctx.Value(headerAuthorize).(string)
//
//	if value == "" {
//		return "", fmt.Errorf("request unauthenticated with %s", expectedScheme)
//	}
//
//	splits := strings.SplitN(value, " ", 2)
//	if len(splits) < 2 {
//		return "", errors.New("bad authorization string")
//	}
//
//	if !strings.EqualFold(splits[0], expectedScheme) {
//		return "", fmt.Errorf("request unauthenticated with %s", expectedScheme)
//	}
//
//	return splits[1], nil
//}
