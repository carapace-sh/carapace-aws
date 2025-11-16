package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeConnectionsCmd = &cobra.Command{
	Use:   "describe-connections",
	Short: "Describes the status of the connections that have been made between the replication instance and an endpoint.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeConnectionsCmd).Standalone()

	dms_describeConnectionsCmd.Flags().String("filters", "", "The filters applied to the connection.")
	dms_describeConnectionsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeConnectionsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeConnectionsCmd)
}
