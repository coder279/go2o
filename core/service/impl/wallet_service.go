package impl

import (
	"context"
	"go2o/core/domain/interface/wallet"
	"go2o/core/service/proto"
	"go2o/core/service/thrift/parser"
)

var _ proto.WalletServiceServer = new(walletServiceImpl)

func NewWalletService(repo wallet.IWalletRepo)*walletServiceImpl {
	return &walletServiceImpl{
		_repo: repo,
	}
}

type walletServiceImpl struct {
	_repo wallet.IWalletRepo
	serviceUtil
}

func (w *walletServiceImpl) CreateWallet(_ context.Context, r *proto.CreateWalletRequest) (*proto.Result, error) {
	v := &wallet.Wallet{
		UserId:     r.UserId,
		WalletType: int(r.WalletType),
		WalletFlag: int(r.Flag),
		Remark:     r.Remark,
	}
	iw := w._repo.CreateWallet(v)
	_, err := iw.Save()
	return w.result(err), nil
}

func (w *walletServiceImpl) GetWalletId(_ context.Context, r *proto.GetWalletRequest) (*proto.Int64, error) {
	iw := w._repo.GetWalletByUserId(r.UserId, int(r.WalletType))
	if iw != nil {
		return &proto.Int64{Value: iw.GetAggregateRootId()}, nil
	}
	return &proto.Int64{Value: 0}, nil
}

func (w *walletServiceImpl) GetWallet(_ context.Context, walletId *proto.Int64) (*proto.SWallet, error) {
	iw := w._repo.GetWallet(walletId.Value)
	if iw != nil {
		return w.parseWallet(iw.Get()), nil
	}
	return nil, nil
}

func (w *walletServiceImpl) GetWalletLog(_ context.Context, r *proto.WalletLogIDRequest) (*proto.SWalletLog, error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw != nil {
		if l := iw.GetLog(r.Id); l.ID > 0 {
			return w.parseWalletLog(l), nil
		}
	}
	return nil, nil
}
func (w *walletServiceImpl) Adjust(_ context.Context, r *proto.AdjustRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.Adjust(int(r.Value), r.Title, r.OuterNo, int(r.OpuId), r.OpuName)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) Discount(_ context.Context, r *proto.DiscountRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.Discount(int(r.Value), r.Title, r.OuterNo, r.Must)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) Freeze(_ context.Context, r *proto.FreezeRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.Freeze(int(r.Value), r.Title, r.OuterNo, int(r.OpuId), r.OpuName)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) Unfreeze(_ context.Context, r *proto.UnfreezeRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.Unfreeze(int(r.Value), r.Title, r.OuterNo, int(r.OpuId), r.OpuName)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) Charge(_ context.Context, r *proto.ChargeRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.Charge(int(r.Value), int(r.By), r.Title,
			r.OuterNo, int(r.OpuId), r.OpuName)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) Transfer(_ context.Context, r *proto.TransferRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		title := "钱包转账"
		toTitle := "钱包收款"
		//todo: title
		err = iw.Transfer(r.ToWalletId, int(r.Value),
			int(r.TradeFee), title, toTitle, r.Remark)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) RequestTakeOut(_ context.Context, r *proto.RequestTakeOutRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		_, tradeNo, err1 := iw.RequestTakeOut(int(r.Value),
			int(r.TradeFee), int(r.Kind), r.Title)
		if err1 != nil {
			err = err1
		} else {
			return w.success(map[string]string{
				"TradeNo": tradeNo,
			}), nil
		}
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) ReviewTakeOut(_ context.Context, r *proto.ReviewTakeOutRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.ReviewTakeOut(r.TakeId, r.ReviewPass, r.Remark, int(r.OpuId), r.OpuName)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) FinishTakeOut(_ context.Context, r *proto.FinishTakeOutRequest) (ro *proto.Result, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		err = wallet.ErrNoSuchWalletAccount
	} else {
		err = iw.FinishTakeOut(r.TakeId, r.OuterNo)
	}
	return w.result(err), nil
}

func (w *walletServiceImpl) PagingWalletLog(_ context.Context, r *proto.PagingWalletLogRequest) (ro *proto.SPagingResult, err error) {
	iw := w._repo.GetWallet(r.WalletId)
	if iw == nil {
		return parser.PagingResult(0, nil, wallet.ErrNoSuchWalletAccount), nil
	}
	total, list := iw.PagingLog(int(r.Params.Begin),
		int(r.Params.Over), r.Params.Parameters,
		r.Params.SortBy)
	return parser.PagingResult(total, list, err), nil
}

func (w *walletServiceImpl) parseWallet(v wallet.Wallet) *proto.SWallet {
	return &proto.SWallet{
		ID:             v.ID,
		HashCode:       v.HashCode,
		NodeId:         int32(v.NodeId),
		UserId:         v.UserId,
		WalletType:     int32(v.WalletType),
		WalletFlag:     int32(v.WalletFlag),
		Balance:        int32(v.Balance),
		PresentBalance: int32(v.PresentBalance),
		AdjustAmount:   int32(v.AdjustAmount),
		FreezeAmount:   int32(v.FreezeAmount),
		LatestAmount:   int32(v.LatestAmount),
		ExpiredAmount:  int32(v.ExpiredAmount),
		TotalCharge:    int32(v.TotalCharge),
		TotalPresent:   int32(v.TotalPresent),
		TotalPay:       int32(v.TotalPay),
		State:          int32(v.State),
		Remark:         v.Remark,
		CreateTime:     v.CreateTime,
		UpdateTime:     v.UpdateTime,
	}
}
func (w *walletServiceImpl) parseWalletLog(l wallet.WalletLog) *proto.SWalletLog {
	return &proto.SWalletLog{
		ID:           l.ID,
		WalletId:     l.WalletId,
		Kind:         int32(l.Kind),
		Title:        l.Title,
		OuterChan:    l.OuterChan,
		OuterNo:      l.OuterNo,
		Value:        int32(l.Value),
		Balance:      int32(l.Balance),
		TradeFee:     int32(l.TradeFee),
		OperatorId:   int32(l.OperatorId),
		OperatorName: l.OperatorName,
		Remark:       l.Remark,
		ReviewState:  int32(l.ReviewState),
		ReviewRemark: l.ReviewRemark,
		ReviewTime:   l.ReviewTime,
		CreateTime:   l.CreateTime,
		UpdateTime:   l.UpdateTime,
	}
}
