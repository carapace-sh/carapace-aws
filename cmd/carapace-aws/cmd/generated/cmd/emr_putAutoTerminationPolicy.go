package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_putAutoTerminationPolicyCmd = &cobra.Command{
	Use:   "put-auto-termination-policy",
	Short: "Auto-termination is supported in Amazon EMR releases 5.30.0 and 6.1.0 and later.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_putAutoTerminationPolicyCmd).Standalone()

	emr_putAutoTerminationPolicyCmd.Flags().String("auto-termination-policy", "", "Specifies the auto-termination policy to attach to the cluster.")
	emr_putAutoTerminationPolicyCmd.Flags().String("cluster-id", "", "Specifies the ID of the Amazon EMR cluster to which the auto-termination policy will be attached.")
	emr_putAutoTerminationPolicyCmd.MarkFlagRequired("cluster-id")
	emrCmd.AddCommand(emr_putAutoTerminationPolicyCmd)
}
