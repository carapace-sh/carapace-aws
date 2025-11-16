package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_deleteApplicationCmd).Standalone()

	appconfig_deleteApplicationCmd.Flags().String("application-id", "", "The ID of the application to delete.")
	appconfig_deleteApplicationCmd.MarkFlagRequired("application-id")
	appconfigCmd.AddCommand(appconfig_deleteApplicationCmd)
}
