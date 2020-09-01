package spotinst

import (
	"github.com/YaleSpinup/apierror"
)

func ErrCode(msg string, err error) error {
	return apierror.New(apierror.ErrBadRequest, msg, err)
}
