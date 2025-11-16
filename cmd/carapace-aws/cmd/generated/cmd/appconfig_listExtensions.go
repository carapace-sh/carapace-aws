package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listExtensionsCmd = &cobra.Command{
	Use:   "list-extensions",
	Short: "Lists all custom and Amazon Web Services authored AppConfig extensions in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listExtensionsCmd).Standalone()

	appconfig_listExtensionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	appconfig_listExtensionsCmd.Flags().String("name", "", "The extension name.")
	appconfig_listExtensionsCmd.Flags().String("next-token", "", "A token to start the list.")
	appconfigCmd.AddCommand(appconfig_listExtensionsCmd)
}
