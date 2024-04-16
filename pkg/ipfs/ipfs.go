package ipfs

import (
	"bytes"
	"ipfs-file-api/internal/config"

	shell "github.com/ipfs/go-ipfs-api"
)

func UploadBytes(buf *bytes.Buffer) (string, error) {
	sh := shell.NewShell(config.IPFSHost)

	cid, err := sh.Add(buf)
	if err != nil {
		return "", err
	}

	return cid, nil
}

func DownloadBytes(cid string) ([]byte, error) {
	sh := shell.NewShell(config.IPFSHost)

	data, err := sh.Cat(cid)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
