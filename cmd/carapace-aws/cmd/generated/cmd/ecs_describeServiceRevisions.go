package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_describeServiceRevisionsCmd = &cobra.Command{
	Use:   "describe-service-revisions",
	Short: "Describes one or more service revisions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_describeServiceRevisionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_describeServiceRevisionsCmd).Standalone()

		ecs_describeServiceRevisionsCmd.Flags().String("service-revision-arns", "", "The ARN of the service revision.")
		ecs_describeServiceRevisionsCmd.MarkFlagRequired("service-revision-arns")
	})
	ecsCmd.AddCommand(ecs_describeServiceRevisionsCmd)
}
