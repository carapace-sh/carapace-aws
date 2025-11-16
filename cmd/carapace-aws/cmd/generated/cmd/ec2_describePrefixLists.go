package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_describePrefixListsCmd = &cobra.Command{
	Use:   "describe-prefix-lists",
	Short: "Describes available Amazon Web Services services in a prefix list format, which includes the prefix list name and prefix list ID of the service and the IP address range for the service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_describePrefixListsCmd).Standalone()

	ec2_describePrefixListsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describePrefixListsCmd.Flags().String("filters", "", "One or more filters.")
	ec2_describePrefixListsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_describePrefixListsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_describePrefixListsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_describePrefixListsCmd.Flags().String("prefix-list-ids", "", "One or more prefix list IDs.")
	ec2_describePrefixListsCmd.Flag("no-dry-run").Hidden = true
	ec2Cmd.AddCommand(ec2_describePrefixListsCmd)
}
