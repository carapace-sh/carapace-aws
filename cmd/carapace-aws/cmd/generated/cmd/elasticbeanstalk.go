package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalkCmd = &cobra.Command{
	Use:   "elasticbeanstalk",
	Short: "AWS Elastic Beanstalk\n\nAWS Elastic Beanstalk makes it easy for you to create, deploy, and manage scalable, fault-tolerant applications running on the Amazon Web Services cloud.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalkCmd).Standalone()

	rootCmd.AddCommand(elasticbeanstalkCmd)
}
