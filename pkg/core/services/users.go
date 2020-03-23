package services

import (
	"acuser/pkg/models"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type UserSvc struct {
	pool *pgxpool.Pool
}

func NewUserSvc(pool *pgxpool.Pool) *UserSvc {
	if pool == nil {
		panic(errors.New("pool can't be nil")) // <- be accurate
	}
	return &UserSvc{pool: pool}
}

func (receiver *UserSvc) DbInit() (err error){
	ddls := []string{rolesDDL, usersDDL}
	for _, ddl := range ddls{
		_, err := receiver.pool.Exec(context.Background(), ddl)
		if err != nil {
			log.Printf("err, %e\n", err)
			return err
		}
	}
	return nil
}

func (receiver *UserSvc) Save(User models.User) (err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		return err
	}
	defer conn.Release()
	//log.Printf("User Pass = %s\n", User.Password)
	password, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	//log.Printf("Pass = %s\n", password)
	if err != nil {
		err = errors.New("server error")
		return
	}
	_, err = conn.Exec(context.Background(), userSaveDml, User.Name, User.Surname, User.Login, password, User.Address, User.Email, User.Phone, User.Role, User.Remove)
	if err != nil {
		log.Print("can't add to db")
		return err
	}
	return nil
}

func (receiver *UserSvc) GetUserByLogin(login string) (User models.User, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	row := conn.QueryRow(context.Background(), getUserByLoginDml, login)
	err = row.Scan(
		&User.Id,
		&User.Name,
		&User.Surname,
		&User.Login,
		&User.Password,  ///TODO without Password
		&User.Address,
		&User.Email,
		&User.Phone,
		&User.Remove,
		&User.Role,
	)
	if err != nil {
		fmt.Printf("can't read from db %e", err)
		return
	}
	return User, nil
}

func (receiver *UserSvc) RemoveUserByLogin(login string) (err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	_, err = conn.Exec(context.Background(), removeUserByLoginDml, login)
	if err != nil {
		fmt.Printf("can't update user %e", err)
		return
	}
	return nil
}

func (receiver *UserSvc) GetUserList() (users []models.User, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), getUserListDml)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		User := models.User{}
		err2 := rows.Scan(
			&User.Id,
			&User.Name,
			&User.Surname,
			&User.Login,
			&User.Password,
			&User.Address,
			&User.Email,
			&User.Phone,
			&User.Remove,
			&User.Role,
		)
		///
		if err2 != nil {
			fmt.Printf("can't read from db %e", err)
			return
		}
		//
		users = append(users, User)
	}
	if rows.Err() != nil {
		log.Printf("rows err %s", err)
		return nil, rows.Err()
	}
	return
}

func (receiver *UserSvc) SaveUser(User models.User) (dto models.UserDTO, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		return
	}
	defer conn.Release()
	password, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	_, err = conn.Exec(context.Background(), userSaveDml, User.Name, User.Surname, User.Login, password, User.Address, User.Email, User.Phone, User.Role, User.Remove)
	if err != nil {
		log.Print("can't add to db")
		return
	}
	dto.Name = User.Name
	dto.Surname = User.Surname
	dto.Login =	User.Login
	dto.Address = User.Address
	dto.Email = 	User.Email
	dto.Phone =	User.Phone
	dto.Role = User.Role
	return dto, nil
}

func (receiver *UserSvc) GetUsers() (users []models.UserDTO, err error){
	conn, err := receiver.pool.Acquire(context.Background())
	if err != nil {
		log.Printf("can't get connection %e", err)
		return
	}
	defer conn.Release()
	rows, err := conn.Query(context.Background(), getUsersDml)
	if err != nil {
		fmt.Printf("can't read user rows %e", err)
		return
	}
	defer rows.Close()
	for rows.Next(){
		User := models.UserDTO{}
		err2 := rows.Scan(
			&User.Name,
			&User.Surname,
			&User.Login,
			&User.Address,
			&User.Email,
			&User.Phone,
			&User.Role,
		)
		///
		if err2 != nil {
			fmt.Printf("can't read from db %e", err)
			return
		}
		//
		users = append(users, User)
	}
	if rows.Err() != nil {
		log.Printf("rows err %s", err)
		return nil, rows.Err()
	}
	return
}