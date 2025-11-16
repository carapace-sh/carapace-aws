package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var braket_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "Creates an Amazon Braket hybrid job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(braket_createJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(braket_createJobCmd).Standalone()

		braket_createJobCmd.Flags().String("algorithm-specification", "", "Definition of the Amazon Braket job to be created.")
		braket_createJobCmd.Flags().String("associations", "", "The list of Amazon Braket resources associated with the hybrid job.")
		braket_createJobCmd.Flags().String("checkpoint-config", "", "Information about the output locations for hybrid job checkpoint data.")
		braket_createJobCmd.Flags().String("client-token", "", "The client token associated with this request that guarantees that the request is idempotent.")
		braket_createJobCmd.Flags().String("device-config", "", "The quantum processing unit (QPU) or simulator used to create an Amazon Braket hybrid job.")
		braket_createJobCmd.Flags().String("hyper-parameters", "", "Algorithm-specific parameters used by an Amazon Braket hybrid job that influence the quality of the training job.")
		braket_createJobCmd.Flags().String("input-data-config", "", "A list of parameters that specify the name and type of input data and where it is located.")
		braket_createJobCmd.Flags().String("instance-config", "", "Configuration of the resource instances to use while running the hybrid job on Amazon Braket.")
		braket_createJobCmd.Flags().String("job-name", "", "The name of the Amazon Braket hybrid job.")
		braket_createJobCmd.Flags().String("output-data-config", "", "The path to the S3 location where you want to store hybrid job artifacts and the encryption key used to store them.")
		braket_createJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of an IAM role that Amazon Braket can assume to perform tasks on behalf of a user.")
		braket_createJobCmd.Flags().String("stopping-condition", "", "The user-defined criteria that specifies when a hybrid job stops running.")
		braket_createJobCmd.Flags().String("tags", "", "Tags to be added to the hybrid job you're creating.")
		braket_createJobCmd.MarkFlagRequired("algorithm-specification")
		braket_createJobCmd.MarkFlagRequired("client-token")
		braket_createJobCmd.MarkFlagRequired("device-config")
		braket_createJobCmd.MarkFlagRequired("instance-config")
		braket_createJobCmd.MarkFlagRequired("job-name")
		braket_createJobCmd.MarkFlagRequired("output-data-config")
		braket_createJobCmd.MarkFlagRequired("role-arn")
	})
	braketCmd.AddCommand(braket_createJobCmd)
}
