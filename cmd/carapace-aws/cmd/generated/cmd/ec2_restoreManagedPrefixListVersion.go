package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_restoreManagedPrefixListVersionCmd = &cobra.Command{
	Use:   "restore-managed-prefix-list-version",
	Short: "Restores the entries from a previous version of a managed prefix list to a new version of the prefix list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_restoreManagedPrefixListVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ec2_restoreManagedPrefixListVersionCmd).Standalone()

		ec2_restoreManagedPrefixListVersionCmd.Flags().String("current-version", "", "The current version number for the prefix list.")
		ec2_restoreManagedPrefixListVersionCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_restoreManagedPrefixListVersionCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
		ec2_restoreManagedPrefixListVersionCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
		ec2_restoreManagedPrefixListVersionCmd.Flags().String("previous-version", "", "The version to restore.")
		ec2_restoreManagedPrefixListVersionCmd.MarkFlagRequired("current-version")
		ec2_restoreManagedPrefixListVersionCmd.Flag("no-dry-run").Hidden = true
		ec2_restoreManagedPrefixListVersionCmd.MarkFlagRequired("prefix-list-id")
		ec2_restoreManagedPrefixListVersionCmd.MarkFlagRequired("previous-version")
	})
	ec2Cmd.AddCommand(ec2_restoreManagedPrefixListVersionCmd)
}
