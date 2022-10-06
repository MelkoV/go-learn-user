package cmd

import (
	"fmt"
	"github.com/MelkoV/go-learn-common/app"
	"github.com/MelkoV/go-learn-logger/logger"
	"github.com/MelkoV/go-learn-user/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// apiCmd represents the api command
var (
	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "Run API gRPC server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			port := viper.GetInt("api.port")
			l := logger.NewCategoryLogger("user/api", app.SYSTEM_UUID, logger.NewStreamLog())
			l.Info("starting API server on port %d", port)
			dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s",
				viper.GetString("db.host"),
				viper.GetString("db.user"),
				viper.GetString("db.password"),
				viper.GetString("db.name"),
				viper.GetString("db.port"),
				viper.GetString("db.timeZone"),
			)
			api.Serve(port, l, dsn)
		},
	}
)

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
