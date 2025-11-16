package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_listAttacksCmd = &cobra.Command{
	Use:   "list-attacks",
	Short: "Returns all ongoing DDoS attacks or all DDoS attacks during a specified time period.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_listAttacksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_listAttacksCmd).Standalone()

		shield_listAttacksCmd.Flags().String("end-time", "", "The end of the time period for the attacks.")
		shield_listAttacksCmd.Flags().String("max-results", "", "The greatest number of objects that you want Shield Advanced to return to the list request.")
		shield_listAttacksCmd.Flags().String("next-token", "", "When you request a list of objects from Shield Advanced, if the response does not include all of the remaining available objects, Shield Advanced includes a `NextToken` value in the response.")
		shield_listAttacksCmd.Flags().String("resource-arns", "", "The ARNs (Amazon Resource Names) of the resources that were attacked.")
		shield_listAttacksCmd.Flags().String("start-time", "", "The start of the time period for the attacks.")
	})
	shieldCmd.AddCommand(shield_listAttacksCmd)
}
