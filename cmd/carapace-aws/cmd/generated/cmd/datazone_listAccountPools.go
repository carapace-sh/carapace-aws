package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_listAccountPoolsCmd = &cobra.Command{
	Use:   "list-account-pools",
	Short: "Lists existing account pools.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_listAccountPoolsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_listAccountPoolsCmd).Standalone()

		datazone_listAccountPoolsCmd.Flags().String("domain-identifier", "", "The ID of the domain where exsting account pools are to be listed.")
		datazone_listAccountPoolsCmd.Flags().String("max-results", "", "The maximum number of account pools to return in a single call to ListAccountPools.")
		datazone_listAccountPoolsCmd.Flags().String("name", "", "The name of the account pool to be listed.")
		datazone_listAccountPoolsCmd.Flags().String("next-token", "", "When the number of account pools is greater than the default value for the MaxResults parameter, or if you explicitly specify a value for MaxResults that is less than the number of account pools, the response includes a pagination token named NextToken.")
		datazone_listAccountPoolsCmd.Flags().String("sort-by", "", "The sort by mechanism in which the existing account pools are to be listed.")
		datazone_listAccountPoolsCmd.Flags().String("sort-order", "", "The sort order in which the existing account pools are to be listed.")
		datazone_listAccountPoolsCmd.MarkFlagRequired("domain-identifier")
	})
	datazoneCmd.AddCommand(datazone_listAccountPoolsCmd)
}
