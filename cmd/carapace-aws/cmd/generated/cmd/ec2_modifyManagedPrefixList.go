package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_modifyManagedPrefixListCmd = &cobra.Command{
	Use:   "modify-managed-prefix-list",
	Short: "Modifies the specified managed prefix list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_modifyManagedPrefixListCmd).Standalone()

	ec2_modifyManagedPrefixListCmd.Flags().String("add-entries", "", "One or more entries to add to the prefix list.")
	ec2_modifyManagedPrefixListCmd.Flags().String("current-version", "", "The current version of the prefix list.")
	ec2_modifyManagedPrefixListCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyManagedPrefixListCmd.Flags().String("ipam-prefix-list-resolver-sync-enabled", "", "Indicates whether synchronization with an IPAM prefix list resolver should be enabled for this managed prefix list.")
	ec2_modifyManagedPrefixListCmd.Flags().String("max-entries", "", "The maximum number of entries for the prefix list.")
	ec2_modifyManagedPrefixListCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_modifyManagedPrefixListCmd.Flags().String("prefix-list-id", "", "The ID of the prefix list.")
	ec2_modifyManagedPrefixListCmd.Flags().String("prefix-list-name", "", "A name for the prefix list.")
	ec2_modifyManagedPrefixListCmd.Flags().String("remove-entries", "", "One or more entries to remove from the prefix list.")
	ec2_modifyManagedPrefixListCmd.Flag("no-dry-run").Hidden = true
	ec2_modifyManagedPrefixListCmd.MarkFlagRequired("prefix-list-id")
	ec2Cmd.AddCommand(ec2_modifyManagedPrefixListCmd)
}
