package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeAvailablePatchesCmd = &cobra.Command{
	Use:   "describe-available-patches",
	Short: "Lists all patches eligible to be included in a patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeAvailablePatchesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describeAvailablePatchesCmd).Standalone()

		ssm_describeAvailablePatchesCmd.Flags().String("filters", "", "Each element in the array is a structure containing a key-value pair.")
		ssm_describeAvailablePatchesCmd.Flags().String("max-results", "", "The maximum number of patches to return (per page).")
		ssm_describeAvailablePatchesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	})
	ssmCmd.AddCommand(ssm_describeAvailablePatchesCmd)
}
