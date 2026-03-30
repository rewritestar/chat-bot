package repository

import (
	"errors"
	"strings"

	"github.com/go-sql-driver/mysql"

	"chat-bot/src/core/cerror"
)

func duplicatedEmailError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		if mysqlErr.Number == cerror.MysqlDuplicatedErrorNumber &&
			strings.Contains(mysqlErr.Message, "email") {
			return cerror.CommonError{
				Comment: "email is duplicated",
				Message: "같은 이메일의 계정이 이미 존재합니다. 다른 이메일을 사용해주세요.",
			}
		}
	}
	return err
}

func emailNotFoundError(err error) error {
	if strings.Contains(err.Error(), "record not found") {
		return cerror.CommonError{
			Comment: "email is not found",
			Message: "이메일에 해당하는 계정이 존재하지 않습니다.",
		}
	}
	return err
}
