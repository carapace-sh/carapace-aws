package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codedeploy_batchGetOnPremisesInstancesCmd = &cobra.Command{
	Use:   "batch-get-on-premises-instances",
	Short: "Gets information about one or more on-premises instances.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codedeploy_batchGetOnPremisesInstancesCmd).Standalone()

	codedeploy_batchGetOnPremisesInstancesCmd.Flags().String("instance-names", "", "The names of the on-premises instances about which to get information.")
	codedeploy_batchGetOnPremisesInstancesCmd.MarkFlagRequired("instance-names")
	codedeployCmd.AddCommand(codedeploy_batchGetOnPremisesInstancesCmd)
}
