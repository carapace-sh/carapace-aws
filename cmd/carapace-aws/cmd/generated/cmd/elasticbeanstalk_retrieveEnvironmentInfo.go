package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_retrieveEnvironmentInfoCmd = &cobra.Command{
	Use:   "retrieve-environment-info",
	Short: "Retrieves the compiled information from a [RequestEnvironmentInfo]() request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_retrieveEnvironmentInfoCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_retrieveEnvironmentInfoCmd).Standalone()

		elasticbeanstalk_retrieveEnvironmentInfoCmd.Flags().String("environment-id", "", "The ID of the data's environment.")
		elasticbeanstalk_retrieveEnvironmentInfoCmd.Flags().String("environment-name", "", "The name of the data's environment.")
		elasticbeanstalk_retrieveEnvironmentInfoCmd.Flags().String("info-type", "", "The type of information to retrieve.")
		elasticbeanstalk_retrieveEnvironmentInfoCmd.MarkFlagRequired("info-type")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_retrieveEnvironmentInfoCmd)
}
