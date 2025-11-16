package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_createQuantumTaskCmd = &cobra.Command{
	Use:   "create-quantum-task",
	Short: "Creates a quantum task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_createQuantumTaskCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_createQuantumTaskCmd).Standalone()

		braket_createQuantumTaskCmd.Flags().String("action", "", "The action associated with the quantum task.")
		braket_createQuantumTaskCmd.Flags().String("associations", "", "The list of Amazon Braket resources associated with the quantum task.")
		braket_createQuantumTaskCmd.Flags().String("client-token", "", "The client token associated with the request.")
		braket_createQuantumTaskCmd.Flags().String("device-arn", "", "The ARN of the device to run the quantum task on.")
		braket_createQuantumTaskCmd.Flags().String("device-parameters", "", "The parameters for the device to run the quantum task on.")
		braket_createQuantumTaskCmd.Flags().String("experimental-capabilities", "", "Enable experimental capabilities for the quantum task.")
		braket_createQuantumTaskCmd.Flags().String("job-token", "", "The token for an Amazon Braket hybrid job that associates it with the quantum task.")
		braket_createQuantumTaskCmd.Flags().String("output-s3-bucket", "", "The S3 bucket to store quantum task result files in.")
		braket_createQuantumTaskCmd.Flags().String("output-s3-key-prefix", "", "The key prefix for the location in the S3 bucket to store quantum task results in.")
		braket_createQuantumTaskCmd.Flags().String("shots", "", "The number of shots to use for the quantum task.")
		braket_createQuantumTaskCmd.Flags().String("tags", "", "Tags to be added to the quantum task you're creating.")
		braket_createQuantumTaskCmd.MarkFlagRequired("action")
		braket_createQuantumTaskCmd.MarkFlagRequired("client-token")
		braket_createQuantumTaskCmd.MarkFlagRequired("device-arn")
		braket_createQuantumTaskCmd.MarkFlagRequired("output-s3-bucket")
		braket_createQuantumTaskCmd.MarkFlagRequired("output-s3-key-prefix")
		braket_createQuantumTaskCmd.MarkFlagRequired("shots")
	})
	braketCmd.AddCommand(braket_createQuantumTaskCmd)
}
