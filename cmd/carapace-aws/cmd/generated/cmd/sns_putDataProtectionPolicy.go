package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_putDataProtectionPolicyCmd = &cobra.Command{
	Use:   "put-data-protection-policy",
	Short: "Adds or updates an inline policy document that is stored in the specified Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_putDataProtectionPolicyCmd).Standalone()

	sns_putDataProtectionPolicyCmd.Flags().String("data-protection-policy", "", "The JSON serialization of the topic's `DataProtectionPolicy`.")
	sns_putDataProtectionPolicyCmd.Flags().String("resource-arn", "", "The ARN of the topic whose `DataProtectionPolicy` you want to add or update.")
	sns_putDataProtectionPolicyCmd.MarkFlagRequired("data-protection-policy")
	sns_putDataProtectionPolicyCmd.MarkFlagRequired("resource-arn")
	snsCmd.AddCommand(sns_putDataProtectionPolicyCmd)
}
