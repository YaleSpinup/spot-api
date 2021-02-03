package spotinst

import (
	"github.com/YaleSpinup/apierror"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spotinst/spotinst-sdk-go/spotinst/client"
)

func ErrCode(msg string, err error) error {
	log.Debugf("processing error code with message '%s' and error '%s'", msg, err)

	if spotErrs, ok := errors.Cause(err).(client.Errors); ok && len(spotErrs) > 0 {
		if len(spotErrs) > 1 {
			log.Warnf("got %d errors, only the first one will be returned", len(spotErrs))
		}

		aerr := spotErrs[0]

		switch aerr.Code {
		case "MANAGED_INSTANCE_DOESNT_EXIST":
			return apierror.New(apierror.ErrNotFound, aerr.Message, aerr)
		default:
			m := msg + ": " + aerr.Message
			return apierror.New(apierror.ErrBadRequest, m, aerr)
		}
	}

	return apierror.New(apierror.ErrInternalError, msg, err)
}
