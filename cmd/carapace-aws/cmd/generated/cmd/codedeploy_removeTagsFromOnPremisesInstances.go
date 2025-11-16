package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_removeTagsFromOnPremisesInstancesCmd = &cobra.Command{
	Use:   "remove-tags-from-on-premises-instances",
	Short: "Removes one or more tags from one or more on-premises instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_removeTagsFromOnPremisesInstancesCmd).Standalone()

	codedeploy_removeTagsFromOnPremisesInstancesCmd.Flags().String("instance-names", "", "The names of the on-premises instances from which to remove tags.")
	codedeploy_removeTagsFromOnPremisesInstancesCmd.Flags().String("tags", "", "The tag key-value pairs to remove from the on-premises instances.")
	codedeploy_removeTagsFromOnPremisesInstancesCmd.MarkFlagRequired("instance-names")
	codedeploy_removeTagsFromOnPremisesInstancesCmd.MarkFlagRequired("tags")
	codedeployCmd.AddCommand(codedeploy_removeTagsFromOnPremisesInstancesCmd)
}
