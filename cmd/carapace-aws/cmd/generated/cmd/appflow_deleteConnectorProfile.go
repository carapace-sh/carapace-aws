package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_deleteConnectorProfileCmd = &cobra.Command{
	Use:   "delete-connector-profile",
	Short: "Enables you to delete an existing connector profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_deleteConnectorProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_deleteConnectorProfileCmd).Standalone()

		appflow_deleteConnectorProfileCmd.Flags().String("connector-profile-name", "", "The name of the connector profile.")
		appflow_deleteConnectorProfileCmd.Flags().Bool("force-delete", false, "Indicates whether Amazon AppFlow should delete the profile, even if it is currently in use in one or more flows.")
		appflow_deleteConnectorProfileCmd.Flags().Bool("no-force-delete", false, "Indicates whether Amazon AppFlow should delete the profile, even if it is currently in use in one or more flows.")
		appflow_deleteConnectorProfileCmd.MarkFlagRequired("connector-profile-name")
		appflow_deleteConnectorProfileCmd.Flag("no-force-delete").Hidden = true
	})
	appflowCmd.AddCommand(appflow_deleteConnectorProfileCmd)
}
