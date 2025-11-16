package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_putResourcePolicyCmd = &cobra.Command{
	Use:   "put-resource-policy",
	Short: "Attaches a resource-based policy to a data stream or registered consumer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_putResourcePolicyCmd).Standalone()

	kinesis_putResourcePolicyCmd.Flags().String("policy", "", "Details of the resource policy.")
	kinesis_putResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the data stream or consumer.")
	kinesis_putResourcePolicyCmd.MarkFlagRequired("policy")
	kinesis_putResourcePolicyCmd.MarkFlagRequired("resource-arn")
	kinesisCmd.AddCommand(kinesis_putResourcePolicyCmd)
}
