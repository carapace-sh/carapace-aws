package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_putDestinationCmd = &cobra.Command{
	Use:   "put-destination",
	Short: "Creates or updates a destination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_putDestinationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(logs_putDestinationCmd).Standalone()

		logs_putDestinationCmd.Flags().String("destination-name", "", "A name for the destination.")
		logs_putDestinationCmd.Flags().String("role-arn", "", "The ARN of an IAM role that grants CloudWatch Logs permissions to call the Amazon Kinesis `PutRecord` operation on the destination stream.")
		logs_putDestinationCmd.Flags().String("tags", "", "An optional list of key-value pairs to associate with the resource.")
		logs_putDestinationCmd.Flags().String("target-arn", "", "The ARN of an Amazon Kinesis stream to which to deliver matching log events.")
		logs_putDestinationCmd.MarkFlagRequired("destination-name")
		logs_putDestinationCmd.MarkFlagRequired("role-arn")
		logs_putDestinationCmd.MarkFlagRequired("target-arn")
	})
	logsCmd.AddCommand(logs_putDestinationCmd)
}
