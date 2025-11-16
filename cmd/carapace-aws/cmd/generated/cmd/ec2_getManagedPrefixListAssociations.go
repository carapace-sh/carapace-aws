package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_getManagedPrefixListAssociationsCmd = &cobra.Command{
	Use:   "get-managed-prefix-list-associations",
	Short: "Gets information about the resources that are associated with the specified managed prefix list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_getManagedPrefixListAssociationsCmd).Standalone()

	ec2_getManagedPrefixListAssociationsCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getManagedPrefixListAssociationsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
	ec2_getManagedPrefixListAssociationsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	ec2_getManagedPrefixListAssociationsCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_getManagedPrefixListAssociationsCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
	ec2_getManagedPrefixListAssociationsCmd.Flag("no-dry-run").Hidden = true
	ec2_getManagedPrefixListAssociationsCmd.MarkFlagRequired("prefix-list-id")
	ec2Cmd.AddCommand(ec2_getManagedPrefixListAssociationsCmd)
}
