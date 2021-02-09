package main

import (
	"d2-item-sorter/hotkeys"
	"d2-item-sorter/sorter"
	"d2-item-sorter/utils"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

func startServer(port int) {

	fmt.Printf("Starting d2-items-server on port :3000\n")

	saveDir := utils.GetD2SavePath()
	fmt.Printf("D2 save directory: %s\n", saveDir)

	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		fmt.Printf("%s does not exist", saveDir)
		return
	}

	saveDir, _ = filepath.Abs(saveDir)
	hotkeys.Init(keys)

	sharedStash := sorter.ParseSharedStash(saveDir)
	characters, _ := sorter.ParseCharacters(saveDir)
	fmt.Printf("%d characters found\n", len(characters))

	app := fiber.New()

	app.Post("/characters/:name/transferToShared", func(c *fiber.Ctx) error {
		groups := sorter.GetGroups("./groupConfigs/groups.yaml")
		name := c.Params("name")

		if _, ok := characters[name]; ok == true {

			transferItems(name, "_shared", groups)

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

		groups := sorter.GetGroups("./groupConfigs/groups.yaml")
		name := c.Params("name")

		if _, ok := characters[name]; ok == true {

			transferItems("_shared", name, groups)

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

		groups := sorter.GetGroups("./groupConfigs/groups.yaml")
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

		transferItems(nameFrom, nameTo, groups)

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

	app.Post("/startRunRecorder", func(c *fiber.Ctx) error {

		//finishRunner = sorter.StartRunRecorder(&sharedStash, sorter.GetGroups("./groupConfigs/groups.yaml"))
		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	app.Post("/stopRunRecorder", func(c *fiber.Ctx) error {

		finishRunner <- true

		return c.JSON(&fiber.Map{
			"success": true,
		})
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
