package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describePatchBaselinesCmd = &cobra.Command{
	Use:   "describe-patch-baselines",
	Short: "Lists the patch baselines in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describePatchBaselinesCmd).Standalone()

	ssm_describePatchBaselinesCmd.Flags().String("filters", "", "Each element in the array is a structure containing a key-value pair.")
	ssm_describePatchBaselinesCmd.Flags().String("max-results", "", "The maximum number of patch baselines to return (per page).")
	ssm_describePatchBaselinesCmd.Flags().String("next-token", "", "The token for the next set of items to return.")
	ssmCmd.AddCommand(ssm_describePatchBaselinesCmd)
}
