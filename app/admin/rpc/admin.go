package main

import (
	"bodhiadmin/common/utils"
	"flag"
	"fmt"
	"gorm.io/gorm"
	"time"

	"bodhiadmin/app/admin/rpc/internal/config"
	accountServer "bodhiadmin/app/admin/rpc/internal/server/account"
	menuServer "bodhiadmin/app/admin/rpc/internal/server/menu"
	nodeServer "bodhiadmin/app/admin/rpc/internal/server/node"
	nodegroupServer "bodhiadmin/app/admin/rpc/internal/server/nodegroup"
	roleServer "bodhiadmin/app/admin/rpc/internal/server/role"
	userServer "bodhiadmin/app/admin/rpc/internal/server/user"
	userroleServer "bodhiadmin/app/admin/rpc/internal/server/userrole"
	"bodhiadmin/app/admin/rpc/internal/svc"
	"bodhiadmin/app/admin/rpc/proto/admin"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const CONFFILE = "etc/admin.yaml"
const APPNAME = "admin.rpc"
const DEFAULTPORT = "9020"
const APINAME = "admin.api"
const APICONFFILE = "etc/admin-api.yaml"
const APIDEFAULTPORT = "9021"
const TAB = "  "

var (
	masterUuid, secret string
	setEtcd            string
	database           = "bodhi_admin"
	expired            = int64(24) // token expired hour
	refreshExpired     = int64(12) // refresh token expired hour
)

var c config.Config
var configFile *string

func main() {
	//check system status
	err := CheckStatus()
	fmt.Sprintf("check status ", err)
	if err != nil {
		panic(any(fmt.Sprintf("Check Status error: %v", err)))
	}

	if configFile == nil {
		configFile = flag.String("f", CONFFILE, "the config file")
		flag.Parse()
		conf.MustLoad(*configFile, &c)
	}
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		admin.RegisterAccountServer(grpcServer, accountServer.NewAccountServer(ctx))
		admin.RegisterMenuServer(grpcServer, menuServer.NewMenuServer(ctx))
		admin.RegisterNodeServer(grpcServer, nodeServer.NewNodeServer(ctx))
		admin.RegisterNodeGroupServer(grpcServer, nodegroupServer.NewNodeGroupServer(ctx))
		admin.RegisterRoleServer(grpcServer, roleServer.NewRoleServer(ctx))
		admin.RegisterUserServer(grpcServer, userServer.NewUserServer(ctx))
		admin.RegisterUserRoleServer(grpcServer, userroleServer.NewUserRoleServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func CheckStatus() error {
	status := utils.CheckLock()
	if status {
		// installed
		return nil
	}

	// do install
	//1. init conf
	//2. install database
	//3. create api conf
	reset, err := utils.CheckConf(CONFFILE)
	//fmt.Println("check conf finished ", reset)
	if err != nil {
		return err
	}
	//skip install
	if reset != "y" {
		fmt.Sprintf("Skip init conf")
		return nil
	}
	err = InitRpcConf()
	if err != nil {
		fmt.Sprintf("Init rpc conf error!")
		panic(any(err))
	}

	err = InstallDatabase()
	if err != nil {
		fmt.Sprintf("Install database error!")
		panic(any(err))
	}

	err = InitApiConf()
	if err != nil {
		fmt.Sprintf("Init api conf error!")
		return err
	}

	err = utils.CreatLock()
	if err != nil {
		return err
	}
	fmt.Println("install finished")
	time.Sleep(3 * time.Second)
	return nil
}

func InitApiConf() error {
	fmt.Printf("create default api config? default yes(y/n)")
	setApiConf := utils.ScanInput("y")
	if setApiConf == "y" {
		reset, err := utils.CheckConf(APICONFFILE)
		if err != nil {
			return err
		}
		if reset == "y" {
			base := "Name: " + APINAME + "\n" +
				"Host: 0.0.0.0" + "\n" +
				"Port: " + APIDEFAULTPORT + "\n\n"
			err := utils.WriteFile(APICONFFILE, base)
			if err != nil {
				return err
			}
			err = SetRpcConf()
			if err != nil {
				return err
			}
			err = SetAuth(APICONFFILE)
			if err != nil {
				return err
			}
			_ = SetLog(APICONFFILE)
		}
	}
	fmt.Println("create api config is ok")
	return nil
}

func SetRpcConf() error {
	rpcConf := "AdminRpcConf: \n"
	if setEtcd == "y" {
		rpcConf += TAB + "Etcd: \n" +
			TAB + TAB + "Hosts: \n" +
			TAB + TAB + TAB + "- " + c.Etcd.Hosts[0] + "\n" +
			TAB + TAB + "Key: " + APPNAME
	} else {
		rpcConf += TAB + "Endpoints: \n" +
			TAB + TAB + "- 127.0.0.1:" + DEFAULTPORT + "\n" +
			TAB + "NonBlock: true \n\n"
	}
	err := utils.WriteFile(APICONFFILE, rpcConf)
	if err != nil {
		return err
	}
	return nil
}

// InstallDatabase check database & tables, install database
func InstallDatabase() error {
	fmt.Println("Start install database...")

	configFile = flag.String("f", CONFFILE, "the config file")
	flag.Parse()
	conf.MustLoad(*configFile, &c)

	err := CreateDatabase()
	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)

	db, err := utils.InitMySQL(c.MySql.Host, c.MySql.User, c.MySql.Password, c.MySql.Database, c.MySql.Port)
	if err != nil {
		return err
	}

	install, err := CheckTables(db)
	if err != nil {
		return err
	}
	if install == "y" {
		err = InstallTables(db)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
		err = CreateMaster(db)
		if err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}

	fmt.Println("install database is ok...")
	return nil
}

func CreateMaster(db *gorm.DB) error {
	fmt.Println("Start create master...")
	var username, password, repassword, email string
	for {
		fmt.Printf("Please enter username (4-32):")
		username = utils.ScanInput("admin")
		if len(username) == 0 || !utils.ValidatorUserName(username, 4, 32) {
			fmt.Println("Invalid username, please reenter")
			continue
		}
		break
	}
	for {
		fmt.Printf("Please enter email:")
		email = utils.ScanInput("")
		if len(email) == 0 || !utils.ValidatorEmail(email) {
			fmt.Println("Invalid email, please reenter")
			continue
		}
		break
	}
	for {
		fmt.Printf("Please enter password (8-16,[a-zA-Z],[0-9],[.@$!_], choose more than 2 options):")
		password = utils.ScanInput("")
		if len(password) == 0 || !utils.ValidatorPassword(password, 8, 16) {
			fmt.Println("Invalid password, please reenter")
			continue
		}
		break
	}
	for {
		fmt.Printf("Please enter password again:")
		repassword = utils.ScanInput("")
		if len(repassword) == 0 || password != repassword {
			fmt.Println("repassword is inconsistent with password, please reenter")
			continue
		}
		break
	}
	pwd := utils.GetHashedPassword(password, c.AdminConf.Salt)
	masterSql := fmt.Sprintf("INSERT INTO `user` (`user_uuid`, `username`, `password`, `email`, `name`, `avatar`, `remark`, `last_login_ip`, `last_login_time`, `last_active_ip`, `last_active_time`, `custom_data`, `mail_verified`, `is_deleted`, `is_enabled`, `created_at`) VALUES ('%s', '%s', '%s', '%s', '%s', '', '', '', 0, '', 0, '', 0, 0, 1, '%d');", c.AdminConf.Master, username, pwd, email, username, utils.GetTimestamp())
	err := db.Exec(masterSql).Error
	if err != nil {
		return err
	}

	fmt.Println("create master is ok...")
	return nil
}

func InstallTables(db *gorm.DB) error {
	fmt.Println("Start install database...")
	installFile := "etc/install/install.sql"
	err := utils.RunSql(db, installFile)
	if err != nil {
		return err
	}
	fmt.Println("database installed")
	return nil
}

func CheckTables(db *gorm.DB) (string, error) {
	fmt.Println("Check Tables...")
	var reinstall string
	var tables []string
	err := db.Raw("show tables").Scan(&tables).Error
	if err != nil {
		return reinstall, err
	}
	if len(tables) == 0 {
		reinstall = "y"
	} else {
		fmt.Printf("database is not null, clear tables and install? default no (y/n)")
		reinstall = utils.ScanInput("n")
		if reinstall != "y" {
			fmt.Println("skip install tables")
		} else {
			utils.ClearTables(db, tables)
		}
	}
	return reinstall, nil
}

func CreateDatabase() error {
	fmt.Println("Check MySQL conf...")
	db, err := utils.CheckMySQLConf(c.MySql.Host, c.MySql.User, c.MySql.Password, c.MySql.Port)
	if err != nil {
		return err
	}
	fmt.Println("MySQL connect is ok...")
	defer db.Close()
	// create database
	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", c.MySql.Database))
	if err != nil {
		return err
	}

	fmt.Println("Create database is ok...")
	return nil
}

// InitRpcConf init rpc conf
func InitRpcConf() error {
	fmt.Println("Start init rpc conf...")
	var err error
	var steps = []string{
		"base",
		"etcd",
		"admin",
		"auth",
		"database",
		"redis",
		"log",
	}
	i := 1
	for _, step := range steps {
		fmt.Println(i, ". Init", step, "conf...")
		switch step {
		case "base":
			err = SetBase()
			break
		case "etcd":
			err = SetEtcd()
			break
		case "admin":
			err = SetAdminConf()
			break
		case "auth":
			err = SetAuth(CONFFILE)
			break
		case "database":
			err = SetDatabase()
			break
		case "redis":
			err = SetRedis()
			break
		case "log":
			err = SetLog(CONFFILE)
			break
		default:
			break
		}
		if err != nil {
			fmt.Sprintf("set %s conf err: %v", step, err)
		}
		time.Sleep(1 * time.Second)
		i++
	}
	return nil
}

func SetLog(file string) error {
	var conf = "Log:\n" +
		TAB + "Mode: file\n" +
		TAB + "Encoding: plain\n" +
		TAB + "Path: logs\n" +
		TAB + "Level: info # set error in product env\n" +
		TAB + "KeepDays: 1\n\n"
	err := utils.WriteFile(file, conf)
	if err != nil {
		return err
	}
	fmt.Sprintf("set Redis conf is ok")
	return nil
}

func SetRedis() error {
	var setRedis, host, port, password, index string
	fmt.Printf("set Redis conf？default no (y/n)")
	setRedis = utils.ScanInput("n")
	if setRedis == "y" {
		for {
			fmt.Printf("Please enter Redis host (default 127.0.0.1):")
			host = utils.ScanInput("127.0.0.1")
			if len(host) == 0 || !utils.ValidatorIp(host) {
				fmt.Println("Invalid host, please reenter")
				continue
			}
			break
		}
		for {
			fmt.Printf("Please enter Redis port (default 6379):")
			port = utils.ScanInput("6379")
			if len(port) == 0 || !utils.ValidatorPort(utils.StrToInt64(port)) {
				fmt.Println("Invalid port, please reenter")
				continue
			}
			break
		}
		fmt.Printf("Please enter Redis database index (default 1):")
		index = utils.ScanInput("1")
		fmt.Printf("Please enter Redis password:")
		password = utils.ScanInput("")

		var conf = "RedisConf:\n" +
			TAB + "Host: " + host + "\n" +
			TAB + "Port: " + port + "\n" +
			TAB + "DBIndex: " + index + "\n" +
			TAB + "Password: " + password + "\n\n"
		err := utils.WriteFile(CONFFILE, conf)
		if err != nil {
			return err
		}
		fmt.Sprintf("set Redis conf is ok")
		return nil
	} else {
		fmt.Println("skip set redis conf...")
	}
	return nil
}

func SetDatabase() error {
	var host, port, user, password string
	for {
		fmt.Printf("Please enter MySQL host (default 127.0.0.1):")
		host = utils.ScanInput("127.0.0.1")
		if len(host) == 0 || !utils.ValidatorIp(host) {
			fmt.Println("Invalid host, please reenter")
			continue
		}
		break
	}
	for {
		fmt.Printf("Please enter MySQL port (default 3306):")
		port = utils.ScanInput("3306")
		if len(port) == 0 || !utils.ValidatorPort(utils.StrToInt64(port)) {
			fmt.Println("Invalid port, please reenter")
			continue
		}
		break
	}
	for {
		fmt.Printf("Please enter MySQL user:")
		user = utils.ScanInput("root")
		if len(user) == 0 {
			fmt.Println("Invalid user, please reenter")
			continue
		}
		break
	}
	fmt.Printf("Please enter MySQL password:")
	password = utils.ScanInput("")

	var conf = "Mysql:\n" +
		TAB + "Database: " + database + "\n" +
		TAB + "Host: " + host + "\n" +
		TAB + "Port: " + port + "\n" +
		TAB + "User: " + user + "\n" +
		TAB + "Password: '" + password + "'\n\n"
	err := utils.WriteFile(CONFFILE, conf)
	if err != nil {
		return err
	}
	fmt.Sprintf("set MySQL conf is ok")
	return nil
}

func SetAuth(file string) error {
	authKey := "Auth"
	if secret == "" {
		secret = utils.RandomStr(32)
		authKey = "AuthConf"
	}
	expiredTime := utils.Int64ToStr(3600 * expired)
	var conf = authKey + ":\n" +
		TAB + "AccessSecret: " + secret + "\n" +
		TAB + "AccessExpired: " + expiredTime + "\n\n"
	err := utils.WriteFile(file, conf)
	if err != nil {
		return err
	}
	fmt.Sprintf("set auth conf is ok")
	return nil
}

func SetAdminConf() error {
	salt := utils.RandomStr(16)
	masterUuid = utils.CreateUuid()
	expiredTime := utils.Int64ToStr(3600 * refreshExpired)
	var conf = "AdminConf:\n" +
		TAB + "Salt: " + salt + "\n" +
		TAB + "Master: " + masterUuid + "\n" +
		TAB + "RefreshExpired: " + expiredTime + "\n\n"
	err := utils.WriteFile(CONFFILE, conf)
	if err != nil {
		return err
	}
	fmt.Sprintf("set refresh token expired %d hour(%s s)", refreshExpired, expiredTime)
	return nil
}

func SetEtcd() error {
	var host, port string
	fmt.Printf("set etcd service？default no (y/n)")
	setEtcd = utils.ScanInput("n")
	if setEtcd == "y" {
		for {
			fmt.Printf("Please enter ETCD host:")
			host = utils.ScanInput("")
			if len(host) == 0 || !utils.ValidatorIp(host) {
				fmt.Println("Invalid host, please reenter")
				continue
			}
			break
		}
		for {
			fmt.Printf("Please enter ETCD port:")
			port = utils.ScanInput("")
			if len(port) == 0 || !utils.ValidatorPort(utils.StrToInt64(port)) {
				fmt.Println("Invalid port, please reenter")
				continue
			}
			break
		}
		var etcdHost = host + ":" + port
		var conf = "Etcd:\n" +
			TAB + "Hosts: \n" +
			TAB + TAB + "- " + etcdHost + "\n" +
			TAB + "Key: " + APPNAME + "\n\n"
		err := utils.WriteFile(CONFFILE, conf)
		if err != nil {
			return err
		}
		fmt.Sprintf("set etcd service host: %s, key: %s", etcdHost, APPNAME)
	} else {
		fmt.Println("skip set etcd service...")
	}
	return nil
}

func SetBase() error {
	base := "Name: " + APPNAME + "\n" +
		"ListenOn: 0.0.0.0:" + DEFAULTPORT + "\n\n"
	err := utils.WriteFile(CONFFILE, base)
	if err != nil {
		return err
	}

	fmt.Println("set base conf is ok!")
	return nil
}
