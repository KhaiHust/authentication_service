package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IVerifyOtpUseCase interface {
	VerifyOtpRegister(ctx *context.Context, email string, otp string) error
}
type VerifyOtpUseCase struct {
	cachePort                  port.ICachePort
	getUserUsecase             IGetUserUsecase
	updateUserUseCase          IUpdateUserUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (v VerifyOtpUseCase) VerifyOtpRegister(ctx *context.Context, email string, otp string) error {

	existedUser, err := v.getUserUsecase.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(ctx, "VerifyOtpRegister: GetUserByEmail error", err)
		return err
	}
	key := common.BuildCacheKeyOTPRegister(email)
	cachedOtp, err := v.cachePort.GetFromCache(*ctx, key)
	if err != nil && err.Error() != constant.ErrCacheKeyNil {
		log.Error(ctx, "VerifyOtpRegister: GetFromCache error", err)
		return err
	}
	if (err != nil && err.Error() == constant.ErrCacheKeyNil) || cachedOtp == "" {
		return errors.New(constant.ErrOtpNotFound)
	}
	cachedOtpStr, ok := cachedOtp.(string)
	if !ok {
		return errors.New(constant.ErrOtpInvalid)
	}
	cachedOtpStr = cachedOtpStr[1 : len(cachedOtpStr)-1]
	if cachedOtpStr != otp {
		return errors.New(constant.ErrOtpInvalid)
	}
	err = v.cachePort.DeleteFromCache(*ctx, key)
	if err != nil {
		log.Error(ctx, "VerifyOtpRegister: DeleteFromCache error", err)
		return err
	}
	//update status user verified
	existedUser.IsVerified = true
	tx := v.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			log.Error(ctx, "VerifyOtpRegister: Recovered in VerifyOtpRegister", errors.New(r.(string)))
			err = exception.InternalServerErrorException
		}
		if err != nil {
			errRollback := v.databaseTransactionUsecase.Rollback(tx)
			if errRollback != nil {
				log.Error(ctx, "VerifyOtpRegister: Rollback error", errRollback)
			} else {
				log.Info(ctx, "VerifyOtpRegister: Rollback transaction")
			}

		}
	}()
	_, err = v.updateUserUseCase.UpdateVerifiedUser(ctx, existedUser, tx)
	if err != nil {
		log.Error(ctx, "VerifyOtpRegister: UpdateVerifiedUser error", err)
		return err
	}
	err = v.databaseTransactionUsecase.Commit(tx)
	if err != nil {
		log.Error(ctx, "VerifyOtpRegister: Commit error", err)
		return err
	}
	err = v.cachePort.DeleteFromCache(*ctx, common.BuildCacheKeyGetUserInfoByEmail(email))
	if err != nil {
		log.Error(ctx, "VerifyOtpRegister: DeleteFromCache error", err)
		return err
	}
	return nil
}

func NewVerifyOtpUseCase(cachePort port.ICachePort, getUserUsecase IGetUserUsecase, updateUserUseCase IUpdateUserUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) IVerifyOtpUseCase {
	return &VerifyOtpUseCase{
		cachePort:                  cachePort,
		getUserUsecase:             getUserUsecase,
		updateUserUseCase:          updateUserUseCase,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
