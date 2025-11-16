package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ec2_createManagedPrefixListCmd = &cobra.Command{
	Use:   "create-managed-prefix-list",
	Short: "Creates a managed prefix list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ec2_createManagedPrefixListCmd).Standalone()

	ec2_createManagedPrefixListCmd.Flags().String("address-family", "", "The IP address type.")
	ec2_createManagedPrefixListCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier you provide to ensure the idempotency of the request.")
	ec2_createManagedPrefixListCmd.Flags().Bool("dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createManagedPrefixListCmd.Flags().String("entries", "", "One or more entries for the prefix list.")
	ec2_createManagedPrefixListCmd.Flags().String("max-entries", "", "The maximum number of entries for the prefix list.")
	ec2_createManagedPrefixListCmd.Flags().Bool("no-dry-run", false, "Checks whether you have the required permissions for the action, without actually making the request, and provides an error response.")
	ec2_createManagedPrefixListCmd.Flags().String("prefix-list-name", "", "A name for the prefix list.")
	ec2_createManagedPrefixListCmd.Flags().String("tag-specifications", "", "The tags to apply to the prefix list during creation.")
	ec2_createManagedPrefixListCmd.MarkFlagRequired("address-family")
	ec2_createManagedPrefixListCmd.MarkFlagRequired("max-entries")
	ec2_createManagedPrefixListCmd.Flag("no-dry-run").Hidden = true
	ec2_createManagedPrefixListCmd.MarkFlagRequired("prefix-list-name")
	ec2Cmd.AddCommand(ec2_createManagedPrefixListCmd)
}
