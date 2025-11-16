package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_getDataProtectionPolicyCmd = &cobra.Command{
	Use:   "get-data-protection-policy",
	Short: "Retrieves the specified inline `DataProtectionPolicy` document that is stored in the specified Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_getDataProtectionPolicyCmd).Standalone()

	sns_getDataProtectionPolicyCmd.Flags().String("resource-arn", "", "The ARN of the topic whose `DataProtectionPolicy` you want to get.")
	sns_getDataProtectionPolicyCmd.MarkFlagRequired("resource-arn")
	snsCmd.AddCommand(sns_getDataProtectionPolicyCmd)
}
