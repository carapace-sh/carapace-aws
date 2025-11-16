package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchGetFindingDetailsCmd = &cobra.Command{
	Use:   "batch-get-finding-details",
	Short: "Gets vulnerability details for findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchGetFindingDetailsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_batchGetFindingDetailsCmd).Standalone()

		inspector2_batchGetFindingDetailsCmd.Flags().String("finding-arns", "", "A list of finding ARNs.")
		inspector2_batchGetFindingDetailsCmd.MarkFlagRequired("finding-arns")
	})
	inspector2Cmd.AddCommand(inspector2_batchGetFindingDetailsCmd)
}
