package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_setTerminationProtectionCmd = &cobra.Command{
	Use:   "set-termination-protection",
	Short: "SetTerminationProtection locks a cluster (job flow) so the Amazon EC2 instances in the cluster cannot be terminated by user intervention, an API call, or in the event of a job-flow error.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_setTerminationProtectionCmd).Standalone()

	emr_setTerminationProtectionCmd.Flags().String("job-flow-ids", "", "A list of strings that uniquely identify the clusters to protect.")
	emr_setTerminationProtectionCmd.Flags().Bool("no-termination-protected", false, "A Boolean that indicates whether to protect the cluster and prevent the Amazon EC2 instances in the cluster from shutting down due to API calls, user intervention, or job-flow error.")
	emr_setTerminationProtectionCmd.Flags().Bool("termination-protected", false, "A Boolean that indicates whether to protect the cluster and prevent the Amazon EC2 instances in the cluster from shutting down due to API calls, user intervention, or job-flow error.")
	emr_setTerminationProtectionCmd.MarkFlagRequired("job-flow-ids")
	emr_setTerminationProtectionCmd.Flag("no-termination-protected").Hidden = true
	emr_setTerminationProtectionCmd.MarkFlagRequired("no-termination-protected")
	emr_setTerminationProtectionCmd.MarkFlagRequired("termination-protected")
	emrCmd.AddCommand(emr_setTerminationProtectionCmd)
}
