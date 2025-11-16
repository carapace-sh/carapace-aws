package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_createApplicationCmd).Standalone()

	appconfig_createApplicationCmd.Flags().String("description", "", "A description of the application.")
	appconfig_createApplicationCmd.Flags().String("name", "", "A name for the application.")
	appconfig_createApplicationCmd.Flags().String("tags", "", "Metadata to assign to the application.")
	appconfig_createApplicationCmd.MarkFlagRequired("name")
	appconfigCmd.AddCommand(appconfig_createApplicationCmd)
}
