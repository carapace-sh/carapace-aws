package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appintegrations_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes the Application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appintegrations_deleteApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appintegrations_deleteApplicationCmd).Standalone()

		appintegrations_deleteApplicationCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the Application.")
		appintegrations_deleteApplicationCmd.MarkFlagRequired("arn")
	})
	appintegrationsCmd.AddCommand(appintegrations_deleteApplicationCmd)
}
