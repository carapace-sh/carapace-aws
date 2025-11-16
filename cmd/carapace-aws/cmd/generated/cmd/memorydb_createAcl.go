package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_createAclCmd = &cobra.Command{
	Use:   "create-acl",
	Short: "Creates an Access Control List.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_createAclCmd).Standalone()

	memorydb_createAclCmd.Flags().String("aclname", "", "The name of the Access Control List.")
	memorydb_createAclCmd.Flags().String("tags", "", "A list of tags to be added to this resource.")
	memorydb_createAclCmd.Flags().String("user-names", "", "The list of users that belong to the Access Control List.")
	memorydb_createAclCmd.MarkFlagRequired("aclname")
	memorydbCmd.AddCommand(memorydb_createAclCmd)
}
