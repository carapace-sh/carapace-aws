package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteWebAppCmd = &cobra.Command{
	Use:   "delete-web-app",
	Short: "Deletes the specified web app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteWebAppCmd).Standalone()

	transfer_deleteWebAppCmd.Flags().String("web-app-id", "", "Provide the unique identifier for the web app that you are deleting.")
	transfer_deleteWebAppCmd.MarkFlagRequired("web-app-id")
	transferCmd.AddCommand(transfer_deleteWebAppCmd)
}
