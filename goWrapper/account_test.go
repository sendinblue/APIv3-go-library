package newtest

import (
	"context"
	"fmt"
	"testing"
)

func TestGetAccount(t *testing.T) {
	var ctx context.Context
	var aa = APIClient{
		cfg: NewConfiguration(),
	}
	//aa.cfg = NewConfiguration()
	fmt.Println(aa)
	sib := NewAPIClient(aa.cfg)
	result, resp, err := sib.AccountApi.GetAccount(ctx)
	if err != nil {
		fmt.Println("===in get account error===")
		t.Fatal(err)
	}
	fmt.Println("====Account Result:", result, "     ====    resp:", resp)
}
