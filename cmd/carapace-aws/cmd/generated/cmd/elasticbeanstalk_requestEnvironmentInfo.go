package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_requestEnvironmentInfoCmd = &cobra.Command{
	Use:   "request-environment-info",
	Short: "Initiates a request to compile the specified type of information of the deployed environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_requestEnvironmentInfoCmd).Standalone()

	elasticbeanstalk_requestEnvironmentInfoCmd.Flags().String("environment-id", "", "The ID of the environment of the requested data.")
	elasticbeanstalk_requestEnvironmentInfoCmd.Flags().String("environment-name", "", "The name of the environment of the requested data.")
	elasticbeanstalk_requestEnvironmentInfoCmd.Flags().String("info-type", "", "The type of information to request.")
	elasticbeanstalk_requestEnvironmentInfoCmd.MarkFlagRequired("info-type")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_requestEnvironmentInfoCmd)
}
