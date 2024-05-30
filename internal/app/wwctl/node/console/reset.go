package console

import (
    "github.com/spf13/cobra"
    "github.com/warewulf/warewulf/internal/pkg/warewulfd/power"
    "log"
)

// ResetCmd represents the reset command
var ResetCmd = &cobra.Command{
    Use:   "reset",
    Short: "Reset the node console",
    Long:  `This command resets the node console.`,
    Run: func(cmd *cobra.Command, args []string) {
        // Implement the reset functionality here
        if len(args) < 1 {
            log.Fatalf("You must specify the node(s) to reset")
        }

        for _, node := range args {
            err := power.Reset(node)
            if err != nil {
                log.Fatalf("Failed to reset node %s: %v", node, err)
            }
            log.Printf("Node %s has been reset successfully", node)
        }
    },
}

func init() {
    ConsoleCmd.AddCommand(ResetCmd)
}