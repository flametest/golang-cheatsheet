package main

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"

	"github.com/flametest/test/ent"
	"github.com/flametest/test/ent/user"
	"github.com/flametest/test/types"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nickdu2009/copier"
)

type TestZero struct {
	A int32
}

// func isZhiShu(n int) bool {
// 	if n <= 1 {
// 		return false
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

// func getZhiShu() {
// 	var list []int
// 	for i := 0; i <= 100; i++ {
// 		res := isZhiShu(i)
// 		if res {
// 			list = append(list, i)
// 		}
// 	}

// 	fmt.Println(list)
// }

type UserA struct {
	Id      int64               `json:"id,omitempty"`
	Name    string              `json:"name,omitempty"`
	Meta    *A                  `json:"meta,omitempty"`
	Money   types.EntityDecimal `json:"money"`
	Age     int64               `json:age`
	Address string              `json:address`
}

type A struct {
	Xxx []string `json:xxx,omitempty`
	Yyy string   `json:yyy,omitempty`
	Zzz types.EntityDecimal `json:"zzz"`
}

func (p *UserA) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	srcBytes, ok := src.([]byte)
	if !ok {
		return errors.New("only support []byte type")
	}
	return json.Unmarshal(srcBytes, p)
}

func (p *UserA) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func (p *A) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	srcBytes, ok := src.([]byte)
	if !ok {
		return errors.New("only support []byte type")
	}
	return json.Unmarshal(srcBytes, p)
}

func (p *A) Value() (driver.Value, error) {
	if p == nil {
		return nil, nil
	}
	return json.Marshal(p)
}

func main() {
	// client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	// if err != nil {
	// 	log.Fatalf("failed opening connection to sqlite: %v", err)
	// }
	// defer client.Close()
	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)

	// }
	// a := list.New()
	// a.PushBack(1)
	// fmt.Println(a)
	// fmt.Println(time.Unix(121424, 0))

	// x := [3]int{1, 2, 3}

	// func(arr [3]int) {
	// 	arr[0] = 7
	// 	fmt.Println(arr)
	// }(x)

	// fmt.Println(x)

	// testZero := TestZero{}
	// fmt.Println(testZero)
	// y := []string{"x", "y"}
	// fmt.Println(reflect.TypeOf(x))
	// fmt.Println(reflect.TypeOf(y))
	// getZhiShu()
	x := &A{}
	fmt.Println(x.Xxx, x.Yyy)
	client, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/test?parseTime=true")
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	tx, err := client.Tx(ctx)

	u, _ := createUser(ctx, tx)
	fmt.Println(u)

	if err := tx.Commit(); err != nil {
		fmt.Println(err)
	}
	u, err = client.User.Query().Where(user.ID(1)).Only(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	user1 := &UserA{Age: 12}
	copier.Copy(user1, u)
	fmt.Println(user1.Money.Decimal())
	fmt.Println(user1.Age)
	fmt.Println(user1.Address)

	user1.Money = types.NewEntityDecimalFromInt64(1)
	u1 := &ent.User{}
	copier.Copy(u1, user1)
	fmt.Println(u1)

	uu, err := client.User.Query().Where(user.Name("jj")).First(ctx)
	fmt.Println("first", uu)
}

func createUser(ctx context.Context, tx *ent.Tx) (*ent.User, error) {
	b, _ := json.Marshal(A{Zzz: types.NewEntityDecimalFromInt64(1)})
	u, err := tx.User.Create().SetName("jj").SetMeta(b).Save(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return u, nil
}
