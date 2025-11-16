package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listCloudVmClustersCmd = &cobra.Command{
	Use:   "list-cloud-vm-clusters",
	Short: "Returns information about the VM clusters owned by your Amazon Web Services account or only the ones on the specified Exadata infrastructure.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listCloudVmClustersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_listCloudVmClustersCmd).Standalone()

		odb_listCloudVmClustersCmd.Flags().String("cloud-exadata-infrastructure-id", "", "The unique identifier of the Oracle Exadata infrastructure.")
		odb_listCloudVmClustersCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
		odb_listCloudVmClustersCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	})
	odbCmd.AddCommand(odb_listCloudVmClustersCmd)
}
