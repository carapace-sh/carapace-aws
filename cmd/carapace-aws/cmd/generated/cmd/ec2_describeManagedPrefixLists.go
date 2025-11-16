package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describeManagedPrefixListsCmd = &cobra.Command{
	Use:   "describe-managed-prefix-lists",
	Short: "Describes your managed prefix lists and any Amazon Web Services-managed prefix lists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describeManagedPrefixListsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_describeManagedPrefixListsCmd).Standalone()

		ec2_describeManagedPrefixListsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeManagedPrefixListsCmd.Flags().String("filters", "", "One or more filters.")
		ec2_describeManagedPrefixListsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		ec2_describeManagedPrefixListsCmd.Flags().String("next-token", "", "The token for the next page of results.")
		ec2_describeManagedPrefixListsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_describeManagedPrefixListsCmd.Flags().String("prefix-list-ids", "", "One or more prefix list IDs.")
		ec2_describeManagedPrefixListsCmd.Flag("no-dry-run").Hidden = true
	})
	ec2Cmd.AddCommand(ec2_describeManagedPrefixListsCmd)
}
