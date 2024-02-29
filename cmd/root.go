package cmd

import (
	"fmt"
	"mordys/lsgo"
	"os"

	cobra "github.com/spf13/cobra"
)

var rootCommand = &cobra.Command{
	Use:   "mordys",
	Short: "Mordenkainen’s Magnificent Mod Manager for Baldur's Gate 3",
	Long:  `The ultimate BG3 mod mananger, the best in all the realms!`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: If init hasn't been run before, run automatically.
		lsgo.ReadPak("./test_files/BlackDye.pak") // read pak as a default
	},
}

func Execute() {
	if err := rootCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
