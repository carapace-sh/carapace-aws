package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_describeDataProvidersCmd = &cobra.Command{
	Use:   "describe-data-providers",
	Short: "Returns a paginated list of data providers for your account in the current region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_describeDataProvidersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_describeDataProvidersCmd).Standalone()

		dms_describeDataProvidersCmd.Flags().String("filters", "", "Filters applied to the data providers described in the form of key-value pairs.")
		dms_describeDataProvidersCmd.Flags().String("marker", "", "Specifies the unique pagination token that makes it possible to display the next page of results.")
		dms_describeDataProvidersCmd.Flags().String("max-records", "", "The maximum number of records to include in the response.")
	})
	dmsCmd.AddCommand(dms_describeDataProvidersCmd)
}
