package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeOrderableReplicationInstancesCmd = &cobra.Command{
	Use:   "describe-orderable-replication-instances",
	Short: "Returns information about the replication instance types that can be created in the specified region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeOrderableReplicationInstancesCmd).Standalone()

	dms_describeOrderableReplicationInstancesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeOrderableReplicationInstancesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeOrderableReplicationInstancesCmd)
}
