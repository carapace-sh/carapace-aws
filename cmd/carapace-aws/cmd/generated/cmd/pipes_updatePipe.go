package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pipes_updatePipeCmd = &cobra.Command{
	Use:   "update-pipe",
	Short: "Update an existing pipe.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pipes_updatePipeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(pipes_updatePipeCmd).Standalone()

		pipes_updatePipeCmd.Flags().String("description", "", "A description of the pipe.")
		pipes_updatePipeCmd.Flags().String("desired-state", "", "The state the pipe should be in.")
		pipes_updatePipeCmd.Flags().String("enrichment", "", "The ARN of the enrichment resource.")
		pipes_updatePipeCmd.Flags().String("enrichment-parameters", "", "The parameters required to set up enrichment on your pipe.")
		pipes_updatePipeCmd.Flags().String("kms-key-identifier", "", "The identifier of the KMS customer managed key for EventBridge to use, if you choose to use a customer managed key to encrypt pipe data.")
		pipes_updatePipeCmd.Flags().String("log-configuration", "", "The logging configuration settings for the pipe.")
		pipes_updatePipeCmd.Flags().String("name", "", "The name of the pipe.")
		pipes_updatePipeCmd.Flags().String("role-arn", "", "The ARN of the role that allows the pipe to send data to the target.")
		pipes_updatePipeCmd.Flags().String("source-parameters", "", "The parameters required to set up a source for your pipe.")
		pipes_updatePipeCmd.Flags().String("target", "", "The ARN of the target resource.")
		pipes_updatePipeCmd.Flags().String("target-parameters", "", "The parameters required to set up a target for your pipe.")
		pipes_updatePipeCmd.MarkFlagRequired("name")
		pipes_updatePipeCmd.MarkFlagRequired("role-arn")
	})
	pipesCmd.AddCommand(pipes_updatePipeCmd)
}
