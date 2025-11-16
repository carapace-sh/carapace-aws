package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_deleteAttributesCmd = &cobra.Command{
	Use:   "delete-attributes",
	Short: "Deletes one or more custom attributes from an Amazon ECS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_deleteAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecs_deleteAttributesCmd).Standalone()

		ecs_deleteAttributesCmd.Flags().String("attributes", "", "The attributes to delete from your resource.")
		ecs_deleteAttributesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that contains the resource to delete attributes.")
		ecs_deleteAttributesCmd.MarkFlagRequired("attributes")
	})
	ecsCmd.AddCommand(ecs_deleteAttributesCmd)
}
