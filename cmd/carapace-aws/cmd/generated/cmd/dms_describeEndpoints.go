package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEndpointsCmd = &cobra.Command{
	Use:   "describe-endpoints",
	Short: "Returns information about the endpoints for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEndpointsCmd).Standalone()

	dms_describeEndpointsCmd.Flags().String("filters", "", "Filters applied to the endpoints.")
	dms_describeEndpointsCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeEndpointsCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeEndpointsCmd)
}
