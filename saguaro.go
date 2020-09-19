package main

import (
  "log"
  "os"

  "github.com/urfave/cli/v2"
  "github.com/raulcodes/saguaro/app_creator"
  "github.com/raulcodes/saguaro/app_builder"
  "github.com/raulcodes/saguaro/app_server"
  "github.com/raulcodes/saguaro/app_watcher"
)

func main() {
  app := &cli.App{
    Commands: []*cli.Command{
      {
        Name:    "create-app",
        Aliases: []string{"ca"},
        Usage:   "create a new saguaro app",
        Action:  func(c *cli.Context) error {
          name := "my-saguaro-app"
          if c.NArg() > 0 {
            name = c.Args().Get(0)
          }
          err := app_creator.CreateApp(name)
          if err != nil {
            return err
          }
          return nil        
        },
      },
      {
        Name:    "build",
        Aliases: []string{"b"},
        Usage:   "build your app",
        Action:  func(c *cli.Context) error {
          err := app_builder.BuildApp()
          if err != nil {
            return err
          }
          return nil
        },
      },
      {
        Name:    "serve",
        Aliases: []string{"s"},
        Usage:   "serve your app",
        Action:  func(c *cli.Context) error {
          err := app_server.ServeApp()
          if err != nil {
            return err
          }
          return nil
        },
      },
      {
        Name:    "watch",
        Aliases: []string{"w"},
        Usage:   "watch your app",
        Action:  func(c *cli.Context) error {
          app_watcher.WatchDocuments()
          return nil
        },
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
