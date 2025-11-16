package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_getApplicationCmd = &cobra.Command{
	Use:   "get-application",
	Short: "Get an Application resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_getApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_getApplicationCmd).Standalone()

		appintegrations_getApplicationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Application.")
		appintegrations_getApplicationCmd.MarkFlagRequired("arn")
	})
	appintegrationsCmd.AddCommand(appintegrations_getApplicationCmd)
}
