package cmd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go-project/app/route"
	"go-project/lib/config"
)

var (
	httpCommand = &cobra.Command{
		Use:   "http",
		Short: "command http",
		Run:   httpServer,
	}
)

func init() {
	rootCmd.AddCommand(httpCommand)
}

func httpServer(cmd *cobra.Command, args []string) {
	configFile, _ := rootFlags.GetString(flagsConfig)
	fmt.Println("configFile-->", configFile)

	viperConfig := viper.New()
	viperConfig.SetConfigFile(configFile)
	viperConfig.AddConfigPath(".")
	err := viperConfig.ReadInConfig()
	config.SetViper(viperConfig)
	if err != nil {
		panic(fmt.Errorf("fatal error c file: %w", err))
	}

	g := gin.New()
	h := httpGin{
		g: g,
	}
	h.g.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
	route.Register(h.g)
	err = h.g.Run(config.Viper().GetString("app_addr"))
	if err != nil {
		panic(err)
	}
}

type httpGin struct {
	g *gin.Engine
}
