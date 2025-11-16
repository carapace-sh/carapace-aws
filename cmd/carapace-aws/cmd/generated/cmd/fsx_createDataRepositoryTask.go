package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fsx_createDataRepositoryTaskCmd = &cobra.Command{
	Use:   "create-data-repository-task",
	Short: "Creates an Amazon FSx for Lustre data repository task.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fsx_createDataRepositoryTaskCmd).Standalone()

	fsx_createDataRepositoryTaskCmd.Flags().String("capacity-to-release", "", "Specifies the amount of data to release, in GiB, by an Amazon File Cache `AUTO_RELEASE_DATA` task that automatically releases files from the cache.")
	fsx_createDataRepositoryTaskCmd.Flags().String("client-request-token", "", "")
	fsx_createDataRepositoryTaskCmd.Flags().String("file-system-id", "", "")
	fsx_createDataRepositoryTaskCmd.Flags().String("paths", "", "A list of paths for the data repository task to use when the task is processed.")
	fsx_createDataRepositoryTaskCmd.Flags().String("release-configuration", "", "The configuration that specifies the last accessed time criteria for files that will be released from an Amazon FSx for Lustre file system.")
	fsx_createDataRepositoryTaskCmd.Flags().String("report", "", "Defines whether or not Amazon FSx provides a CompletionReport once the task has completed.")
	fsx_createDataRepositoryTaskCmd.Flags().String("tags", "", "")
	fsx_createDataRepositoryTaskCmd.Flags().String("type", "", "Specifies the type of data repository task to create.")
	fsx_createDataRepositoryTaskCmd.MarkFlagRequired("file-system-id")
	fsx_createDataRepositoryTaskCmd.MarkFlagRequired("report")
	fsx_createDataRepositoryTaskCmd.MarkFlagRequired("type")
	fsxCmd.AddCommand(fsx_createDataRepositoryTaskCmd)
}
