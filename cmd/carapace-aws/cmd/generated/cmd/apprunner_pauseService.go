package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_pauseServiceCmd = &cobra.Command{
	Use:   "pause-service",
	Short: "Pause an active App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_pauseServiceCmd).Standalone()

	apprunner_pauseServiceCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to pause.")
	apprunner_pauseServiceCmd.MarkFlagRequired("service-arn")
	apprunnerCmd.AddCommand(apprunner_pauseServiceCmd)
}
