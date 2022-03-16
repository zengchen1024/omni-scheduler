package scheduler

import (
	"fmt"

	"github.com/opensourceways/scheduler/models"
	"github.com/opensourceways/scheduler/solv"
)

type eventHandler struct {
	gctx *GCtx
}

func (h eventHandler) eventScanRepo(e models.EventScanRepo) error {
	fmt.Println(h.gctx.arch)

	c := NewChecker(h.gctx, fmt.Sprintf("%s/%s", e.Project, e.Repository), h.gctx.arch)
	pool := solv.NewPool()
	defer pool.Free()

	return c.addRepo(pool, c.prp)
}

func HandleEvent(gctx *GCtx, e interface{}) error {
	h := eventHandler{gctx}

	switch v := e.(type) {
	case models.EventScanRepo:
		return h.eventScanRepo(v)
	default:
		return fmt.Errorf("unkonw event type: %T", v)
	}
}
