package codec
/*
import (
	"bytes"
	"gitlab.alipay-inc.com/afe/mosn/pkg/log"
	"gitlab.alipay-inc.com/afe/mosn/pkg/network/buffer"
	"testing"
)
func Test_TrRequestCommand_Parse(t *testing.T) {

	log.InitDefaultLogger("", log.INFO)

	strEchoBytes := []byte{0x0d, 0x00, 0x01, 0x02, 0x00, 0x00, 0x00, 0x00, 0x8f, 0x2d, 0x00, 0x00, 0x00, 0x57, 0x4d, 0x74, 0x00, 0x2a, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x6f, 0x62, 0x61, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x00, 0x03, 0x63, 0x74, 0x78, 0x4d, 0x74, 0x00, 0x39, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x6f, 0x62, 0x61, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x24, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x00, 0x02, 0x69, 0x64, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x53, 0x00, 0x06, 0x74, 0x68, 0x69, 0x73, 0x24, 0x30, 0x52, 0x00, 0x00, 0x00, 0x00, 0x7a, 0x7a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x74, 0x00, 0x2d, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x00, 0x02, 0x69, 0x64, 0x49, 0x00, 0x00, 0x00, 0x01, 0x53, 0x00, 0x03, 0x6d, 0x73, 0x67, 0x53, 0x00, 0x12, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2c, 0x20, 0x74, 0x73, 0x75, 0x69, 0x2e, 0x7a}

	trCodec := trCodec{}
	data := &buffer.IoBuffer{}

	data.ReadOne(bytes.NewBuffer(strEchoBytes))
	trRequestCommands := make([]trRequestCommand, 0)
	trCodec.Decode(data)

	command := trRequestCommands[0]
	if command.appClassName != "com.alipay.remoting.rpc.tr.util.ClientRequest" {

		log.DefaultLogger.Errorf("some error occurs")
	} else {
		log.DefaultLogger.Errorf("some decode success")

	}
}

func Test_TrResponseCommand_Parse(t *testing.T) {

	log.InitDefaultLogger("", log.INFO)

	strEchoBytes := []byte{0x0d, 0x01, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0xd3, 0x2e, 0x00, 0x00, 0x00, 0x62, 0x4d, 0x74, 0x00, 0x2b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x6f, 0x62, 0x61, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x00, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x53, 0x00, 0x0f, 0x31, 0x32, 0x37, 0x2e, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x3a, 0x35, 0x37, 0x34, 0x32, 0x38, 0x53, 0x00, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x00, 0x00, 0x00, 0x00, 0x53, 0x00, 0x08, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x73, 0x67, 0x4e, 0x53, 0x00, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4e, 0x53, 0x00, 0x03, 0x63, 0x74, 0x78, 0x4d, 0x74, 0x00, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x61, 0x6f, 0x62, 0x61, 0x6f, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x6d, 0x70, 0x6c, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x24, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x53, 0x00, 0x02, 0x69, 0x64, 0x4c, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x53, 0x00, 0x06, 0x74, 0x68, 0x69, 0x73, 0x24, 0x30, 0x52, 0x00, 0x00, 0x00, 0x00, 0x7a, 0x7a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x74, 0x00, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x6c, 0x69, 0x70, 0x61, 0x79, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x74, 0x72, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x00, 0x02, 0x69, 0x64, 0x49, 0x00, 0x00, 0x00, 0x63, 0x53, 0x00, 0x03, 0x6d, 0x73, 0x67, 0x53, 0x00, 0x1c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x20, 0x66, 0x72, 0x6f, 0x6d, 0x20, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2c, 0x20, 0x74, 0x73, 0x75, 0x69, 0x2e, 0x7a}

	trCodec := trCodec{}
	data := &buffer.IoBuffer{}

	data.ReadOne(bytes.NewBuffer(strEchoBytes))
	trResponseCommands := make([]trResponseCommand, 0)
	trCodec.Decode(nil, data, &trResponseCommands)

	command := trResponseCommands[0]
	if command.appClassName != "com.alipay.remoting.rpc.tr.util.ClientResponse" {

		log.DefaultLogger.Errorf("some error occurs")
	} else {
		log.DefaultLogger.Errorf("some decode success")

	}
}
*/