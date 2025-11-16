package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_createPatchBaselineCmd = &cobra.Command{
	Use:   "create-patch-baseline",
	Short: "Creates a patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_createPatchBaselineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_createPatchBaselineCmd).Standalone()

		ssm_createPatchBaselineCmd.Flags().String("approval-rules", "", "A set of rules used to include patches in the baseline.")
		ssm_createPatchBaselineCmd.Flags().String("approved-patches", "", "A list of explicitly approved patches for the baseline.")
		ssm_createPatchBaselineCmd.Flags().String("approved-patches-compliance-level", "", "Defines the compliance level for approved patches.")
		ssm_createPatchBaselineCmd.Flags().Bool("approved-patches-enable-non-security", false, "Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.")
		ssm_createPatchBaselineCmd.Flags().String("available-security-updates-compliance-status", "", "Indicates the status you want to assign to security patches that are available but not approved because they don't meet the installation criteria specified in the patch baseline.")
		ssm_createPatchBaselineCmd.Flags().String("client-token", "", "User-provided idempotency token.")
		ssm_createPatchBaselineCmd.Flags().String("description", "", "A description of the patch baseline.")
		ssm_createPatchBaselineCmd.Flags().String("global-filters", "", "A set of global filters used to include patches in the baseline.")
		ssm_createPatchBaselineCmd.Flags().String("name", "", "The name of the patch baseline.")
		ssm_createPatchBaselineCmd.Flags().Bool("no-approved-patches-enable-non-security", false, "Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.")
		ssm_createPatchBaselineCmd.Flags().String("operating-system", "", "Defines the operating system the patch baseline applies to.")
		ssm_createPatchBaselineCmd.Flags().String("rejected-patches", "", "A list of explicitly rejected patches for the baseline.")
		ssm_createPatchBaselineCmd.Flags().String("rejected-patches-action", "", "The action for Patch Manager to take on patches included in the `RejectedPackages` list.")
		ssm_createPatchBaselineCmd.Flags().String("sources", "", "Information about the patches to use to update the managed nodes, including target operating systems and source repositories.")
		ssm_createPatchBaselineCmd.Flags().String("tags", "", "Optional metadata that you assign to a resource.")
		ssm_createPatchBaselineCmd.MarkFlagRequired("name")
		ssm_createPatchBaselineCmd.Flag("no-approved-patches-enable-non-security").Hidden = true
	})
	ssmCmd.AddCommand(ssm_createPatchBaselineCmd)
}
