package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_deleteConnectionCmd = &cobra.Command{
	Use:   "delete-connection",
	Short: "Delete an App Runner connection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_deleteConnectionCmd).Standalone()

	apprunner_deleteConnectionCmd.Flags().String("connection-arn", "", "The Amazon Resource Name (ARN) of the App Runner connection that you want to delete.")
	apprunner_deleteConnectionCmd.MarkFlagRequired("connection-arn")
	apprunnerCmd.AddCommand(apprunner_deleteConnectionCmd)
}
