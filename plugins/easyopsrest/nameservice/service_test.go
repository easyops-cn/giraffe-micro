package nameservice_test

import (
	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/nameservice"
	"testing"
)

func TestSetAppName(t *testing.T) {
	nameservice.SetAppName("test")
}
