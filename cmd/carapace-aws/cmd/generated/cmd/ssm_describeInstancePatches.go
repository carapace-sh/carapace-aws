package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describeInstancePatchesCmd = &cobra.Command{
	Use:   "describe-instance-patches",
	Short: "Retrieves information about the patches on the specified managed node and their state relative to the patch baseline being used for the node.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describeInstancePatchesCmd).Standalone()

	ssm_describeInstancePatchesCmd.Flags().String("filters", "", "Each element in the array is a structure containing a key-value pair.")
	ssm_describeInstancePatchesCmd.Flags().String("instance-id", "", "The ID of the managed node whose patch state information should be retrieved.")
	ssm_describeInstancePatchesCmd.Flags().String("max-results", "", "The maximum number of patches to return (per page).")
	ssm_describeInstancePatchesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssm_describeInstancePatchesCmd.MarkFlagRequired("instance-id")
	ssmCmd.AddCommand(ssm_describeInstancePatchesCmd)
}
