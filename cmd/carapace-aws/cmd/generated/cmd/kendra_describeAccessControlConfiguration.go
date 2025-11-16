package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kendra_describeAccessControlConfigurationCmd = &cobra.Command{
	Use:   "describe-access-control-configuration",
	Short: "Gets information about an access control configuration that you created for your documents in an index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kendra_describeAccessControlConfigurationCmd).Standalone()

	kendra_describeAccessControlConfigurationCmd.Flags().String("id", "", "The identifier of the access control configuration you want to get information on.")
	kendra_describeAccessControlConfigurationCmd.Flags().String("index-id", "", "The identifier of the index for an access control configuration.")
	kendra_describeAccessControlConfigurationCmd.MarkFlagRequired("id")
	kendra_describeAccessControlConfigurationCmd.MarkFlagRequired("index-id")
	kendraCmd.AddCommand(kendra_describeAccessControlConfigurationCmd)
}
