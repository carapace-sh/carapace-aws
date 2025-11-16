package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var xray_getSamplingRulesCmd = &cobra.Command{
	Use:   "get-sampling-rules",
	Short: "Retrieves all sampling rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xray_getSamplingRulesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(xray_getSamplingRulesCmd).Standalone()

		xray_getSamplingRulesCmd.Flags().String("next-token", "", "Pagination token.")
	})
	xrayCmd.AddCommand(xray_getSamplingRulesCmd)
}
