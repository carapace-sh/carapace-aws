package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_listTrustAnchorsCmd = &cobra.Command{
	Use:   "list-trust-anchors",
	Short: "Lists the trust anchors in the authenticated account and Amazon Web Services Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_listTrustAnchorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_listTrustAnchorsCmd).Standalone()

		rolesanywhere_listTrustAnchorsCmd.Flags().String("next-token", "", "A token that indicates where the output should continue from, if a previous request did not show all results.")
		rolesanywhere_listTrustAnchorsCmd.Flags().String("page-size", "", "The number of resources in the paginated list.")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_listTrustAnchorsCmd)
}
