package middlewire

// func VerifyUserToken(c *gin.Context) (string, error) {

// 	fmt.Println("user middlewiere")
// 	id, err := helper.VerifyAcessToken(accessToken, u.token.UserSecurityKey)
// 	if err != nil {
// 		fmt.Println("error at accestoken", err)
// 	}

// 	err = helper.VerifyRefreshToken(refreshToken, u.token.UserSecurityKey)
// 	if err != nil {
// 		return "", err
// 	}
// 	return id, nil
// }


// func (u *adminUsecase) VerifyAdminToken(token string) (*emptypb.Empty, error) {
// 	fmt.Println("admin middlewire auth")
// 	err := service_auth_svc.VerifyRefreshToken(token, u.tokenSecurityKey.AdminSecurityKey)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &emptypb.Empty{}, nil
// }
