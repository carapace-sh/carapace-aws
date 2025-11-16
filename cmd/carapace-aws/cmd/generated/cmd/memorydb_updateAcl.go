package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_updateAclCmd = &cobra.Command{
	Use:   "update-acl",
	Short: "Changes the list of users that belong to the Access Control List.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_updateAclCmd).Standalone()

	memorydb_updateAclCmd.Flags().String("aclname", "", "The name of the Access Control List.")
	memorydb_updateAclCmd.Flags().String("user-names-to-add", "", "The list of users to add to the Access Control List.")
	memorydb_updateAclCmd.Flags().String("user-names-to-remove", "", "The list of users to remove from the Access Control List.")
	memorydb_updateAclCmd.MarkFlagRequired("aclname")
	memorydbCmd.AddCommand(memorydb_updateAclCmd)
}
