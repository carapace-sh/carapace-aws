package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getManagedPrefixListEntriesCmd = &cobra.Command{
	Use:   "get-managed-prefix-list-entries",
	Short: "Gets information about the entries for a specified managed prefix list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getManagedPrefixListEntriesCmd).Standalone()

	ec2_getManagedPrefixListEntriesCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getManagedPrefixListEntriesCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_getManagedPrefixListEntriesCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getManagedPrefixListEntriesCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getManagedPrefixListEntriesCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
	ec2_getManagedPrefixListEntriesCmd.Flags().String("target-version", "", "The version of the prefix list for which to return the entries.")
	ec2_getManagedPrefixListEntriesCmd.Flag("no-dry-run").Hidden = true
	ec2_getManagedPrefixListEntriesCmd.MarkFlagRequired("prefix-list-id")
	ec2Cmd.AddCommand(ec2_getManagedPrefixListEntriesCmd)
}
