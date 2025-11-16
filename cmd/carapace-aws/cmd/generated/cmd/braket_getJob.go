package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_getJobCmd = &cobra.Command{
	Use:   "get-job",
	Short: "Retrieves the specified Amazon Braket hybrid job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_getJobCmd).Standalone()

	braket_getJobCmd.Flags().String("additional-attribute-names", "", "A list of attributes to return additional information for.")
	braket_getJobCmd.Flags().String("job-arn", "", "The ARN of the hybrid job to retrieve.")
	braket_getJobCmd.MarkFlagRequired("job-arn")
	braketCmd.AddCommand(braket_getJobCmd)
}
