package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_deleteAclCmd = &cobra.Command{
	Use:   "delete-acl",
	Short: "Deletes an Access Control List.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_deleteAclCmd).Standalone()

	memorydb_deleteAclCmd.Flags().String("aclname", "", "The name of the Access Control List to delete.")
	memorydb_deleteAclCmd.MarkFlagRequired("aclname")
	memorydbCmd.AddCommand(memorydb_deleteAclCmd)
}
