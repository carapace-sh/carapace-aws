package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "Cancels an Amazon Braket hybrid job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_cancelJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_cancelJobCmd).Standalone()

		braket_cancelJobCmd.Flags().String("job-arn", "", "The ARN of the Amazon Braket hybrid job to cancel.")
		braket_cancelJobCmd.MarkFlagRequired("job-arn")
	})
	braketCmd.AddCommand(braket_cancelJobCmd)
}
