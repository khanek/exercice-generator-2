package main

import (
	"fmt"
	"khanek/exercise-generator/api"
	"khanek/exercise-generator/config/bindatafs"
	"khanek/exercise-generator/core/admin"
	"khanek/exercise-generator/core/database"
	"khanek/exercise-generator/words"
	"log"
	"net/http"
	"os"

	// "github.com/gin-contrib/pprof"  // profiling
	"github.com/gin-gonic/gin"

	"golang.org/x/sync/errgroup"

	_ "net/http/pprof"
)

var (
	g errgroup.Group
)

func runServer() {
	db := database.Initialize()
	defer db.Close()
	admin.Initialize(db)

	ginEngine := gin.Default()
	// pprof.Register(ginEngine)  // profiling

	initializeAPI(ginEngine)
	initializeAdmin(ginEngine)

	ginEngine.Run(":9000")
}

func compileStatic() {
	assetFS := bindatafs.AssetFS.NameSpace("admin")
	// Register view paths into AssetFS
	assetFS.RegisterPath("public/admin")
	// Compile templates under registered view paths into binary
	if err := assetFS.Compile(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cmd := "runServer"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}
	switch cmd {
	case "runServer":
		runServer()
	case "compileStatic":
		compileStatic()
	default:
		fmt.Println("expected 'runServer' or 'compileStatic' subcommands")
		os.Exit(1)
	}
}

func initializeAPI(engine *gin.Engine) error {
	api.AddUrls(engine)
	return nil
}

func initializeAdmin(engine *gin.Engine) error {
	handler := admin.GetAdmin()
	admin.Register(&words.Tag{})
	admin.Register(&words.Word{})
	mux := http.NewServeMux()
	handler.MountTo("/admin", mux)
	engine.Any("/admin/*resources", gin.WrapH(mux))
	return nil
}
