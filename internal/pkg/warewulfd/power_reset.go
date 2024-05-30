package power

import (
    "log"
    "os/exec"
)

// Reset function to reset the node
func Reset(node string) error {
    cmd := exec.Command("ipmitool", "-I", "lanplus", "-H", node, "-U", "admin", "-P", "password", "power", "reset")
    output, err := cmd.CombinedOutput()
    if err != nil {
        log.Printf("Error resetting node %s: %v\nOutput: %s", node, err, string(output))
        return err
    }
    log.Printf("Node %s has been reset successfully\nOutput: %s", node, string(output))
    return nil
}
