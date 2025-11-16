package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteCapacityManagerDataExportCmd = &cobra.Command{
	Use:   "delete-capacity-manager-data-export",
	Short: "Deletes an existing Capacity Manager data export configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteCapacityManagerDataExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteCapacityManagerDataExportCmd).Standalone()

		ec2_deleteCapacityManagerDataExportCmd.Flags().String("capacity-manager-data-export-id", "", "The unique identifier of the data export configuration to delete.")
		ec2_deleteCapacityManagerDataExportCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCapacityManagerDataExportCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteCapacityManagerDataExportCmd.MarkFlagRequired("capacity-manager-data-export-id")
		ec2_deleteCapacityManagerDataExportCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_deleteCapacityManagerDataExportCmd)
}
