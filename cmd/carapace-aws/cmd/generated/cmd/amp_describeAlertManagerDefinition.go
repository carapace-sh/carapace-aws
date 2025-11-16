package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_describeAlertManagerDefinitionCmd = &cobra.Command{
	Use:   "describe-alert-manager-definition",
	Short: "Retrieves the full information about the alert manager definition for a workspace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_describeAlertManagerDefinitionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(amp_describeAlertManagerDefinitionCmd).Standalone()

		amp_describeAlertManagerDefinitionCmd.Flags().String("workspace-id", "", "The ID of the workspace to retrieve the alert manager definition from.")
		amp_describeAlertManagerDefinitionCmd.MarkFlagRequired("workspace-id")
	})
	ampCmd.AddCommand(amp_describeAlertManagerDefinitionCmd)
}
