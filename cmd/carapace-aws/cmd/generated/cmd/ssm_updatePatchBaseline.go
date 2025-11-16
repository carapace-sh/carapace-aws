package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssm_updatePatchBaselineCmd = &cobra.Command{
	Use:   "update-patch-baseline",
	Short: "Modifies an existing patch baseline.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssm_updatePatchBaselineCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssm_updatePatchBaselineCmd).Standalone()

		ssm_updatePatchBaselineCmd.Flags().String("approval-rules", "", "A set of rules used to include patches in the baseline.")
		ssm_updatePatchBaselineCmd.Flags().String("approved-patches", "", "A list of explicitly approved patches for the baseline.")
		ssm_updatePatchBaselineCmd.Flags().String("approved-patches-compliance-level", "", "Assigns a new compliance severity level to an existing patch baseline.")
		ssm_updatePatchBaselineCmd.Flags().Bool("approved-patches-enable-non-security", false, "Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.")
		ssm_updatePatchBaselineCmd.Flags().String("available-security-updates-compliance-status", "", "Indicates the status to be assigned to security patches that are available but not approved because they don't meet the installation criteria specified in the patch baseline.")
		ssm_updatePatchBaselineCmd.Flags().String("baseline-id", "", "The ID of the patch baseline to update.")
		ssm_updatePatchBaselineCmd.Flags().String("description", "", "A description of the patch baseline.")
		ssm_updatePatchBaselineCmd.Flags().String("global-filters", "", "A set of global filters used to include patches in the baseline.")
		ssm_updatePatchBaselineCmd.Flags().String("name", "", "The name of the patch baseline.")
		ssm_updatePatchBaselineCmd.Flags().Bool("no-approved-patches-enable-non-security", false, "Indicates whether the list of approved patches includes non-security updates that should be applied to the managed nodes.")
		ssm_updatePatchBaselineCmd.Flags().Bool("no-replace", false, "If True, then all fields that are required by the [CreatePatchBaseline]() operation are also required for this API request.")
		ssm_updatePatchBaselineCmd.Flags().String("rejected-patches", "", "A list of explicitly rejected patches for the baseline.")
		ssm_updatePatchBaselineCmd.Flags().String("rejected-patches-action", "", "The action for Patch Manager to take on patches included in the `RejectedPackages` list.")
		ssm_updatePatchBaselineCmd.Flags().Bool("replace", false, "If True, then all fields that are required by the [CreatePatchBaseline]() operation are also required for this API request.")
		ssm_updatePatchBaselineCmd.Flags().String("sources", "", "Information about the patches to use to update the managed nodes, including target operating systems and source repositories.")
		ssm_updatePatchBaselineCmd.MarkFlagRequired("baseline-id")
		ssm_updatePatchBaselineCmd.Flag("no-approved-patches-enable-non-security").Hidden = true
		ssm_updatePatchBaselineCmd.Flag("no-replace").Hidden = true
	})
	ssmCmd.AddCommand(ssm_updatePatchBaselineCmd)
}
