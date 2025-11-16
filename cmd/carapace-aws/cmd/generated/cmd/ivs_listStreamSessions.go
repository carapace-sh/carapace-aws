package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_listStreamSessionsCmd = &cobra.Command{
	Use:   "list-stream-sessions",
	Short: "Gets a summary of current and previous streams for a specified channel in your account, in the AWS region where the API request is processed.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_listStreamSessionsCmd).Standalone()

	ivs_listStreamSessionsCmd.Flags().String("channel-arn", "", "Channel ARN used to filter the list.")
	ivs_listStreamSessionsCmd.Flags().String("max-results", "", "Maximum number of streams to return.")
	ivs_listStreamSessionsCmd.Flags().String("next-token", "", "The first stream to retrieve.")
	ivs_listStreamSessionsCmd.MarkFlagRequired("channel-arn")
	ivsCmd.AddCommand(ivs_listStreamSessionsCmd)
}
