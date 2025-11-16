package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector_addAttributesToFindingsCmd = &cobra.Command{
	Use:   "add-attributes-to-findings",
	Short: "Assigns attributes (key and value pairs) to the findings that are specified by the ARNs of the findings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector_addAttributesToFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector_addAttributesToFindingsCmd).Standalone()

		inspector_addAttributesToFindingsCmd.Flags().String("attributes", "", "The array of attributes that you want to assign to specified findings.")
		inspector_addAttributesToFindingsCmd.Flags().String("finding-arns", "", "The ARNs that specify the findings that you want to assign attributes to.")
		inspector_addAttributesToFindingsCmd.MarkFlagRequired("attributes")
		inspector_addAttributesToFindingsCmd.MarkFlagRequired("finding-arns")
	})
	inspectorCmd.AddCommand(inspector_addAttributesToFindingsCmd)
}
