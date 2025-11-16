package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_getCloudExadataInfrastructureCmd = &cobra.Command{
	Use:   "get-cloud-exadata-infrastructure",
	Short: "Returns information about the specified Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_getCloudExadataInfrastructureCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_getCloudExadataInfrastructureCmd).Standalone()

		odb_getCloudExadataInfrastructureCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Exadata infrastructure.")
		odb_getCloudExadataInfrastructureCmd.MarkFlagRequired("cloud-exadata-infrastructure-id")
	})
	odbCmd.AddCommand(odb_getCloudExadataInfrastructureCmd)
}
