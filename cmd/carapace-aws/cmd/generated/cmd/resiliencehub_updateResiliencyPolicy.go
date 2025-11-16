package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_updateResiliencyPolicyCmd = &cobra.Command{
	Use:   "update-resiliency-policy",
	Short: "Updates a resiliency policy.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_updateResiliencyPolicyCmd).Standalone()

	resiliencehub_updateResiliencyPolicyCmd.Flags().String("data-location-constraint", "", "Specifies a high-level geographical location constraint for where your resilience policy data can be stored.")
	resiliencehub_updateResiliencyPolicyCmd.Flags().String("policy", "", "Resiliency policy to be created, including the recovery time objective (RTO) and recovery point objective (RPO) in seconds.")
	resiliencehub_updateResiliencyPolicyCmd.Flags().String("policy-arn", "", "Amazon Resource Name (ARN) of the resiliency policy.")
	resiliencehub_updateResiliencyPolicyCmd.Flags().String("policy-description", "", "Description of the resiliency policy.")
	resiliencehub_updateResiliencyPolicyCmd.Flags().String("policy-name", "", "Name of the resiliency policy.")
	resiliencehub_updateResiliencyPolicyCmd.Flags().String("tier", "", "The tier for this resiliency policy, ranging from the highest severity (`MissionCritical`) to lowest (`NonCritical`).")
	resiliencehub_updateResiliencyPolicyCmd.MarkFlagRequired("policy-arn")
	resiliencehubCmd.AddCommand(resiliencehub_updateResiliencyPolicyCmd)
}
