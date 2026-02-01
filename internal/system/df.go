package system

import (
	"bufio"
	"bytes"
	"context"
	"strings"
)

type FSUsage struct {
	Device     string
	FSType     string
	Size       string
	Used       string
	Avail      string
	UsePercent string
	Mountpoint string
}

func ListFSUsage(ctx context.Context, exec Executor) ([]FSUsage, error) {
	out, err := exec.Run(ctx, "df", "-Th")
	if err != nil {
		return nil, err
	}

	var usages []FSUsage
	scanner := bufio.NewScanner(bytes.NewReader(out))

	// skip header
	scanner.Scan()

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) < 7 {
			continue
		}

		usages = append(usages, FSUsage{
			Device:     fields[0],
			FSType:     fields[1],
			Size:       fields[2],
			Used:       fields[3],
			Avail:      fields[4],
			UsePercent: fields[5],
			Mountpoint: fields[6],
		})
	}

	return usages, nil
}
