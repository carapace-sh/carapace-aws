package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listConnectionGroups2020_05_31Cmd = &cobra.Command{
	Use:   "list-connection-groups2020_05_31",
	Short: "Lists the connection groups in your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listConnectionGroups2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listConnectionGroups2020_05_31Cmd).Standalone()

		cloudfront_listConnectionGroups2020_05_31Cmd.Flags().String("association-filter", "", "Filter by associated Anycast IP list ID.")
		cloudfront_listConnectionGroups2020_05_31Cmd.Flags().String("marker", "", "The marker for the next set of connection groups to retrieve.")
		cloudfront_listConnectionGroups2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of connection groups to return.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listConnectionGroups2020_05_31Cmd)
}
