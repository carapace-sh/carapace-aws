package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_deleteOdbNetworkCmd = &cobra.Command{
	Use:   "delete-odb-network",
	Short: "Deletes the specified ODB network.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_deleteOdbNetworkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(odb_deleteOdbNetworkCmd).Standalone()

		odb_deleteOdbNetworkCmd.Flags().Bool("delete-associated-resources", false, "Specifies whether to delete associated OCI networking resources along with the ODB network.")
		odb_deleteOdbNetworkCmd.Flags().Bool("no-delete-associated-resources", false, "Specifies whether to delete associated OCI networking resources along with the ODB network.")
		odb_deleteOdbNetworkCmd.Flags().String("odb-network-id", "", "The unique identifier of the ODB network to delete.")
		odb_deleteOdbNetworkCmd.MarkFlagRequired("delete-associated-resources")
		odb_deleteOdbNetworkCmd.Flag("no-delete-associated-resources").Hidden = true
		odb_deleteOdbNetworkCmd.MarkFlagRequired("no-delete-associated-resources")
		odb_deleteOdbNetworkCmd.MarkFlagRequired("odb-network-id")
	})
	odbCmd.AddCommand(odb_deleteOdbNetworkCmd)
}
