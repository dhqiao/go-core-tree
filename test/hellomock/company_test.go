package hellomock

import (
	"testing"
	"github.com/golang/mock/gomock"
	"tool-go/library/test/hellomock/mock"
)

func TestCompany_Meeting2(t *testing.T) {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("王尼玛")).Return("这是自定义的返回值，可以是任意类型。")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("王尼玛"))
	//t.Log(company.Meeting("张全蛋"))
}

func Test_Meeting(t *testing.T)  {
	ctl := gomock.NewController(t)
	mock_talker := mock_hellomock.NewMockTalker(ctl)
	mock_talker.EXPECT().SayHello(gomock.Eq("你好")).Return("自定义")

	company := NewCompany(mock_talker)
	t.Log(company.Meeting("你好me "))

}