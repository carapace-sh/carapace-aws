package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var odb_listDbNodesCmd = &cobra.Command{
	Use:   "list-db-nodes",
	Short: "Returns information about the DB nodes for the specified VM cluster.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(odb_listDbNodesCmd).Standalone()

	odb_listDbNodesCmd.Flags().String("cloud-vm-cluster-id", "", "The unique identifier of the VM cluster.")
	odb_listDbNodesCmd.Flags().String("max-results", "", "The maximum number of items to return for this request.")
	odb_listDbNodesCmd.Flags().String("next-token", "", "The token returned from a previous paginated request.")
	odb_listDbNodesCmd.MarkFlagRequired("cloud-vm-cluster-id")
	odbCmd.AddCommand(odb_listDbNodesCmd)
}
