package main

import (
	"log"
	"math/rand"
	"os"
	"regexp"
	"time"

	msql "github.com/dnevsky/hmm-bot/storage/mysql"

	"github.com/SevereCloud/vksdk/v2/api"
	"github.com/SevereCloud/vksdk/v2/api/params"
	"github.com/dnevsky/hmm-bot/pkg"
	"github.com/dnevsky/hmm-bot/storage"
	"github.com/dnevsky/hmm-bot/vkapi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	vk *api.VK
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error init configs: %s", err.Error())
	}

	godotenv.Load(".env")

	// CheckWhiteList := viper.GetBool("checkWhiteList")

	// cfg := postgr.Config{
	// 	Host:     viper.GetString("db.host"),
	// 	Port:     viper.GetString("db.port"),
	// 	Username: viper.GetString("db.username"),
	// 	DBName:   viper.GetString("db.db"),
	// 	SSLMode:  viper.GetString("db.ssl"),
	// 	Password: os.Getenv("DB_PASSWORD"),
	// }

	// db, err := postgr.NewPostgresDB(cfg)
	// if err != nil {
	// 	logrus.Fatalf("error while init db connection: %s", err.Error())
	// }

	cfg := msql.Config{
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		Net:      viper.GetString("db.net"),
		Host:     viper.GetString("db.host"),
		DBName:   viper.GetString("db.db"),
	}

	db, err := msql.NewMySQLDB(cfg)
	if err != nil {
		logrus.Fatalf("error while init db connection: %s", err.Error())
	}

	repos := storage.NewStorage(db)

	openai := pkg.NewOpenAI(os.Getenv("OPENAI_TOKEN"))

	vk, err := vkapi.NewVK(os.Getenv("VK_TOKEN"))
	if err != nil {
		logrus.Fatalf("error while init vk: %s", err.Error())
	}

	lp, err := vkapi.InitLongPool(vk)
	if err != nil {
		logrus.Fatalf("error while start longpoll: %s", err.Error())
	}

	handler := vkapi.NewHandler(vk, lp, repos, openai)
	handler.InitHandler()

	logrus.Println("Run longpoll...")
	handler.Run()

	// lp.GroupJoin(func(_ context.Context, obj events.GroupJoinObject) {
	// 	fmt.Println("groupJoin")
	// 	log.Printf("%#v", obj)
	// })

	// lp.GroupLeave(func(_ context.Context, obj events.GroupLeaveObject) {
	// 	fmt.Println("groupLeave")
	// 	log.Printf("%#v", obj)
	// })

	// lp.MessageNew(func(_ context.Context, obj events.MessageNewObject) {

	// 	var msg = "nil"
	// 	var attachment string
	// 	mention := true
	// 	if obj.Message.Action.Type == "" {
	// 		log.Printf("%d | %d: %s", obj.Message.PeerID, obj.Message.FromID, obj.Message.Text)
	// 		switch cmd := obj.Message.Text; {
	// 		case MatchString("^/test .+", cmd):
	// 			re := regexp.MustCompile(`^/test (.+)`)
	// 			res := re.FindStringSubmatch(cmd)
	// 			if len(res) == 0 {
	// 				msg = "/test [some text]"
	// 				break
	// 			}

	// 			re = regexp.MustCompile(`(\[\d+\])? {.{6}}\[AFK: \d+:\d+(:\d+)?]`)

	// 			s := re.ReplaceAllString(res[1], "")

	// 			re = regexp.MustCompile(`\[\d+\]`)

	// 			s = re.ReplaceAllString(s, "")

	// 			VKSendMessage(s, obj.Message.PeerID, mention, attachment)

	// 		case MatchString("^/sw", cmd):
	// 			if obj.Message.FromID == 197541619 {
	// 				CheckWhiteList = !CheckWhiteList
	// 			}

	// 			if CheckWhiteList {
	// 				msg = "Теперь: Включено"
	// 			} else {
	// 				msg = "Теперь: Выкючено"
	// 			}

	// 		default:
	// 			// if obj.Message.Text == "" {
	// 			// 	msg = "empty message"
	// 			// 	break
	// 			// }
	// 			// msg = obj.Message.Text
	// 		}

	// 		if msg != "nil" || attachment != "" {
	// 			VKSendMessage(msg, obj.Message.PeerID, mention, attachment)
	// 		}

	// 		res, err := repos.CheckUser(obj.Message.FromID)
	// 		if err != nil {
	// 			logrus.Printf("error white run CheckUser: %s", err.Error())
	// 		}
	// 		logrus.Println(res)

	// 	}
	// })

	// log.Println("Start longpoll")
	// if err := lp.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("default")
	return viper.ReadInConfig()
}

func VKSendMessage(msg string, peer int, mention bool, attach string) {
	p := params.NewMessagesSendBuilder()
	rand.Seed(time.Now().UnixNano())

	p.Message(msg)
	p.RandomID(rand.Int())
	p.PeerID(peer)
	p.DisableMentions(mention)
	if attach != "" {
		p.Attachment(attach)
	}

	_, err := vk.MessagesSend(p.Params)

	if err != nil {
		VKSendError(err, peer)
	}
	// return response, err
}

func VKSendError(err error, peer int) {
	VKSendMessage("Во время выполнения комманды произошла ошибка.\n"+err.Error(), peer, true, "")
	log.Println(err.Error())
}

func MatchString(pattern string, s string) (result bool) {
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func MatchStringStrong(pattern string, s string) (result bool) {
	ok, _ := regexp.MatchString(pattern, s)
	return ok
}

func CheckMaksim(check bool, id int) bool {
	if check {
		if id != 310138108 {
			return false
		}
		return true
	} else {
		return false
	}
}
