package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_createPipeCmd = &cobra.Command{
	Use:   "create-pipe",
	Short: "Create a pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_createPipeCmd).Standalone()

	pipes_createPipeCmd.Flags().String("description", "", "A description of the pipe.")
	pipes_createPipeCmd.Flags().String("desired-state", "", "The state the pipe should be in.")
	pipes_createPipeCmd.Flags().String("enrichment", "", "The ARN of the enrichment resource.")
	pipes_createPipeCmd.Flags().String("enrichment-parameters", "", "The parameters required to set up enrichment on your pipe.")
	pipes_createPipeCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt pipe data.")
	pipes_createPipeCmd.Flags().String("log-configuration", "", "The logging configuration settings for the pipe.")
	pipes_createPipeCmd.Flags().String("name", "", "The name of the pipe.")
	pipes_createPipeCmd.Flags().String("role-arn", "", "The ARN of the role that allows the pipe to send data to the target.")
	pipes_createPipeCmd.Flags().String("source", "", "The ARN of the source resource.")
	pipes_createPipeCmd.Flags().String("source-parameters", "", "The parameters required to set up a source for your pipe.")
	pipes_createPipeCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the pipe.")
	pipes_createPipeCmd.Flags().String("target", "", "The ARN of the target resource.")
	pipes_createPipeCmd.Flags().String("target-parameters", "", "The parameters required to set up a target for your pipe.")
	pipes_createPipeCmd.MarkFlagRequired("name")
	pipes_createPipeCmd.MarkFlagRequired("role-arn")
	pipes_createPipeCmd.MarkFlagRequired("source")
	pipes_createPipeCmd.MarkFlagRequired("target")
	pipesCmd.AddCommand(pipes_createPipeCmd)
}
