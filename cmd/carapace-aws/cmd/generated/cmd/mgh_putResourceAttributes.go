package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgh_putResourceAttributesCmd = &cobra.Command{
	Use:   "put-resource-attributes",
	Short: "Provides identifying details of the resource being migrated so that it can be associated in the Application Discovery Service repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgh_putResourceAttributesCmd).Standalone()

	mgh_putResourceAttributesCmd.Flags().String("dry-run", "", "Optional boolean flag to indicate whether any effect should take place.")
	mgh_putResourceAttributesCmd.Flags().String("migration-task-name", "", "Unique identifier that references the migration task.")
	mgh_putResourceAttributesCmd.Flags().String("progress-update-stream", "", "The name of the ProgressUpdateStream.")
	mgh_putResourceAttributesCmd.Flags().String("resource-attribute-list", "", "Information about the resource that is being migrated.")
	mgh_putResourceAttributesCmd.MarkFlagRequired("migration-task-name")
	mgh_putResourceAttributesCmd.MarkFlagRequired("progress-update-stream")
	mgh_putResourceAttributesCmd.MarkFlagRequired("resource-attribute-list")
	mghCmd.AddCommand(mgh_putResourceAttributesCmd)
}
