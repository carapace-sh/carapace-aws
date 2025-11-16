package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_deleteManagedPrefixListCmd = &cobra.Command{
	Use:   "delete-managed-prefix-list",
	Short: "Deletes the specified managed prefix list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_deleteManagedPrefixListCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_deleteManagedPrefixListCmd).Standalone()

		ec2_deleteManagedPrefixListCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteManagedPrefixListCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_deleteManagedPrefixListCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
		ec2_deleteManagedPrefixListCmd.Flag("no-dry-run").Hidden = true
		ec2_deleteManagedPrefixListCmd.MarkFlagRequired("prefix-list-id")
	})
	ec2Cmd.AddCommand(ec2_deleteManagedPrefixListCmd)
}
