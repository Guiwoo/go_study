package main

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"math/rand"
	"semina_entgo/custom"
	"semina_entgo/ent"
	"semina_entgo/ent/car"
	"semina_entgo/ent/group"
	"semina_entgo/ent/tester"
	"semina_entgo/ent/user"
	"time"
)

/**
스키마를 코드로 관리한다.
그래프 탐색이 쉽게 가능  - 쿼리실행, 집계, 그래프구조
정적타입 그리고 명시적인 API
*/

func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	fruit := []string{"apple", "banana", "citron", "durian"}

	u, err := client.User.
		Create().
		SetAge(30).
		SetName(fruit[rand.Int()%len(fruit)] + " guiwoo").
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w\n", err)
	}

	log.Printf("user was created %+v\n", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) ([]*ent.User, error) {
	list, err := client.User.Query().
		Where(user.Name("banana guiwoo")).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to create %+v", err)
	}
	log.Println("car was created ", tesla)

	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("fail to create %+v", err)
	}

	log.Println("car was created ", ford)

	foundUser, err := client.User.Query().
		Where(user.ID(1)).Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("fail to create %+v", err)
	}

	rs, err := foundUser.Update().
		AddCars(tesla, ford).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("fail to update user %+v", err)
	}

	return rs, nil
}

func QueryCars(ctx context.Context, client *ent.Client) error {
	usr, err := client.User.Query().Where(user.ID(1)).Only(ctx)
	if err != nil {
		return err
	}

	cars, err := usr.QueryCars().All(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cars)

	car, err := usr.QueryCars().Where(car.Model("Tesla")).Only(ctx)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", car)

	return nil
}

func QueryCarUsers(ctx context.Context, clinet *ent.Client) error {
	usr, err := clinet.User.Query().Where(user.ID(1)).Only(ctx)
	if err != nil {
		return err
	}

	cars, err := usr.QueryCars().All(ctx)
	if err != nil {
		return err
	}

	for _, v := range cars {
		owner, err := v.QueryOwner().Only(ctx)
		if err != nil {
			log.Printf("fail to get owner %+v\n", v)
			return err
		}
		log.Printf("car %+v owner %+v\n", v, owner)
	}

	return nil
}

func CreateGraph(ctx context.Context, client *ent.Client) error {
	guiwoo, err := client.User.Create().
		SetAge(30).
		SetName("Guiwoo").
		Save(ctx)

	if err != nil {
		log.Printf("fail to create user %+v", err)
		return err
	}

	park, err := client.User.Create().
		SetAge(21).
		SetName("Park").
		Save(ctx)

	if err != nil {
		log.Printf("fail to create user %+v", err)
		return err
	}

	if err := client.Car.Create().
		SetModel("Hyundai").
		SetRegisteredAt(time.Now()).
		SetOwner(guiwoo).
		Exec(ctx); err != nil {
		log.Printf("fail to create car %+v", err)
		return err
	}

	if err := client.Car.Create().
		SetModel("Kia").
		SetRegisteredAt(time.Now().AddDate(-20, -2, -1)).
		SetOwner(park).
		Exec(ctx); err != nil {
		log.Printf("fail to create car %+v", err)
		return err
	}

	if err := client.Group.Create().
		SetName("Korea").
		AddUsers(guiwoo).
		Exec(ctx); err != nil {
		fmt.Printf("fail to create group %+v", err)
		return err
	}

	if err := client.Group.Create().
		SetName("Japan").
		AddUsers(park).
		Exec(ctx); err != nil {
		fmt.Printf("fail to create group %+v", err)
		return err
	}

	return nil
}

func QueryKorea(ctx context.Context, client *ent.Client) error {
	if cars, err := client.Debug().Group.Query().
		Where(group.Name("Korea")).
		QueryUsers().QueryCars().All(ctx); err != nil {
		return fmt.Errorf("fail to query group users cars %+v", err)
	} else {
		fmt.Printf("%+v\n", cars)
		return nil
	}
}

func QueryParkCars(ctx context.Context, client *ent.Client) error {
	usr := client.User.Query().
		Where(user.HasCars(), user.Name("Guiwoo")).
		OnlyX(ctx)

	cars, err := usr.
		QueryGroups().
		QueryUsers().
		QueryCars().
		Where(car.Not(car.Model("Kia"))).All(ctx)
	if err != nil {
		return err
	}
	for _, v := range cars {
		fmt.Printf("%+v\n", v)
	}
	return nil
}

func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
	group, err := client.Group.Query().
		Where(group.HasUsers()).
		All(ctx)
	if err != nil {
		return err
	}
	for _, v := range group {
		rs, err := v.QueryUsers().All(ctx)
		if err != nil {
			log.Printf("fail to get User %+v\n", v)
			return err
		}
		fmt.Printf("%+v\n", rs)
	}
	fmt.Printf("%+v", group)
	return nil
}

func TestInsertValidateReturnError(ctx context.Context, client *ent.Client) error {
	tsr, err := client.Tester.Create().
		SetPascalCase("work?").
		SetLetMeCheck("let me check ?").
		SetSize(tester.SizeBig).
		SetShape(custom.Rectangle).
		SetLevel(custom.High).
		Save(ctx)
	if err != nil {
		return err
	}
	fmt.Println(tsr)
	return nil
}

func CreateCardAndUser(ctx context.Context, client *ent.Client) error {
	guiwoo := client.User.Query().Where(user.ID(1)).OnlyX(ctx)

	card, err := client.Card.Create().
		SetNumber("1234-1234-1234-1234").
		SetOwner(guiwoo).
		SetExpiredAt(time.Now().AddDate(5, 0, 0)).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Printf("card created %+v\n", card)
	return nil
}

func main() {
	client, err := ent.Open("mysql", "guiwoo:guiwoo@tcp(localhost:3306)/guiwoo?parseTime=true")
	if err != nil {
		panic(err)
	}

	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	//auto migrate
	if err := client.Schema.Create(context.TODO()); err != nil {
		log.Fatalf("failed to create entity %+v", err)
	}

	if err := CreateCardAndUser(context.Background(), client); err != nil {
		log.Fatalf("failed to query korea %+v", err)
	}
}
