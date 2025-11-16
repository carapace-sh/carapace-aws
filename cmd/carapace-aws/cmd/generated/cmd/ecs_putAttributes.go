package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecs_putAttributesCmd = &cobra.Command{
	Use:   "put-attributes",
	Short: "Create or update an attribute on an Amazon ECS resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecs_putAttributesCmd).Standalone()

	ecs_putAttributesCmd.Flags().String("attributes", "", "The attributes to apply to your resource.")
	ecs_putAttributesCmd.Flags().String("cluster", "", "The short name or full Amazon Resource Name (ARN) of the cluster that contains the resource to apply attributes.")
	ecs_putAttributesCmd.MarkFlagRequired("attributes")
	ecsCmd.AddCommand(ecs_putAttributesCmd)
}
