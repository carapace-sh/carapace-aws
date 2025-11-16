package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_checkCapacityCmd = &cobra.Command{
	Use:   "check-capacity",
	Short: "Returns the web ACL capacity unit (WCU) requirements for a specified scope and set of rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_checkCapacityCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_checkCapacityCmd).Standalone()

		wafv2_checkCapacityCmd.Flags().String("rules", "", "An array of [Rule]() that you're configuring to use in a rule group or web ACL.")
		wafv2_checkCapacityCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_checkCapacityCmd.MarkFlagRequired("rules")
		wafv2_checkCapacityCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_checkCapacityCmd)
}
