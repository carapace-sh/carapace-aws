package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_createResiliencyPolicyCmd = &cobra.Command{
	Use:   "create-resiliency-policy",
	Short: "Creates a resiliency policy for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_createResiliencyPolicyCmd).Standalone()

	resiliencehub_createResiliencyPolicyCmd.Flags().String("client-token", "", "Used for an idempotency token.")
	resiliencehub_createResiliencyPolicyCmd.Flags().String("data-location-constraint", "", "Specifies a high-level geographical location constraint for where your resilience policy data can be stored.")
	resiliencehub_createResiliencyPolicyCmd.Flags().String("policy", "", "The type of resiliency policy to be created, including the recovery time objective (RTO) and recovery point objective (RPO) in seconds.")
	resiliencehub_createResiliencyPolicyCmd.Flags().String("policy-description", "", "Description of the resiliency policy.")
	resiliencehub_createResiliencyPolicyCmd.Flags().String("policy-name", "", "Name of the resiliency policy.")
	resiliencehub_createResiliencyPolicyCmd.Flags().String("tags", "", "Tags assigned to the resource.")
	resiliencehub_createResiliencyPolicyCmd.Flags().String("tier", "", "The tier for this resiliency policy, ranging from the highest severity (`MissionCritical`) to lowest (`NonCritical`).")
	resiliencehub_createResiliencyPolicyCmd.MarkFlagRequired("policy")
	resiliencehub_createResiliencyPolicyCmd.MarkFlagRequired("policy-name")
	resiliencehub_createResiliencyPolicyCmd.MarkFlagRequired("tier")
	resiliencehubCmd.AddCommand(resiliencehub_createResiliencyPolicyCmd)
}
