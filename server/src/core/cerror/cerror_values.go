package cerror

var (
	BadRequestMsg          = "요청하신 내용이 잘못되었습니다. 다시 시도해주세요."
	NotFoundMsg            = "요청하신 내용을 찾을 수가 없습니다."
	InternalServerErrorMsg = "서버에 오류가 발생했습니다. 관리자에게 문의해주세요."
)

var (
	MysqlDuplicatedErrorNumber uint16 = 1062
)
