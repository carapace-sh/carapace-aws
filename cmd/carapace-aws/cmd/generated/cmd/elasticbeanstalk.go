package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalkCmd = &cobra.Command{
	Use:   "elasticbeanstalk",
	Short: "AWS Elastic Beanstalk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalkCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalkCmd).Standalone()

	})
	rootCmd.AddCommand(elasticbeanstalkCmd)
}
