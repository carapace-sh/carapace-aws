package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_getOutpostBillingInformationCmd = &cobra.Command{
	Use:   "get-outpost-billing-information",
	Short: "Gets current and historical billing information about the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_getOutpostBillingInformationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_getOutpostBillingInformationCmd).Standalone()

		outposts_getOutpostBillingInformationCmd.Flags().String("max-results", "", "")
		outposts_getOutpostBillingInformationCmd.Flags().String("next-token", "", "")
		outposts_getOutpostBillingInformationCmd.Flags().String("outpost-identifier", "", "The ID or ARN of the Outpost.")
		outposts_getOutpostBillingInformationCmd.MarkFlagRequired("outpost-identifier")
	})
	outpostsCmd.AddCommand(outposts_getOutpostBillingInformationCmd)
}
