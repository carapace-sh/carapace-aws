package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_deleteAccessControlConfigurationCmd = &cobra.Command{
	Use:   "delete-access-control-configuration",
	Short: "Deletes an access control configuration that you created for your documents in an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_deleteAccessControlConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(kendra_deleteAccessControlConfigurationCmd).Standalone()

		kendra_deleteAccessControlConfigurationCmd.Flags().String("id", "", "The identifier of the access control configuration you want to delete.")
		kendra_deleteAccessControlConfigurationCmd.Flags().String("index-id", "", "The identifier of the index for an access control configuration.")
		kendra_deleteAccessControlConfigurationCmd.MarkFlagRequired("id")
		kendra_deleteAccessControlConfigurationCmd.MarkFlagRequired("index-id")
	})
	kendraCmd.AddCommand(kendra_deleteAccessControlConfigurationCmd)
}
