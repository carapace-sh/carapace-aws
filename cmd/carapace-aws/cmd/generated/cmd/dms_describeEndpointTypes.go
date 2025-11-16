package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeEndpointTypesCmd = &cobra.Command{
	Use:   "describe-endpoint-types",
	Short: "Returns information about the type of endpoints available.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeEndpointTypesCmd).Standalone()

	dms_describeEndpointTypesCmd.Flags().String("filters", "", "Filters applied to the endpoint types.")
	dms_describeEndpointTypesCmd.Flags().String("marker", "", "An optional pagination token provided by a previous request.")
	dms_describeEndpointTypesCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	dmsCmd.AddCommand(dms_describeEndpointTypesCmd)
}
