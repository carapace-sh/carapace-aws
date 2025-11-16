package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_swapEnvironmentCnamesCmd = &cobra.Command{
	Use:   "swap-environment-cnames",
	Short: "Swaps the CNAMEs of two environments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_swapEnvironmentCnamesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_swapEnvironmentCnamesCmd).Standalone()

		elasticbeanstalk_swapEnvironmentCnamesCmd.Flags().String("destination-environment-id", "", "The ID of the destination environment.")
		elasticbeanstalk_swapEnvironmentCnamesCmd.Flags().String("destination-environment-name", "", "The name of the destination environment.")
		elasticbeanstalk_swapEnvironmentCnamesCmd.Flags().String("source-environment-id", "", "The ID of the source environment.")
		elasticbeanstalk_swapEnvironmentCnamesCmd.Flags().String("source-environment-name", "", "The name of the source environment.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_swapEnvironmentCnamesCmd)
}
