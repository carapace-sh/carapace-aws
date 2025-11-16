package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rbin_createRuleCmd = &cobra.Command{
	Use:   "create-rule",
	Short: "Creates a Recycle Bin retention rule.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rbin_createRuleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rbin_createRuleCmd).Standalone()

		rbin_createRuleCmd.Flags().String("description", "", "The retention rule description.")
		rbin_createRuleCmd.Flags().String("exclude-resource-tags", "", "\\[Region-level retention rules only] Specifies the exclusion tags to use to identify resources that are to be excluded, or ignored, by a Region-level retention rule.")
		rbin_createRuleCmd.Flags().String("lock-configuration", "", "Information about the retention rule lock configuration.")
		rbin_createRuleCmd.Flags().String("resource-tags", "", "\\[Tag-level retention rules only] Specifies the resource tags to use to identify resources that are to be retained by a tag-level retention rule.")
		rbin_createRuleCmd.Flags().String("resource-type", "", "The resource type to be retained by the retention rule.")
		rbin_createRuleCmd.Flags().String("retention-period", "", "Information about the retention period for which the retention rule is to retain resources.")
		rbin_createRuleCmd.Flags().String("tags", "", "Information about the tags to assign to the retention rule.")
		rbin_createRuleCmd.MarkFlagRequired("resource-type")
		rbin_createRuleCmd.MarkFlagRequired("retention-period")
	})
	rbinCmd.AddCommand(rbin_createRuleCmd)
}
