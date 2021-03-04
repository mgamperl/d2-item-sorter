package server

import (
	"bufio"
	"d2-item-sorter/pkg/api"
	"d2-item-sorter/pkg/hotkeys"
	"d2-item-sorter/pkg/internal/reader"
	"d2-item-sorter/pkg/internal/runRecorder"
	"d2-item-sorter/pkg/internal/utils"
	"embed"
	_ "embed" // using embedded ressources in this package
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed ui/build/* ui/build/static/*
var embeddedFiles embed.FS

var runRecorderChannel = make(chan bool)

func getFileSystem(useOS bool) http.FileSystem {
	/*if useOS {
		log.Print("using live mode")
		return http.FS(os.DirFS("static"))
	}*/

	fsys, err := fs.Sub(embeddedFiles, "ui/build")
	if err != nil {
		panic(err)
	}

	return http.FS(fsys)
}

var keys = map[int16]*hotkeys.Hotkey{
	1: {Id: 1, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'O'}, // ALT+CTRL+O
	2: {Id: 2, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'L'}, // ALT+SHIFT+M
	3: {Id: 3, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'X'}, // ALT+CTRL+X
}

func StartServer(port int) {

	fmt.Printf("Starting d2-grailer on port %d\n", port)

	saveDir := utils.GetD2SavePath()
	fmt.Printf("D2 save directory: %s\n", saveDir)

	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		fmt.Printf("%s does not exist", saveDir)
		return
	}

	saveDir, _ = filepath.Abs(saveDir)
	hotkeys.Init(keys)

	sharedStash := reader.ReadSharedStash(saveDir)
	characters, _ := reader.ReadAllCharactersFromPath(saveDir)
	fmt.Printf("%d characters found\n", len(characters))

	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	apiGroup := app.Group("/api")

	apiGroup.Post("/characters/:name/transferToShared", func(c *fiber.Ctx) error {
		groups := api.GetGroups("./config/groups.yaml")
		name := c.Params("name")

		if _, ok := characters[name]; ok == true {

			api.TransferItems(name, "_shared", groups)

			return c.JSON(&fiber.Map{
				"success": true,
			})
		}

		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Character %s not found", name),
		})
	})

	apiGroup.Post("/shared/transfer/:name", func(c *fiber.Ctx) error {

		groups := api.GetGroups("./config/groups.yaml")
		name := c.Params("name")

		if _, ok := characters[name]; ok == true {

			api.TransferItems("_shared", name, groups)

			return c.JSON(&fiber.Map{
				"success": true,
			})
		}

		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Character %s not found", name),
		})
	})

	apiGroup.Post("/characters/:nameFrom/transfer/:nameTo", func(c *fiber.Ctx) error {

		groups := api.GetGroups("./config/groups.yaml")
		nameFrom := c.Params("nameFrom")
		nameTo := c.Params("nameTo")

		if _, ok := characters[nameFrom]; ok == false {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Character %s not found", nameFrom),
			})
		}

		if _, ok := characters[nameTo]; ok == false {
			return c.Status(404).JSON(&fiber.Map{
				"success": false,
				"error":   fmt.Sprintf("Character %s not found", nameTo),
			})
		}

		api.TransferItems(nameFrom, nameTo, groups)

		return c.JSON(&fiber.Map{
			"success": true,
		})

	})

	apiGroup.Get("/characters", func(c *fiber.Ctx) error {
		return c.JSON(characters)
	})

	apiGroup.Get("/characters/:name", func(c *fiber.Ctx) error {

		name := c.Params("name")

		if char, ok := characters[name]; ok == true {
			return c.JSON(char)
		}

		return c.SendStatus(404)

	})

	apiGroup.Get("/characters/:name/items", func(c *fiber.Ctx) error {

		name := c.Params("name")

		if char, ok := characters[name]; ok == true {
			return c.JSON(char.Items)
		}

		return c.SendStatus(404)

	})

	apiGroup.Get("/sharedStash", func(c *fiber.Ctx) error {
		return c.JSON(sharedStash)
	})

	apiGroup.Get("/sharedStash/items", func(c *fiber.Ctx) error {
		return c.JSON(sharedStash.GetAllItems())
	})

	apiGroup.Get("/grailStatus", func(c *fiber.Ctx) error {
		return c.JSON(api.GetGrailStatusList())
	})

	apiGroup.Post("/sharedStash/sort", func(c *fiber.Ctx) error {

		groups := api.GetGroups("./config/groups.yaml")

		api.SortItems("_shared", groups)

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
		})
	})

	apiGroup.Post("/startRunRecorder", func(c *fiber.Ctx) error {

		runRecorder.StartRunRecorder(&sharedStash, api.GetGroups("./config/groups.yaml"))
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	apiGroup.Post("/stopRunRecorder", func(c *fiber.Ctx) error {

		runRecorder.StopRunRecorder()
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Use("/", filesystem.New(filesystem.Config{
		Root:   getFileSystem(false),
		Browse: false,
		Index:  "index.html",
	}))

	app.Get("/*", func(ctx *fiber.Ctx) error {
		f, _ := getFileSystem(false).Open("index.html")
		defer f.Close()
		reader := bufio.NewReader(f)
		content, _ := ioutil.ReadAll(reader)

		ctx.Response().Header.Add("Content-Type", "text/html")

		return ctx.Send(content)
	})

	app.Listen(fmt.Sprintf(":%d", port))
}
