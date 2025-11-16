package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehendmedical_startPhidetectionJobCmd = &cobra.Command{
	Use:   "start-phidetection-job",
	Short: "Starts an asynchronous job to detect protected health information (PHI).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehendmedical_startPhidetectionJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(comprehendmedical_startPhidetectionJobCmd).Standalone()

		comprehendmedical_startPhidetectionJobCmd.Flags().String("client-request-token", "", "A unique identifier for the request.")
		comprehendmedical_startPhidetectionJobCmd.Flags().String("data-access-role-arn", "", "The Amazon Resource Name (ARN) of the AWS Identity and Access Management (IAM) role that grants Amazon Comprehend Medical read access to your input data.")
		comprehendmedical_startPhidetectionJobCmd.Flags().String("input-data-config", "", "Specifies the format and location of the input data for the job.")
		comprehendmedical_startPhidetectionJobCmd.Flags().String("job-name", "", "The identifier of the job.")
		comprehendmedical_startPhidetectionJobCmd.Flags().String("kmskey", "", "An AWS Key Management Service key to encrypt your output files.")
		comprehendmedical_startPhidetectionJobCmd.Flags().String("language-code", "", "The language of the input documents.")
		comprehendmedical_startPhidetectionJobCmd.Flags().String("output-data-config", "", "Specifies where to send the output files.")
		comprehendmedical_startPhidetectionJobCmd.MarkFlagRequired("data-access-role-arn")
		comprehendmedical_startPhidetectionJobCmd.MarkFlagRequired("input-data-config")
		comprehendmedical_startPhidetectionJobCmd.MarkFlagRequired("language-code")
		comprehendmedical_startPhidetectionJobCmd.MarkFlagRequired("output-data-config")
	})
	comprehendmedicalCmd.AddCommand(comprehendmedical_startPhidetectionJobCmd)
}
