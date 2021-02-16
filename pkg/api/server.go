package api

import (
	"d2-item-sorter/pkg/hotkeys"
	"d2-item-sorter/pkg/internal"
	"d2-item-sorter/pkg/internal/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

var keys = map[int16]*hotkeys.Hotkey{
	1: {Id: 1, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'O'}, // ALT+CTRL+O
	2: {Id: 2, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'L'}, // ALT+SHIFT+M
	3: {Id: 3, Modifiers: hotkeys.ModWin + hotkeys.ModAlt, KeyCode: 'X'}, // ALT+CTRL+X
}

func StartServer(port int) {

	fmt.Printf("Starting d2-items-server on port :3000\n")

	saveDir := utils.GetD2SavePath()
	fmt.Printf("D2 save directory: %s\n", saveDir)

	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		fmt.Printf("%s does not exist", saveDir)
		return
	}

	saveDir, _ = filepath.Abs(saveDir)
	hotkeys.Init(keys)

	sharedStash := internal.ParseSharedStash(saveDir)
	characters, _ := internal.ParseCharacters(saveDir)
	fmt.Printf("%d characters found\n", len(characters))

	app := fiber.New()

	app.Static("/", "./public")

	app.Post("/characters/:name/transferToShared", func(c *fiber.Ctx) error {
		groups := GetGroups("./config/groups.yaml")
		name := c.Params("name")

		if _, ok := characters[name]; ok == true {

			TransferItems(name, "_shared", groups)

			return c.JSON(&fiber.Map{
				"success": true,
			})
		}

		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Character %s not found", name),
		})
	})

	app.Post("/shared/transfer/:name", func(c *fiber.Ctx) error {

		groups := GetGroups("./config/groups.yaml")
		name := c.Params("name")

		if _, ok := characters[name]; ok == true {

			TransferItems("_shared", name, groups)

			return c.JSON(&fiber.Map{
				"success": true,
			})
		}

		return c.Status(404).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Character %s not found", name),
		})
	})

	app.Post("/characters/:nameFrom/transfer/:nameTo", func(c *fiber.Ctx) error {

		groups := GetGroups("./config/groups.yaml")
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

		TransferItems(nameFrom, nameTo, groups)

		return c.JSON(&fiber.Map{
			"success": true,
		})

	})

	app.Get("/characters", func(c *fiber.Ctx) error {
		return c.JSON(characters)
	})

	app.Get("/characters/:name", func(c *fiber.Ctx) error {

		name := c.Params("name")

		if char, ok := characters[name]; ok == true {
			return c.JSON(char)
		}

		return c.SendStatus(404)

	})

	app.Get("/characters/:name/items", func(c *fiber.Ctx) error {

		name := c.Params("name")

		if char, ok := characters[name]; ok == true {
			return c.JSON(char.GetAllItems())
		}

		return c.SendStatus(404)

	})

	app.Get("/sharedStash", func(c *fiber.Ctx) error {
		return c.JSON(sharedStash)
	})

	app.Get("/sharedStash/items", func(c *fiber.Ctx) error {
		return c.JSON(sharedStash.GetAllItems())
	})

	app.Get("/grailStatus", func(c *fiber.Ctx) error {
		return c.JSON(GetGrailStatusList())
	})

	app.Post("/sharedStash/sort", func(c *fiber.Ctx) error {

		groups := GetGroups("./config/groups.yaml")

		SortItems("_shared", groups)

		return c.Status(200).JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Post("/startRunRecorder", func(c *fiber.Ctx) error {

		internal.StartRunRecorder(&sharedStash, GetGroups("./config/groups.yaml"))
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Post("/stopRunRecorder", func(c *fiber.Ctx) error {

		internal.StopRunRecorder()
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
