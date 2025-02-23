package lib

import (
	"fmt"
	"os"
	"time"
)

func BufferToFile(buffer []byte) error {
	err := os.WriteFile(fmt.Sprintf("screenshot_%d.png", time.Now().UnixMilli()), buffer, 0644)
	if err != nil {
		return err
	}

	fmt.Println("- Screenshot Saved Successfully ✅")
	return nil
}
