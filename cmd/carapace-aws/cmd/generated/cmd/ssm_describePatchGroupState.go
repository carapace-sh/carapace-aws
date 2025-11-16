package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_describePatchGroupStateCmd = &cobra.Command{
	Use:   "describe-patch-group-state",
	Short: "Returns high-level aggregated patch compliance state information for a patch group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_describePatchGroupStateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_describePatchGroupStateCmd).Standalone()

		ssm_describePatchGroupStateCmd.Flags().String("patch-group", "", "The name of the patch group whose patch snapshot should be retrieved.")
		ssm_describePatchGroupStateCmd.MarkFlagRequired("patch-group")
	})
	ssmCmd.AddCommand(ssm_describePatchGroupStateCmd)
}
