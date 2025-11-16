package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeReplicationsCmd = &cobra.Command{
	Use:   "describe-replications",
	Short: "Provides details on replication progress by returning status information for one or more provisioned DMS Serverless replications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeReplicationsCmd).Standalone()

	dms_describeReplicationsCmd.Flags().String("filters", "", "Filters applied to the replications.")
	dms_describeReplicationsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeReplicationsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeReplicationsCmd)
}
