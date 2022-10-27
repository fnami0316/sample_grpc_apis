// アカウント関連サービスの実装

package service // serviceディレクトリにあるので・・・

import (
	"log" // 何かと使うので。
	"github.com/fnami0316/sample_grpc_apis/pb" // このパッケージの「AccountServiceServer」インターフェースの実装をするので。
  "google.golang.org/protobuf/types/known/emptypb"  // emptypb.Empty使うので
	"context" // context.Context使うので
)

// 以下、本質的な部分になる「account_grpc.pb.go」の「AccountServiceServer」インターフェースの実装

// -- 構造体
type AccountServiceServer struct { 
	pb.UnimplementedAccountServiceServer
}

// [メモ]
// 引数やレシーバの変数は、s, ctx, reqにするのが普通？
// pbパッケージをインポートしてるので、pb.をつけてけば_grpc.pb.go, .pb.goにかかれている型は使えるようになる
// pbパッケージの中で、contextもインポートしてるのでcontextも大丈夫なはずなのだが波線が・・・？
// コンストラクタ

// -- 構造体をレシーバとするメソッド
// アカウント作成
// [レシーバ] AccountServiceServerのポインタ
// [引数] 
//  ctx: ??? (Context)
//  req: アカウント作成リクエストメッセージ (CreateAccountRequestのポインタ)
// [返り値]
//  アカウントリソース (Accountのポインタ)
//  エラー (error)
func (s *AccountServiceServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.Account, error) {
	log.Printf("[start]CreateAccount req: %v\n", req)

	// IDを生成
	account_id := "1"

  log.Println("[end]CreateAccount")
	return &pb.Account{ Id: account_id, Email: req.Account.Email, Password: req.Account.Password }, nil
}

// アカウント情報表示
// [レシーバ] AccountServiceServerのポインタ
// [引数] 
//  ctx: ??? (Context)
//  req: アカウント作成リクエストメッセージ (GetAccountRequestのポインタ)
// [返り値]
//  アカウントリソース (Accountのポインタ)
//  エラー (error)
func (s *AccountServiceServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.Account, error) {
	log.Printf("[start]GetAccount req: %v\n", req)

  log.Println("[end]GetAccount")
	return &pb.Account{ Id: req.Id, Email: "test@example.com", Password: "test" }, nil
}

// アカウント情報更新
// [レシーバ] AccountServiceServerのポインタ
// [引数] 
//  ctx: ??? (Context)
//  req: アカウント作成リクエストメッセージ (UpdateAccountRequestのポインタ)
// [返り値]
//  アカウントリソース (Accountのポインタ)
//  エラー (error)
func (s *AccountServiceServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.Account, error) {
	log.Printf("[start]UpdateAccount req: %v\n", req)

	// アカウント情報更新処理をここへ書く

  log.Println("[end]UpdateAccount")
	return &pb.Account{ Id: req.Account.Id, Email: req.Account.Email, Password: req.Account.Password }, nil
}

// アカウント削除
// [レシーバ] DeleteServiceServerのポインタ
// [引数] 
//  ctx: ??? (Context)
//  req: アカウント作成リクエストメッセージ (DeleteAccountRequestのポインタ)
// [返り値]
//  アカウントリソース (Accountのポインタ)
//  エラー (error)
func (s *AccountServiceServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*emptypb.Empty, error) {
	log.Printf("[start]DeleteAccount req: %v\n", req)

	// アカウント情報削除処理をここへ書く

  log.Println("[end]DeleteAccount")
	return &emptypb.Empty{}, nil
}