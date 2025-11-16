package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_describeEnvironmentHealthCmd = &cobra.Command{
	Use:   "describe-environment-health",
	Short: "Returns information about the overall health of the specified environment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_describeEnvironmentHealthCmd).Standalone()

	elasticbeanstalk_describeEnvironmentHealthCmd.Flags().String("attribute-names", "", "Specify the response elements to return.")
	elasticbeanstalk_describeEnvironmentHealthCmd.Flags().String("environment-id", "", "Specify the environment by ID.")
	elasticbeanstalk_describeEnvironmentHealthCmd.Flags().String("environment-name", "", "Specify the environment by name.")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_describeEnvironmentHealthCmd)
}
