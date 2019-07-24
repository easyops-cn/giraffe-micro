package nameservice_test

import (
	"testing"

	"github.com/easyops-cn/giraffe-micro/plugins/easyopsrest/nameservice"
)

func TestSetAppName(t *testing.T) {
	nameservice.SetAppName("test")
}
