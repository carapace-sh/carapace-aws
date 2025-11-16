package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var kinesis_getResourcePolicyCmd = &cobra.Command{
	Use:   "get-resource-policy",
	Short: "Returns a policy attached to the specified data stream or consumer.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(kinesis_getResourcePolicyCmd).Standalone()

	kinesis_getResourcePolicyCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the data stream or consumer.")
	kinesis_getResourcePolicyCmd.MarkFlagRequired("resource-arn")
	kinesisCmd.AddCommand(kinesis_getResourcePolicyCmd)
}
