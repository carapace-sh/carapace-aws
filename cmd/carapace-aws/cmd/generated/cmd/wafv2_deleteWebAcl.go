package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var wafv2_deleteWebAclCmd = &cobra.Command{
	Use:   "delete-web-acl",
	Short: "Deletes the specified [WebACL]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(wafv2_deleteWebAclCmd).Standalone()

	wafv2_deleteWebAclCmd.Flags().String("id", "", "The unique identifier for the web ACL.")
	wafv2_deleteWebAclCmd.Flags().String("lock-token", "", "A token used for optimistic locking.")
	wafv2_deleteWebAclCmd.Flags().String("name", "", "The name of the web ACL.")
	wafv2_deleteWebAclCmd.Flags().String("scope", "", "Specifies whether this is for a global resource type, such as a Amazon CloudFront distribution.")
	wafv2_deleteWebAclCmd.MarkFlagRequired("id")
	wafv2_deleteWebAclCmd.MarkFlagRequired("lock-token")
	wafv2_deleteWebAclCmd.MarkFlagRequired("name")
	wafv2_deleteWebAclCmd.MarkFlagRequired("scope")
	wafv2Cmd.AddCommand(wafv2_deleteWebAclCmd)
}
