package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pi_getDimensionKeyDetailsCmd = &cobra.Command{
	Use:   "get-dimension-key-details",
	Short: "Get the attributes of the specified dimension group for a DB instance or data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pi_getDimensionKeyDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pi_getDimensionKeyDetailsCmd).Standalone()

		pi_getDimensionKeyDetailsCmd.Flags().String("group", "", "The name of the dimension group.")
		pi_getDimensionKeyDetailsCmd.Flags().String("group-identifier", "", "The ID of the dimension group from which to retrieve dimension details.")
		pi_getDimensionKeyDetailsCmd.Flags().String("identifier", "", "The ID for a data source from which to gather dimension data.")
		pi_getDimensionKeyDetailsCmd.Flags().String("requested-dimensions", "", "A list of dimensions to retrieve the detail data for within the given dimension group.")
		pi_getDimensionKeyDetailsCmd.Flags().String("service-type", "", "The Amazon Web Services service for which Performance Insights returns data.")
		pi_getDimensionKeyDetailsCmd.MarkFlagRequired("group")
		pi_getDimensionKeyDetailsCmd.MarkFlagRequired("group-identifier")
		pi_getDimensionKeyDetailsCmd.MarkFlagRequired("identifier")
		pi_getDimensionKeyDetailsCmd.MarkFlagRequired("service-type")
	})
	piCmd.AddCommand(pi_getDimensionKeyDetailsCmd)
}
