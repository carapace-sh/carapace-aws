package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptunedata_listMlendpointsCmd = &cobra.Command{
	Use:   "list-mlendpoints",
	Short: "Lists existing inference endpoints.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptunedata_listMlendpointsCmd).Standalone()

	neptunedata_listMlendpointsCmd.Flags().String("max-items", "", "The maximum number of items to return (from 1 to 1024; the default is 10.")
	neptunedata_listMlendpointsCmd.Flags().String("neptune-iam-role-arn", "", "The ARN of an IAM role that provides Neptune access to SageMaker and Amazon S3 resources.")
	neptunedataCmd.AddCommand(neptunedata_listMlendpointsCmd)
}
