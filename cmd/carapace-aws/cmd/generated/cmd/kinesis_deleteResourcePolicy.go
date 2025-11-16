package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_deleteResourcePolicyCmd = &cobra.Command{
	Use:   "delete-resource-policy",
	Short: "Delete a policy for the specified data stream or consumer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_deleteResourcePolicyCmd).Standalone()

	kinesis_deleteResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the data stream or consumer.")
	kinesis_deleteResourcePolicyCmd.MarkFlagRequired("resource-arn")
	kinesisCmd.AddCommand(kinesis_deleteResourcePolicyCmd)
}
