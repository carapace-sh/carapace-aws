package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteIpsetCmd = &cobra.Command{
	Use:   "delete-ipset",
	Short: "Deletes the specified [IPSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteIpsetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(wafv2_deleteIpsetCmd).Standalone()

		wafv2_deleteIpsetCmd.Flags().String("id", "", "A unique identifier for the set.")
		wafv2_deleteIpsetCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
		wafv2_deleteIpsetCmd.Flags().String("name", "", "The name of the IP set.")
		wafv2_deleteIpsetCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
		wafv2_deleteIpsetCmd.MarkFlagRequired("id")
		wafv2_deleteIpsetCmd.MarkFlagRequired("lock-token")
		wafv2_deleteIpsetCmd.MarkFlagRequired("name")
		wafv2_deleteIpsetCmd.MarkFlagRequired("scope")
	})
	wafv2Cmd.AddCommand(wafv2_deleteIpsetCmd)
}
