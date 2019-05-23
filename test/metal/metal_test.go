package metal

import (
	"tool-go/library/test/metal/mock_metal"
	"github.com/golang/mock/gomock"
	"testing"
	"fmt"
)

func GetMetalName(mi Imetal) string {
	mi.GetName()
	return mi.GetName()
}

func SetMetalName(mi Imetal,name string) string {
	return  mi.SetName(name)
}


func TestMetalName(t *testing.T)  {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()
	mockMetal := mock_metal.NewMockImetal(mockCtl)   //mock_metal就是生成的mock代码，以包的形式存在

	m:=new(Metal)
	mockCtl.RecordCall(m,"GetName").Times(1)
	mockCtl.Call(m,"GetName")

	call:=mockMetal.EXPECT().GetName().Return("apple")
	mockMetal.EXPECT().GetName().Return("peer").After(call)   //注入期望的返回值

	mockedBrand:=GetMetalName(mockMetal)

	mockMetal.EXPECT().SetName(gomock.Eq("al")).Do(func(format string) {   //入参校验
		fmt.Println("recv param :",format)
	}).Return("setdone")

	mockMetal.EXPECT().SetName(gomock.Any()).Do(func(format string) {    //入参不做校验
		fmt.Println("recv param :",format)
	}).Return("setdone")

	mockedSetName:=SetMetalName(mockMetal,"al")
	fmt.Println(mockedSetName)

	if "peer"!=mockedBrand{
		t.Error("Get wrong name:", mockedBrand)
	}

	if "setdone"!=mockedSetName{
		t.Error("Set wrong name:", mockedSetName)
	}
}