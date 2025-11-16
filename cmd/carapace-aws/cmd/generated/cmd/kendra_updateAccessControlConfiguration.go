package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_updateAccessControlConfigurationCmd = &cobra.Command{
	Use:   "update-access-control-configuration",
	Short: "Updates an access control configuration for your documents in an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_updateAccessControlConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_updateAccessControlConfigurationCmd).Standalone()

		kendra_updateAccessControlConfigurationCmd.Flags().String("access-control-list", "", "Information you want to update on principals (users and/or groups) and which documents they should have access to.")
		kendra_updateAccessControlConfigurationCmd.Flags().String("description", "", "A new description for the access control configuration.")
		kendra_updateAccessControlConfigurationCmd.Flags().String("hierarchical-access-control-list", "", "The updated list of [principal](https://docs.aws.amazon.com/kendra/latest/dg/API_Principal.html) lists that define the hierarchy for which documents users should have access to.")
		kendra_updateAccessControlConfigurationCmd.Flags().String("id", "", "The identifier of the access control configuration you want to update.")
		kendra_updateAccessControlConfigurationCmd.Flags().String("index-id", "", "The identifier of the index for an access control configuration.")
		kendra_updateAccessControlConfigurationCmd.Flags().String("name", "", "A new name for the access control configuration.")
		kendra_updateAccessControlConfigurationCmd.MarkFlagRequired("id")
		kendra_updateAccessControlConfigurationCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_updateAccessControlConfigurationCmd)
}
