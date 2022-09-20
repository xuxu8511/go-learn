package xnet

type Request interface {
	GetConnection() Connection

	GetData() Message

	//GetMsgID() uint32
}

type RequestImpl struct {
	conn Connection
	msg  Message
}

func (r *RequestImpl) GetConnection() Connection {
	return r.conn
}

func (r *RequestImpl) GetData() Message {
	return r.msg
}

//func (r *RequestImpl) GetMsgID() uint32 {
//
//}
