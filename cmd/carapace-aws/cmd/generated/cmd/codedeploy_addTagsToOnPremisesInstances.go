package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_addTagsToOnPremisesInstancesCmd = &cobra.Command{
	Use:   "add-tags-to-on-premises-instances",
	Short: "Adds tags to on-premises instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_addTagsToOnPremisesInstancesCmd).Standalone()

	codedeploy_addTagsToOnPremisesInstancesCmd.Flags().String("instance-names", "", "The names of the on-premises instances to which to add tags.")
	codedeploy_addTagsToOnPremisesInstancesCmd.Flags().String("tags", "", "The tag key-value pairs to add to the on-premises instances.")
	codedeploy_addTagsToOnPremisesInstancesCmd.MarkFlagRequired("instance-names")
	codedeploy_addTagsToOnPremisesInstancesCmd.MarkFlagRequired("tags")
	codedeployCmd.AddCommand(codedeploy_addTagsToOnPremisesInstancesCmd)
}
