package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_createAccessControlConfigurationCmd = &cobra.Command{
	Use:   "create-access-control-configuration",
	Short: "Creates an access configuration for your documents.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_createAccessControlConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_createAccessControlConfigurationCmd).Standalone()

		kendra_createAccessControlConfigurationCmd.Flags().String("access-control-list", "", "Information on principals (users and/or groups) and which documents they should have access to.")
		kendra_createAccessControlConfigurationCmd.Flags().String("client-token", "", "A token that you provide to identify the request to create an access control configuration.")
		kendra_createAccessControlConfigurationCmd.Flags().String("description", "", "A description for the access control configuration.")
		kendra_createAccessControlConfigurationCmd.Flags().String("hierarchical-access-control-list", "", "The list of [principal](https://docs.aws.amazon.com/kendra/latest/dg/API_Principal.html) lists that define the hierarchy for which documents users should have access to.")
		kendra_createAccessControlConfigurationCmd.Flags().String("index-id", "", "The identifier of the index to create an access control configuration for your documents.")
		kendra_createAccessControlConfigurationCmd.Flags().String("name", "", "A name for the access control configuration.")
		kendra_createAccessControlConfigurationCmd.MarkFlagRequired("index-id")
		kendra_createAccessControlConfigurationCmd.MarkFlagRequired("name")
	})
	kendraCmd.AddCommand(kendra_createAccessControlConfigurationCmd)
}
