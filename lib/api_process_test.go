package lib

import (
	"context"
	"fmt"
	"testing"

	"github.com/antihax/optional"
)

func TestGetProcess(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	sib := NewAPIClient(cli.cfg)
	result, resp, err := sib.ProcessApi.GetProcess(ctx, 10)
	if err != nil {
		fmt.Println("===in get process error===")
		t.Fatal(err)
	}
	fmt.Println("====Process Result:", result, "     ====    resp:", resp)
}

func TestGetProcesses(t *testing.T) {
	var ctx context.Context
	var cli = APIClient{
		cfg: NewConfiguration(),
	}
	cli.cfg.AddDefaultHeader("api-key", "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")
	sib := NewAPIClient(cli.cfg)
	params := &ProcessApiGetProcessesOpts{
		Sort: optional.NewString("asc"),
	}
	result, resp, err := sib.ProcessApi.GetProcesses(ctx, params)
	if err != nil {
		fmt.Println("===in get process error===")
		t.Fatal(err)
	}
	fmt.Println("====All Process Result:", result, "     ====    resp:", resp)
}
