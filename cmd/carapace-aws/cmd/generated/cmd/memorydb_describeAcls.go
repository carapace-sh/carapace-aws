package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var memorydb_describeAclsCmd = &cobra.Command{
	Use:   "describe-acls",
	Short: "Returns a list of ACLs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(memorydb_describeAclsCmd).Standalone()

	memorydb_describeAclsCmd.Flags().String("aclname", "", "The name of the ACL.")
	memorydb_describeAclsCmd.Flags().String("max-results", "", "The maximum number of records to include in the response.")
	memorydb_describeAclsCmd.Flags().String("next-token", "", "An optional argument to pass in case the total number of records exceeds the value of MaxResults.")
	memorydbCmd.AddCommand(memorydb_describeAclsCmd)
}
