package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInventoryDeletionsCmd = &cobra.Command{
	Use:   "describe-inventory-deletions",
	Short: "Describes a specific delete inventory operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInventoryDeletionsCmd).Standalone()

	ssm_describeInventoryDeletionsCmd.Flags().String("deletion-id", "", "Specify the delete inventory ID for which you want information.")
	ssm_describeInventoryDeletionsCmd.Flags().String("max-results", "", "The maximum number of items to return for this call.")
	ssm_describeInventoryDeletionsCmd.Flags().String("next-token", "", "A token to start the list.")
	ssmCmd.AddCommand(ssm_describeInventoryDeletionsCmd)
}
