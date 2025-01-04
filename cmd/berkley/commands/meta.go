package commands

import (
	"encoding/binary"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Metadata struct {
	LSN         uint64
	PageNo      uint32
	Magic       uint32
	Version     uint32
	PageSize    uint32
	Ec          uint8
	Ty          uint8
	Mf          uint8
	Free        uint32
	LastPageNo  uint32
	NParts      uint32
	KeyCount    uint32
	RecordCount uint32
	Flags       []byte // unit32
	UID         []byte // 20 bytes
	Minkey      uint32
	ReLen       uint32
	RePad       uint32
	Root        uint32
	CryptoMagic uint32
	Iv          []byte // unit128
	ChkSum      []byte
}

func GetMetaCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "meta [file]",
		Short: "Read metadata from Berkeley DB file",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			filePath := args[0]

			file, err := os.Open(filePath)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				return
			}
			defer file.Close()

			// Read first 512 bytes
			header := make([]byte, 512)
			_, err = file.Read(header)
			if err != nil {
				fmt.Printf("Error reading file: %v\n", err)
				return
			}

			metadata := &Metadata{}

			lsn := header[0:8]
			pageNo := header[8:12]
			magic := header[12:16]

			metadata.LSN = binary.LittleEndian.Uint64(lsn)
			metadata.PageNo = binary.LittleEndian.Uint32(pageNo)
			metadata.Magic = binary.LittleEndian.Uint32(magic)

			fmt.Printf("Metadata: %v\n", metadata)
		},
	}

	return cmd
}
